package model

import (
	"encoding/json"
	"net/http"
)

// CoprocessorStage represents the enum for stages
type CoprocessorStage string

// CoprocessorBody represents the body structure
type CoprocessorBody struct {
	Query         string      `json:"query,omitempty"`
	OperationName string      `json:"operationName,omitempty"`
	Variables     interface{} `json:"variables,omitempty"`
	Errors        []Error     `json:"errors,omitempty"`
}

// ErrorBody represents the structure of the error object
type Error struct {
	HTTPStatus int        `json:"httpstatus"`
	Message    string     `json:"message"`
	Code       string     `json:"code"`
	Extensions Extensions `json:"extensions"`
}

type Extensions struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	DeveloperText string `json:"developer_text"`
}

// CoprocessorContext represents the context structure
type CoprocessorContext struct {
	Entries map[string]interface{} `json:"entries"` // Context entries as a map
}

type CoprocessorPayload struct {
	Version     int                `json:"version,omitempty"`
	Stage       CoprocessorStage   `json:"stage,omitempty"`
	Control     any                `json:"control,omitempty"`
	ID          string             `json:"id,omitempty"`
	Headers     http.Header        `json:"headers,omitempty"`
	Body        CoprocessorBody    `json:"body,omitempty"`
	Context     CoprocessorContext `json:"context,omitempty"`
	SDL         string             `json:"sdl,omitempty"`
	Method      string             `json:"method,omitempty"`
	Path        string             `json:"path,omitempty"`        // The RouterService or SupergraphService path that this coprocessor request pertains to.
	ServiceName string             `json:"serviceName,omitempty"` //SubgraphRequest Stage
	URI         string             `json:"uri,omitempty"`         //SubgraphRequest Stage
}

type RouterResponsePayload struct {
	Version     int              `json:"version,omitempty"`
	Stage       CoprocessorStage `json:"stage,omitempty"`
	Control     any              `json:"control,omitempty"`
	ID          string           `json:"id,omitempty"`
	Headers     http.Header      `json:"headers,omitempty"`
	Body        json.RawMessage  `json:"body,omitempty"`
	Context     json.RawMessage  `json:"context,omitempty"`
	SDL         string           `json:"sdl,omitempty"`
	Method      string           `json:"method,omitempty"`
	Path        string           `json:"path,omitempty"`        // The RouterService or SupergraphService path that this coprocessor request pertains to.
	ServiceName string           `json:"serviceName,omitempty"` //SubgraphRequest Stage
	URI         string           `json:"uri,omitempty"`         //SubgraphRequest Stage
}
