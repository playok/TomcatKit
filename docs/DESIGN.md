# TomcatKit Design Document

## Overview

TomcatKit is designed as a modular, extensible TUI application for managing Apache Tomcat 9.0 configurations. This document outlines the architectural decisions and design patterns used throughout the project.

## Goals

1. **Comprehensive Coverage**: Support all user-configurable aspects of Tomcat 9.0
2. **Safety First**: Never modify files without creating backups
3. **User-Friendly**: Intuitive navigation with keyboard shortcuts
4. **Maintainable**: Clean separation between configuration logic and UI

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                        CLI Entry                             │
│                    cmd/tomcatkit/main.go                     │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      TUI Application                         │
│                    internal/tui/app.go                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │ Instance    │  │ Main Menu   │  │ Settings    │         │
│  │ Selector    │  │             │  │ Manager     │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                         Views                                │
│                   internal/tui/views/                        │
│  ┌──────────┐ ┌───────────┐ ┌──────────┐ ┌──────────┐      │
│  │ Server   │ │ Connector │ │ Security │ │ ...more  │      │
│  │ View     │ │ View      │ │ View     │ │ views    │      │
│  └──────────┘ └───────────┘ └──────────┘ └──────────┘      │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   Configuration Layer                        │
│                    internal/config/                          │
│  ┌──────────┐ ┌───────────┐ ┌──────────┐ ┌──────────┐      │
│  │ server/  │ │connector/ │ │ realm/   │ │ jndi/    │      │
│  │ Types    │ │ Types     │ │ Types    │ │ Types    │      │
│  │ Service  │ │ Defaults  │ │ Service  │ │          │      │
│  └──────────┘ └───────────┘ └──────────┘ └──────────┘      │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Tomcat Configuration Files                │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │ server.xml   │  │tomcat-users  │  │ context.xml  │      │
│  │              │  │    .xml      │  │              │      │
│  └──────────────┘  └──────────────┘  └──────────────┘      │
└─────────────────────────────────────────────────────────────┘
```

## Layer Responsibilities

### CLI Layer (`cmd/tomcatkit/`)

- Parse command-line flags
- Initialize settings manager
- Create and run TUI application
- Handle version and help output

### TUI Layer (`internal/tui/`)

- Manage application lifecycle
- Handle navigation between views
- Manage Tomcat instance selection
- Coordinate settings persistence

### View Layer (`internal/tui/views/`)

- Present configuration data in interactive forms
- Handle user input and validation
- Coordinate with configuration services
- Provide feedback on operations (success/error)

### Configuration Layer (`internal/config/`)

- Define type-safe structures for XML elements
- Provide services for loading/saving configurations
- Implement business logic for configuration management
- Create backups before modifications

### Detection Layer (`internal/detector/`)

- Auto-detect Tomcat installations
- Query environment variables
- Search common installation paths
- Detect from running processes

## Design Patterns

### Service Pattern

Each configuration module uses a Service struct that encapsulates:
- File path management
- Loading and parsing
- Saving with backup
- CRUD operations

```go
type ConfigService struct {
    catalinaBase string
    filePath     string
    config       *ConfigType
}

func (s *ConfigService) Load() error { ... }
func (s *ConfigService) Save() error { ... }
func (s *ConfigService) Get() *ConfigType { ... }
```

### View Pattern

Each view follows a consistent structure:

```go
type ModuleView struct {
    app      *tview.Application
    pages    *tview.Pages
    onReturn func()

    // Module-specific state
    service  *ServiceType
    instance *config.TomcatInstance
}

func NewModuleView(app *tview.Application, instance *config.TomcatInstance, onReturn func()) *ModuleView { ... }
func (v *ModuleView) GetRoot() tview.Primitive { ... }
```

### Callback Pattern

Views use callbacks for:
- Returning to the main menu (`onReturn`)
- Refreshing lists after edits
- Switching between pages

### Backup Strategy

Before any configuration modification:

1. Create backup directory if not exists: `$CATALINA_BASE/conf/backup/`
2. Copy current file to backup location
3. Proceed with modification
4. On error, backup remains available for manual recovery

## Data Flow

### Loading Configuration

```
User selects menu item
        │
        ▼
