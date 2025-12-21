package detector

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/playok/tomcatkit/internal/config"
)

// Detector finds Tomcat installations on the system
type Detector struct{}

// NewDetector creates a new Tomcat detector
func NewDetector() *Detector {
	return &Detector{}
}

// DetectAll finds all Tomcat installations
func (d *Detector) DetectAll() ([]*config.TomcatInstance, error) {
	var instances []*config.TomcatInstance

	// Check environment variables first
	if envInstance := d.detectFromEnv(); envInstance != nil {
		instances = append(instances, envInstance)
	}

	// Check common installation paths
	commonPaths := d.getCommonPaths()
	for _, path := range commonPaths {
		if instance := d.detectAtPath(path); instance != nil {
			// Avoid duplicates
			if !d.isDuplicate(instances, instance) {
				instances = append(instances, instance)
			}
		}
	}

	// Check running processes
	runningInstances := d.detectRunningInstances()
	for _, instance := range runningInstances {
		if !d.isDuplicate(instances, instance) {
			instances = append(instances, instance)
		} else {
			// Update running status for existing instance
			for _, existing := range instances {
				if existing.CatalinaHome == instance.CatalinaHome {
					existing.IsRunning = true
					existing.PID = instance.PID
				}
			}
		}
	}

	return instances, nil
}

// detectFromEnv checks CATALINA_HOME and CATALINA_BASE environment variables
func (d *Detector) detectFromEnv() *config.TomcatInstance {
	catalinaHome := os.Getenv("CATALINA_HOME")
	catalinaBase := os.Getenv("CATALINA_BASE")

	if catalinaHome == "" {
		return nil
	}

	if catalinaBase == "" {
		catalinaBase = catalinaHome
	}

	if !d.isValidTomcatDir(catalinaHome) {
		return nil
	}

	return &config.TomcatInstance{
		CatalinaHome: catalinaHome,
		CatalinaBase: catalinaBase,
		Version:      d.detectVersion(catalinaHome),
	}
}

// getCommonPaths returns common Tomcat installation paths
func (d *Detector) getCommonPaths() []string {
	var paths []string

	switch runtime.GOOS {
	case "darwin": // macOS
		paths = []string{
			"/usr/local/Cellar/tomcat",
			"/usr/local/opt/tomcat",
			"/opt/homebrew/Cellar/tomcat",
			"/opt/homebrew/opt/tomcat",
			"/Library/Tomcat",
			filepath.Join(os.Getenv("HOME"), "tomcat"),
			filepath.Join(os.Getenv("HOME"), "apache-tomcat"),
		}
		// Check for versioned Homebrew installations
		homebrewPaths := []string{"/usr/local/Cellar/tomcat", "/opt/homebrew/Cellar/tomcat"}
		for _, hp := range homebrewPaths {
			if entries, err := os.ReadDir(hp); err == nil {
				for _, entry := range entries {
					if entry.IsDir() {
						paths = append(paths, filepath.Join(hp, entry.Name(), "libexec"))
					}
				}
			}
		}
	case "linux":
		paths = []string{
			"/opt/tomcat",
			"/opt/apache-tomcat",
			"/usr/share/tomcat",
			"/usr/share/tomcat9",
			"/var/lib/tomcat",
			"/var/lib/tomcat9",
			filepath.Join(os.Getenv("HOME"), "tomcat"),
			filepath.Join(os.Getenv("HOME"), "apache-tomcat"),
		}
	case "windows":
		paths = []string{
			"C:\\Program Files\\Apache Software Foundation\\Tomcat 9.0",
			"C:\\Program Files (x86)\\Apache Software Foundation\\Tomcat 9.0",
			"C:\\tomcat",
			"C:\\apache-tomcat",
		}
	}

	// Add paths from common patterns
	homeDir := os.Getenv("HOME")
	if homeDir != "" {
		// Check for apache-tomcat-* directories
		if entries, err := os.ReadDir(homeDir); err == nil {
			for _, entry := range entries {
				if entry.IsDir() && strings.HasPrefix(entry.Name(), "apache-tomcat-") {
					paths = append(paths, filepath.Join(homeDir, entry.Name()))
				}
			}
		}
	}

	return paths
}

