package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Settings represents the application settings
type Settings struct {
	LastCatalinaHome string           `json:"last_catalina_home"`
	LastCatalinaBase string           `json:"last_catalina_base"`
	RecentPaths      []TomcatInstance `json:"recent_paths"`
	Language         string           `json:"language,omitempty"`
}

// SettingsManager handles loading and saving settings
type SettingsManager struct {
	settingsPath string
	settings     *Settings
}

// NewSettingsManager creates a new settings manager
func NewSettingsManager() *SettingsManager {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".config", "tomcatkit")

	return &SettingsManager{
		settingsPath: filepath.Join(configDir, "settings.json"),
		settings:     &Settings{},
	}
}

// Load loads settings from file
func (m *SettingsManager) Load() error {
	data, err := os.ReadFile(m.settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create default settings
			m.settings = &Settings{}
			return nil
		}
		return err
	}

	return json.Unmarshal(data, m.settings)
}

// Save saves settings to file
func (m *SettingsManager) Save() error {
	// Ensure directory exists
	dir := filepath.Dir(m.settingsPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(m.settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.settingsPath, data, 0644)
}

// GetSettings returns the current settings
func (m *SettingsManager) GetSettings() *Settings {
	return m.settings
}

// SetLastInstance saves the last used Tomcat instance
func (m *SettingsManager) SetLastInstance(instance *TomcatInstance) {
	m.settings.LastCatalinaHome = instance.CatalinaHome
	m.settings.LastCatalinaBase = instance.CatalinaBase

	// Add to recent paths if not already there
	found := false
	for i, recent := range m.settings.RecentPaths {
		if recent.CatalinaHome == instance.CatalinaHome {
			// Move to front
			m.settings.RecentPaths = append(
				[]TomcatInstance{*instance},
				append(m.settings.RecentPaths[:i], m.settings.RecentPaths[i+1:]...)...,
			)
			found = true
			break
		}
	}

	if !found {
		m.settings.RecentPaths = append([]TomcatInstance{*instance}, m.settings.RecentPaths...)
	}

	// Keep only last 5 recent paths
	if len(m.settings.RecentPaths) > 5 {
		m.settings.RecentPaths = m.settings.RecentPaths[:5]
	}
}

// GetLastInstance returns the last used Tomcat instance
func (m *SettingsManager) GetLastInstance() *TomcatInstance {
	if m.settings.LastCatalinaHome == "" {
		return nil
	}
	return &TomcatInstance{
		CatalinaHome: m.settings.LastCatalinaHome,
		CatalinaBase: m.settings.LastCatalinaBase,
	}
}

// GetRecentInstances returns recent Tomcat instances
func (m *SettingsManager) GetRecentInstances() []TomcatInstance {
	return m.settings.RecentPaths
}

// ClearRecentInstances clears all recent instances
func (m *SettingsManager) ClearRecentInstances() {
	m.settings.RecentPaths = nil
	m.settings.LastCatalinaHome = ""
	m.settings.LastCatalinaBase = ""
}

// GetLanguage returns the saved language setting
func (m *SettingsManager) GetLanguage() string {
	return m.settings.Language
}

// SetLanguage sets the language setting
func (m *SettingsManager) SetLanguage(lang string) {
	m.settings.Language = lang
}
