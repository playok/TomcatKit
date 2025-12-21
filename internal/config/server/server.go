package server

// Server represents the root server.xml element
type Server struct {
	Port      int                    `xml:"port,attr"`
	Shutdown  string                 `xml:"shutdown,attr"`
	Listeners []Listener             `xml:"Listener"`
	Resources *GlobalNamingResources `xml:"GlobalNamingResources"`
	Services  []Service              `xml:"Service"`
}

// Listener represents a lifecycle listener
type Listener struct {
	ClassName string `xml:"className,attr"`
	// Common listener attributes
	SSLEngine     string `xml:"SSLEngine,attr,omitempty"`
	SSLRandomSeed string `xml:"SSLRandomSeed,attr,omitempty"`
}

// GlobalNamingResources represents global JNDI resources
type GlobalNamingResources struct {
	Resources []Resource `xml:"Resource"`
}

// Resource represents a JNDI resource
type Resource struct {
	Name        string `xml:"name,attr"`
	Auth        string `xml:"auth,attr"`
	Type        string `xml:"type,attr"`
	Description string `xml:"description,attr,omitempty"`
	Factory     string `xml:"factory,attr,omitempty"`
	Pathname    string `xml:"pathname,attr,omitempty"`
}

// Service represents a Tomcat service
type Service struct {
	Name       string      `xml:"name,attr"`
	Executors  []Executor  `xml:"Executor"`
	Connectors []Connector `xml:"Connector"`
	Engine     Engine      `xml:"Engine"`
}

// Executor represents a thread pool executor
type Executor struct {
	ClassName       string `xml:"className,attr,omitempty"`
	Name            string `xml:"name,attr"`
	NamePrefix      string `xml:"namePrefix,attr,omitempty"`
	MaxThreads      int    `xml:"maxThreads,attr,omitempty"`
	MinSpareThreads int    `xml:"minSpareThreads,attr,omitempty"`
	MaxIdleTime     int    `xml:"maxIdleTime,attr,omitempty"`
	// Virtual Thread specific (Tomcat 11+)
	MaxQueueSize int `xml:"maxQueueSize,attr,omitempty"`
}

// Executor class names
const (
	ExecutorStandardThreadPool = "org.apache.catalina.core.StandardThreadPoolExecutor"
	ExecutorVirtualThread      = "org.apache.catalina.core.StandardVirtualThreadExecutor"
)

// NewStandardExecutor creates a standard thread pool executor
func NewStandardExecutor(name string) *Executor {
	return &Executor{
		ClassName:       ExecutorStandardThreadPool,
		Name:            name,
		NamePrefix:      "catalina-exec-",
		MaxThreads:      200,
		MinSpareThreads: 25,
		MaxIdleTime:     60000,
	}
}

// NewVirtualThreadExecutor creates a virtual thread executor (Java 21+, Tomcat 11+)
func NewVirtualThreadExecutor(name string) *Executor {
	return &Executor{
		ClassName:    ExecutorVirtualThread,
		Name:         name,
		NamePrefix:   "catalina-virt-",
		MaxQueueSize: 100,
	}
}

// IsVirtualThread returns true if this is a virtual thread executor
func (e *Executor) IsVirtualThread() bool {
	return e.ClassName == ExecutorVirtualThread
}

// Connector represents HTTP/AJP connector
type Connector struct {
	Port              int    `xml:"port,attr"`
	Protocol          string `xml:"protocol,attr,omitempty"`
	ConnectionTimeout int    `xml:"connectionTimeout,attr,omitempty"`
	RedirectPort      int    `xml:"redirectPort,attr,omitempty"`
	// Common attributes
	Address string `xml:"address,attr,omitempty"`
	// HTTP specific
	MaxThreads      int `xml:"maxThreads,attr,omitempty"`
	MinSpareThreads int `xml:"minSpareThreads,attr,omitempty"`
	AcceptCount     int `xml:"acceptCount,attr,omitempty"`
	// SSL specific
	SSLEnabled   bool   `xml:"SSLEnabled,attr,omitempty"`
	Scheme       string `xml:"scheme,attr,omitempty"`
	Secure       bool   `xml:"secure,attr,omitempty"`
	KeystoreFile string `xml:"keystoreFile,attr,omitempty"`
	KeystorePass string `xml:"keystorePass,attr,omitempty"`
	KeystoreType string `xml:"keystoreType,attr,omitempty"`
	ClientAuth   string `xml:"clientAuth,attr,omitempty"`
	SSLProtocol  string `xml:"sslProtocol,attr,omitempty"`
	// AJP specific
	SecretRequired                  bool   `xml:"secretRequired,attr,omitempty"`
	Secret                          string `xml:"secret,attr,omitempty"`
	AllowedRequestAttributesPattern string `xml:"allowedRequestAttributesPattern,attr,omitempty"`
	// Executor reference
	Executor string `xml:"executor,attr,omitempty"`
	// Virtual Thread (Tomcat 11+ direct support)
	UseVirtualThreads string `xml:"useVirtualThreads,attr,omitempty"`
	// Compression settings
	Compression             string `xml:"compression,attr,omitempty"`
	CompressionMinSize      int    `xml:"compressionMinSize,attr,omitempty"`
	CompressibleMimeType    string `xml:"compressibleMimeType,attr,omitempty"`
	NoCompressionUserAgents string `xml:"noCompressionUserAgents,attr,omitempty"`
	// Nested SSL configuration
	SSLHostConfig *SSLHostConfig `xml:"SSLHostConfig,omitempty"`
}