View.GetRoot() called
        │
        ▼
Service.Load() reads XML file
        │
        ▼
xml.Unmarshal() parses to structs
        │
        ▼
View displays data in tview components
```

### Saving Configuration

```
User clicks Save button
        │
        ▼
Form values extracted
        │
        ▼
Struct fields updated
        │
        ▼
Service.Save() called
        │
        ▼
createBackup() runs first
        │
        ▼
xml.MarshalIndent() generates XML
        │
        ▼
os.WriteFile() saves to disk
        │
        ▼
View shows success message
```

## XML Mapping Strategy

### Attributes vs Elements

- Simple values → XML attributes (`xml:"name,attr"`)
- Complex nested structures → XML elements
- Optional values → `omitempty` tag

### Example Mapping

```go
type Server struct {
    XMLName  xml.Name  `xml:"Server"`
    Port     int       `xml:"port,attr"`
    Shutdown string    `xml:"shutdown,attr"`
    Services []Service `xml:"Service"`          // Nested element
}
```

### Handling Unknown Attributes

For extensibility, consider using:
```go
ExtraAttrs []xml.Attr `xml:",any,attr"`
```

## Navigation Model

```
┌─────────────────┐
│ Instance Select │ ◄─── First run or "Change Instance"
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Main Menu     │ ◄─── Escape from any view
└────────┬────────┘
         │
    ┌────┴────┬────────┬──────────┐
    ▼         ▼        ▼          ▼
┌───────┐ ┌───────┐ ┌───────┐ ┌───────┐
│Server │ │Connect│ │Security│ │ ...   │
│ View  │ │ View  │ │ View  │ │       │
└───┬───┘ └───┬───┘ └───┬───┘ └───────┘
    │         │        │
    ▼         ▼        ▼
┌───────┐ ┌───────┐ ┌───────┐
│ Edit  │ │ Edit  │ │ Edit  │
│ Forms │ │ Forms │ │ Forms │
└───────┘ └───────┘ └───────┘
```

## Error Handling

### User-Facing Errors

- Display modal dialogs for error messages
- Include actionable information
- Allow dismissal with Enter or Escape

### Logging Strategy

- Errors are displayed to users via TUI
- Future: Add file logging for debugging

## Configuration File Locations

| File | Location | Purpose |
|------|----------|---------|
| server.xml | $CATALINA_BASE/conf/ | Main server configuration |
| tomcat-users.xml | $CATALINA_BASE/conf/ | User/role definitions |
| context.xml | $CATALINA_BASE/conf/ | Default context settings |
| web.xml | $CATALINA_BASE/conf/ | Default servlet configuration |
| logging.properties | $CATALINA_BASE/conf/ | JULI logging configuration |
| settings.json | ~/.config/tomcatkit/ | Application settings |

## Future Considerations

### Configuration Validation

- Validate port numbers are in valid range
- Check for duplicate connector ports
- Validate file paths exist
- Warn about insecure configurations

### Diff View

- Show changes before saving
- Allow selective application of changes
- Highlight modified fields

### Import/Export

- Export configuration as portable package
- Import configurations from other instances
- Template-based configuration generation

### Multi-Instance Management

- Compare configurations across instances
- Sync settings between instances
- Bulk operations

## Module Checklist

When implementing a new module:

- [ ] Define types in `internal/config/<module>/<module>.go`
- [ ] Add XML struct tags for all fields
- [ ] Create service with Load/Save methods
- [ ] Implement backup before save
- [ ] Create view in `internal/tui/views/<module>.go`
- [ ] Add list view for collections
- [ ] Add edit forms for items
- [ ] Handle Escape key navigation
- [ ] Register in main menu (`internal/tui/app.go`)
- [ ] Update CLAUDE.md with new module info
- [ ] Update README.md feature list
