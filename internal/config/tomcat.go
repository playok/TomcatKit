package config

// TomcatInstance represents a detected Tomcat installation
type TomcatInstance struct {
	CatalinaHome string
	CatalinaBase string
	Version      string
	IsRunning    bool
	PID          int
}
