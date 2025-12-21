package web

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ConfigService handles web.xml operations
type ConfigService struct {
	catalinaBase string
	filePath     string
	webApp       *WebApp
}

// NewConfigService creates a new web.xml configuration service
func NewConfigService(catalinaBase string) *ConfigService {
	return &ConfigService{
		catalinaBase: catalinaBase,
		filePath:     filepath.Join(catalinaBase, "conf", "web.xml"),
	}
}

// GetFilePath returns the path to web.xml
func (s *ConfigService) GetFilePath() string {
	return s.filePath
}

// Load reads and parses web.xml
func (s *ConfigService) Load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create default web.xml if file doesn't exist
			s.webApp = NewWebApp()
			return nil
		}
		return fmt.Errorf("failed to read web.xml: %w", err)
	}

	var webapp WebApp
	if err := xml.Unmarshal(data, &webapp); err != nil {
		return fmt.Errorf("failed to parse web.xml: %w", err)
	}

	s.webApp = &webapp
	return nil
}

// Save writes the configuration back to web.xml
func (s *ConfigService) Save() error {
	if s.webApp == nil {
		return fmt.Errorf("no web.xml configuration loaded")
	}

	// Create backup first
	if err := s.createBackup(); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	data, err := xml.MarshalIndent(s.webApp, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal web.xml: %w", err)
	}

	// Add XML declaration and license comment
	output := []byte(xml.Header)
	output = append(output, []byte(`<!--
  Licensed to the Apache Software Foundation (ASF) under one or more
  contributor license agreements.  See the NOTICE file distributed with
  this work for additional information regarding copyright ownership.
  The ASF licenses this file to You under the Apache License, Version 2.0
  (the "License"); you may not use this file except in compliance with
  the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
-->
`)...)
	output = append(output, data...)

	if err := os.WriteFile(s.filePath, output, 0644); err != nil {
		return fmt.Errorf("failed to write web.xml: %w", err)
	}

	return nil
}

// createBackup creates a backup of the current web.xml
func (s *ConfigService) createBackup() error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return nil // No file to backup
	}

	backupDir := filepath.Join(s.catalinaBase, "conf", "backup")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(backupDir, fmt.Sprintf("web.xml.%s", timestamp))
	return os.WriteFile(backupPath, data, 0644)
}

// GetWebApp returns the current web application configuration
func (s *ConfigService) GetWebApp() *WebApp {
	return s.webApp
}

// --- Servlet operations ---

// GetServlets returns all servlets
func (s *ConfigService) GetServlets() []Servlet {
	if s.webApp != nil {
		return s.webApp.Servlets
	}
	return nil
}

// GetServlet returns a servlet by name
func (s *ConfigService) GetServlet(name string) *Servlet {
	if s.webApp != nil {
		for i := range s.webApp.Servlets {
			if s.webApp.Servlets[i].ServletName == name {
				return &s.webApp.Servlets[i]
			}
		}
	}
	return nil
}

// AddServlet adds a new servlet
func (s *ConfigService) AddServlet(servlet Servlet) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, srv := range s.webApp.Servlets {
		if srv.ServletName == servlet.ServletName {
			return fmt.Errorf("servlet '%s' already exists", servlet.ServletName)
		}
	}

	s.webApp.Servlets = append(s.webApp.Servlets, servlet)
	return nil
}

// UpdateServlet updates an existing servlet
func (s *ConfigService) UpdateServlet(name string, servlet Servlet) error {
	if s.webApp != nil {
		for i := range s.webApp.Servlets {
			if s.webApp.Servlets[i].ServletName == name {
				s.webApp.Servlets[i] = servlet
				return nil
			}
		}
	}
	return fmt.Errorf("servlet '%s' not found", name)
}

// DeleteServlet deletes a servlet and its mappings
func (s *ConfigService) DeleteServlet(name string) error {
	if s.webApp == nil {
		return fmt.Errorf("servlet '%s' not found", name)
	}

	// Remove servlet
	found := false
	var newServlets []Servlet
	for _, srv := range s.webApp.Servlets {
		if srv.ServletName == name {
			found = true
		} else {
			newServlets = append(newServlets, srv)
		}
	}
	if !found {
		return fmt.Errorf("servlet '%s' not found", name)
	}
	s.webApp.Servlets = newServlets

	// Remove mappings
	var newMappings []ServletMapping
	for _, m := range s.webApp.ServletMappings {
		if m.ServletName != name {
			newMappings = append(newMappings, m)
		}
	}
	s.webApp.ServletMappings = newMappings

	return nil
}

