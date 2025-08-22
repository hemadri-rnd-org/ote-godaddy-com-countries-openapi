package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Stack []string `json:"stack,omitempty"` // Stack trace indicating where the error occurred.<br/> NOTE: This attribute <strong>MAY</strong> be included for Development and Test environments. However, it <strong>MUST NOT</strong> be exposed from OTE nor Production systems
}

// ErrorField represents the ErrorField schema from the OpenAPI specification
type ErrorField struct {
	Message string `json:"message,omitempty"` // Human-readable, English description of the problem with the contents of the field
	Path string `json:"path"` // JSONPath referring to the field within the submitted data containing an error
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
}

// ErrorLimit represents the ErrorLimit schema from the OpenAPI specification
type ErrorLimit struct {
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Retryaftersec int `json:"retryAfterSec"` // Number of seconds to wait before attempting a similar request
	Stack []string `json:"stack,omitempty"` // Stack trace indicating where the error occurred.<br/> NOTE: This attribute <strong>MAY</strong> be included for Development and Test environments. However, it <strong>MUST NOT</strong> be exposed from OTE nor Production systems
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
}

// State represents the State schema from the OpenAPI specification
type State struct {
	Label string `json:"label,omitempty"` // The localized name of the state, province, or territory
	Statekey string `json:"stateKey,omitempty"` // The state code
}

// Country represents the Country schema from the OpenAPI specification
type Country struct {
	Label string `json:"label,omitempty"` // The localized name of the country
	States []State `json:"states,omitempty"` // List of states/provinces in this country
	Callingcode string `json:"callingCode,omitempty"` // The calling code prefix used for phone numbers in this country
	Countrykey string `json:"countryKey,omitempty"` // The ISO country-code
}

// CountrySummary represents the CountrySummary schema from the OpenAPI specification
type CountrySummary struct {
	Callingcode string `json:"callingCode,omitempty"` // The calling code prefix used for phone numbers in this country
	Countrykey string `json:"countryKey,omitempty"` // The ISO country-code
	Label string `json:"label,omitempty"` // The localized name of the country
}
