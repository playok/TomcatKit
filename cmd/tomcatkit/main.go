package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/playok/tomcatkit/internal/config"
	"github.com/playok/tomcatkit/internal/tui"
)

var (
	version = "0.1.0"
)

func main() {
	// Define CLI flags
	catalinaHome := flag.String("home", "", "Path to CATALINA_HOME (Tomcat installation directory)")
	catalinaBase := flag.String("base", "", "Path to CATALINA_BASE (defaults to CATALINA_HOME if not specified)")
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help information")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `TomcatKit - Apache Tomcat Configuration Helper v%s

Usage:
  tomcatkit [options]

Options:
  -home string    Path to CATALINA_HOME (Tomcat installation directory)
  -base string    Path to CATALINA_BASE (defaults to CATALINA_HOME)
  -version        Show version information
  -help           Show this help message

Examples:
  tomcatkit                              # Auto-detect or select Tomcat instance
  tomcatkit -home /opt/tomcat            # Specify Tomcat home directory
  tomcatkit -home /opt/tomcat -base /var/tomcat  # Specify both home and base

Environment Variables:
  CATALINA_HOME   Tomcat installation directory
  CATALINA_BASE   Tomcat instance directory (optional)

`, version)
	}

	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Printf("TomcatKit v%s\n", version)
		os.Exit(0)
	}

	// Handle help flag
	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	// Create options for the app
	opts := &tui.AppOptions{}

	// Check CLI flags first, then environment variables
	if *catalinaHome != "" {
		opts.CatalinaHome = *catalinaHome
		if *catalinaBase != "" {
			opts.CatalinaBase = *catalinaBase
		} else {
			opts.CatalinaBase = *catalinaHome
		}
	} else if envHome := os.Getenv("CATALINA_HOME"); envHome != "" {
		opts.CatalinaHome = envHome
		if envBase := os.Getenv("CATALINA_BASE"); envBase != "" {
			opts.CatalinaBase = envBase
		} else {
			opts.CatalinaBase = envHome
		}
	}

	// Load settings manager
	settingsManager := config.NewSettingsManager()
	if err := settingsManager.Load(); err != nil {
		log.Printf("Warning: Failed to load settings: %v", err)
	}
	opts.SettingsManager = settingsManager

	// Create and run app
	app := tui.NewAppWithOptions(opts)
	if err := app.Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
