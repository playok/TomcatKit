package logging

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ConfigService provides operations for logging.properties
type ConfigService struct {
	catalinaBase string
	configPath   string
	config       *LoggingConfig
	rawLines     []string // Preserve original lines for comments
}

// NewConfigService creates a new logging configuration service
func NewConfigService(catalinaBase string) *ConfigService {
	return &ConfigService{
		catalinaBase: catalinaBase,
		configPath:   filepath.Join(catalinaBase, "conf", "logging.properties"),
	}
}

// GetConfigPath returns the path to logging.properties
func (s *ConfigService) GetConfigPath() string {
	return s.configPath
}

// Load reads and parses the logging.properties file
func (s *ConfigService) Load() error {
	file, err := os.Open(s.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default config if file doesn't exist
			s.config = DefaultLoggingConfig()
			return nil
		}
		return fmt.Errorf("failed to open logging.properties: %w", err)
	}
	defer file.Close()

	s.config = &LoggingConfig{
		FileHandlers:   []FileHandler{},
		ConsoleHandler: NewConsoleHandler(),
		Loggers:        []Logger{},
	}
	s.rawLines = []string{}

	// Read all lines
	scanner := bufio.NewScanner(file)
	properties := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		s.rawLines = append(s.rawLines, line)

		// Skip comments and empty lines
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Handle continuation lines
		for strings.HasSuffix(trimmed, "\\") {
			if scanner.Scan() {
				nextLine := scanner.Text()
				s.rawLines = append(s.rawLines, nextLine)
				trimmed = strings.TrimSuffix(trimmed, "\\") + strings.TrimSpace(nextLine)
			}
		}

		// Parse property
		if idx := strings.Index(trimmed, "="); idx > 0 {
			key := strings.TrimSpace(trimmed[:idx])
			value := strings.TrimSpace(trimmed[idx+1:])
			properties[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading logging.properties: %w", err)
	}

	// Parse handlers list
	if handlers, ok := properties["handlers"]; ok {
		s.config.Handlers = parseHandlerList(handlers)
	}

	// Parse root logger handlers
	if rootHandlers, ok := properties[".handlers"]; ok {
		s.config.RootHandlers = parseHandlerList(rootHandlers)
	}

	// Parse file handlers
	s.parseFileHandlers(properties)

	// Parse console handler
	s.parseConsoleHandler(properties)

	// Parse loggers
	s.parseLoggers(properties)

	return nil
}

// parseHandlerList splits a comma-separated handler list
func parseHandlerList(handlers string) []string {
	var result []string
	for _, h := range strings.Split(handlers, ",") {
		h = strings.TrimSpace(h)
		if h != "" {
			result = append(result, h)
		}
	}
	return result
}

// parseFileHandlers extracts file handler configurations
func (s *ConfigService) parseFileHandlers(props map[string]string) {
	// Find all file handlers by looking for .level properties with prefixes
	handlerPrefixes := make(map[string]bool)

	for key := range props {
		// Match patterns like "1catalina.org.apache.juli.AsyncFileHandler.level"
		if strings.Contains(key, HandlerAsyncFileHandler) || strings.Contains(key, HandlerFileHandler) {
			// Extract prefix (e.g., "1catalina")
			if idx := strings.Index(key, ".org.apache.juli"); idx > 0 {
				prefix := key[:idx]
				handlerPrefixes[prefix] = true
			}
		}
	}

	for prefix := range handlerPrefixes {
		handler := s.parseFileHandler(prefix, props)
		if handler != nil {
			s.config.FileHandlers = append(s.config.FileHandlers, *handler)
		}
	}
}

