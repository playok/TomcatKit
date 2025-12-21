# CLAUDE.md

This file provides guidance for Claude Code (claude.ai/claude-code) when working with this repository.

## Project Overview

TomcatKit is a CLI-based TUI (Text User Interface) utility for managing Apache Tomcat 9.0 configuration files. It provides an interactive ncurses-style interface for editing server.xml, tomcat-users.xml, and other Tomcat configuration files.

## Build Commands

```bash
# Build the application
make build
# or
go build -o bin/tomcatkit ./cmd/tomcatkit

# Run the application
./bin/tomcatkit

# Run with specific Tomcat path
./bin/tomcatkit -home /path/to/tomcat

# Tidy dependencies
go mod tidy
```

## Architecture

### Design Principles

1. **Separation of Concerns**: Configuration logic is separate from TUI presentation
2. **Type Safety**: All Tomcat XML elements are represented as Go structs with proper XML tags
3. **Safe Editing**: Always create backups before modifying configuration files
4. **Auto-detection**: Support multiple methods to locate Tomcat installations

### Package Structure

```
internal/
├── config/           # Configuration types and services
│   ├── tomcat.go     # TomcatInstance struct (CATALINA_HOME, CATALINA_BASE, version)
│   ├── settings.go   # Application settings persistence
│   ├── server/       # server.xml types and operations
│   ├── connector/    # Connector protocol constants and defaults
│   ├── realm/        # Realm types and tomcat-users.xml operations
│   ├── jndi/         # JNDI resource types and context.xml operations
│   └── logging/      # Logging configuration types
├── detector/         # Tomcat installation auto-detection
├── parser/           # XML parsing utilities
└── tui/              # Terminal UI
    ├── app.go        # Main application, navigation, instance selection
    └── views/        # Individual configuration views
        ├── server.go     # Server configuration view
        ├── connector.go  # Connector configuration view
        ├── security.go   # Security/Realm view
        ├── jndi.go       # JNDI Resources view
        ├── host.go       # Virtual Hosts & Contexts view
        └── valve.go      # Valves configuration view
```

### Key Types

#### TomcatInstance (`internal/config/tomcat.go`)
```go
type TomcatInstance struct {
    CatalinaHome string  // Tomcat installation directory
    CatalinaBase string  // Instance directory (can differ from Home)
    Version      string  // Detected Tomcat version
}
```

#### Server Configuration (`internal/config/server/server.go`)
- `Server`: Root element with shutdown port and command
- `Service`: Contains connectors and engine
- `Connector`: HTTP/AJP/SSL connector configuration
- `Engine`: Servlet engine with hosts
- `Host`: Virtual host configuration
- `Context`: Web application context
- `Realm`: Authentication realm configuration
- `Valve`: Request processing valves

#### TUI Application (`internal/tui/app.go`)
- Uses `tview` library for terminal UI
- Page-based navigation with main menu
- Instance selector for Tomcat path management
- Each configuration module has its own view

### Module Implementation Pattern

Each configuration module follows this pattern:

1. **Types file** (`internal/config/<module>/<module>.go`):
   - Define structs with XML tags for marshaling/unmarshaling
   - Define constants for valid values
   - Define factory functions for defaults

2. **Service file** (`internal/config/<module>/service.go`):
   - `Load()` - Read and parse XML file
   - `Save()` - Marshal and write XML file with backup
   - CRUD operations for nested elements

3. **View file** (`internal/tui/views/<module>.go`):
   - Create `<Module>View` struct with tview components
   - Implement list views for collections
   - Implement forms for editing individual items
   - Handle navigation with Escape key

### XML Handling

- Use Go's `encoding/xml` package
- Struct tags define XML element/attribute mapping
- Handle nested elements (e.g., Realm can contain nested Realms)
- Preserve original structure when saving

Example:
```go
type Connector struct {
    Port              int    `xml:"port,attr"`
    Protocol          string `xml:"protocol,attr,omitempty"`
    ConnectionTimeout int    `xml:"connectionTimeout,attr,omitempty"`
    // ...
}
```

### Settings Persistence