// SSLHostConfig represents SSL host configuration
type SSLHostConfig struct {
	Protocols               string        `xml:"protocols,attr,omitempty"`
	CertificateVerification string        `xml:"certificateVerification,attr,omitempty"`
	Certificates            []Certificate `xml:"Certificate"`
}

// Certificate represents an SSL certificate
type Certificate struct {
	CertificateFile             string `xml:"certificateFile,attr,omitempty"`
	CertificateKeyFile          string `xml:"certificateKeyFile,attr,omitempty"`
	CertificateKeystoreFile     string `xml:"certificateKeystoreFile,attr,omitempty"`
	CertificateKeystorePassword string `xml:"certificateKeystorePassword,attr,omitempty"`
	Type                        string `xml:"type,attr,omitempty"`
}

// Engine represents the Catalina engine
type Engine struct {
	Name        string   `xml:"name,attr"`
	DefaultHost string   `xml:"defaultHost,attr"`
	JvmRoute    string   `xml:"jvmRoute,attr,omitempty"`
	Realm       *Realm   `xml:"Realm,omitempty"`
	Hosts       []Host   `xml:"Host"`
	Cluster     *Cluster `xml:"Cluster,omitempty"`
	Valves      []Valve  `xml:"Valve"`
}

// Host represents a virtual host
type Host struct {
	Name             string `xml:"name,attr"`
	AppBase          string `xml:"appBase,attr"`
	UnpackWARs       bool   `xml:"unpackWARs,attr,omitempty"`
	AutoDeploy       bool   `xml:"autoDeploy,attr,omitempty"`
	DeployOnStartup  bool   `xml:"deployOnStartup,attr,omitempty"`
	CreateDirs       bool   `xml:"createDirs,attr,omitempty"`
	DeployXML        bool   `xml:"deployXML,attr,omitempty"`
	CopyXML          bool   `xml:"copyXML,attr,omitempty"`
	WorkDir          string `xml:"workDir,attr,omitempty"`
	DeployIgnore     string `xml:"deployIgnore,attr,omitempty"`
	StartStopThreads int    `xml:"startStopThreads,attr,omitempty"`
	// Error handling
	ErrorReportValveClass string `xml:"errorReportValveClass,attr,omitempty"`
	// Nested elements
	Aliases  []Alias   `xml:"Alias"`
	Valves   []Valve   `xml:"Valve"`
	Contexts []Context `xml:"Context"`
	Realm    *Realm    `xml:"Realm,omitempty"`
}

// Alias represents a host alias (alternative hostname)
type Alias struct {
	Name string `xml:",chardata"`
}

// Context represents a web application context
type Context struct {
	Path         string `xml:"path,attr"`
	DocBase      string `xml:"docBase,attr"`
	Reloadable   bool   `xml:"reloadable,attr,omitempty"`
	CrossContext bool   `xml:"crossContext,attr,omitempty"`
	Privileged   bool   `xml:"privileged,attr,omitempty"`
	// Session configuration
	Cookies                            bool   `xml:"cookies,attr,omitempty"`
	SessionCookieName                  string `xml:"sessionCookieName,attr,omitempty"`
	SessionCookiePath                  string `xml:"sessionCookiePath,attr,omitempty"`
	SessionCookieDomain                string `xml:"sessionCookieDomain,attr,omitempty"`
	SessionCookiePathUsesTrailingSlash bool   `xml:"sessionCookiePathUsesTrailingSlash,attr,omitempty"`
	UseHttpOnly                        bool   `xml:"useHttpOnly,attr,omitempty"`
	// Resource handling
	AntiResourceLocking bool `xml:"antiResourceLocking,attr,omitempty"`
	SwallowOutput       bool `xml:"swallowOutput,attr,omitempty"`
	// Override settings
	Override                    bool `xml:"override,attr,omitempty"`
	AllowCasualMultipartParsing bool `xml:"allowCasualMultipartParsing,attr,omitempty"`
	// Cache settings
	CachingAllowed bool `xml:"cachingAllowed,attr,omitempty"`
	CacheMaxSize   int  `xml:"cacheMaxSize,attr,omitempty"`
	CacheTTL       int  `xml:"cacheTTL,attr,omitempty"`
	// Nested elements
	Resources        []Resource    `xml:"Resource"`
	Environments     []Environment `xml:"Environment"`
	Parameters       []Parameter   `xml:"Parameter"`
	Valves           []Valve       `xml:"Valve"`
	Realm            *Realm        `xml:"Realm,omitempty"`
	Manager          *Manager      `xml:"Manager,omitempty"`
	Loader           *Loader       `xml:"Loader,omitempty"`
	WatchedResources []string      `xml:"WatchedResource"`
}

