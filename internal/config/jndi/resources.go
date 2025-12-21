package jndi

import "encoding/xml"

// ResourceType represents the type of JNDI resource
type ResourceType string

const (
	ResourceTypeDataSource   ResourceType = "javax.sql.DataSource"
	ResourceTypeMailSession  ResourceType = "javax.mail.Session"
	ResourceTypeUserDatabase ResourceType = "org.apache.catalina.UserDatabase"
	ResourceTypeBean         ResourceType = "java.lang.Object"
)

// Auth types
const (
	AuthContainer   = "Container"
	AuthApplication = "Application"
)

// Resource represents a generic JNDI resource in XML
type Resource struct {
	XMLName     xml.Name `xml:"Resource"`
	Name        string   `xml:"name,attr"`
	Auth        string   `xml:"auth,attr,omitempty"`
	Type        string   `xml:"type,attr"`
	Description string   `xml:"description,attr,omitempty"`
	Factory     string   `xml:"factory,attr,omitempty"`
	// UserDatabase specific
	Pathname string `xml:"pathname,attr,omitempty"`
	Readonly bool   `xml:"readonly,attr,omitempty"`
	// DataSource specific (DBCP2)
	DriverClassName               string `xml:"driverClassName,attr,omitempty"`
	URL                           string `xml:"url,attr,omitempty"`
	Username                      string `xml:"username,attr,omitempty"`
	Password                      string `xml:"password,attr,omitempty"`
	InitialSize                   int    `xml:"initialSize,attr,omitempty"`
	MaxTotal                      int    `xml:"maxTotal,attr,omitempty"`
	MaxIdle                       int    `xml:"maxIdle,attr,omitempty"`
	MinIdle                       int    `xml:"minIdle,attr,omitempty"`
	MaxWaitMillis                 int    `xml:"maxWaitMillis,attr,omitempty"`
	ValidationQuery               string `xml:"validationQuery,attr,omitempty"`
	ValidationQueryTimeout        int    `xml:"validationQueryTimeout,attr,omitempty"`
	TestOnBorrow                  bool   `xml:"testOnBorrow,attr,omitempty"`
	TestOnReturn                  bool   `xml:"testOnReturn,attr,omitempty"`
	TestWhileIdle                 bool   `xml:"testWhileIdle,attr,omitempty"`
	TimeBetweenEvictionRunsMillis int    `xml:"timeBetweenEvictionRunsMillis,attr,omitempty"`
	NumTestsPerEvictionRun        int    `xml:"numTestsPerEvictionRun,attr,omitempty"`
	MinEvictableIdleTimeMillis    int    `xml:"minEvictableIdleTimeMillis,attr,omitempty"`
	RemoveAbandonedOnBorrow       bool   `xml:"removeAbandonedOnBorrow,attr,omitempty"`
	RemoveAbandonedOnMaintenance  bool   `xml:"removeAbandonedOnMaintenance,attr,omitempty"`
	RemoveAbandonedTimeout        int    `xml:"removeAbandonedTimeout,attr,omitempty"`
	LogAbandoned                  bool   `xml:"logAbandoned,attr,omitempty"`
	DefaultAutoCommit             string `xml:"defaultAutoCommit,attr,omitempty"`
	DefaultReadOnly               string `xml:"defaultReadOnly,attr,omitempty"`
	DefaultTransactionIsolation   string `xml:"defaultTransactionIsolation,attr,omitempty"`
	PoolPreparedStatements        bool   `xml:"poolPreparedStatements,attr,omitempty"`
	MaxOpenPreparedStatements     int    `xml:"maxOpenPreparedStatements,attr,omitempty"`
	ConnectionInitSqls            string `xml:"connectionInitSqls,attr,omitempty"`
	ConnectionProperties          string `xml:"connectionProperties,attr,omitempty"`
	// Mail Session specific
	MailSmtpHost          string `xml:"mail.smtp.host,attr,omitempty"`
	MailSmtpPort          string `xml:"mail.smtp.port,attr,omitempty"`
	MailSmtpAuth          string `xml:"mail.smtp.auth,attr,omitempty"`
	MailSmtpStartTLS      string `xml:"mail.smtp.starttls.enable,attr,omitempty"`
	MailSmtpUser          string `xml:"mail.smtp.user,attr,omitempty"`
	MailTransportProtocol string `xml:"mail.transport.protocol,attr,omitempty"`
	MailDebug             string `xml:"mail.debug,attr,omitempty"`
}

