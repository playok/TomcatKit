package realm

import (
	"github.com/playok/tomcatkit/internal/config/server"
)

// RealmType represents the type of realm
type RealmType string

const (
	RealmTypeUserDatabase RealmType = "UserDatabaseRealm"
	RealmTypeDataSource   RealmType = "DataSourceRealm"
	RealmTypeJNDI         RealmType = "JNDIRealm"
	RealmTypeJAAS         RealmType = "JAASRealm"
	RealmTypeCombined     RealmType = "CombinedRealm"
	RealmTypeLockOut      RealmType = "LockOutRealm"
	RealmTypeMemory       RealmType = "MemoryRealm"
)

// Realm class names
const (
	ClassUserDatabaseRealm = "org.apache.catalina.realm.UserDatabaseRealm"
	ClassDataSourceRealm   = "org.apache.catalina.realm.DataSourceRealm"
	ClassJNDIRealm         = "org.apache.catalina.realm.JNDIRealm"
	ClassJAASRealm         = "org.apache.catalina.realm.JAASRealm"
	ClassCombinedRealm     = "org.apache.catalina.realm.CombinedRealm"
	ClassLockOutRealm      = "org.apache.catalina.realm.LockOutRealm"
	ClassMemoryRealm       = "org.apache.catalina.realm.MemoryRealm"
)

// GetRealmType determines the realm type from class name
func GetRealmType(className string) RealmType {
	switch className {
	case ClassUserDatabaseRealm:
		return RealmTypeUserDatabase
	case ClassDataSourceRealm:
		return RealmTypeDataSource
	case ClassJNDIRealm:
		return RealmTypeJNDI
	case ClassJAASRealm:
		return RealmTypeJAAS
	case ClassCombinedRealm:
		return RealmTypeCombined
	case ClassLockOutRealm:
		return RealmTypeLockOut
	case ClassMemoryRealm:
		return RealmTypeMemory
	default:
		return RealmType(className)
	}
}

// GetRealmDescription returns a human-readable realm description
func GetRealmDescription(className string) string {
	descriptions := map[string]string{
		ClassUserDatabaseRealm: "File-based authentication using tomcat-users.xml",
		ClassDataSourceRealm:   "Database authentication via JDBC DataSource",
		ClassJNDIRealm:         "LDAP/Directory server authentication",
		ClassJAASRealm:         "JAAS (Java Authentication and Authorization Service)",
		ClassCombinedRealm:     "Combines multiple realms with fallback",
		ClassLockOutRealm:      "Wraps realm with brute-force protection",
		ClassMemoryRealm:       "Simple in-memory authentication (demo only)",
	}
	if desc, ok := descriptions[className]; ok {
		return desc
	}
	return className
}

// GetShortRealmName returns a short name for the realm
func GetShortRealmName(className string) string {
	names := map[string]string{
		ClassUserDatabaseRealm: "UserDatabaseRealm",
		ClassDataSourceRealm:   "DataSourceRealm",
		ClassJNDIRealm:         "JNDIRealm",
		ClassJAASRealm:         "JAASRealm",
		ClassCombinedRealm:     "CombinedRealm",
		ClassLockOutRealm:      "LockOutRealm",
		ClassMemoryRealm:       "MemoryRealm",
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

// AvailableRealmTypes returns all available realm types
func AvailableRealmTypes() []string {
	return []string{
		ClassUserDatabaseRealm,
		ClassDataSourceRealm,
		ClassJNDIRealm,
		ClassJAASRealm,
		ClassCombinedRealm,
		ClassLockOutRealm,
	}
}

// DefaultUserDatabaseRealm creates a default UserDatabaseRealm
func DefaultUserDatabaseRealm() server.Realm {
	return server.Realm{
		ClassName:    ClassUserDatabaseRealm,
		ResourceName: "UserDatabase",
	}
}

// DefaultDataSourceRealm creates a default DataSourceRealm
func DefaultDataSourceRealm() server.Realm {
	return server.Realm{
		ClassName:      ClassDataSourceRealm,
		DataSourceName: "jdbc/UserDB",
		UserTable:      "users",
		UserNameCol:    "user_name",
		UserCredCol:    "user_pass",
		UserRoleTable:  "user_roles",
		RoleNameCol:    "role_name",
	}
}

// DefaultJNDIRealm creates a default JNDIRealm
func DefaultJNDIRealm() server.Realm {
	return server.Realm{
		ClassName:     ClassJNDIRealm,
		ConnectionURL: "ldap://localhost:389",
		UserPattern:   "uid={0},ou=people,dc=example,dc=com",
		RoleBase:      "ou=groups,dc=example,dc=com",
		RoleName:      "cn",
		RoleSearch:    "(uniqueMember={0})",
	}
}

// DefaultJAASRealm creates a default JAASRealm
func DefaultJAASRealm() server.Realm {
	return server.Realm{
		ClassName: ClassJAASRealm,
	}
}

// DefaultLockOutRealm creates a default LockOutRealm wrapping UserDatabaseRealm
func DefaultLockOutRealm() server.Realm {
	return server.Realm{
		ClassName: ClassLockOutRealm,
		NestedRealms: []server.Realm{
			DefaultUserDatabaseRealm(),
		},
	}
}

// DefaultCombinedRealm creates a default CombinedRealm
func DefaultCombinedRealm() server.Realm {
	return server.Realm{
		ClassName:    ClassCombinedRealm,
		NestedRealms: []server.Realm{},
	}
}

// CredentialHandlerAlgorithms returns available password hashing algorithms
func CredentialHandlerAlgorithms() []string {
	return []string{
		"SHA-256",
		"SHA-512",
		"MD5",
		"PBKDF2WithHmacSHA256",
		"PBKDF2WithHmacSHA512",
	}
}

// CredentialHandlerClasses returns available credential handler classes
func CredentialHandlerClasses() []string {
	return []string{
		"org.apache.catalina.realm.MessageDigestCredentialHandler",
		"org.apache.catalina.realm.SecretKeyCredentialHandler",
		"org.apache.catalina.realm.NestedCredentialHandler",
	}
}
