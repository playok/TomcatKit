package realm

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TomcatUsers represents the tomcat-users.xml structure
type TomcatUsers struct {
	XMLName  xml.Name `xml:"tomcat-users"`
	Xmlns    string   `xml:"xmlns,attr,omitempty"`
	XmlnsXsi string   `xml:"xmlns:xsi,attr,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
	Roles    []Role   `xml:"role"`
	Users    []User   `xml:"user"`
}

// Role represents a role definition
type Role struct {
	RoleName    string `xml:"rolename,attr"`
	Description string `xml:"description,attr,omitempty"`
}

// User represents a user definition
type User struct {
	Username string `xml:"username,attr"`
	Password string `xml:"password,attr"`
	Roles    string `xml:"roles,attr"`
}

// GetRolesList returns the roles as a slice
func (u *User) GetRolesList() []string {
	if u.Roles == "" {
		return []string{}
	}
	roles := strings.Split(u.Roles, ",")
	for i, r := range roles {
		roles[i] = strings.TrimSpace(r)
	}
	return roles
}

// SetRolesList sets the roles from a slice
func (u *User) SetRolesList(roles []string) {
	u.Roles = strings.Join(roles, ",")
}

// HasRole checks if the user has a specific role
func (u *User) HasRole(role string) bool {
	for _, r := range u.GetRolesList() {
		if r == role {
			return true
		}
	}
	return false
}

// UsersService handles tomcat-users.xml operations
type UsersService struct {
	catalinaBase string
	filePath     string
	users        *TomcatUsers
}

// NewUsersService creates a new users service
func NewUsersService(catalinaBase string) *UsersService {
	return &UsersService{
		catalinaBase: catalinaBase,
		filePath:     filepath.Join(catalinaBase, "conf", "tomcat-users.xml"),
	}
}

// Load reads and parses tomcat-users.xml
func (s *UsersService) Load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to read tomcat-users.xml: %w", err)
	}

	var users TomcatUsers
	if err := xml.Unmarshal(data, &users); err != nil {
		return fmt.Errorf("failed to parse tomcat-users.xml: %w", err)
	}

	s.users = &users
	return nil
}

// Save writes the configuration back to tomcat-users.xml
func (s *UsersService) Save() error {
	if s.users == nil {
		return fmt.Errorf("no users configuration loaded")
	}

	// Create backup first
	if err := s.createBackup(); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	// Set default namespace if not present
	if s.users.Xmlns == "" {
		s.users.Xmlns = "http://tomcat.apache.org/xml"
	}

	data, err := xml.MarshalIndent(s.users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tomcat-users.xml: %w", err)
	}

	// Add XML declaration
	output := []byte(xml.Header)
	output = append(output, data...)

	if err := os.WriteFile(s.filePath, output, 0640); err != nil {
		return fmt.Errorf("failed to write tomcat-users.xml: %w", err)
	}

	return nil
}

// createBackup creates a backup of the current tomcat-users.xml
func (s *UsersService) createBackup() error {
	backupDir := filepath.Join(s.catalinaBase, "conf", "backup")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	backupPath := filepath.Join(backupDir, "tomcat-users.xml.bak")
	return os.WriteFile(backupPath, data, 0640)
}

// GetUsers returns all users
func (s *UsersService) GetUsers() []User {
	if s.users != nil {
		return s.users.Users
	}
	return nil
}

// GetUser returns a user by username
func (s *UsersService) GetUser(username string) *User {
	if s.users != nil {
		for i := range s.users.Users {
			if s.users.Users[i].Username == username {
				return &s.users.Users[i]
			}
		}
	}
	return nil
}

// AddUser adds a new user
func (s *UsersService) AddUser(user User) error {
	if s.users == nil {
		s.users = &TomcatUsers{}
	}

	// Check for duplicate
	for _, u := range s.users.Users {
		if u.Username == user.Username {
			return fmt.Errorf("user '%s' already exists", user.Username)
		}
	}

	s.users.Users = append(s.users.Users, user)
	return nil
}

// UpdateUser updates an existing user
func (s *UsersService) UpdateUser(username string, user User) error {
	if s.users != nil {
		for i := range s.users.Users {
			if s.users.Users[i].Username == username {
				s.users.Users[i] = user
				return nil
			}
		}
	}
	return fmt.Errorf("user '%s' not found", username)
}

// DeleteUser deletes a user
func (s *UsersService) DeleteUser(username string) error {
	if s.users != nil {
		for i := range s.users.Users {
			if s.users.Users[i].Username == username {
				s.users.Users = append(s.users.Users[:i], s.users.Users[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("user '%s' not found", username)
}

// GetRoles returns all roles
func (s *UsersService) GetRoles() []Role {
	if s.users != nil {
		return s.users.Roles
	}
	return nil
}

// GetRole returns a role by name
func (s *UsersService) GetRole(roleName string) *Role {
	if s.users != nil {
		for i := range s.users.Roles {
			if s.users.Roles[i].RoleName == roleName {
				return &s.users.Roles[i]
			}
		}
	}
	return nil
}

// AddRole adds a new role
func (s *UsersService) AddRole(role Role) error {
	if s.users == nil {
		s.users = &TomcatUsers{}
	}

	// Check for duplicate
	for _, r := range s.users.Roles {
		if r.RoleName == role.RoleName {
			return fmt.Errorf("role '%s' already exists", role.RoleName)
		}
	}

	s.users.Roles = append(s.users.Roles, role)
	return nil
}

// DeleteRole deletes a role
func (s *UsersService) DeleteRole(roleName string) error {
	if s.users != nil {
		for i := range s.users.Roles {
			if s.users.Roles[i].RoleName == roleName {
				s.users.Roles = append(s.users.Roles[:i], s.users.Roles[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("role '%s' not found", roleName)
}

// GetAllRoleNames returns all unique role names (from roles and users)
func (s *UsersService) GetAllRoleNames() []string {
	roleMap := make(map[string]bool)

	if s.users != nil {
		// From role definitions
		for _, r := range s.users.Roles {
			roleMap[r.RoleName] = true
		}
		// From user role assignments
		for _, u := range s.users.Users {
			for _, r := range u.GetRolesList() {
				roleMap[r] = true
			}
		}
	}

	roles := make([]string, 0, len(roleMap))
	for r := range roleMap {
		roles = append(roles, r)
	}
	return roles
}

// Common Tomcat roles
var CommonRoles = []Role{
	{RoleName: "manager-gui", Description: "Access to HTML interface of Manager app"},
	{RoleName: "manager-script", Description: "Access to text interface of Manager app"},
	{RoleName: "manager-jmx", Description: "Access to JMX proxy interface of Manager app"},
	{RoleName: "manager-status", Description: "Access to read-only status pages of Manager app"},
	{RoleName: "admin-gui", Description: "Access to HTML interface of Host Manager app"},
	{RoleName: "admin-script", Description: "Access to text interface of Host Manager app"},
}