// detectAtPath checks if a valid Tomcat installation exists at the given path
func (d *Detector) detectAtPath(path string) *config.TomcatInstance {
	if !d.isValidTomcatDir(path) {
		return nil
	}

	return &config.TomcatInstance{
		CatalinaHome: path,
		CatalinaBase: path,
		Version:      d.detectVersion(path),
	}
}

// isValidTomcatDir checks if a directory is a valid Tomcat installation
func (d *Detector) isValidTomcatDir(path string) bool {
	// Check for essential Tomcat files/directories
	requiredPaths := []string{
		filepath.Join(path, "conf", "server.xml"),
		filepath.Join(path, "lib", "catalina.jar"),
	}

	for _, p := range requiredPaths {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// detectVersion reads Tomcat version from catalina.jar or version script
func (d *Detector) detectVersion(catalinaHome string) string {
	// Try running version.sh or version.bat
	var versionCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		versionCmd = exec.Command(filepath.Join(catalinaHome, "bin", "version.bat"))
	} else {
		versionCmd = exec.Command(filepath.Join(catalinaHome, "bin", "version.sh"))
	}

	if output, err := versionCmd.Output(); err == nil {
		// Parse version from output
		re := regexp.MustCompile(`Server version:\s*Apache Tomcat/(\d+\.\d+\.\d+)`)
		if matches := re.FindStringSubmatch(string(output)); len(matches) > 1 {
			return matches[1]
		}
	}

	// Fallback: try to detect from directory name
	baseName := filepath.Base(catalinaHome)
	re := regexp.MustCompile(`(\d+\.\d+\.\d+)`)
	if matches := re.FindStringSubmatch(baseName); len(matches) > 1 {
		return matches[1]
	}

	return "unknown"
}

// detectRunningInstances finds running Tomcat processes
func (d *Detector) detectRunningInstances() []*config.TomcatInstance {
	var instances []*config.TomcatInstance

	// Use ps command to find Java processes with Catalina
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("wmic", "process", "where", "name='java.exe'", "get", "processid,commandline")
	} else {
		cmd = exec.Command("ps", "aux")
	}

	output, err := cmd.Output()
	if err != nil {
		return instances
	}

	lines := strings.Split(string(output), "\n")
	catalinaHomeRe := regexp.MustCompile(`-Dcatalina\.home=([^\s]+)`)
	catalinaBaseRe := regexp.MustCompile(`-Dcatalina\.base=([^\s]+)`)
	pidRe := regexp.MustCompile(`^\S+\s+(\d+)`)

	for _, line := range lines {
		if !strings.Contains(line, "catalina") && !strings.Contains(line, "Catalina") {
			continue
		}

		var catalinaHome, catalinaBase string
		var pid int

		if matches := catalinaHomeRe.FindStringSubmatch(line); len(matches) > 1 {
			catalinaHome = matches[1]
		}
		if matches := catalinaBaseRe.FindStringSubmatch(line); len(matches) > 1 {
			catalinaBase = matches[1]
		}
		if matches := pidRe.FindStringSubmatch(line); len(matches) > 1 {
			// Parse PID
			if p, err := strconv.Atoi(matches[1]); err == nil {
				pid = p
			}
		}

		if catalinaHome != "" {
			if catalinaBase == "" {
				catalinaBase = catalinaHome
			}
			instances = append(instances, &config.TomcatInstance{
				CatalinaHome: catalinaHome,
				CatalinaBase: catalinaBase,
				Version:      d.detectVersion(catalinaHome),
				IsRunning:    true,
				PID:          pid,
			})
		}
	}

	return instances
}

// isDuplicate checks if an instance already exists in the list
func (d *Detector) isDuplicate(instances []*config.TomcatInstance, instance *config.TomcatInstance) bool {
	for _, existing := range instances {
		if existing.CatalinaHome == instance.CatalinaHome {
			return true
		}
	}
	return false
}