// Parameter represents a context init parameter
type Parameter struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	Override    bool   `xml:"override,attr,omitempty"`
	Description string `xml:"description,attr,omitempty"`
}

// Manager represents a session manager
type Manager struct {
	ClassName               string `xml:"className,attr,omitempty"`
	MaxActiveSessions       int    `xml:"maxActiveSessions,attr,omitempty"`
	SessionIdLength         int    `xml:"sessionIdLength,attr,omitempty"`
	MaxInactiveInterval     int    `xml:"maxInactiveInterval,attr,omitempty"`
	Pathname                string `xml:"pathname,attr,omitempty"`
	ProcessExpiresFrequency int    `xml:"processExpiresFrequency,attr,omitempty"`
	SecureRandomClass       string `xml:"secureRandomClass,attr,omitempty"`
	SecureRandomAlgorithm   string `xml:"secureRandomAlgorithm,attr,omitempty"`
}

// Loader represents a web application class loader
type Loader struct {
	ClassName          string `xml:"className,attr,omitempty"`
	Delegate           bool   `xml:"delegate,attr,omitempty"`
	Reloadable         bool   `xml:"reloadable,attr,omitempty"`
	SearchVirtualFirst bool   `xml:"searchVirtualFirst,attr,omitempty"`
}

// Environment represents an environment entry
type Environment struct {
	Name     string `xml:"name,attr"`
	Value    string `xml:"value,attr"`
	Type     string `xml:"type,attr"`
	Override bool   `xml:"override,attr,omitempty"`
}

// Realm represents authentication realm
type Realm struct {
	ClassName string `xml:"className,attr"`
	// Common attributes
	ResourceName string `xml:"resourceName,attr,omitempty"`
	// DataSourceRealm
	DataSourceName string `xml:"dataSourceName,attr,omitempty"`
	UserTable      string `xml:"userTable,attr,omitempty"`
	UserNameCol    string `xml:"userNameCol,attr,omitempty"`
	UserCredCol    string `xml:"userCredCol,attr,omitempty"`
	UserRoleTable  string `xml:"userRoleTable,attr,omitempty"`
	RoleNameCol    string `xml:"roleNameCol,attr,omitempty"`
	// JNDIRealm
	ConnectionURL      string `xml:"connectionURL,attr,omitempty"`
	ConnectionName     string `xml:"connectionName,attr,omitempty"`
	ConnectionPassword string `xml:"connectionPassword,attr,omitempty"`
	UserPattern        string `xml:"userPattern,attr,omitempty"`
	UserBase           string `xml:"userBase,attr,omitempty"`
	UserSearch         string `xml:"userSearch,attr,omitempty"`
	RoleBase           string `xml:"roleBase,attr,omitempty"`
	RoleName           string `xml:"roleName,attr,omitempty"`
	RoleSearch         string `xml:"roleSearch,attr,omitempty"`
	// Nested realms (CombinedRealm, LockOutRealm)
	NestedRealms []Realm `xml:"Realm,omitempty"`
	// CredentialHandler
	CredentialHandler *CredentialHandler `xml:"CredentialHandler,omitempty"`
}

// CredentialHandler represents password hashing configuration
type CredentialHandler struct {
	ClassName  string `xml:"className,attr"`
	Algorithm  string `xml:"algorithm,attr,omitempty"`
	Iterations int    `xml:"iterations,attr,omitempty"`
	SaltLength int    `xml:"saltLength,attr,omitempty"`
}

