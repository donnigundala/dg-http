package contracts

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response.
type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Code    string      `json:"code,omitempty"`
	Fields  interface{} `json:"fields,omitempty"`
}
