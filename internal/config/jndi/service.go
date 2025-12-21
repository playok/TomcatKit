package jndi

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

// ContextService handles context.xml operations
type ContextService struct {
	catalinaBase string
	filePath     string
	context      *Context
}

// NewContextService creates a new context service
func NewContextService(catalinaBase string) *ContextService {
	return &ContextService{
		catalinaBase: catalinaBase,
		filePath:     filepath.Join(catalinaBase, "conf", "context.xml"),
	}
}

// Load reads and parses context.xml
func (s *ContextService) Load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create empty context if file doesn't exist
			s.context = &Context{
				WatchedResources: []string{"WEB-INF/web.xml", "${catalina.base}/conf/web.xml"},
			}
			return nil
		}
		return fmt.Errorf("failed to read context.xml: %w", err)
	}

	var ctx Context
	if err := xml.Unmarshal(data, &ctx); err != nil {
		return fmt.Errorf("failed to parse context.xml: %w", err)
	}

	s.context = &ctx
	return nil
}

// Save writes the configuration back to context.xml
func (s *ContextService) Save() error {
	if s.context == nil {
		return fmt.Errorf("no context configuration loaded")
	}

	// Create backup first
	if err := s.createBackup(); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	data, err := xml.MarshalIndent(s.context, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal context.xml: %w", err)
	}

	// Add XML declaration and license comment
	output := []byte(xml.Header)
	output = append(output, []byte(`<!-- Licensed to the Apache Software Foundation (ASF) -->
`)...)
	output = append(output, data...)

	if err := os.WriteFile(s.filePath, output, 0644); err != nil {
		return fmt.Errorf("failed to write context.xml: %w", err)
	}

	return nil
}

// createBackup creates a backup of the current context.xml
func (s *ContextService) createBackup() error {
	backupDir := filepath.Join(s.catalinaBase, "conf", "backup")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	backupPath := filepath.Join(backupDir, "context.xml.bak")
	return os.WriteFile(backupPath, data, 0644)
}

// GetContext returns the current context
func (s *ContextService) GetContext() *Context {
	return s.context
}

// GetResources returns all resources
func (s *ContextService) GetResources() []Resource {
	if s.context != nil {
		return s.context.Resources
	}
	return nil
}

// GetResource returns a resource by name
func (s *ContextService) GetResource(name string) *Resource {
	if s.context != nil {
		for i := range s.context.Resources {
			if s.context.Resources[i].Name == name {
				return &s.context.Resources[i]
			}
		}
	}
	return nil
}

// AddResource adds a new resource
func (s *ContextService) AddResource(resource Resource) error {
	if s.context == nil {
		s.context = &Context{}
	}

	// Check for duplicate
	for _, r := range s.context.Resources {
		if r.Name == resource.Name {
			return fmt.Errorf("resource '%s' already exists", resource.Name)
		}
	}

	s.context.Resources = append(s.context.Resources, resource)
	return nil
}

// UpdateResource updates an existing resource
func (s *ContextService) UpdateResource(name string, resource Resource) error {
	if s.context != nil {
		for i := range s.context.Resources {
			if s.context.Resources[i].Name == name {
				s.context.Resources[i] = resource
				return nil
			}
		}
	}
	return fmt.Errorf("resource '%s' not found", name)
}