// Valve represents a request processing valve
type Valve struct {
	ClassName string `xml:"className,attr"`

	// AccessLogValve attributes
	Directory                string `xml:"directory,attr,omitempty"`
	Prefix                   string `xml:"prefix,attr,omitempty"`
	Suffix                   string `xml:"suffix,attr,omitempty"`
	Pattern                  string `xml:"pattern,attr,omitempty"`
	Rotatable                bool   `xml:"rotatable,attr,omitempty"`
	RenameOnRotate           bool   `xml:"renameOnRotate,attr,omitempty"`
	FileDateFormat           string `xml:"fileDateFormat,attr,omitempty"`
	Encoding                 string `xml:"encoding,attr,omitempty"`
	Locale                   string `xml:"locale,attr,omitempty"`
	RequestAttributesEnabled bool   `xml:"requestAttributesEnabled,attr,omitempty"`
	Buffered                 bool   `xml:"buffered,attr,omitempty"`
	MaxLogMessageBufferSize  int    `xml:"maxLogMessageBufferSize,attr,omitempty"`
	ConditionIf              string `xml:"conditionIf,attr,omitempty"`
	ConditionUnless          string `xml:"conditionUnless,attr,omitempty"`

	// RemoteAddrValve / RemoteCIDRValve / RemoteHostValve
	Allow                         string `xml:"allow,attr,omitempty"`
	Deny                          string `xml:"deny,attr,omitempty"`
	DenyStatus                    int    `xml:"denyStatus,attr,omitempty"`
	AddConnectorPort              bool   `xml:"addConnectorPort,attr,omitempty"`
	InvalidAuthenticationWhenDeny bool   `xml:"invalidAuthenticationWhenDeny,attr,omitempty"`

	// RemoteIpValve
	RemoteIpHeader           string `xml:"remoteIpHeader,attr,omitempty"`
	ProtocolHeader           string `xml:"protocolHeader,attr,omitempty"`
	ProtocolHeaderHttpsValue string `xml:"protocolHeaderHttpsValue,attr,omitempty"`
	PortHeader               string `xml:"portHeader,attr,omitempty"`
	ProxiesHeader            string `xml:"proxiesHeader,attr,omitempty"`
	RemoteIpProxiesHeader    string `xml:"remoteIpProxiesHeader,attr,omitempty"`
	InternalProxies          string `xml:"internalProxies,attr,omitempty"`
	TrustedProxies           string `xml:"trustedProxies,attr,omitempty"`
	ChangeLocalPort          bool   `xml:"changeLocalPort,attr,omitempty"`
	ChangeLocalName          bool   `xml:"changeLocalName,attr,omitempty"`

	// ErrorReportValve
	ShowServerInfo bool `xml:"showServerInfo,attr,omitempty"`
	ShowReport     bool `xml:"showReport,attr,omitempty"`

	// SingleSignOn
	CookieDomain            string `xml:"cookieDomain,attr,omitempty"`
	CookieName              string `xml:"cookieName,attr,omitempty"`
	RequireReauthentication bool   `xml:"requireReauthentication,attr,omitempty"`

	// ReplicationValve
	Filter               string `xml:"filter,attr,omitempty"`
	PrimaryIndicator     bool   `xml:"primaryIndicator,attr,omitempty"`
	PrimaryIndicatorName string `xml:"primaryIndicatorName,attr,omitempty"`

	// StuckThreadDetectionValve
	Threshold                int `xml:"threshold,attr,omitempty"`
	InterruptThreadThreshold int `xml:"interruptThreadThreshold,attr,omitempty"`

	// CrawlerSessionManagerValve
	CrawlerUserAgents       string `xml:"crawlerUserAgents,attr,omitempty"`
	SessionInactiveInterval int    `xml:"sessionInactiveInterval,attr,omitempty"`

	// SemaphoreValve
	Concurrency int  `xml:"concurrency,attr,omitempty"`
	Fairness    bool `xml:"fairness,attr,omitempty"`
	Block       bool `xml:"block,attr,omitempty"`

	// AuthenticatorValve common
	AlwaysUseSession                bool   `xml:"alwaysUseSession,attr,omitempty"`
	Cache                           bool   `xml:"cache,attr,omitempty"`
	CacheSize                       int    `xml:"cacheSize,attr,omitempty"`
	ChangeSessionIdOnAuthentication bool   `xml:"changeSessionIdOnAuthentication,attr,omitempty"`
	DisableProxyCaching             bool   `xml:"disableProxyCaching,attr,omitempty"`
	SecurePagesWithPragma           bool   `xml:"securePagesWithPragma,attr,omitempty"`
	SecureRandomClass               string `xml:"secureRandomClass,attr,omitempty"`
	SecureRandomAlgorithm           string `xml:"secureRandomAlgorithm,attr,omitempty"`
}

// Valve class names
const (
	ValveAccessLog             = "org.apache.catalina.valves.AccessLogValve"
	ValveExtendedAccessLog     = "org.apache.catalina.valves.ExtendedAccessLogValve"
	ValveRemoteAddr            = "org.apache.catalina.valves.RemoteAddrValve"
	ValveRemoteCIDR            = "org.apache.catalina.valves.RemoteCIDRValve"
	ValveRemoteHost            = "org.apache.catalina.valves.RemoteHostValve"
	ValveRemoteIp              = "org.apache.catalina.valves.RemoteIpValve"
	ValveErrorReport           = "org.apache.catalina.valves.ErrorReportValve"
	ValveSingleSignOn          = "org.apache.catalina.authenticator.SingleSignOn"
	ValveReplication           = "org.apache.catalina.ha.tcp.ReplicationValve"
	ValveStuckThreadDetection  = "org.apache.catalina.valves.StuckThreadDetectionValve"
	ValveCrawlerSessionManager = "org.apache.catalina.valves.CrawlerSessionManagerValve"
	ValveRequestDumper         = "org.apache.catalina.valves.RequestDumperValve"
	ValveSemaphore             = "org.apache.catalina.valves.SemaphoreValve"
	ValveHealthCheck           = "org.apache.catalina.valves.HealthCheckValve"
	ValvePersistentValve       = "org.apache.catalina.valves.PersistentValve"
)