// parseFileHandler parses a single file handler configuration
func (s *ConfigService) parseFileHandler(prefix string, props map[string]string) *FileHandler {
	// Determine handler class
	className := HandlerAsyncFileHandler
	baseKey := prefix + "." + HandlerAsyncFileHandler
	if _, ok := props[baseKey+".level"]; !ok {
		className = HandlerFileHandler
		baseKey = prefix + "." + HandlerFileHandler
		if _, ok := props[baseKey+".level"]; !ok {
			return nil
		}
	}

	handler := &FileHandler{
		Prefix:    prefix,
		ClassName: className,
		Level:     LogLevelAll,
		Directory: "${catalina.base}/logs",
		Suffix:    ".log",
		MaxDays:   90,
		Encoding:  "UTF-8",
		Rotatable: true,
	}

	// Parse properties
	if level, ok := props[baseKey+".level"]; ok {
		handler.Level = LogLevel(level)
	}
	if dir, ok := props[baseKey+".directory"]; ok {
		handler.Directory = dir
	}
	if prefix, ok := props[baseKey+".prefix"]; ok {
		handler.FilePrefix = prefix
	}
	if suffix, ok := props[baseKey+".suffix"]; ok {
		handler.Suffix = suffix
	}
	if maxDays, ok := props[baseKey+".maxDays"]; ok {
		if days, err := strconv.Atoi(maxDays); err == nil {
			handler.MaxDays = days
		}
	}
	if encoding, ok := props[baseKey+".encoding"]; ok {
		handler.Encoding = encoding
	}
	if bufSize, ok := props[baseKey+".bufferSize"]; ok {
		if size, err := strconv.Atoi(bufSize); err == nil {
			handler.BufferSize = size
		}
	}
	if formatter, ok := props[baseKey+".formatter"]; ok {
		handler.Formatter = formatter
	}
	if rotatable, ok := props[baseKey+".rotatable"]; ok {
		handler.Rotatable = strings.ToLower(rotatable) == "true"
	}

	return handler
}

// parseConsoleHandler parses console handler configuration
func (s *ConfigService) parseConsoleHandler(props map[string]string) {
	baseKey := HandlerConsoleHandler

	if level, ok := props[baseKey+".level"]; ok {
		s.config.ConsoleHandler.Level = LogLevel(level)
	}
	if formatter, ok := props[baseKey+".formatter"]; ok {
		s.config.ConsoleHandler.Formatter = formatter
	}
	if encoding, ok := props[baseKey+".encoding"]; ok {
		s.config.ConsoleHandler.Encoding = encoding
	}
}

// parseLoggers extracts logger configurations
func (s *ConfigService) parseLoggers(props map[string]string) {
	// Find all logger entries by looking for .level or .handlers
	loggerNames := make(map[string]bool)

	levelRegex := regexp.MustCompile(`^([a-zA-Z0-9._\[\]/]+)\.level$`)
	handlersRegex := regexp.MustCompile(`^([a-zA-Z0-9._\[\]/]+)\.handlers$`)

	for key := range props {
		// Skip handler properties
		if strings.Contains(key, "org.apache.juli") || strings.Contains(key, "java.util.logging") {
			continue
		}

		if matches := levelRegex.FindStringSubmatch(key); len(matches) > 1 {
			name := matches[1]
			// Skip root logger (empty name represented by just ".handlers" or ".level")
			if name != "" && name != "." {
				loggerNames[name] = true
			}
		}
		if matches := handlersRegex.FindStringSubmatch(key); len(matches) > 1 {
			name := matches[1]
			if name != "" && name != "." {
				loggerNames[name] = true
			}
		}
	}

	for name := range loggerNames {
		logger := Logger{
			Name:              name,
			Level:             LogLevelInfo,
			UseParentHandlers: true,
		}

		if level, ok := props[name+".level"]; ok {
			logger.Level = LogLevel(level)
		}
		if handlers, ok := props[name+".handlers"]; ok {
			logger.Handlers = parseHandlerList(handlers)
		}
		if useParent, ok := props[name+".useParentHandlers"]; ok {
			logger.UseParentHandlers = strings.ToLower(useParent) != "false"
		}

		s.config.Loggers = append(s.config.Loggers, logger)
	}
}

// GetConfig returns the current logging configuration
func (s *ConfigService) GetConfig() *LoggingConfig {
	return s.config
}

