# TomcatKit

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/playok/TomcatKit)](https://goreportcard.com/report/github.com/playok/TomcatKit)
[![Tomcat](https://img.shields.io/badge/Tomcat-9.0-F8DC75?style=flat&logo=apache-tomcat)](https://tomcat.apache.org/)
[![AI Generated](https://img.shields.io/badge/AI%20Generated-Claude%20Code-blueviolet?style=flat&logo=anthropic)](https://claude.ai/claude-code)

A CLI-based TUI (Text User Interface) helper utility for Apache Tomcat 9.0 configuration management.

## Demo

![TomcatKit Demo](docs/assets/demo.gif)

## Features

- **Interactive TUI**: ncurses-style terminal interface using [tview](https://github.com/rivo/tview)
- **Comprehensive Configuration**: Covers all major Tomcat 9.0 configuration areas
- **Auto-detection**: Automatically detects Tomcat installations from environment variables, common paths, and running processes
- **Safe Editing**: Creates automatic backups before modifying configuration files
- **Multi-instance Support**: Remembers recently used Tomcat instances

## Supported Configuration Modules

| Module | Status | Description |
|--------|--------|-------------|
| Server | Complete | server.xml core settings (Server, Service, Engine, Host) |
| Connector | Complete | HTTP, AJP, SSL/TLS connectors and thread pools |
| Security/Realm | Complete | Authentication realms and tomcat-users.xml management |
| JNDI Resources | Complete | DataSource, Mail Session, Environment entries, Resource Links |
| Virtual Hosts | Complete | Host, Context, Parameters, Session Manager configuration |
| Valves | Complete | AccessLog, RemoteAddr, RemoteIp, ErrorReport, SSO, StuckThread valves |
| Clustering | Complete | Session replication, membership, interceptors, farm deployer |
| Logging | Complete | JULI logging.properties, file handlers, loggers |
| Context | Complete | context.xml settings, resources, cookies, session manager |
| Web | Complete | web.xml servlets, filters, session, security constraints |
| Quick Templates | Complete | Virtual Threads, HTTPS, Connection Pool, Gzip, Security |

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/playok/tomcatkit.git
cd tomcatkit

# Build
make build

# Or using go directly
go build -o bin/tomcatkit ./cmd/tomcatkit
```

### Requirements

- Go 1.21 or later
- Apache Tomcat 9.0 installation

## Usage

### Basic Usage

```bash
# Run with auto-detection
./bin/tomcatkit

# Specify Tomcat paths explicitly
./bin/tomcatkit -home /opt/tomcat -base /var/tomcat/instance1

# Show version
./bin/tomcatkit -version

# Show help
./bin/tomcatkit -help
```

### CLI Options

| Option | Description |
|--------|-------------|
| `-home` | Path to CATALINA_HOME (Tomcat installation directory) |
| `-base` | Path to CATALINA_BASE (instance directory, defaults to CATALINA_HOME) |
| `-version` | Show version information |
| `-help` | Show help message |

### Navigation

- **Arrow Keys**: Navigate menus and lists
- **Enter**: Select item or confirm action
- **Escape**: Go back / Cancel
- **Tab**: Move between form fields
- **Ctrl+C**: Exit application

## Project Structure

```
tomcatkit/
├── cmd/
│   └── tomcatkit/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   ├── tomcat.go         # Tomcat instance configuration
│   │   ├── settings.go       # Application settings persistence
│   │   ├── server/
│   │   │   ├── server.go     # server.xml type definitions
│   │   │   └── service.go    # server.xml operations
│   │   ├── connector/
│   │   │   └── connector.go  # Connector protocols and defaults
│   │   ├── realm/
│   │   │   ├── realm.go      # Realm types and utilities
│   │   │   └── users.go      # tomcat-users.xml operations
│   │   ├── jndi/
│   │   │   ├── resources.go  # JNDI resource types
│   │   │   └── service.go    # context.xml operations
│   │   ├── logging/
│   │   │   ├── logging.go    # Logging configuration types
│   │   │   └── service.go    # logging.properties operations
│   │   └── web/
│   │       ├── types.go      # web.xml type definitions
│   │       └── service.go    # web.xml operations
│   ├── detector/
│   │   └── detector.go       # Tomcat auto-detection
│   ├── parser/
│   │   └── xml.go            # XML parsing utilities
│   └── tui/
│       ├── app.go            # Main TUI application
│       └── views/
│           ├── server.go     # Server configuration view
│           ├── connector.go  # Connector configuration view
│           ├── security.go   # Security/Realm view
│           ├── jndi.go       # JNDI Resources view
│           ├── host.go       # Virtual Hosts & Contexts view
│           ├── valve.go      # Valves configuration view
│           ├── cluster.go    # Clustering configuration view
│           ├── logging.go    # Logging configuration view
│           ├── context.go    # Context configuration view
│           ├── web.go        # Web application configuration view
│           └── quicktemplates.go  # Quick configuration templates
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Configuration Files

TomcatKit manages the following Tomcat configuration files:

- `$CATALINA_BASE/conf/server.xml` - Main server configuration
- `$CATALINA_BASE/conf/tomcat-users.xml` - User and role definitions
- `$CATALINA_BASE/conf/context.xml` - Default context configuration
- `$CATALINA_BASE/conf/web.xml` - Default web application configuration
- `$CATALINA_BASE/conf/logging.properties` - JULI logging configuration

## Settings

Application settings are stored in:
- Linux/macOS: `~/.config/tomcatkit/settings.json`
- Windows: `%APPDATA%\tomcatkit\settings.json`

## About This Project

This is a **hobby project** created for fun and learning purposes. It was built to explore AI-assisted development and to provide a useful tool for Tomcat administrators.

### AI-Generated

This project was entirely created by AI using **[Claude Code](https://claude.ai/claude-code)** (Anthropic's Claude).

- **AI Model**: Claude Opus 4.5 (`claude-opus-4-5-20251101`)
- **Development Tool**: Claude Code CLI
- **Human Role**: Project direction, requirements specification, and review

All code, documentation, and configuration in this repository were generated through AI-assisted development. The AI handled architecture design, implementation, debugging, and documentation while the human provided guidance and validation.

> **Note**: This is a personal hobby project and is not affiliated with or endorsed by the Apache Software Foundation.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

[playok](https://github.com/playok)

---

<p align="center">
  <sub>Built with AI assistance from Claude Code</sub>
</p>