// GetValveDescription returns a human-readable description for a valve
func GetValveDescription(className string) string {
	descriptions := map[string]string{
		ValveAccessLog:             "Logs access information in a customizable format",
		ValveExtendedAccessLog:     "Extended W3C format access logging",
		ValveRemoteAddr:            "IP address-based access control",
		ValveRemoteCIDR:            "CIDR-based access control",
		ValveRemoteHost:            "Hostname-based access control",
		ValveRemoteIp:              "Handles X-Forwarded-For and similar headers from proxies",
		ValveErrorReport:           "Customizes error page generation",
		ValveSingleSignOn:          "Enables single sign-on across web applications",
		ValveReplication:           "Triggers session replication in cluster",
		ValveStuckThreadDetection:  "Detects and reports stuck request threads",
		ValveCrawlerSessionManager: "Special session handling for web crawlers",
		ValveRequestDumper:         "Debug valve that logs request/response details",
		ValveSemaphore:             "Limits concurrent access to a resource",
		ValveHealthCheck:           "Provides health check endpoint",
		ValvePersistentValve:       "Handles persistent sessions",
	}
	if desc, ok := descriptions[className]; ok {
		return desc
	}
	return className
}

// GetValveShortName returns a short name for the valve
func GetValveShortName(className string) string {
	names := map[string]string{
		ValveAccessLog:             "AccessLogValve",
		ValveExtendedAccessLog:     "ExtendedAccessLogValve",
		ValveRemoteAddr:            "RemoteAddrValve",
		ValveRemoteCIDR:            "RemoteCIDRValve",
		ValveRemoteHost:            "RemoteHostValve",
		ValveRemoteIp:              "RemoteIpValve",
		ValveErrorReport:           "ErrorReportValve",
		ValveSingleSignOn:          "SingleSignOn",
		ValveReplication:           "ReplicationValve",
		ValveStuckThreadDetection:  "StuckThreadDetectionValve",
		ValveCrawlerSessionManager: "CrawlerSessionManagerValve",
		ValveRequestDumper:         "RequestDumperValve",
		ValveSemaphore:             "SemaphoreValve",
		ValveHealthCheck:           "HealthCheckValve",
		ValvePersistentValve:       "PersistentValve",
	}
	if name, ok := names[className]; ok {
		return name
	}
	// Extract class name from full path
	for i := len(className) - 1; i >= 0; i-- {
		if className[i] == '.' {
			return className[i+1:]
		}
	}
	return className
}

// AvailableValveTypes returns all available valve class names
func AvailableValveTypes() []string {
	return []string{
		ValveAccessLog,
		ValveRemoteAddr,
		ValveRemoteCIDR,
		ValveRemoteIp,
		ValveErrorReport,
		ValveSingleSignOn,
		ValveStuckThreadDetection,
		ValveCrawlerSessionManager,
		ValveRequestDumper,
		ValveSemaphore,
	}
}

// Common access log patterns
var AccessLogPatterns = map[string]string{
	"common":   "%h %l %u %t \"%r\" %s %b",
	"combined": "%h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-Agent}i\"",
	"default":  "%h %l %u %t \"%r\" %s %b",
}

// DefaultAccessLogValve creates a default access log valve
func DefaultAccessLogValve() Valve {
	return Valve{
		ClassName:      ValveAccessLog,
		Directory:      "logs",
		Prefix:         "localhost_access_log",
		Suffix:         ".txt",
		Pattern:        "%h %l %u %t \"%r\" %s %b",
		Rotatable:      true,
		FileDateFormat: ".yyyy-MM-dd",
	}
}

// DefaultRemoteAddrValve creates a default remote address valve
func DefaultRemoteAddrValve() Valve {
	return Valve{
		ClassName:  ValveRemoteAddr,
		Allow:      "127\\.\\d+\\.\\d+\\.\\d+|::1|0:0:0:0:0:0:0:1",
		DenyStatus: 403,
	}
}

// DefaultRemoteIpValve creates a default remote IP valve
func DefaultRemoteIpValve() Valve {
	return Valve{
		ClassName:       ValveRemoteIp,
		RemoteIpHeader:  "X-Forwarded-For",
		ProtocolHeader:  "X-Forwarded-Proto",
		InternalProxies: "10\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}|192\\.168\\.\\d{1,3}\\.\\d{1,3}|127\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}",
	}
}

// DefaultErrorReportValve creates a default error report valve
func DefaultErrorReportValve() Valve {
	return Valve{
		ClassName:      ValveErrorReport,
		ShowServerInfo: false,
		ShowReport:     true,
	}
}

// DefaultSingleSignOnValve creates a default SSO valve
func DefaultSingleSignOnValve() Valve {
	return Valve{
		ClassName: ValveSingleSignOn,
	}
}

// DefaultStuckThreadDetectionValve creates a default stuck thread detection valve
func DefaultStuckThreadDetectionValve() Valve {
	return Valve{
		ClassName: ValveStuckThreadDetection,
		Threshold: 600,
	}
}

