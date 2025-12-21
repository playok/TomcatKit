package parser

import (
	"encoding/xml"
	"os"

	"github.com/playok/tomcatkit/internal/config/server"
)

// ParseServerXML parses a server.xml file
func ParseServerXML(filePath string) (*server.Server, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var s server.Server
	if err := xml.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

// WriteServerXML writes server configuration to a file
func WriteServerXML(filePath string, s *server.Server) error {
	data, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	// Add XML declaration
	output := []byte(xml.Header)
	output = append(output, data...)

	return os.WriteFile(filePath, output, 0644)
}

// ParseContextXML parses a context.xml file
func ParseContextXML(filePath string) (*server.Context, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ctx server.Context
	if err := xml.Unmarshal(data, &ctx); err != nil {
		return nil, err
	}

	return &ctx, nil
}

// WriteContextXML writes context configuration to a file
func WriteContextXML(filePath string, ctx *server.Context) error {
	data, err := xml.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return err
	}

	output := []byte(xml.Header)
	output = append(output, data...)

	return os.WriteFile(filePath, output, 0644)
}