Settings stored in `~/.config/tomcatkit/settings.json`:
```json
{
  "last_instance": {
    "catalina_home": "/opt/tomcat",
    "catalina_base": "/opt/tomcat",
    "version": "9.0.x"
  },
  "recent_paths": ["/opt/tomcat", "/var/tomcat/instance1"]
}
```

## Tomcat Configuration Reference

### server.xml Structure
```xml
<Server port="8005" shutdown="SHUTDOWN">
  <Listener className="..." />
  <GlobalNamingResources>
    <Resource name="UserDatabase" ... />
  </GlobalNamingResources>
  <Service name="Catalina">
    <Connector port="8080" protocol="HTTP/1.1" ... />
    <Connector port="8443" protocol="HTTP/1.1" SSLEnabled="true" ... />
    <Connector port="8009" protocol="AJP/1.3" ... />
    <Engine name="Catalina" defaultHost="localhost">
      <Realm className="org.apache.catalina.realm.LockOutRealm">
        <Realm className="org.apache.catalina.realm.UserDatabaseRealm" ... />
      </Realm>
      <Host name="localhost" appBase="webapps" ...>
        <Valve className="org.apache.catalina.valves.AccessLogValve" ... />
      </Host>
    </Engine>
  </Service>
</Server>
```

### tomcat-users.xml Structure
```xml
<tomcat-users>
  <role rolename="manager-gui"/>
  <role rolename="admin-gui"/>
  <user username="admin" password="secret" roles="manager-gui,admin-gui"/>
</tomcat-users>
```

### Supported Realm Types
- `UserDatabaseRealm`: File-based (tomcat-users.xml)
- `DataSourceRealm`: JDBC database authentication
- `JNDIRealm`: LDAP/Directory server
- `JAASRealm`: Java Authentication and Authorization Service
- `CombinedRealm`: Multiple realms with fallback
- `LockOutRealm`: Brute-force protection wrapper

### Connector Protocols
- HTTP: `HTTP/1.1`, `org.apache.coyote.http11.Http11NioProtocol`, `Http11Nio2Protocol`, `Http11AprProtocol`
- AJP: `AJP/1.3`, `org.apache.coyote.ajp.AjpNioProtocol`, `AjpNio2Protocol`, `AjpAprProtocol`

## Implementation Status

### Completed
- [x] Project structure and build system
- [x] Tomcat instance detection and selection
- [x] Settings persistence
- [x] Server configuration (server.xml core elements)
- [x] Connector configuration (HTTP, AJP, SSL, Executors)
- [x] Security/Realm configuration
- [x] tomcat-users.xml management (Users & Roles)
- [x] JNDI Resources (DataSource, Mail Session, Environment, Resource Links)
- [x] Virtual Hosts & Contexts (Host, Context, Parameters, Session Manager)
- [x] Valves (AccessLog, RemoteAddr, RemoteIp, ErrorReport, SSO, StuckThread, etc.)

### Planned
- [ ] Clustering configuration
- [ ] Logging configuration (JULI)
- [ ] Import/Export configuration
- [ ] Configuration validation
- [ ] Diff view for changes

## Common Patterns

### Adding a New Configuration Module

1. Create type definitions in `internal/config/<module>/<module>.go`
2. Create service with Load/Save in `internal/config/<module>/service.go`
3. Create TUI view in `internal/tui/views/<module>.go`
4. Register view in `internal/tui/app.go` menu

### Form Pattern in Views
```go
func (v *ModuleView) showEditForm(item *ItemType) {
    form := tview.NewForm()
    form.AddInputField("Field", item.Value, 40, nil, nil)
    form.AddButton("Save", func() {
        item.Value = form.GetFormItemByLabel("Field").(*tview.InputField).GetText()
        // Save and refresh
    })
    form.AddButton("Cancel", func() {
        v.pages.SwitchToPage("list")
    })
    form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyEscape {
            v.pages.SwitchToPage("list")
            return nil
        }
        return event
    })
}
```

### Navigation Pattern
- Main menu uses `tview.List` with keyboard shortcuts
- Each view manages its own pages via `tview.Pages`
- Escape key returns to previous view/page
- All views call `app.SetRoot()` to switch context
