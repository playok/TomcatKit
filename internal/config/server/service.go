package server

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigService handles server.xml configuration operations
type ConfigService struct {
	catalinaBase string
	server       *Server
	filePath     string
}

// NewConfigService creates a new server configuration service
func NewConfigService(catalinaBase string) *ConfigService {
	return &ConfigService{
		catalinaBase: catalinaBase,
		filePath:     filepath.Join(catalinaBase, "conf", "server.xml"),
	}
}

// Load reads and parses server.xml
func (s *ConfigService) Load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to read server.xml: %w", err)
	}

	var server Server
	if err := xml.Unmarshal(data, &server); err != nil {
		return fmt.Errorf("failed to parse server.xml: %w", err)
	}

	s.server = &server
	return nil
}

// Save writes the configuration back to server.xml
func (s *ConfigService) Save() error {
	if s.server == nil {
		return fmt.Errorf("no server configuration loaded")
	}

	// Create backup first
	if err := s.createBackup(); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	data, err := xml.MarshalIndent(s.server, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal server.xml: %w", err)
	}

	// Add XML declaration
	output := []byte(xml.Header)
	output = append(output, data...)

	if err := os.WriteFile(s.filePath, output, 0644); err != nil {
		return fmt.Errorf("failed to write server.xml: %w", err)
	}

	return nil
}

// createBackup creates a backup of the current server.xml
func (s *ConfigService) createBackup() error {
	backupDir := filepath.Join(s.catalinaBase, "conf", "backup")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file to backup
		}
		return err
	}

	backupPath := filepath.Join(backupDir, "server.xml.bak")
	return os.WriteFile(backupPath, data, 0644)
}

// GetServer returns the current server configuration
func (s *ConfigService) GetServer() *Server {
	return s.server
}

// GetConfig returns the current server configuration (alias for GetServer)
func (s *ConfigService) GetConfig() *Server {
	return s.server
}

// GetFilePath returns the server.xml file path
func (s *ConfigService) GetFilePath() string {
	return s.filePath
}

// UpdateServerPort updates the shutdown port
func (s *ConfigService) UpdateServerPort(port int) {
	if s.server != nil {
		s.server.Port = port
	}
}

// UpdateShutdownCommand updates the shutdown command
func (s *ConfigService) UpdateShutdownCommand(cmd string) {
	if s.server != nil {
		s.server.Shutdown = cmd
	}
}

// GetListeners returns all listeners
func (s *ConfigService) GetListeners() []Listener {
	if s.server != nil {
		return s.server.Listeners
	}
	return nil
}

// AddListener adds a new listener
func (s *ConfigService) AddListener(listener Listener) {
	if s.server != nil {
		s.server.Listeners = append(s.server.Listeners, listener)
	}
}

// RemoveListener removes a listener by index
func (s *ConfigService) RemoveListener(index int) {
	if s.server != nil && index >= 0 && index < len(s.server.Listeners) {
		s.server.Listeners = append(s.server.Listeners[:index], s.server.Listeners[index+1:]...)
	}
}

// GetServices returns all services
func (s *ConfigService) GetServices() []Service {
	if s.server != nil {
		return s.server.Services
	}
	return nil
}

// GetService returns a service by index
func (s *ConfigService) GetService(index int) *Service {
	if s.server != nil && index >= 0 && index < len(s.server.Services) {
		return &s.server.Services[index]
	}
	return nil
}

// UpdateService updates a service at the given index
func (s *ConfigService) UpdateService(index int, service Service) {
	if s.server != nil && index >= 0 && index < len(s.server.Services) {
		s.server.Services[index] = service
	}
}

// GetGlobalResources returns global naming resources
func (s *ConfigService) GetGlobalResources() *GlobalNamingResources {
	if s.server != nil {
		return s.server.Resources
	}
	return nil
}

// SetGlobalResources sets global naming resources
func (s *ConfigService) SetGlobalResources(resources *GlobalNamingResources) {
	if s.server != nil {
		s.server.Resources = resources
	}
}

// AddGlobalResource adds a global resource
func (s *ConfigService) AddGlobalResource(resource Resource) {
	if s.server != nil {
		if s.server.Resources == nil {
			s.server.Resources = &GlobalNamingResources{}
		}
		s.server.Resources.Resources = append(s.server.Resources.Resources, resource)
	}
}

// RemoveGlobalResource removes a global resource by index
func (s *ConfigService) RemoveGlobalResource(index int) {
	if s.server != nil && s.server.Resources != nil {
		resources := s.server.Resources.Resources
		if index >= 0 && index < len(resources) {
			s.server.Resources.Resources = append(resources[:index], resources[index+1:]...)
		}
	}
}

// Common Listener class names
const (
	ListenerVersionLogger         = "org.apache.catalina.startup.VersionLoggerListener"
	ListenerAprLifecycle          = "org.apache.catalina.core.AprLifecycleListener"
	ListenerJreMemoryLeak         = "org.apache.catalina.core.JreMemoryLeakPreventionListener"
	ListenerGlobalResources       = "org.apache.catalina.mbeans.GlobalResourcesLifecycleListener"
	ListenerThreadLocalLeak       = "org.apache.catalina.core.ThreadLocalLeakPreventionListener"
	ListenerNamingContextListener = "org.apache.catalina.startup.NamingContextListener"
)

// GetListenerDescription returns a human-readable description for a listener
func GetListenerDescription(className string) string {
	descriptions := map[string]string{
		ListenerVersionLogger:         "Version Logger - Logs server version at startup",
		ListenerAprLifecycle:          "APR Lifecycle - Initializes APR/Native library",
		ListenerJreMemoryLeak:         "JRE Memory Leak Prevention - Prevents JRE memory leaks",
		ListenerGlobalResources:       "Global Resources - Initializes global JNDI resources",
		ListenerThreadLocalLeak:       "ThreadLocal Leak Prevention - Cleans up ThreadLocal variables",
		ListenerNamingContextListener: "Naming Context - Creates naming context for JNDI",
	}
	if desc, ok := descriptions[className]; ok {
		return desc
	}
	return className
}