// GetServletMappings returns all servlet mappings
func (s *ConfigService) GetServletMappings() []ServletMapping {
	if s.webApp != nil {
		return s.webApp.ServletMappings
	}
	return nil
}

// AddServletMapping adds a servlet mapping
func (s *ConfigService) AddServletMapping(mapping ServletMapping) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.ServletMappings = append(s.webApp.ServletMappings, mapping)
}

// --- Filter operations ---

// GetFilters returns all filters
func (s *ConfigService) GetFilters() []Filter {
	if s.webApp != nil {
		return s.webApp.Filters
	}
	return nil
}

// GetFilter returns a filter by name
func (s *ConfigService) GetFilter(name string) *Filter {
	if s.webApp != nil {
		for i := range s.webApp.Filters {
			if s.webApp.Filters[i].FilterName == name {
				return &s.webApp.Filters[i]
			}
		}
	}
	return nil
}

// AddFilter adds a new filter
func (s *ConfigService) AddFilter(filter Filter) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, f := range s.webApp.Filters {
		if f.FilterName == filter.FilterName {
			return fmt.Errorf("filter '%s' already exists", filter.FilterName)
		}
	}

	s.webApp.Filters = append(s.webApp.Filters, filter)
	return nil
}

// UpdateFilter updates an existing filter
func (s *ConfigService) UpdateFilter(name string, filter Filter) error {
	if s.webApp != nil {
		for i := range s.webApp.Filters {
			if s.webApp.Filters[i].FilterName == name {
				s.webApp.Filters[i] = filter
				return nil
			}
		}
	}
	return fmt.Errorf("filter '%s' not found", name)
}

// DeleteFilter deletes a filter and its mappings
func (s *ConfigService) DeleteFilter(name string) error {
	if s.webApp == nil {
		return fmt.Errorf("filter '%s' not found", name)
	}

	// Remove filter
	found := false
	var newFilters []Filter
	for _, f := range s.webApp.Filters {
		if f.FilterName == name {
			found = true
		} else {
			newFilters = append(newFilters, f)
		}
	}
	if !found {
		return fmt.Errorf("filter '%s' not found", name)
	}
	s.webApp.Filters = newFilters

	// Remove mappings
	var newMappings []FilterMapping
	for _, m := range s.webApp.FilterMappings {
		if m.FilterName != name {
			newMappings = append(newMappings, m)
		}
	}
	s.webApp.FilterMappings = newMappings

	return nil
}

// GetFilterMappings returns all filter mappings
func (s *ConfigService) GetFilterMappings() []FilterMapping {
	if s.webApp != nil {
		return s.webApp.FilterMappings
	}
	return nil
}

// AddFilterMapping adds a filter mapping
func (s *ConfigService) AddFilterMapping(mapping FilterMapping) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.FilterMappings = append(s.webApp.FilterMappings, mapping)
}

// --- Listener operations ---

// GetListeners returns all listeners
func (s *ConfigService) GetListeners() []Listener {
	if s.webApp != nil {
		return s.webApp.Listeners
	}
	return nil
}

// AddListener adds a new listener
func (s *ConfigService) AddListener(listener Listener) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, l := range s.webApp.Listeners {
		if l.ListenerClass == listener.ListenerClass {
			return fmt.Errorf("listener '%s' already exists", listener.ListenerClass)
		}
	}

	s.webApp.Listeners = append(s.webApp.Listeners, listener)
	return nil
}

// DeleteListener deletes a listener by class
func (s *ConfigService) DeleteListener(listenerClass string) error {
	if s.webApp == nil {
		return fmt.Errorf("listener '%s' not found", listenerClass)
	}

	var newListeners []Listener
	found := false
	for _, l := range s.webApp.Listeners {
		if l.ListenerClass == listenerClass {
			found = true
		} else {
			newListeners = append(newListeners, l)
		}
	}
	if !found {
		return fmt.Errorf("listener '%s' not found", listenerClass)
	}
	s.webApp.Listeners = newListeners
	return nil
}

// --- Session config operations ---

// GetSessionConfig returns the session configuration
func (s *ConfigService) GetSessionConfig() *SessionConfig {
	if s.webApp != nil {
		return s.webApp.SessionConfig
	}
	return nil
}

