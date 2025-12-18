package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donnigundala/dg-core/ctxutil"
	"github.com/donnigundala/dg-http/middleware"
	"github.com/gin-gonic/gin"
)

func setupRequestIDRouter(m gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(m)
	return r
}

// TestRequestID_GeneratesID tests that a request ID is generated when not provided
func TestRequestID_GeneratesID(t *testing.T) {
	r := setupRequestIDRouter(middleware.RequestID())
	r.GET("/test", func(c *gin.Context) {
		// Verify request ID is in context
		requestID := ctxutil.RequestIDFromContext(c.Request.Context())
		if requestID == "" {
			t.Error("Expected request ID to be set in context")
		}

		// Verify request ID is in Gin context
		ginRequestID, exists := c.Get("request_id")
		if !exists {
			t.Error("Expected request ID to be set in Gin context")
		}
		if ginRequestID != requestID {
			t.Errorf("Request ID mismatch: context=%s, gin=%s", requestID, ginRequestID)
		}

		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Verify response header contains request ID
	responseID := w.Header().Get("X-Request-ID")
	if responseID == "" {
		t.Error("Expected X-Request-ID header in response")
	}
}

// TestRequestID_UsesProvidedID tests that provided request ID is used
func TestRequestID_UsesProvidedID(t *testing.T) {
	r := setupRequestIDRouter(middleware.RequestID())
	r.GET("/test", func(c *gin.Context) {
		requestID := ctxutil.RequestIDFromContext(c.Request.Context())
		if requestID != "test-request-id-123" {
			t.Errorf("Expected request ID 'test-request-id-123', got '%s'", requestID)
		}
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Request-ID", "test-request-id-123")
	r.ServeHTTP(w, req)

	// Verify the same ID is returned in response
	responseID := w.Header().Get("X-Request-ID")
	if responseID != "test-request-id-123" {
		t.Errorf("Expected response ID 'test-request-id-123', got '%s'", responseID)
	}
}

// TestRequestIDWithDefault tests the default configuration
func TestRequestIDWithDefault(t *testing.T) {
	r := setupRequestIDRouter(middleware.RequestIDWithDefault())
	r.GET("/test", func(c *gin.Context) {
		requestID := ctxutil.RequestIDFromContext(c.Request.Context())
		if requestID == "" {
			t.Error("Expected request ID to be set")
		}
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