// Environment represents an environment entry
type Environment struct {
	XMLName     xml.Name `xml:"Environment"`
	Name        string   `xml:"name,attr"`
	Value       string   `xml:"value,attr"`
	Type        string   `xml:"type,attr"`
	Override    bool     `xml:"override,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
}

// ResourceLink represents a link to a global resource
type ResourceLink struct {
	XMLName xml.Name `xml:"ResourceLink"`
	Name    string   `xml:"name,attr"`
	Global  string   `xml:"global,attr"`
	Type    string   `xml:"type,attr"`
}

// Context represents the context.xml structure
type Context struct {
	XMLName      xml.Name `xml:"Context"`
	Path         string   `xml:"path,attr,omitempty"`
	DocBase      string   `xml:"docBase,attr,omitempty"`
	Reloadable   bool     `xml:"reloadable,attr,omitempty"`
	CrossContext bool     `xml:"crossContext,attr,omitempty"`
	Privileged   bool     `xml:"privileged,attr,omitempty"`

	// Cookie settings
	Cookies             string `xml:"cookies,attr,omitempty"`
	UseHttpOnly         string `xml:"useHttpOnly,attr,omitempty"`
	SessionCookieName   string `xml:"sessionCookieName,attr,omitempty"`
	SessionCookiePath   string `xml:"sessionCookiePath,attr,omitempty"`
	SessionCookieDomain string `xml:"sessionCookieDomain,attr,omitempty"`

	// Resource locking
	AntiResourceLocking string `xml:"antiResourceLocking,attr,omitempty"`
	AntiJARLocking      string `xml:"antiJARLocking,attr,omitempty"`

	// Caching
	CachingAllowed string `xml:"cachingAllowed,attr,omitempty"`
	CacheMaxSize   int    `xml:"cacheMaxSize,attr,omitempty"`
	CacheTTL       int    `xml:"cacheTTL,attr,omitempty"`

	// Misc settings
	SwallowOutput                    string `xml:"swallowOutput,attr,omitempty"`
	MapperContextRootRedirectEnabled string `xml:"mapperContextRootRedirectEnabled,attr,omitempty"`
	MapperDirectoryRedirectEnabled   string `xml:"mapperDirectoryRedirectEnabled,attr,omitempty"`
	AllowCasualMultipartParsing      string `xml:"allowCasualMultipartParsing,attr,omitempty"`

	// Class loading
	Delegate                   string `xml:"delegate,attr,omitempty"`
	ParallelAnnotationScanning string `xml:"parallelAnnotationScanning,attr,omitempty"`

	// Nested elements
	Resources        []Resource         `xml:"Resource"`
	Environments     []Environment      `xml:"Environment"`
	ResourceLinks    []ResourceLink     `xml:"ResourceLink"`
	Parameters       []ContextParameter `xml:"Parameter"`
	WatchedResources []string           `xml:"WatchedResource"`
	Manager          *ContextManager    `xml:"Manager,omitempty"`
	Loader           *ContextLoader     `xml:"Loader,omitempty"`
	JarScanner       *JarScanner        `xml:"JarScanner,omitempty"`
	CookieProcessor  *CookieProcessor   `xml:"CookieProcessor,omitempty"`
	Valves           []ContextValve     `xml:"Valve"`
}

// ContextParameter represents a context parameter
type ContextParameter struct {
	XMLName     xml.Name `xml:"Parameter"`
	Name        string   `xml:"name,attr"`
	Value       string   `xml:"value,attr"`
	Override    bool     `xml:"override,attr,omitempty"`
	Description string   `xml:"description,attr,omitempty"`
}

// ContextManager represents a session manager for context
type ContextManager struct {
	XMLName                              xml.Name `xml:"Manager"`
	ClassName                            string   `xml:"className,attr,omitempty"`
	MaxActiveSessions                    int      `xml:"maxActiveSessions,attr,omitempty"`
	SessionIdLength                      int      `xml:"sessionIdLength,attr,omitempty"`
	MaxInactiveInterval                  int      `xml:"maxInactiveInterval,attr,omitempty"`
	SessionAttributeValueClassNameFilter string   `xml:"sessionAttributeValueClassNameFilter,attr,omitempty"`
	// PersistentManager specific
	SaveOnRestart string        `xml:"saveOnRestart,attr,omitempty"`
	MinIdleSwap   int           `xml:"minIdleSwap,attr,omitempty"`
	MaxIdleSwap   int           `xml:"maxIdleSwap,attr,omitempty"`
	MaxIdleBackup int           `xml:"maxIdleBackup,attr,omitempty"`
	Store         *SessionStore `xml:"Store,omitempty"`
}

// SessionStore represents a session store for persistent manager
type SessionStore struct {
	XMLName   xml.Name `xml:"Store"`
	ClassName string   `xml:"className,attr"`
	Directory string   `xml:"directory,attr,omitempty"`
	// JDBCStore specific
	DriverName         string `xml:"driverName,attr,omitempty"`
	ConnectionURL      string `xml:"connectionURL,attr,omitempty"`
	ConnectionName     string `xml:"connectionName,attr,omitempty"`
	ConnectionPassword string `xml:"connectionPassword,attr,omitempty"`
	SessionTable       string `xml:"sessionTable,attr,omitempty"`
}

// ContextLoader represents a class loader for context
type ContextLoader struct {
	XMLName             xml.Name `xml:"Loader"`
	ClassName           string   `xml:"className,attr,omitempty"`
	Delegate            string   `xml:"delegate,attr,omitempty"`
	Reloadable          string   `xml:"reloadable,attr,omitempty"`
	SearchExternalFirst string   `xml:"searchExternalFirst,attr,omitempty"`
}

// JarScanner represents a JAR scanner configuration
type JarScanner struct {
	XMLName                xml.Name       `xml:"JarScanner"`
	ClassName              string         `xml:"className,attr,omitempty"`
	ScanClassPath          string         `xml:"scanClassPath,attr,omitempty"`
	ScanManifest           string         `xml:"scanManifest,attr,omitempty"`
	ScanAllFiles           string         `xml:"scanAllFiles,attr,omitempty"`
	ScanAllDirectories     string         `xml:"scanAllDirectories,attr,omitempty"`
	ScanBootstrapClassPath string         `xml:"scanBootstrapClassPath,attr,omitempty"`
	JarScanFilter          *JarScanFilter `xml:"JarScanFilter,omitempty"`
}

// JarScanFilter represents a JAR scan filter
type JarScanFilter struct {
	XMLName                 xml.Name `xml:"JarScanFilter"`
	DefaultPluggabilityScan string   `xml:"defaultPluggabilityScan,attr,omitempty"`
	DefaultTldScan          string   `xml:"defaultTldScan,attr,omitempty"`
	PluggabilitySkip        string   `xml:"pluggabilitySkip,attr,omitempty"`
	PluggabilityScan        string   `xml:"pluggabilityScan,attr,omitempty"`
	TldSkip                 string   `xml:"tldSkip,attr,omitempty"`
	TldScan                 string   `xml:"tldScan,attr,omitempty"`
}

// CookieProcessor represents a cookie processor configuration
type CookieProcessor struct {
	XMLName                 xml.Name `xml:"CookieProcessor"`
	ClassName               string   `xml:"className,attr,omitempty"`
	SameSiteCookies         string   `xml:"sameSiteCookies,attr,omitempty"`
	AllowEqualsInValue      string   `xml:"allowEqualsInValue,attr,omitempty"`
	AllowHttpSepsInV0       string   `xml:"allowHttpSepsInV0,attr,omitempty"`
	AllowNameOnly           string   `xml:"allowNameOnly,attr,omitempty"`
	ForwardSlashIsSeparator string   `xml:"forwardSlashIsSeparator,attr,omitempty"`
}

// ContextValve represents a valve in context
type ContextValve struct {
	XMLName   xml.Name `xml:"Valve"`
	ClassName string   `xml:"className,attr"`
	// Common valve attributes
	Pattern   string `xml:"pattern,attr,omitempty"`
	Directory string `xml:"directory,attr,omitempty"`
	Prefix    string `xml:"prefix,attr,omitempty"`
	Suffix    string `xml:"suffix,attr,omitempty"`
	Rotatable string `xml:"rotatable,attr,omitempty"`
	Buffered  string `xml:"buffered,attr,omitempty"`
	// Additional attributes stored as raw
	Allow string `xml:"allow,attr,omitempty"`
	Deny  string `xml:"deny,attr,omitempty"`
}

// Manager class names
const (
	ManagerStandard   = "org.apache.catalina.session.StandardManager"
	ManagerPersistent = "org.apache.catalina.session.PersistentManager"
)

// Store class names
const (
	StoreFile = "org.apache.catalina.session.FileStore"
	StoreJDBC = "org.apache.catalina.session.JDBCStore"
)

// CookieProcessor class names
const (
	CookieProcessorRfc6265 = "org.apache.tomcat.util.http.Rfc6265CookieProcessor"
	CookieProcessorLegacy  = "org.apache.tomcat.util.http.LegacyCookieProcessor"
)

// SameSite cookie values
var SameSiteCookieValues = []string{"unset", "none", "lax", "strict"}

// NewContextParameter creates a new context parameter
func NewContextParameter(name, value string) *ContextParameter {
	return &ContextParameter{
		Name:     name,
		Value:    value,
		Override: true,
	}
}

// NewContextManager creates a default session manager
func NewContextManager() *ContextManager {
	return &ContextManager{
		ClassName:           ManagerStandard,
		MaxActiveSessions:   -1,
		SessionIdLength:     16,
		MaxInactiveInterval: 1800,
	}
}

// NewPersistentManager creates a persistent session manager
func NewPersistentManager() *ContextManager {
	return &ContextManager{
		ClassName:         ManagerPersistent,
		MaxActiveSessions: -1,
		SaveOnRestart:     "true",
		MinIdleSwap:       -1,
		MaxIdleSwap:       -1,
		MaxIdleBackup:     -1,
		Store: &SessionStore{
			ClassName: StoreFile,
			Directory: "sessions",
		},
	}
}

// NewJarScanner creates a default JAR scanner
func NewJarScanner() *JarScanner {
	return &JarScanner{
		ScanClassPath:          "true",
		ScanManifest:           "false",
		ScanAllFiles:           "false",
		ScanBootstrapClassPath: "false",
	}
}

// NewCookieProcessor creates a default cookie processor
func NewCookieProcessor() *CookieProcessor {
	return &CookieProcessor{
		ClassName:       CookieProcessorRfc6265,
		SameSiteCookies: "unset",
	}
}

// NewDataSourceResource creates a DataSource resource with default values
func NewDataSourceResource(name string) *Resource {
	return &Resource{
		Name:                          name,
		Auth:                          AuthContainer,
		Type:                          string(ResourceTypeDataSource),
		Factory:                       "org.apache.tomcat.jdbc.pool.DataSourceFactory",
		InitialSize:                   10,
		MaxTotal:                      100,
		MaxIdle:                       30,
		MinIdle:                       10,
		MaxWaitMillis:                 10000,
		TestOnBorrow:                  true,
		TestWhileIdle:                 true,
		TimeBetweenEvictionRunsMillis: 30000,
		MinEvictableIdleTimeMillis:    60000,
	}
}

// NewMailSessionResource creates a Mail Session resource with default values
func NewMailSessionResource(name string) *Resource {
	return &Resource{
		Name:                  name,
		Auth:                  AuthContainer,
		Type:                  string(ResourceTypeMailSession),
		MailSmtpHost:          "localhost",
		MailSmtpPort:          "25",
		MailTransportProtocol: "smtp",
	}
}

// NewUserDatabaseResource creates a UserDatabase resource with default values
func NewUserDatabaseResource(name string) *Resource {
	return &Resource{
		Name:        name,
		Auth:        AuthContainer,
		Type:        string(ResourceTypeUserDatabase),
		Description: "User database that can be updated and saved",
		Factory:     "org.apache.catalina.users.MemoryUserDatabaseFactory",
		Pathname:    "conf/tomcat-users.xml",
	}
}

// NewEnvironment creates an Environment with default values
func NewEnvironment(name, value, valueType string) *Environment {
	return &Environment{
		Name:     name,
		Value:    value,
		Type:     valueType,
		Override: true,
	}
}

// NewResourceLink creates a ResourceLink
func NewResourceLink(name, global, linkType string) *ResourceLink {
	return &ResourceLink{
		Name:   name,
		Global: global,
		Type:   linkType,
	}
}

// Common JDBC drivers
var CommonDrivers = map[string]string{
	"MySQL":      "com.mysql.cj.jdbc.Driver",
	"PostgreSQL": "org.postgresql.Driver",
	"Oracle":     "oracle.jdbc.OracleDriver",
	"SQL Server": "com.microsoft.sqlserver.jdbc.SQLServerDriver",
	"MariaDB":    "org.mariadb.jdbc.Driver",
	"H2":         "org.h2.Driver",
	"HSQLDB":     "org.hsqldb.jdbc.JDBCDriver",
	"Derby":      "org.apache.derby.jdbc.ClientDriver",
	"SQLite":     "org.sqlite.JDBC",
}

// JDBC URL templates
var JDBCURLTemplates = map[string]string{
	"MySQL":      "jdbc:mysql://localhost:3306/dbname?useSSL=false&serverTimezone=UTC",
	"PostgreSQL": "jdbc:postgresql://localhost:5432/dbname",
	"Oracle":     "jdbc:oracle:thin:@localhost:1521:ORCL",
	"SQL Server": "jdbc:sqlserver://localhost:1433;databaseName=dbname",
	"MariaDB":    "jdbc:mariadb://localhost:3306/dbname",
	"H2":         "jdbc:h2:mem:testdb;DB_CLOSE_DELAY=-1",
	"HSQLDB":     "jdbc:hsqldb:mem:testdb",
	"Derby":      "jdbc:derby://localhost:1527/dbname",
	"SQLite":     "jdbc:sqlite:path/to/database.db",
}

// Validation queries for different databases
var ValidationQueries = map[string]string{
	"MySQL":      "SELECT 1",
	"PostgreSQL": "SELECT 1",
	"Oracle":     "SELECT 1 FROM DUAL",
	"SQL Server": "SELECT 1",
	"MariaDB":    "SELECT 1",
	"H2":         "SELECT 1",
	"HSQLDB":     "SELECT 1 FROM INFORMATION_SCHEMA.SYSTEM_USERS",
	"Derby":      "VALUES 1",
	"SQLite":     "SELECT 1",
}

// Environment value types
var EnvironmentTypes = []string{
	"java.lang.String",
	"java.lang.Integer",
	"java.lang.Long",
	"java.lang.Boolean",
	"java.lang.Double",
	"java.lang.Float",
	"java.lang.Short",
	"java.lang.Byte",
	"java.lang.Character",
}

// Transaction isolation levels
var TransactionIsolationLevels = []string{
	"NONE",
	"READ_UNCOMMITTED",
	"READ_COMMITTED",
	"REPEATABLE_READ",
	"SERIALIZABLE",
}

// DataSource factory classes
var DataSourceFactories = []string{
	"org.apache.tomcat.jdbc.pool.DataSourceFactory",
	"org.apache.tomcat.dbcp.dbcp2.BasicDataSourceFactory",
	"org.apache.commons.dbcp2.BasicDataSourceFactory",
}