// Cluster represents session clustering configuration
type Cluster struct {
	ClassName                        string            `xml:"className,attr"`
	ChannelSendOptions               string            `xml:"channelSendOptions,attr,omitempty"`
	ChannelStartOptions              string            `xml:"channelStartOptions,attr,omitempty"`
	NotifyLifecycleListenerOnFailure bool              `xml:"notifyLifecycleListenerOnFailure,attr,omitempty"`
	Manager                          *ClusterManager   `xml:"Manager,omitempty"`
	Channel                          *Channel          `xml:"Channel,omitempty"`
	Valves                           []Valve           `xml:"Valve"`
	ClusterListeners                 []ClusterListener `xml:"ClusterListener"`
	Deployer                         *FarmWarDeployer  `xml:"Deployer,omitempty"`
}

// ClusterManager represents cluster session manager
type ClusterManager struct {
	ClassName string `xml:"className,attr"`
	Name      string `xml:"name,attr,omitempty"`
	// Common attributes
	ExpireSessionsOnShutdown              bool `xml:"expireSessionsOnShutdown,attr,omitempty"`
	NotifyListenersOnReplication          bool `xml:"notifyListenersOnReplication,attr,omitempty"`
	NotifyContainerListenersOnReplication bool `xml:"notifyContainerListenersOnReplication,attr,omitempty"`
	NotifySessionListenersOnReplication   bool `xml:"notifySessionListenersOnReplication,attr,omitempty"`
	// DeltaManager specific
	StateTransferTimeout    int  `xml:"stateTransferTimeout,attr,omitempty"`
	SendAllSessions         bool `xml:"sendAllSessions,attr,omitempty"`
	SendAllSessionsSize     int  `xml:"sendAllSessionsSize,attr,omitempty"`
	SendAllSessionsWaitTime int  `xml:"sendAllSessionsWaitTime,attr,omitempty"`
	// BackupManager specific
	MapSendOptions          string `xml:"mapSendOptions,attr,omitempty"`
	RpcTimeout              int    `xml:"rpcTimeout,attr,omitempty"`
	TerminateOnStartFailure bool   `xml:"terminateOnStartFailure,attr,omitempty"`
	AccessTimeout           int    `xml:"accessTimeout,attr,omitempty"`
}

// Channel represents cluster communication channel
type Channel struct {
	ClassName    string        `xml:"className,attr"`
	Membership   *Membership   `xml:"Membership,omitempty"`
	Receiver     *Receiver     `xml:"Receiver,omitempty"`
	Sender       *Sender       `xml:"Sender,omitempty"`
	Interceptors []Interceptor `xml:"Interceptor"`
}

// Membership represents cluster membership service
type Membership struct {
	ClassName             string `xml:"className,attr"`
	Address               string `xml:"address,attr,omitempty"`
	Port                  int    `xml:"port,attr,omitempty"`
	Frequency             int    `xml:"frequency,attr,omitempty"`
	DropTime              int    `xml:"dropTime,attr,omitempty"`
	Bind                  string `xml:"bind,attr,omitempty"`
	RecoveryEnabled       bool   `xml:"recoveryEnabled,attr,omitempty"`
	RecoveryCounter       int    `xml:"recoveryCounter,attr,omitempty"`
	RecoverySleepTime     int    `xml:"recoverySleepTime,attr,omitempty"`
	LocalLoopbackDisabled bool   `xml:"localLoopbackDisabled,attr,omitempty"`
}

// Receiver represents cluster message receiver
type Receiver struct {
	ClassName       string `xml:"className,attr"`
	Address         string `xml:"address,attr,omitempty"`
	Port            int    `xml:"port,attr,omitempty"`
	AutoBind        int    `xml:"autoBind,attr,omitempty"`
	SelectorTimeout int    `xml:"selectorTimeout,attr,omitempty"`
	MaxThreads      int    `xml:"maxThreads,attr,omitempty"`
	MinThreads      int    `xml:"minThreads,attr,omitempty"`
	MaxIdleTime     int    `xml:"maxIdleTime,attr,omitempty"`
	OoBInline       bool   `xml:"ooBInline,attr,omitempty"`
	RxBufSize       int    `xml:"rxBufSize,attr,omitempty"`
	TxBufSize       int    `xml:"txBufSize,attr,omitempty"`
	UdpRxBufSize    int    `xml:"udpRxBufSize,attr,omitempty"`
	UdpTxBufSize    int    `xml:"udpTxBufSize,attr,omitempty"`
	SoTimeout       int    `xml:"soTimeout,attr,omitempty"`
	Timeout         int    `xml:"timeout,attr,omitempty"`
}

// Sender represents cluster message sender
type Sender struct {
	ClassName string     `xml:"className,attr"`
	Transport *Transport `xml:"Transport,omitempty"`
}

