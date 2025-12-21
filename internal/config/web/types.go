package web

import "encoding/xml"

// WebApp represents the web.xml structure
type WebApp struct {
	XMLName              xml.Name              `xml:"web-app"`
	XMLNS                string                `xml:"xmlns,attr,omitempty"`
	XSI                  string                `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation       string                `xml:"xsi:schemaLocation,attr,omitempty"`
	Version              string                `xml:"version,attr,omitempty"`
	MetadataComplete     string                `xml:"metadata-complete,attr,omitempty"`
	DisplayName          string                `xml:"display-name,omitempty"`
	Description          string                `xml:"description,omitempty"`
	DistributableElement *Distributable        `xml:"distributable,omitempty"`
	ContextParams        []ContextParam        `xml:"context-param"`
	Filters              []Filter              `xml:"filter"`
	FilterMappings       []FilterMapping       `xml:"filter-mapping"`
	Listeners            []Listener            `xml:"listener"`
	Servlets             []Servlet             `xml:"servlet"`
	ServletMappings      []ServletMapping      `xml:"servlet-mapping"`
	SessionConfig        *SessionConfig        `xml:"session-config,omitempty"`
	MimeMappings         []MimeMapping         `xml:"mime-mapping"`
	WelcomeFileList      *WelcomeFileList      `xml:"welcome-file-list,omitempty"`
	ErrorPages           []ErrorPage           `xml:"error-page"`
	JspConfig            *JspConfig            `xml:"jsp-config,omitempty"`
	SecurityConstraints  []SecurityConstraint  `xml:"security-constraint"`
	LoginConfig          *LoginConfig          `xml:"login-config,omitempty"`
	SecurityRoles        []SecurityRole        `xml:"security-role"`
	ResourceRefs         []ResourceRef         `xml:"resource-ref"`
	ResourceEnvRefs      []ResourceEnvRef      `xml:"resource-env-ref"`
	MessageDestRefs      []MessageDestRef      `xml:"message-destination-ref"`
	AbsoluteOrdering     *AbsoluteOrdering     `xml:"absolute-ordering,omitempty"`
	DenyUncoveredMethods *DenyUncoveredMethods `xml:"deny-uncovered-http-methods,omitempty"`
}

// Distributable marks the application as distributable
type Distributable struct {
	XMLName xml.Name `xml:"distributable"`
}

// ContextParam represents a context parameter
type ContextParam struct {
	XMLName     xml.Name `xml:"context-param"`
	ParamName   string   `xml:"param-name"`
	ParamValue  string   `xml:"param-value"`
	Description string   `xml:"description,omitempty"`
}

// Filter represents a filter definition
type Filter struct {
	XMLName        xml.Name    `xml:"filter"`
	FilterName     string      `xml:"filter-name"`
	FilterClass    string      `xml:"filter-class"`
	Description    string      `xml:"description,omitempty"`
	DisplayName    string      `xml:"display-name,omitempty"`
	AsyncSupported string      `xml:"async-supported,omitempty"`
	InitParams     []InitParam `xml:"init-param"`
}

// FilterMapping represents a filter URL/servlet mapping
type FilterMapping struct {
	XMLName      xml.Name `xml:"filter-mapping"`
	FilterName   string   `xml:"filter-name"`
	URLPatterns  []string `xml:"url-pattern"`
	ServletNames []string `xml:"servlet-name"`
	Dispatchers  []string `xml:"dispatcher"`
}

// Listener represents a listener definition
type Listener struct {
	XMLName       xml.Name `xml:"listener"`
	ListenerClass string   `xml:"listener-class"`
	Description   string   `xml:"description,omitempty"`
}

// Servlet represents a servlet definition
type Servlet struct {
	XMLName          xml.Name          `xml:"servlet"`
	ServletName      string            `xml:"servlet-name"`
	ServletClass     string            `xml:"servlet-class,omitempty"`
	JspFile          string            `xml:"jsp-file,omitempty"`
	Description      string            `xml:"description,omitempty"`
	DisplayName      string            `xml:"display-name,omitempty"`
	LoadOnStartup    string            `xml:"load-on-startup,omitempty"`
	AsyncSupported   string            `xml:"async-supported,omitempty"`
	Enabled          string            `xml:"enabled,omitempty"`
	InitParams       []InitParam       `xml:"init-param"`
	MultipartConfig  *MultipartConfig  `xml:"multipart-config,omitempty"`
	SecurityRoleRefs []SecurityRoleRef `xml:"security-role-ref"`
}

// InitParam represents an initialization parameter
type InitParam struct {
	XMLName     xml.Name `xml:"init-param"`
	ParamName   string   `xml:"param-name"`
	ParamValue  string   `xml:"param-value"`
	Description string   `xml:"description,omitempty"`
}

// MultipartConfig represents multipart file upload configuration
type MultipartConfig struct {
	XMLName           xml.Name `xml:"multipart-config"`
	Location          string   `xml:"location,omitempty"`
	MaxFileSize       int64    `xml:"max-file-size,omitempty"`
	MaxRequestSize    int64    `xml:"max-request-size,omitempty"`
	FileSizeThreshold int      `xml:"file-size-threshold,omitempty"`
}

// SecurityRoleRef represents a security role reference
type SecurityRoleRef struct {
	XMLName     xml.Name `xml:"security-role-ref"`
	RoleName    string   `xml:"role-name"`
	RoleLink    string   `xml:"role-link,omitempty"`
	Description string   `xml:"description,omitempty"`
}

// ServletMapping represents a servlet URL mapping
type ServletMapping struct {
	XMLName     xml.Name `xml:"servlet-mapping"`
	ServletName string   `xml:"servlet-name"`
	URLPatterns []string `xml:"url-pattern"`
}

// SessionConfig represents session configuration
type SessionConfig struct {
	XMLName        xml.Name      `xml:"session-config"`
	SessionTimeout int           `xml:"session-timeout,omitempty"`
	CookieConfig   *CookieConfig `xml:"cookie-config,omitempty"`
	TrackingModes  []string      `xml:"tracking-mode"`
}

// CookieConfig represents session cookie configuration
type CookieConfig struct {
	XMLName  xml.Name `xml:"cookie-config"`
	Name     string   `xml:"name,omitempty"`
	Domain   string   `xml:"domain,omitempty"`
	Path     string   `xml:"path,omitempty"`
	Comment  string   `xml:"comment,omitempty"`
	HttpOnly string   `xml:"http-only,omitempty"`
	Secure   string   `xml:"secure,omitempty"`
	MaxAge   int      `xml:"max-age,omitempty"`
}

// MimeMapping represents a MIME type mapping
type MimeMapping struct {
	XMLName   xml.Name `xml:"mime-mapping"`
	Extension string   `xml:"extension"`
	MimeType  string   `xml:"mime-type"`
}

// WelcomeFileList represents welcome file configuration
type WelcomeFileList struct {
	XMLName      xml.Name `xml:"welcome-file-list"`
	WelcomeFiles []string `xml:"welcome-file"`
}

// ErrorPage represents an error page mapping
type ErrorPage struct {
	XMLName       xml.Name `xml:"error-page"`
	ErrorCode     string   `xml:"error-code,omitempty"`
	ExceptionType string   `xml:"exception-type,omitempty"`
	Location      string   `xml:"location"`
}

// JspConfig represents JSP configuration
type JspConfig struct {
	XMLName           xml.Name           `xml:"jsp-config"`
	TagLibs           []TagLib           `xml:"taglib"`
	JspPropertyGroups []JspPropertyGroup `xml:"jsp-property-group"`
}

// TagLib represents a taglib definition
type TagLib struct {
	XMLName        xml.Name `xml:"taglib"`
	TaglibURI      string   `xml:"taglib-uri"`
	TaglibLocation string   `xml:"taglib-location"`
}

// JspPropertyGroup represents JSP property group
type JspPropertyGroup struct {
	XMLName               xml.Name `xml:"jsp-property-group"`
	Description           string   `xml:"description,omitempty"`
	DisplayName           string   `xml:"display-name,omitempty"`
	URLPatterns           []string `xml:"url-pattern"`
	ElIgnored             string   `xml:"el-ignored,omitempty"`
	PageEncoding          string   `xml:"page-encoding,omitempty"`
	ScriptingInvalid      string   `xml:"scripting-invalid,omitempty"`
	IsXml                 string   `xml:"is-xml,omitempty"`
	IncludePreludes       []string `xml:"include-prelude"`
	IncludeCodas          []string `xml:"include-coda"`
	DeferredSyntaxAllowed string   `xml:"deferred-syntax-allowed-as-literal,omitempty"`
	TrimWhitespace        string   `xml:"trim-directive-whitespaces,omitempty"`
	DefaultContentType    string   `xml:"default-content-type,omitempty"`
	Buffer                string   `xml:"buffer,omitempty"`
	ErrorOnUndeclaredNS   string   `xml:"error-on-undeclared-namespace,omitempty"`
}

// SecurityConstraint represents a security constraint
type SecurityConstraint struct {
	XMLName                xml.Name                `xml:"security-constraint"`
	DisplayName            string                  `xml:"display-name,omitempty"`
	WebResourceCollections []WebResourceCollection `xml:"web-resource-collection"`
	AuthConstraint         *AuthConstraint         `xml:"auth-constraint,omitempty"`
	UserDataConstraint     *UserDataConstraint     `xml:"user-data-constraint,omitempty"`
}

// WebResourceCollection represents a collection of web resources
type WebResourceCollection struct {
	XMLName             xml.Name `xml:"web-resource-collection"`
	WebResourceName     string   `xml:"web-resource-name"`
	Description         string   `xml:"description,omitempty"`
	URLPatterns         []string `xml:"url-pattern"`
	HTTPMethods         []string `xml:"http-method"`
	HTTPMethodOmissions []string `xml:"http-method-omission"`
}

// AuthConstraint represents authorization constraints
type AuthConstraint struct {
	XMLName     xml.Name `xml:"auth-constraint"`
	Description string   `xml:"description,omitempty"`
	RoleNames   []string `xml:"role-name"`
}

// UserDataConstraint represents transport guarantee
type UserDataConstraint struct {
	XMLName            xml.Name `xml:"user-data-constraint"`
	Description        string   `xml:"description,omitempty"`
	TransportGuarantee string   `xml:"transport-guarantee"`
}

// LoginConfig represents login configuration
type LoginConfig struct {
	XMLName         xml.Name         `xml:"login-config"`
	AuthMethod      string           `xml:"auth-method,omitempty"`
	RealmName       string           `xml:"realm-name,omitempty"`
	FormLoginConfig *FormLoginConfig `xml:"form-login-config,omitempty"`
}

// FormLoginConfig represents form-based login configuration
type FormLoginConfig struct {
	XMLName       xml.Name `xml:"form-login-config"`
	FormLoginPage string   `xml:"form-login-page"`
	FormErrorPage string   `xml:"form-error-page"`
}

// SecurityRole represents a security role
type SecurityRole struct {
	XMLName     xml.Name `xml:"security-role"`
	RoleName    string   `xml:"role-name"`
	Description string   `xml:"description,omitempty"`
}

// ResourceRef represents a resource reference
type ResourceRef struct {
	XMLName         xml.Name `xml:"resource-ref"`
	Description     string   `xml:"description,omitempty"`
	ResRefName      string   `xml:"res-ref-name"`
	ResType         string   `xml:"res-type"`
	ResAuth         string   `xml:"res-auth,omitempty"`
	ResSharingScope string   `xml:"res-sharing-scope,omitempty"`
}

// ResourceEnvRef represents a resource environment reference
type ResourceEnvRef struct {
	XMLName            xml.Name `xml:"resource-env-ref"`
	Description        string   `xml:"description,omitempty"`
	ResourceEnvRefName string   `xml:"resource-env-ref-name"`
	ResourceEnvRefType string   `xml:"resource-env-ref-type"`
}

// MessageDestRef represents a message destination reference
type MessageDestRef struct {
	XMLName            xml.Name `xml:"message-destination-ref"`
	Description        string   `xml:"description,omitempty"`
	MessageDestRefName string   `xml:"message-destination-ref-name"`
	MessageDestType    string   `xml:"message-destination-type"`
	MessageDestUsage   string   `xml:"message-destination-usage,omitempty"`
	MessageDestLink    string   `xml:"message-destination-link,omitempty"`
}

// AbsoluteOrdering represents fragment ordering
type AbsoluteOrdering struct {
	XMLName xml.Name `xml:"absolute-ordering"`
	Names   []string `xml:"name"`
	Others  *Others  `xml:"others,omitempty"`
}

// Others represents the <others/> element in ordering
type Others struct {
	XMLName xml.Name `xml:"others"`
}

// DenyUncoveredMethods element
type DenyUncoveredMethods struct {
	XMLName xml.Name `xml:"deny-uncovered-http-methods"`
}

// Authentication methods
const (
	AuthMethodBasic      = "BASIC"
	AuthMethodDigest     = "DIGEST"
	AuthMethodForm       = "FORM"
	AuthMethodClientCert = "CLIENT-CERT"
)

// Transport guarantees
const (
	TransportNone         = "NONE"
	TransportIntegral     = "INTEGRAL"
	TransportConfidential = "CONFIDENTIAL"
)

// Dispatcher types
const (
	DispatcherRequest = "REQUEST"
	DispatcherForward = "FORWARD"
	DispatcherInclude = "INCLUDE"
	DispatcherError   = "ERROR"
	DispatcherAsync   = "ASYNC"
)

// Session tracking modes
const (
	TrackingCookie = "COOKIE"
	TrackingURL    = "URL"
	TrackingSSL    = "SSL"
)

// Common servlet classes
var CommonServletClasses = []string{
	"org.apache.catalina.servlets.DefaultServlet",
	"org.apache.jasper.servlet.JspServlet",
	"org.apache.catalina.servlets.CGIServlet",
	"org.apache.catalina.servlets.SSIServlet",
	"org.apache.catalina.servlets.WebdavServlet",
}

// Common filter classes
var CommonFilterClasses = []string{
	"org.apache.catalina.filters.AddDefaultCharsetFilter",
	"org.apache.catalina.filters.CorsFilter",
	"org.apache.catalina.filters.CsrfPreventionFilter",
	"org.apache.catalina.filters.ExpiresFilter",
	"org.apache.catalina.filters.FailedRequestFilter",
	"org.apache.catalina.filters.HttpHeaderSecurityFilter",
	"org.apache.catalina.filters.RemoteAddrFilter",
	"org.apache.catalina.filters.RemoteHostFilter",
	"org.apache.catalina.filters.RemoteIpFilter",
	"org.apache.catalina.filters.RequestDumperFilter",
	"org.apache.catalina.filters.RestCsrfPreventionFilter",
	"org.apache.catalina.filters.SetCharacterEncodingFilter",
	"org.apache.tomcat.websocket.server.WsFilter",
}

// Common listener classes
var CommonListenerClasses = []string{
	"org.apache.catalina.core.AprLifecycleListener",
	"org.apache.catalina.core.JreMemoryLeakPreventionListener",
	"org.apache.catalina.core.ThreadLocalLeakPreventionListener",
	"org.apache.catalina.mbeans.GlobalResourcesLifecycleListener",
	"org.apache.catalina.security.SecurityListener",
}

// Common MIME mappings
var CommonMimeMappings = map[string]string{
	"html":  "text/html",
	"htm":   "text/html",
	"css":   "text/css",
	"js":    "application/javascript",
	"json":  "application/json",
	"xml":   "application/xml",
	"txt":   "text/plain",
	"jpg":   "image/jpeg",
	"jpeg":  "image/jpeg",
	"png":   "image/png",
	"gif":   "image/gif",
	"svg":   "image/svg+xml",
	"ico":   "image/x-icon",
	"pdf":   "application/pdf",
	"zip":   "application/zip",
	"woff":  "font/woff",
	"woff2": "font/woff2",
	"ttf":   "font/ttf",
	"eot":   "application/vnd.ms-fontobject",
}

// NewWebApp creates a default web application configuration
func NewWebApp() *WebApp {
	return &WebApp{
		XMLNS:          "http://xmlns.jcp.org/xml/ns/javaee",
		XSI:            "http://www.w3.org/2001/XMLSchema-instance",
		SchemaLocation: "http://xmlns.jcp.org/xml/ns/javaee http://xmlns.jcp.org/xml/ns/javaee/web-app_4_0.xsd",
		Version:        "4.0",
		WelcomeFileList: &WelcomeFileList{
			WelcomeFiles: []string{"index.html", "index.htm", "index.jsp"},
		},
		SessionConfig: &SessionConfig{
			SessionTimeout: 30,
		},
	}
}

// NewServlet creates a new servlet with defaults
func NewServlet(name, class string) *Servlet {
	return &Servlet{
		ServletName:  name,
		ServletClass: class,
	}
}

// NewFilter creates a new filter with defaults
func NewFilter(name, class string) *Filter {
	return &Filter{
		FilterName:  name,
		FilterClass: class,
	}
}

// NewSecurityConstraint creates a new security constraint
func NewSecurityConstraint(name string, urlPatterns []string) *SecurityConstraint {
	return &SecurityConstraint{
		WebResourceCollections: []WebResourceCollection{
			{
				WebResourceName: name,
				URLPatterns:     urlPatterns,
			},
		},
	}
}

// NewLoginConfig creates a login configuration
func NewLoginConfig(authMethod string) *LoginConfig {
	return &LoginConfig{
		AuthMethod: authMethod,
	}
}

// NewFormLoginConfig creates a form-based login configuration
func NewFormLoginConfig(loginPage, errorPage string) *LoginConfig {
	return &LoginConfig{
		AuthMethod: AuthMethodForm,
		FormLoginConfig: &FormLoginConfig{
			FormLoginPage: loginPage,
			FormErrorPage: errorPage,
		},
	}
}

// NewErrorPage creates an error page mapping
func NewErrorPage(errorCode, location string) *ErrorPage {
	return &ErrorPage{
		ErrorCode: errorCode,
		Location:  location,
	}
}

// NewExceptionErrorPage creates an exception error page mapping
func NewExceptionErrorPage(exceptionType, location string) *ErrorPage {
	return &ErrorPage{
		ExceptionType: exceptionType,
		Location:      location,
	}
}