// SetSessionConfig sets the session configuration
func (s *ConfigService) SetSessionConfig(config *SessionConfig) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.SessionConfig = config
}

// --- Welcome file operations ---

// GetWelcomeFiles returns the welcome file list
func (s *ConfigService) GetWelcomeFiles() []string {
	if s.webApp != nil && s.webApp.WelcomeFileList != nil {
		return s.webApp.WelcomeFileList.WelcomeFiles
	}
	return nil
}

// SetWelcomeFiles sets the welcome file list
func (s *ConfigService) SetWelcomeFiles(files []string) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	if s.webApp.WelcomeFileList == nil {
		s.webApp.WelcomeFileList = &WelcomeFileList{}
	}
	s.webApp.WelcomeFileList.WelcomeFiles = files
}

// --- Error page operations ---

// GetErrorPages returns all error pages
func (s *ConfigService) GetErrorPages() []ErrorPage {
	if s.webApp != nil {
		return s.webApp.ErrorPages
	}
	return nil
}

// AddErrorPage adds an error page
func (s *ConfigService) AddErrorPage(errorPage ErrorPage) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.ErrorPages = append(s.webApp.ErrorPages, errorPage)
}

// DeleteErrorPage deletes an error page
func (s *ConfigService) DeleteErrorPage(index int) error {
	if s.webApp == nil || index < 0 || index >= len(s.webApp.ErrorPages) {
		return fmt.Errorf("error page at index %d not found", index)
	}
	s.webApp.ErrorPages = append(s.webApp.ErrorPages[:index], s.webApp.ErrorPages[index+1:]...)
	return nil
}

// --- MIME mapping operations ---

// GetMimeMappings returns all MIME mappings
func (s *ConfigService) GetMimeMappings() []MimeMapping {
	if s.webApp != nil {
		return s.webApp.MimeMappings
	}
	return nil
}

// AddMimeMapping adds a MIME mapping
func (s *ConfigService) AddMimeMapping(mapping MimeMapping) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, m := range s.webApp.MimeMappings {
		if m.Extension == mapping.Extension {
			return fmt.Errorf("MIME mapping for '%s' already exists", mapping.Extension)
		}
	}

	s.webApp.MimeMappings = append(s.webApp.MimeMappings, mapping)
	return nil
}

// DeleteMimeMapping deletes a MIME mapping by extension
func (s *ConfigService) DeleteMimeMapping(extension string) error {
	if s.webApp == nil {
		return fmt.Errorf("MIME mapping for '%s' not found", extension)
	}

	var newMappings []MimeMapping
	found := false
	for _, m := range s.webApp.MimeMappings {
		if m.Extension == extension {
			found = true
		} else {
			newMappings = append(newMappings, m)
		}
	}
	if !found {
		return fmt.Errorf("MIME mapping for '%s' not found", extension)
	}
	s.webApp.MimeMappings = newMappings
	return nil
}

// --- Security constraint operations ---

// GetSecurityConstraints returns all security constraints
func (s *ConfigService) GetSecurityConstraints() []SecurityConstraint {
	if s.webApp != nil {
		return s.webApp.SecurityConstraints
	}
	return nil
}

// AddSecurityConstraint adds a security constraint
func (s *ConfigService) AddSecurityConstraint(constraint SecurityConstraint) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.SecurityConstraints = append(s.webApp.SecurityConstraints, constraint)
}

// UpdateSecurityConstraint updates a security constraint at index
func (s *ConfigService) UpdateSecurityConstraint(index int, constraint SecurityConstraint) error {
	if s.webApp == nil || index < 0 || index >= len(s.webApp.SecurityConstraints) {
		return fmt.Errorf("security constraint at index %d not found", index)
	}
	s.webApp.SecurityConstraints[index] = constraint
	return nil
}

// DeleteSecurityConstraint deletes a security constraint
func (s *ConfigService) DeleteSecurityConstraint(index int) error {
	if s.webApp == nil || index < 0 || index >= len(s.webApp.SecurityConstraints) {
		return fmt.Errorf("security constraint at index %d not found", index)
	}
	s.webApp.SecurityConstraints = append(s.webApp.SecurityConstraints[:index], s.webApp.SecurityConstraints[index+1:]...)
	return nil
}

// --- Login config operations ---