// Save writes the logging configuration to file
func (s *ConfigService) Save() error {
	// Create backup
	if err := s.createBackup(); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	// Generate content
	content := s.generateContent()

	// Write file
	if err := os.WriteFile(s.configPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write logging.properties: %w", err)
	}

	return nil
}

// createBackup creates a backup of the current logging.properties
func (s *ConfigService) createBackup() error {
	if _, err := os.Stat(s.configPath); os.IsNotExist(err) {
		return nil // No file to backup
	}

	backupDir := filepath.Join(s.catalinaBase, "conf", "backup")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(backupDir, fmt.Sprintf("logging.properties.%s", timestamp))

	data, err := os.ReadFile(s.configPath)
	if err != nil {
		return err
	}

	return os.WriteFile(backupPath, data, 0644)
}

// generateContent creates the logging.properties content
func (s *ConfigService) generateContent() string {
	var sb strings.Builder

	// Header
	sb.WriteString("# Tomcat JULI Logging Configuration\n")
	sb.WriteString("# Generated by TomcatKit\n")
	sb.WriteString("# " + time.Now().Format("2006-01-02 15:04:05") + "\n")
	sb.WriteString("\n")

	// Handlers list
	sb.WriteString("# Registered handlers\n")
	sb.WriteString("handlers = ")
	sb.WriteString(strings.Join(s.config.Handlers, ", "))
	sb.WriteString("\n\n")

	// Root logger handlers
	sb.WriteString("# Root logger handlers\n")
	sb.WriteString(".handlers = ")
	sb.WriteString(strings.Join(s.config.RootHandlers, ", "))
	sb.WriteString("\n\n")

	// File handlers
	sb.WriteString("############################################################\n")
	sb.WriteString("# File Handler Configuration\n")
	sb.WriteString("############################################################\n\n")

	for _, handler := range s.config.FileHandlers {
		sb.WriteString(s.generateFileHandlerConfig(&handler))
		sb.WriteString("\n")
	}

	// Console handler
	sb.WriteString("############################################################\n")
	sb.WriteString("# Console Handler Configuration\n")
	sb.WriteString("############################################################\n\n")
	sb.WriteString(s.generateConsoleHandlerConfig())
	sb.WriteString("\n")

	// Loggers
	if len(s.config.Loggers) > 0 {
		sb.WriteString("############################################################\n")
		sb.WriteString("# Logger Configuration\n")
		sb.WriteString("############################################################\n\n")

		for _, logger := range s.config.Loggers {
			sb.WriteString(s.generateLoggerConfig(&logger))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// generateFileHandlerConfig generates config for a file handler
func (s *ConfigService) generateFileHandlerConfig(handler *FileHandler) string {
	var sb strings.Builder
	baseKey := handler.GetHandlerName()

	sb.WriteString(fmt.Sprintf("# %s file handler\n", handler.Prefix))
	sb.WriteString(fmt.Sprintf("%s.level = %s\n", baseKey, handler.Level))
	sb.WriteString(fmt.Sprintf("%s.directory = %s\n", baseKey, handler.Directory))
	sb.WriteString(fmt.Sprintf("%s.prefix = %s\n", baseKey, handler.FilePrefix))
	sb.WriteString(fmt.Sprintf("%s.suffix = %s\n", baseKey, handler.Suffix))
	sb.WriteString(fmt.Sprintf("%s.maxDays = %d\n", baseKey, handler.MaxDays))
	sb.WriteString(fmt.Sprintf("%s.encoding = %s\n", baseKey, handler.Encoding))

	if handler.IsAsync() && handler.BufferSize > 0 {
		sb.WriteString(fmt.Sprintf("%s.bufferSize = %d\n", baseKey, handler.BufferSize))
	}
	if handler.Formatter != "" {
		sb.WriteString(fmt.Sprintf("%s.formatter = %s\n", baseKey, handler.Formatter))
	}

	return sb.String()
}

// generateConsoleHandlerConfig generates config for console handler
func (s *ConfigService) generateConsoleHandlerConfig() string {
	var sb strings.Builder
	baseKey := HandlerConsoleHandler

	sb.WriteString(fmt.Sprintf("%s.level = %s\n", baseKey, s.config.ConsoleHandler.Level))
	sb.WriteString(fmt.Sprintf("%s.formatter = %s\n", baseKey, s.config.ConsoleHandler.Formatter))
	if s.config.ConsoleHandler.Encoding != "" {
		sb.WriteString(fmt.Sprintf("%s.encoding = %s\n", baseKey, s.config.ConsoleHandler.Encoding))
	}

	return sb.String()
}

// generateLoggerConfig generates config for a logger
func (s *ConfigService) generateLoggerConfig(logger *Logger) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n", logger.Name))
	sb.WriteString(fmt.Sprintf("%s.level = %s\n", logger.Name, logger.Level))
	if len(logger.Handlers) > 0 {
		sb.WriteString(fmt.Sprintf("%s.handlers = %s\n", logger.Name, strings.Join(logger.Handlers, ", ")))
	}
	if !logger.UseParentHandlers {
		sb.WriteString(fmt.Sprintf("%s.useParentHandlers = false\n", logger.Name))
	}

	return sb.String()
}

// AddFileHandler adds a new file handler
func (s *ConfigService) AddFileHandler(handler *FileHandler) {
	s.config.FileHandlers = append(s.config.FileHandlers, *handler)

	// Add to handlers list if not present
	handlerName := handler.GetHandlerName()
	found := false
	for _, h := range s.config.Handlers {
		if h == handlerName {
			found = true
			break
		}
	}
	if !found {
		s.config.Handlers = append(s.config.Handlers, handlerName)
	}
}

// RemoveFileHandler removes a file handler by prefix
func (s *ConfigService) RemoveFileHandler(prefix string) {
	var newHandlers []FileHandler
	var removedHandlerName string

	for _, h := range s.config.FileHandlers {
		if h.Prefix == prefix {
			removedHandlerName = h.GetHandlerName()
		} else {
			newHandlers = append(newHandlers, h)
		}
	}
	s.config.FileHandlers = newHandlers

	// Remove from handlers list
	if removedHandlerName != "" {
		var newList []string
		for _, h := range s.config.Handlers {
			if h != removedHandlerName {
				newList = append(newList, h)
			}
		}
		s.config.Handlers = newList

		// Also remove from root handlers
		var newRootList []string
		for _, h := range s.config.RootHandlers {
			if h != removedHandlerName {
				newRootList = append(newRootList, h)
			}
		}
		s.config.RootHandlers = newRootList
	}
}

// AddLogger adds a new logger
func (s *ConfigService) AddLogger(logger *Logger) {
	s.config.Loggers = append(s.config.Loggers, *logger)
}

// RemoveLogger removes a logger by name
func (s *ConfigService) RemoveLogger(name string) {
	var newLoggers []Logger
	for _, l := range s.config.Loggers {
		if l.Name != name {
			newLoggers = append(newLoggers, l)
		}
	}
	s.config.Loggers = newLoggers
}

// UpdateLogger updates an existing logger
func (s *ConfigService) UpdateLogger(name string, updated *Logger) {
	for i, l := range s.config.Loggers {
		if l.Name == name {
			s.config.Loggers[i] = *updated
			return
		}
	}
}

// GetFileHandler returns a file handler by prefix
func (s *ConfigService) GetFileHandler(prefix string) *FileHandler {
	for i := range s.config.FileHandlers {
		if s.config.FileHandlers[i].Prefix == prefix {
			return &s.config.FileHandlers[i]
		}
	}
	return nil
}

// GetLogger returns a logger by name
func (s *ConfigService) GetLogger(name string) *Logger {
	for i := range s.config.Loggers {
		if s.config.Loggers[i].Name == name {
			return &s.config.Loggers[i]
		}
	}
	return nil
}

// SetRootHandlers sets the root logger handlers
func (s *ConfigService) SetRootHandlers(handlers []string) {
	s.config.RootHandlers = handlers
}

// SetConsoleHandler updates the console handler configuration
func (s *ConfigService) SetConsoleHandler(handler *ConsoleHandler) {
	s.config.ConsoleHandler = handler
}