// Transport represents sender transport
type Transport struct {
	ClassName        string `xml:"className,attr"`
	RxBufSize        int    `xml:"rxBufSize,attr,omitempty"`
	TxBufSize        int    `xml:"txBufSize,attr,omitempty"`
	UdpRxBufSize     int    `xml:"udpRxBufSize,attr,omitempty"`
	UdpTxBufSize     int    `xml:"udpTxBufSize,attr,omitempty"`
	DirectBuffer     bool   `xml:"directBuffer,attr,omitempty"`
	KeepAliveCount   int    `xml:"keepAliveCount,attr,omitempty"`
	KeepAliveTime    int    `xml:"keepAliveTime,attr,omitempty"`
	Timeout          int    `xml:"timeout,attr,omitempty"`
	MaxRetryAttempts int    `xml:"maxRetryAttempts,attr,omitempty"`
	OoBInline        bool   `xml:"ooBInline,attr,omitempty"`
	SoKeepAlive      bool   `xml:"soKeepAlive,attr,omitempty"`
	SoLingerOn       bool   `xml:"soLingerOn,attr,omitempty"`
	SoLingerTime     int    `xml:"soLingerTime,attr,omitempty"`
	SoReuseAddress   bool   `xml:"soReuseAddress,attr,omitempty"`
	SoTrafficClass   int    `xml:"soTrafficClass,attr,omitempty"`
	TcpNoDelay       bool   `xml:"tcpNoDelay,attr,omitempty"`
	ThrowOnFailedAck bool   `xml:"throwOnFailedAck,attr,omitempty"`
}

// Interceptor represents channel interceptor
type Interceptor struct {
	ClassName string `xml:"className,attr"`
	// TcpFailureDetector
	ConnectTimeout        int  `xml:"connectTimeout,attr,omitempty"`
	PerformSendTest       bool `xml:"performSendTest,attr,omitempty"`
	PerformReadTest       bool `xml:"performReadTest,attr,omitempty"`
	ReadTestTimeout       int  `xml:"readTestTimeout,attr,omitempty"`
	RemoveSuspectsTimeout int  `xml:"removeSuspectsTimeout,attr,omitempty"`
	// MessageDispatchInterceptor
	MaxQueueSize  int  `xml:"maxQueueSize,attr,omitempty"`
	OptionalQueue bool `xml:"optionalQueue,attr,omitempty"`
	AlwaysSend    bool `xml:"alwaysSend,attr,omitempty"`
	// ThroughputInterceptor
	Interval int `xml:"interval,attr,omitempty"`
	// StaticMembershipInterceptor
	// (Members are defined as nested Member elements)
	// EncryptInterceptor
	EncryptionAlgorithm string `xml:"encryptionAlgorithm,attr,omitempty"`
	EncryptionKey       string `xml:"encryptionKey,attr,omitempty"`
	EncryptionKeyFile   string `xml:"encryptionKeyFile,attr,omitempty"`
	// GzipInterceptor (no additional attrs)
}

// ClusterListener represents cluster event listener
type ClusterListener struct {
	ClassName string `xml:"className,attr"`
}

// FarmWarDeployer represents cluster deployment
type FarmWarDeployer struct {
	ClassName              string `xml:"className,attr"`
	TempDir                string `xml:"tempDir,attr,omitempty"`
	DeployDir              string `xml:"deployDir,attr,omitempty"`
	WatchDir               string `xml:"watchDir,attr,omitempty"`
	WatchEnabled           bool   `xml:"watchEnabled,attr,omitempty"`
	ProcessDeployFrequency int    `xml:"processDeployFrequency,attr,omitempty"`
}

// Cluster class names
const (
	ClusterSimpleTcpCluster = "org.apache.catalina.ha.tcp.SimpleTcpCluster"
)

// Cluster manager class names
const (
	ClusterManagerDelta  = "org.apache.catalina.ha.session.DeltaManager"
	ClusterManagerBackup = "org.apache.catalina.ha.session.BackupManager"
)

// Channel class names
const (
	ChannelGroupChannel = "org.apache.catalina.tribes.group.GroupChannel"
)

// Membership class names
const (
	MembershipMcastService            = "org.apache.catalina.tribes.membership.McastService"
	MembershipStaticMembershipService = "org.apache.catalina.tribes.membership.StaticMembershipService"
)

// Receiver class names
const (
	ReceiverNioReceiver = "org.apache.catalina.tribes.transport.nio.NioReceiver"
	ReceiverBioReceiver = "org.apache.catalina.tribes.transport.bio.BioReceiver"
)

// Sender class names
const (
	SenderReplicationTransmitter = "org.apache.catalina.tribes.transport.ReplicationTransmitter"
)

// Transport class names
const (
	TransportPooledParallelSender = "org.apache.catalina.tribes.transport.nio.PooledParallelSender"
)

