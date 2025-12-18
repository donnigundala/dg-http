package response

import (
	"encoding/json"
	"net/http"

	"github.com/donnigundala/dg-core/errors"
)

// HTTPError represents an error response for HTTP APIs.
type HTTPError struct {
	Code    string                 `json:"code,omitempty"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

// ToHTTPError converts an Error to HTTPError.
func ToHTTPError(err error) (int, *HTTPError) {
	if err == nil {
		return http.StatusOK, nil
	}

	// Check if it's our custom Error type
	if e, ok := err.(*errors.Error); ok {
		return e.StatusCode(), &HTTPError{
			Code:    e.Code(),
			Message: e.Message(),
			Fields:  e.Fields(),
		}
	}

	// Default error response
	return http.StatusInternalServerError, &HTTPError{
		Code:    "INTERNAL_ERROR",
		Message: err.Error(),
	}
}

// WriteHTTPError writes an error response to http.ResponseWriter.
func WriteHTTPError(w http.ResponseWriter, err error) {
	status, httpErr := ToHTTPError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(httpErr)
}

// HTTPStatusFromError extracts HTTP status code from error.
func HTTPStatusFromError(err error) int {
	if e, ok := err.(*errors.Error); ok {
		return e.StatusCode()
	}
	return http.StatusInternalServerError
}
