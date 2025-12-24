package views

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/playok/tomcatkit/internal/config/jndi"
	"github.com/playok/tomcatkit/internal/config/server"
	"github.com/playok/tomcatkit/internal/i18n"
	"github.com/rivo/tview"
)

// PreviewPanel is a panel that shows XML/properties preview
type PreviewPanel struct {
	*tview.TextView
}

// NewPreviewPanel creates a new preview panel
func NewPreviewPanel() *PreviewPanel {
	panel := &PreviewPanel{
		TextView: tview.NewTextView().
			SetDynamicColors(true).
			SetWordWrap(true).
			SetScrollable(true),
	}
	panel.SetBorder(true).
		SetTitle(" " + i18n.T("preview.title") + " ").
		SetBorderColor(tcell.ColorBlue)

	return panel
}

// SetPreview updates the preview content
func (p *PreviewPanel) SetPreview(content string) {
	p.SetText(content)
}

// SetXMLPreview formats and sets XML content with syntax highlighting
func (p *PreviewPanel) SetXMLPreview(content string) {
	// Add simple syntax highlighting for XML
	highlighted := highlightXML(content)
	p.SetText(highlighted)
}

// highlightXML adds color codes for XML syntax
func highlightXML(xmlContent string) string {
	// Simple XML syntax highlighting
	result := xmlContent

	// Highlight tag names
	result = strings.ReplaceAll(result, "<Connector", "[yellow]<Connector[white]")
	result = strings.ReplaceAll(result, "<Executor", "[yellow]<Executor[white]")
	result = strings.ReplaceAll(result, "<Host", "[yellow]<Host[white]")
	result = strings.ReplaceAll(result, "<Context", "[yellow]<Context[white]")
	result = strings.ReplaceAll(result, "<Parameter", "[yellow]<Parameter[white]")
	result = strings.ReplaceAll(result, "<Manager", "[yellow]<Manager[white]")
	result = strings.ReplaceAll(result, "<Valve", "[yellow]<Valve[white]")
	result = strings.ReplaceAll(result, "<Realm", "[yellow]<Realm[white]")
	result = strings.ReplaceAll(result, "<Resource", "[yellow]<Resource[white]")
	result = strings.ReplaceAll(result, "<Environment", "[yellow]<Environment[white]")
	result = strings.ReplaceAll(result, "<ResourceLink", "[yellow]<ResourceLink[white]")
	result = strings.ReplaceAll(result, "<Listener", "[yellow]<Listener[white]")
	result = strings.ReplaceAll(result, "<Server", "[yellow]<Server[white]")
	result = strings.ReplaceAll(result, "<Service", "[yellow]<Service[white]")
	result = strings.ReplaceAll(result, "<Engine", "[yellow]<Engine[white]")
	result = strings.ReplaceAll(result, "<user", "[yellow]<user[white]")
	result = strings.ReplaceAll(result, "<role", "[yellow]<role[white]")
	result = strings.ReplaceAll(result, "/>", "[yellow]/>[white]")
	result = strings.ReplaceAll(result, ">", "[yellow]>[white]")

	// Highlight attribute values (between quotes)
	lines := strings.Split(result, "\n")
	for i, line := range lines {
		// Find and highlight quoted values
		newLine := ""
		inQuote := false
		for j := 0; j < len(line); j++ {
			if line[j] == '"' {
				if !inQuote {
					newLine += "[green]\""
					inQuote = true
				} else {
					newLine += "\"[white]"
					inQuote = false
				}
			} else {
				newLine += string(line[j])
			}
		}
		lines[i] = newLine
	}

	return strings.Join(lines, "\n")
}

// CreateFormWithPreview creates a flex layout with form on top and preview panel on bottom
func CreateFormWithPreview(form *tview.Form, preview *PreviewPanel) *tview.Flex {
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 2, true).
		AddItem(preview, 0, 1, false)

	return flex
}