// Interceptor class names
const (
	InterceptorTcpFailureDetector = "org.apache.catalina.tribes.group.interceptors.TcpFailureDetector"
	InterceptorMessageDispatch    = "org.apache.catalina.tribes.group.interceptors.MessageDispatchInterceptor"
	InterceptorThroughput         = "org.apache.catalina.tribes.group.interceptors.ThroughputInterceptor"
	InterceptorStaticMembership   = "org.apache.catalina.tribes.group.interceptors.StaticMembershipInterceptor"
	InterceptorTcpPingInterceptor = "org.apache.catalina.tribes.group.interceptors.TcpPingInterceptor"
	InterceptorEncrypt            = "org.apache.catalina.tribes.group.interceptors.EncryptInterceptor"
	InterceptorGzip               = "org.apache.catalina.tribes.group.interceptors.GzipInterceptor"
	InterceptorFragmentation      = "org.apache.catalina.tribes.group.interceptors.FragmentationInterceptor"
	InterceptorOrderInterceptor   = "org.apache.catalina.tribes.group.interceptors.OrderInterceptor"
)

// Cluster listener class names
const (
	ClusterListenerSessionListener = "org.apache.catalina.ha.session.ClusterSessionListener"
)

// Deployer class names
const (
	DeployerFarmWarDeployer = "org.apache.catalina.ha.deploy.FarmWarDeployer"
)

// GetClusterManagerDescription returns description for cluster manager
func GetClusterManagerDescription(className string) string {
	descriptions := map[string]string{
		ClusterManagerDelta:  "All-to-all replication, simpler but higher network traffic",
		ClusterManagerBackup: "Primary-backup replication, more efficient for large clusters",
	}
	if desc, ok := descriptions[className]; ok {
		return desc
	}
	return className
}

// GetInterceptorDescription returns description for interceptor
func GetInterceptorDescription(className string) string {
	descriptions := map[string]string{
		InterceptorTcpFailureDetector: "Detects failed members using TCP",
		InterceptorMessageDispatch:    "Asynchronous message dispatch",
		InterceptorThroughput:         "Logs throughput statistics",
		InterceptorStaticMembership:   "Static cluster membership",
		InterceptorTcpPingInterceptor: "Periodic TCP pings to members",
		InterceptorEncrypt:            "Encrypts cluster communication",
		InterceptorGzip:               "Compresses cluster messages",
		InterceptorFragmentation:      "Fragments large messages",
		InterceptorOrderInterceptor:   "Ensures message ordering",
	}
	if desc, ok := descriptions[className]; ok {
		return desc
	}
	return className
}

// GetInterceptorShortName returns short name for interceptor
func GetInterceptorShortName(className string) string {
	names := map[string]string{
		InterceptorTcpFailureDetector: "TcpFailureDetector",
		InterceptorMessageDispatch:    "MessageDispatchInterceptor",
		InterceptorThroughput:         "ThroughputInterceptor",
		InterceptorStaticMembership:   "StaticMembershipInterceptor",
		InterceptorTcpPingInterceptor: "TcpPingInterceptor",
		InterceptorEncrypt:            "EncryptInterceptor",
		InterceptorGzip:               "GzipInterceptor",
		InterceptorFragmentation:      "FragmentationInterceptor",
		InterceptorOrderInterceptor:   "OrderInterceptor",
	}
	if name, ok := names[className]; ok {
		return name
	}
	// Extract from class name
	for i := len(className) - 1; i >= 0; i-- {
		if className[i] == '.' {
			return className[i+1:]
		}
	}
	return className
}

// AvailableInterceptorTypes returns all available interceptor class names
func AvailableInterceptorTypes() []string {
	return []string{
		InterceptorTcpFailureDetector,
		InterceptorMessageDispatch,
		InterceptorThroughput,
		InterceptorTcpPingInterceptor,
		InterceptorEncrypt,
		InterceptorGzip,
		InterceptorFragmentation,
		InterceptorOrderInterceptor,
	}
}

// DefaultCluster creates a default cluster configuration
func DefaultCluster() *Cluster {
	return &Cluster{
		ClassName:          ClusterSimpleTcpCluster,
		ChannelSendOptions: "8",
		Manager: &ClusterManager{
			ClassName:                    ClusterManagerDelta,
			ExpireSessionsOnShutdown:     false,
			NotifyListenersOnReplication: true,
		},
		Channel: &Channel{
			ClassName: ChannelGroupChannel,
			Membership: &Membership{
				ClassName: MembershipMcastService,
				Address:   "228.0.0.4",
				Port:      45564,
				Frequency: 500,
				DropTime:  3000,
			},
			Receiver: &Receiver{
				ClassName:       ReceiverNioReceiver,
				Address:         "auto",
				Port:            4000,
				AutoBind:        100,
				SelectorTimeout: 5000,
				MaxThreads:      6,
			},
			Sender: &Sender{
				ClassName: SenderReplicationTransmitter,
				Transport: &Transport{
					ClassName: TransportPooledParallelSender,
				},
			},
			Interceptors: []Interceptor{
				{ClassName: InterceptorTcpFailureDetector},
				{ClassName: InterceptorMessageDispatch},
			},
		},
		ClusterListeners: []ClusterListener{
			{ClassName: ClusterListenerSessionListener},
		},
	}
}
