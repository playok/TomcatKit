package connector

import (
	"github.com/playok/tomcatkit/internal/config/server"
)

// ConnectorType represents the type of connector
type ConnectorType string

const (
	ConnectorTypeHTTP  ConnectorType = "HTTP/1.1"
	ConnectorTypeHTTP2 ConnectorType = "HTTP/2"
	ConnectorTypeAJP   ConnectorType = "AJP/1.3"
)

// Protocol constants
const (
	ProtocolHTTP11Nio  = "org.apache.coyote.http11.Http11NioProtocol"
	ProtocolHTTP11Nio2 = "org.apache.coyote.http11.Http11Nio2Protocol"
	ProtocolHTTP11Apr  = "org.apache.coyote.http11.Http11AprProtocol"
	ProtocolHTTP2      = "org.apache.coyote.http2.Http2Protocol"
	ProtocolAJPNio     = "org.apache.coyote.ajp.AjpNioProtocol"
	ProtocolAJPNio2    = "org.apache.coyote.ajp.AjpNio2Protocol"
	ProtocolAJPApr     = "org.apache.coyote.ajp.AjpAprProtocol"
)

// GetConnectorType determines the connector type from protocol
func GetConnectorType(protocol string) ConnectorType {
	switch protocol {
	case "AJP/1.3", ProtocolAJPNio, ProtocolAJPNio2, ProtocolAJPApr:
		return ConnectorTypeAJP
	case ProtocolHTTP2:
		return ConnectorTypeHTTP2
	default:
		return ConnectorTypeHTTP
	}
}

// GetProtocolDescription returns a human-readable protocol description
func GetProtocolDescription(protocol string) string {
	descriptions := map[string]string{
		"":                 "HTTP/1.1 (Auto - blocking Java connector)",
		"HTTP/1.1":         "HTTP/1.1 (Auto)",
		ProtocolHTTP11Nio:  "HTTP/1.1 NIO (Non-blocking)",
		ProtocolHTTP11Nio2: "HTTP/1.1 NIO2 (Async non-blocking)",
		ProtocolHTTP11Apr:  "HTTP/1.1 APR (Native)",
		ProtocolHTTP2:      "HTTP/2 (Requires SSL)",
		"AJP/1.3":          "AJP/1.3 (Auto)",
		ProtocolAJPNio:     "AJP NIO (Non-blocking)",
		ProtocolAJPNio2:    "AJP NIO2 (Async non-blocking)",
		ProtocolAJPApr:     "AJP APR (Native)",
	}
	if desc, ok := descriptions[protocol]; ok {
		return desc
	}
	return protocol
}

// DefaultHTTPConnector creates a default HTTP connector
func DefaultHTTPConnector() server.Connector {
	return server.Connector{
		Port:              8080,
		Protocol:          ProtocolHTTP11Nio,
		ConnectionTimeout: 20000,
		RedirectPort:      8443,
		MaxThreads:        200,
		MinSpareThreads:   10,
	}
}

// DefaultHTTPSConnector creates a default HTTPS connector
func DefaultHTTPSConnector() server.Connector {
	return server.Connector{
		Port:              8443,
		Protocol:          ProtocolHTTP11Nio,
		SSLEnabled:        true,
		Scheme:            "https",
		Secure:            true,
		ConnectionTimeout: 20000,
		MaxThreads:        200,
		MinSpareThreads:   10,
		KeystoreFile:      "${user.home}/.keystore",
		KeystorePass:      "changeit",
		KeystoreType:      "JKS",
		ClientAuth:        "false",
		SSLProtocol:       "TLS",
	}
}

// DefaultAJPConnector creates a default AJP connector
func DefaultAJPConnector() server.Connector {
	return server.Connector{
		Port:           8009,
		Protocol:       ProtocolAJPNio,
		RedirectPort:   8443,
		SecretRequired: true,
		Secret:         "",
	}
}

// AvailableProtocols returns available HTTP protocols
func AvailableHTTPProtocols() []string {
	return []string{
		"HTTP/1.1",
		ProtocolHTTP11Nio,
		ProtocolHTTP11Nio2,
		ProtocolHTTP11Apr,
	}
}

// AvailableAJPProtocols returns available AJP protocols
func AvailableAJPProtocols() []string {
	return []string{
		"AJP/1.3",
		ProtocolAJPNio,
		ProtocolAJPNio2,
		ProtocolAJPApr,
	}
}

// SSLProtocols returns available SSL protocols
func SSLProtocols() []string {
	return []string{
		"TLS",
		"TLSv1",
		"TLSv1.1",
		"TLSv1.2",
		"TLSv1.3",
	}
}

// KeystoreTypes returns available keystore types
func KeystoreTypes() []string {
	return []string{
		"JKS",
		"PKCS12",
		"PKCS11",
	}
}

// ClientAuthOptions returns available client auth options
func ClientAuthOptions() []string {
	return []string{
		"false",
		"true",
		"want",
	}
}