// CreateFormWithHelpAndPreview creates a layout with form (left), help (right top), preview (right bottom)
func CreateFormWithHelpAndPreview(form *tview.Form, helpPanel *DynamicHelpPanel, preview *PreviewPanel) *tview.Flex {
	rightPanel := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(helpPanel, 0, 1, false).
		AddItem(preview, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(rightPanel, 0, 1, false)

	return flex
}

// GenerateConnectorXML generates XML preview for a connector
func GenerateConnectorXML(conn *server.Connector) string {
	// Create a simplified view for preview
	type ConnectorPreview struct {
		XMLName           xml.Name `xml:"Connector"`
		Port              int      `xml:"port,attr"`
		Protocol          string   `xml:"protocol,attr,omitempty"`
		ConnectionTimeout int      `xml:"connectionTimeout,attr,omitempty"`
		RedirectPort      int      `xml:"redirectPort,attr,omitempty"`
		MaxThreads        int      `xml:"maxThreads,attr,omitempty"`
		MinSpareThreads   int      `xml:"minSpareThreads,attr,omitempty"`
		AcceptCount       int      `xml:"acceptCount,attr,omitempty"`
		Executor          string   `xml:"executor,attr,omitempty"`
		SSLEnabled        bool     `xml:"SSLEnabled,attr,omitempty"`
		Scheme            string   `xml:"scheme,attr,omitempty"`
		Secure            bool     `xml:"secure,attr,omitempty"`
		KeystoreFile      string   `xml:"keystoreFile,attr,omitempty"`
		KeystoreType      string   `xml:"keystoreType,attr,omitempty"`
		SSLProtocol       string   `xml:"sslProtocol,attr,omitempty"`
		ClientAuth        string   `xml:"clientAuth,attr,omitempty"`
		SecretRequired    bool     `xml:"secretRequired,attr,omitempty"`
		Secret            string   `xml:"secret,attr,omitempty"`
	}

	preview := ConnectorPreview{
		Port:              conn.Port,
		Protocol:          conn.Protocol,
		ConnectionTimeout: conn.ConnectionTimeout,
		RedirectPort:      conn.RedirectPort,
		MaxThreads:        conn.MaxThreads,
		MinSpareThreads:   conn.MinSpareThreads,
		AcceptCount:       conn.AcceptCount,
		Executor:          conn.Executor,
		SSLEnabled:        conn.SSLEnabled,
		Scheme:            conn.Scheme,
		Secure:            conn.Secure,
		KeystoreFile:      conn.KeystoreFile,
		KeystoreType:      conn.KeystoreType,
		SSLProtocol:       conn.SSLProtocol,
		ClientAuth:        conn.ClientAuth,
		SecretRequired:    conn.SecretRequired,
		Secret:            conn.Secret,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateExecutorXML generates XML preview for an executor
func GenerateExecutorXML(exec *server.Executor) string {
	type ExecutorPreview struct {
		XMLName         xml.Name `xml:"Executor"`
		ClassName       string   `xml:"className,attr,omitempty"`
		Name            string   `xml:"name,attr"`
		NamePrefix      string   `xml:"namePrefix,attr,omitempty"`
		MaxThreads      int      `xml:"maxThreads,attr,omitempty"`
		MinSpareThreads int      `xml:"minSpareThreads,attr,omitempty"`
		MaxIdleTime     int      `xml:"maxIdleTime,attr,omitempty"`
		MaxQueueSize    int      `xml:"maxQueueSize,attr,omitempty"`
	}

	preview := ExecutorPreview{
		ClassName:       exec.ClassName,
		Name:            exec.Name,
		NamePrefix:      exec.NamePrefix,
		MaxThreads:      exec.MaxThreads,
		MinSpareThreads: exec.MinSpareThreads,
		MaxIdleTime:     exec.MaxIdleTime,
		MaxQueueSize:    exec.MaxQueueSize,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateHostXML generates XML preview for a host
func GenerateHostXML(host *server.Host) string {
	type HostPreview struct {
		XMLName         xml.Name `xml:"Host"`
		Name            string   `xml:"name,attr"`
		AppBase         string   `xml:"appBase,attr,omitempty"`
		UnpackWARs      bool     `xml:"unpackWARs,attr,omitempty"`
		AutoDeploy      bool     `xml:"autoDeploy,attr,omitempty"`
		DeployOnStartup bool     `xml:"deployOnStartup,attr,omitempty"`
	}

	preview := HostPreview{
		Name:            host.Name,
		AppBase:         host.AppBase,
		UnpackWARs:      host.UnpackWARs,
		AutoDeploy:      host.AutoDeploy,
		DeployOnStartup: host.DeployOnStartup,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateServerXML generates XML preview for server settings
func GenerateServerXML(srv *server.Server) string {
	type ServerPreview struct {
		XMLName  xml.Name `xml:"Server"`
		Port     int      `xml:"port,attr"`
		Shutdown string   `xml:"shutdown,attr"`
	}

	preview := ServerPreview{
		Port:     srv.Port,
		Shutdown: srv.Shutdown,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateListenerXML generates XML preview for a listener
func GenerateListenerXML(listener *server.Listener) string {
	type ListenerPreview struct {
		XMLName       xml.Name `xml:"Listener"`
		ClassName     string   `xml:"className,attr"`
		SSLEngine     string   `xml:"SSLEngine,attr,omitempty"`
		SSLRandomSeed string   `xml:"SSLRandomSeed,attr,omitempty"`
	}

	preview := ListenerPreview{
		ClassName:     listener.ClassName,
		SSLEngine:     listener.SSLEngine,
		SSLRandomSeed: listener.SSLRandomSeed,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateServiceXML generates XML preview for a service
func GenerateServiceXML(svc *server.Service) string {
	type ServicePreview struct {
		XMLName xml.Name `xml:"Service"`
		Name    string   `xml:"name,attr"`
	}

	preview := ServicePreview{
		Name: svc.Name,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateEngineXML generates XML preview for an engine
func GenerateEngineXML(engine *server.Engine) string {
	type EnginePreview struct {
		XMLName     xml.Name `xml:"Engine"`
		Name        string   `xml:"name,attr"`
		DefaultHost string   `xml:"defaultHost,attr,omitempty"`
		JvmRoute    string   `xml:"jvmRoute,attr,omitempty"`
	}

	preview := EnginePreview{
		Name:        engine.Name,
		DefaultHost: engine.DefaultHost,
		JvmRoute:    engine.JvmRoute,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateValveXML generates XML preview for a valve
func GenerateValveXML(valve *server.Valve) string {
	output, err := xml.MarshalIndent(valve, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateRealmXML generates XML preview for a realm
func GenerateRealmXML(realm *server.Realm) string {
	output, err := xml.MarshalIndent(realm, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateResourceXML generates XML preview for a global resource
func GenerateResourceXML(res *server.Resource) string {
	type ResourcePreview struct {
		XMLName     xml.Name `xml:"Resource"`
		Name        string   `xml:"name,attr"`
		Auth        string   `xml:"auth,attr,omitempty"`
		Type        string   `xml:"type,attr,omitempty"`
		Description string   `xml:"description,attr,omitempty"`
		Factory     string   `xml:"factory,attr,omitempty"`
		Pathname    string   `xml:"pathname,attr,omitempty"`
	}

	preview := ResourcePreview{
		Name:        res.Name,
		Auth:        res.Auth,
		Type:        res.Type,
		Description: res.Description,
		Factory:     res.Factory,
		Pathname:    res.Pathname,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateContextXML generates XML preview for a context
func GenerateContextXML(ctx *server.Context) string {
	type ContextPreview struct {
		XMLName             xml.Name `xml:"Context"`
		Path                string   `xml:"path,attr"`
		DocBase             string   `xml:"docBase,attr,omitempty"`
		Reloadable          bool     `xml:"reloadable,attr,omitempty"`
		CrossContext        bool     `xml:"crossContext,attr,omitempty"`
		Privileged          bool     `xml:"privileged,attr,omitempty"`
		Cookies             bool     `xml:"cookies,attr,omitempty"`
		SessionCookieName   string   `xml:"sessionCookieName,attr,omitempty"`
		SessionCookiePath   string   `xml:"sessionCookiePath,attr,omitempty"`
		SessionCookieDomain string   `xml:"sessionCookieDomain,attr,omitempty"`
		UseHttpOnly         bool     `xml:"useHttpOnly,attr,omitempty"`
		AntiResourceLocking bool     `xml:"antiResourceLocking,attr,omitempty"`
		SwallowOutput       bool     `xml:"swallowOutput,attr,omitempty"`
		CachingAllowed      bool     `xml:"cachingAllowed,attr,omitempty"`
		CacheMaxSize        int      `xml:"cacheMaxSize,attr,omitempty"`
		CacheTTL            int      `xml:"cacheTTL,attr,omitempty"`
	}

	preview := ContextPreview{
		Path:                ctx.Path,
		DocBase:             ctx.DocBase,
		Reloadable:          ctx.Reloadable,
		CrossContext:        ctx.CrossContext,
		Privileged:          ctx.Privileged,
		Cookies:             ctx.Cookies,
		SessionCookieName:   ctx.SessionCookieName,
		SessionCookiePath:   ctx.SessionCookiePath,
		SessionCookieDomain: ctx.SessionCookieDomain,
		UseHttpOnly:         ctx.UseHttpOnly,
		AntiResourceLocking: ctx.AntiResourceLocking,
		SwallowOutput:       ctx.SwallowOutput,
		CachingAllowed:      ctx.CachingAllowed,
		CacheMaxSize:        ctx.CacheMaxSize,
		CacheTTL:            ctx.CacheTTL,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateParameterXML generates XML preview for a parameter
func GenerateParameterXML(param *server.Parameter) string {
	type ParameterPreview struct {
		XMLName     xml.Name `xml:"Parameter"`
		Name        string   `xml:"name,attr"`
		Value       string   `xml:"value,attr"`
		Override    bool     `xml:"override,attr,omitempty"`
		Description string   `xml:"description,attr,omitempty"`
	}

	preview := ParameterPreview{
		Name:        param.Name,
		Value:       param.Value,
		Override:    param.Override,
		Description: param.Description,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateManagerXML generates XML preview for a manager
func GenerateManagerXML(mgr *server.Manager) string {
	type ManagerPreview struct {
		XMLName                 xml.Name `xml:"Manager"`
		ClassName               string   `xml:"className,attr,omitempty"`
		MaxActiveSessions       int      `xml:"maxActiveSessions,attr,omitempty"`
		SessionIdLength         int      `xml:"sessionIdLength,attr,omitempty"`
		MaxInactiveInterval     int      `xml:"maxInactiveInterval,attr,omitempty"`
		Pathname                string   `xml:"pathname,attr,omitempty"`
		ProcessExpiresFrequency int      `xml:"processExpiresFrequency,attr,omitempty"`
	}

	preview := ManagerPreview{
		ClassName:               mgr.ClassName,
		MaxActiveSessions:       mgr.MaxActiveSessions,
		SessionIdLength:         mgr.SessionIdLength,
		MaxInactiveInterval:     mgr.MaxInactiveInterval,
		Pathname:                mgr.Pathname,
		ProcessExpiresFrequency: mgr.ProcessExpiresFrequency,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GeneratePropertiesPreview generates a properties file preview
func GeneratePropertiesPreview(props map[string]string) string {
	var lines []string
	for key, value := range props {
		lines = append(lines, fmt.Sprintf("[cyan]%s[white] = [green]%s[white]", key, value))
	}
	return strings.Join(lines, "\n")
}

// GenerateUserXML generates XML preview for a user
func GenerateUserXML(username, password, roles string) string {
	type UserPreview struct {
		XMLName  xml.Name `xml:"user"`
		Username string   `xml:"username,attr"`
		Password string   `xml:"password,attr"`
		Roles    string   `xml:"roles,attr"`
	}

	preview := UserPreview{
		Username: username,
		Password: password,
		Roles:    roles,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateRoleXML generates XML preview for a role
func GenerateRoleXML(roleName, description string) string {
	type RolePreview struct {
		XMLName     xml.Name `xml:"role"`
		RoleName    string   `xml:"rolename,attr"`
		Description string   `xml:"description,attr,omitempty"`
	}

	preview := RolePreview{
		RoleName:    roleName,
		Description: description,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateJNDIResourceXML generates XML preview for a JNDI Resource (DataSource/MailSession)
func GenerateJNDIResourceXML(res *jndi.Resource) string {
	output, err := xml.MarshalIndent(res, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateEnvironmentXML generates XML preview for an Environment entry
func GenerateEnvironmentXML(env *jndi.Environment) string {
	output, err := xml.MarshalIndent(env, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateResourceLinkXML generates XML preview for a ResourceLink
func GenerateResourceLinkXML(link *jndi.ResourceLink) string {
	output, err := xml.MarshalIndent(link, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateContextSettingsXML generates XML preview for context settings (from context.xml)
func GenerateContextSettingsXML(ctx *jndi.Context) string {
	type ContextSettingsPreview struct {
		XMLName             xml.Name `xml:"Context"`
		Reloadable          bool     `xml:"reloadable,attr,omitempty"`
		CrossContext        bool     `xml:"crossContext,attr,omitempty"`
		Privileged          bool     `xml:"privileged,attr,omitempty"`
		Cookies             string   `xml:"cookies,attr,omitempty"`
		UseHttpOnly         string   `xml:"useHttpOnly,attr,omitempty"`
		SessionCookieName   string   `xml:"sessionCookieName,attr,omitempty"`
		CachingAllowed      string   `xml:"cachingAllowed,attr,omitempty"`
		CacheMaxSize        int      `xml:"cacheMaxSize,attr,omitempty"`
		AntiResourceLocking string   `xml:"antiResourceLocking,attr,omitempty"`
		SwallowOutput       string   `xml:"swallowOutput,attr,omitempty"`
	}

	preview := ContextSettingsPreview{
		Reloadable:          ctx.Reloadable,
		CrossContext:        ctx.CrossContext,
		Privileged:          ctx.Privileged,
		Cookies:             ctx.Cookies,
		UseHttpOnly:         ctx.UseHttpOnly,
		SessionCookieName:   ctx.SessionCookieName,
		CachingAllowed:      ctx.CachingAllowed,
		CacheMaxSize:        ctx.CacheMaxSize,
		AntiResourceLocking: ctx.AntiResourceLocking,
		SwallowOutput:       ctx.SwallowOutput,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateContextParameterXML generates XML preview for a context parameter
func GenerateContextParameterXML(param *jndi.ContextParameter) string {
	type ParameterPreview struct {
		XMLName     xml.Name `xml:"Parameter"`
		Name        string   `xml:"name,attr"`
		Value       string   `xml:"value,attr"`
		Override    bool     `xml:"override,attr,omitempty"`
		Description string   `xml:"description,attr,omitempty"`
	}

	preview := ParameterPreview{
		Name:        param.Name,
		Value:       param.Value,
		Override:    param.Override,
		Description: param.Description,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateWatchedResourceXML generates XML preview for a watched resource
func GenerateWatchedResourceXML(resource string) string {
	type WatchedResourcePreview struct {
		XMLName xml.Name `xml:"WatchedResource"`
		Value   string   `xml:",chardata"`
	}

	preview := WatchedResourcePreview{
		Value: resource,
	}

	output, err := xml.MarshalIndent(preview, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateContextManagerXML generates XML preview for a context manager
func GenerateContextManagerXML(mgr *jndi.ContextManager) string {
	output, err := xml.MarshalIndent(mgr, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateCookieProcessorXML generates XML preview for a cookie processor
func GenerateCookieProcessorXML(processor *jndi.CookieProcessor) string {
	output, err := xml.MarshalIndent(processor, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateJarScannerXML generates XML preview for a jar scanner
func GenerateJarScannerXML(scanner *jndi.JarScanner) string {
	output, err := xml.MarshalIndent(scanner, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}

	return string(output)
}

// GenerateServletXML generates XML preview for a servlet
func GenerateServletXML(servlet interface{}) string {
	output, err := xml.MarshalIndent(servlet, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}
	return string(output)
}

// GenerateFilterXML generates XML preview for a filter
func GenerateFilterXML(filter interface{}) string {
	output, err := xml.MarshalIndent(filter, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}
	return string(output)
}

// GenerateSessionConfigXML generates XML preview for session configuration
func GenerateSessionConfigXML(sessionConfig interface{}) string {
	output, err := xml.MarshalIndent(sessionConfig, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}
	return string(output)
}

// GenerateSecurityConstraintXML generates XML preview for a security constraint
func GenerateSecurityConstraintXML(constraint interface{}) string {
	output, err := xml.MarshalIndent(constraint, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error generating preview: %v", err)
	}
	return string(output)
}
