package i18n

// translations contains all message translations
var translations = map[Language]map[string]string{
	English: {
		// Common
		"app.title":            "TomcatKit - Tomcat Configuration Manager",
		"app.status.ready":     "Ready",
		"app.status.saved":     "Configuration saved successfully",
		"app.status.error":     "Error: %s",
		"app.lang.select":      "Select Language",
		"app.lang.current":     "Current: %s",
		"common.back":          "Back",
		"common.cancel":        "Cancel",
		"common.save":          "Save Configuration",
		"common.save.short":    "Save",
		"common.apply":         "Apply",
		"common.delete":        "Delete",
		"common.add":           "Add",
		"common.edit":          "Edit",
		"common.remove":        "Remove",
		"common.yes":           "Yes",
		"common.no":            "No",
		"common.confirm":       "Confirm",
		"common.warning":       "Warning",
		"common.error":         "Error",
		"common.success":       "Success",
		"common.loading":       "Loading...",
		"common.return":        "Return to main menu",
		"common.enabled":       "Enabled",
		"common.disabled":      "Disabled",
		"common.notconfigured": "Not configured",
		"common.minutes":       "minutes",
		"common.selecthost":    "Select Host",
		"common.selecttype":    "Select Type",
		"common.settings":      "Settings",
		"common.parameters":    "Parameters",
		"common.contexts":      "Contexts",
		"common.addall":        "Add All",
		"help.title":           "Help",
		"preview.title":        "XML Preview",
		"preview.properties":   "Properties Preview",

		// Main Menu
		"menu.title":               "Main Menu",
		"menu.server":              "Server",
		"menu.server.desc":         "Configure server.xml core settings",
		"menu.connector":           "Connector",
		"menu.connector.desc":      "HTTP, AJP, SSL/TLS connectors",
		"menu.security":            "Security / Realm",
		"menu.security.desc":       "Authentication realms and users",
		"menu.jndi":                "JNDI Resources",
		"menu.jndi.desc":           "DataSource, Mail Session, Environment",
		"menu.host":                "Virtual Hosts",
		"menu.host.desc":           "Host, Context, Session Manager",
		"menu.valve":               "Valves",
		"menu.valve.desc":          "AccessLog, RemoteAddr, SSO valves",
		"menu.cluster":             "Clustering",
		"menu.cluster.desc":        "Session replication, membership",
		"menu.logging":             "Logging",
		"menu.logging.desc":        "JULI logging.properties",
		"menu.context":             "Context",
		"menu.context.desc":        "context.xml configuration",
		"menu.web":                 "Web Application",
		"menu.web.desc":            "web.xml servlets, filters, security",
		"menu.quicktemplates":      "Quick Templates",
		"menu.quicktemplates.desc": "Apply common configurations quickly",
		"menu.exit":                "Exit",
		"menu.exit.desc":           "Exit TomcatKit",

		// Footer
		"footer.navigate": "Navigate",
		"footer.select":   "Select",
		"footer.back":     "Back",
		"footer.lang":     "Lang",
		"footer.quit":     "Quit",

		// Quick Templates
		"qt.title":              "Quick Templates",
		"qt.select":             "Select a quick template to apply",
		"qt.virtualthread":      "Virtual Threads",
		"qt.virtualthread.desc": "Enable Virtual Thread executor (Java 21+, Tomcat 11+)",
		"qt.https":              "HTTPS with SSL",
		"qt.https.desc":         "Configure HTTPS connector with SSL/TLS",
		"qt.connpool":           "Connection Pool Tuning",
		"qt.connpool.desc":      "Optimize connection pool settings",
		"qt.gzip":               "Gzip Compression",
		"qt.gzip.desc":          "Enable response compression",
		"qt.accesslog":          "Access Log",
		"qt.accesslog.desc":     "Enable access logging",
		"qt.security":           "Security Hardening",
		"qt.security.desc":      "Apply security best practices",
		"qt.apache":             "Apache httpd (mod_jk/AJP)",
		"qt.apache.desc":        "Configure AJP connector for Apache httpd",
		"qt.nginx":              "nginx Reverse Proxy",
		"qt.nginx.desc":         "Configure for nginx proxy_pass",
		"qt.haproxy":            "HAProxy Load Balancer",
		"qt.haproxy.desc":       "Configure for HAProxy load balancing",

		// Virtual Thread Template
		"qt.vt.title":            "Virtual Thread Template",
		"qt.vt.info":             "Virtual Thread Executor (Java 21+, Tomcat 11+)\n\nVirtual threads are lightweight threads that can significantly\nimprove throughput for I/O-bound applications.",
		"qt.vt.requirements":     "Requirements:",
		"qt.vt.req.java":         "Java 21 or later",
		"qt.vt.req.tomcat":       "Tomcat 11.0 or later (or Tomcat 10.1.25+)",
		"qt.vt.willdo":           "This template will:",
		"qt.vt.willdo.create":    "Create a Virtual Thread Executor",
		"qt.vt.willdo.configure": "Configure the HTTP connector to use the executor",
		"qt.vt.warning.exists":   "Warning: Virtual Thread executor already exists!",
		"qt.vt.executor.name":    "Executor Name",
		"qt.vt.thread.prefix":    "Thread Name Prefix",
		"qt.vt.max.queue":        "Max Queue Size",
		"qt.vt.apply.connector":  "Apply to Connector",
		"qt.vt.apply":            "Apply Template",
		"qt.vt.success":          "Virtual Thread executor applied successfully!",

		// HTTPS Template
		"qt.https.title":         "HTTPS Template",
		"qt.https.info":          "HTTPS Connector Configuration\n\nThis template creates an HTTPS connector with SSL/TLS.",
		"qt.https.need":          "You will need:",
		"qt.https.need.keystore": "A keystore file (.jks or .p12)",
		"qt.https.need.password": "Keystore password",
		"qt.https.port":          "HTTPS Port",
		"qt.https.keystore.file": "Keystore File",
		"qt.https.keystore.pass": "Keystore Password",
		"qt.https.keystore.type": "Keystore Type",
		"qt.https.success":       "HTTPS connector added successfully!",

		// Connection Pool Template
		"qt.pool.title":           "Connection Pool Tuning",
		"qt.pool.info":            "Connection Pool Optimization\n\nTune thread pool settings for better performance.",
		"qt.pool.recommended":     "Recommended settings:",
		"qt.pool.dev":             "Development: 25-100 threads",
		"qt.pool.prod":            "Production: 150-400 threads",
		"qt.pool.high":            "High traffic: 400-800 threads",
		"qt.pool.profile":         "Profile",
		"qt.pool.maxthreads":      "Max Threads",
		"qt.pool.minsparethreads": "Min Spare Threads",
		"qt.pool.acceptcount":     "Accept Count",
		"qt.pool.conntimeout":     "Connection Timeout (ms)",
		"qt.pool.apply":           "Apply to All HTTP Connectors",
		"qt.pool.success":         "Connection pool settings applied!",

		// Gzip Template
		"qt.gzip.title":   "Gzip Compression",
		"qt.gzip.info":    "Gzip Compression\n\nEnable response compression for text-based content.",
		"qt.gzip.willdo":  "This will add compression attributes to HTTP connectors:",
		"qt.gzip.minsize": "Min Compression Size (bytes)",
		"qt.gzip.success": "Gzip compression enabled!",

		// Access Log Template
		"qt.al.title":            "Access Log Template",
		"qt.al.info":             "Access Log Valve\n\nConfigure access logging for HTTP requests.",
		"qt.al.patterns":         "Common patterns:",
		"qt.al.pattern.common":   "common: %h %l %u %t \"%r\" %s %b",
		"qt.al.pattern.combined": "combined: %h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-Agent}i\"",
		"qt.al.pattern":          "Log Pattern",
		"qt.al.directory":        "Directory",
		"qt.al.prefix":           "File Prefix",
		"qt.al.suffix":           "File Suffix",
		"qt.al.success":          "Access log configured successfully!",

		// Security Hardening Template
		"qt.sec.title":             "Security Hardening",
		"qt.sec.info":              "Security Hardening\n\nApply security best practices to your Tomcat configuration.",
		"qt.sec.willdo":            "This template will:",
		"qt.sec.willdo.shutdown":   "Change default shutdown port from 8005 to -1 (disabled)",
		"qt.sec.willdo.command":    "Change shutdown command",
		"qt.sec.willdo.version":    "Remove server version from error pages",
		"qt.sec.willdo.listener":   "Add security-related listeners",
		"qt.sec.disable.shutdown":  "Disable Shutdown Port",
		"qt.sec.remove.serverinfo": "Remove Server Info from Errors",
		"qt.sec.add.listener":      "Add Security Listener",
		"qt.sec.success":           "Security hardening applied!",

		// Apache httpd Template
		"qt.ajp.title":          "Apache httpd (AJP) Template",
		"qt.ajp.info":           "Apache httpd Integration (mod_jk / mod_proxy_ajp)\n\nConfigure AJP connector for Apache httpd reverse proxy.",
		"qt.ajp.modules":        "Supported Apache modules:",
		"qt.ajp.modjk":          "mod_jk: Traditional Tomcat connector",
		"qt.ajp.modproxy":       "mod_proxy_ajp: Apache proxy module for AJP",
		"qt.ajp.willdo":         "This template will:",
		"qt.ajp.willdo.create":  "Create AJP/1.3 connector on specified port",
		"qt.ajp.willdo.secret":  "Configure secret authentication (Tomcat 9.0.31+)",
		"qt.ajp.willdo.valve":   "Add RemoteIpValve for proper client IP handling",
		"qt.ajp.port":           "AJP Port",
		"qt.ajp.address":        "Bind Address",
		"qt.ajp.secret":         "AJP Secret",
		"qt.ajp.remoteip":       "Add RemoteIpValve",
		"qt.ajp.success":        "Apache httpd AJP connector configured successfully!",
		"qt.ajp.config.title":   "Apache httpd Configuration",
		"qt.ajp.config.applied": "Apache httpd Configuration Applied!",
		"qt.ajp.config.created": "AJP Connector created on port %s",
		"qt.ajp.config.copy":    "Copy the following configuration to your Apache httpd:",
		"qt.ajp.config.option1": "Option 1: mod_proxy_ajp (Recommended)",
		"qt.ajp.config.option2": "Option 2: mod_jk (workers.properties)",
		"qt.ajp.config.note":    "Note: Restart Apache httpd after applying configuration.",

		// nginx Template
		"qt.nginx.title":          "nginx Reverse Proxy Template",
		"qt.nginx.info":           "nginx Reverse Proxy Configuration\n\nConfigure Tomcat for nginx proxy_pass.",
		"qt.nginx.willdo":         "This template will:",
		"qt.nginx.willdo.valve":   "Add RemoteIpValve for X-Forwarded-* headers",
		"qt.nginx.willdo.ip":      "Configure proper client IP handling",
		"qt.nginx.willdo.http":    "Optionally adjust HTTP connector settings",
		"qt.nginx.proxy.note":     "nginx will proxy to Tomcat's HTTP connector",
		"qt.nginx.connector":      "HTTP Connector",
		"qt.nginx.internal":       "Internal Proxies (regex)",
		"qt.nginx.proto":          "Handle X-Forwarded-Proto",
		"qt.nginx.success":        "nginx reverse proxy configured successfully!",
		"qt.nginx.config.title":   "nginx Configuration",
		"qt.nginx.config.applied": "nginx Reverse Proxy Configuration Applied!",
		"qt.nginx.config.valve":   "RemoteIpValve configured for nginx proxy",
		"qt.nginx.config.copy":    "Copy the following configuration to your nginx:",
		"qt.nginx.config.basic":   "Basic Configuration",
		"qt.nginx.config.https":   "HTTPS Configuration (with SSL termination)",
		"qt.nginx.config.note":    "Note: Restart nginx after applying configuration.",

		// HAProxy Template
		"qt.haproxy.title":           "HAProxy Load Balancer Template",
		"qt.haproxy.info":            "HAProxy Load Balancer Configuration\n\nConfigure Tomcat for HAProxy load balancing.",
		"qt.haproxy.willdo":          "This template will:",
		"qt.haproxy.willdo.valve":    "Add RemoteIpValve for X-Forwarded-* headers",
		"qt.haproxy.willdo.jvm":      "Configure jvmRoute for sticky sessions (optional)",
		"qt.haproxy.willdo.ip":       "Set up proper client IP handling",
		"qt.haproxy.modes":           "Load balancing modes supported:",
		"qt.haproxy.mode.http":       "HTTP mode (Layer 7) - recommended",
		"qt.haproxy.mode.tcp":        "TCP mode (Layer 4) - for SSL passthrough",
		"qt.haproxy.connector":       "HTTP Connector",
		"qt.haproxy.sticky":          "Enable Sticky Sessions",
		"qt.haproxy.jvmroute":        "JVM Route (node ID)",
		"qt.haproxy.internal":        "Internal Proxies (regex)",
		"qt.haproxy.success":         "HAProxy load balancer configured successfully!",
		"qt.haproxy.config.title":    "HAProxy Configuration",
		"qt.haproxy.config.applied":  "HAProxy Load Balancer Configuration Applied!",
		"qt.haproxy.config.valve":    "RemoteIpValve configured for HAProxy",
		"qt.haproxy.config.sticky":   "Sticky Session Notes:",
		"qt.haproxy.config.jvmroute": "jvmRoute \"%s\" configured in Tomcat Engine",
		"qt.haproxy.config.cookie":   "HAProxy uses JSESSIONID cookie for session affinity",
		"qt.haproxy.config.format":   "Session ID format: <session-id>.<jvmRoute>",
		"qt.haproxy.config.copy":     "Copy the following configuration to your HAProxy:",
		"qt.haproxy.config.http":     "HTTP Mode (Layer 7) - Recommended",
		"qt.haproxy.config.tcp":      "TCP Mode (Layer 4) - SSL Passthrough",
		"qt.haproxy.config.health":   "Health Check Endpoint (Optional)",
		"qt.haproxy.config.stats":    "HAProxy Stats (Optional)",
		"qt.haproxy.config.note":     "Note: Restart HAProxy after applying configuration.",

		// Continue prompt
		"prompt.continue": "Press Enter or Escape to continue",
		"prompt.nohttp":   "No HTTP connector found!",

		// Instance Selection
		"instance.title":           "Select Tomcat Instance",
		"instance.recent":          "Recent Instances",
		"instance.detected":        "Detected Instances",
		"instance.none":            "No Tomcat installations detected",
		"instance.manual":          "Enter Path Manually",
		"instance.manual.desc":     "Specify CATALINA_HOME path",
		"instance.running":         "Running",
		"instance.noselected":      "No Tomcat instance selected",
		"instance.pressT":          "Press 't' to select a Tomcat instance",
		"instance.path.title":      "Enter Tomcat Path",
		"instance.path.home":       "CATALINA_HOME",
		"instance.path.base":       "CATALINA_BASE (optional)",
		"instance.path.validate":   "Validate & Select",
		"instance.path.required":   "CATALINA_HOME is required",
		"instance.path.invalid":    "Invalid path: server.xml not found",
		"instance.selected":        "Tomcat instance selected successfully",
		"instance.info":            "Tomcat Instance",
		"instance.version":         "Version",
		"instance.status":          "Status",
		"instance.stopped":         "Stopped",
		"instance.ready":           "Ready to configure",
		"instance.path.help.home":  "CATALINA_HOME: Tomcat installation directory (contains bin, lib, conf)",
		"instance.path.help.base":  "CATALINA_BASE: Instance directory (optional, defaults to CATALINA_HOME)",
		"instance.path.help.xml":   "The path should contain conf/server.xml",
		"instance.info.noselected": "No Tomcat instance selected",
		"instance.info.getstarted": "To get started:",
		"instance.info.step1":      "Press 't' to select a Tomcat instance",
		"instance.info.step2":      "Or run with: tomcatkit -home /path/to/tomcat",
		"instance.info.autodetect": "TomcatKit will auto-detect installed Tomcat instances.",

		// Server View
		"server.title":                      "Server Configuration",
		"server.port":                       "Shutdown Port",
		"server.port.desc":                  "Server shutdown listener port",
		"server.shutdown":                   "Shutdown Command",
		"server.listeners":                  "Listeners",
		"server.listeners.desc":             "Lifecycle listeners",
		"server.services":                   "Services",
		"server.services.desc":              "Service configurations",
		"server.globalresources":            "Global Resources",
		"server.globalresources.desc":       "Global JNDI resources",
		"server.listener.add":               "Add Listener",
		"server.listener.edit":              "Edit Listener",
		"server.listener.classname":         "Class Name",
		"server.service":                    "Service",
		"server.service.name":               "Service Name",
		"server.engine":                     "Engine",
		"server.engine.name":                "Engine Name",
		"server.engine.defaulthost":         "Default Host",
		"server.engine.jvmroute":            "JVM Route",
		"server.executor":                   "Executors",
		"server.executor.add":               "Add Executor",
		"server.executor.edit":              "Edit Executor",
		"server.executor.name":              "Executor Name",
		"server.executor.prefix":            "Name Prefix",
		"server.executor.maxthreads":        "Max Threads",
		"server.executor.minthreads":        "Min Spare Threads",
		"server.executor.maxidle":           "Max Idle Time (ms)",
		"server.executor.threads":           "Threads: %d-%d, MaxIdle: %dms",
		"server.executor.updated":           "Executor updated",
		"server.executor.deleted":           "Executor deleted",
		"server.executor.added":             "Executor added",
		"server.listener.delete":            "Delete",
		"server.listener.custom":            "Custom",
		"server.listener.custom.desc":       "Enter a custom listener class",
		"server.listener.custom.title":      "Custom Listener",
		"server.listener.sslengine":         "SSL Engine",
		"server.listener.sslseed":           "SSL Random Seed",
		"server.listener.updated":           "Listener updated successfully",
		"server.listener.deleted":           "Listener deleted",
		"server.listener.added":             "Listener added",
		"server.listener.detail":            "Listener Detail",
		"server.listener.classrequired":     "Class name is required",
		"server.service.edit":               "Edit Service Name",
		"server.service.engine.desc":        "Default Host: %s, jvmRoute: %s",
		"server.service.connectors":         "Connectors",
		"server.service.connectors.desc":    "HTTP/AJP connectors (configure in Connector menu)",
		"server.service.hosts":              "Hosts",
		"server.service.hosts.desc":         "Virtual hosts (configure in Host menu)",
		"server.service.updated":            "Service name updated",
		"server.engine.settings":            "Engine Settings",
		"server.engine.saved":               "Engine settings saved",
		"server.globalresource.add":         "Add Resource",
		"server.globalresource.add.desc":    "Add a new global resource",
		"server.globalresource.edit":        "Edit Resource",
		"server.globalresource.auth":        "Auth",
		"server.globalresource.type":        "Type",
		"server.globalresource.description": "Description",
		"server.globalresource.factory":     "Factory",
		"server.globalresource.pathname":    "Pathname",
		"server.globalresource.updated":     "Resource updated",
		"server.globalresource.deleted":     "Resource deleted",
		"server.globalresource.added":       "Resource added",
		"server.settings":                   "Server Settings",
		"server.settings.saved":             "Server settings saved successfully",
		"server.settings.invalidport":       "Invalid port number",
		"server.confirm.delete":             "Are you sure you want to delete this %s?",
		"server.confirm.yes":                "Yes",
		"server.confirm.no":                 "No",

		// Connector View
		"connector.title":                   "Connector Configuration",
		"connector.http":                    "HTTP Connectors",
		"connector.http.desc":               "HTTP/1.1, HTTP/2 connector settings",
		"connector.ajp":                     "AJP Connectors",
		"connector.ajp.desc":                "Apache JServ Protocol connector settings",
		"connector.ssl":                     "SSL/TLS Connectors",
		"connector.ssl.desc":                "HTTPS and SSL certificate settings",
		"connector.executor":                "Executors",
		"connector.executor.desc":           "Shared thread pool configuration",
		"connector.list":                    "Connectors",
		"connector.add":                     "Add Connector",
		"connector.edit":                    "Edit Connector",
		"connector.port":                    "Port",
		"connector.protocol":                "Protocol",
		"connector.timeout":                 "Connection Timeout",
		"connector.redirect":                "Redirect Port",
		"connector.maxthreads":              "Max Threads",
		"connector.minthreads":              "Min Spare Threads",
		"connector.acceptcount":             "Accept Count",
		"connector.ssl.enabled":             "SSL Enabled",
		"connector.ssl.keystore":            "Keystore File",
		"connector.ssl.password":            "Keystore Password",
		"connector.ssl.type":                "Keystore Type",
		"connector.ssl.protocol":            "SSL Protocol",
		"connector.ssl.clientauth":          "Client Auth",
		"connector.http.add":                "Add HTTP Connector",
		"connector.http.add.desc":           "Create a new HTTP connector",
		"connector.ajp.add":                 "Add AJP Connector",
		"connector.ajp.add.desc":            "Create a new AJP connector",
		"connector.ssl.add":                 "Add HTTPS Connector",
		"connector.ssl.add.desc":            "Create a new SSL/TLS connector",
		"connector.executor.add":            "Add Executor",
		"connector.executor.add.desc":       "Create a new thread pool",
		"connector.executor.title":          "Executors (Thread Pools)",
		"connector.service":                 "Service",
		"connector.secret":                  "Secret",
		"connector.secret.required":         "Secret Required",
		"connector.secret.none":             "No Secret",
		"connector.secret.set":              "Secret Set",
		"connector.keystore.notconfig":      "Keystore not configured",
		"connector.noservices":              "No services configured",
		"connector.updated.http":            "HTTP connector updated successfully",
		"connector.updated.ajp":             "AJP connector updated successfully",
		"connector.updated.ssl":             "SSL connector updated successfully",
		"connector.deleted":                 "Connector deleted",
		"connector.added":                   "Connector added successfully",
		"connector.ssl.added":               "SSL connector added successfully",
		"connector.executor.updated":        "Executor updated successfully",
		"connector.executor.deleted":        "Executor deleted",
		"connector.executor.added":          "Executor added successfully",
		"connector.delete.title":            "Delete Connector",
		"connector.delete.confirm":          "Delete connector on port %d?",
		"connector.delete.ajp.confirm":      "Delete AJP connector on port %d?",
		"connector.delete.ssl.confirm":      "Delete SSL connector on port %d?",
		"connector.executor.delete.title":   "Delete Executor",
		"connector.executor.delete.confirm": "Delete executor '%s'?",
		"connector.executor.name":           "Name",
		"connector.executor.nameprefix":     "Name Prefix",
		"connector.executor.maxidle":        "Max Idle Time (ms)",
		"connector.executor.optional":       "Executor (optional)",
		"connector.executor.add.title":      "Add Executor",
		"connector.returnmenu":              "Return to connector menu",
		"connector.secretrequired":          "Secret Required",
		"connector.sslprotocol":             "SSL Protocol",
		"connector.keystorefile":            "Keystore File",
		"connector.keystorepass":            "Keystore Password",
		"connector.keystoretype":            "Keystore Type",
		"connector.clientauth":              "Client Auth",
		"connector.http.add.title":          "Add HTTP Connector",
		"connector.ajp.add.title":           "Add AJP Connector",
		"connector.ssl.add.title":           "Add HTTPS Connector",
		"connector.error.noservices":        "No services configured",
		"connector.added.ssl":               "SSL connector added successfully",

		// Security View
		"security.title":                "Security & Authentication",
		"security.realm":                "Realm Configuration",
		"security.realm.desc":           "Configure authentication realm",
		"security.realm.add":            "Add Realm",
		"security.realm.edit":           "Edit Realm",
		"security.realm.type":           "Realm Type",
		"security.realm.current":        "Current",
		"security.realm.nested":         "Nested Realms",
		"security.realm.set":            "Set Realm Type",
		"security.realm.set.desc":       "Configure a different realm type",
		"security.realm.remove":         "Remove Realm",
		"security.realm.remove.desc":    "Remove current realm configuration",
		"security.realm.remove.confirm": "Remove the current realm configuration?",
		"security.realm.removed":        "Realm removed",
		"security.realm.config":         "Realm Configuration",
		"security.realm.selecttype":     "Select Realm Type",
		"security.users":                "Users & Roles",
		"security.users.desc":           "Manage tomcat-users.xml",
		"security.users.title":          "Users & Roles (tomcat-users.xml)",
		"security.users.list":           "Users",
		"security.users.list.desc":      "Manage user accounts",
		"security.credential":           "Credential Handler",
		"security.credential.desc":      "Password hashing configuration",
		"security.user.add":             "Add User",
		"security.user.edit":            "Edit User",
		"security.user.name":            "Username",
		"security.user.password":        "Password",
		"security.user.roles":           "Roles",
		"security.roles":                "Roles",
		"security.roles.list":           "Roles",
		"security.roles.list.desc":      "Manage role definitions",
		"security.role.add":             "Add Role",
		"security.role.name":            "Role Name",

		// JNDI View
		"jndi.title":                "JNDI Resources - context.xml",
		"jndi.resources":            "Resources",
		"jndi.resource.add":         "Add Resource",
		"jndi.resource.edit":        "Edit Resource",
		"jndi.resource.name":        "Resource Name",
		"jndi.resource.type":        "Resource Type",
		"jndi.resource.auth":        "Auth",
		"jndi.datasource":           "DataSource (JDBC)",
		"jndi.datasource.desc":      "Database connection pools",
		"jndi.datasource.driver":    "Driver Class",
		"jndi.datasource.url":       "JDBC URL",
		"jndi.datasource.username":  "Username",
		"jndi.datasource.password":  "Password",
		"jndi.datasource.maxactive": "Max Active",
		"jndi.datasource.maxidle":   "Max Idle",
		"jndi.mail":                 "Mail Session",
		"jndi.mail.desc":            "JavaMail configuration",
		"jndi.environment":          "Environment Entries",
		"jndi.environment.desc":     "Environment variables",
		"jndi.environment.add":      "Add Environment Entry",
		"jndi.environment.name":     "Entry Name",
		"jndi.environment.value":    "Value",
		"jndi.environment.type":     "Type",
		"jndi.resourcelink":         "Resource Links",
		"jndi.resourcelink.desc":    "Links to global resources",

		// Host View
		"host.title":                  "Virtual Hosts & Contexts",
		"host.list":                   "Hosts",
		"host.add":                    "Add Host",
		"host.edit":                   "Edit Host",
		"host.name":                   "Host Name",
		"host.appbase":                "App Base",
		"host.unpackwars":             "Unpack WARs",
		"host.autodeploy":             "Auto Deploy",
		"host.aliases":                "Aliases",
		"host.alias.add":              "Add Alias",
		"host.virtualhost":            "Virtual Hosts",
		"host.virtualhost.desc":       "Manage virtual hosts",
		"host.context":                "Contexts",
		"host.context.desc":           "Manage web application contexts",
		"host.engine":                 "Engine Settings",
		"host.engine.desc":            "Configure the Catalina engine",
		"context.title":               "Context Configuration (context.xml)",
		"context.list":                "Contexts",
		"context.add":                 "Add Context",
		"context.edit":                "Edit Context",
		"context.path":                "Context Path",
		"context.docbase":             "Doc Base",
		"context.reloadable":          "Reloadable",
		"context.settings":            "Context Settings",
		"context.settings.desc":       "Basic context attributes and options",
		"context.resources":           "JNDI Resources",
		"context.resources.count":     "%d resources (DataSource, MailSession, etc.)",
		"context.environment":         "Environment Entries",
		"context.environment.count":   "%d environment entries",
		"context.resourcelinks":       "Resource Links",
		"context.resourcelinks.count": "%d resource links",
		"context.parameters":          "Parameters",
		"context.parameters.count":    "%d context parameters",
		"context.watched":             "Watched Resources",
		"context.watched.count":       "%d watched resources",
		"context.manager":             "Session Manager",
		"context.cookie":              "Cookie Processor",
		"context.cookie.desc":         "SameSite, cookie settings",
		"context.jarscanner":          "JAR Scanner",
		"context.jarscanner.desc":     "Class scanning configuration",
		"context.save.desc":           "Write changes to context.xml",
		"host.sessionmanager":         "Session Manager",

		// Valve View
		"valve.title":             "Valves Configuration",
		"valve.list":              "Valves",
		"valve.add":               "Add Valve",
		"valve.edit":              "Edit Valve",
		"valve.type":              "Valve Type",
		"valve.engine":            "Engine Valves",
		"valve.engine.desc":       "Valves applied to all requests",
		"valve.host":              "Host Valves",
		"valve.host.desc":         "Valves for specific virtual hosts",
		"valve.context":           "Context Valves",
		"valve.context.desc":      "Valves for specific applications",
		"valve.quickadd":          "Quick Add Common Valves",
		"valve.quickadd.desc":     "Add commonly used valves",
		"valve.accesslog":         "Access Log Valve",
		"valve.accesslog.dir":     "Directory",
		"valve.accesslog.prefix":  "Prefix",
		"valve.accesslog.suffix":  "Suffix",
		"valve.accesslog.pattern": "Pattern",
		"valve.remoteaddr":        "Remote Address Valve",
		"valve.remoteaddr.allow":  "Allow",
		"valve.remoteaddr.deny":   "Deny",
		"valve.remoteip":          "Remote IP Valve",
		"valve.remoteip.header":   "Remote IP Header",
		"valve.remoteip.protocol": "Protocol Header",
		"valve.sso":               "Single Sign-On",
		"valve.error":             "Error Report Valve",

		// Cluster View
		"cluster.title":              "Clustering Configuration",
		"cluster.enable":             "Enable Clustering",
		"cluster.disable":            "Disable Clustering",
		"cluster.status":             "Cluster Status",
		"cluster.status.desc":        "Enable or disable clustering",
		"cluster.settings":           "Cluster Settings",
		"cluster.settings.desc":      "Basic cluster configuration",
		"cluster.manager":            "Session Manager",
		"cluster.manager.desc":       "DeltaManager or BackupManager",
		"cluster.manager.type":       "Manager Type",
		"cluster.manager.delta":      "DeltaManager",
		"cluster.manager.backup":     "BackupManager",
		"cluster.channel":            "Channel",
		"cluster.membership":         "Membership",
		"cluster.membership.desc":    "Multicast membership settings",
		"cluster.membership.address": "Multicast Address",
		"cluster.membership.port":    "Multicast Port",
		"cluster.receiver":           "Receiver",
		"cluster.receiver.desc":      "Message receiver configuration",
		"cluster.receiver.address":   "Address",
		"cluster.receiver.port":      "Port",
		"cluster.sender":             "Sender",
		"cluster.sender.desc":        "Message sender configuration",
		"cluster.interceptors":       "Interceptors",
		"cluster.interceptors.desc":  "Channel interceptors",
		"cluster.interceptor.add":    "Add Interceptor",
		"cluster.deployer":           "Farm Deployer",
		"cluster.deployer.desc":      "Cluster deployment settings",
		"cluster.deployer.remove":    "Remove Deployer",

		// Logging View
		"logging.title":              "Logging Configuration (logging.properties)",
		"logging.handlers":           "Handlers",
		"logging.handler.add":        "Add Handler",
		"logging.handler.edit":       "Edit Handler",
		"logging.handler.type":       "Handler Type",
		"logging.handler.level":      "Level",
		"logging.handler.directory":  "Directory",
		"logging.handler.prefix":     "Prefix",
		"logging.filehandlers":       "File Handlers",
		"logging.filehandlers.count": "%d file handlers configured",
		"logging.console":            "Console Handler",
		"logging.loggers":            "Loggers",
		"logging.loggers.count":      "%d loggers configured",
		"logging.logger.add":         "Add Logger",
		"logging.logger.edit":        "Edit Logger",
		"logging.logger.name":        "Logger Name",
		"logging.logger.level":       "Level",
		"logging.logger.handlers":    "Handlers",
		"logging.rootlogger":         "Root Logger",
		"logging.rootlogger.count":   "%d handlers assigned",
		"logging.save.desc":          "Write changes to logging.properties",

		// Context View (context.xml)
		"ctxxml.title":            "Context Configuration",
		"ctxxml.settings":         "Basic Settings",
		"ctxxml.reloadable":       "Reloadable",
		"ctxxml.crosscontext":     "Cross Context",
		"ctxxml.privileged":       "Privileged",
		"ctxxml.cookies":          "Cookie Settings",
		"ctxxml.cookies.httponly": "HTTP Only",
		"ctxxml.cookies.name":     "Session Cookie Name",
		"ctxxml.resources":        "Resources",
		"ctxxml.parameters":       "Parameters",
		"ctxxml.watched":          "Watched Resources",
		"ctxxml.manager":          "Session Manager",
		"ctxxml.loader":           "Class Loader",
		"ctxxml.jarscanner":       "JAR Scanner",

		// Web View (web.xml)
		"web.title":              "Web Application Configuration (web.xml)",
		"web.servlets":           "Servlets",
		"web.servlets.count":     "%d servlets configured",
		"web.filters":            "Filters",
		"web.filters.count":      "%d filters configured",
		"web.listeners":          "Listeners",
		"web.listeners.count":    "%d listeners configured",
		"web.session":            "Session Config",
		"web.welcomefiles":       "Welcome Files",
		"web.welcomefiles.desc":  "Default page files",
		"web.errorpages":         "Error Pages",
		"web.errorpages.count":   "%d error pages",
		"web.mime":               "MIME Mappings",
		"web.mime.count":         "%d MIME types",
		"web.security":           "Security Constraints",
		"web.security.count":     "%d constraints",
		"web.login":              "Login Config",
		"web.login.desc":         "Authentication method",
		"web.roles":              "Security Roles",
		"web.roles.desc":         "Define security roles",
		"web.contextparams":      "Context Parameters",
		"web.contextparams.desc": "Web app init params",
		"web.save.desc":          "Write changes to web.xml",
		// Web sub-menu keys
		"web.servlet.add":                     "Add Servlet",
		"web.servlet.add.desc":                "Create a new servlet",
		"web.servlet.edit":                    "Edit Servlet",
		"web.servlet.name":                    "Servlet Name",
		"web.servlet.class":                   "Servlet Class",
		"web.servlet.jsp":                     "JSP File (optional)",
		"web.servlet.loadonstartup":           "Load On Startup",
		"web.servlet.async":                   "Async Supported",
		"web.servlet.initparams":              "Init Params (name=value per line)",
		"web.servlet.urlpatterns":             "URL Patterns (one per line)",
		"web.servlet.quickdefault":            "Quick Add Default Servlet",
		"web.servlet.quickdefault.desc":       "Add Tomcat default servlet",
		"web.servlet.quickjsp":                "Quick Add JSP Servlet",
		"web.servlet.quickjsp.desc":           "Add Tomcat JSP servlet",
		"web.servlet.added":                   "Servlet added",
		"web.servlet.updated":                 "Servlet updated",
		"web.servlet.deleted":                 "Servlet deleted",
		"web.servlet.error.name":              "Servlet name is required",
		"web.servlet.error.class":             "Either servlet class or JSP file is required",
		"web.filter.add":                      "Add Filter",
		"web.filter.add.desc":                 "Create a new filter",
		"web.filter.edit":                     "Edit Filter",
		"web.filter.name":                     "Filter Name",
		"web.filter.class":                    "Filter Class",
		"web.filter.async":                    "Async Supported",
		"web.filter.initparams":               "Init Params (name=value per line)",
		"web.filter.urlpatterns":              "URL Patterns (one per line)",
		"web.filter.quickcors":                "Quick Add CORS Filter",
		"web.filter.quickcors.desc":           "Add CORS filter",
		"web.filter.quickencoding":            "Quick Add Encoding Filter",
		"web.filter.quickencoding.desc":       "Add character encoding filter",
		"web.filter.added":                    "Filter added",
		"web.filter.updated":                  "Filter updated",
		"web.filter.deleted":                  "Filter deleted",
		"web.filter.error.required":           "Filter name and class are required",
		"web.listener.add":                    "Add Listener",
		"web.listener.add.desc":               "Create a new listener",
		"web.listener.edit":                   "Edit Listener",
		"web.listener.class":                  "Listener Class",
		"web.listener.description":            "Description",
		"web.listener.added":                  "Listener added",
		"web.listener.deleted":                "Listener deleted",
		"web.listener.error.class":            "Listener class is required",
		"web.session.title":                   "Session Configuration",
		"web.session.timeout":                 "Session Timeout (minutes)",
		"web.session.tracking.cookie":         "Tracking: COOKIE",
		"web.session.tracking.url":            "Tracking: URL",
		"web.session.tracking.ssl":            "Tracking: SSL",
		"web.session.cookie.name":             "Cookie Name",
		"web.session.cookie.domain":           "Cookie Domain",
		"web.session.cookie.path":             "Cookie Path",
		"web.session.cookie.httponly":         "HttpOnly Cookie",
		"web.session.cookie.secure":           "Secure Cookie",
		"web.session.saved":                   "Session configuration saved",
		"web.welcomefiles.title":              "Welcome Files",
		"web.welcomefiles.perline":            "Welcome Files (one per line)",
		"web.welcomefiles.adddefaults":        "Add Defaults",
		"web.welcomefiles.saved":              "Welcome files saved",
		"web.welcomefiles.defaultadded":       "Default welcome files added",
		"web.errorpage.add":                   "Add Error Page",
		"web.errorpage.add.desc":              "Create a new error page",
		"web.errorpage.edit":                  "Edit Error Page",
		"web.errorpage.type":                  "Error Type",
		"web.errorpage.code":                  "Error Code (e.g., 404)",
		"web.errorpage.exception":             "Exception Type",
		"web.errorpage.location":              "Location (page path)",
		"web.errorpage.quickcommon":           "Add Common Error Pages",
		"web.errorpage.quickcommon.desc":      "Add 404, 500 error pages",
		"web.errorpage.added":                 "Error page added",
		"web.errorpage.deleted":               "Error page deleted",
		"web.errorpage.commonadded":           "Common error pages added",
		"web.errorpage.error.location":        "Location is required",
		"web.mime.add":                        "Add MIME Mapping",
		"web.mime.add.desc":                   "Create a new MIME mapping",
		"web.mime.edit":                       "Edit MIME Mapping",
		"web.mime.extension":                  "Extension (without dot)",
		"web.mime.type":                       "MIME Type",
		"web.mime.quickcommon":                "Add Common MIME Types",
		"web.mime.quickcommon.desc":           "Add common web MIME types",
		"web.mime.added":                      "MIME mapping added",
		"web.mime.deleted":                    "MIME mapping deleted",
		"web.mime.commonadded":                "Common MIME types added",
		"web.mime.error.required":             "Extension and MIME type are required",
		"web.securityconstraint.add":          "Add Security Constraint",
		"web.securityconstraint.add.desc":     "Create a new security constraint",
		"web.securityconstraint.edit":         "Edit Security Constraint",
		"web.securityconstraint.resourcename": "Resource Name",
		"web.securityconstraint.urlpatterns":  "URL Patterns (one per line)",
		"web.securityconstraint.httpmethods":  "HTTP Methods (comma-separated, empty=all)",
		"web.securityconstraint.roles":        "Required Roles (one per line)",
		"web.securityconstraint.transport":    "Transport Guarantee",
		"web.securityconstraint.added":        "Security constraint added",
		"web.securityconstraint.updated":      "Security constraint updated",
		"web.securityconstraint.deleted":      "Security constraint deleted",
		"web.login.title":                     "Login Configuration",
		"web.login.authmethod":                "Auth Method",
		"web.login.realmname":                 "Realm Name",
		"web.login.formloginpage":             "Form Login Page",
		"web.login.formerrorpage":             "Form Error Page",
		"web.login.saved":                     "Login configuration saved",
		"web.login.removed":                   "Login configuration removed",
		"web.role.add":                        "Add Security Role",
		"web.role.add.desc":                   "Create a new security role",
		"web.role.edit":                       "Edit Security Role",
		"web.role.name":                       "Role Name",
		"web.role.description":                "Description",
		"web.role.quickcommon":                "Add Common Roles",
		"web.role.quickcommon.desc":           "Add admin, user, manager roles",
		"web.role.added":                      "Security role added",
		"web.role.deleted":                    "Security role deleted",
		"web.role.commonadded":                "Common roles added",
		"web.role.error.name":                 "Role name is required",
		"web.contextparam.add":                "Add Context Parameter",
		"web.contextparam.add.desc":           "Create a new parameter",
		"web.contextparam.edit":               "Edit Context Parameter",
		"web.contextparam.name":               "Parameter Name",
		"web.contextparam.value":              "Parameter Value",
		"web.contextparam.description":        "Description",
		"web.contextparam.added":              "Context parameter added",
		"web.contextparam.updated":            "Context parameter updated",
		"web.contextparam.deleted":            "Context parameter deleted",
		"web.contextparam.error.name":         "Parameter name is required",
		"webxml.title":                        "Web Application Configuration",
		"webxml.servlets":                     "Servlets",
		"webxml.servlet.add":                  "Add Servlet",
		"webxml.servlet.edit":                 "Edit Servlet",
		"webxml.servlet.name":                 "Servlet Name",
		"webxml.servlet.class":                "Servlet Class",
		"webxml.servlet.mapping":              "URL Pattern",
		"webxml.filters":                      "Filters",
		"webxml.filter.add":                   "Add Filter",
		"webxml.filter.edit":                  "Edit Filter",
		"webxml.filter.name":                  "Filter Name",
		"webxml.filter.class":                 "Filter Class",
		"webxml.listeners":                    "Listeners",
		"webxml.listener.add":                 "Add Listener",
		"webxml.listener.class":               "Listener Class",
		"webxml.session":                      "Session Config",
		"webxml.session.timeout":              "Session Timeout (min)",
		"webxml.welcome":                      "Welcome Files",
		"webxml.error":                        "Error Pages",
		"webxml.mime":                         "MIME Mappings",
		"webxml.security":                     "Security Constraints",
		"webxml.security.add":                 "Add Security Constraint",
		"webxml.login":                        "Login Config",
		"webxml.login.method":                 "Auth Method",
		"webxml.roles":                        "Security Roles",

		// ==================== DETAILED HELP DESCRIPTIONS ====================
		// Server Configuration Help
		"help.server.port": `[::b]Shutdown Port[::-]
The TCP/IP port number on which Tomcat listens for shutdown commands.

[yellow]Default:[-] 8005
[yellow]Range:[-] 1-65535 (or -1 to disable)

[green]Security Note:[-]
• Set to -1 in production to disable remote shutdown
• Only listens on localhost (127.0.0.1)
• Must match the port in shutdown scripts

[gray]Example: <Server port="8005" shutdown="SHUTDOWN">[-]`,

		"help.server.shutdown": `[::b]Shutdown Command[::-]
The command string that must be received to trigger shutdown.

[yellow]Default:[-] SHUTDOWN

[green]Security Recommendations:[-]
• Change from default "SHUTDOWN" to a complex, random string
• Combined with port=-1, provides maximum security
• Used by catalina.sh stop / shutdown.bat

[gray]Example: <Server port="8005" shutdown="COMPLEX_SECRET_STRING">[-]`,

		"help.server.listener": `[::b]Lifecycle Listeners[::-]
Listeners respond to specific events in the server lifecycle.

[green]Common Listeners:[-]
• [yellow]VersionLoggerListener[-]: Logs Tomcat version info at startup
• [yellow]AprLifecycleListener[-]: Enables Apache Portable Runtime (APR)
• [yellow]JreMemoryLeakPreventionListener[-]: Prevents JRE memory leaks
• [yellow]GlobalResourcesLifecycleListener[-]: Required for JNDI resources
• [yellow]ThreadLocalLeakPreventionListener[-]: Cleans up thread-local vars

[gray]All listeners are optional but recommended for production.[-]`,

		"help.server.service": `[::b]Service[::-]
A Service groups one or more Connectors with a single Engine.

[yellow]Default name:[-] Catalina

[green]Components:[-]
• [yellow]Engine[-]: Request processing engine (one per Service)
• [yellow]Connectors[-]: HTTP, HTTPS, AJP (one or more)
• [yellow]Executors[-]: Shared thread pools (optional)

[gray]Most installations use a single Service named "Catalina".[-]`,

		"help.server.engine": `[::b]Engine[::-]
The Engine receives and processes all requests from Connectors.

[green]Attributes:[-]
• [yellow]name[-]: Logical name (default: Catalina)
• [yellow]defaultHost[-]: Default Host for unmatched requests
• [yellow]jvmRoute[-]: Unique ID for load balancer sticky sessions

[green]jvmRoute Usage:[-]
Session IDs become: <session-id>.<jvmRoute>
Example: ABC123.node1

[gray]Required for session affinity in clustered environments.[-]`,

		"help.server.executor": `[::b]Executor (Thread Pool)[::-]
A shared thread pool for Connectors.

[green]Attributes:[-]
• [yellow]name[-]: Unique identifier (referenced by Connectors)
• [yellow]maxThreads[-]: Maximum worker threads (default: 200)
• [yellow]minSpareThreads[-]: Minimum idle threads (default: 25)
• [yellow]maxIdleTime[-]: Idle thread timeout in ms (default: 60000)

[green]Sizing Guidelines:[-]
• Development: 25-100 threads
• Production: 150-400 threads
• High-traffic: 400-800 threads

[gray]Connectors reference Executors via executor="name" attribute.[-]`,

		// Connector Help
		"help.connector.http": `[::b]HTTP Connector[::-]
Handles HTTP/1.1 and HTTP/2 client connections.

[green]Key Attributes:[-]
• [yellow]port[-]: TCP port (default: 8080)
• [yellow]protocol[-]: HTTP/1.1, org.apache.coyote.http11.Http11NioProtocol
• [yellow]connectionTimeout[-]: Socket timeout in ms (default: 20000)
• [yellow]maxThreads[-]: Max request threads (default: 200)
• [yellow]acceptCount[-]: Backlog queue size (default: 100)

[green]Protocol Options:[-]
• HTTP/1.1 - Auto-selects NIO implementation
• Http11NioProtocol - Non-blocking I/O (recommended)
• Http11Nio2Protocol - NIO2/AIO implementation
• Http11AprProtocol - APR/native (requires APR library)`,

		"help.connector.https": `[::b]HTTPS Connector (SSL/TLS)[::-]
Secure HTTP connections with SSL/TLS encryption.

[green]Key Attributes:[-]
• [yellow]port[-]: TCP port (default: 8443)
• [yellow]SSLEnabled[-]: Must be "true"
• [yellow]keystoreFile[-]: Path to keystore (.jks, .p12)
• [yellow]keystorePass[-]: Keystore password
• [yellow]keystoreType[-]: JKS, PKCS12 (recommended)

[green]SSL Protocol Options:[-]
• TLS (auto-negotiates highest version)
• TLSv1.2 (minimum recommended)
• TLSv1.3 (most secure, Java 11+)

[green]Security Recommendations:[-]
• Use TLSv1.2 or TLSv1.3 only
• Disable weak ciphers
• Use 2048-bit+ RSA or 256-bit+ ECC keys`,

		"help.connector.ajp": `[::b]AJP Connector[::-]
Apache JServ Protocol for Apache httpd integration.

[green]Key Attributes:[-]
• [yellow]port[-]: TCP port (default: 8009)
• [yellow]protocol[-]: AJP/1.3
• [yellow]secretRequired[-]: Require shared secret (default: true in 9.0.31+)
• [yellow]secret[-]: Shared authentication secret
• [yellow]address[-]: Bind address (default: all interfaces)

[green]Security (Tomcat 9.0.31+):[-]
• secretRequired="true" enforces authentication
• secret must match Apache httpd's ProxyPass secret
• address="127.0.0.1" limits to localhost

[gray]Example: ProxyPass /app ajp://localhost:8009/app secret=mySecret[-]`,

		"help.connector.port": `[::b]Port[::-]
TCP port number for incoming connections.

[yellow]Default Ports:[-]
• HTTP: 8080
• HTTPS: 8443
• AJP: 8009

[green]Notes:[-]
• Ports below 1024 require root/admin privileges
• Use iptables/firewall redirect for 80→8080
• Each Connector must use a unique port`,

		"help.connector.protocol": `[::b]Protocol[::-]
The protocol handler implementation.

[green]HTTP Protocols:[-]
• [yellow]HTTP/1.1[-]: Auto-selects best implementation
• [yellow]Http11NioProtocol[-]: Non-blocking I/O (default)
• [yellow]Http11Nio2Protocol[-]: Java NIO2 (async)
• [yellow]Http11AprProtocol[-]: APR/native (requires APR)

[green]AJP Protocols:[-]
• [yellow]AJP/1.3[-]: Standard AJP protocol
• [yellow]AjpNioProtocol[-]: NIO implementation
• [yellow]AjpNio2Protocol[-]: NIO2 implementation

[gray]NIO is recommended for most deployments.[-]`,

		"help.connector.maxthreads": `[::b]Max Threads[::-]
Maximum number of request processing threads.

[yellow]Default:[-] 200

[green]Sizing Guidelines:[-]
• Low traffic: 25-100
• Medium traffic: 150-300
• High traffic: 400-800

[green]Formula:[-]
maxThreads ≈ (peak_concurrent_users × avg_response_time_sec)

[yellow]Warning:[-]
• Too few: requests queue and timeout
• Too many: memory exhaustion, context switching overhead
• Monitor thread pool usage in production`,

		"help.connector.minsparethreads": `[::b]Min Spare Threads[::-]
Minimum number of idle threads maintained.

[yellow]Default:[-] 10

[green]Purpose:[-]
• Keeps threads ready for incoming requests
• Reduces latency for initial requests
• Prevents thread creation overhead during bursts

[green]Recommendation:[-]
• Set to 10-25 for most applications
• Higher values for bursty traffic patterns
• Lower values to reduce memory usage`,

		"help.connector.connectiontimeout": `[::b]Connection Timeout[::-]
Milliseconds to wait for request data after connection.

[yellow]Default:[-] 20000 (20 seconds)

[green]Behavior:[-]
• Timer starts after TCP connection established
• Resets with each data received
• Closes connection if exceeded

[green]Recommendations:[-]
• 20000-60000 for normal applications
• Lower (5000-10000) behind load balancer
• Higher for slow clients or large uploads`,

		"help.connector.acceptcount": `[::b]Accept Count[::-]
Maximum queue length for incoming connections.

[yellow]Default:[-] 100

[green]Behavior:[-]
• Connections queue when all threads busy
• Exceeding this returns "connection refused"
• OS may have its own lower limit

[green]Tuning:[-]
• Increase for traffic spikes (200-500)
• Lower to fail-fast when overloaded
• Monitor queue depth in production`,

		"help.connector.redirectport": `[::b]Redirect Port[::-]
Port for automatic HTTPS redirects.

[yellow]Default:[-] 8443

[green]Usage:[-]
• Used when security-constraint requires HTTPS
• Redirects HTTP→HTTPS automatically
• Must match your HTTPS Connector port

[gray]Example: HTTP on 8080 redirects to HTTPS on 8443[-]`,

		"help.connector.executor": `[::b]Executor (Thread Pool)[::-]
Reference to a shared thread pool.

[yellow]Usage:[-]
• Leave empty to use connector's own thread pool
• Set executor name to share threads across connectors

[green]Benefits:[-]
• Centralized thread pool management
• Better resource utilization
• Easier monitoring and tuning

[gray]Define Executor in server.xml:
<Executor name="tomcatThreadPool"
  maxThreads="200" minSpareThreads="10"/>[-]`,

		"help.connector.secretrequired": `[::b]Secret Required[::-]
Enable secret-based authentication for AJP connections.

[yellow]Default:[-] true (Tomcat 9.0.31+)

[green]Purpose:[-]
• Prevents unauthorized access to AJP port
• Mitigates Ghostcat vulnerability (CVE-2020-1938)
• Requires matching secret on both sides

[red]Security:[-]
• Always enable in production
• Use strong, random secret values
• Never expose AJP port to untrusted networks`,

		"help.connector.secret": `[::b]Secret[::-]
Shared secret for AJP connector authentication.

[yellow]Requirements:[-]
• Must match Apache mod_proxy_ajp secret
• Use strong, random value (32+ chars)
• Keep confidential

[green]Configuration:[-]
• Tomcat: secret="yourSecretValue"
• Apache: ProxyPass ajp://host:8009 secret=yourSecretValue

[gray]Generate with: openssl rand -base64 32[-]`,

		// Security Help
		"help.security.realm": `[::b]Realm[::-]
A Realm connects Tomcat to a user/role database for authentication.

[green]Realm Types:[-]
• [yellow]UserDatabaseRealm[-]: Uses tomcat-users.xml (default)
• [yellow]JDBCRealm[-]: Database via JDBC
• [yellow]DataSourceRealm[-]: Database via JNDI DataSource
• [yellow]JNDIRealm[-]: LDAP/Active Directory
• [yellow]JAASRealm[-]: Java Authentication (JAAS)

[green]Placement:[-]
• Engine level: Applies to all Hosts/Contexts
• Host level: Applies to all Contexts in Host
• Context level: Applies to single application

[gray]Realms can be nested with CombinedRealm or LockOutRealm.[-]`,

		"help.security.userdatabase": `[::b]UserDatabaseRealm[::-]
File-based authentication using tomcat-users.xml.

[green]Configuration:[-]
• File: conf/tomcat-users.xml
• Roles: Define access permissions
• Users: Username, password, and role assignments

[green]Default Roles:[-]
• manager-gui: Access Manager web interface
• manager-script: Access Manager text/script interface
• admin-gui: Access Host Manager interface

[yellow]Security Note:[-]
Passwords in tomcat-users.xml should be digested:
1. Run: digest.sh -a SHA-256 mypassword
2. Use output in password attribute`,

		"help.security.jdbcrealm": `[::b]JDBCRealm / DataSourceRealm[::-]
Database-backed authentication.

[green]Required Tables:[-]
• Users table: username, password columns
• Roles table: username, role_name columns

[green]DataSourceRealm Attributes:[-]
• dataSourceName: JNDI name (e.g., jdbc/UserDB)
• userTable: Table containing usernames
• userNameCol: Username column name
• userCredCol: Password column name
• userRoleTable: Table containing roles
• roleNameCol: Role column name

[gray]DataSourceRealm is preferred over JDBCRealm for connection pooling.[-]`,

		"help.security.ldaprealm": `[::b]JNDIRealm (LDAP/Active Directory)[::-]
LDAP-based authentication and authorization.

[green]Key Attributes:[-]
• connectionURL: ldap://server:389 or ldaps://server:636
• userPattern: DN pattern, e.g., uid={0},ou=users,dc=example,dc=com
• userSearch: Search filter for finding users
• roleBase: Base DN for role searches
• roleName: Attribute containing role name

[green]Active Directory Example:[-]
• userPattern: {0}@domain.com
• userSearch: (sAMAccountName={0})
• roleSearch: (member={0})

[gray]Use connectionPoolSize for high-volume authentication.[-]`,

		// JNDI Help
		"help.jndi.datasource": `[::b]JNDI DataSource[::-]
Database connection pool accessible via JNDI lookup.

[green]Key Attributes:[-]
• [yellow]name[-]: JNDI name (e.g., jdbc/MyDB)
• [yellow]type[-]: javax.sql.DataSource
• [yellow]driverClassName[-]: JDBC driver class
• [yellow]url[-]: JDBC connection URL
• [yellow]username/password[-]: Database credentials

[green]Connection Pool Settings:[-]
• maxTotal: Maximum connections (default: 8)
• maxIdle: Maximum idle connections
• minIdle: Minimum idle connections
• maxWaitMillis: Wait timeout for connection

[green]Usage in Application:[-]
Context ctx = new InitialContext();
DataSource ds = (DataSource) ctx.lookup("java:comp/env/jdbc/MyDB");`,

		"help.jndi.environment": `[::b]Environment Entry[::-]
Simple values accessible via JNDI for configuration.

[green]Supported Types:[-]
• java.lang.String
• java.lang.Integer
• java.lang.Boolean
• java.lang.Double

[green]Example:[-]
<Environment name="maxItems" value="100" type="java.lang.Integer"/>

[green]Usage in Application:[-]
Context ctx = new InitialContext();
Integer maxItems = (Integer) ctx.lookup("java:comp/env/maxItems");

[gray]Useful for externalizing application configuration.[-]`,

		// Host Help
		"help.host": `[::b]Virtual Host[::-]
A Host represents a virtual host with its own applications.

[green]Key Attributes:[-]
• [yellow]name[-]: Host name (e.g., www.example.com)
• [yellow]appBase[-]: Application directory (default: webapps)
• [yellow]unpackWARs[-]: Auto-extract WAR files (default: true)
• [yellow]autoDeploy[-]: Hot-deploy applications (default: true)

[green]Aliases:[-]
Additional hostnames that map to this Host.
Example: example.com → www.example.com

[green]Directory Structure:[-]
• $CATALINA_BASE/webapps/ - Default appBase
• Each subdirectory or WAR is an application`,

		"help.host.appbase": `[::b]Application Base[::-]
Directory containing web applications for this Host.

[yellow]Default:[-] webapps

[green]Path Types:[-]
• Relative: Relative to $CATALINA_BASE
• Absolute: Full filesystem path

[green]Deployment Methods:[-]
• Drop WAR file in appBase
• Create directory with web application
• Use Context descriptor in conf/Catalina/[host]/`,

		"help.host.autodeploy": `[::b]Auto Deploy[::-]
Automatically deploy applications when files change.

[yellow]Default:[-] true

[green]Behavior:[-]
• Monitors appBase for new/modified files
• Deploys new applications automatically
• Reloads modified applications

[yellow]Production Recommendation:[-]
Set to "false" in production for:
• Better security (prevent unauthorized deployments)
• Improved performance (no file watching)
• Predictable deployment timing`,

		// Valve Help
		"help.valve.accesslog": `[::b]Access Log Valve[::-]
Logs HTTP requests to a file.

[green]Key Attributes:[-]
• [yellow]directory[-]: Log directory (default: logs)
• [yellow]prefix[-]: Log filename prefix (default: localhost_access_log)
• [yellow]suffix[-]: Log filename suffix (default: .txt)
• [yellow]pattern[-]: Log format pattern

[green]Pattern Variables:[-]
• %h - Remote hostname/IP
• %l - Remote logical username (always -)
• %u - Authenticated username
• %t - Date/time in Common Log Format
• %r - First line of request
• %s - HTTP status code
• %b - Bytes sent (- if zero)
• %D - Time to process request (ms)
• %T - Time to process request (sec)

[green]Common Patterns:[-]
• common: %h %l %u %t "%r" %s %b
• combined: + Referer and User-Agent`,

		"help.valve.remoteip": `[::b]Remote IP Valve[::-]
Extracts real client IP from proxy headers.

[green]Key Attributes:[-]
• [yellow]remoteIpHeader[-]: Header containing client IP (X-Forwarded-For)
• [yellow]protocolHeader[-]: Header for protocol (X-Forwarded-Proto)
• [yellow]internalProxies[-]: Regex for trusted proxy IPs

[green]Use Case:[-]
Behind reverse proxy (nginx, Apache, HAProxy):
• Real IP in X-Forwarded-For header
• request.getRemoteAddr() returns proxy IP

[green]Configuration:[-]
<Valve className="...RemoteIpValve"
       remoteIpHeader="X-Forwarded-For"
       protocolHeader="X-Forwarded-Proto"/>

[gray]Essential for accurate logging and security behind proxies.[-]`,

		"help.valve.remoteaddr": `[::b]Remote Address Filter[::-]
Restricts access based on client IP address.

[green]Attributes:[-]
• [yellow]allow[-]: Regex of allowed IPs
• [yellow]deny[-]: Regex of denied IPs

[green]Evaluation Order:[-]
1. Check deny patterns first
2. Check allow patterns
3. If no match: deny if allow specified, allow otherwise

[green]Examples:[-]
• Allow only localhost: allow="127\.\d+\.\d+\.\d+"
• Allow internal network: allow="192\.168\.\d+\.\d+"
• Block specific IP: deny="10\.0\.0\.1"

[yellow]Note:[-] Use with RemoteIpValve when behind proxy.`,

		"help.valve.sso": `[::b]Single Sign-On Valve[::-]
Enables single sign-on across applications.

[green]Behavior:[-]
• User logs in once
• Session shared across all apps in same Host
• Logout from one app logs out from all

[green]Configuration:[-]
Place at Host level:
<Valve className="...SingleSignOn"/>

[green]Requirements:[-]
• All apps must use same Realm
• Apps must be in same virtual Host
• Use secure cookies in production`,

		// Cluster Help
		"help.cluster": `[::b]Tomcat Clustering[::-]
Session replication across multiple Tomcat instances.

[green]Components:[-]
• [yellow]Manager[-]: Manages session replication
• [yellow]Channel[-]: Group communication
• [yellow]Membership[-]: Cluster member discovery
• [yellow]Receiver[-]: Receives replicated data
• [yellow]Sender[-]: Sends replicated data

[green]Manager Types:[-]
• [yellow]DeltaManager[-]: Replicates to all nodes (small clusters)
• [yellow]BackupManager[-]: Replicates to one backup (large clusters)

[green]Requirements:[-]
• <distributable/> in web.xml
• All session attributes must be Serializable
• Multicast network or static membership`,

		"help.cluster.membership": `[::b]Cluster Membership[::-]
Discovers and tracks cluster members.

[green]Multicast Membership:[-]
• address: Multicast group (default: 228.0.0.4)
• port: Multicast port (default: 45564)
• frequency: Heartbeat interval (ms)
• dropTime: Member dropout timeout (ms)

[green]Static Membership:[-]
For networks without multicast:
• Configure each member explicitly
• Use StaticMembershipInterceptor

[yellow]Network Requirements:[-]
• UDP multicast must be enabled
• Firewall must allow multicast traffic
• All nodes on same multicast group`,

		// Logging Help
		"help.logging": `[::b]JULI Logging[::-]
Tomcat uses Java Util Logging (JULI) by default.

[green]Configuration File:[-]
conf/logging.properties

[green]Handler Types:[-]
• FileHandler: Writes to rotating log files
• ConsoleHandler: Writes to console/stdout

[green]Log Levels (ascending):[-]
• FINEST - Most detailed
• FINER
• FINE
• CONFIG
• INFO
• WARNING
• SEVERE - Only errors

[green]Common Loggers:[-]
• org.apache.catalina - Tomcat core
• org.apache.coyote - Connectors
• org.apache.jasper - JSP engine`,

		// Context Help
		"help.context": `[::b]Context Configuration[::-]
Defines a web application and its settings.

[green]Key Attributes:[-]
• [yellow]path[-]: URL path (e.g., /myapp)
• [yellow]docBase[-]: Application directory or WAR
• [yellow]reloadable[-]: Auto-reload on class changes
• [yellow]crossContext[-]: Allow cross-app dispatching

[green]Locations:[-]
• $CATALINA_BASE/conf/context.xml - Global defaults
• $CATALINA_BASE/conf/[engine]/[host]/ - Per-app
• META-INF/context.xml - Embedded in application

[yellow]Production Settings:[-]
• reloadable="false" (performance)
• privileged="false" (security)`,

		"help.context.reloadable": `[::b]Reloadable[::-]
Automatically reload when classes change.

[yellow]Default:[-] false

[green]Behavior:[-]
• Monitors /WEB-INF/classes and /WEB-INF/lib
• Reloads application when changes detected
• Useful during development

[yellow]Production Recommendation:[-]
Set to "false" for:
• Better performance (no file monitoring)
• Stability (no unexpected reloads)
• Memory efficiency`,

		// Web.xml Help
		"help.webxml.servlet": `[::b]Servlet Configuration[::-]
Defines a servlet and its mappings.

[green]Elements:[-]
• [yellow]servlet-name[-]: Unique identifier
• [yellow]servlet-class[-]: Fully qualified class name
• [yellow]load-on-startup[-]: Initialization order (optional)
• [yellow]init-param[-]: Initialization parameters

[green]Mapping:[-]
• [yellow]url-pattern[-]: URL pattern to match
• Patterns: /exact, /path/*, *.extension

[green]Load Order:[-]
• Negative or absent: Load on first request
• 0 or positive: Load on startup (lower = earlier)`,

		"help.webxml.filter": `[::b]Filter Configuration[::-]
Filters process requests before servlets.

[green]Common Uses:[-]
• Authentication/Authorization
• Request/Response modification
• Logging and auditing
• Compression
• Character encoding

[green]Filter Chain:[-]
• Filters execute in order defined
• Each filter calls chain.doFilter()
• Can modify request before servlet
• Can modify response after servlet

[green]Mapping Types:[-]
• url-pattern: Match URL patterns
• servlet-name: Apply to specific servlet
• dispatcher: REQUEST, FORWARD, INCLUDE, ERROR`,

		"help.webxml.session": `[::b]Session Configuration[::-]
Configures HTTP session behavior.

[green]session-timeout:[-]
• Time in minutes before session expires
• -1 = sessions never timeout
• Default: 30 minutes

[green]cookie-config:[-]
• name: Cookie name (default: JSESSIONID)
• http-only: Prevent JavaScript access (recommended)
• secure: Only send over HTTPS
• max-age: Cookie lifetime in seconds

[green]tracking-mode:[-]
• COOKIE - Use cookies (recommended)
• URL - URL rewriting (security risk)
• SSL - SSL session ID`,

		"help.webxml.security": `[::b]Security Constraints[::-]
Defines protected resources and access rules.

[green]Components:[-]
• [yellow]web-resource-collection[-]: URLs to protect
• [yellow]auth-constraint[-]: Required roles
• [yellow]user-data-constraint[-]: Transport guarantee

[green]Transport Guarantee:[-]
• NONE - No encryption required
• INTEGRAL - Data integrity (HTTPS)
• CONFIDENTIAL - Confidentiality (HTTPS)

[green]Example:[-]
<security-constraint>
  <web-resource-collection>
    <url-pattern>/admin/*</url-pattern>
  </web-resource-collection>
  <auth-constraint>
    <role-name>admin</role-name>
  </auth-constraint>
</security-constraint>`,

		// Sub-menu help for Security
		"help.security.realm.current": `[::b]Current Realm[::-]
The currently configured authentication realm.

[green]Click to:[-]
• View realm configuration details
• Modify realm settings
• Configure nested realms (for LockOutRealm/CombinedRealm)`,

		"help.security.realm.set": `[::b]Set Realm Type[::-]
Choose and configure an authentication realm.

[green]Available Types:[-]
• UserDatabaseRealm - Uses tomcat-users.xml
• DataSourceRealm - Database via JNDI DataSource
• JNDIRealm - LDAP/Active Directory
• LockOutRealm - Prevents brute-force attacks
• CombinedRealm - Combines multiple realms`,

		"help.security.realm.remove": `[::b]Remove Realm[::-]
Removes the current realm configuration.

[yellow]Warning:[-]
Removing the realm will disable authentication.
Applications requiring security-constraints will fail.`,

		"help.security.realm.userdatabase": `[::b]UserDatabaseRealm[::-]
File-based authentication using tomcat-users.xml.

[green]Best for:[-]
• Development environments
• Small deployments
• Simple authentication needs

[yellow]Configuration:[-]
• Users defined in conf/tomcat-users.xml
• Supports role-based access control
• Easy to set up and manage`,

		"help.security.realm.datasource": `[::b]DataSourceRealm[::-]
Database-backed authentication via JNDI DataSource.

[green]Best for:[-]
• Production environments
• Large number of users
• Integration with existing databases

[yellow]Requirements:[-]
• Configure DataSource in context.xml
• Create users and roles tables
• Set table/column mappings`,

		"help.security.realm.jndi": `[::b]JNDIRealm (LDAP)[::-]
Authentication via LDAP or Active Directory.

[green]Best for:[-]
• Enterprise environments
• Centralized user management
• Active Directory integration

[yellow]Key Settings:[-]
• connectionURL: ldap://server:389
• userPattern or userSearch for user lookup
• roleBase and roleSearch for roles`,

		"help.security.realm.lockout": `[::b]LockOutRealm[::-]
Wrapper realm that prevents brute-force attacks.

[green]Features:[-]
• Locks user after failed login attempts
• Configurable lockout duration
• Works with any nested realm

[yellow]Settings:[-]
• failureCount: Attempts before lockout (default: 5)
• lockOutTime: Lockout duration in seconds`,

		"help.security.realm.combined": `[::b]CombinedRealm[::-]
Combines multiple realms into one.

[green]Use Cases:[-]
• Authenticate from multiple sources
• Fallback authentication
• Migration between realms

[yellow]Behavior:[-]
• Tries each nested realm in order
• First successful auth wins
• All nested realms must be configured`,

		"help.security.users.list": `[::b]Users[::-]
Manage user accounts in tomcat-users.xml.

[green]User Properties:[-]
• Username: Unique identifier
• Password: User credential (can be digested)
• Roles: Comma-separated list of roles

[yellow]Common Roles:[-]
• manager-gui: Access Manager web app
• manager-script: Access Manager text interface
• admin-gui: Access Host Manager`,

		"help.security.roles.list": `[::b]Roles[::-]
Define security roles in tomcat-users.xml.

[green]Purpose:[-]
• Group permissions together
• Assign multiple roles to users
• Reference in security-constraints

[yellow]Standard Roles:[-]
• manager-gui, manager-script, manager-jmx
• admin-gui, admin-script
• Custom roles for your applications`,

		// Sub-menu help for JNDI
		"help.jndi.mail": `[::b]Mail Session[::-]
JavaMail configuration for sending emails.

[green]Key Properties:[-]
• mail.smtp.host: SMTP server hostname
• mail.smtp.port: SMTP port (25, 587, 465)
• mail.smtp.auth: Enable authentication

[yellow]Usage:[-]
Session session = (Session) ctx.lookup("java:comp/env/mail/Session");
Transport.send(message);`,

		"help.jndi.resourcelink": `[::b]Resource Links[::-]
Links to resources defined in GlobalNamingResources.

[green]Purpose:[-]
• Share resources across applications
• Centralize configuration
• Reference global DataSources

[yellow]Example:[-]
<ResourceLink name="jdbc/MyDB"
              global="jdbc/GlobalDB"
              type="javax.sql.DataSource"/>`,

		// Sub-menu help for Host
		"help.host.virtualhost": `[::b]Virtual Hosts[::-]
Define virtual hosts for different domain names.

[green]Key Attributes:[-]
• name: Host name (e.g., www.example.com)
• appBase: Application directory
• autoDeploy: Hot-deploy applications

[yellow]Usage:[-]
• Single Tomcat serving multiple domains
• Separate applications per domain`,

		"help.host.context": `[::b]Contexts[::-]
Web applications deployed on hosts.

[green]Key Attributes:[-]
• path: URL path (e.g., /myapp)
• docBase: Application directory or WAR

[yellow]Deployment Methods:[-]
• Drop WAR in webapps/
• Create directory structure
• Use context descriptor`,

		"help.host.engine": `[::b]Engine Settings[::-]
Configure the Catalina Engine.

[green]Key Attributes:[-]
• name: Engine name (default: Catalina)
• defaultHost: Default virtual host
• jvmRoute: Node ID for load balancing

[yellow]jvmRoute Usage:[-]
Required for sticky sessions in clusters
Session ID format: <sessionId>.<jvmRoute>`,

		// Sub-menu help for Valve
		"help.valve.engine": `[::b]Engine Valves[::-]
Valves applied to all requests entering the Engine.

[green]Common Engine Valves:[-]
• AccessLogValve: Log all HTTP requests
• RemoteIpValve: Handle proxy headers
• ErrorReportValve: Customize error pages

[yellow]Execution Order:[-]
Engine valves run first, before Host and Context valves.`,

		"help.valve.host": `[::b]Host Valves[::-]
Valves applied to requests for specific virtual hosts.

[green]Use Cases:[-]
• Per-host access logging
• Host-specific IP filtering
• Single sign-on within host

[yellow]Scope:[-]
Only affects requests to this virtual host.`,

		"help.valve.context": `[::b]Context Valves[::-]
Valves applied to specific web applications.

[green]Use Cases:[-]
• Application-specific logging
• Request/response modification
• Custom security checks

[yellow]Scope:[-]
Only affects requests to this application.`,

		"help.valve.quickadd": `[::b]Quick Add Common Valves[::-]
Quickly add frequently used valve configurations.

[green]Available Templates:[-]
• Access Log Valve
• Remote IP Valve (for proxies)
• Remote Address Filter
• Single Sign-On

[yellow]Benefits:[-]
Pre-configured with recommended settings.`,

		// Sub-menu help for Cluster
		"help.cluster.status": `[::b]Cluster Status[::-]
Enable or disable session replication.

[green]When Enabled:[-]
• Session data replicated across nodes
• Failover without session loss
• Requires multicast network

[yellow]Requirements:[-]
• <distributable/> in web.xml
• Serializable session attributes
• Network multicast support`,

		"help.cluster.settings": `[::b]Cluster Settings[::-]
Basic cluster configuration options.

[green]Key Settings:[-]
• Cluster class name
• Channel timeout
• Session notification options

[yellow]Best Practices:[-]
• Use DeltaManager for small clusters
• Use BackupManager for large clusters`,

		"help.cluster.manager": `[::b]Session Manager[::-]
Controls how sessions are replicated.

[green]Manager Types:[-]
• [yellow]DeltaManager[-]: All-to-all replication
  - Best for small clusters (<5 nodes)
  - Higher network overhead

• [yellow]BackupManager[-]: Primary-backup replication
  - Better for large clusters
  - Lower network overhead`,

		"help.cluster.receiver": `[::b]Receiver[::-]
Receives replicated session data.

[green]Settings:[-]
• address: Listen address (auto-detect or specific)
• port: Listen port (default: 4000-4100)
• timeout: Receive timeout

[yellow]NIO vs BIO:[-]
NioReceiver recommended for performance.`,

		"help.cluster.sender": `[::b]Sender[::-]
Sends session data to other members.

[green]Settings:[-]
• Transport type: NIO or Pooled
• Timeout and retry settings

[yellow]Pooled Sender:[-]
Uses connection pooling for better performance.`,

		"help.cluster.interceptors": `[::b]Interceptors[::-]
Channel interceptors for message processing.

[green]Common Interceptors:[-]
• TcpFailureDetector: Detects member failures
• StaticMembershipInterceptor: Static member list
• MessageDispatch15Interceptor: Async delivery

[yellow]Order:[-]
Interceptors execute in configured order.`,

		"help.cluster.deployer": `[::b]Farm Deployer[::-]
Deploy applications across cluster.

[green]Features:[-]
• Deploy WAR to all nodes
• Synchronized deployment
• Undeploy from all nodes

[yellow]Settings:[-]
• tempDir: Temporary file directory
• deployDir: Deployment directory
• watchDir: Watch for new WARs`,

		// Sub-menu help for Logging
		"help.logging.filehandlers": `[::b]File Handlers[::-]
Write logs to rotating log files.

[green]Key Properties:[-]
• directory: Log directory (default: logs)
• prefix: Filename prefix
• suffix: Filename suffix
• level: Logging level

[yellow]Default Handlers:[-]
• catalina - Tomcat core logs
• localhost - Application logs
• manager/host-manager - Manager logs`,

		"help.logging.console": `[::b]Console Handler[::-]
Write logs to console/stdout.

[green]Properties:[-]
• level: Minimum log level
• formatter: Log format

[yellow]Usage:[-]
Useful for development and Docker environments.
In production, prefer file handlers.`,

		"help.logging.loggers": `[::b]Loggers[::-]
Configure logging for specific packages.

[green]Common Loggers:[-]
• org.apache.catalina - Tomcat core
• org.apache.coyote - Connectors
• org.apache.jasper - JSP engine
• org.apache.tomcat - Tomcat utilities

[yellow]Levels:[-]
SEVERE > WARNING > INFO > CONFIG > FINE > FINER > FINEST`,

		"help.logging.rootlogger": `[::b]Root Logger[::-]
Default logger for all packages.

[green]Configuration:[-]
• Assign handlers for root logger
• Set default logging level
• Affects all loggers without specific config

[yellow]Best Practice:[-]
Set root logger to INFO or WARNING for production.`,

		// Sub-menu help for Context
		"help.context.settings": `[::b]Context Settings[::-]
Basic configuration for this application.

[green]Key Attributes:[-]
• reloadable: Auto-reload on class changes
• crossContext: Allow dispatcher to other contexts
• privileged: Access Tomcat internals

[yellow]Production:[-]
Set reloadable=false for performance.`,

		"help.context.resources": `[::b]JNDI Resources[::-]
Application-specific JNDI resources.

[green]Resource Types:[-]
• DataSource: Database connections
• MailSession: Email configuration
• Custom resources

[yellow]Scope:[-]
Available only to this application via java:comp/env.`,

		"help.context.environment": `[::b]Environment Entries[::-]
Simple configuration values via JNDI.

[green]Supported Types:[-]
• String, Integer, Boolean, Double

[green]Usage:[-]
Context ctx = new InitialContext();
String value = (String) ctx.lookup("java:comp/env/myParam");`,

		"help.context.resourcelinks": `[::b]Resource Links[::-]
Links to global resources.

[green]Purpose:[-]
• Reference GlobalNamingResources
• Share DataSources across apps

[yellow]Example:[-]
<ResourceLink name="jdbc/MyDB" global="jdbc/SharedDB"/>`,

		"help.context.parameters": `[::b]Context Parameters[::-]
Application initialization parameters.

[green]Accessible via:[-]
• ServletContext.getInitParameter()
• ${initParam.name} in JSP

[yellow]Difference from web.xml:[-]
Context parameters override web.xml context-params.`,

		"help.context.watched": `[::b]Watched Resources[::-]
Files that trigger application reload when modified.

[green]Default Watched:[-]
• WEB-INF/web.xml
• WEB-INF/classes/
• WEB-INF/lib/

[yellow]Add Custom:[-]
<WatchedResource>WEB-INF/custom.properties</WatchedResource>`,

		"help.context.manager": `[::b]Session Manager[::-]
Configure session persistence and handling.

[green]Manager Types:[-]
• StandardManager: In-memory sessions
• PersistentManager: Session persistence

[yellow]Persistence:[-]
Save sessions to files or database for restart recovery.`,

		"help.context.cookie": `[::b]Cookie Processor[::-]
Configure session cookie behavior.

[green]Key Attributes:[-]
• sameSiteCookies: None, Lax, Strict
• Affects JSESSIONID cookie

[yellow]SameSite:[-]
• Strict: Most secure, may break some workflows
• Lax: Good balance (recommended)
• None: Required for cross-site (needs Secure)`,

		"help.context.jarscanner": `[::b]JAR Scanner[::-]
Configure class and annotation scanning.

[green]Settings:[-]
• scanClassPath: Scan classpath JARs
• scanManifest: Check manifests
• scanAllDirectories: Scan all dirs

[yellow]Performance:[-]
Disable scanning for faster startup if not using annotations.`,

		// Sub-menu help for Web.xml
		"help.web.servlets": `[::b]Servlets[::-]
HTTP request handlers for your application.

[green]Configuration:[-]
• servlet-name: Unique identifier
• servlet-class: Java class
• url-pattern: URL mapping
• load-on-startup: Initialization order

[yellow]Default Servlets:[-]
• default: Static content
• jsp: JSP processing`,

		"help.web.filters": `[::b]Filters[::-]
Process requests before/after servlets.

[green]Common Uses:[-]
• Character encoding
• Authentication
• Logging
• Compression

[yellow]Filter Chain:[-]
Filters execute in order defined in web.xml.`,

		"help.web.listeners": `[::b]Listeners[::-]
Respond to servlet container events.

[green]Listener Types:[-]
• ServletContextListener: App startup/shutdown
• HttpSessionListener: Session events
• ServletRequestListener: Request events

[yellow]Use Cases:[-]
Initialize resources, cleanup on shutdown.`,

		"help.web.session": `[::b]Session Config[::-]
Configure HTTP session behavior.

[green]Settings:[-]
• session-timeout: Inactivity timeout (minutes)
• cookie-config: Session cookie settings
• tracking-mode: COOKIE, URL, SSL

[yellow]Security:[-]
Use http-only and secure cookies in production.`,

		"help.web.welcomefiles": `[::b]Welcome Files[::-]
Default pages when accessing directories.

[green]Default Order:[-]
1. index.html
2. index.htm
3. index.jsp

[yellow]Behavior:[-]
First matching file in list is served.`,

		"help.web.errorpages": `[::b]Error Pages[::-]
Custom pages for error responses.

[green]Map by:[-]
• error-code: HTTP status (404, 500, etc.)
• exception-type: Java exception class

[yellow]Location:[-]
Path relative to web app root (e.g., /error/404.html)`,

		"help.web.mime": `[::b]MIME Mappings[::-]
Map file extensions to content types.

[green]Format:[-]
• extension: File extension (without dot)
• mime-type: Content type

[yellow]Examples:[-]
• json → application/json
• woff2 → font/woff2`,

		"help.web.security": `[::b]Security Constraints[::-]
Protect URLs and require authentication.

[green]Components:[-]
• web-resource-collection: URLs to protect
• auth-constraint: Required roles
• user-data-constraint: HTTPS requirement

[yellow]Transport:[-]
Use CONFIDENTIAL for HTTPS redirect.`,

		"help.web.login": `[::b]Login Config[::-]
Configure authentication method.

[green]Auth Methods:[-]
• BASIC: Browser popup
• FORM: Custom login page
• DIGEST: Hashed password
• CLIENT-CERT: Certificate auth

[yellow]FORM Login:[-]
Requires form-login-page and form-error-page.`,

		"help.web.roles": `[::b]Security Roles[::-]
Define roles referenced in constraints.

[green]Purpose:[-]
• Declare roles used by application
• Map to Realm-defined roles
• Document role usage

[yellow]Note:[-]
Roles must exist in the configured Realm.`,

		"help.web.contextparams": `[::b]Context Parameters[::-]
Application-wide initialization parameters.

[green]Access:[-]
• ServletContext.getInitParameter("name")
• ${initParam.name} in JSP/JSTL

[yellow]Use Cases:[-]
Database URLs, feature flags, configuration.`,

		// Additional common keys
		"help.security.users": `[::b]Users & Roles[::-]
Manage authentication users in tomcat-users.xml.

[green]Features:[-]
• Add/Edit/Delete users
• Assign roles to users
• Manage role definitions

[yellow]File Location:[-]
$CATALINA_BASE/conf/tomcat-users.xml`,

		"help.security.credential": `[::b]Credential Handler[::-]
Configure password hashing algorithm.

[green]Algorithms:[-]
• SHA-256: Recommended minimum
• SHA-512: More secure
• PBKDF2: Industry standard

[yellow]Settings:[-]
• iterations: Hash iterations (1000+ recommended)
• saltLength: Salt size in bytes`,

		// DataSource Property Help
		"help.ds.name": `[yellow]JNDI Name[white]

The JNDI name used to look up this DataSource.

[aqua]Example:[white]
  jdbc/MyDatabase

[aqua]Usage in code:[white]
  Context ctx = new InitialContext();
  DataSource ds = (DataSource)
    ctx.lookup("java:comp/env/jdbc/MyDatabase");`,

		"help.ds.auth": `[yellow]Authentication[white]

Specifies who manages the authentication.

[aqua]Container:[white]
  The container (Tomcat) manages sign-on
  to the resource. Credentials are stored
  in the DataSource configuration.

[aqua]Application:[white]
  The application provides credentials
  programmatically when getting connections.`,

		"help.ds.factory": `[yellow]Factory Class[white]

The JNDI object factory class.

[aqua]Default:[white]
  org.apache.tomcat.dbcp.dbcp2.
    BasicDataSourceFactory

[aqua]Other options:[white]
  • org.apache.commons.dbcp2.BasicDataSourceFactory
  • com.zaxxer.hikari.HikariJNDIFactory`,

		"help.ds.driver": `[yellow]JDBC Driver Class[white]

The fully qualified class name of the JDBC driver.

[aqua]Common drivers:[white]
  MySQL 8.x: com.mysql.cj.jdbc.Driver
  PostgreSQL: org.postgresql.Driver
  Oracle: oracle.jdbc.OracleDriver
  SQL Server: com.microsoft.sqlserver.
    jdbc.SQLServerDriver
  MariaDB: org.mariadb.jdbc.Driver`,

		"help.ds.url": `[yellow]JDBC URL[white]

The connection URL for the database.

[aqua]Format examples:[white]
  MySQL:
    jdbc:mysql://host:3306/dbname
  PostgreSQL:
    jdbc:postgresql://host:5432/dbname
  Oracle:
    jdbc:oracle:thin:@host:1521:SID
  SQL Server:
    jdbc:sqlserver://host:1433;
      databaseName=dbname`,

		"help.ds.username": `[yellow]Database Username[white]

The username for database authentication.

This user should have appropriate permissions
for your application's database operations.`,

		"help.ds.password": `[yellow]Database Password[white]

The password for database authentication.

[red]Security Note:[white]
Consider using encrypted passwords or
external secret management for
production environments.`,

		"help.ds.initialsize": `[yellow]Initial Pool Size[white]

Number of connections created when
the pool is initialized.

[aqua]Default:[white] 0
[aqua]Recommended:[white] 5-10

Set based on your application's
baseline connection needs.`,

		"help.ds.maxtotal": `[yellow]Maximum Total Connections[white]

Maximum number of active connections
in the pool.

[aqua]Default:[white] 8
[aqua]Recommended:[white] 20-100

Consider your database's max_connections
setting and number of app instances.`,

		"help.ds.maxidle": `[yellow]Maximum Idle Connections[white]

Maximum connections that can remain
idle in the pool.

[aqua]Default:[white] 8
[aqua]Recommended:[white] Same as Initial Size

Higher values keep connections ready
but consume more database resources.`,

		"help.ds.minidle": `[yellow]Minimum Idle Connections[white]

Minimum connections to keep idle
in the pool.

[aqua]Default:[white] 0
[aqua]Recommended:[white] 5-10

Ensures quick response for burst traffic.`,

		"help.ds.maxwait": `[yellow]Maximum Wait Time[white]

Maximum time (ms) to wait for a
connection from the pool.

[aqua]Default:[white] -1 (infinite)
[aqua]Recommended:[white] 10000-30000

Set a timeout to prevent threads from
blocking indefinitely.`,

		"help.ds.validationquery": `[yellow]Validation Query[white]

SQL query to validate connections before use.

[aqua]Examples by database:[white]
  MySQL/MariaDB: SELECT 1
  PostgreSQL: SELECT 1
  Oracle: SELECT 1 FROM DUAL
  SQL Server: SELECT 1
  H2/HSQLDB: SELECT 1`,

		"help.ds.testonborrow": `[yellow]Test On Borrow[white]

Validate connections before lending
them to the application.

[aqua]Default:[white] false
[aqua]Recommended:[white] true

Ensures application gets valid
connections but adds slight overhead.`,

		"help.ds.testwhileidle": `[yellow]Test While Idle[white]

Validate idle connections periodically
in the background.

[aqua]Default:[white] false
[aqua]Recommended:[white] true

Removes stale connections proactively
without impacting request latency.`,

		// Mail Session Property Help
		"help.mail.name": `[yellow]JNDI Name[white]

The JNDI name used to look up this Mail Session.

[aqua]Example:[white]
  mail/Session

[aqua]Usage in code:[white]
  Session session = (Session)
    ctx.lookup("java:comp/env/mail/Session");`,

		"help.mail.auth": `[yellow]Authentication[white]

Specifies who manages the authentication.

[aqua]Container:[white]
  Tomcat manages SMTP authentication.

[aqua]Application:[white]
  Application provides credentials.`,

		"help.mail.host": `[yellow]SMTP Host[white]

The hostname or IP address of the SMTP server.

[aqua]Examples:[white]
  • smtp.gmail.com
  • smtp.office365.com
  • localhost`,

		"help.mail.port": `[yellow]SMTP Port[white]

The port number for the SMTP server.

[aqua]Common ports:[white]
  • 25: Standard SMTP (unencrypted)
  • 465: SMTPS (SSL/TLS)
  • 587: Submission (STARTTLS)`,

		"help.mail.user": `[yellow]SMTP User[white]

The username for SMTP authentication.

Usually the email address or account name.`,

		"help.mail.protocol": `[yellow]Protocol[white]

The mail transport protocol.

[aqua]smtp:[white]
  Standard SMTP, optionally with STARTTLS

[aqua]smtps:[white]
  SMTP over SSL/TLS (implicit)`,

		"help.mail.smtpauth": `[yellow]SMTP Authentication[white]

Enable SMTP authentication.

[aqua]Default:[white] false

Set to true if your SMTP server requires
username/password authentication.`,

		"help.mail.starttls": `[yellow]StartTLS[white]

Enable STARTTLS encryption.

[aqua]Default:[white] false

Upgrades the connection to TLS after
initial plain text handshake.
Required by many modern SMTP servers.`,

		"help.mail.debug": `[yellow]Debug Mode[white]

Enable JavaMail debug output.

[aqua]Default:[white] false

Prints detailed protocol information
to System.out for troubleshooting.`,

		// Environment Property Help
		"help.env.name": `[yellow]JNDI Name[white]

The JNDI name for this environment entry.

[aqua]Example:[white]
  myapp/config/maxItems

[aqua]Usage in code:[white]
  Integer max = (Integer)
    ctx.lookup("java:comp/env/myapp/config/maxItems");`,

		"help.env.value": `[yellow]Value[white]

The value of this environment entry.

The value will be converted to the
specified type when looked up.`,

		"help.env.type": `[yellow]Type[white]

The Java type of this environment entry.

[aqua]Common types:[white]
  • java.lang.String
  • java.lang.Integer
  • java.lang.Boolean
  • java.lang.Double`,

		"help.env.override": `[yellow]Override[white]

Allow application to override this value.

[aqua]true:[white]
  Application can override via web.xml

[aqua]false:[white]
  Value is fixed by server configuration`,

		"help.env.description": `[yellow]Description[white]

Optional description of this entry.

Documents the purpose and usage of
this configuration value.`,

		// ResourceLink Property Help
		"help.reslink.name": `[yellow]Local Name[white]

The JNDI name used by the web application.

[aqua]Example:[white]
  jdbc/LocalDB

This is the name the application uses
to look up the resource.`,

		"help.reslink.global": `[yellow]Global Name[white]

The name of the global resource in server.xml.

[aqua]Example:[white]
  jdbc/GlobalDB

Links to a resource defined in
<GlobalNamingResources>.`,

		"help.reslink.type": `[yellow]Resource Type[white]

The Java type of the linked resource.

[aqua]Common types:[white]
  • javax.sql.DataSource
  • javax.mail.Session
  • org.apache.catalina.UserDatabase`,

		// Connector Property Help
		"help.conn.port": `[yellow]Port[white]

The TCP port number on which this
connector listens for requests.

[aqua]Common ports:[white]
  • 8080: HTTP (development)
  • 80: HTTP (production)
  • 8443/443: HTTPS
  • 8009: AJP`,

		"help.conn.protocol": `[yellow]Protocol[white]

The protocol handler implementation.

[aqua]HTTP:[white]
  • HTTP/1.1 (auto-detect NIO/APR)
  • org.apache.coyote.http11.Http11NioProtocol

[aqua]AJP:[white]
  • AJP/1.3
  • org.apache.coyote.ajp.AjpNioProtocol`,

		"help.conn.timeout": `[yellow]Connection Timeout[white]

Milliseconds to wait after connection
for first request data.

[aqua]Default:[white] 60000 (60 seconds)
[aqua]Recommended:[white] 20000-60000

Lower values free resources faster
from slow clients.`,

		"help.conn.redirectport": `[yellow]Redirect Port[white]

Port to redirect to when SSL is required.

[aqua]Default:[white] 8443

Used when a request requires security
but arrived on a non-SSL connector.`,

		"help.conn.maxthreads": `[yellow]Max Threads[white]

Maximum threads for request processing.

[aqua]Default:[white] 200
[aqua]Recommended:[white] 200-800

Higher values handle more concurrent
requests but use more memory.`,

		"help.conn.minsparethreads": `[yellow]Min Spare Threads[white]

Minimum threads kept running.

[aqua]Default:[white] 10
[aqua]Recommended:[white] 25-50

Higher values improve response time
for sudden traffic spikes.`,

		"help.conn.acceptcount": `[yellow]Accept Count[white]

Maximum queue length for incoming
connections when all threads are busy.

[aqua]Default:[white] 100

Connections beyond this are refused.`,

		"help.conn.executor": `[yellow]Executor[white]

Name of a shared thread pool executor.

Leave empty to use connector-specific
thread pool, or specify an executor
name defined in the Service.`,

		"help.conn.sslenabled": `[yellow]SSL Enabled[white]

Enable SSL/TLS for this connector.

Requires keystore configuration.`,

		"help.conn.scheme": `[yellow]Scheme[white]

Protocol scheme for request URLs.

[aqua]Values:[white]
  • http: Non-secure connections
  • https: Secure SSL/TLS connections`,

		"help.conn.secure": `[yellow]Secure[white]

Mark requests as secure.

Set to true for SSL/TLS connectors
so request.isSecure() returns true.`,

		"help.conn.keystorefile": `[yellow]Keystore File[white]

Path to the SSL keystore file.

[aqua]Example:[white]
  conf/localhost-rsa.jks
  ${catalina.base}/conf/keystore.p12`,

		"help.conn.keystorepass": `[yellow]Keystore Password[white]

Password for the keystore file.

[red]Security:[white]
Consider using external secret
management in production.`,

		"help.conn.keystoretype": `[yellow]Keystore Type[white]

Type of the keystore file.

[aqua]Types:[white]
  • JKS: Java KeyStore (legacy)
  • PKCS12: Modern standard (recommended)`,

		"help.conn.sslprotocol": `[yellow]SSL Protocol[white]

SSL/TLS protocol version.

[aqua]Recommended:[white] TLS
[aqua]Specific versions:[white] TLSv1.2, TLSv1.3

Avoid SSLv3 and TLSv1.0/1.1.`,

		"help.conn.clientauth": `[yellow]Client Auth[white]

Client certificate authentication mode.

[aqua]false:[white] No client cert required
[aqua]want:[white] Request but don't require
[aqua]true:[white] Require client certificate`,

		"help.conn.secret": `[yellow]AJP Secret[white]

Shared secret for AJP authentication.

Required when secretRequired is true.
Must match the web server configuration.`,

		"help.conn.secretrequired": `[yellow]Secret Required[white]

Require AJP secret for connections.

[aqua]Default:[white] true (Tomcat 9.0.31+)

Set to false only in trusted networks.`,

		"help.conn.address": `[yellow]Address[white]

IP address to bind this connector to.

Leave empty to bind to all interfaces.

[aqua]Examples:[white]
  • 127.0.0.1 (localhost only)
  • 0.0.0.0 (all interfaces)`,

		// Executor Property Help
		"help.exec.name": `[yellow]Executor Name[white]

Unique name for this thread pool.

Used by connectors to reference
this shared executor.

[aqua]Example:[white]
  tomcatThreadPool`,

		"help.exec.classname": `[yellow]Class Name[white]

The executor implementation class.

[aqua]Standard:[white]
  org.apache.catalina.core.
    StandardThreadExecutor

[aqua]Virtual Threads (Java 21+):[white]
  org.apache.catalina.core.
    StandardVirtualThreadExecutor`,

		"help.exec.nameprefix": `[yellow]Name Prefix[white]

Prefix for thread names.

Useful for identifying threads
in logs and thread dumps.

[aqua]Example:[white]
  catalina-exec-`,

		"help.exec.maxthreads": `[yellow]Max Threads[white]

Maximum threads in the pool.

[aqua]Default:[white] 200
[aqua]Recommended:[white] 200-800`,

		"help.exec.minsparethreads": `[yellow]Min Spare Threads[white]

Minimum idle threads kept alive.

[aqua]Default:[white] 25`,

		"help.exec.maxidletime": `[yellow]Max Idle Time[white]

Milliseconds before an idle thread
is terminated.

[aqua]Default:[white] 60000 (1 minute)`,

		"help.exec.maxqueuesize": `[yellow]Max Queue Size[white]

Maximum pending requests in queue.

[aqua]Default:[white] Integer.MAX_VALUE

Lower values reject requests earlier
when overloaded.`,

		"help.exec.prestartminsparethreads": `[yellow]Prestart Min Spare Threads[white]

Start minimum threads at startup.

[aqua]Default:[white] false

Set true to avoid cold start latency.`,

		// Server Settings Property Help
		"help.server.address": `[yellow]Shutdown Address[white]

Address for shutdown listener.

[aqua]Default:[white] localhost

Never bind to external interfaces.`,

		// Listener Property Help
		"help.listener.classname": `[yellow]Class Name[white]

Fully qualified class name of the listener.

[aqua]Common Listeners:[white]
  • VersionLoggerListener
  • AprLifecycleListener
  • JreMemoryLeakPreventionListener
  • ThreadLocalLeakPreventionListener`,

		"help.listener.sslengine": `[yellow]SSL Engine[white]

SSL engine for APR/native connector.

[aqua]Values:[white]
  • on: Enable OpenSSL engine
  • off: Use JSSE`,

		// Service/Engine Property Help
		"help.service.name": `[yellow]Service Name[white]

Name of this service container.

[aqua]Default:[white] Catalina

Multiple services can exist in one
server for different configurations.`,

		"help.engine.name": `[yellow]Engine Name[white]

Name of this Catalina engine.

[aqua]Default:[white] Catalina

Used in logging and JMX.`,

		"help.engine.defaulthost": `[yellow]Default Host[white]

Host to use for unmatched requests.

[aqua]Default:[white] localhost

Must match a configured <Host> name.`,

		"help.engine.jvmroute": `[yellow]JVM Route[white]

Route ID for session affinity.

Used by load balancers to identify
this Tomcat instance.

[aqua]Example:[white] node1`,

		// Host Property Help
		"help.host.name": `[yellow]Host Name[white]

Virtual host name (domain).

[aqua]Examples:[white]
  • localhost
  • www.example.com
  • *.example.com (wildcard)`,

		"help.host.unpackwars": `[yellow]Unpack WARs[white]

Extract WAR files before running.

[aqua]Default:[white] true

Unpacked apps start faster.`,

		"help.host.deployonstart": `[yellow]Deploy On Startup[white]

Deploy applications when Tomcat starts.

[aqua]Default:[white] true`,

		// Context Property Help
		"help.context.path": `[yellow]Context Path[white]

URL path for this application.

[aqua]Examples:[white]
  • "" (ROOT application)
  • /myapp
  • /api/v1`,

		"help.context.docbase": `[yellow]Document Base[white]

Path to application files.

Can be WAR file or directory.
Relative to Host's appBase.`,

		"help.context.crosscontext": `[yellow]Cross Context[white]

Allow access to other contexts.

[aqua]Default:[white] false

Enables getContext() calls.`,

		"help.context.cookies": `[yellow]Cookies[white]

Use cookies for session tracking.

[aqua]Default:[white] true`,

		"help.context.privileged": `[yellow]Privileged[white]

Access Tomcat internal classes.

[aqua]Default:[white] false

Required for manager apps.`,

		// Valve Property Help
		"help.valve.classname": `[yellow]Valve Class Name[white]

The valve implementation class.

[aqua]Common Valves:[white]
  • AccessLogValve
  • RemoteAddrValve
  • RemoteIpValve
  • ErrorReportValve
  • StuckThreadDetectionValve`,

		"help.valve.accesslog.pattern": `[yellow]Log Pattern[white]

Access log format pattern.

[aqua]Common patterns:[white]
  %h: Remote host
  %l: Remote user (identd)
  %u: Authenticated user
  %t: Date/time
  %r: Request line
  %s: Status code
  %b: Bytes sent

[aqua]Combined format:[white]
  %h %l %u %t "%r" %s %b`,

		"help.valve.accesslog.directory": `[yellow]Directory[white]

Log file directory.

[aqua]Default:[white] logs

Relative to CATALINA_BASE.`,

		"help.valve.accesslog.prefix": `[yellow]Prefix[white]

Log file name prefix.

[aqua]Default:[white] localhost_access_log`,

		"help.valve.accesslog.suffix": `[yellow]Suffix[white]

Log file name suffix.

[aqua]Default:[white] .txt`,

		"help.valve.accesslog.rotate": `[yellow]Rotatable[white]

Enable daily log rotation.

[aqua]Default:[white] true`,

		"help.valve.remoteaddr.allow": `[yellow]Allow Pattern[white]

Regex for allowed IP addresses.

[aqua]Examples:[white]
  • 127\\.0\\.0\\.1
  • 192\\.168\\.\\d+\\.\\d+
  • 10\\.\\d+\\.\\d+\\.\\d+`,

		"help.valve.remoteaddr.deny": `[yellow]Deny Pattern[white]

Regex for denied IP addresses.

Checked after allow pattern.`,

		"help.valve.stuckthread.threshold": `[yellow]Threshold[white]

Seconds before thread is stuck.

[aqua]Default:[white] 600 (10 minutes)

Logs warning when exceeded.`,

		// Realm Property Help
		"help.realm.classname": `[yellow]Realm Class Name[white]

Authentication realm implementation.

[aqua]Common Realms:[white]
  • UserDatabaseRealm (file-based)
  • DataSourceRealm (JDBC)
  • JNDIRealm (LDAP)
  • JAASRealm (JAAS)
  • CombinedRealm (multiple)`,

		"help.realm.userdatabase.resource": `[yellow]Resource Name[white]

JNDI name of UserDatabase resource.

[aqua]Default:[white] UserDatabase

Defined in GlobalNamingResources.`,

		"help.realm.datasource.name": `[yellow]DataSource Name[white]

JNDI name of the DataSource.

[aqua]Example:[white]
  jdbc/AuthDB`,

		"help.realm.datasource.usertable": `[yellow]User Table[white]

Table containing user credentials.

[aqua]Default:[white] users`,

		"help.realm.datasource.usernameCol": `[yellow]Username Column[white]

Column name for username.

[aqua]Default:[white] user_name`,

		"help.realm.datasource.passwordCol": `[yellow]Password Column[white]

Column name for password.

[aqua]Default:[white] user_pass`,

		"help.realm.datasource.roletable": `[yellow]Role Table[white]

Table containing user roles.

[aqua]Default:[white] user_roles`,

		"help.realm.datasource.rolenameCol": `[yellow]Role Name Column[white]

Column name for role name.

[aqua]Default:[white] role_name`,

		"help.realm.jndi.connectionURL": `[yellow]Connection URL[white]

LDAP server URL.

[aqua]Example:[white]
  ldap://ldap.example.com:389`,

		"help.realm.jndi.userbase": `[yellow]User Base[white]

Base DN for user searches.

[aqua]Example:[white]
  ou=users,dc=example,dc=com`,

		"help.realm.jndi.userpattern": `[yellow]User Pattern[white]

DN pattern for direct user lookup.

[aqua]Example:[white]
  uid={0},ou=users,dc=example,dc=com`,

		"help.realm.jndi.rolebase": `[yellow]Role Base[white]

Base DN for role searches.

[aqua]Example:[white]
  ou=groups,dc=example,dc=com`,

		// User/Role Property Help
		"help.user.username": `[yellow]Username[white]

Unique identifier for the user.

Used for authentication and
referenced in role assignments.`,

		"help.user.password": `[yellow]Password[white]

User's password.

[red]Security:[white]
Use digested passwords in production.
See CredentialHandler configuration.`,

		"help.user.roles": `[yellow]Roles[white]

Comma-separated list of roles.

[aqua]Common roles:[white]
  • manager-gui
  • admin-gui
  • manager-script`,

		"help.role.name": `[yellow]Role Name[white]

Unique identifier for the role.

Referenced in web.xml security
constraints and user assignments.`,

		"help.role.description": `[yellow]Description[white]

Optional role description.

Documents the purpose and
permissions of this role.`,

		"help.default": `[gray]Select a field to see help information.[-]`,
	},

	Korean: {
		// Common
		"app.title":            "TomcatKit - Tomcat 설정 관리자",
		"app.status.ready":     "준비됨",
		"app.status.saved":     "설정이 저장되었습니다",
		"app.status.error":     "오류: %s",
		"app.lang.select":      "언어 선택",
		"app.lang.current":     "현재: %s",
		"common.back":          "뒤로",
		"common.cancel":        "취소",
		"common.save":          "설정 저장",
		"common.apply":         "적용",
		"common.delete":        "삭제",
		"common.add":           "추가",
		"common.edit":          "편집",
		"common.yes":           "예",
		"common.no":            "아니오",
		"common.confirm":       "확인",
		"common.warning":       "경고",
		"common.error":         "오류",
		"common.success":       "성공",
		"common.loading":       "로딩중...",
		"common.return":        "메인 메뉴로 돌아가기",
		"common.enabled":       "활성화됨",
		"common.disabled":      "비활성화됨",
		"common.notconfigured": "설정되지 않음",
		"common.minutes":       "분",
		"help.title":           "도움말",
		"preview.title":        "XML 미리보기",
		"preview.properties":   "Properties 미리보기",

		// Main Menu
		"menu.title":               "메인 메뉴",
		"menu.server":              "서버",
		"menu.server.desc":         "server.xml 핵심 설정",
		"menu.connector":           "커넥터",
		"menu.connector.desc":      "HTTP, AJP, SSL/TLS 커넥터",
		"menu.security":            "보안 / Realm",
		"menu.security.desc":       "인증 영역 및 사용자 관리",
		"menu.jndi":                "JNDI 리소스",
		"menu.jndi.desc":           "DataSource, Mail Session, 환경변수",
		"menu.host":                "가상 호스트",
		"menu.host.desc":           "Host, Context, 세션 매니저",
		"menu.valve":               "밸브",
		"menu.valve.desc":          "AccessLog, RemoteAddr, SSO 밸브",
		"menu.cluster":             "클러스터링",
		"menu.cluster.desc":        "세션 복제, 멤버십",
		"menu.logging":             "로깅",
		"menu.logging.desc":        "JULI logging.properties",
		"menu.context":             "Context",
		"menu.context.desc":        "context.xml 설정",
		"menu.web":                 "웹 애플리케이션",
		"menu.web.desc":            "web.xml 서블릿, 필터, 보안",
		"menu.quicktemplates":      "빠른 템플릿",
		"menu.quicktemplates.desc": "일반 설정을 빠르게 적용",
		"menu.exit":                "종료",
		"menu.exit.desc":           "TomcatKit 종료",

		// Footer
		"footer.navigate": "이동",
		"footer.select":   "선택",
		"footer.back":     "뒤로",
		"footer.lang":     "언어",
		"footer.quit":     "종료",

		// Quick Templates
		"qt.title":              "빠른 템플릿",
		"qt.select":             "적용할 빠른 템플릿을 선택하세요",
		"qt.virtualthread":      "가상 스레드",
		"qt.virtualthread.desc": "가상 스레드 Executor 활성화 (Java 21+, Tomcat 11+)",
		"qt.https":              "HTTPS와 SSL",
		"qt.https.desc":         "SSL/TLS로 HTTPS 커넥터 구성",
		"qt.connpool":           "커넥션 풀 튜닝",
		"qt.connpool.desc":      "스레드 풀 설정 최적화",
		"qt.gzip":               "Gzip 압축",
		"qt.gzip.desc":          "응답 압축 활성화",
		"qt.accesslog":          "접근 로그",
		"qt.accesslog.desc":     "접근 로깅 활성화",
		"qt.security":           "보안 강화",
		"qt.security.desc":      "보안 모범 사례 적용",
		"qt.apache":             "Apache httpd (mod_jk/AJP)",
		"qt.apache.desc":        "Apache httpd용 AJP 커넥터 구성",
		"qt.nginx":              "nginx 리버스 프록시",
		"qt.nginx.desc":         "nginx proxy_pass 구성",
		"qt.haproxy":            "HAProxy 로드밸런서",
		"qt.haproxy.desc":       "HAProxy 로드밸런싱 구성",

		// Virtual Thread Template
		"qt.vt.title":            "가상 스레드 템플릿",
		"qt.vt.info":             "가상 스레드 Executor (Java 21+, Tomcat 11+)\n\n가상 스레드는 I/O 바운드 애플리케이션의 처리량을\n크게 향상시킬 수 있는 경량 스레드입니다.",
		"qt.vt.requirements":     "요구사항:",
		"qt.vt.req.java":         "Java 21 이상",
		"qt.vt.req.tomcat":       "Tomcat 11.0 이상 (또는 Tomcat 10.1.25+)",
		"qt.vt.willdo":           "이 템플릿은 다음을 수행합니다:",
		"qt.vt.willdo.create":    "가상 스레드 Executor 생성",
		"qt.vt.willdo.configure": "HTTP 커넥터에 Executor 연결",
		"qt.vt.warning.exists":   "경고: 가상 스레드 Executor가 이미 존재합니다!",
		"qt.vt.executor.name":    "Executor 이름",
		"qt.vt.thread.prefix":    "스레드 이름 접두사",
		"qt.vt.max.queue":        "최대 큐 크기",
		"qt.vt.apply.connector":  "적용할 커넥터",
		"qt.vt.apply":            "템플릿 적용",
		"qt.vt.success":          "가상 스레드 Executor가 적용되었습니다!",

		// HTTPS Template
		"qt.https.title":         "HTTPS 템플릿",
		"qt.https.info":          "HTTPS 커넥터 설정\n\n이 템플릿은 SSL/TLS로 HTTPS 커넥터를 생성합니다.",
		"qt.https.need":          "필요한 것:",
		"qt.https.need.keystore": "키스토어 파일 (.jks 또는 .p12)",
		"qt.https.need.password": "키스토어 비밀번호",
		"qt.https.port":          "HTTPS 포트",
		"qt.https.keystore.file": "키스토어 파일",
		"qt.https.keystore.pass": "키스토어 비밀번호",
		"qt.https.keystore.type": "키스토어 유형",
		"qt.https.success":       "HTTPS 커넥터가 추가되었습니다!",

		// Connection Pool Template
		"qt.pool.title":           "커넥션 풀 튜닝",
		"qt.pool.info":            "커넥션 풀 최적화\n\n더 나은 성능을 위해 스레드 풀 설정을 조정합니다.",
		"qt.pool.recommended":     "권장 설정:",
		"qt.pool.dev":             "개발: 25-100 스레드",
		"qt.pool.prod":            "운영: 150-400 스레드",
		"qt.pool.high":            "고트래픽: 400-800 스레드",
		"qt.pool.profile":         "프로필",
		"qt.pool.maxthreads":      "최대 스레드",
		"qt.pool.minsparethreads": "최소 여유 스레드",
		"qt.pool.acceptcount":     "Accept Count",
		"qt.pool.conntimeout":     "연결 타임아웃 (ms)",
		"qt.pool.apply":           "모든 HTTP 커넥터에 적용",
		"qt.pool.success":         "커넥션 풀 설정이 적용되었습니다!",

		// Gzip Template
		"qt.gzip.title":   "Gzip 압축",
		"qt.gzip.info":    "Gzip 압축\n\n텍스트 기반 콘텐츠에 대한 응답 압축을 활성화합니다.",
		"qt.gzip.willdo":  "HTTP 커넥터에 압축 속성을 추가합니다:",
		"qt.gzip.minsize": "최소 압축 크기 (바이트)",
		"qt.gzip.success": "Gzip 압축이 활성화되었습니다!",

		// Access Log Template
		"qt.al.title":            "접근 로그 템플릿",
		"qt.al.info":             "접근 로그 밸브\n\nHTTP 요청에 대한 접근 로깅을 구성합니다.",
		"qt.al.patterns":         "일반 패턴:",
		"qt.al.pattern.common":   "common: %h %l %u %t \"%r\" %s %b",
		"qt.al.pattern.combined": "combined: %h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-Agent}i\"",
		"qt.al.pattern":          "로그 패턴",
		"qt.al.directory":        "디렉토리",
		"qt.al.prefix":           "파일 접두사",
		"qt.al.suffix":           "파일 접미사",
		"qt.al.success":          "접근 로그가 구성되었습니다!",

		// Security Hardening Template
		"qt.sec.title":             "보안 강화",
		"qt.sec.info":              "보안 강화\n\nTomcat 설정에 보안 모범 사례를 적용합니다.",
		"qt.sec.willdo":            "이 템플릿은 다음을 수행합니다:",
		"qt.sec.willdo.shutdown":   "기본 shutdown 포트를 8005에서 -1로 변경 (비활성화)",
		"qt.sec.willdo.command":    "shutdown 명령 변경",
		"qt.sec.willdo.version":    "오류 페이지에서 서버 버전 제거",
		"qt.sec.willdo.listener":   "보안 관련 리스너 추가",
		"qt.sec.disable.shutdown":  "Shutdown 포트 비활성화",
		"qt.sec.remove.serverinfo": "오류에서 서버 정보 제거",
		"qt.sec.add.listener":      "보안 리스너 추가",
		"qt.sec.success":           "보안 강화가 적용되었습니다!",

		// Apache httpd Template
		"qt.ajp.title":          "Apache httpd (AJP) 템플릿",
		"qt.ajp.info":           "Apache httpd 연동 (mod_jk / mod_proxy_ajp)\n\nApache httpd 리버스 프록시용 AJP 커넥터를 구성합니다.",
		"qt.ajp.modules":        "지원되는 Apache 모듈:",
		"qt.ajp.modjk":          "mod_jk: 전통적인 Tomcat 커넥터",
		"qt.ajp.modproxy":       "mod_proxy_ajp: AJP용 Apache 프록시 모듈",
		"qt.ajp.willdo":         "이 템플릿은 다음을 수행합니다:",
		"qt.ajp.willdo.create":  "지정된 포트에 AJP/1.3 커넥터 생성",
		"qt.ajp.willdo.secret":  "시크릿 인증 구성 (Tomcat 9.0.31+)",
		"qt.ajp.willdo.valve":   "클라이언트 IP 처리를 위한 RemoteIpValve 추가",
		"qt.ajp.port":           "AJP 포트",
		"qt.ajp.address":        "바인드 주소",
		"qt.ajp.secret":         "AJP 시크릿",
		"qt.ajp.remoteip":       "RemoteIpValve 추가",
		"qt.ajp.success":        "Apache httpd AJP 커넥터가 구성되었습니다!",
		"qt.ajp.config.title":   "Apache httpd 설정",
		"qt.ajp.config.applied": "Apache httpd 설정이 적용되었습니다!",
		"qt.ajp.config.created": "AJP 커넥터가 포트 %s에 생성됨",
		"qt.ajp.config.copy":    "다음 설정을 Apache httpd에 복사하세요:",
		"qt.ajp.config.option1": "옵션 1: mod_proxy_ajp (권장)",
		"qt.ajp.config.option2": "옵션 2: mod_jk (workers.properties)",
		"qt.ajp.config.note":    "참고: 설정 적용 후 Apache httpd를 재시작하세요.",

		// nginx Template
		"qt.nginx.title":          "nginx 리버스 프록시 템플릿",
		"qt.nginx.info":           "nginx 리버스 프록시 설정\n\nnginx proxy_pass용 Tomcat을 구성합니다.",
		"qt.nginx.willdo":         "이 템플릿은 다음을 수행합니다:",
		"qt.nginx.willdo.valve":   "X-Forwarded-* 헤더용 RemoteIpValve 추가",
		"qt.nginx.willdo.ip":      "적절한 클라이언트 IP 처리 구성",
		"qt.nginx.willdo.http":    "선택적으로 HTTP 커넥터 설정 조정",
		"qt.nginx.proxy.note":     "nginx가 Tomcat의 HTTP 커넥터로 프록시합니다",
		"qt.nginx.connector":      "HTTP 커넥터",
		"qt.nginx.internal":       "내부 프록시 (정규식)",
		"qt.nginx.proto":          "X-Forwarded-Proto 처리",
		"qt.nginx.success":        "nginx 리버스 프록시가 구성되었습니다!",
		"qt.nginx.config.title":   "nginx 설정",
		"qt.nginx.config.applied": "nginx 리버스 프록시 설정이 적용되었습니다!",
		"qt.nginx.config.valve":   "nginx 프록시용 RemoteIpValve 구성됨",
		"qt.nginx.config.copy":    "다음 설정을 nginx에 복사하세요:",
		"qt.nginx.config.basic":   "기본 설정",
		"qt.nginx.config.https":   "HTTPS 설정 (SSL 종료)",
		"qt.nginx.config.note":    "참고: 설정 적용 후 nginx를 재시작하세요.",

		// HAProxy Template
		"qt.haproxy.title":           "HAProxy 로드밸런서 템플릿",
		"qt.haproxy.info":            "HAProxy 로드밸런서 설정\n\nHAProxy 로드밸런싱용 Tomcat을 구성합니다.",
		"qt.haproxy.willdo":          "이 템플릿은 다음을 수행합니다:",
		"qt.haproxy.willdo.valve":    "X-Forwarded-* 헤더용 RemoteIpValve 추가",
		"qt.haproxy.willdo.jvm":      "스티키 세션용 jvmRoute 구성 (선택)",
		"qt.haproxy.willdo.ip":       "적절한 클라이언트 IP 처리 설정",
		"qt.haproxy.modes":           "지원되는 로드밸런싱 모드:",
		"qt.haproxy.mode.http":       "HTTP 모드 (Layer 7) - 권장",
		"qt.haproxy.mode.tcp":        "TCP 모드 (Layer 4) - SSL 패스스루용",
		"qt.haproxy.connector":       "HTTP 커넥터",
		"qt.haproxy.sticky":          "스티키 세션 활성화",
		"qt.haproxy.jvmroute":        "JVM Route (노드 ID)",
		"qt.haproxy.internal":        "내부 프록시 (정규식)",
		"qt.haproxy.success":         "HAProxy 로드밸런서가 구성되었습니다!",
		"qt.haproxy.config.title":    "HAProxy 설정",
		"qt.haproxy.config.applied":  "HAProxy 로드밸런서 설정이 적용되었습니다!",
		"qt.haproxy.config.valve":    "HAProxy용 RemoteIpValve 구성됨",
		"qt.haproxy.config.sticky":   "스티키 세션 참고:",
		"qt.haproxy.config.jvmroute": "Tomcat Engine에 jvmRoute \"%s\" 구성됨",
		"qt.haproxy.config.cookie":   "HAProxy가 세션 어피니티에 JSESSIONID 쿠키 사용",
		"qt.haproxy.config.format":   "세션 ID 형식: <session-id>.<jvmRoute>",
		"qt.haproxy.config.copy":     "다음 설정을 HAProxy에 복사하세요:",
		"qt.haproxy.config.http":     "HTTP 모드 (Layer 7) - 권장",
		"qt.haproxy.config.tcp":      "TCP 모드 (Layer 4) - SSL 패스스루",
		"qt.haproxy.config.health":   "헬스체크 엔드포인트 (선택)",
		"qt.haproxy.config.stats":    "HAProxy 통계 (선택)",
		"qt.haproxy.config.note":     "참고: 설정 적용 후 HAProxy를 재시작하세요.",

		// Continue prompt
		"prompt.continue": "계속하려면 Enter 또는 Escape를 누르세요",
		"prompt.nohttp":   "HTTP 커넥터를 찾을 수 없습니다!",

		// Instance Selection
		"instance.title":           "Tomcat 인스턴스 선택",
		"instance.recent":          "최근 인스턴스",
		"instance.detected":        "감지된 인스턴스",
		"instance.none":            "Tomcat 설치를 찾을 수 없습니다",
		"instance.manual":          "경로 직접 입력",
		"instance.manual.desc":     "CATALINA_HOME 경로 지정",
		"instance.running":         "실행중",
		"instance.noselected":      "Tomcat 인스턴스가 선택되지 않았습니다",
		"instance.pressT":          "'t'를 눌러 Tomcat 인스턴스를 선택하세요",
		"instance.path.title":      "Tomcat 경로 입력",
		"instance.path.home":       "CATALINA_HOME",
		"instance.path.base":       "CATALINA_BASE (선택)",
		"instance.path.validate":   "검증 및 선택",
		"instance.path.required":   "CATALINA_HOME은 필수입니다",
		"instance.path.invalid":    "잘못된 경로: server.xml을 찾을 수 없습니다",
		"instance.selected":        "Tomcat 인스턴스가 선택되었습니다",
		"instance.info":            "Tomcat 인스턴스",
		"instance.version":         "버전",
		"instance.status":          "상태",
		"instance.stopped":         "중지됨",
		"instance.ready":           "설정 준비됨",
		"instance.path.help.home":  "CATALINA_HOME: Tomcat 설치 디렉토리 (bin, lib, conf 포함)",
		"instance.path.help.base":  "CATALINA_BASE: 인스턴스 디렉토리 (선택, 기본값은 CATALINA_HOME)",
		"instance.path.help.xml":   "경로에 conf/server.xml이 있어야 합니다",
		"instance.info.noselected": "Tomcat 인스턴스가 선택되지 않았습니다",
		"instance.info.getstarted": "시작하려면:",
		"instance.info.step1":      "'t'를 눌러 Tomcat 인스턴스를 선택하세요",
		"instance.info.step2":      "또는 실행: tomcatkit -home /path/to/tomcat",
		"instance.info.autodetect": "TomcatKit이 설치된 Tomcat 인스턴스를 자동 감지합니다.",

		// Server View
		"server.title":                      "서버 설정",
		"server.port":                       "Shutdown 포트",
		"server.port.desc":                  "서버 종료 리스너 포트",
		"server.shutdown":                   "Shutdown 명령",
		"server.listeners":                  "리스너",
		"server.listeners.desc":             "라이프사이클 리스너",
		"server.services":                   "서비스",
		"server.services.desc":              "서비스 설정",
		"server.globalresources":            "전역 리소스",
		"server.globalresources.desc":       "전역 JNDI 리소스",
		"server.listener.add":               "리스너 추가",
		"server.listener.edit":              "리스너 편집",
		"server.listener.classname":         "클래스 이름",
		"server.service":                    "서비스",
		"server.service.name":               "서비스 이름",
		"server.engine":                     "엔진",
		"server.engine.name":                "엔진 이름",
		"server.engine.defaulthost":         "기본 호스트",
		"server.engine.jvmroute":            "JVM Route",
		"server.executor":                   "Executor",
		"server.executor.add":               "Executor 추가",
		"server.executor.edit":              "Executor 편집",
		"server.executor.name":              "Executor 이름",
		"server.executor.prefix":            "이름 접두사",
		"server.executor.maxthreads":        "최대 스레드",
		"server.executor.minthreads":        "최소 여유 스레드",
		"server.executor.maxidle":           "최대 유휴 시간 (ms)",
		"server.executor.threads":           "스레드: %d-%d, 최대 유휴: %dms",
		"server.executor.updated":           "Executor가 업데이트됨",
		"server.executor.deleted":           "Executor가 삭제됨",
		"server.executor.added":             "Executor가 추가됨",
		"server.listener.delete":            "삭제",
		"server.listener.custom":            "커스텀",
		"server.listener.custom.desc":       "커스텀 리스너 클래스 입력",
		"server.listener.custom.title":      "커스텀 리스너",
		"server.listener.sslengine":         "SSL 엔진",
		"server.listener.sslseed":           "SSL 랜덤 시드",
		"server.listener.updated":           "리스너가 업데이트됨",
		"server.listener.deleted":           "리스너가 삭제됨",
		"server.listener.added":             "리스너가 추가됨",
		"server.listener.detail":            "리스너 상세",
		"server.listener.classrequired":     "클래스 이름이 필요합니다",
		"server.service.edit":               "서비스 이름 편집",
		"server.service.engine.desc":        "기본 호스트: %s, jvmRoute: %s",
		"server.service.connectors":         "커넥터",
		"server.service.connectors.desc":    "HTTP/AJP 커넥터 (커넥터 메뉴에서 설정)",
		"server.service.hosts":              "호스트",
		"server.service.hosts.desc":         "가상 호스트 (호스트 메뉴에서 설정)",
		"server.service.updated":            "서비스 이름이 업데이트됨",
		"server.engine.settings":            "엔진 설정",
		"server.engine.saved":               "엔진 설정이 저장됨",
		"server.globalresource.add":         "리소스 추가",
		"server.globalresource.add.desc":    "새 전역 리소스 추가",
		"server.globalresource.edit":        "리소스 편집",
		"server.globalresource.auth":        "인증",
		"server.globalresource.type":        "유형",
		"server.globalresource.description": "설명",
		"server.globalresource.factory":     "팩토리",
		"server.globalresource.pathname":    "경로명",
		"server.globalresource.updated":     "리소스가 업데이트됨",
		"server.globalresource.deleted":     "리소스가 삭제됨",
		"server.globalresource.added":       "리소스가 추가됨",
		"server.settings":                   "서버 설정",
		"server.settings.saved":             "서버 설정이 저장됨",
		"server.settings.invalidport":       "잘못된 포트 번호",
		"server.confirm.delete":             "이 %s을(를) 삭제하시겠습니까?",
		"server.confirm.yes":                "예",
		"server.confirm.no":                 "아니오",

		// Connector View
		"connector.title":                   "커넥터 설정",
		"connector.http":                    "HTTP 커넥터",
		"connector.http.desc":               "HTTP/1.1, HTTP/2 커넥터 설정",
		"connector.ajp":                     "AJP 커넥터",
		"connector.ajp.desc":                "Apache JServ Protocol 커넥터 설정",
		"connector.ssl":                     "SSL/TLS 커넥터",
		"connector.ssl.desc":                "HTTPS 및 SSL 인증서 설정",
		"connector.executor":                "Executor",
		"connector.executor.desc":           "공유 스레드 풀 설정",
		"connector.list":                    "커넥터",
		"connector.add":                     "커넥터 추가",
		"connector.edit":                    "커넥터 편집",
		"connector.port":                    "포트",
		"connector.protocol":                "프로토콜",
		"connector.timeout":                 "연결 타임아웃",
		"connector.redirect":                "리다이렉트 포트",
		"connector.maxthreads":              "최대 스레드",
		"connector.minthreads":              "최소 여유 스레드",
		"connector.acceptcount":             "Accept Count",
		"connector.ssl.enabled":             "SSL 활성화",
		"connector.ssl.keystore":            "키스토어 파일",
		"connector.ssl.password":            "키스토어 비밀번호",
		"connector.ssl.type":                "키스토어 유형",
		"connector.ssl.protocol":            "SSL 프로토콜",
		"connector.ssl.clientauth":          "클라이언트 인증",
		"connector.http.add":                "HTTP 커넥터 추가",
		"connector.http.add.desc":           "새 HTTP 커넥터 생성",
		"connector.ajp.add":                 "AJP 커넥터 추가",
		"connector.ajp.add.desc":            "새 AJP 커넥터 생성",
		"connector.ssl.add":                 "HTTPS 커넥터 추가",
		"connector.ssl.add.desc":            "새 SSL/TLS 커넥터 생성",
		"connector.executor.add":            "Executor 추가",
		"connector.executor.add.desc":       "새 스레드 풀 생성",
		"connector.executor.title":          "Executor (스레드 풀)",
		"connector.service":                 "서비스",
		"connector.secret":                  "Secret",
		"connector.secret.required":         "Secret 필요",
		"connector.secret.none":             "Secret 없음",
		"connector.secret.set":              "Secret 설정됨",
		"connector.keystore.notconfig":      "키스토어가 설정되지 않음",
		"connector.noservices":              "설정된 서비스가 없습니다",
		"connector.updated.http":            "HTTP 커넥터가 업데이트됨",
		"connector.updated.ajp":             "AJP 커넥터가 업데이트됨",
		"connector.updated.ssl":             "SSL 커넥터가 업데이트됨",
		"connector.deleted":                 "커넥터가 삭제됨",
		"connector.added":                   "커넥터가 추가됨",
		"connector.ssl.added":               "SSL 커넥터가 추가됨",
		"connector.executor.updated":        "Executor가 업데이트됨",
		"connector.executor.deleted":        "Executor가 삭제됨",
		"connector.executor.added":          "Executor가 추가됨",
		"connector.delete.title":            "커넥터 삭제",
		"connector.delete.confirm":          "포트 %d의 커넥터를 삭제하시겠습니까?",
		"connector.delete.ajp.confirm":      "포트 %d의 AJP 커넥터를 삭제하시겠습니까?",
		"connector.delete.ssl.confirm":      "포트 %d의 SSL 커넥터를 삭제하시겠습니까?",
		"connector.executor.delete.title":   "Executor 삭제",
		"connector.executor.delete.confirm": "Executor '%s'을(를) 삭제하시겠습니까?",
		"connector.executor.name":           "이름",
		"connector.executor.nameprefix":     "이름 접두사",
		"connector.executor.maxidle":        "최대 유휴 시간 (ms)",
		"connector.executor.optional":       "Executor (선택사항)",
		"connector.executor.add.title":      "Executor 추가",
		"connector.returnmenu":              "커넥터 메뉴로 돌아가기",
		"connector.secretrequired":          "Secret 필요",
		"connector.sslprotocol":             "SSL 프로토콜",
		"connector.keystorefile":            "키스토어 파일",
		"connector.keystorepass":            "키스토어 비밀번호",
		"connector.keystoretype":            "키스토어 유형",
		"connector.clientauth":              "클라이언트 인증",
		"connector.http.add.title":          "HTTP 커넥터 추가",
		"connector.ajp.add.title":           "AJP 커넥터 추가",
		"connector.ssl.add.title":           "HTTPS 커넥터 추가",
		"connector.error.noservices":        "설정된 서비스가 없습니다",
		"connector.added.ssl":               "SSL 커넥터가 추가됨",

		// Security View
		"security.title":                "보안 및 인증",
		"security.realm":                "Realm 설정",
		"security.realm.desc":           "인증 Realm 설정",
		"security.realm.add":            "Realm 추가",
		"security.realm.edit":           "Realm 편집",
		"security.realm.type":           "Realm 유형",
		"security.realm.current":        "현재",
		"security.realm.nested":         "중첩 Realm",
		"security.realm.set":            "Realm 유형 설정",
		"security.realm.set.desc":       "다른 Realm 유형 설정",
		"security.realm.remove":         "Realm 제거",
		"security.realm.remove.desc":    "현재 Realm 설정 제거",
		"security.realm.remove.confirm": "현재 Realm 설정을 제거하시겠습니까?",
		"security.realm.removed":        "Realm이 제거되었습니다",
		"security.realm.config":         "Realm 설정",
		"security.realm.selecttype":     "Realm 유형 선택",
		"security.users":                "사용자 및 역할",
		"security.users.desc":           "tomcat-users.xml 관리",
		"security.users.title":          "사용자 및 역할 (tomcat-users.xml)",
		"security.users.list":           "사용자",
		"security.users.list.desc":      "사용자 계정 관리",
		"security.credential":           "자격 증명 핸들러",
		"security.credential.desc":      "비밀번호 해싱 설정",
		"security.user.add":             "사용자 추가",
		"security.user.edit":            "사용자 편집",
		"security.user.name":            "사용자명",
		"security.user.password":        "비밀번호",
		"security.user.roles":           "역할",
		"security.roles":                "역할",
		"security.roles.list":           "역할",
		"security.roles.list.desc":      "역할 정의 관리",
		"security.role.add":             "역할 추가",
		"security.role.name":            "역할 이름",

		// JNDI View
		"jndi.title":                "JNDI 리소스 - context.xml",
		"jndi.resources":            "리소스",
		"jndi.resource.add":         "리소스 추가",
		"jndi.resource.edit":        "리소스 편집",
		"jndi.resource.name":        "리소스 이름",
		"jndi.resource.type":        "리소스 유형",
		"jndi.resource.auth":        "인증",
		"jndi.datasource":           "DataSource (JDBC)",
		"jndi.datasource.desc":      "데이터베이스 커넥션 풀",
		"jndi.datasource.driver":    "드라이버 클래스",
		"jndi.datasource.url":       "JDBC URL",
		"jndi.datasource.username":  "사용자명",
		"jndi.datasource.password":  "비밀번호",
		"jndi.datasource.maxactive": "최대 활성",
		"jndi.datasource.maxidle":   "최대 유휴",
		"jndi.mail":                 "Mail Session",
		"jndi.mail.desc":            "JavaMail 설정",
		"jndi.environment":          "환경 변수",
		"jndi.environment.desc":     "환경 변수 설정",
		"jndi.environment.add":      "환경 변수 추가",
		"jndi.environment.name":     "항목 이름",
		"jndi.environment.value":    "값",
		"jndi.environment.type":     "유형",
		"jndi.resourcelink":         "리소스 링크",
		"jndi.resourcelink.desc":    "전역 리소스 링크",

		// Host View
		"host.title":                  "가상 호스트 및 Context",
		"host.list":                   "호스트",
		"host.add":                    "호스트 추가",
		"host.edit":                   "호스트 편집",
		"host.name":                   "호스트 이름",
		"host.appbase":                "앱 베이스",
		"host.unpackwars":             "WAR 압축 해제",
		"host.autodeploy":             "자동 배포",
		"host.aliases":                "별칭",
		"host.alias.add":              "별칭 추가",
		"host.virtualhost":            "가상 호스트",
		"host.virtualhost.desc":       "가상 호스트 관리",
		"host.context":                "Context",
		"host.context.desc":           "웹 애플리케이션 Context 관리",
		"host.engine":                 "엔진 설정",
		"host.engine.desc":            "Catalina 엔진 설정",
		"context.title":               "Context 설정 (context.xml)",
		"context.list":                "Context 목록",
		"context.add":                 "Context 추가",
		"context.edit":                "Context 편집",
		"context.path":                "Context 경로",
		"context.docbase":             "문서 베이스",
		"context.reloadable":          "리로드 가능",
		"context.settings":            "Context 설정",
		"context.settings.desc":       "기본 Context 속성 및 옵션",
		"context.resources":           "JNDI 리소스",
		"context.resources.count":     "%d개 리소스 (DataSource, MailSession 등)",
		"context.environment":         "환경 변수",
		"context.environment.count":   "%d개 환경 변수",
		"context.resourcelinks":       "리소스 링크",
		"context.resourcelinks.count": "%d개 리소스 링크",
		"context.parameters":          "파라미터",
		"context.parameters.count":    "%d개 Context 파라미터",
		"context.watched":             "감시 리소스",
		"context.watched.count":       "%d개 감시 리소스",
		"context.manager":             "세션 매니저",
		"context.cookie":              "쿠키 프로세서",
		"context.cookie.desc":         "SameSite, 쿠키 설정",
		"context.jarscanner":          "JAR 스캐너",
		"context.jarscanner.desc":     "클래스 스캐닝 설정",
		"context.save.desc":           "context.xml에 변경사항 저장",
		"host.sessionmanager":         "세션 매니저",

		// Valve View
		"valve.title":             "밸브 설정",
		"valve.list":              "밸브",
		"valve.add":               "밸브 추가",
		"valve.edit":              "밸브 편집",
		"valve.type":              "밸브 유형",
		"valve.engine":            "엔진 밸브",
		"valve.engine.desc":       "모든 요청에 적용되는 밸브",
		"valve.host":              "호스트 밸브",
		"valve.host.desc":         "특정 가상 호스트용 밸브",
		"valve.context":           "Context 밸브",
		"valve.context.desc":      "특정 애플리케이션용 밸브",
		"valve.quickadd":          "자주 사용하는 밸브 빠른 추가",
		"valve.quickadd.desc":     "자주 사용하는 밸브 추가",
		"valve.accesslog":         "접근 로그 밸브",
		"valve.accesslog.dir":     "디렉토리",
		"valve.accesslog.prefix":  "접두사",
		"valve.accesslog.suffix":  "접미사",
		"valve.accesslog.pattern": "패턴",
		"valve.remoteaddr":        "원격 주소 밸브",
		"valve.remoteaddr.allow":  "허용",
		"valve.remoteaddr.deny":   "거부",
		"valve.remoteip":          "원격 IP 밸브",
		"valve.remoteip.header":   "원격 IP 헤더",
		"valve.remoteip.protocol": "프로토콜 헤더",
		"valve.sso":               "싱글 사인온",
		"valve.error":             "오류 보고 밸브",

		// Cluster View
		"cluster.title":              "클러스터링 설정",
		"cluster.enable":             "클러스터링 활성화",
		"cluster.disable":            "클러스터링 비활성화",
		"cluster.status":             "클러스터 상태",
		"cluster.status.desc":        "클러스터링 활성화/비활성화",
		"cluster.settings":           "클러스터 설정",
		"cluster.settings.desc":      "기본 클러스터 설정",
		"cluster.manager":            "세션 매니저",
		"cluster.manager.desc":       "DeltaManager 또는 BackupManager",
		"cluster.manager.type":       "매니저 유형",
		"cluster.manager.delta":      "DeltaManager",
		"cluster.manager.backup":     "BackupManager",
		"cluster.channel":            "채널",
		"cluster.membership":         "멤버십",
		"cluster.membership.desc":    "멀티캐스트 멤버십 설정",
		"cluster.membership.address": "멀티캐스트 주소",
		"cluster.membership.port":    "멀티캐스트 포트",
		"cluster.receiver":           "수신기",
		"cluster.receiver.desc":      "메시지 수신기 설정",
		"cluster.receiver.address":   "주소",
		"cluster.receiver.port":      "포트",
		"cluster.sender":             "송신기",
		"cluster.sender.desc":        "메시지 송신기 설정",
		"cluster.interceptors":       "인터셉터",
		"cluster.interceptors.desc":  "채널 인터셉터",
		"cluster.interceptor.add":    "인터셉터 추가",
		"cluster.deployer":           "팜 디플로이어",
		"cluster.deployer.desc":      "클러스터 배포 설정",
		"cluster.deployer.remove":    "디플로이어 제거",

		// Logging View
		"logging.title":              "로깅 설정 (logging.properties)",
		"logging.handlers":           "핸들러",
		"logging.handler.add":        "핸들러 추가",
		"logging.handler.edit":       "핸들러 편집",
		"logging.handler.type":       "핸들러 유형",
		"logging.handler.level":      "레벨",
		"logging.handler.directory":  "디렉토리",
		"logging.handler.prefix":     "접두사",
		"logging.filehandlers":       "파일 핸들러",
		"logging.filehandlers.count": "%d개 파일 핸들러 설정됨",
		"logging.console":            "콘솔 핸들러",
		"logging.loggers":            "로거",
		"logging.loggers.count":      "%d개 로거 설정됨",
		"logging.logger.add":         "로거 추가",
		"logging.logger.edit":        "로거 편집",
		"logging.logger.name":        "로거 이름",
		"logging.logger.level":       "레벨",
		"logging.logger.handlers":    "핸들러",
		"logging.rootlogger":         "루트 로거",
		"logging.rootlogger.count":   "%d개 핸들러 할당됨",
		"logging.save.desc":          "logging.properties에 변경사항 저장",

		// Context View (context.xml)
		"ctxxml.title":            "Context 설정",
		"ctxxml.settings":         "기본 설정",
		"ctxxml.reloadable":       "리로드 가능",
		"ctxxml.crosscontext":     "Cross Context",
		"ctxxml.privileged":       "Privileged",
		"ctxxml.cookies":          "쿠키 설정",
		"ctxxml.cookies.httponly": "HTTP Only",
		"ctxxml.cookies.name":     "세션 쿠키 이름",
		"ctxxml.resources":        "리소스",
		"ctxxml.parameters":       "파라미터",
		"ctxxml.watched":          "감시 리소스",
		"ctxxml.manager":          "세션 매니저",
		"ctxxml.loader":           "클래스 로더",
		"ctxxml.jarscanner":       "JAR 스캐너",

		// Web View (web.xml)
		"web.title":              "웹 애플리케이션 설정 (web.xml)",
		"web.servlets":           "서블릿",
		"web.servlets.count":     "%d개 서블릿 설정됨",
		"web.filters":            "필터",
		"web.filters.count":      "%d개 필터 설정됨",
		"web.listeners":          "리스너",
		"web.listeners.count":    "%d개 리스너 설정됨",
		"web.session":            "세션 설정",
		"web.welcomefiles":       "환영 파일",
		"web.welcomefiles.desc":  "기본 페이지 파일",
		"web.errorpages":         "오류 페이지",
		"web.errorpages.count":   "%d개 오류 페이지",
		"web.mime":               "MIME 매핑",
		"web.mime.count":         "%d개 MIME 유형",
		"web.security":           "보안 제약",
		"web.security.count":     "%d개 제약",
		"web.login":              "로그인 설정",
		"web.login.desc":         "인증 방법",
		"web.roles":              "보안 역할",
		"web.roles.desc":         "보안 역할 정의",
		"web.contextparams":      "Context 파라미터",
		"web.contextparams.desc": "웹 앱 초기화 파라미터",
		"web.save.desc":          "web.xml에 변경사항 저장",
		// Web sub-menu keys
		"web.servlet.add":                     "서블릿 추가",
		"web.servlet.add.desc":                "새 서블릿 생성",
		"web.servlet.edit":                    "서블릿 편집",
		"web.servlet.name":                    "서블릿 이름",
		"web.servlet.class":                   "서블릿 클래스",
		"web.servlet.jsp":                     "JSP 파일 (선택사항)",
		"web.servlet.loadonstartup":           "시작 시 로드",
		"web.servlet.async":                   "비동기 지원",
		"web.servlet.initparams":              "초기화 파라미터 (줄당 name=value)",
		"web.servlet.urlpatterns":             "URL 패턴 (줄당 하나)",
		"web.servlet.quickdefault":            "기본 서블릿 빠른 추가",
		"web.servlet.quickdefault.desc":       "Tomcat 기본 서블릿 추가",
		"web.servlet.quickjsp":                "JSP 서블릿 빠른 추가",
		"web.servlet.quickjsp.desc":           "Tomcat JSP 서블릿 추가",
		"web.servlet.added":                   "서블릿이 추가됨",
		"web.servlet.updated":                 "서블릿이 업데이트됨",
		"web.servlet.deleted":                 "서블릿이 삭제됨",
		"web.servlet.error.name":              "서블릿 이름이 필요합니다",
		"web.servlet.error.class":             "서블릿 클래스 또는 JSP 파일이 필요합니다",
		"web.filter.add":                      "필터 추가",
		"web.filter.add.desc":                 "새 필터 생성",
		"web.filter.edit":                     "필터 편집",
		"web.filter.name":                     "필터 이름",
		"web.filter.class":                    "필터 클래스",
		"web.filter.async":                    "비동기 지원",
		"web.filter.initparams":               "초기화 파라미터 (줄당 name=value)",
		"web.filter.urlpatterns":              "URL 패턴 (줄당 하나)",
		"web.filter.quickcors":                "CORS 필터 빠른 추가",
		"web.filter.quickcors.desc":           "CORS 필터 추가",
		"web.filter.quickencoding":            "인코딩 필터 빠른 추가",
		"web.filter.quickencoding.desc":       "문자 인코딩 필터 추가",
		"web.filter.added":                    "필터가 추가됨",
		"web.filter.updated":                  "필터가 업데이트됨",
		"web.filter.deleted":                  "필터가 삭제됨",
		"web.filter.error.required":           "필터 이름과 클래스가 필요합니다",
		"web.listener.add":                    "리스너 추가",
		"web.listener.add.desc":               "새 리스너 생성",
		"web.listener.edit":                   "리스너 편집",
		"web.listener.class":                  "리스너 클래스",
		"web.listener.description":            "설명",
		"web.listener.added":                  "리스너가 추가됨",
		"web.listener.deleted":                "리스너가 삭제됨",
		"web.listener.error.class":            "리스너 클래스가 필요합니다",
		"web.session.title":                   "세션 설정",
		"web.session.timeout":                 "세션 타임아웃 (분)",
		"web.session.tracking.cookie":         "추적: COOKIE",
		"web.session.tracking.url":            "추적: URL",
		"web.session.tracking.ssl":            "추적: SSL",
		"web.session.cookie.name":             "쿠키 이름",
		"web.session.cookie.domain":           "쿠키 도메인",
		"web.session.cookie.path":             "쿠키 경로",
		"web.session.cookie.httponly":         "HttpOnly 쿠키",
		"web.session.cookie.secure":           "Secure 쿠키",
		"web.session.saved":                   "세션 설정이 저장됨",
		"web.welcomefiles.title":              "환영 파일",
		"web.welcomefiles.perline":            "환영 파일 (줄당 하나)",
		"web.welcomefiles.adddefaults":        "기본값 추가",
		"web.welcomefiles.saved":              "환영 파일이 저장됨",
		"web.welcomefiles.defaultadded":       "기본 환영 파일이 추가됨",
		"web.errorpage.add":                   "오류 페이지 추가",
		"web.errorpage.add.desc":              "새 오류 페이지 생성",
		"web.errorpage.edit":                  "오류 페이지 편집",
		"web.errorpage.type":                  "오류 유형",
		"web.errorpage.code":                  "오류 코드 (예: 404)",
		"web.errorpage.exception":             "예외 유형",
		"web.errorpage.location":              "위치 (페이지 경로)",
		"web.errorpage.quickcommon":           "일반 오류 페이지 추가",
		"web.errorpage.quickcommon.desc":      "404, 500 오류 페이지 추가",
		"web.errorpage.added":                 "오류 페이지가 추가됨",
		"web.errorpage.deleted":               "오류 페이지가 삭제됨",
		"web.errorpage.commonadded":           "일반 오류 페이지가 추가됨",
		"web.errorpage.error.location":        "위치가 필요합니다",
		"web.mime.add":                        "MIME 매핑 추가",
		"web.mime.add.desc":                   "새 MIME 매핑 생성",
		"web.mime.edit":                       "MIME 매핑 편집",
		"web.mime.extension":                  "확장자 (점 제외)",
		"web.mime.type":                       "MIME 유형",
		"web.mime.quickcommon":                "일반 MIME 유형 추가",
		"web.mime.quickcommon.desc":           "일반 웹 MIME 유형 추가",
		"web.mime.added":                      "MIME 매핑이 추가됨",
		"web.mime.deleted":                    "MIME 매핑이 삭제됨",
		"web.mime.commonadded":                "일반 MIME 유형이 추가됨",
		"web.mime.error.required":             "확장자와 MIME 유형이 필요합니다",
		"web.securityconstraint.add":          "보안 제약 추가",
		"web.securityconstraint.add.desc":     "새 보안 제약 생성",
		"web.securityconstraint.edit":         "보안 제약 편집",
		"web.securityconstraint.resourcename": "리소스 이름",
		"web.securityconstraint.urlpatterns":  "URL 패턴 (줄당 하나)",
		"web.securityconstraint.httpmethods":  "HTTP 메서드 (쉼표로 구분, 비우면 전체)",
		"web.securityconstraint.roles":        "필요 역할 (줄당 하나)",
		"web.securityconstraint.transport":    "전송 보장",
		"web.securityconstraint.added":        "보안 제약이 추가됨",
		"web.securityconstraint.updated":      "보안 제약이 업데이트됨",
		"web.securityconstraint.deleted":      "보안 제약이 삭제됨",
		"web.login.title":                     "로그인 설정",
		"web.login.authmethod":                "인증 방법",
		"web.login.realmname":                 "Realm 이름",
		"web.login.formloginpage":             "폼 로그인 페이지",
		"web.login.formerrorpage":             "폼 오류 페이지",
		"web.login.saved":                     "로그인 설정이 저장됨",
		"web.login.removed":                   "로그인 설정이 제거됨",
		"web.role.add":                        "보안 역할 추가",
		"web.role.add.desc":                   "새 보안 역할 생성",
		"web.role.edit":                       "보안 역할 편집",
		"web.role.name":                       "역할 이름",
		"web.role.description":                "설명",
		"web.role.quickcommon":                "일반 역할 추가",
		"web.role.quickcommon.desc":           "admin, user, manager 역할 추가",
		"web.role.added":                      "보안 역할이 추가됨",
		"web.role.deleted":                    "보안 역할이 삭제됨",
		"web.role.commonadded":                "일반 역할이 추가됨",
		"web.role.error.name":                 "역할 이름이 필요합니다",
		"web.contextparam.add":                "Context 파라미터 추가",
		"web.contextparam.add.desc":           "새 파라미터 생성",
		"web.contextparam.edit":               "Context 파라미터 편집",
		"web.contextparam.name":               "파라미터 이름",
		"web.contextparam.value":              "파라미터 값",
		"web.contextparam.description":        "설명",
		"web.contextparam.added":              "Context 파라미터가 추가됨",
		"web.contextparam.updated":            "Context 파라미터가 업데이트됨",
		"web.contextparam.deleted":            "Context 파라미터가 삭제됨",
		"web.contextparam.error.name":         "파라미터 이름이 필요합니다",
		"webxml.title":                        "웹 애플리케이션 설정",
		"webxml.servlets":                     "서블릿",
		"webxml.servlet.add":                  "서블릿 추가",
		"webxml.servlet.edit":                 "서블릿 편집",
		"webxml.servlet.name":                 "서블릿 이름",
		"webxml.servlet.class":                "서블릿 클래스",
		"webxml.servlet.mapping":              "URL 패턴",
		"webxml.filters":                      "필터",
		"webxml.filter.add":                   "필터 추가",
		"webxml.filter.edit":                  "필터 편집",
		"webxml.filter.name":                  "필터 이름",
		"webxml.filter.class":                 "필터 클래스",
		"webxml.listeners":                    "리스너",
		"webxml.listener.add":                 "리스너 추가",
		"webxml.listener.class":               "리스너 클래스",
		"webxml.session":                      "세션 설정",
		"webxml.session.timeout":              "세션 타임아웃 (분)",
		"webxml.welcome":                      "환영 파일",
		"webxml.error":                        "오류 페이지",
		"webxml.mime":                         "MIME 매핑",
		"webxml.security":                     "보안 제약",
		"webxml.security.add":                 "보안 제약 추가",
		"webxml.login":                        "로그인 설정",
		"webxml.login.method":                 "인증 방법",
		"webxml.roles":                        "보안 역할",

		// ==================== 상세 도움말 ====================
		// 서버 설정 도움말
		"help.server.port": `[::b]Shutdown 포트[::-]
Tomcat이 종료 명령을 수신하는 TCP/IP 포트 번호입니다.

[yellow]기본값:[-] 8005
[yellow]범위:[-] 1-65535 (또는 -1로 비활성화)

[green]보안 참고사항:[-]
• 운영 환경에서는 -1로 설정하여 원격 종료 비활성화
• localhost(127.0.0.1)에서만 수신
• 종료 스크립트의 포트와 일치해야 함

[gray]예: <Server port="8005" shutdown="SHUTDOWN">[-]`,

		"help.server.shutdown": `[::b]Shutdown 명령[::-]
종료를 트리거하기 위해 수신해야 하는 명령 문자열입니다.

[yellow]기본값:[-] SHUTDOWN

[green]보안 권장사항:[-]
• 기본 "SHUTDOWN"을 복잡하고 무작위한 문자열로 변경
• port=-1과 함께 사용하면 최대 보안 제공
• catalina.sh stop / shutdown.bat에서 사용

[gray]예: <Server port="8005" shutdown="복잡한_비밀_문자열">[-]`,

		"help.server.listener": `[::b]라이프사이클 리스너[::-]
리스너는 서버 라이프사이클의 특정 이벤트에 응답합니다.

[green]일반적인 리스너:[-]
• [yellow]VersionLoggerListener[-]: 시작 시 Tomcat 버전 정보 로깅
• [yellow]AprLifecycleListener[-]: Apache Portable Runtime (APR) 활성화
• [yellow]JreMemoryLeakPreventionListener[-]: JRE 메모리 누수 방지
• [yellow]GlobalResourcesLifecycleListener[-]: JNDI 리소스에 필요
• [yellow]ThreadLocalLeakPreventionListener[-]: 스레드 로컬 변수 정리

[gray]모든 리스너는 선택 사항이지만 운영 환경에서 권장됩니다.[-]`,

		"help.server.service": `[::b]서비스[::-]
서비스는 하나 이상의 커넥터를 단일 엔진과 그룹화합니다.

[yellow]기본 이름:[-] Catalina

[green]구성 요소:[-]
• [yellow]Engine[-]: 요청 처리 엔진 (서비스당 하나)
• [yellow]Connectors[-]: HTTP, HTTPS, AJP (하나 이상)
• [yellow]Executors[-]: 공유 스레드 풀 (선택사항)

[gray]대부분의 설치는 "Catalina"라는 단일 서비스를 사용합니다.[-]`,

		"help.server.engine": `[::b]엔진[::-]
엔진은 커넥터에서 모든 요청을 수신하고 처리합니다.

[green]속성:[-]
• [yellow]name[-]: 논리적 이름 (기본값: Catalina)
• [yellow]defaultHost[-]: 일치하지 않는 요청의 기본 호스트
• [yellow]jvmRoute[-]: 로드 밸런서 스티키 세션용 고유 ID

[green]jvmRoute 사용:[-]
세션 ID 형식: <session-id>.<jvmRoute>
예: ABC123.node1

[gray]클러스터 환경에서 세션 어피니티에 필요합니다.[-]`,

		"help.server.executor": `[::b]Executor (스레드 풀)[::-]
커넥터를 위한 공유 스레드 풀입니다.

[green]속성:[-]
• [yellow]name[-]: 고유 식별자 (커넥터에서 참조)
• [yellow]maxThreads[-]: 최대 작업자 스레드 (기본값: 200)
• [yellow]minSpareThreads[-]: 최소 유휴 스레드 (기본값: 25)
• [yellow]maxIdleTime[-]: 유휴 스레드 타임아웃 ms (기본값: 60000)

[green]크기 조정 가이드라인:[-]
• 개발: 25-100 스레드
• 운영: 150-400 스레드
• 고트래픽: 400-800 스레드

[gray]커넥터는 executor="name" 속성으로 Executor를 참조합니다.[-]`,

		// 커넥터 도움말
		"help.connector.http": `[::b]HTTP 커넥터[::-]
HTTP/1.1 및 HTTP/2 클라이언트 연결을 처리합니다.

[green]주요 속성:[-]
• [yellow]port[-]: TCP 포트 (기본값: 8080)
• [yellow]protocol[-]: HTTP/1.1, org.apache.coyote.http11.Http11NioProtocol
• [yellow]connectionTimeout[-]: 소켓 타임아웃 ms (기본값: 20000)
• [yellow]maxThreads[-]: 최대 요청 스레드 (기본값: 200)
• [yellow]acceptCount[-]: 백로그 큐 크기 (기본값: 100)

[green]프로토콜 옵션:[-]
• HTTP/1.1 - 최적 구현 자동 선택
• Http11NioProtocol - 논블로킹 I/O (권장)
• Http11Nio2Protocol - NIO2/AIO 구현
• Http11AprProtocol - APR/native (APR 라이브러리 필요)`,

		"help.connector.https": `[::b]HTTPS 커넥터 (SSL/TLS)[::-]
SSL/TLS 암호화로 보안 HTTP 연결을 제공합니다.

[green]주요 속성:[-]
• [yellow]port[-]: TCP 포트 (기본값: 8443)
• [yellow]SSLEnabled[-]: "true"여야 함
• [yellow]keystoreFile[-]: 키스토어 경로 (.jks, .p12)
• [yellow]keystorePass[-]: 키스토어 비밀번호
• [yellow]keystoreType[-]: JKS, PKCS12 (권장)

[green]SSL 프로토콜 옵션:[-]
• TLS (최고 버전 자동 협상)
• TLSv1.2 (최소 권장)
• TLSv1.3 (가장 안전, Java 11+)

[green]보안 권장사항:[-]
• TLSv1.2 또는 TLSv1.3만 사용
• 약한 암호 비활성화
• 2048비트+ RSA 또는 256비트+ ECC 키 사용`,

		"help.connector.ajp": `[::b]AJP 커넥터[::-]
Apache httpd 연동을 위한 Apache JServ Protocol입니다.

[green]주요 속성:[-]
• [yellow]port[-]: TCP 포트 (기본값: 8009)
• [yellow]protocol[-]: AJP/1.3
• [yellow]secretRequired[-]: 공유 시크릿 필요 (9.0.31+에서 기본값: true)
• [yellow]secret[-]: 공유 인증 시크릿
• [yellow]address[-]: 바인드 주소 (기본값: 모든 인터페이스)

[green]보안 (Tomcat 9.0.31+):[-]
• secretRequired="true"는 인증을 강제
• secret은 Apache httpd의 ProxyPass secret과 일치해야 함
• address="127.0.0.1"은 localhost로 제한

[gray]예: ProxyPass /app ajp://localhost:8009/app secret=mySecret[-]`,

		"help.connector.port": `[::b]포트[::-]
수신 연결을 위한 TCP 포트 번호입니다.

[yellow]기본 포트:[-]
• HTTP: 8080
• HTTPS: 8443
• AJP: 8009

[green]참고사항:[-]
• 1024 미만 포트는 root/admin 권한 필요
• 80→8080 리디렉션에 iptables/방화벽 사용
• 각 커넥터는 고유한 포트를 사용해야 함`,

		"help.connector.protocol": `[::b]프로토콜[::-]
프로토콜 핸들러 구현입니다.

[green]HTTP 프로토콜:[-]
• [yellow]HTTP/1.1[-]: 최적 구현 자동 선택
• [yellow]Http11NioProtocol[-]: 논블로킹 I/O (기본값)
• [yellow]Http11Nio2Protocol[-]: Java NIO2 (비동기)
• [yellow]Http11AprProtocol[-]: APR/native (APR 필요)

[green]AJP 프로토콜:[-]
• [yellow]AJP/1.3[-]: 표준 AJP 프로토콜
• [yellow]AjpNioProtocol[-]: NIO 구현
• [yellow]AjpNio2Protocol[-]: NIO2 구현

[gray]대부분의 배포에는 NIO가 권장됩니다.[-]`,

		"help.connector.maxthreads": `[::b]최대 스레드[::-]
요청 처리 스레드의 최대 수입니다.

[yellow]기본값:[-] 200

[green]크기 조정 가이드라인:[-]
• 낮은 트래픽: 25-100
• 중간 트래픽: 150-300
• 높은 트래픽: 400-800

[green]공식:[-]
maxThreads ≈ (피크_동시_사용자 × 평균_응답_시간_초)

[yellow]경고:[-]
• 너무 적음: 요청이 큐에 쌓이고 타임아웃
• 너무 많음: 메모리 고갈, 컨텍스트 스위칭 오버헤드
• 운영 환경에서 스레드 풀 사용량 모니터링`,

		"help.connector.minsparethreads": `[::b]최소 여유 스레드[::-]
유지되는 최소 유휴 스레드 수입니다.

[yellow]기본값:[-] 10

[green]목적:[-]
• 수신 요청을 위해 스레드를 준비 상태로 유지
• 초기 요청의 지연 시간 감소
• 버스트 시 스레드 생성 오버헤드 방지

[green]권장사항:[-]
• 대부분의 애플리케이션에 10-25 설정
• 버스티한 트래픽 패턴에는 더 높은 값
• 메모리 사용량 줄이려면 더 낮은 값`,

		"help.connector.connectiontimeout": `[::b]연결 타임아웃[::-]
연결 후 요청 데이터를 기다리는 밀리초입니다.

[yellow]기본값:[-] 20000 (20초)

[green]동작:[-]
• TCP 연결 설정 후 타이머 시작
• 데이터 수신 시마다 리셋
• 초과 시 연결 종료

[green]권장사항:[-]
• 일반 애플리케이션: 20000-60000
• 로드 밸런서 뒤: 더 낮게 (5000-10000)
• 느린 클라이언트나 대용량 업로드: 더 높게`,

		"help.connector.acceptcount": `[::b]Accept Count[::-]
수신 연결의 최대 큐 길이입니다.

[yellow]기본값:[-] 100

[green]동작:[-]
• 모든 스레드가 바쁠 때 연결이 큐에 대기
• 초과 시 "connection refused" 반환
• OS 자체 하한이 있을 수 있음

[green]튜닝:[-]
• 트래픽 급증에 대비해 증가 (200-500)
• 과부하 시 빠른 실패를 위해 감소
• 운영 환경에서 큐 깊이 모니터링`,

		"help.connector.redirectport": `[::b]리다이렉트 포트[::-]
자동 HTTPS 리디렉션을 위한 포트입니다.

[yellow]기본값:[-] 8443

[green]사용:[-]
• security-constraint가 HTTPS를 요구할 때 사용
• 자동으로 HTTP→HTTPS 리디렉션
• HTTPS 커넥터 포트와 일치해야 함

[gray]예: 8080의 HTTP가 8443의 HTTPS로 리디렉션[-]`,

		"help.connector.executor": `[::b]Executor (스레드 풀)[::-]
공유 스레드 풀에 대한 참조입니다.

[yellow]사용법:[-]
• 비워두면 커넥터 자체 스레드 풀 사용
• Executor 이름을 설정하면 커넥터 간 스레드 공유

[green]장점:[-]
• 중앙화된 스레드 풀 관리
• 더 나은 리소스 활용
• 쉬운 모니터링 및 튜닝

[gray]server.xml에서 Executor 정의:
<Executor name="tomcatThreadPool"
  maxThreads="200" minSpareThreads="10"/>[-]`,

		"help.connector.secretrequired": `[::b]Secret Required[::-]
AJP 연결에 대한 비밀 기반 인증을 활성화합니다.

[yellow]기본값:[-] true (Tomcat 9.0.31+)

[green]목적:[-]
• AJP 포트에 대한 무단 접근 방지
• Ghostcat 취약점 완화 (CVE-2020-1938)
• 양쪽에서 일치하는 비밀 필요

[red]보안:[-]
• 운영 환경에서는 항상 활성화
• 강력하고 무작위한 비밀 값 사용
• AJP 포트를 신뢰할 수 없는 네트워크에 노출하지 않음`,

		"help.connector.secret": `[::b]Secret[::-]
AJP 커넥터 인증을 위한 공유 비밀입니다.

[yellow]요구사항:[-]
• Apache mod_proxy_ajp 비밀과 일치해야 함
• 강력하고 무작위한 값 사용 (32자 이상)
• 기밀 유지

[green]설정:[-]
• Tomcat: secret="yourSecretValue"
• Apache: ProxyPass ajp://host:8009 secret=yourSecretValue

[gray]생성 명령: openssl rand -base64 32[-]`,

		// 보안 도움말
		"help.security.realm": `[::b]Realm[::-]
Realm은 인증을 위해 Tomcat을 사용자/역할 데이터베이스에 연결합니다.

[green]Realm 유형:[-]
• [yellow]UserDatabaseRealm[-]: tomcat-users.xml 사용 (기본값)
• [yellow]JDBCRealm[-]: JDBC를 통한 데이터베이스
• [yellow]DataSourceRealm[-]: JNDI DataSource를 통한 데이터베이스
• [yellow]JNDIRealm[-]: LDAP/Active Directory
• [yellow]JAASRealm[-]: Java 인증 (JAAS)

[green]배치:[-]
• Engine 레벨: 모든 Host/Context에 적용
• Host 레벨: Host의 모든 Context에 적용
• Context 레벨: 단일 애플리케이션에 적용

[gray]CombinedRealm 또는 LockOutRealm으로 중첩 가능합니다.[-]`,

		"help.security.userdatabase": `[::b]UserDatabaseRealm[::-]
tomcat-users.xml을 사용하는 파일 기반 인증입니다.

[green]설정:[-]
• 파일: conf/tomcat-users.xml
• Roles: 액세스 권한 정의
• Users: 사용자명, 비밀번호, 역할 할당

[green]기본 역할:[-]
• manager-gui: Manager 웹 인터페이스 액세스
• manager-script: Manager 텍스트/스크립트 인터페이스 액세스
• admin-gui: Host Manager 인터페이스 액세스

[yellow]보안 참고:[-]
tomcat-users.xml의 비밀번호는 다이제스트되어야 합니다:
1. 실행: digest.sh -a SHA-256 mypassword
2. 출력을 password 속성에 사용`,

		"help.security.jdbcrealm": `[::b]JDBCRealm / DataSourceRealm[::-]
데이터베이스 기반 인증입니다.

[green]필요한 테이블:[-]
• Users 테이블: username, password 컬럼
• Roles 테이블: username, role_name 컬럼

[green]DataSourceRealm 속성:[-]
• dataSourceName: JNDI 이름 (예: jdbc/UserDB)
• userTable: 사용자명을 포함하는 테이블
• userNameCol: 사용자명 컬럼 이름
• userCredCol: 비밀번호 컬럼 이름
• userRoleTable: 역할을 포함하는 테이블
• roleNameCol: 역할 컬럼 이름

[gray]연결 풀링을 위해 JDBCRealm보다 DataSourceRealm이 선호됩니다.[-]`,

		"help.security.ldaprealm": `[::b]JNDIRealm (LDAP/Active Directory)[::-]
LDAP 기반 인증 및 권한 부여입니다.

[green]주요 속성:[-]
• connectionURL: ldap://server:389 또는 ldaps://server:636
• userPattern: DN 패턴, 예: uid={0},ou=users,dc=example,dc=com
• userSearch: 사용자 검색 필터
• roleBase: 역할 검색의 기본 DN
• roleName: 역할 이름을 포함하는 속성

[green]Active Directory 예:[-]
• userPattern: {0}@domain.com
• userSearch: (sAMAccountName={0})
• roleSearch: (member={0})

[gray]대량 인증에는 connectionPoolSize를 사용하세요.[-]`,

		// JNDI 도움말
		"help.jndi.datasource": `[::b]JNDI DataSource[::-]
JNDI 조회를 통해 액세스 가능한 데이터베이스 연결 풀입니다.

[green]주요 속성:[-]
• [yellow]name[-]: JNDI 이름 (예: jdbc/MyDB)
• [yellow]type[-]: javax.sql.DataSource
• [yellow]driverClassName[-]: JDBC 드라이버 클래스
• [yellow]url[-]: JDBC 연결 URL
• [yellow]username/password[-]: 데이터베이스 자격증명

[green]연결 풀 설정:[-]
• maxTotal: 최대 연결 수 (기본값: 8)
• maxIdle: 최대 유휴 연결 수
• minIdle: 최소 유휴 연결 수
• maxWaitMillis: 연결 대기 타임아웃

[green]애플리케이션에서 사용:[-]
Context ctx = new InitialContext();
DataSource ds = (DataSource) ctx.lookup("java:comp/env/jdbc/MyDB");`,

		"help.jndi.environment": `[::b]환경 항목[::-]
설정을 위해 JNDI를 통해 액세스 가능한 간단한 값입니다.

[green]지원 타입:[-]
• java.lang.String
• java.lang.Integer
• java.lang.Boolean
• java.lang.Double

[green]예:[-]
<Environment name="maxItems" value="100" type="java.lang.Integer"/>

[green]애플리케이션에서 사용:[-]
Context ctx = new InitialContext();
Integer maxItems = (Integer) ctx.lookup("java:comp/env/maxItems");

[gray]애플리케이션 설정을 외부화하는데 유용합니다.[-]`,

		// Host 도움말
		"help.host": `[::b]가상 호스트[::-]
Host는 자체 애플리케이션을 가진 가상 호스트를 나타냅니다.

[green]주요 속성:[-]
• [yellow]name[-]: 호스트 이름 (예: www.example.com)
• [yellow]appBase[-]: 애플리케이션 디렉토리 (기본값: webapps)
• [yellow]unpackWARs[-]: WAR 파일 자동 추출 (기본값: true)
• [yellow]autoDeploy[-]: 애플리케이션 핫 배포 (기본값: true)

[green]별칭:[-]
이 Host에 매핑되는 추가 호스트 이름.
예: example.com → www.example.com

[green]디렉토리 구조:[-]
• $CATALINA_BASE/webapps/ - 기본 appBase
• 각 하위 디렉토리 또는 WAR는 애플리케이션`,

		"help.host.appbase": `[::b]애플리케이션 베이스[::-]
이 Host의 웹 애플리케이션을 포함하는 디렉토리입니다.

[yellow]기본값:[-] webapps

[green]경로 유형:[-]
• 상대: $CATALINA_BASE 기준
• 절대: 전체 파일시스템 경로

[green]배포 방법:[-]
• appBase에 WAR 파일 드롭
• 웹 애플리케이션이 있는 디렉토리 생성
• conf/Catalina/[host]/에서 Context 설명자 사용`,

		"help.host.autodeploy": `[::b]자동 배포[::-]
파일 변경 시 애플리케이션을 자동으로 배포합니다.

[yellow]기본값:[-] true

[green]동작:[-]
• 새/수정된 파일에 대해 appBase 모니터링
• 새 애플리케이션 자동 배포
• 수정된 애플리케이션 리로드

[yellow]운영 환경 권장사항:[-]
다음을 위해 "false"로 설정:
• 향상된 보안 (무단 배포 방지)
• 향상된 성능 (파일 감시 없음)
• 예측 가능한 배포 타이밍`,

		// Valve 도움말
		"help.valve.accesslog": `[::b]액세스 로그 Valve[::-]
HTTP 요청을 파일에 로깅합니다.

[green]주요 속성:[-]
• [yellow]directory[-]: 로그 디렉토리 (기본값: logs)
• [yellow]prefix[-]: 로그 파일 접두사 (기본값: localhost_access_log)
• [yellow]suffix[-]: 로그 파일 접미사 (기본값: .txt)
• [yellow]pattern[-]: 로그 포맷 패턴

[green]패턴 변수:[-]
• %h - 원격 호스트명/IP
• %l - 원격 논리적 사용자명 (항상 -)
• %u - 인증된 사용자명
• %t - 공통 로그 포맷의 날짜/시간
• %r - 요청의 첫 줄
• %s - HTTP 상태 코드
• %b - 전송된 바이트 (0이면 -)
• %D - 요청 처리 시간 (ms)
• %T - 요청 처리 시간 (초)

[green]일반적인 패턴:[-]
• common: %h %l %u %t "%r" %s %b
• combined: + Referer와 User-Agent`,

		"help.valve.remoteip": `[::b]Remote IP Valve[::-]
프록시 헤더에서 실제 클라이언트 IP를 추출합니다.

[green]주요 속성:[-]
• [yellow]remoteIpHeader[-]: 클라이언트 IP를 포함하는 헤더 (X-Forwarded-For)
• [yellow]protocolHeader[-]: 프로토콜 헤더 (X-Forwarded-Proto)
• [yellow]internalProxies[-]: 신뢰할 수 있는 프록시 IP의 정규식

[green]사용 사례:[-]
리버스 프록시 (nginx, Apache, HAProxy) 뒤:
• 실제 IP가 X-Forwarded-For 헤더에
• request.getRemoteAddr()는 프록시 IP 반환

[green]설정:[-]
<Valve className="...RemoteIpValve"
       remoteIpHeader="X-Forwarded-For"
       protocolHeader="X-Forwarded-Proto"/>

[gray]프록시 뒤에서 정확한 로깅과 보안에 필수적입니다.[-]`,

		"help.valve.remoteaddr": `[::b]원격 주소 필터[::-]
클라이언트 IP 주소에 따라 액세스를 제한합니다.

[green]속성:[-]
• [yellow]allow[-]: 허용된 IP의 정규식
• [yellow]deny[-]: 거부된 IP의 정규식

[green]평가 순서:[-]
1. 먼저 거부 패턴 확인
2. 허용 패턴 확인
3. 일치 없음: allow가 지정되면 거부, 그렇지 않으면 허용

[green]예:[-]
• localhost만 허용: allow="127\.\d+\.\d+\.\d+"
• 내부 네트워크 허용: allow="192\.168\.\d+\.\d+"
• 특정 IP 차단: deny="10\.0\.0\.1"

[yellow]참고:[-] 프록시 뒤에서는 RemoteIpValve와 함께 사용하세요.`,

		"help.valve.sso": `[::b]싱글 사인온 Valve[::-]
애플리케이션 간 싱글 사인온을 활성화합니다.

[green]동작:[-]
• 사용자가 한 번 로그인
• 같은 Host의 모든 앱에서 세션 공유
• 하나의 앱에서 로그아웃하면 모두에서 로그아웃

[green]설정:[-]
Host 레벨에 배치:
<Valve className="...SingleSignOn"/>

[green]요구사항:[-]
• 모든 앱이 동일한 Realm 사용
• 앱이 동일한 가상 Host에 있어야 함
• 운영 환경에서는 보안 쿠키 사용`,

		// 클러스터 도움말
		"help.cluster": `[::b]Tomcat 클러스터링[::-]
여러 Tomcat 인스턴스 간의 세션 복제입니다.

[green]구성 요소:[-]
• [yellow]Manager[-]: 세션 복제 관리
• [yellow]Channel[-]: 그룹 통신
• [yellow]Membership[-]: 클러스터 멤버 검색
• [yellow]Receiver[-]: 복제된 데이터 수신
• [yellow]Sender[-]: 복제된 데이터 전송

[green]Manager 유형:[-]
• [yellow]DeltaManager[-]: 모든 노드로 복제 (소규모 클러스터)
• [yellow]BackupManager[-]: 하나의 백업으로 복제 (대규모 클러스터)

[green]요구사항:[-]
• web.xml에 <distributable/>
• 모든 세션 속성은 Serializable이어야 함
• 멀티캐스트 네트워크 또는 정적 멤버십`,

		"help.cluster.membership": `[::b]클러스터 멤버십[::-]
클러스터 멤버를 검색하고 추적합니다.

[green]멀티캐스트 멤버십:[-]
• address: 멀티캐스트 그룹 (기본값: 228.0.0.4)
• port: 멀티캐스트 포트 (기본값: 45564)
• frequency: 하트비트 간격 (ms)
• dropTime: 멤버 드롭아웃 타임아웃 (ms)

[green]정적 멤버십:[-]
멀티캐스트가 없는 네트워크용:
• 각 멤버를 명시적으로 설정
• StaticMembershipInterceptor 사용

[yellow]네트워크 요구사항:[-]
• UDP 멀티캐스트 활성화 필요
• 방화벽이 멀티캐스트 트래픽 허용
• 모든 노드가 동일한 멀티캐스트 그룹에`,

		// 로깅 도움말
		"help.logging": `[::b]JULI 로깅[::-]
Tomcat은 기본적으로 Java Util Logging (JULI)을 사용합니다.

[green]설정 파일:[-]
conf/logging.properties

[green]Handler 유형:[-]
• FileHandler: 로테이팅 로그 파일에 쓰기
• ConsoleHandler: 콘솔/stdout에 쓰기

[green]로그 레벨 (오름차순):[-]
• FINEST - 가장 상세함
• FINER
• FINE
• CONFIG
• INFO
• WARNING
• SEVERE - 오류만

[green]일반적인 로거:[-]
• org.apache.catalina - Tomcat 코어
• org.apache.coyote - 커넥터
• org.apache.jasper - JSP 엔진`,

		// Context 도움말
		"help.context": `[::b]Context 설정[::-]
웹 애플리케이션과 그 설정을 정의합니다.

[green]주요 속성:[-]
• [yellow]path[-]: URL 경로 (예: /myapp)
• [yellow]docBase[-]: 애플리케이션 디렉토리 또는 WAR
• [yellow]reloadable[-]: 클래스 변경 시 자동 리로드
• [yellow]crossContext[-]: 앱 간 디스패칭 허용

[green]위치:[-]
• $CATALINA_BASE/conf/context.xml - 전역 기본값
• $CATALINA_BASE/conf/[engine]/[host]/ - 앱별
• META-INF/context.xml - 애플리케이션에 임베디드

[yellow]운영 환경 설정:[-]
• reloadable="false" (성능)
• privileged="false" (보안)`,

		"help.context.reloadable": `[::b]리로드 가능[::-]
클래스 변경 시 자동으로 리로드합니다.

[yellow]기본값:[-] false

[green]동작:[-]
• /WEB-INF/classes 및 /WEB-INF/lib 모니터링
• 변경 감지 시 애플리케이션 리로드
• 개발 중에 유용

[yellow]운영 환경 권장사항:[-]
다음을 위해 "false"로 설정:
• 향상된 성능 (파일 모니터링 없음)
• 안정성 (예상치 못한 리로드 없음)
• 메모리 효율성`,

		// Web.xml 도움말
		"help.webxml.servlet": `[::b]서블릿 설정[::-]
서블릿과 그 매핑을 정의합니다.

[green]요소:[-]
• [yellow]servlet-name[-]: 고유 식별자
• [yellow]servlet-class[-]: 전체 클래스 이름
• [yellow]load-on-startup[-]: 초기화 순서 (선택사항)
• [yellow]init-param[-]: 초기화 매개변수

[green]매핑:[-]
• [yellow]url-pattern[-]: 일치시킬 URL 패턴
• 패턴: /exact, /path/*, *.extension

[green]로드 순서:[-]
• 음수 또는 없음: 첫 요청 시 로드
• 0 또는 양수: 시작 시 로드 (낮을수록 먼저)`,

		"help.webxml.filter": `[::b]필터 설정[::-]
필터는 서블릿 전에 요청을 처리합니다.

[green]일반적인 용도:[-]
• 인증/권한 부여
• 요청/응답 수정
• 로깅 및 감사
• 압축
• 문자 인코딩

[green]필터 체인:[-]
• 정의된 순서대로 필터 실행
• 각 필터가 chain.doFilter() 호출
• 서블릿 전 요청 수정 가능
• 서블릿 후 응답 수정 가능

[green]매핑 유형:[-]
• url-pattern: URL 패턴 일치
• servlet-name: 특정 서블릿에 적용
• dispatcher: REQUEST, FORWARD, INCLUDE, ERROR`,

		"help.webxml.session": `[::b]세션 설정[::-]
HTTP 세션 동작을 구성합니다.

[green]session-timeout:[-]
• 세션 만료 전 시간 (분)
• -1 = 세션이 만료되지 않음
• 기본값: 30분

[green]cookie-config:[-]
• name: 쿠키 이름 (기본값: JSESSIONID)
• http-only: JavaScript 액세스 방지 (권장)
• secure: HTTPS를 통해서만 전송
• max-age: 쿠키 수명 (초)

[green]tracking-mode:[-]
• COOKIE - 쿠키 사용 (권장)
• URL - URL 재작성 (보안 위험)
• SSL - SSL 세션 ID`,

		"help.webxml.security": `[::b]보안 제약[::-]
보호된 리소스와 액세스 규칙을 정의합니다.

[green]구성 요소:[-]
• [yellow]web-resource-collection[-]: 보호할 URL
• [yellow]auth-constraint[-]: 필요한 역할
• [yellow]user-data-constraint[-]: 전송 보장

[green]전송 보장:[-]
• NONE - 암호화 불필요
• INTEGRAL - 데이터 무결성 (HTTPS)
• CONFIDENTIAL - 기밀성 (HTTPS)

[green]예:[-]
<security-constraint>
  <web-resource-collection>
    <url-pattern>/admin/*</url-pattern>
  </web-resource-collection>
  <auth-constraint>
    <role-name>admin</role-name>
  </auth-constraint>
</security-constraint>`,

		// DataSource 속성 도움말
		"help.ds.name": `[yellow]JNDI 이름[white]

이 DataSource를 조회하는 데 사용되는 JNDI 이름입니다.

[aqua]예:[white]
  jdbc/MyDatabase

[aqua]코드 사용법:[white]
  Context ctx = new InitialContext();
  DataSource ds = (DataSource)
    ctx.lookup("java:comp/env/jdbc/MyDatabase");`,

		"help.ds.auth": `[yellow]인증[white]

인증을 관리하는 주체를 지정합니다.

[aqua]Container:[white]
  컨테이너(Tomcat)가 리소스 로그인을
  관리합니다. 자격 증명은 DataSource
  설정에 저장됩니다.

[aqua]Application:[white]
  애플리케이션이 연결 시 자격 증명을
  프로그래밍 방식으로 제공합니다.`,

		"help.ds.factory": `[yellow]Factory 클래스[white]

JNDI 객체 팩토리 클래스입니다.

[aqua]기본값:[white]
  org.apache.tomcat.dbcp.dbcp2.
    BasicDataSourceFactory

[aqua]기타 옵션:[white]
  • org.apache.commons.dbcp2.BasicDataSourceFactory
  • com.zaxxer.hikari.HikariJNDIFactory`,

		"help.ds.driver": `[yellow]JDBC 드라이버 클래스[white]

JDBC 드라이버의 정규화된 클래스 이름입니다.

[aqua]일반 드라이버:[white]
  MySQL 8.x: com.mysql.cj.jdbc.Driver
  PostgreSQL: org.postgresql.Driver
  Oracle: oracle.jdbc.OracleDriver
  SQL Server: com.microsoft.sqlserver.
    jdbc.SQLServerDriver
  MariaDB: org.mariadb.jdbc.Driver`,

		"help.ds.url": `[yellow]JDBC URL[white]

데이터베이스 연결 URL입니다.

[aqua]형식 예:[white]
  MySQL:
    jdbc:mysql://host:3306/dbname
  PostgreSQL:
    jdbc:postgresql://host:5432/dbname
  Oracle:
    jdbc:oracle:thin:@host:1521:SID
  SQL Server:
    jdbc:sqlserver://host:1433;
      databaseName=dbname`,

		"help.ds.username": `[yellow]데이터베이스 사용자명[white]

데이터베이스 인증을 위한 사용자명입니다.

이 사용자는 애플리케이션의 데이터베이스
작업에 적절한 권한이 있어야 합니다.`,

		"help.ds.password": `[yellow]데이터베이스 비밀번호[white]

데이터베이스 인증을 위한 비밀번호입니다.

[red]보안 참고:[white]
프로덕션 환경에서는 암호화된 비밀번호나
외부 비밀 관리 사용을 고려하세요.`,

		"help.ds.initialsize": `[yellow]초기 풀 크기[white]

풀이 초기화될 때 생성되는 연결 수입니다.

[aqua]기본값:[white] 0
[aqua]권장:[white] 5-10

애플리케이션의 기본 연결 요구 사항에
따라 설정하세요.`,

		"help.ds.maxtotal": `[yellow]최대 총 연결 수[white]

풀의 최대 활성 연결 수입니다.

[aqua]기본값:[white] 8
[aqua]권장:[white] 20-100

데이터베이스의 max_connections 설정과
앱 인스턴스 수를 고려하세요.`,

		"help.ds.maxidle": `[yellow]최대 유휴 연결 수[white]

풀에서 유휴 상태로 유지될 수 있는
최대 연결 수입니다.

[aqua]기본값:[white] 8
[aqua]권장:[white] 초기 크기와 동일

높은 값은 연결을 준비 상태로 유지하지만
더 많은 데이터베이스 리소스를 소비합니다.`,

		"help.ds.minidle": `[yellow]최소 유휴 연결 수[white]

풀에서 유휴 상태로 유지할
최소 연결 수입니다.

[aqua]기본값:[white] 0
[aqua]권장:[white] 5-10

갑작스러운 트래픽 증가에 빠른 응답을
보장합니다.`,

		"help.ds.maxwait": `[yellow]최대 대기 시간[white]

풀에서 연결을 가져오기 위해 대기하는
최대 시간(ms)입니다.

[aqua]기본값:[white] -1 (무제한)
[aqua]권장:[white] 10000-30000

스레드가 무한정 차단되는 것을 방지하기
위해 타임아웃을 설정하세요.`,

		"help.ds.validationquery": `[yellow]유효성 검사 쿼리[white]

사용 전 연결을 검증하는 SQL 쿼리입니다.

[aqua]데이터베이스별 예:[white]
  MySQL/MariaDB: SELECT 1
  PostgreSQL: SELECT 1
  Oracle: SELECT 1 FROM DUAL
  SQL Server: SELECT 1
  H2/HSQLDB: SELECT 1`,

		"help.ds.testonborrow": `[yellow]대여 시 테스트[white]

애플리케이션에 연결을 빌려주기 전에
연결을 검증합니다.

[aqua]기본값:[white] false
[aqua]권장:[white] true

유효한 연결을 보장하지만 약간의
오버헤드가 추가됩니다.`,

		"help.ds.testwhileidle": `[yellow]유휴 시 테스트[white]

백그라운드에서 유휴 연결을 주기적으로
검증합니다.

[aqua]기본값:[white] false
[aqua]권장:[white] true

요청 지연에 영향을 주지 않고 끊어진
연결을 사전에 제거합니다.`,

		// Mail Session 속성 도움말
		"help.mail.name": `[yellow]JNDI 이름[white]

이 Mail Session을 조회하는 데 사용되는
JNDI 이름입니다.

[aqua]예:[white]
  mail/Session

[aqua]코드 사용법:[white]
  Session session = (Session)
    ctx.lookup("java:comp/env/mail/Session");`,

		"help.mail.auth": `[yellow]인증[white]

인증을 관리하는 주체를 지정합니다.

[aqua]Container:[white]
  Tomcat이 SMTP 인증을 관리합니다.

[aqua]Application:[white]
  애플리케이션이 자격 증명을 제공합니다.`,

		"help.mail.host": `[yellow]SMTP 호스트[white]

SMTP 서버의 호스트명 또는 IP 주소입니다.

[aqua]예:[white]
  • smtp.gmail.com
  • smtp.office365.com
  • localhost`,

		"help.mail.port": `[yellow]SMTP 포트[white]

SMTP 서버의 포트 번호입니다.

[aqua]일반 포트:[white]
  • 25: 표준 SMTP (암호화 없음)
  • 465: SMTPS (SSL/TLS)
  • 587: Submission (STARTTLS)`,

		"help.mail.user": `[yellow]SMTP 사용자[white]

SMTP 인증을 위한 사용자명입니다.

일반적으로 이메일 주소 또는 계정 이름입니다.`,

		"help.mail.protocol": `[yellow]프로토콜[white]

메일 전송 프로토콜입니다.

[aqua]smtp:[white]
  표준 SMTP, 선택적으로 STARTTLS 사용

[aqua]smtps:[white]
  SSL/TLS를 통한 SMTP (암시적)`,

		"help.mail.smtpauth": `[yellow]SMTP 인증[white]

SMTP 인증을 활성화합니다.

[aqua]기본값:[white] false

SMTP 서버가 사용자명/비밀번호 인증을
요구하는 경우 true로 설정하세요.`,

		"help.mail.starttls": `[yellow]StartTLS[white]

STARTTLS 암호화를 활성화합니다.

[aqua]기본값:[white] false

초기 평문 핸드셰이크 후 연결을 TLS로
업그레이드합니다. 많은 현대 SMTP 서버에서
필요합니다.`,

		"help.mail.debug": `[yellow]디버그 모드[white]

JavaMail 디버그 출력을 활성화합니다.

[aqua]기본값:[white] false

문제 해결을 위해 상세한 프로토콜 정보를
System.out에 출력합니다.`,

		// Environment 속성 도움말
		"help.env.name": `[yellow]JNDI 이름[white]

이 환경 항목의 JNDI 이름입니다.

[aqua]예:[white]
  myapp/config/maxItems

[aqua]코드 사용법:[white]
  Integer max = (Integer)
    ctx.lookup("java:comp/env/myapp/config/maxItems");`,

		"help.env.value": `[yellow]값[white]

이 환경 항목의 값입니다.

조회 시 지정된 타입으로 변환됩니다.`,

		"help.env.type": `[yellow]타입[white]

이 환경 항목의 Java 타입입니다.

[aqua]일반 타입:[white]
  • java.lang.String
  • java.lang.Integer
  • java.lang.Boolean
  • java.lang.Double`,

		"help.env.override": `[yellow]재정의[white]

애플리케이션이 이 값을 재정의할 수 있습니다.

[aqua]true:[white]
  web.xml을 통해 재정의 가능

[aqua]false:[white]
  서버 설정으로 값이 고정됨`,

		"help.env.description": `[yellow]설명[white]

이 항목의 선택적 설명입니다.

이 설정 값의 목적과 사용법을 문서화합니다.`,

		// ResourceLink 속성 도움말
		"help.reslink.name": `[yellow]로컬 이름[white]

웹 애플리케이션에서 사용하는 JNDI 이름입니다.

[aqua]예:[white]
  jdbc/LocalDB

애플리케이션이 리소스를 조회하는 데
사용하는 이름입니다.`,

		"help.reslink.global": `[yellow]전역 이름[white]

server.xml의 전역 리소스 이름입니다.

[aqua]예:[white]
  jdbc/GlobalDB

<GlobalNamingResources>에 정의된
리소스로 연결됩니다.`,

		"help.reslink.type": `[yellow]리소스 타입[white]

연결된 리소스의 Java 타입입니다.

[aqua]일반 타입:[white]
  • javax.sql.DataSource
  • javax.mail.Session
  • org.apache.catalina.UserDatabase`,

		// Connector 속성 도움말
		"help.conn.port": `[yellow]포트[white]

이 커넥터가 요청을 수신하는
TCP 포트 번호입니다.

[aqua]일반 포트:[white]
  • 8080: HTTP (개발)
  • 80: HTTP (프로덕션)
  • 8443/443: HTTPS
  • 8009: AJP`,

		"help.conn.protocol": `[yellow]프로토콜[white]

프로토콜 핸들러 구현입니다.

[aqua]HTTP:[white]
  • HTTP/1.1 (NIO/APR 자동 감지)
  • org.apache.coyote.http11.Http11NioProtocol

[aqua]AJP:[white]
  • AJP/1.3
  • org.apache.coyote.ajp.AjpNioProtocol`,

		"help.conn.timeout": `[yellow]연결 타임아웃[white]

연결 후 첫 번째 요청 데이터를 기다리는
시간(밀리초)입니다.

[aqua]기본값:[white] 60000 (60초)
[aqua]권장:[white] 20000-60000

낮은 값은 느린 클라이언트로부터
리소스를 더 빨리 해제합니다.`,

		"help.conn.redirectport": `[yellow]리다이렉트 포트[white]

SSL이 필요할 때 리다이렉트할 포트입니다.

[aqua]기본값:[white] 8443

요청이 보안을 요구하지만 비SSL 커넥터로
도착했을 때 사용됩니다.`,

		"help.conn.maxthreads": `[yellow]최대 스레드[white]

요청 처리를 위한 최대 스레드 수입니다.

[aqua]기본값:[white] 200
[aqua]권장:[white] 200-800

높은 값은 더 많은 동시 요청을 처리하지만
더 많은 메모리를 사용합니다.`,

		"help.conn.minsparethreads": `[yellow]최소 여유 스레드[white]

계속 실행되는 최소 스레드 수입니다.

[aqua]기본값:[white] 10
[aqua]권장:[white] 25-50

높은 값은 갑작스러운 트래픽 급증에
응답 시간을 개선합니다.`,

		"help.conn.acceptcount": `[yellow]수락 대기열[white]

모든 스레드가 사용 중일 때 들어오는
연결의 최대 대기열 길이입니다.

[aqua]기본값:[white] 100

이를 초과하는 연결은 거부됩니다.`,

		"help.conn.executor": `[yellow]Executor[white]

공유 스레드 풀 executor의 이름입니다.

비워두면 커넥터 전용 스레드 풀을 사용하고,
Service에 정의된 executor 이름을
지정할 수 있습니다.`,

		"help.conn.sslenabled": `[yellow]SSL 활성화[white]

이 커넥터에 SSL/TLS를 활성화합니다.

키스토어 설정이 필요합니다.`,

		"help.conn.scheme": `[yellow]스키마[white]

요청 URL의 프로토콜 스키마입니다.

[aqua]값:[white]
  • http: 비보안 연결
  • https: 보안 SSL/TLS 연결`,

		"help.conn.secure": `[yellow]보안[white]

요청을 보안으로 표시합니다.

SSL/TLS 커넥터에 true로 설정하면
request.isSecure()가 true를 반환합니다.`,

		"help.conn.keystorefile": `[yellow]키스토어 파일[white]

SSL 키스토어 파일 경로입니다.

[aqua]예:[white]
  conf/localhost-rsa.jks
  ${catalina.base}/conf/keystore.p12`,

		"help.conn.keystorepass": `[yellow]키스토어 비밀번호[white]

키스토어 파일의 비밀번호입니다.

[red]보안:[white]
프로덕션에서는 외부 비밀 관리
사용을 고려하세요.`,

		"help.conn.keystoretype": `[yellow]키스토어 타입[white]

키스토어 파일의 타입입니다.

[aqua]타입:[white]
  • JKS: Java KeyStore (레거시)
  • PKCS12: 현대 표준 (권장)`,

		"help.conn.sslprotocol": `[yellow]SSL 프로토콜[white]

SSL/TLS 프로토콜 버전입니다.

[aqua]권장:[white] TLS
[aqua]특정 버전:[white] TLSv1.2, TLSv1.3

SSLv3 및 TLSv1.0/1.1은 피하세요.`,

		"help.conn.clientauth": `[yellow]클라이언트 인증[white]

클라이언트 인증서 인증 모드입니다.

[aqua]false:[white] 클라이언트 인증서 불필요
[aqua]want:[white] 요청하지만 필수 아님
[aqua]true:[white] 클라이언트 인증서 필수`,

		"help.conn.secret": `[yellow]AJP 비밀[white]

AJP 인증을 위한 공유 비밀입니다.

secretRequired가 true일 때 필요합니다.
웹 서버 설정과 일치해야 합니다.`,

		"help.conn.secretrequired": `[yellow]비밀 필수[white]

연결에 AJP 비밀을 요구합니다.

[aqua]기본값:[white] true (Tomcat 9.0.31+)

신뢰할 수 있는 네트워크에서만
false로 설정하세요.`,

		"help.conn.address": `[yellow]주소[white]

이 커넥터를 바인딩할 IP 주소입니다.

비워두면 모든 인터페이스에 바인딩합니다.

[aqua]예:[white]
  • 127.0.0.1 (localhost만)
  • 0.0.0.0 (모든 인터페이스)`,

		// Executor 속성 도움말
		"help.exec.name": `[yellow]Executor 이름[white]

이 스레드 풀의 고유 이름입니다.

커넥터가 이 공유 executor를
참조하는 데 사용됩니다.

[aqua]예:[white]
  tomcatThreadPool`,

		"help.exec.classname": `[yellow]클래스 이름[white]

executor 구현 클래스입니다.

[aqua]표준:[white]
  org.apache.catalina.core.
    StandardThreadExecutor

[aqua]가상 스레드 (Java 21+):[white]
  org.apache.catalina.core.
    StandardVirtualThreadExecutor`,

		"help.exec.nameprefix": `[yellow]이름 접두사[white]

스레드 이름의 접두사입니다.

로그와 스레드 덤프에서 스레드를
식별하는 데 유용합니다.

[aqua]예:[white]
  catalina-exec-`,

		"help.exec.maxthreads": `[yellow]최대 스레드[white]

풀의 최대 스레드 수입니다.

[aqua]기본값:[white] 200
[aqua]권장:[white] 200-800`,

		"help.exec.minsparethreads": `[yellow]최소 여유 스레드[white]

유지되는 최소 유휴 스레드 수입니다.

[aqua]기본값:[white] 25`,

		"help.exec.maxidletime": `[yellow]최대 유휴 시간[white]

유휴 스레드가 종료되기 전
대기하는 시간(밀리초)입니다.

[aqua]기본값:[white] 60000 (1분)`,

		"help.exec.maxqueuesize": `[yellow]최대 대기열 크기[white]

대기열의 최대 대기 요청 수입니다.

[aqua]기본값:[white] Integer.MAX_VALUE

낮은 값은 과부하 시 요청을 더 빨리
거부합니다.`,

		"help.exec.prestartminsparethreads": `[yellow]최소 스레드 사전 시작[white]

시작 시 최소 스레드를 시작합니다.

[aqua]기본값:[white] false

콜드 스타트 지연을 피하려면
true로 설정하세요.`,

		// Server 속성 도움말
		"help.server.address": `[yellow]종료 주소[white]

종료 리스너의 주소입니다.

[aqua]기본값:[white] localhost

외부 인터페이스에 절대 바인딩하지 마세요.`,

		// Listener 속성 도움말
		"help.listener.classname": `[yellow]클래스 이름[white]

리스너의 정규화된 클래스 이름입니다.

[aqua]일반 리스너:[white]
  • VersionLoggerListener
  • AprLifecycleListener
  • JreMemoryLeakPreventionListener
  • ThreadLocalLeakPreventionListener`,

		"help.listener.sslengine": `[yellow]SSL 엔진[white]

APR/네이티브 커넥터용 SSL 엔진입니다.

[aqua]값:[white]
  • on: OpenSSL 엔진 활성화
  • off: JSSE 사용`,

		// Service/Engine 속성 도움말
		"help.service.name": `[yellow]Service 이름[white]

이 서비스 컨테이너의 이름입니다.

[aqua]기본값:[white] Catalina

하나의 서버에 다른 설정을 위해
여러 서비스가 존재할 수 있습니다.`,

		"help.engine.name": `[yellow]Engine 이름[white]

이 Catalina 엔진의 이름입니다.

[aqua]기본값:[white] Catalina

로깅과 JMX에서 사용됩니다.`,

		"help.engine.defaulthost": `[yellow]기본 호스트[white]

일치하지 않는 요청에 사용할 호스트입니다.

[aqua]기본값:[white] localhost

설정된 <Host> 이름과 일치해야 합니다.`,

		"help.engine.jvmroute": `[yellow]JVM 라우트[white]

세션 어피니티를 위한 라우트 ID입니다.

로드 밸런서가 이 Tomcat 인스턴스를
식별하는 데 사용됩니다.

[aqua]예:[white] node1`,

		// Host 속성 도움말
		"help.host.name": `[yellow]호스트 이름[white]

가상 호스트 이름 (도메인)입니다.

[aqua]예:[white]
  • localhost
  • www.example.com
  • *.example.com (와일드카드)`,

		"help.host.unpackwars": `[yellow]WAR 압축 해제[white]

실행 전 WAR 파일을 추출합니다.

[aqua]기본값:[white] true

압축 해제된 앱이 더 빨리 시작됩니다.`,

		"help.host.deployonstart": `[yellow]시작 시 배포[white]

Tomcat 시작 시 애플리케이션을 배포합니다.

[aqua]기본값:[white] true`,

		// Context 속성 도움말
		"help.context.path": `[yellow]컨텍스트 경로[white]

이 애플리케이션의 URL 경로입니다.

[aqua]예:[white]
  • "" (ROOT 애플리케이션)
  • /myapp
  • /api/v1`,

		"help.context.docbase": `[yellow]문서 베이스[white]

애플리케이션 파일 경로입니다.

WAR 파일 또는 디렉토리일 수 있습니다.
Host의 appBase 기준 상대 경로.`,

		"help.context.crosscontext": `[yellow]교차 컨텍스트[white]

다른 컨텍스트에 대한 액세스를 허용합니다.

[aqua]기본값:[white] false

getContext() 호출을 활성화합니다.`,

		"help.context.cookies": `[yellow]쿠키[white]

세션 추적에 쿠키를 사용합니다.

[aqua]기본값:[white] true`,

		"help.context.privileged": `[yellow]권한 있음[white]

Tomcat 내부 클래스에 액세스합니다.

[aqua]기본값:[white] false

manager 앱에 필요합니다.`,

		// Valve 속성 도움말
		"help.valve.classname": `[yellow]Valve 클래스 이름[white]

valve 구현 클래스입니다.

[aqua]일반 Valve:[white]
  • AccessLogValve
  • RemoteAddrValve
  • RemoteIpValve
  • ErrorReportValve
  • StuckThreadDetectionValve`,

		"help.valve.accesslog.pattern": `[yellow]로그 패턴[white]

액세스 로그 형식 패턴입니다.

[aqua]일반 패턴:[white]
  %h: 원격 호스트
  %l: 원격 사용자 (identd)
  %u: 인증된 사용자
  %t: 날짜/시간
  %r: 요청 라인
  %s: 상태 코드
  %b: 전송 바이트

[aqua]Combined 형식:[white]
  %h %l %u %t "%r" %s %b`,

		"help.valve.accesslog.directory": `[yellow]디렉토리[white]

로그 파일 디렉토리입니다.

[aqua]기본값:[white] logs

CATALINA_BASE 기준 상대 경로.`,

		"help.valve.accesslog.prefix": `[yellow]접두사[white]

로그 파일 이름 접두사입니다.

[aqua]기본값:[white] localhost_access_log`,

		"help.valve.accesslog.suffix": `[yellow]접미사[white]

로그 파일 이름 접미사입니다.

[aqua]기본값:[white] .txt`,

		"help.valve.accesslog.rotate": `[yellow]로테이션[white]

일별 로그 로테이션을 활성화합니다.

[aqua]기본값:[white] true`,

		"help.valve.remoteaddr.allow": `[yellow]허용 패턴[white]

허용된 IP 주소의 정규식입니다.

[aqua]예:[white]
  • 127\\.0\\.0\\.1
  • 192\\.168\\.\\d+\\.\\d+
  • 10\\.\\d+\\.\\d+\\.\\d+`,

		"help.valve.remoteaddr.deny": `[yellow]거부 패턴[white]

거부된 IP 주소의 정규식입니다.

허용 패턴 후에 확인됩니다.`,

		"help.valve.stuckthread.threshold": `[yellow]임계값[white]

스레드가 멈춘 것으로 간주되는 시간(초)입니다.

[aqua]기본값:[white] 600 (10분)

초과 시 경고를 기록합니다.`,

		// Realm 속성 도움말
		"help.realm.classname": `[yellow]Realm 클래스 이름[white]

인증 realm 구현입니다.

[aqua]일반 Realm:[white]
  • UserDatabaseRealm (파일 기반)
  • DataSourceRealm (JDBC)
  • JNDIRealm (LDAP)
  • JAASRealm (JAAS)
  • CombinedRealm (다중)`,

		"help.realm.userdatabase.resource": `[yellow]리소스 이름[white]

UserDatabase 리소스의 JNDI 이름입니다.

[aqua]기본값:[white] UserDatabase

GlobalNamingResources에 정의됩니다.`,

		"help.realm.datasource.name": `[yellow]DataSource 이름[white]

DataSource의 JNDI 이름입니다.

[aqua]예:[white]
  jdbc/AuthDB`,

		"help.realm.datasource.usertable": `[yellow]사용자 테이블[white]

사용자 자격 증명이 있는 테이블입니다.

[aqua]기본값:[white] users`,

		"help.realm.datasource.usernameCol": `[yellow]사용자명 컬럼[white]

사용자명의 컬럼 이름입니다.

[aqua]기본값:[white] user_name`,

		"help.realm.datasource.passwordCol": `[yellow]비밀번호 컬럼[white]

비밀번호의 컬럼 이름입니다.

[aqua]기본값:[white] user_pass`,

		"help.realm.datasource.roletable": `[yellow]역할 테이블[white]

사용자 역할이 있는 테이블입니다.

[aqua]기본값:[white] user_roles`,

		"help.realm.datasource.rolenameCol": `[yellow]역할 이름 컬럼[white]

역할 이름의 컬럼 이름입니다.

[aqua]기본값:[white] role_name`,

		"help.realm.jndi.connectionURL": `[yellow]연결 URL[white]

LDAP 서버 URL입니다.

[aqua]예:[white]
  ldap://ldap.example.com:389`,

		"help.realm.jndi.userbase": `[yellow]사용자 베이스[white]

사용자 검색을 위한 기본 DN입니다.

[aqua]예:[white]
  ou=users,dc=example,dc=com`,

		"help.realm.jndi.userpattern": `[yellow]사용자 패턴[white]

직접 사용자 조회를 위한 DN 패턴입니다.

[aqua]예:[white]
  uid={0},ou=users,dc=example,dc=com`,

		"help.realm.jndi.rolebase": `[yellow]역할 베이스[white]

역할 검색을 위한 기본 DN입니다.

[aqua]예:[white]
  ou=groups,dc=example,dc=com`,

		// User/Role 속성 도움말
		"help.user.username": `[yellow]사용자명[white]

사용자의 고유 식별자입니다.

인증에 사용되며 역할 할당에서
참조됩니다.`,

		"help.user.password": `[yellow]비밀번호[white]

사용자의 비밀번호입니다.

[red]보안:[white]
프로덕션에서는 해시된 비밀번호를 사용하세요.
CredentialHandler 설정을 참조하세요.`,

		"help.user.roles": `[yellow]역할[white]

쉼표로 구분된 역할 목록입니다.

[aqua]일반 역할:[white]
  • manager-gui
  • admin-gui
  • manager-script`,

		"help.role.name": `[yellow]역할 이름[white]

역할의 고유 식별자입니다.

web.xml 보안 제약 조건과 사용자
할당에서 참조됩니다.`,

		"help.role.description": `[yellow]설명[white]

선택적 역할 설명입니다.

이 역할의 목적과 권한을 문서화합니다.`,

		"help.default": `[gray]도움말 정보를 보려면 필드를 선택하세요.[-]`,
	},

	Japanese: {
		// Common
		"app.title":            "TomcatKit - Tomcat設定マネージャー",
		"app.status.ready":     "準備完了",
		"app.status.saved":     "設定が保存されました",
		"app.status.error":     "エラー: %s",
		"app.lang.select":      "言語を選択",
		"app.lang.current":     "現在: %s",
		"common.back":          "戻る",
		"common.cancel":        "キャンセル",
		"common.save":          "設定を保存",
		"common.apply":         "適用",
		"common.delete":        "削除",
		"common.add":           "追加",
		"common.edit":          "編集",
		"common.yes":           "はい",
		"common.no":            "いいえ",
		"common.confirm":       "確認",
		"common.warning":       "警告",
		"common.error":         "エラー",
		"common.success":       "成功",
		"common.loading":       "読み込み中...",
		"common.return":        "メインメニューに戻る",
		"common.enabled":       "有効",
		"common.disabled":      "無効",
		"common.notconfigured": "未設定",
		"common.minutes":       "分",
		"help.title":           "ヘルプ",
		"preview.title":        "XMLプレビュー",
		"preview.properties":   "Propertiesプレビュー",

		// Main Menu
		"menu.title":               "メインメニュー",
		"menu.server":              "サーバー",
		"menu.server.desc":         "server.xml コア設定",
		"menu.connector":           "コネクタ",
		"menu.connector.desc":      "HTTP、AJP、SSL/TLSコネクタ",
		"menu.security":            "セキュリティ / Realm",
		"menu.security.desc":       "認証レルムとユーザー",
		"menu.jndi":                "JNDIリソース",
		"menu.jndi.desc":           "DataSource、Mail Session、環境変数",
		"menu.host":                "バーチャルホスト",
		"menu.host.desc":           "Host、Context、セッションマネージャー",
		"menu.valve":               "バルブ",
		"menu.valve.desc":          "AccessLog、RemoteAddr、SSOバルブ",
		"menu.cluster":             "クラスタリング",
		"menu.cluster.desc":        "セッション複製、メンバーシップ",
		"menu.logging":             "ログ",
		"menu.logging.desc":        "JULI logging.properties",
		"menu.context":             "Context",
		"menu.context.desc":        "context.xml 設定",
		"menu.web":                 "Webアプリケーション",
		"menu.web.desc":            "web.xml サーブレット、フィルター、セキュリティ",
		"menu.quicktemplates":      "クイックテンプレート",
		"menu.quicktemplates.desc": "一般的な設定を素早く適用",
		"menu.exit":                "終了",
		"menu.exit.desc":           "TomcatKitを終了",

		// Footer
		"footer.navigate": "移動",
		"footer.select":   "選択",
		"footer.back":     "戻る",
		"footer.lang":     "言語",
		"footer.quit":     "終了",

		// Quick Templates
		"qt.title":              "クイックテンプレート",
		"qt.select":             "適用するクイックテンプレートを選択してください",
		"qt.virtualthread":      "バーチャルスレッド",
		"qt.virtualthread.desc": "バーチャルスレッドExecutorを有効化 (Java 21+、Tomcat 11+)",
		"qt.https":              "HTTPSとSSL",
		"qt.https.desc":         "SSL/TLSでHTTPSコネクタを構成",
		"qt.connpool":           "接続プールチューニング",
		"qt.connpool.desc":      "スレッドプール設定の最適化",
		"qt.gzip":               "Gzip圧縮",
		"qt.gzip.desc":          "レスポンス圧縮を有効化",
		"qt.accesslog":          "アクセスログ",
		"qt.accesslog.desc":     "アクセスログを有効化",
		"qt.security":           "セキュリティ強化",
		"qt.security.desc":      "セキュリティベストプラクティスを適用",
		"qt.apache":             "Apache httpd (mod_jk/AJP)",
		"qt.apache.desc":        "Apache httpd用AJPコネクタを構成",
		"qt.nginx":              "nginxリバースプロキシ",
		"qt.nginx.desc":         "nginx proxy_pass設定",
		"qt.haproxy":            "HAProxyロードバランサー",
		"qt.haproxy.desc":       "HAProxyロードバランシング設定",

		// Virtual Thread Template
		"qt.vt.title":            "バーチャルスレッドテンプレート",
		"qt.vt.info":             "バーチャルスレッドExecutor (Java 21+、Tomcat 11+)\n\nバーチャルスレッドは、I/Oバウンドアプリケーションの\nスループットを大幅に向上させる軽量スレッドです。",
		"qt.vt.requirements":     "要件:",
		"qt.vt.req.java":         "Java 21以上",
		"qt.vt.req.tomcat":       "Tomcat 11.0以上 (またはTomcat 10.1.25+)",
		"qt.vt.willdo":           "このテンプレートは以下を実行します:",
		"qt.vt.willdo.create":    "バーチャルスレッドExecutorを作成",
		"qt.vt.willdo.configure": "HTTPコネクタにExecutorを設定",
		"qt.vt.warning.exists":   "警告: バーチャルスレッドExecutorは既に存在します!",
		"qt.vt.executor.name":    "Executor名",
		"qt.vt.thread.prefix":    "スレッド名プレフィックス",
		"qt.vt.max.queue":        "最大キューサイズ",
		"qt.vt.apply.connector":  "適用するコネクタ",
		"qt.vt.apply":            "テンプレートを適用",
		"qt.vt.success":          "バーチャルスレッドExecutorが適用されました!",

		// HTTPS Template
		"qt.https.title":         "HTTPSテンプレート",
		"qt.https.info":          "HTTPSコネクタ設定\n\nこのテンプレートはSSL/TLSでHTTPSコネクタを作成します。",
		"qt.https.need":          "必要なもの:",
		"qt.https.need.keystore": "キーストアファイル (.jks または .p12)",
		"qt.https.need.password": "キーストアパスワード",
		"qt.https.port":          "HTTPSポート",
		"qt.https.keystore.file": "キーストアファイル",
		"qt.https.keystore.pass": "キーストアパスワード",
		"qt.https.keystore.type": "キーストアタイプ",
		"qt.https.success":       "HTTPSコネクタが追加されました!",

		// Connection Pool Template
		"qt.pool.title":           "接続プールチューニング",
		"qt.pool.info":            "接続プール最適化\n\nより良いパフォーマンスのためにスレッドプール設定を調整します。",
		"qt.pool.recommended":     "推奨設定:",
		"qt.pool.dev":             "開発: 25-100スレッド",
		"qt.pool.prod":            "本番: 150-400スレッド",
		"qt.pool.high":            "高トラフィック: 400-800スレッド",
		"qt.pool.profile":         "プロファイル",
		"qt.pool.maxthreads":      "最大スレッド",
		"qt.pool.minsparethreads": "最小スペアスレッド",
		"qt.pool.acceptcount":     "Accept Count",
		"qt.pool.conntimeout":     "接続タイムアウト (ms)",
		"qt.pool.apply":           "全HTTPコネクタに適用",
		"qt.pool.success":         "接続プール設定が適用されました!",

		// Gzip Template
		"qt.gzip.title":   "Gzip圧縮",
		"qt.gzip.info":    "Gzip圧縮\n\nテキストベースのコンテンツに対するレスポンス圧縮を有効にします。",
		"qt.gzip.willdo":  "HTTPコネクタに圧縮属性を追加します:",
		"qt.gzip.minsize": "最小圧縮サイズ (バイト)",
		"qt.gzip.success": "Gzip圧縮が有効になりました!",

		// Access Log Template
		"qt.al.title":            "アクセスログテンプレート",
		"qt.al.info":             "アクセスログバルブ\n\nHTTPリクエストのアクセスログを構成します。",
		"qt.al.patterns":         "一般的なパターン:",
		"qt.al.pattern.common":   "common: %h %l %u %t \"%r\" %s %b",
		"qt.al.pattern.combined": "combined: %h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-Agent}i\"",
		"qt.al.pattern":          "ログパターン",
		"qt.al.directory":        "ディレクトリ",
		"qt.al.prefix":           "ファイルプレフィックス",
		"qt.al.suffix":           "ファイルサフィックス",
		"qt.al.success":          "アクセスログが構成されました!",

		// Security Hardening Template
		"qt.sec.title":             "セキュリティ強化",
		"qt.sec.info":              "セキュリティ強化\n\nTomcat設定にセキュリティベストプラクティスを適用します。",
		"qt.sec.willdo":            "このテンプレートは以下を実行します:",
		"qt.sec.willdo.shutdown":   "デフォルトのshutdownポートを8005から-1に変更 (無効化)",
		"qt.sec.willdo.command":    "shutdownコマンドを変更",
		"qt.sec.willdo.version":    "エラーページからサーバーバージョンを削除",
		"qt.sec.willdo.listener":   "セキュリティ関連のリスナーを追加",
		"qt.sec.disable.shutdown":  "Shutdownポートを無効化",
		"qt.sec.remove.serverinfo": "エラーからサーバー情報を削除",
		"qt.sec.add.listener":      "セキュリティリスナーを追加",
		"qt.sec.success":           "セキュリティ強化が適用されました!",

		// Apache httpd Template
		"qt.ajp.title":          "Apache httpd (AJP) テンプレート",
		"qt.ajp.info":           "Apache httpd連携 (mod_jk / mod_proxy_ajp)\n\nApache httpdリバースプロキシ用AJPコネクタを構成します。",
		"qt.ajp.modules":        "サポートされるApacheモジュール:",
		"qt.ajp.modjk":          "mod_jk: 従来のTomcatコネクタ",
		"qt.ajp.modproxy":       "mod_proxy_ajp: AJP用Apacheプロキシモジュール",
		"qt.ajp.willdo":         "このテンプレートは以下を実行します:",
		"qt.ajp.willdo.create":  "指定ポートにAJP/1.3コネクタを作成",
		"qt.ajp.willdo.secret":  "シークレット認証を構成 (Tomcat 9.0.31+)",
		"qt.ajp.willdo.valve":   "クライアントIP処理用RemoteIpValveを追加",
		"qt.ajp.port":           "AJPポート",
		"qt.ajp.address":        "バインドアドレス",
		"qt.ajp.secret":         "AJPシークレット",
		"qt.ajp.remoteip":       "RemoteIpValveを追加",
		"qt.ajp.success":        "Apache httpd AJPコネクタが構成されました!",
		"qt.ajp.config.title":   "Apache httpd設定",
		"qt.ajp.config.applied": "Apache httpd設定が適用されました!",
		"qt.ajp.config.created": "AJPコネクタがポート%sに作成されました",
		"qt.ajp.config.copy":    "以下の設定をApache httpdにコピーしてください:",
		"qt.ajp.config.option1": "オプション1: mod_proxy_ajp (推奨)",
		"qt.ajp.config.option2": "オプション2: mod_jk (workers.properties)",
		"qt.ajp.config.note":    "注意: 設定適用後、Apache httpdを再起動してください。",

		// nginx Template
		"qt.nginx.title":          "nginxリバースプロキシテンプレート",
		"qt.nginx.info":           "nginxリバースプロキシ設定\n\nnginx proxy_pass用Tomcatを構成します。",
		"qt.nginx.willdo":         "このテンプレートは以下を実行します:",
		"qt.nginx.willdo.valve":   "X-Forwarded-*ヘッダー用RemoteIpValveを追加",
		"qt.nginx.willdo.ip":      "適切なクライアントIP処理を構成",
		"qt.nginx.willdo.http":    "オプションでHTTPコネクタ設定を調整",
		"qt.nginx.proxy.note":     "nginxがTomcatのHTTPコネクタにプロキシします",
		"qt.nginx.connector":      "HTTPコネクタ",
		"qt.nginx.internal":       "内部プロキシ (正規表現)",
		"qt.nginx.proto":          "X-Forwarded-Protoを処理",
		"qt.nginx.success":        "nginxリバースプロキシが構成されました!",
		"qt.nginx.config.title":   "nginx設定",
		"qt.nginx.config.applied": "nginxリバースプロキシ設定が適用されました!",
		"qt.nginx.config.valve":   "nginxプロキシ用RemoteIpValve構成済み",
		"qt.nginx.config.copy":    "以下の設定をnginxにコピーしてください:",
		"qt.nginx.config.basic":   "基本設定",
		"qt.nginx.config.https":   "HTTPS設定 (SSL終端)",
		"qt.nginx.config.note":    "注意: 設定適用後、nginxを再起動してください。",

		// HAProxy Template
		"qt.haproxy.title":           "HAProxyロードバランサーテンプレート",
		"qt.haproxy.info":            "HAProxyロードバランサー設定\n\nHAProxyロードバランシング用Tomcatを構成します。",
		"qt.haproxy.willdo":          "このテンプレートは以下を実行します:",
		"qt.haproxy.willdo.valve":    "X-Forwarded-*ヘッダー用RemoteIpValveを追加",
		"qt.haproxy.willdo.jvm":      "スティッキーセッション用jvmRouteを構成 (オプション)",
		"qt.haproxy.willdo.ip":       "適切なクライアントIP処理を設定",
		"qt.haproxy.modes":           "サポートされるロードバランシングモード:",
		"qt.haproxy.mode.http":       "HTTPモード (Layer 7) - 推奨",
		"qt.haproxy.mode.tcp":        "TCPモード (Layer 4) - SSLパススルー用",
		"qt.haproxy.connector":       "HTTPコネクタ",
		"qt.haproxy.sticky":          "スティッキーセッションを有効化",
		"qt.haproxy.jvmroute":        "JVM Route (ノードID)",
		"qt.haproxy.internal":        "内部プロキシ (正規表現)",
		"qt.haproxy.success":         "HAProxyロードバランサーが構成されました!",
		"qt.haproxy.config.title":    "HAProxy設定",
		"qt.haproxy.config.applied":  "HAProxyロードバランサー設定が適用されました!",
		"qt.haproxy.config.valve":    "HAProxy用RemoteIpValve構成済み",
		"qt.haproxy.config.sticky":   "スティッキーセッション注記:",
		"qt.haproxy.config.jvmroute": "Tomcat EngineにjvmRoute \"%s\" 構成済み",
		"qt.haproxy.config.cookie":   "HAProxyがセッションアフィニティにJSESSIONIDクッキーを使用",
		"qt.haproxy.config.format":   "セッションID形式: <session-id>.<jvmRoute>",
		"qt.haproxy.config.copy":     "以下の設定をHAProxyにコピーしてください:",
		"qt.haproxy.config.http":     "HTTPモード (Layer 7) - 推奨",
		"qt.haproxy.config.tcp":      "TCPモード (Layer 4) - SSLパススルー",
		"qt.haproxy.config.health":   "ヘルスチェックエンドポイント (オプション)",
		"qt.haproxy.config.stats":    "HAProxy統計 (オプション)",
		"qt.haproxy.config.note":     "注意: 設定適用後、HAProxyを再起動してください。",

		// Continue prompt
		"prompt.continue": "続行するにはEnterまたはEscapeを押してください",
		"prompt.nohttp":   "HTTPコネクタが見つかりません!",

		// Instance Selection
		"instance.title":           "Tomcatインスタンスを選択",
		"instance.recent":          "最近のインスタンス",
		"instance.detected":        "検出されたインスタンス",
		"instance.none":            "Tomcatインストールが見つかりません",
		"instance.manual":          "パスを手動入力",
		"instance.manual.desc":     "CATALINA_HOMEパスを指定",
		"instance.running":         "実行中",
		"instance.noselected":      "Tomcatインスタンスが選択されていません",
		"instance.pressT":          "'t'を押してTomcatインスタンスを選択してください",
		"instance.path.title":      "Tomcatパスを入力",
		"instance.path.home":       "CATALINA_HOME",
		"instance.path.base":       "CATALINA_BASE (オプション)",
		"instance.path.validate":   "検証して選択",
		"instance.path.required":   "CATALINA_HOMEは必須です",
		"instance.path.invalid":    "無効なパス: server.xmlが見つかりません",
		"instance.selected":        "Tomcatインスタンスが選択されました",
		"instance.info":            "Tomcatインスタンス",
		"instance.version":         "バージョン",
		"instance.status":          "ステータス",
		"instance.stopped":         "停止中",
		"instance.ready":           "設定準備完了",
		"instance.path.help.home":  "CATALINA_HOME: Tomcatインストールディレクトリ (bin, lib, confを含む)",
		"instance.path.help.base":  "CATALINA_BASE: インスタンスディレクトリ (オプション、デフォルトはCATALINA_HOME)",
		"instance.path.help.xml":   "パスにはconf/server.xmlが含まれている必要があります",
		"instance.info.noselected": "Tomcatインスタンスが選択されていません",
		"instance.info.getstarted": "開始するには:",
		"instance.info.step1":      "'t'を押してTomcatインスタンスを選択してください",
		"instance.info.step2":      "または実行: tomcatkit -home /path/to/tomcat",
		"instance.info.autodetect": "TomcatKitはインストールされたTomcatインスタンスを自動検出します。",

		// Server View
		"server.title":                      "サーバー設定",
		"server.port":                       "Shutdownポート",
		"server.port.desc":                  "サーバーシャットダウンリスナーポート",
		"server.shutdown":                   "Shutdownコマンド",
		"server.listeners":                  "リスナー",
		"server.listeners.desc":             "ライフサイクルリスナー",
		"server.services":                   "サービス",
		"server.services.desc":              "サービス設定",
		"server.globalresources":            "グローバルリソース",
		"server.globalresources.desc":       "グローバルJNDIリソース",
		"server.listener.add":               "リスナーを追加",
		"server.listener.edit":              "リスナーを編集",
		"server.listener.classname":         "クラス名",
		"server.service":                    "サービス",
		"server.service.name":               "サービス名",
		"server.engine":                     "エンジン",
		"server.engine.name":                "エンジン名",
		"server.engine.defaulthost":         "デフォルトホスト",
		"server.engine.jvmroute":            "JVM Route",
		"server.executor":                   "Executor",
		"server.executor.add":               "Executorを追加",
		"server.executor.edit":              "Executorを編集",
		"server.executor.name":              "Executor名",
		"server.executor.prefix":            "名前プレフィックス",
		"server.executor.maxthreads":        "最大スレッド",
		"server.executor.minthreads":        "最小スペアスレッド",
		"server.executor.maxidle":           "最大アイドル時間 (ms)",
		"server.executor.threads":           "スレッド: %d-%d, 最大アイドル: %dms",
		"server.executor.updated":           "Executorが更新されました",
		"server.executor.deleted":           "Executorが削除されました",
		"server.executor.added":             "Executorが追加されました",
		"server.listener.delete":            "削除",
		"server.listener.custom":            "カスタム",
		"server.listener.custom.desc":       "カスタムリスナークラスを入力",
		"server.listener.custom.title":      "カスタムリスナー",
		"server.listener.sslengine":         "SSLエンジン",
		"server.listener.sslseed":           "SSLランダムシード",
		"server.listener.updated":           "リスナーが更新されました",
		"server.listener.deleted":           "リスナーが削除されました",
		"server.listener.added":             "リスナーが追加されました",
		"server.listener.detail":            "リスナー詳細",
		"server.listener.classrequired":     "クラス名は必須です",
		"server.service.edit":               "サービス名を編集",
		"server.service.engine.desc":        "デフォルトホスト: %s, jvmRoute: %s",
		"server.service.connectors":         "コネクタ",
		"server.service.connectors.desc":    "HTTP/AJPコネクタ (コネクタメニューで設定)",
		"server.service.hosts":              "ホスト",
		"server.service.hosts.desc":         "仮想ホスト (ホストメニューで設定)",
		"server.service.updated":            "サービス名が更新されました",
		"server.engine.settings":            "エンジン設定",
		"server.engine.saved":               "エンジン設定が保存されました",
		"server.globalresource.add":         "リソースを追加",
		"server.globalresource.add.desc":    "新しいグローバルリソースを追加",
		"server.globalresource.edit":        "リソースを編集",
		"server.globalresource.auth":        "認証",
		"server.globalresource.type":        "タイプ",
		"server.globalresource.description": "説明",
		"server.globalresource.factory":     "ファクトリ",
		"server.globalresource.pathname":    "パス名",
		"server.globalresource.updated":     "リソースが更新されました",
		"server.globalresource.deleted":     "リソースが削除されました",
		"server.globalresource.added":       "リソースが追加されました",
		"server.settings":                   "サーバー設定",
		"server.settings.saved":             "サーバー設定が保存されました",
		"server.settings.invalidport":       "無効なポート番号",
		"server.confirm.delete":             "この%sを削除しますか？",
		"server.confirm.yes":                "はい",
		"server.confirm.no":                 "いいえ",

		// Connector View
		"connector.title":                   "コネクタ設定",
		"connector.http":                    "HTTPコネクタ",
		"connector.http.desc":               "HTTP/1.1, HTTP/2 コネクタ設定",
		"connector.ajp":                     "AJPコネクタ",
		"connector.ajp.desc":                "Apache JServ Protocol コネクタ設定",
		"connector.ssl":                     "SSL/TLSコネクタ",
		"connector.ssl.desc":                "HTTPSおよびSSL証明書設定",
		"connector.executor":                "Executor",
		"connector.executor.desc":           "共有スレッドプール設定",
		"connector.list":                    "コネクタ",
		"connector.add":                     "コネクタを追加",
		"connector.edit":                    "コネクタを編集",
		"connector.port":                    "ポート",
		"connector.protocol":                "プロトコル",
		"connector.timeout":                 "接続タイムアウト",
		"connector.redirect":                "リダイレクトポート",
		"connector.maxthreads":              "最大スレッド",
		"connector.minthreads":              "最小スペアスレッド",
		"connector.acceptcount":             "Accept Count",
		"connector.ssl.enabled":             "SSL有効",
		"connector.ssl.keystore":            "キーストアファイル",
		"connector.ssl.password":            "キーストアパスワード",
		"connector.ssl.type":                "キーストアタイプ",
		"connector.ssl.protocol":            "SSLプロトコル",
		"connector.ssl.clientauth":          "クライアント認証",
		"connector.http.add":                "HTTPコネクタを追加",
		"connector.http.add.desc":           "新しいHTTPコネクタを作成",
		"connector.ajp.add":                 "AJPコネクタを追加",
		"connector.ajp.add.desc":            "新しいAJPコネクタを作成",
		"connector.ssl.add":                 "HTTPSコネクタを追加",
		"connector.ssl.add.desc":            "新しいSSL/TLSコネクタを作成",
		"connector.executor.add":            "Executorを追加",
		"connector.executor.add.desc":       "新しいスレッドプールを作成",
		"connector.executor.title":          "Executor (スレッドプール)",
		"connector.service":                 "サービス",
		"connector.secret":                  "Secret",
		"connector.secret.required":         "Secret必須",
		"connector.secret.none":             "Secretなし",
		"connector.secret.set":              "Secret設定済み",
		"connector.keystore.notconfig":      "キーストアが設定されていません",
		"connector.noservices":              "サービスが設定されていません",
		"connector.updated.http":            "HTTPコネクタが更新されました",
		"connector.updated.ajp":             "AJPコネクタが更新されました",
		"connector.updated.ssl":             "SSLコネクタが更新されました",
		"connector.deleted":                 "コネクタが削除されました",
		"connector.added":                   "コネクタが追加されました",
		"connector.ssl.added":               "SSLコネクタが追加されました",
		"connector.executor.updated":        "Executorが更新されました",
		"connector.executor.deleted":        "Executorが削除されました",
		"connector.executor.added":          "Executorが追加されました",
		"connector.delete.title":            "コネクタを削除",
		"connector.delete.confirm":          "ポート%dのコネクタを削除しますか？",
		"connector.delete.ajp.confirm":      "ポート%dのAJPコネクタを削除しますか？",
		"connector.delete.ssl.confirm":      "ポート%dのSSLコネクタを削除しますか？",
		"connector.executor.delete.title":   "Executorを削除",
		"connector.executor.delete.confirm": "Executor '%s'を削除しますか？",
		"connector.executor.name":           "名前",
		"connector.executor.nameprefix":     "名前プレフィックス",
		"connector.executor.maxidle":        "最大アイドル時間 (ms)",
		"connector.executor.optional":       "Executor (オプション)",
		"connector.executor.add.title":      "Executorを追加",
		"connector.returnmenu":              "コネクタメニューに戻る",
		"connector.secretrequired":          "Secret必須",
		"connector.sslprotocol":             "SSLプロトコル",
		"connector.keystorefile":            "キーストアファイル",
		"connector.keystorepass":            "キーストアパスワード",
		"connector.keystoretype":            "キーストアタイプ",
		"connector.clientauth":              "クライアント認証",
		"connector.http.add.title":          "HTTPコネクタを追加",
		"connector.ajp.add.title":           "AJPコネクタを追加",
		"connector.ssl.add.title":           "HTTPSコネクタを追加",
		"connector.error.noservices":        "サービスが設定されていません",
		"connector.added.ssl":               "SSLコネクタが追加されました",

		// Security View
		"security.title":                "セキュリティと認証",
		"security.realm":                "Realm設定",
		"security.realm.desc":           "認証Realmの設定",
		"security.realm.add":            "Realmを追加",
		"security.realm.edit":           "Realmを編集",
		"security.realm.type":           "Realmタイプ",
		"security.realm.current":        "現在",
		"security.realm.nested":         "ネストされたRealm",
		"security.realm.set":            "Realmタイプを設定",
		"security.realm.set.desc":       "別のRealmタイプを設定",
		"security.realm.remove":         "Realmを削除",
		"security.realm.remove.desc":    "現在のRealm設定を削除",
		"security.realm.remove.confirm": "現在のRealm設定を削除しますか？",
		"security.realm.removed":        "Realmが削除されました",
		"security.realm.config":         "Realm設定",
		"security.realm.selecttype":     "Realmタイプを選択",
		"security.users":                "ユーザーとロール",
		"security.users.desc":           "tomcat-users.xmlの管理",
		"security.users.title":          "ユーザーとロール (tomcat-users.xml)",
		"security.users.list":           "ユーザー",
		"security.users.list.desc":      "ユーザーアカウントの管理",
		"security.credential":           "資格情報ハンドラー",
		"security.credential.desc":      "パスワードハッシュの設定",
		"security.user.add":             "ユーザーを追加",
		"security.user.edit":            "ユーザーを編集",
		"security.user.name":            "ユーザー名",
		"security.user.password":        "パスワード",
		"security.user.roles":           "ロール",
		"security.roles":                "ロール",
		"security.roles.list":           "ロール",
		"security.roles.list.desc":      "ロール定義の管理",
		"security.role.add":             "ロールを追加",
		"security.role.name":            "ロール名",

		// JNDI View
		"jndi.title":                "JNDIリソース - context.xml",
		"jndi.resources":            "リソース",
		"jndi.resource.add":         "リソースを追加",
		"jndi.resource.edit":        "リソースを編集",
		"jndi.resource.name":        "リソース名",
		"jndi.resource.type":        "リソースタイプ",
		"jndi.resource.auth":        "認証",
		"jndi.datasource":           "DataSource (JDBC)",
		"jndi.datasource.desc":      "データベース接続プール",
		"jndi.datasource.driver":    "ドライバクラス",
		"jndi.datasource.url":       "JDBC URL",
		"jndi.datasource.username":  "ユーザー名",
		"jndi.datasource.password":  "パスワード",
		"jndi.datasource.maxactive": "最大アクティブ",
		"jndi.datasource.maxidle":   "最大アイドル",
		"jndi.mail":                 "メールセッション",
		"jndi.mail.desc":            "JavaMail設定",
		"jndi.environment":          "環境変数",
		"jndi.environment.desc":     "環境変数の設定",
		"jndi.environment.add":      "環境変数を追加",
		"jndi.environment.name":     "エントリ名",
		"jndi.environment.value":    "値",
		"jndi.environment.type":     "タイプ",
		"jndi.resourcelink":         "リソースリンク",
		"jndi.resourcelink.desc":    "グローバルリソースへのリンク",

		// Host View
		"host.title":                  "バーチャルホストとContext",
		"host.list":                   "ホスト",
		"host.add":                    "ホストを追加",
		"host.edit":                   "ホストを編集",
		"host.name":                   "ホスト名",
		"host.appbase":                "アプリベース",
		"host.unpackwars":             "WAR展開",
		"host.autodeploy":             "自動デプロイ",
		"host.aliases":                "エイリアス",
		"host.alias.add":              "エイリアスを追加",
		"host.virtualhost":            "バーチャルホスト",
		"host.virtualhost.desc":       "バーチャルホストの管理",
		"host.context":                "Context",
		"host.context.desc":           "WebアプリケーションContextの管理",
		"host.engine":                 "エンジン設定",
		"host.engine.desc":            "Catalinaエンジンの設定",
		"context.title":               "Context設定 (context.xml)",
		"context.list":                "Context一覧",
		"context.add":                 "Contextを追加",
		"context.edit":                "Contextを編集",
		"context.path":                "Contextパス",
		"context.docbase":             "ドキュメントベース",
		"context.reloadable":          "リロード可能",
		"context.settings":            "Context設定",
		"context.settings.desc":       "基本Context属性とオプション",
		"context.resources":           "JNDIリソース",
		"context.resources.count":     "%d個のリソース (DataSource、MailSessionなど)",
		"context.environment":         "環境変数",
		"context.environment.count":   "%d個の環境変数",
		"context.resourcelinks":       "リソースリンク",
		"context.resourcelinks.count": "%d個のリソースリンク",
		"context.parameters":          "パラメータ",
		"context.parameters.count":    "%d個のContextパラメータ",
		"context.watched":             "監視リソース",
		"context.watched.count":       "%d個の監視リソース",
		"context.manager":             "セッションマネージャー",
		"context.cookie":              "クッキープロセッサ",
		"context.cookie.desc":         "SameSite、クッキー設定",
		"context.jarscanner":          "JARスキャナー",
		"context.jarscanner.desc":     "クラススキャン設定",
		"context.save.desc":           "context.xmlに変更を保存",
		"host.sessionmanager":         "セッションマネージャー",

		// Valve View
		"valve.title":             "バルブ設定",
		"valve.list":              "バルブ",
		"valve.add":               "バルブを追加",
		"valve.edit":              "バルブを編集",
		"valve.type":              "バルブタイプ",
		"valve.engine":            "エンジンバルブ",
		"valve.engine.desc":       "すべてのリクエストに適用されるバルブ",
		"valve.host":              "ホストバルブ",
		"valve.host.desc":         "特定のバーチャルホスト用バルブ",
		"valve.context":           "Contextバルブ",
		"valve.context.desc":      "特定のアプリケーション用バルブ",
		"valve.quickadd":          "よく使うバルブをクイック追加",
		"valve.quickadd.desc":     "よく使うバルブを追加",
		"valve.accesslog":         "アクセスログバルブ",
		"valve.accesslog.dir":     "ディレクトリ",
		"valve.accesslog.prefix":  "プレフィックス",
		"valve.accesslog.suffix":  "サフィックス",
		"valve.accesslog.pattern": "パターン",
		"valve.remoteaddr":        "リモートアドレスバルブ",
		"valve.remoteaddr.allow":  "許可",
		"valve.remoteaddr.deny":   "拒否",
		"valve.remoteip":          "リモートIPバルブ",
		"valve.remoteip.header":   "リモートIPヘッダー",
		"valve.remoteip.protocol": "プロトコルヘッダー",
		"valve.sso":               "シングルサインオン",
		"valve.error":             "エラーレポートバルブ",

		// Cluster View
		"cluster.title":              "クラスタリング設定",
		"cluster.enable":             "クラスタリングを有効化",
		"cluster.disable":            "クラスタリングを無効化",
		"cluster.status":             "クラスタ状態",
		"cluster.status.desc":        "クラスタリングの有効化/無効化",
		"cluster.settings":           "クラスタ設定",
		"cluster.settings.desc":      "基本クラスタ設定",
		"cluster.manager":            "セッションマネージャー",
		"cluster.manager.desc":       "DeltaManagerまたはBackupManager",
		"cluster.manager.type":       "マネージャータイプ",
		"cluster.manager.delta":      "DeltaManager",
		"cluster.manager.backup":     "BackupManager",
		"cluster.channel":            "チャネル",
		"cluster.membership":         "メンバーシップ",
		"cluster.membership.desc":    "マルチキャストメンバーシップ設定",
		"cluster.membership.address": "マルチキャストアドレス",
		"cluster.membership.port":    "マルチキャストポート",
		"cluster.receiver":           "レシーバー",
		"cluster.receiver.desc":      "メッセージレシーバー設定",
		"cluster.receiver.address":   "アドレス",
		"cluster.receiver.port":      "ポート",
		"cluster.sender":             "センダー",
		"cluster.sender.desc":        "メッセージセンダー設定",
		"cluster.interceptors":       "インターセプター",
		"cluster.interceptors.desc":  "チャネルインターセプター",
		"cluster.interceptor.add":    "インターセプターを追加",
		"cluster.deployer":           "ファームデプロイヤー",
		"cluster.deployer.desc":      "クラスタデプロイメント設定",
		"cluster.deployer.remove":    "デプロイヤーを削除",

		// Logging View
		"logging.title":              "ログ設定 (logging.properties)",
		"logging.handlers":           "ハンドラー",
		"logging.handler.add":        "ハンドラーを追加",
		"logging.handler.edit":       "ハンドラーを編集",
		"logging.handler.type":       "ハンドラータイプ",
		"logging.handler.level":      "レベル",
		"logging.handler.directory":  "ディレクトリ",
		"logging.handler.prefix":     "プレフィックス",
		"logging.filehandlers":       "ファイルハンドラー",
		"logging.filehandlers.count": "%d個のファイルハンドラー設定",
		"logging.console":            "コンソールハンドラー",
		"logging.loggers":            "ロガー",
		"logging.loggers.count":      "%d個のロガー設定",
		"logging.logger.add":         "ロガーを追加",
		"logging.logger.edit":        "ロガーを編集",
		"logging.logger.name":        "ロガー名",
		"logging.logger.level":       "レベル",
		"logging.logger.handlers":    "ハンドラー",
		"logging.rootlogger":         "ルートロガー",
		"logging.rootlogger.count":   "%d個のハンドラー割り当て",
		"logging.save.desc":          "logging.propertiesに変更を保存",

		// Context View (context.xml)
		"ctxxml.title":            "Context設定",
		"ctxxml.settings":         "基本設定",
		"ctxxml.reloadable":       "リロード可能",
		"ctxxml.crosscontext":     "Cross Context",
		"ctxxml.privileged":       "Privileged",
		"ctxxml.cookies":          "クッキー設定",
		"ctxxml.cookies.httponly": "HTTP Only",
		"ctxxml.cookies.name":     "セッションクッキー名",
		"ctxxml.resources":        "リソース",
		"ctxxml.parameters":       "パラメータ",
		"ctxxml.watched":          "監視リソース",
		"ctxxml.manager":          "セッションマネージャー",
		"ctxxml.loader":           "クラスローダー",
		"ctxxml.jarscanner":       "JARスキャナー",

		// Web View (web.xml)
		"web.title":              "Webアプリケーション設定 (web.xml)",
		"web.servlets":           "サーブレット",
		"web.servlets.count":     "%d個のサーブレット設定",
		"web.filters":            "フィルター",
		"web.filters.count":      "%d個のフィルター設定",
		"web.listeners":          "リスナー",
		"web.listeners.count":    "%d個のリスナー設定",
		"web.session":            "セッション設定",
		"web.welcomefiles":       "ウェルカムファイル",
		"web.welcomefiles.desc":  "デフォルトページファイル",
		"web.errorpages":         "エラーページ",
		"web.errorpages.count":   "%d個のエラーページ",
		"web.mime":               "MIMEマッピング",
		"web.mime.count":         "%d個のMIMEタイプ",
		"web.security":           "セキュリティ制約",
		"web.security.count":     "%d個の制約",
		"web.login":              "ログイン設定",
		"web.login.desc":         "認証方法",
		"web.roles":              "セキュリティロール",
		"web.roles.desc":         "セキュリティロールを定義",
		"web.contextparams":      "Contextパラメータ",
		"web.contextparams.desc": "Webアプリ初期化パラメータ",
		"web.save.desc":          "web.xmlに変更を保存",
		// Web sub-menu keys
		"web.servlet.add":                     "サーブレットを追加",
		"web.servlet.add.desc":                "新しいサーブレットを作成",
		"web.servlet.edit":                    "サーブレットを編集",
		"web.servlet.name":                    "サーブレット名",
		"web.servlet.class":                   "サーブレットクラス",
		"web.servlet.jsp":                     "JSPファイル (オプション)",
		"web.servlet.loadonstartup":           "起動時にロード",
		"web.servlet.async":                   "非同期サポート",
		"web.servlet.initparams":              "初期化パラメータ (1行にname=value)",
		"web.servlet.urlpatterns":             "URLパターン (1行に1つ)",
		"web.servlet.quickdefault":            "デフォルトサーブレットを追加",
		"web.servlet.quickdefault.desc":       "Tomcatデフォルトサーブレットを追加",
		"web.servlet.quickjsp":                "JSPサーブレットを追加",
		"web.servlet.quickjsp.desc":           "Tomcat JSPサーブレットを追加",
		"web.servlet.added":                   "サーブレットが追加されました",
		"web.servlet.updated":                 "サーブレットが更新されました",
		"web.servlet.deleted":                 "サーブレットが削除されました",
		"web.servlet.error.name":              "サーブレット名は必須です",
		"web.servlet.error.class":             "サーブレットクラスまたはJSPファイルが必要です",
		"web.filter.add":                      "フィルターを追加",
		"web.filter.add.desc":                 "新しいフィルターを作成",
		"web.filter.edit":                     "フィルターを編集",
		"web.filter.name":                     "フィルター名",
		"web.filter.class":                    "フィルタークラス",
		"web.filter.async":                    "非同期サポート",
		"web.filter.initparams":               "初期化パラメータ (1行にname=value)",
		"web.filter.urlpatterns":              "URLパターン (1行に1つ)",
		"web.filter.quickcors":                "CORSフィルターを追加",
		"web.filter.quickcors.desc":           "CORSフィルターを追加",
		"web.filter.quickencoding":            "エンコーディングフィルターを追加",
		"web.filter.quickencoding.desc":       "文字エンコーディングフィルターを追加",
		"web.filter.added":                    "フィルターが追加されました",
		"web.filter.updated":                  "フィルターが更新されました",
		"web.filter.deleted":                  "フィルターが削除されました",
		"web.filter.error.required":           "フィルター名とクラスは必須です",
		"web.listener.add":                    "リスナーを追加",
		"web.listener.add.desc":               "新しいリスナーを作成",
		"web.listener.edit":                   "リスナーを編集",
		"web.listener.class":                  "リスナークラス",
		"web.listener.description":            "説明",
		"web.listener.added":                  "リスナーが追加されました",
		"web.listener.deleted":                "リスナーが削除されました",
		"web.listener.error.class":            "リスナークラスは必須です",
		"web.session.title":                   "セッション設定",
		"web.session.timeout":                 "セッションタイムアウト (分)",
		"web.session.tracking.cookie":         "トラッキング: COOKIE",
		"web.session.tracking.url":            "トラッキング: URL",
		"web.session.tracking.ssl":            "トラッキング: SSL",
		"web.session.cookie.name":             "Cookie名",
		"web.session.cookie.domain":           "Cookieドメイン",
		"web.session.cookie.path":             "Cookieパス",
		"web.session.cookie.httponly":         "HttpOnly Cookie",
		"web.session.cookie.secure":           "Secure Cookie",
		"web.session.saved":                   "セッション設定が保存されました",
		"web.welcomefiles.title":              "ウェルカムファイル",
		"web.welcomefiles.perline":            "ウェルカムファイル (1行に1つ)",
		"web.welcomefiles.adddefaults":        "デフォルトを追加",
		"web.welcomefiles.saved":              "ウェルカムファイルが保存されました",
		"web.welcomefiles.defaultadded":       "デフォルトウェルカムファイルが追加されました",
		"web.errorpage.add":                   "エラーページを追加",
		"web.errorpage.add.desc":              "新しいエラーページを作成",
		"web.errorpage.edit":                  "エラーページを編集",
		"web.errorpage.type":                  "エラータイプ",
		"web.errorpage.code":                  "エラーコード (例: 404)",
		"web.errorpage.exception":             "例外タイプ",
		"web.errorpage.location":              "場所 (ページパス)",
		"web.errorpage.quickcommon":           "一般的なエラーページを追加",
		"web.errorpage.quickcommon.desc":      "404, 500エラーページを追加",
		"web.errorpage.added":                 "エラーページが追加されました",
		"web.errorpage.deleted":               "エラーページが削除されました",
		"web.errorpage.commonadded":           "一般的なエラーページが追加されました",
		"web.errorpage.error.location":        "場所は必須です",
		"web.mime.add":                        "MIMEマッピングを追加",
		"web.mime.add.desc":                   "新しいMIMEマッピングを作成",
		"web.mime.edit":                       "MIMEマッピングを編集",
		"web.mime.extension":                  "拡張子 (ドットなし)",
		"web.mime.type":                       "MIMEタイプ",
		"web.mime.quickcommon":                "一般的なMIMEタイプを追加",
		"web.mime.quickcommon.desc":           "一般的なWebのMIMEタイプを追加",
		"web.mime.added":                      "MIMEマッピングが追加されました",
		"web.mime.deleted":                    "MIMEマッピングが削除されました",
		"web.mime.commonadded":                "一般的なMIMEタイプが追加されました",
		"web.mime.error.required":             "拡張子とMIMEタイプは必須です",
		"web.securityconstraint.add":          "セキュリティ制約を追加",
		"web.securityconstraint.add.desc":     "新しいセキュリティ制約を作成",
		"web.securityconstraint.edit":         "セキュリティ制約を編集",
		"web.securityconstraint.resourcename": "リソース名",
		"web.securityconstraint.urlpatterns":  "URLパターン (1行に1つ)",
		"web.securityconstraint.httpmethods":  "HTTPメソッド (カンマ区切り、空=すべて)",
		"web.securityconstraint.roles":        "必要なロール (1行に1つ)",
		"web.securityconstraint.transport":    "転送保証",
		"web.securityconstraint.added":        "セキュリティ制約が追加されました",
		"web.securityconstraint.updated":      "セキュリティ制約が更新されました",
		"web.securityconstraint.deleted":      "セキュリティ制約が削除されました",
		"web.login.title":                     "ログイン設定",
		"web.login.authmethod":                "認証方法",
		"web.login.realmname":                 "Realm名",
		"web.login.formloginpage":             "フォームログインページ",
		"web.login.formerrorpage":             "フォームエラーページ",
		"web.login.saved":                     "ログイン設定が保存されました",
		"web.login.removed":                   "ログイン設定が削除されました",
		"web.role.add":                        "セキュリティロールを追加",
		"web.role.add.desc":                   "新しいセキュリティロールを作成",
		"web.role.edit":                       "セキュリティロールを編集",
		"web.role.name":                       "ロール名",
		"web.role.description":                "説明",
		"web.role.quickcommon":                "一般的なロールを追加",
		"web.role.quickcommon.desc":           "admin, user, managerロールを追加",
		"web.role.added":                      "セキュリティロールが追加されました",
		"web.role.deleted":                    "セキュリティロールが削除されました",
		"web.role.commonadded":                "一般的なロールが追加されました",
		"web.role.error.name":                 "ロール名は必須です",
		"web.contextparam.add":                "Contextパラメータを追加",
		"web.contextparam.add.desc":           "新しいパラメータを作成",
		"web.contextparam.edit":               "Contextパラメータを編集",
		"web.contextparam.name":               "パラメータ名",
		"web.contextparam.value":              "パラメータ値",
		"web.contextparam.description":        "説明",
		"web.contextparam.added":              "Contextパラメータが追加されました",
		"web.contextparam.updated":            "Contextパラメータが更新されました",
		"web.contextparam.deleted":            "Contextパラメータが削除されました",
		"web.contextparam.error.name":         "パラメータ名は必須です",
		"webxml.title":                        "Webアプリケーション設定",
		"webxml.servlets":                     "サーブレット",
		"webxml.servlet.add":                  "サーブレットを追加",
		"webxml.servlet.edit":                 "サーブレットを編集",
		"webxml.servlet.name":                 "サーブレット名",
		"webxml.servlet.class":                "サーブレットクラス",
		"webxml.servlet.mapping":              "URLパターン",
		"webxml.filters":                      "フィルター",
		"webxml.filter.add":                   "フィルターを追加",
		"webxml.filter.edit":                  "フィルターを編集",
		"webxml.filter.name":                  "フィルター名",
		"webxml.filter.class":                 "フィルタークラス",
		"webxml.listeners":                    "リスナー",
		"webxml.listener.add":                 "リスナーを追加",
		"webxml.listener.class":               "リスナークラス",
		"webxml.session":                      "セッション設定",
		"webxml.session.timeout":              "セッションタイムアウト (分)",
		"webxml.welcome":                      "ウェルカムファイル",
		"webxml.error":                        "エラーページ",
		"webxml.mime":                         "MIMEマッピング",
		"webxml.security":                     "セキュリティ制約",
		"webxml.security.add":                 "セキュリティ制約を追加",
		"webxml.login":                        "ログイン設定",
		"webxml.login.method":                 "認証方法",
		"webxml.roles":                        "セキュリティロール",

		// ==================== 詳細ヘルプ ====================
		// サーバー設定ヘルプ
		"help.server.port": `[::b]Shutdownポート[::-]
Tomcatがシャットダウンコマンドを受信するTCP/IPポート番号です。

[yellow]デフォルト:[-] 8005
[yellow]範囲:[-] 1-65535 (または-1で無効化)

[green]セキュリティ注意:[-]
• 本番環境では-1に設定してリモートシャットダウンを無効化
• localhost (127.0.0.1) でのみリッスン
• シャットダウンスクリプトのポートと一致させる必要あり

[gray]例: <Server port="8005" shutdown="SHUTDOWN">[-]`,

		"help.server.shutdown": `[::b]Shutdownコマンド[::-]
シャットダウンをトリガーするために受信する必要があるコマンド文字列です。

[yellow]デフォルト:[-] SHUTDOWN

[green]セキュリティ推奨事項:[-]
• デフォルトの"SHUTDOWN"を複雑でランダムな文字列に変更
• port=-1と組み合わせて最大のセキュリティを提供
• catalina.sh stop / shutdown.batで使用

[gray]例: <Server port="8005" shutdown="複雑な秘密の文字列">[-]`,

		"help.server.listener": `[::b]ライフサイクルリスナー[::-]
リスナーはサーバーライフサイクルの特定のイベントに応答します。

[green]一般的なリスナー:[-]
• [yellow]VersionLoggerListener[-]: 起動時にTomcatバージョン情報をログ
• [yellow]AprLifecycleListener[-]: Apache Portable Runtime (APR) を有効化
• [yellow]JreMemoryLeakPreventionListener[-]: JREメモリリークを防止
• [yellow]GlobalResourcesLifecycleListener[-]: JNDIリソースに必要
• [yellow]ThreadLocalLeakPreventionListener[-]: スレッドローカル変数をクリーンアップ

[gray]すべてのリスナーはオプションですが、本番環境では推奨されます。[-]`,

		"help.server.service": `[::b]サービス[::-]
サービスは1つ以上のコネクタを単一のエンジンとグループ化します。

[yellow]デフォルト名:[-] Catalina

[green]コンポーネント:[-]
• [yellow]Engine[-]: リクエスト処理エンジン (サービスごとに1つ)
• [yellow]Connectors[-]: HTTP, HTTPS, AJP (1つ以上)
• [yellow]Executors[-]: 共有スレッドプール (オプション)

[gray]ほとんどのインストールは"Catalina"という単一のサービスを使用します。[-]`,

		"help.server.engine": `[::b]エンジン[::-]
エンジンはコネクタからすべてのリクエストを受信して処理します。

[green]属性:[-]
• [yellow]name[-]: 論理名 (デフォルト: Catalina)
• [yellow]defaultHost[-]: マッチしないリクエストのデフォルトホスト
• [yellow]jvmRoute[-]: ロードバランサースティッキーセッション用の一意のID

[green]jvmRouteの使用:[-]
セッションID形式: <session-id>.<jvmRoute>
例: ABC123.node1

[gray]クラスター環境でのセッションアフィニティに必要です。[-]`,

		"help.server.executor": `[::b]Executor (スレッドプール)[::-]
コネクタ用の共有スレッドプールです。

[green]属性:[-]
• [yellow]name[-]: 一意の識別子 (コネクタから参照)
• [yellow]maxThreads[-]: 最大ワーカースレッド (デフォルト: 200)
• [yellow]minSpareThreads[-]: 最小アイドルスレッド (デフォルト: 25)
• [yellow]maxIdleTime[-]: アイドルスレッドタイムアウト ms (デフォルト: 60000)

[green]サイジングガイドライン:[-]
• 開発: 25-100スレッド
• 本番: 150-400スレッド
• 高トラフィック: 400-800スレッド

[gray]コネクタはexecutor="name"属性でExecutorを参照します。[-]`,

		// コネクタヘルプ
		"help.connector.http": `[::b]HTTPコネクタ[::-]
HTTP/1.1およびHTTP/2クライアント接続を処理します。

[green]主な属性:[-]
• [yellow]port[-]: TCPポート (デフォルト: 8080)
• [yellow]protocol[-]: HTTP/1.1, org.apache.coyote.http11.Http11NioProtocol
• [yellow]connectionTimeout[-]: ソケットタイムアウト ms (デフォルト: 20000)
• [yellow]maxThreads[-]: 最大リクエストスレッド (デフォルト: 200)
• [yellow]acceptCount[-]: バックログキューサイズ (デフォルト: 100)

[green]プロトコルオプション:[-]
• HTTP/1.1 - 最適な実装を自動選択
• Http11NioProtocol - ノンブロッキングI/O (推奨)
• Http11Nio2Protocol - NIO2/AIO実装
• Http11AprProtocol - APR/native (APRライブラリが必要)`,

		"help.connector.https": `[::b]HTTPSコネクタ (SSL/TLS)[::-]
SSL/TLS暗号化によるセキュアなHTTP接続を提供します。

[green]主な属性:[-]
• [yellow]port[-]: TCPポート (デフォルト: 8443)
• [yellow]SSLEnabled[-]: "true"である必要あり
• [yellow]keystoreFile[-]: キーストアへのパス (.jks, .p12)
• [yellow]keystorePass[-]: キーストアパスワード
• [yellow]keystoreType[-]: JKS, PKCS12 (推奨)

[green]SSLプロトコルオプション:[-]
• TLS (最高バージョンを自動ネゴシエート)
• TLSv1.2 (最小推奨)
• TLSv1.3 (最も安全、Java 11+)

[green]セキュリティ推奨事項:[-]
• TLSv1.2またはTLSv1.3のみ使用
• 弱い暗号を無効化
• 2048ビット以上のRSAまたは256ビット以上のECCキーを使用`,

		"help.connector.ajp": `[::b]AJPコネクタ[::-]
Apache httpd連携用のApache JServ Protocolです。

[green]主な属性:[-]
• [yellow]port[-]: TCPポート (デフォルト: 8009)
• [yellow]protocol[-]: AJP/1.3
• [yellow]secretRequired[-]: 共有シークレットが必要 (9.0.31+でデフォルト: true)
• [yellow]secret[-]: 共有認証シークレット
• [yellow]address[-]: バインドアドレス (デフォルト: すべてのインターフェース)

[green]セキュリティ (Tomcat 9.0.31+):[-]
• secretRequired="true"で認証を強制
• secretはApache httpdのProxyPass secretと一致させる
• address="127.0.0.1"でlocalhostに制限

[gray]例: ProxyPass /app ajp://localhost:8009/app secret=mySecret[-]`,

		"help.connector.port": `[::b]ポート[::-]
受信接続用のTCPポート番号です。

[yellow]デフォルトポート:[-]
• HTTP: 8080
• HTTPS: 8443
• AJP: 8009

[green]注意事項:[-]
• 1024未満のポートはroot/admin権限が必要
• 80→8080リダイレクトにiptables/ファイアウォールを使用
• 各コネクタは一意のポートを使用する必要あり`,

		"help.connector.protocol": `[::b]プロトコル[::-]
プロトコルハンドラー実装です。

[green]HTTPプロトコル:[-]
• [yellow]HTTP/1.1[-]: 最適な実装を自動選択
• [yellow]Http11NioProtocol[-]: ノンブロッキングI/O (デフォルト)
• [yellow]Http11Nio2Protocol[-]: Java NIO2 (非同期)
• [yellow]Http11AprProtocol[-]: APR/native (APRが必要)

[green]AJPプロトコル:[-]
• [yellow]AJP/1.3[-]: 標準AJPプロトコル
• [yellow]AjpNioProtocol[-]: NIO実装
• [yellow]AjpNio2Protocol[-]: NIO2実装

[gray]ほとんどのデプロイメントにはNIOが推奨されます。[-]`,

		"help.connector.maxthreads": `[::b]最大スレッド[::-]
リクエスト処理スレッドの最大数です。

[yellow]デフォルト:[-] 200

[green]サイジングガイドライン:[-]
• 低トラフィック: 25-100
• 中トラフィック: 150-300
• 高トラフィック: 400-800

[green]計算式:[-]
maxThreads ≈ (ピーク同時ユーザー × 平均応答時間秒)

[yellow]警告:[-]
• 少なすぎ: リクエストがキューに溜まりタイムアウト
• 多すぎ: メモリ枯渇、コンテキストスイッチングオーバーヘッド
• 本番環境でスレッドプール使用率を監視`,

		"help.connector.minsparethreads": `[::b]最小スペアスレッド[::-]
維持される最小アイドルスレッド数です。

[yellow]デフォルト:[-] 10

[green]目的:[-]
• 受信リクエストに備えてスレッドを準備状態に維持
• 初期リクエストのレイテンシを削減
• バースト時のスレッド作成オーバーヘッドを防止

[green]推奨事項:[-]
• ほとんどのアプリケーションで10-25を設定
• バースト性のあるトラフィックパターンには高い値
• メモリ使用量を減らすには低い値`,

		"help.connector.connectiontimeout": `[::b]接続タイムアウト[::-]
接続後にリクエストデータを待つミリ秒数です。

[yellow]デフォルト:[-] 20000 (20秒)

[green]動作:[-]
• TCP接続確立後にタイマー開始
• データ受信ごとにリセット
• 超過時に接続をクローズ

[green]推奨事項:[-]
• 通常のアプリケーション: 20000-60000
• ロードバランサーの背後: 低め (5000-10000)
• 遅いクライアントや大容量アップロード: 高め`,

		"help.connector.acceptcount": `[::b]Accept Count[::-]
受信接続の最大キュー長です。

[yellow]デフォルト:[-] 100

[green]動作:[-]
• すべてのスレッドがビジーのとき接続がキューに待機
• 超過すると"connection refused"を返す
• OSに独自の下限がある場合あり

[green]チューニング:[-]
• トラフィック急増に備えて増加 (200-500)
• 過負荷時の高速フェイルのために減少
• 本番環境でキュー深度を監視`,

		"help.connector.redirectport": `[::b]リダイレクトポート[::-]
自動HTTPSリダイレクト用のポートです。

[yellow]デフォルト:[-] 8443

[green]使用方法:[-]
• security-constraintがHTTPSを要求するときに使用
• 自動的にHTTP→HTTPSリダイレクト
• HTTPSコネクタポートと一致させる必要あり

[gray]例: 8080のHTTPが8443のHTTPSにリダイレクト[-]`,

		"help.connector.executor": `[::b]Executor (スレッドプール)[::-]
共有スレッドプールへの参照です。

[yellow]使用方法:[-]
• 空のままでコネクタ独自のスレッドプールを使用
• Executor名を設定してコネクタ間でスレッドを共有

[green]利点:[-]
• 集中化されたスレッドプール管理
• リソースの効率的な活用
• 監視とチューニングが容易

[gray]server.xmlでExecutorを定義:
<Executor name="tomcatThreadPool"
  maxThreads="200" minSpareThreads="10"/>[-]`,

		"help.connector.secretrequired": `[::b]Secret Required[::-]
AJP接続のシークレットベース認証を有効にします。

[yellow]デフォルト:[-] true (Tomcat 9.0.31+)

[green]目的:[-]
• AJPポートへの不正アクセスを防止
• Ghostcat脆弱性を軽減 (CVE-2020-1938)
• 両側で一致するシークレットが必要

[red]セキュリティ:[-]
• 本番環境では常に有効化
• 強力でランダムなシークレット値を使用
• 信頼できないネットワークにAJPポートを公開しない`,

		"help.connector.secret": `[::b]Secret[::-]
AJPコネクタ認証用の共有シークレットです。

[yellow]要件:[-]
• Apache mod_proxy_ajpのシークレットと一致する必要あり
• 強力でランダムな値を使用 (32文字以上)
• 機密保持

[green]設定:[-]
• Tomcat: secret="yourSecretValue"
• Apache: ProxyPass ajp://host:8009 secret=yourSecretValue

[gray]生成コマンド: openssl rand -base64 32[-]`,

		// セキュリティヘルプ
		"help.security.realm": `[::b]Realm[::-]
RealmはTomcatを認証用のユーザー/ロールデータベースに接続します。

[green]Realmタイプ:[-]
• [yellow]UserDatabaseRealm[-]: tomcat-users.xml使用 (デフォルト)
• [yellow]JDBCRealm[-]: JDBC経由のデータベース
• [yellow]DataSourceRealm[-]: JNDI DataSource経由のデータベース
• [yellow]JNDIRealm[-]: LDAP/Active Directory
• [yellow]JAASRealm[-]: Java認証 (JAAS)

[green]配置:[-]
• Engineレベル: すべてのHost/Contextに適用
• Hostレベル: Host内のすべてのContextに適用
• Contextレベル: 単一のアプリケーションに適用

[gray]CombinedRealmまたはLockOutRealmでネスト可能です。[-]`,

		"help.security.userdatabase": `[::b]UserDatabaseRealm[::-]
tomcat-users.xmlを使用するファイルベースの認証です。

[green]設定:[-]
• ファイル: conf/tomcat-users.xml
• Roles: アクセス権限を定義
• Users: ユーザー名、パスワード、ロール割り当て

[green]デフォルトロール:[-]
• manager-gui: Managerウェブインターフェースアクセス
• manager-script: Managerテキスト/スクリプトインターフェースアクセス
• admin-gui: Host Managerインターフェースアクセス

[yellow]セキュリティ注意:[-]
tomcat-users.xmlのパスワードはダイジェストする必要があります:
1. 実行: digest.sh -a SHA-256 mypassword
2. 出力をpassword属性に使用`,

		"help.security.jdbcrealm": `[::b]JDBCRealm / DataSourceRealm[::-]
データベースベースの認証です。

[green]必要なテーブル:[-]
• Usersテーブル: username, passwordカラム
• Rolesテーブル: username, role_nameカラム

[green]DataSourceRealm属性:[-]
• dataSourceName: JNDI名 (例: jdbc/UserDB)
• userTable: ユーザー名を含むテーブル
• userNameCol: ユーザー名カラム名
• userCredCol: パスワードカラム名
• userRoleTable: ロールを含むテーブル
• roleNameCol: ロールカラム名

[gray]接続プーリングのためJDBCRealmよりDataSourceRealmが推奨されます。[-]`,

		"help.security.ldaprealm": `[::b]JNDIRealm (LDAP/Active Directory)[::-]
LDAPベースの認証と認可です。

[green]主な属性:[-]
• connectionURL: ldap://server:389 または ldaps://server:636
• userPattern: DNパターン、例: uid={0},ou=users,dc=example,dc=com
• userSearch: ユーザー検索フィルター
• roleBase: ロール検索の基本DN
• roleName: ロール名を含む属性

[green]Active Directory例:[-]
• userPattern: {0}@domain.com
• userSearch: (sAMAccountName={0})
• roleSearch: (member={0})

[gray]大量認証にはconnectionPoolSizeを使用してください。[-]`,

		// JNDIヘルプ
		"help.jndi.datasource": `[::b]JNDI DataSource[::-]
JNDIルックアップでアクセス可能なデータベース接続プールです。

[green]主な属性:[-]
• [yellow]name[-]: JNDI名 (例: jdbc/MyDB)
• [yellow]type[-]: javax.sql.DataSource
• [yellow]driverClassName[-]: JDBCドライバクラス
• [yellow]url[-]: JDBC接続URL
• [yellow]username/password[-]: データベース資格情報

[green]接続プール設定:[-]
• maxTotal: 最大接続数 (デフォルト: 8)
• maxIdle: 最大アイドル接続数
• minIdle: 最小アイドル接続数
• maxWaitMillis: 接続待機タイムアウト

[green]アプリケーションでの使用:[-]
Context ctx = new InitialContext();
DataSource ds = (DataSource) ctx.lookup("java:comp/env/jdbc/MyDB");`,

		"help.jndi.environment": `[::b]環境エントリ[::-]
設定のためにJNDIでアクセス可能な単純な値です。

[green]サポートされるタイプ:[-]
• java.lang.String
• java.lang.Integer
• java.lang.Boolean
• java.lang.Double

[green]例:[-]
<Environment name="maxItems" value="100" type="java.lang.Integer"/>

[green]アプリケーションでの使用:[-]
Context ctx = new InitialContext();
Integer maxItems = (Integer) ctx.lookup("java:comp/env/maxItems");

[gray]アプリケーション設定を外部化するのに便利です。[-]`,

		// Hostヘルプ
		"help.host": `[::b]バーチャルホスト[::-]
Hostは独自のアプリケーションを持つバーチャルホストを表します。

[green]主な属性:[-]
• [yellow]name[-]: ホスト名 (例: www.example.com)
• [yellow]appBase[-]: アプリケーションディレクトリ (デフォルト: webapps)
• [yellow]unpackWARs[-]: WARファイル自動展開 (デフォルト: true)
• [yellow]autoDeploy[-]: アプリケーションホットデプロイ (デフォルト: true)

[green]エイリアス:[-]
このHostにマッピングされる追加のホスト名。
例: example.com → www.example.com

[green]ディレクトリ構造:[-]
• $CATALINA_BASE/webapps/ - デフォルトappBase
• 各サブディレクトリまたはWARはアプリケーション`,

		"help.host.appbase": `[::b]アプリケーションベース[::-]
このHostのWebアプリケーションを含むディレクトリです。

[yellow]デフォルト:[-] webapps

[green]パスタイプ:[-]
• 相対: $CATALINA_BASE基準
• 絶対: 完全なファイルシステムパス

[green]デプロイ方法:[-]
• appBaseにWARファイルをドロップ
• Webアプリケーションを含むディレクトリを作成
• conf/Catalina/[host]/でContext記述子を使用`,

		"help.host.autodeploy": `[::b]自動デプロイ[::-]
ファイル変更時にアプリケーションを自動的にデプロイします。

[yellow]デフォルト:[-] true

[green]動作:[-]
• 新規/変更されたファイルのappBaseを監視
• 新しいアプリケーションを自動デプロイ
• 変更されたアプリケーションをリロード

[yellow]本番環境推奨事項:[-]
以下のために"false"に設定:
• セキュリティ向上 (不正なデプロイを防止)
• パフォーマンス向上 (ファイル監視なし)
• 予測可能なデプロイタイミング`,

		// Valveヘルプ
		"help.valve.accesslog": `[::b]アクセスログValve[::-]
HTTPリクエストをファイルにログします。

[green]主な属性:[-]
• [yellow]directory[-]: ログディレクトリ (デフォルト: logs)
• [yellow]prefix[-]: ログファイルプレフィックス (デフォルト: localhost_access_log)
• [yellow]suffix[-]: ログファイルサフィックス (デフォルト: .txt)
• [yellow]pattern[-]: ログフォーマットパターン

[green]パターン変数:[-]
• %h - リモートホスト名/IP
• %l - リモート論理ユーザー名 (常に -)
• %u - 認証されたユーザー名
• %t - Common Log Formatの日時
• %r - リクエストの最初の行
• %s - HTTPステータスコード
• %b - 送信バイト (0の場合は -)
• %D - リクエスト処理時間 (ms)
• %T - リクエスト処理時間 (秒)

[green]一般的なパターン:[-]
• common: %h %l %u %t "%r" %s %b
• combined: + RefererとUser-Agent`,

		"help.valve.remoteip": `[::b]Remote IP Valve[::-]
プロキシヘッダーから実際のクライアントIPを抽出します。

[green]主な属性:[-]
• [yellow]remoteIpHeader[-]: クライアントIPを含むヘッダー (X-Forwarded-For)
• [yellow]protocolHeader[-]: プロトコルヘッダー (X-Forwarded-Proto)
• [yellow]internalProxies[-]: 信頼できるプロキシIPの正規表現

[green]ユースケース:[-]
リバースプロキシ (nginx, Apache, HAProxy) の背後:
• 実際のIPがX-Forwarded-Forヘッダーに
• request.getRemoteAddr()はプロキシIPを返す

[green]設定:[-]
<Valve className="...RemoteIpValve"
       remoteIpHeader="X-Forwarded-For"
       protocolHeader="X-Forwarded-Proto"/>

[gray]プロキシの背後での正確なログとセキュリティに不可欠です。[-]`,

		"help.valve.remoteaddr": `[::b]リモートアドレスフィルター[::-]
クライアントIPアドレスに基づいてアクセスを制限します。

[green]属性:[-]
• [yellow]allow[-]: 許可されたIPの正規表現
• [yellow]deny[-]: 拒否されたIPの正規表現

[green]評価順序:[-]
1. まず拒否パターンを確認
2. 許可パターンを確認
3. マッチなし: allowが指定されていれば拒否、そうでなければ許可

[green]例:[-]
• localhostのみ許可: allow="127\.\d+\.\d+\.\d+"
• 内部ネットワーク許可: allow="192\.168\.\d+\.\d+"
• 特定IP拒否: deny="10\.0\.0\.1"

[yellow]注意:[-] プロキシの背後ではRemoteIpValveと一緒に使用してください。`,

		"help.valve.sso": `[::b]シングルサインオンValve[::-]
アプリケーション間のシングルサインオンを有効にします。

[green]動作:[-]
• ユーザーが一度ログイン
• 同じHost内のすべてのアプリでセッション共有
• 1つのアプリからログアウトするとすべてからログアウト

[green]設定:[-]
Hostレベルに配置:
<Valve className="...SingleSignOn"/>

[green]要件:[-]
• すべてのアプリが同じRealmを使用
• アプリが同じバーチャルHostにある
• 本番環境ではセキュアクッキーを使用`,

		// クラスターヘルプ
		"help.cluster": `[::b]Tomcatクラスタリング[::-]
複数のTomcatインスタンス間のセッションレプリケーションです。

[green]コンポーネント:[-]
• [yellow]Manager[-]: セッションレプリケーション管理
• [yellow]Channel[-]: グループ通信
• [yellow]Membership[-]: クラスターメンバー検出
• [yellow]Receiver[-]: レプリケートされたデータを受信
• [yellow]Sender[-]: レプリケートされたデータを送信

[green]Managerタイプ:[-]
• [yellow]DeltaManager[-]: すべてのノードにレプリケート (小規模クラスター)
• [yellow]BackupManager[-]: 1つのバックアップにレプリケート (大規模クラスター)

[green]要件:[-]
• web.xmlに<distributable/>
• すべてのセッション属性はSerializableである必要
• マルチキャストネットワークまたは静的メンバーシップ`,

		"help.cluster.membership": `[::b]クラスターメンバーシップ[::-]
クラスターメンバーを検出して追跡します。

[green]マルチキャストメンバーシップ:[-]
• address: マルチキャストグループ (デフォルト: 228.0.0.4)
• port: マルチキャストポート (デフォルト: 45564)
• frequency: ハートビート間隔 (ms)
• dropTime: メンバードロップアウトタイムアウト (ms)

[green]静的メンバーシップ:[-]
マルチキャストのないネットワーク用:
• 各メンバーを明示的に設定
• StaticMembershipInterceptorを使用

[yellow]ネットワーク要件:[-]
• UDPマルチキャストを有効にする必要あり
• ファイアウォールがマルチキャストトラフィックを許可
• すべてのノードが同じマルチキャストグループに`,

		// ロギングヘルプ
		"help.logging": `[::b]JULIロギング[::-]
TomcatはデフォルトでJava Util Logging (JULI) を使用します。

[green]設定ファイル:[-]
conf/logging.properties

[green]Handlerタイプ:[-]
• FileHandler: ローテーティングログファイルに書き込み
• ConsoleHandler: コンソール/stdoutに書き込み

[green]ログレベル (昇順):[-]
• FINEST - 最も詳細
• FINER
• FINE
• CONFIG
• INFO
• WARNING
• SEVERE - エラーのみ

[green]一般的なロガー:[-]
• org.apache.catalina - Tomcatコア
• org.apache.coyote - コネクタ
• org.apache.jasper - JSPエンジン`,

		// Contextヘルプ
		"help.context": `[::b]Context設定[::-]
Webアプリケーションとその設定を定義します。

[green]主な属性:[-]
• [yellow]path[-]: URLパス (例: /myapp)
• [yellow]docBase[-]: アプリケーションディレクトリまたはWAR
• [yellow]reloadable[-]: クラス変更時の自動リロード
• [yellow]crossContext[-]: アプリ間ディスパッチを許可

[green]場所:[-]
• $CATALINA_BASE/conf/context.xml - グローバルデフォルト
• $CATALINA_BASE/conf/[engine]/[host]/ - アプリごと
• META-INF/context.xml - アプリケーションに埋め込み

[yellow]本番環境設定:[-]
• reloadable="false" (パフォーマンス)
• privileged="false" (セキュリティ)`,

		"help.context.reloadable": `[::b]リロード可能[::-]
クラス変更時に自動的にリロードします。

[yellow]デフォルト:[-] false

[green]動作:[-]
• /WEB-INF/classesと/WEB-INF/libを監視
• 変更検出時にアプリケーションをリロード
• 開発中に便利

[yellow]本番環境推奨事項:[-]
以下のために"false"に設定:
• パフォーマンス向上 (ファイル監視なし)
• 安定性 (予期しないリロードなし)
• メモリ効率`,

		// Web.xmlヘルプ
		"help.webxml.servlet": `[::b]サーブレット設定[::-]
サーブレットとそのマッピングを定義します。

[green]要素:[-]
• [yellow]servlet-name[-]: 一意の識別子
• [yellow]servlet-class[-]: 完全修飾クラス名
• [yellow]load-on-startup[-]: 初期化順序 (オプション)
• [yellow]init-param[-]: 初期化パラメータ

[green]マッピング:[-]
• [yellow]url-pattern[-]: マッチするURLパターン
• パターン: /exact, /path/*, *.extension

[green]ロード順序:[-]
• 負またはなし: 最初のリクエスト時にロード
• 0または正: 起動時にロード (小さいほど早い)`,

		"help.webxml.filter": `[::b]フィルター設定[::-]
フィルターはサーブレットの前にリクエストを処理します。

[green]一般的な用途:[-]
• 認証/認可
• リクエスト/レスポンス変更
• ログと監査
• 圧縮
• 文字エンコーディング

[green]フィルターチェーン:[-]
• 定義された順序でフィルターを実行
• 各フィルターがchain.doFilter()を呼び出し
• サーブレット前にリクエストを変更可能
• サーブレット後にレスポンスを変更可能

[green]マッピングタイプ:[-]
• url-pattern: URLパターンにマッチ
• servlet-name: 特定のサーブレットに適用
• dispatcher: REQUEST, FORWARD, INCLUDE, ERROR`,

		"help.webxml.session": `[::b]セッション設定[::-]
HTTPセッションの動作を構成します。

[green]session-timeout:[-]
• セッション有効期限までの時間 (分)
• -1 = セッションはタイムアウトしない
• デフォルト: 30分

[green]cookie-config:[-]
• name: クッキー名 (デフォルト: JSESSIONID)
• http-only: JavaScriptアクセス防止 (推奨)
• secure: HTTPS経由でのみ送信
• max-age: クッキー有効期間 (秒)

[green]tracking-mode:[-]
• COOKIE - クッキー使用 (推奨)
• URL - URL書き換え (セキュリティリスク)
• SSL - SSLセッションID`,

		"help.webxml.security": `[::b]セキュリティ制約[::-]
保護されたリソースとアクセスルールを定義します。

[green]コンポーネント:[-]
• [yellow]web-resource-collection[-]: 保護するURL
• [yellow]auth-constraint[-]: 必要なロール
• [yellow]user-data-constraint[-]: トランスポート保証

[green]トランスポート保証:[-]
• NONE - 暗号化不要
• INTEGRAL - データ整合性 (HTTPS)
• CONFIDENTIAL - 機密性 (HTTPS)

[green]例:[-]
<security-constraint>
  <web-resource-collection>
    <url-pattern>/admin/*</url-pattern>
  </web-resource-collection>
  <auth-constraint>
    <role-name>admin</role-name>
  </auth-constraint>
</security-constraint>`,

		// DataSource プロパティヘルプ
		"help.ds.name": `[yellow]JNDI名[white]

このDataSourceを参照するためのJNDI名です。

[aqua]例:[white]
  jdbc/MyDatabase

[aqua]コードでの使用:[white]
  Context ctx = new InitialContext();
  DataSource ds = (DataSource)
    ctx.lookup("java:comp/env/jdbc/MyDatabase");`,

		"help.ds.auth": `[yellow]認証[white]

認証を管理する主体を指定します。

[aqua]Container:[white]
  コンテナ(Tomcat)がリソースへの
  サインオンを管理します。資格情報は
  DataSource設定に保存されます。

[aqua]Application:[white]
  アプリケーションが接続時に資格情報を
  プログラムで提供します。`,

		"help.ds.factory": `[yellow]Factoryクラス[white]

JNDIオブジェクトファクトリクラスです。

[aqua]デフォルト:[white]
  org.apache.tomcat.dbcp.dbcp2.
    BasicDataSourceFactory

[aqua]他のオプション:[white]
  • org.apache.commons.dbcp2.BasicDataSourceFactory
  • com.zaxxer.hikari.HikariJNDIFactory`,

		"help.ds.driver": `[yellow]JDBCドライバクラス[white]

JDBCドライバの完全修飾クラス名です。

[aqua]一般的なドライバ:[white]
  MySQL 8.x: com.mysql.cj.jdbc.Driver
  PostgreSQL: org.postgresql.Driver
  Oracle: oracle.jdbc.OracleDriver
  SQL Server: com.microsoft.sqlserver.
    jdbc.SQLServerDriver
  MariaDB: org.mariadb.jdbc.Driver`,

		"help.ds.url": `[yellow]JDBC URL[white]

データベースへの接続URLです。

[aqua]フォーマット例:[white]
  MySQL:
    jdbc:mysql://host:3306/dbname
  PostgreSQL:
    jdbc:postgresql://host:5432/dbname
  Oracle:
    jdbc:oracle:thin:@host:1521:SID
  SQL Server:
    jdbc:sqlserver://host:1433;
      databaseName=dbname`,

		"help.ds.username": `[yellow]データベースユーザー名[white]

データベース認証用のユーザー名です。

このユーザーはアプリケーションの
データベース操作に適切な権限が必要です。`,

		"help.ds.password": `[yellow]データベースパスワード[white]

データベース認証用のパスワードです。

[red]セキュリティ注意:[white]
本番環境では暗号化されたパスワードや
外部シークレット管理の使用を検討してください。`,

		"help.ds.initialsize": `[yellow]初期プールサイズ[white]

プール初期化時に作成される接続数です。

[aqua]デフォルト:[white] 0
[aqua]推奨:[white] 5-10

アプリケーションの基本的な接続ニーズに
基づいて設定してください。`,

		"help.ds.maxtotal": `[yellow]最大総接続数[white]

プール内のアクティブ接続の最大数です。

[aqua]デフォルト:[white] 8
[aqua]推奨:[white] 20-100

データベースのmax_connections設定と
アプリインスタンス数を考慮してください。`,

		"help.ds.maxidle": `[yellow]最大アイドル接続数[white]

プール内でアイドル状態を維持できる
最大接続数です。

[aqua]デフォルト:[white] 8
[aqua]推奨:[white] 初期サイズと同じ

高い値は接続を準備状態に保ちますが、
より多くのDBリソースを消費します。`,

		"help.ds.minidle": `[yellow]最小アイドル接続数[white]

プール内でアイドル状態を維持する
最小接続数です。

[aqua]デフォルト:[white] 0
[aqua]推奨:[white] 5-10

突発的なトラフィックに対する
迅速な応答を保証します。`,

		"help.ds.maxwait": `[yellow]最大待機時間[white]

プールから接続を取得するための
最大待機時間(ミリ秒)です。

[aqua]デフォルト:[white] -1 (無制限)
[aqua]推奨:[white] 10000-30000

スレッドの無限ブロックを防ぐため
タイムアウトを設定してください。`,

		"help.ds.validationquery": `[yellow]検証クエリ[white]

使用前に接続を検証するSQLクエリです。

[aqua]データベース別の例:[white]
  MySQL/MariaDB: SELECT 1
  PostgreSQL: SELECT 1
  Oracle: SELECT 1 FROM DUAL
  SQL Server: SELECT 1
  H2/HSQLDB: SELECT 1`,

		"help.ds.testonborrow": `[yellow]借用時テスト[white]

アプリケーションに接続を貸し出す前に
接続を検証します。

[aqua]デフォルト:[white] false
[aqua]推奨:[white] true

有効な接続を保証しますが、
若干のオーバーヘッドが追加されます。`,

		"help.ds.testwhileidle": `[yellow]アイドル時テスト[white]

バックグラウンドでアイドル接続を
定期的に検証します。

[aqua]デフォルト:[white] false
[aqua]推奨:[white] true

リクエスト遅延に影響を与えずに
古い接続を事前に削除します。`,

		// Mail Session プロパティヘルプ
		"help.mail.name": `[yellow]JNDI名[white]

このMail Sessionを参照するためのJNDI名です。

[aqua]例:[white]
  mail/Session

[aqua]コードでの使用:[white]
  Session session = (Session)
    ctx.lookup("java:comp/env/mail/Session");`,

		"help.mail.auth": `[yellow]認証[white]

認証を管理する主体を指定します。

[aqua]Container:[white]
  TomcatがSMTP認証を管理します。

[aqua]Application:[white]
  アプリケーションが資格情報を提供します。`,

		"help.mail.host": `[yellow]SMTPホスト[white]

SMTPサーバーのホスト名またはIPアドレスです。

[aqua]例:[white]
  • smtp.gmail.com
  • smtp.office365.com
  • localhost`,

		"help.mail.port": `[yellow]SMTPポート[white]

SMTPサーバーのポート番号です。

[aqua]一般的なポート:[white]
  • 25: 標準SMTP (非暗号化)
  • 465: SMTPS (SSL/TLS)
  • 587: Submission (STARTTLS)`,

		"help.mail.user": `[yellow]SMTPユーザー[white]

SMTP認証用のユーザー名です。

通常はメールアドレスまたはアカウント名です。`,

		"help.mail.protocol": `[yellow]プロトコル[white]

メール転送プロトコルです。

[aqua]smtp:[white]
  標準SMTP、オプションでSTARTTLS使用

[aqua]smtps:[white]
  SSL/TLS経由のSMTP (暗黙的)`,

		"help.mail.smtpauth": `[yellow]SMTP認証[white]

SMTP認証を有効にします。

[aqua]デフォルト:[white] false

SMTPサーバーがユーザー名/パスワード
認証を要求する場合はtrueに設定します。`,

		"help.mail.starttls": `[yellow]StartTLS[white]

STARTTLS暗号化を有効にします。

[aqua]デフォルト:[white] false

初期プレーンテキストハンドシェイク後に
接続をTLSにアップグレードします。
多くの最新SMTPサーバーで必要です。`,

		"help.mail.debug": `[yellow]デバッグモード[white]

JavaMailデバッグ出力を有効にします。

[aqua]デフォルト:[white] false

トラブルシューティング用に詳細な
プロトコル情報をSystem.outに出力します。`,

		// Environment プロパティヘルプ
		"help.env.name": `[yellow]JNDI名[white]

この環境エントリのJNDI名です。

[aqua]例:[white]
  myapp/config/maxItems

[aqua]コードでの使用:[white]
  Integer max = (Integer)
    ctx.lookup("java:comp/env/myapp/config/maxItems");`,

		"help.env.value": `[yellow]値[white]

この環境エントリの値です。

ルックアップ時に指定された型に変換されます。`,

		"help.env.type": `[yellow]型[white]

この環境エントリのJava型です。

[aqua]一般的な型:[white]
  • java.lang.String
  • java.lang.Integer
  • java.lang.Boolean
  • java.lang.Double`,

		"help.env.override": `[yellow]オーバーライド[white]

アプリケーションがこの値を上書きできます。

[aqua]true:[white]
  web.xmlで上書き可能

[aqua]false:[white]
  サーバー設定で値が固定`,

		"help.env.description": `[yellow]説明[white]

このエントリのオプション説明です。

この設定値の目的と使用法を文書化します。`,

		// ResourceLink プロパティヘルプ
		"help.reslink.name": `[yellow]ローカル名[white]

Webアプリケーションで使用するJNDI名です。

[aqua]例:[white]
  jdbc/LocalDB

アプリケーションがリソースを参照する
ために使用する名前です。`,

		"help.reslink.global": `[yellow]グローバル名[white]

server.xmlのグローバルリソース名です。

[aqua]例:[white]
  jdbc/GlobalDB

<GlobalNamingResources>で定義された
リソースにリンクします。`,

		"help.reslink.type": `[yellow]リソース型[white]

リンクされたリソースのJava型です。

[aqua]一般的な型:[white]
  • javax.sql.DataSource
  • javax.mail.Session
  • org.apache.catalina.UserDatabase`,

		// Connector プロパティヘルプ
		"help.conn.port": `[yellow]ポート[white]

このコネクタがリクエストを受け付ける
TCPポート番号です。

[aqua]一般的なポート:[white]
  • 8080: HTTP (開発)
  • 80: HTTP (本番)
  • 8443/443: HTTPS
  • 8009: AJP`,

		"help.conn.protocol": `[yellow]プロトコル[white]

プロトコルハンドラの実装です。

[aqua]HTTP:[white]
  • HTTP/1.1 (NIO/APR自動検出)
  • org.apache.coyote.http11.Http11NioProtocol

[aqua]AJP:[white]
  • AJP/1.3
  • org.apache.coyote.ajp.AjpNioProtocol`,

		"help.conn.timeout": `[yellow]接続タイムアウト[white]

接続後、最初のリクエストデータを
待機する時間(ミリ秒)です。

[aqua]デフォルト:[white] 60000 (60秒)
[aqua]推奨:[white] 20000-60000

低い値は遅いクライアントからの
リソースをより早く解放します。`,

		"help.conn.redirectport": `[yellow]リダイレクトポート[white]

SSLが必要な場合にリダイレクトするポートです。

[aqua]デフォルト:[white] 8443

リクエストがセキュリティを要求するが
非SSLコネクタに到着した場合に使用されます。`,

		"help.conn.maxthreads": `[yellow]最大スレッド[white]

リクエスト処理用の最大スレッド数です。

[aqua]デフォルト:[white] 200
[aqua]推奨:[white] 200-800

高い値はより多くの同時リクエストを
処理しますが、より多くのメモリを使用します。`,

		"help.conn.minsparethreads": `[yellow]最小スペアスレッド[white]

継続的に実行される最小スレッド数です。

[aqua]デフォルト:[white] 10
[aqua]推奨:[white] 25-50

高い値は突発的なトラフィック増加への
応答時間を改善します。`,

		"help.conn.acceptcount": `[yellow]受付キュー[white]

全スレッドが使用中の場合の
着信接続の最大キュー長です。

[aqua]デフォルト:[white] 100

これを超える接続は拒否されます。`,

		"help.conn.executor": `[yellow]Executor[white]

共有スレッドプールExecutorの名前です。

空欄にするとコネクタ専用のスレッドプールを
使用し、Serviceで定義されたExecutor名を
指定できます。`,

		"help.conn.sslenabled": `[yellow]SSL有効[white]

このコネクタでSSL/TLSを有効にします。

キーストア設定が必要です。`,

		"help.conn.scheme": `[yellow]スキーム[white]

リクエストURLのプロトコルスキームです。

[aqua]値:[white]
  • http: 非セキュア接続
  • https: セキュアSSL/TLS接続`,

		"help.conn.secure": `[yellow]セキュア[white]

リクエストをセキュアとしてマークします。

SSL/TLSコネクタでtrueに設定すると
request.isSecure()がtrueを返します。`,

		"help.conn.keystorefile": `[yellow]キーストアファイル[white]

SSLキーストアファイルのパスです。

[aqua]例:[white]
  conf/localhost-rsa.jks
  ${catalina.base}/conf/keystore.p12`,

		"help.conn.keystorepass": `[yellow]キーストアパスワード[white]

キーストアファイルのパスワードです。

[red]セキュリティ:[white]
本番環境では外部シークレット管理の
使用を検討してください。`,

		"help.conn.keystoretype": `[yellow]キーストアタイプ[white]

キーストアファイルのタイプです。

[aqua]タイプ:[white]
  • JKS: Java KeyStore (レガシー)
  • PKCS12: 現代標準 (推奨)`,

		"help.conn.sslprotocol": `[yellow]SSLプロトコル[white]

SSL/TLSプロトコルバージョンです。

[aqua]推奨:[white] TLS
[aqua]特定バージョン:[white] TLSv1.2, TLSv1.3

SSLv3とTLSv1.0/1.1は避けてください。`,

		"help.conn.clientauth": `[yellow]クライアント認証[white]

クライアント証明書認証モードです。

[aqua]false:[white] クライアント証明書不要
[aqua]want:[white] 要求するが必須ではない
[aqua]true:[white] クライアント証明書必須`,

		"help.conn.secret": `[yellow]AJPシークレット[white]

AJP認証用の共有シークレットです。

secretRequiredがtrueの場合に必要です。
Webサーバー設定と一致する必要があります。`,

		"help.conn.secretrequired": `[yellow]シークレット必須[white]

接続にAJPシークレットを要求します。

[aqua]デフォルト:[white] true (Tomcat 9.0.31+)

信頼できるネットワークでのみ
falseに設定してください。`,

		"help.conn.address": `[yellow]アドレス[white]

このコネクタをバインドするIPアドレスです。

空欄にすると全インターフェースにバインドします。

[aqua]例:[white]
  • 127.0.0.1 (localhostのみ)
  • 0.0.0.0 (全インターフェース)`,

		// Executor プロパティヘルプ
		"help.exec.name": `[yellow]Executor名[white]

このスレッドプールの一意な名前です。

コネクタがこの共有Executorを
参照するために使用します。

[aqua]例:[white]
  tomcatThreadPool`,

		"help.exec.classname": `[yellow]クラス名[white]

Executor実装クラスです。

[aqua]標準:[white]
  org.apache.catalina.core.
    StandardThreadExecutor

[aqua]仮想スレッド (Java 21+):[white]
  org.apache.catalina.core.
    StandardVirtualThreadExecutor`,

		"help.exec.nameprefix": `[yellow]名前プレフィックス[white]

スレッド名のプレフィックスです。

ログやスレッドダンプでスレッドを
識別するのに便利です。

[aqua]例:[white]
  catalina-exec-`,

		"help.exec.maxthreads": `[yellow]最大スレッド[white]

プール内の最大スレッド数です。

[aqua]デフォルト:[white] 200
[aqua]推奨:[white] 200-800`,

		"help.exec.minsparethreads": `[yellow]最小スペアスレッド[white]

維持される最小アイドルスレッド数です。

[aqua]デフォルト:[white] 25`,

		"help.exec.maxidletime": `[yellow]最大アイドル時間[white]

アイドルスレッドが終了するまでの
待機時間(ミリ秒)です。

[aqua]デフォルト:[white] 60000 (1分)`,

		"help.exec.maxqueuesize": `[yellow]最大キューサイズ[white]

キュー内の最大保留リクエスト数です。

[aqua]デフォルト:[white] Integer.MAX_VALUE

低い値は過負荷時により早く
リクエストを拒否します。`,

		"help.exec.prestartminsparethreads": `[yellow]最小スレッド事前起動[white]

起動時に最小スレッドを開始します。

[aqua]デフォルト:[white] false

コールドスタート遅延を避けるには
trueに設定してください。`,

		"help.server.address": `[yellow]シャットダウンアドレス[white]

シャットダウンリスナーのアドレスです。

[aqua]デフォルト:[white] localhost

外部インターフェースには絶対にバインドしないでください。`,

		// Listener プロパティヘルプ
		"help.listener.classname": `[yellow]クラス名[white]

リスナーの完全修飾クラス名です。

[aqua]一般的なリスナー:[white]
  • VersionLoggerListener
  • AprLifecycleListener
  • JreMemoryLeakPreventionListener
  • ThreadLocalLeakPreventionListener`,

		"help.listener.sslengine": `[yellow]SSLエンジン[white]

APR/ネイティブコネクタ用SSLエンジンです。

[aqua]値:[white]
  • on: OpenSSLエンジン有効
  • off: JSSE使用`,

		// Service/Engine プロパティヘルプ
		"help.service.name": `[yellow]Service名[white]

このサービスコンテナの名前です。

[aqua]デフォルト:[white] Catalina

1つのサーバーに異なる設定用の
複数のサービスが存在できます。`,

		"help.engine.name": `[yellow]Engine名[white]

このCatalinaエンジンの名前です。

[aqua]デフォルト:[white] Catalina

ロギングとJMXで使用されます。`,

		"help.engine.defaulthost": `[yellow]デフォルトホスト[white]

一致しないリクエストに使用するホストです。

[aqua]デフォルト:[white] localhost

設定された<Host>名と一致する必要があります。`,

		"help.engine.jvmroute": `[yellow]JVMルート[white]

セッションアフィニティ用のルートIDです。

ロードバランサーがこのTomcatインスタンスを
識別するために使用します。

[aqua]例:[white] node1`,

		// Host プロパティヘルプ
		"help.host.name": `[yellow]ホスト名[white]

仮想ホスト名(ドメイン)です。

[aqua]例:[white]
  • localhost
  • www.example.com
  • *.example.com (ワイルドカード)`,

		"help.host.unpackwars": `[yellow]WAR展開[white]

実行前にWARファイルを展開します。

[aqua]デフォルト:[white] true

展開されたアプリはより速く起動します。`,

		"help.host.deployonstart": `[yellow]起動時デプロイ[white]

Tomcat起動時にアプリケーションをデプロイします。

[aqua]デフォルト:[white] true`,

		// Context プロパティヘルプ
		"help.context.path": `[yellow]コンテキストパス[white]

このアプリケーションのURLパスです。

[aqua]例:[white]
  • "" (ROOTアプリケーション)
  • /myapp
  • /api/v1`,

		"help.context.docbase": `[yellow]ドキュメントベース[white]

アプリケーションファイルへのパスです。

WARファイルまたはディレクトリ。
HostのappBaseからの相対パス。`,

		"help.context.crosscontext": `[yellow]クロスコンテキスト[white]

他のコンテキストへのアクセスを許可します。

[aqua]デフォルト:[white] false

getContext()呼び出しを有効にします。`,

		"help.context.cookies": `[yellow]クッキー[white]

セッション追跡にクッキーを使用します。

[aqua]デフォルト:[white] true`,

		"help.context.privileged": `[yellow]特権[white]

Tomcat内部クラスにアクセスします。

[aqua]デフォルト:[white] false

マネージャーアプリに必要です。`,

		// Valve プロパティヘルプ
		"help.valve.classname": `[yellow]Valveクラス名[white]

Valve実装クラスです。

[aqua]一般的なValve:[white]
  • AccessLogValve
  • RemoteAddrValve
  • RemoteIpValve
  • ErrorReportValve
  • StuckThreadDetectionValve`,

		"help.valve.accesslog.pattern": `[yellow]ログパターン[white]

アクセスログの形式パターンです。

[aqua]一般的なパターン:[white]
  %h: リモートホスト
  %l: リモートユーザー (identd)
  %u: 認証済みユーザー
  %t: 日時
  %r: リクエスト行
  %s: ステータスコード
  %b: 送信バイト数

[aqua]Combined形式:[white]
  %h %l %u %t "%r" %s %b`,

		"help.valve.accesslog.directory": `[yellow]ディレクトリ[white]

ログファイルのディレクトリです。

[aqua]デフォルト:[white] logs

CATALINA_BASEからの相対パス。`,

		"help.valve.accesslog.prefix": `[yellow]プレフィックス[white]

ログファイル名のプレフィックスです。

[aqua]デフォルト:[white] localhost_access_log`,

		"help.valve.accesslog.suffix": `[yellow]サフィックス[white]

ログファイル名のサフィックスです。

[aqua]デフォルト:[white] .txt`,

		"help.valve.accesslog.rotate": `[yellow]ローテーション[white]

日次ログローテーションを有効にします。

[aqua]デフォルト:[white] true`,

		"help.valve.remoteaddr.allow": `[yellow]許可パターン[white]

許可するIPアドレスの正規表現です。

[aqua]例:[white]
  • 127\\.0\\.0\\.1
  • 192\\.168\\.\\d+\\.\\d+
  • 10\\.\\d+\\.\\d+\\.\\d+`,

		"help.valve.remoteaddr.deny": `[yellow]拒否パターン[white]

拒否するIPアドレスの正規表現です。

許可パターンの後にチェックされます。`,

		"help.valve.stuckthread.threshold": `[yellow]しきい値[white]

スレッドがスタックとみなされる時間(秒)です。

[aqua]デフォルト:[white] 600 (10分)

超過時に警告がログされます。`,

		// Realm プロパティヘルプ
		"help.realm.classname": `[yellow]Realmクラス名[white]

認証Realm実装です。

[aqua]一般的なRealm:[white]
  • UserDatabaseRealm (ファイルベース)
  • DataSourceRealm (JDBC)
  • JNDIRealm (LDAP)
  • JAASRealm (JAAS)
  • CombinedRealm (複数)`,

		"help.realm.userdatabase.resource": `[yellow]リソース名[white]

UserDatabaseリソースのJNDI名です。

[aqua]デフォルト:[white] UserDatabase

GlobalNamingResourcesで定義されます。`,

		"help.realm.datasource.name": `[yellow]DataSource名[white]

DataSourceのJNDI名です。

[aqua]例:[white]
  jdbc/AuthDB`,

		"help.realm.datasource.usertable": `[yellow]ユーザーテーブル[white]

ユーザー資格情報を含むテーブルです。

[aqua]デフォルト:[white] users`,

		"help.realm.datasource.usernameCol": `[yellow]ユーザー名カラム[white]

ユーザー名のカラム名です。

[aqua]デフォルト:[white] user_name`,

		"help.realm.datasource.passwordCol": `[yellow]パスワードカラム[white]

パスワードのカラム名です。

[aqua]デフォルト:[white] user_pass`,

		"help.realm.datasource.roletable": `[yellow]ロールテーブル[white]

ユーザーロールを含むテーブルです。

[aqua]デフォルト:[white] user_roles`,

		"help.realm.datasource.rolenameCol": `[yellow]ロール名カラム[white]

ロール名のカラム名です。

[aqua]デフォルト:[white] role_name`,

		"help.realm.jndi.connectionURL": `[yellow]接続URL[white]

LDAPサーバーURLです。

[aqua]例:[white]
  ldap://ldap.example.com:389`,

		"help.realm.jndi.userbase": `[yellow]ユーザーベース[white]

ユーザー検索のベースDNです。

[aqua]例:[white]
  ou=users,dc=example,dc=com`,

		"help.realm.jndi.userpattern": `[yellow]ユーザーパターン[white]

直接ユーザー検索用のDNパターンです。

[aqua]例:[white]
  uid={0},ou=users,dc=example,dc=com`,

		"help.realm.jndi.rolebase": `[yellow]ロールベース[white]

ロール検索のベースDNです。

[aqua]例:[white]
  ou=groups,dc=example,dc=com`,

		// User/Role プロパティヘルプ
		"help.user.username": `[yellow]ユーザー名[white]

ユーザーの一意な識別子です。

認証に使用され、ロール割り当てで
参照されます。`,

		"help.user.password": `[yellow]パスワード[white]

ユーザーのパスワードです。

[red]セキュリティ:[white]
本番環境ではハッシュ化されたパスワードを
使用してください。CredentialHandler設定を
参照してください。`,

		"help.user.roles": `[yellow]ロール[white]

カンマ区切りのロールリストです。

[aqua]一般的なロール:[white]
  • manager-gui
  • admin-gui
  • manager-script`,

		"help.role.name": `[yellow]ロール名[white]

ロールの一意な識別子です。

web.xmlセキュリティ制約と
ユーザー割り当てで参照されます。`,

		"help.role.description": `[yellow]説明[white]

オプションのロール説明です。

このロールの目的と権限を文書化します。`,

		"help.default": `[gray]フィールドを選択するとヘルプ情報が表示されます。[-]`,
	},
}