// DeleteResource deletes a resource
func (s *ContextService) DeleteResource(name string) error {
	if s.context != nil {
		for i := range s.context.Resources {
			if s.context.Resources[i].Name == name {
				s.context.Resources = append(s.context.Resources[:i], s.context.Resources[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("resource '%s' not found", name)
}

// GetEnvironments returns all environment entries
func (s *ContextService) GetEnvironments() []Environment {
	if s.context != nil {
		return s.context.Environments
	}
	return nil
}

// GetEnvironment returns an environment by name
func (s *ContextService) GetEnvironment(name string) *Environment {
	if s.context != nil {
		for i := range s.context.Environments {
			if s.context.Environments[i].Name == name {
				return &s.context.Environments[i]
			}
		}
	}
	return nil
}

// AddEnvironment adds a new environment entry
func (s *ContextService) AddEnvironment(env Environment) error {
	if s.context == nil {
		s.context = &Context{}
	}

	// Check for duplicate
	for _, e := range s.context.Environments {
		if e.Name == env.Name {
			return fmt.Errorf("environment '%s' already exists", env.Name)
		}
	}

	s.context.Environments = append(s.context.Environments, env)
	return nil
}

// UpdateEnvironment updates an existing environment entry
func (s *ContextService) UpdateEnvironment(name string, env Environment) error {
	if s.context != nil {
		for i := range s.context.Environments {
			if s.context.Environments[i].Name == name {
				s.context.Environments[i] = env
				return nil
			}
		}
	}
	return fmt.Errorf("environment '%s' not found", name)
}

// DeleteEnvironment deletes an environment entry
func (s *ContextService) DeleteEnvironment(name string) error {
	if s.context != nil {
		for i := range s.context.Environments {
			if s.context.Environments[i].Name == name {
				s.context.Environments = append(s.context.Environments[:i], s.context.Environments[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("environment '%s' not found", name)
}

// GetResourceLinks returns all resource links
func (s *ContextService) GetResourceLinks() []ResourceLink {
	if s.context != nil {
		return s.context.ResourceLinks
	}
	return nil
}

// GetResourceLink returns a resource link by name
func (s *ContextService) GetResourceLink(name string) *ResourceLink {
	if s.context != nil {
		for i := range s.context.ResourceLinks {
			if s.context.ResourceLinks[i].Name == name {
				return &s.context.ResourceLinks[i]
			}
		}
	}
	return nil
}

// AddResourceLink adds a new resource link
func (s *ContextService) AddResourceLink(link ResourceLink) error {
	if s.context == nil {
		s.context = &Context{}
	}

	// Check for duplicate
	for _, l := range s.context.ResourceLinks {
		if l.Name == link.Name {
			return fmt.Errorf("resource link '%s' already exists", link.Name)
		}
	}

	s.context.ResourceLinks = append(s.context.ResourceLinks, link)
	return nil
}

// UpdateResourceLink updates an existing resource link
func (s *ContextService) UpdateResourceLink(name string, link ResourceLink) error {
	if s.context != nil {
		for i := range s.context.ResourceLinks {
			if s.context.ResourceLinks[i].Name == name {
				s.context.ResourceLinks[i] = link
				return nil
			}
		}
	}
	return fmt.Errorf("resource link '%s' not found", name)
}

// DeleteResourceLink deletes a resource link
func (s *ContextService) DeleteResourceLink(name string) error {
	if s.context != nil {
		for i := range s.context.ResourceLinks {
			if s.context.ResourceLinks[i].Name == name {
				s.context.ResourceLinks = append(s.context.ResourceLinks[:i], s.context.ResourceLinks[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("resource link '%s' not found", name)
}

// GetDataSources returns all DataSource resources
func (s *ContextService) GetDataSources() []Resource {
	var dataSources []Resource
	if s.context != nil {
		for _, r := range s.context.Resources {
			if r.Type == string(ResourceTypeDataSource) {
				dataSources = append(dataSources, r)
			}
		}
	}
	return dataSources
}

// GetMailSessions returns all Mail Session resources
func (s *ContextService) GetMailSessions() []Resource {
	var mailSessions []Resource
	if s.context != nil {
		for _, r := range s.context.Resources {
			if r.Type == string(ResourceTypeMailSession) {
				mailSessions = append(mailSessions, r)
			}
		}
	}
	return mailSessions
}

// GetFilePath returns the path to context.xml
func (s *ContextService) GetFilePath() string {
	return s.filePath
}

// GetParameters returns all context parameters
func (s *ContextService) GetParameters() []ContextParameter {
	if s.context != nil {
		return s.context.Parameters
	}
	return nil
}

// GetParameter returns a parameter by name
func (s *ContextService) GetParameter(name string) *ContextParameter {
	if s.context != nil {
		for i := range s.context.Parameters {
			if s.context.Parameters[i].Name == name {
				return &s.context.Parameters[i]
			}
		}
	}
	return nil
}

// AddParameter adds a new context parameter
func (s *ContextService) AddParameter(param ContextParameter) error {
	if s.context == nil {
		s.context = &Context{}
	}

	for _, p := range s.context.Parameters {
		if p.Name == param.Name {
			return fmt.Errorf("parameter '%s' already exists", param.Name)
		}
	}

	s.context.Parameters = append(s.context.Parameters, param)
	return nil
}

// UpdateParameter updates an existing parameter
func (s *ContextService) UpdateParameter(name string, param ContextParameter) error {
	if s.context != nil {
		for i := range s.context.Parameters {
			if s.context.Parameters[i].Name == name {
				s.context.Parameters[i] = param
				return nil
			}
		}
	}
	return fmt.Errorf("parameter '%s' not found", name)
}

// DeleteParameter deletes a parameter
func (s *ContextService) DeleteParameter(name string) error {
	if s.context != nil {
		for i := range s.context.Parameters {
			if s.context.Parameters[i].Name == name {
				s.context.Parameters = append(s.context.Parameters[:i], s.context.Parameters[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("parameter '%s' not found", name)
}

// GetWatchedResources returns all watched resources
func (s *ContextService) GetWatchedResources() []string {
	if s.context != nil {
		return s.context.WatchedResources
	}
	return nil
}

// SetWatchedResources sets the watched resources
func (s *ContextService) SetWatchedResources(resources []string) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.WatchedResources = resources
}

// AddWatchedResource adds a watched resource
func (s *ContextService) AddWatchedResource(resource string) {
	if s.context == nil {
		s.context = &Context{}
	}
	// Check for duplicate
	for _, r := range s.context.WatchedResources {
		if r == resource {
			return
		}
	}
	s.context.WatchedResources = append(s.context.WatchedResources, resource)
}

// RemoveWatchedResource removes a watched resource
func (s *ContextService) RemoveWatchedResource(resource string) {
	if s.context != nil {
		for i, r := range s.context.WatchedResources {
			if r == resource {
				s.context.WatchedResources = append(s.context.WatchedResources[:i], s.context.WatchedResources[i+1:]...)
				return
			}
		}
	}
}

// GetManager returns the session manager
func (s *ContextService) GetManager() *ContextManager {
	if s.context != nil {
		return s.context.Manager
	}
	return nil
}

// SetManager sets the session manager
func (s *ContextService) SetManager(manager *ContextManager) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.Manager = manager
}

// RemoveManager removes the session manager
func (s *ContextService) RemoveManager() {
	if s.context != nil {
		s.context.Manager = nil
	}
}

// GetLoader returns the class loader
func (s *ContextService) GetLoader() *ContextLoader {
	if s.context != nil {
		return s.context.Loader
	}
	return nil
}

// SetLoader sets the class loader
func (s *ContextService) SetLoader(loader *ContextLoader) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.Loader = loader
}

// RemoveLoader removes the class loader
func (s *ContextService) RemoveLoader() {
	if s.context != nil {
		s.context.Loader = nil
	}
}

// GetJarScanner returns the JAR scanner
func (s *ContextService) GetJarScanner() *JarScanner {
	if s.context != nil {
		return s.context.JarScanner
	}
	return nil
}

// SetJarScanner sets the JAR scanner
func (s *ContextService) SetJarScanner(scanner *JarScanner) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.JarScanner = scanner
}

// RemoveJarScanner removes the JAR scanner
func (s *ContextService) RemoveJarScanner() {
	if s.context != nil {
		s.context.JarScanner = nil
	}
}

// GetCookieProcessor returns the cookie processor
func (s *ContextService) GetCookieProcessor() *CookieProcessor {
	if s.context != nil {
		return s.context.CookieProcessor
	}
	return nil
}

// SetCookieProcessor sets the cookie processor
func (s *ContextService) SetCookieProcessor(processor *CookieProcessor) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.CookieProcessor = processor
}

// RemoveCookieProcessor removes the cookie processor
func (s *ContextService) RemoveCookieProcessor() {
	if s.context != nil {
		s.context.CookieProcessor = nil
	}
}

// GetValves returns all context valves
func (s *ContextService) GetValves() []ContextValve {
	if s.context != nil {
		return s.context.Valves
	}
	return nil
}

// AddValve adds a valve to context
func (s *ContextService) AddValve(valve ContextValve) {
	if s.context == nil {
		s.context = &Context{}
	}
	s.context.Valves = append(s.context.Valves, valve)
}

// UpdateValve updates a valve at index
func (s *ContextService) UpdateValve(index int, valve ContextValve) error {
	if s.context == nil || index < 0 || index >= len(s.context.Valves) {
		return fmt.Errorf("valve at index %d not found", index)
	}
	s.context.Valves[index] = valve
	return nil
}

// DeleteValve removes a valve at index
func (s *ContextService) DeleteValve(index int) error {
	if s.context == nil || index < 0 || index >= len(s.context.Valves) {
		return fmt.Errorf("valve at index %d not found", index)
	}
	s.context.Valves = append(s.context.Valves[:index], s.context.Valves[index+1:]...)
	return nil
}

// UpdateContextSettings updates the basic context settings
func (s *ContextService) UpdateContextSettings(ctx *Context) {
	if s.context == nil {
		s.context = &Context{}
	}
	// Copy only the settings, not the nested elements
	s.context.Reloadable = ctx.Reloadable
	s.context.CrossContext = ctx.CrossContext
	s.context.Privileged = ctx.Privileged
	s.context.Cookies = ctx.Cookies
	s.context.UseHttpOnly = ctx.UseHttpOnly
	s.context.SessionCookieName = ctx.SessionCookieName
	s.context.SessionCookiePath = ctx.SessionCookiePath
	s.context.SessionCookieDomain = ctx.SessionCookieDomain
	s.context.AntiResourceLocking = ctx.AntiResourceLocking
	s.context.AntiJARLocking = ctx.AntiJARLocking
	s.context.CachingAllowed = ctx.CachingAllowed
	s.context.CacheMaxSize = ctx.CacheMaxSize
	s.context.CacheTTL = ctx.CacheTTL
	s.context.SwallowOutput = ctx.SwallowOutput
	s.context.MapperContextRootRedirectEnabled = ctx.MapperContextRootRedirectEnabled
	s.context.MapperDirectoryRedirectEnabled = ctx.MapperDirectoryRedirectEnabled
	s.context.AllowCasualMultipartParsing = ctx.AllowCasualMultipartParsing
	s.context.Delegate = ctx.Delegate
	s.context.ParallelAnnotationScanning = ctx.ParallelAnnotationScanning
}
