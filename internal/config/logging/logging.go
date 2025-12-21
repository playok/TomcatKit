package logging

// LogLevel represents Java logging levels
type LogLevel string

const (
	LogLevelOff     LogLevel = "OFF"
	LogLevelSevere  LogLevel = "SEVERE"
	LogLevelWarning LogLevel = "WARNING"
	LogLevelInfo    LogLevel = "INFO"
	LogLevelConfig  LogLevel = "CONFIG"
	LogLevelFine    LogLevel = "FINE"
	LogLevelFiner   LogLevel = "FINER"
	LogLevelFinest  LogLevel = "FINEST"
	LogLevelAll     LogLevel = "ALL"
)

// AvailableLogLevels returns all available log levels in order
func AvailableLogLevels() []LogLevel {
	return []LogLevel{
		LogLevelOff,
		LogLevelSevere,
		LogLevelWarning,
		LogLevelInfo,
		LogLevelConfig,
		LogLevelFine,
		LogLevelFiner,
		LogLevelFinest,
		LogLevelAll,
	}
}

// GetLogLevelDescription returns a description for the log level
func GetLogLevelDescription(level LogLevel) string {
	switch level {
	case LogLevelOff:
		return "Logging disabled"
	case LogLevelSevere:
		return "Serious failures only"
	case LogLevelWarning:
		return "Potential problems"
	case LogLevelInfo:
		return "Informational messages"
	case LogLevelConfig:
		return "Configuration messages"
	case LogLevelFine:
		return "Tracing information"
	case LogLevelFiner:
		return "Detailed tracing"
	case LogLevelFinest:
		return "Most detailed tracing"
	case LogLevelAll:
		return "Log all messages"
	default:
		return string(level)
	}
}

// Handler class names
const (
	HandlerAsyncFileHandler = "org.apache.juli.AsyncFileHandler"
	HandlerFileHandler      = "org.apache.juli.FileHandler"
	HandlerConsoleHandler   = "java.util.logging.ConsoleHandler"
)

// Formatter class names
const (
	FormatterOneLineFormatter = "org.apache.juli.OneLineFormatter"
	FormatterVerboseFormatter = "org.apache.juli.VerboseFormatter"
	FormatterSimpleFormatter  = "java.util.logging.SimpleFormatter"
)

// Common logger names
var CommonLoggers = []string{
	"org.apache.catalina.core.ContainerBase.[Catalina].[localhost]",
	"org.apache.catalina.core.ContainerBase.[Catalina].[localhost].[/manager]",
	"org.apache.catalina.core.ContainerBase.[Catalina].[localhost].[/host-manager]",
	"org.apache.catalina.session",
	"org.apache.catalina.startup",
	"org.apache.catalina.connector",
	"org.apache.coyote",
	"org.apache.jasper",
	"org.apache.tomcat.util",
}

// LoggingConfig represents the complete logging configuration
type LoggingConfig struct {
	Handlers       []string
	RootHandlers   []string
	FileHandlers   []FileHandler
	ConsoleHandler *ConsoleHandler
	Loggers        []Logger
}

// FileHandler represents an async or sync file handler
type FileHandler struct {
	Prefix     string // e.g., "1catalina", "2localhost"
	ClassName  string // AsyncFileHandler or FileHandler
	Level      LogLevel
	Directory  string
	FilePrefix string
	Suffix     string
	MaxDays    int
	Encoding   string
	BufferSize int
	Formatter  string
	Rotatable  bool
}

// GetHandlerName returns the full handler name (prefix + className)
func (h *FileHandler) GetHandlerName() string {
	return h.Prefix + "." + h.ClassName
}

// IsAsync returns true if this is an async handler
func (h *FileHandler) IsAsync() bool {
	return h.ClassName == HandlerAsyncFileHandler
}

// NewAsyncFileHandler creates an AsyncFileHandler with defaults
func NewAsyncFileHandler(prefix, filePrefix string) *FileHandler {
	return &FileHandler{
		Prefix:     prefix,
		ClassName:  HandlerAsyncFileHandler,
		Level:      LogLevelAll,
		Directory:  "${catalina.base}/logs",
		FilePrefix: filePrefix,
		Suffix:     ".log",
		MaxDays:    90,
		Encoding:   "UTF-8",
		BufferSize: 16384,
		Rotatable:  true,
	}
}

// NewFileHandler creates a FileHandler with defaults
func NewFileHandler(prefix, filePrefix string) *FileHandler {
	return &FileHandler{
		Prefix:     prefix,
		ClassName:  HandlerFileHandler,
		Level:      LogLevelAll,
		Directory:  "${catalina.base}/logs",
		FilePrefix: filePrefix,
		Suffix:     ".log",
		MaxDays:    90,
		Encoding:   "UTF-8",
		Rotatable:  true,
	}
}

// ConsoleHandler represents console logging handler
type ConsoleHandler struct {
	Level     LogLevel
	Formatter string
	Encoding  string
}

// NewConsoleHandler creates a ConsoleHandler with defaults
func NewConsoleHandler() *ConsoleHandler {
	return &ConsoleHandler{
		Level:     LogLevelAll,
		Formatter: "org.apache.juli.OneLineFormatter",
		Encoding:  "UTF-8",
	}
}

// Logger represents a logger configuration
type Logger struct {
	Name              string
	Level             LogLevel
	Handlers          []string
	UseParentHandlers bool
}

// AccessLogConfig represents access log valve configuration
type AccessLogConfig struct {
	Enabled        bool
	ClassName      string
	Directory      string
	Prefix         string
	Suffix         string
	Pattern        string
	Rotatable      bool
	FileDateFormat string
	Buffered       bool
	Encoding       string
	MaxDays        int
	// Extended options
	ResolveHosts             bool
	RequestAttributesEnabled bool
	ConditionUnless          string
	ConditionIf              string
}

// NewAccessLogConfig creates an AccessLogConfig with defaults
func NewAccessLogConfig() *AccessLogConfig {
	return &AccessLogConfig{
		Enabled:        true,
		ClassName:      "org.apache.catalina.valves.AccessLogValve",
		Directory:      "logs",
		Prefix:         "localhost_access_log",
		Suffix:         ".txt",
		Pattern:        "combined",
		Rotatable:      true,
		FileDateFormat: ".yyyy-MM-dd",
		Buffered:       true,
		Encoding:       "UTF-8",
		MaxDays:        -1,
	}
}

// Common access log patterns
const (
	PatternCommon   = "common"
	PatternCombined = "combined"
	PatternCustom   = "%h %l %u %t \"%r\" %s %b %D"
)

// DefaultLoggingConfig returns the default Tomcat logging configuration
func DefaultLoggingConfig() *LoggingConfig {
	return &LoggingConfig{
		Handlers: []string{
			"1catalina.org.apache.juli.AsyncFileHandler",
			"2localhost.org.apache.juli.AsyncFileHandler",
			"java.util.logging.ConsoleHandler",
		},
		RootHandlers: []string{
			"1catalina.org.apache.juli.AsyncFileHandler",
			"java.util.logging.ConsoleHandler",
		},
		FileHandlers: []FileHandler{
			*NewAsyncFileHandler("1catalina", "catalina."),
			*NewAsyncFileHandler("2localhost", "localhost."),
		},
		ConsoleHandler: NewConsoleHandler(),
		Loggers: []Logger{
			{
				Name:     "org.apache.catalina.core.ContainerBase.[Catalina].[localhost]",
				Level:    LogLevelInfo,
				Handlers: []string{"2localhost.org.apache.juli.AsyncFileHandler"},
			},
		},
	}
}