// GetLoginConfig returns the login configuration
func (s *ConfigService) GetLoginConfig() *LoginConfig {
	if s.webApp != nil {
		return s.webApp.LoginConfig
	}
	return nil
}

// SetLoginConfig sets the login configuration
func (s *ConfigService) SetLoginConfig(config *LoginConfig) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.LoginConfig = config
}

// RemoveLoginConfig removes the login configuration
func (s *ConfigService) RemoveLoginConfig() {
	if s.webApp != nil {
		s.webApp.LoginConfig = nil
	}
}

// --- Security role operations ---

// GetSecurityRoles returns all security roles
func (s *ConfigService) GetSecurityRoles() []SecurityRole {
	if s.webApp != nil {
		return s.webApp.SecurityRoles
	}
	return nil
}

// AddSecurityRole adds a security role
func (s *ConfigService) AddSecurityRole(role SecurityRole) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, r := range s.webApp.SecurityRoles {
		if r.RoleName == role.RoleName {
			return fmt.Errorf("security role '%s' already exists", role.RoleName)
		}
	}

	s.webApp.SecurityRoles = append(s.webApp.SecurityRoles, role)
	return nil
}

// DeleteSecurityRole deletes a security role
func (s *ConfigService) DeleteSecurityRole(roleName string) error {
	if s.webApp == nil {
		return fmt.Errorf("security role '%s' not found", roleName)
	}

	var newRoles []SecurityRole
	found := false
	for _, r := range s.webApp.SecurityRoles {
		if r.RoleName == roleName {
			found = true
		} else {
			newRoles = append(newRoles, r)
		}
	}
	if !found {
		return fmt.Errorf("security role '%s' not found", roleName)
	}
	s.webApp.SecurityRoles = newRoles
	return nil
}

// --- Context param operations ---

// GetContextParams returns all context parameters
func (s *ConfigService) GetContextParams() []ContextParam {
	if s.webApp != nil {
		return s.webApp.ContextParams
	}
	return nil
}

// AddContextParam adds a context parameter
func (s *ConfigService) AddContextParam(param ContextParam) error {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}

	for _, p := range s.webApp.ContextParams {
		if p.ParamName == param.ParamName {
			return fmt.Errorf("context parameter '%s' already exists", param.ParamName)
		}
	}

	s.webApp.ContextParams = append(s.webApp.ContextParams, param)
	return nil
}

// UpdateContextParam updates a context parameter
func (s *ConfigService) UpdateContextParam(name string, param ContextParam) error {
	if s.webApp != nil {
		for i := range s.webApp.ContextParams {
			if s.webApp.ContextParams[i].ParamName == name {
				s.webApp.ContextParams[i] = param
				return nil
			}
		}
	}
	return fmt.Errorf("context parameter '%s' not found", name)
}

// DeleteContextParam deletes a context parameter
func (s *ConfigService) DeleteContextParam(name string) error {
	if s.webApp == nil {
		return fmt.Errorf("context parameter '%s' not found", name)
	}

	var newParams []ContextParam
	found := false
	for _, p := range s.webApp.ContextParams {
		if p.ParamName == name {
			found = true
		} else {
			newParams = append(newParams, p)
		}
	}
	if !found {
		return fmt.Errorf("context parameter '%s' not found", name)
	}
	s.webApp.ContextParams = newParams
	return nil
}

// --- JSP config operations ---

// GetJspConfig returns the JSP configuration
func (s *ConfigService) GetJspConfig() *JspConfig {
	if s.webApp != nil {
		return s.webApp.JspConfig
	}
	return nil
}

// SetJspConfig sets the JSP configuration
func (s *ConfigService) SetJspConfig(config *JspConfig) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.JspConfig = config
}

// --- General settings ---

// SetDisplayName sets the display name
func (s *ConfigService) SetDisplayName(name string) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	s.webApp.DisplayName = name
}

// SetDistributable sets the distributable flag
func (s *ConfigService) SetDistributable(distributable bool) {
	if s.webApp == nil {
		s.webApp = NewWebApp()
	}
	if distributable {
		s.webApp.DistributableElement = &Distributable{}
	} else {
		s.webApp.DistributableElement = nil
	}
}

// IsDistributable returns whether the app is distributable
func (s *ConfigService) IsDistributable() bool {
	return s.webApp != nil && s.webApp.DistributableElement != nil
}
