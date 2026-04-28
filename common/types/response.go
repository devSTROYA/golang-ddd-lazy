package types

type ErrorResponse struct {
	TraceId string      `json:"traceId,omitempty"`
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string       `json:"code"`
	Details []FieldError `json:"details,omitempty"`
}

type FieldError struct {
	Code     string         `json:"code"`
	Field    string         `json:"field"`
	Metadata map[string]any `json:"metadata,omitempty"`
}
