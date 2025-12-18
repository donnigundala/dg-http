package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donnigundala/dg-http/middleware"
	"github.com/gin-gonic/gin"
)

// setupSecurityRouter is a helper to create a router with security middleware
func setupSecurityRouter(mw gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(mw)
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	return r
}

// TestSecurityHeaders_AllHeaders tests that all security headers are set
func TestSecurityHeaders_AllHeaders(t *testing.T) {
	config := middleware.DefaultSecurityConfig()
	r := setupSecurityRouter(middleware.SecurityHeaders(config))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Check that security headers are set
	headers := map[string]string{
		"X-Frame-Options":           "DENY",
		"X-Content-Type-Options":    "nosniff",
		"X-XSS-Protection":          "1; mode=block",
		"Strict-Transport-Security": "max-age=31536000",
	}

	for header, expectedValue := range headers {
		actualValue := w.Header().Get(header)
		if actualValue != expectedValue {
			t.Errorf("Expected %s header to be '%s', got '%s'", header, expectedValue, actualValue)
		}
	}
}

// TestSecurityHeaders_CustomConfig tests custom security configuration
func TestSecurityHeaders_CustomConfig(t *testing.T) {
	config := middleware.SecurityConfig{
		XFrameOptions:         "SAMEORIGIN",
		XContentTypeOptions:   "nosniff",
		XSSProtection:         "0",
		HSTSMaxAge:            63072000,
		HSTSIncludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'",
		ReferrerPolicy:        "no-referrer",
	}

	r := setupSecurityRouter(middleware.SecurityHeaders(config))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Header().Get("X-Frame-Options") != "SAMEORIGIN" {
		t.Errorf("Expected X-Frame-Options to be 'SAMEORIGIN', got '%s'", w.Header().Get("X-Frame-Options"))
	}

	if w.Header().Get("Content-Security-Policy") != "default-src 'self'" {
		t.Errorf("Expected CSP to be 'default-src 'self'', got '%s'", w.Header().Get("Content-Security-Policy"))
	}
}

// TestSecurityHeadersWithDefault tests the default security middleware
func TestSecurityHeadersWithDefault(t *testing.T) {
	r := setupSecurityRouter(middleware.SecurityHeadersWithDefault())

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Should have at least some security headers
	if w.Header().Get("X-Frame-Options") == "" {
		t.Error("Expected X-Frame-Options header to be set")
	}

	if w.Header().Get("X-Content-Type-Options") == "" {
		t.Error("Expected X-Content-Type-Options header to be set")
	}
}

// TestSecurityHeaders_DoesNotOverrideExisting tests that existing headers are not overridden (or are set correctly)
func TestSecurityHeaders_DoesNotOverrideExisting(t *testing.T) {
	config := middleware.DefaultSecurityConfig()

	// Create a router where the handler sets the header
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.SecurityHeaders(config))
	r.GET("/test", func(c *gin.Context) {
		// Try to override logic if middleware sets it AFTER Next or BEFORE Next.
		// Usually security headers middleware sets headers BEFORE calling next.
		// If handler sets it, it should override what middleware set.

		c.Header("X-Frame-Options", "ALLOW-FROM https://example.com")
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// If middleware sets it before Next(), handler should be able to overwrite it.
	// If middleware sets it after Next(), middleware overwrites handler (bad).
	// Let's assume standard behavior: Handler wins if it explicitly sets it.

	frameOptions := w.Header().Get("X-Frame-Options")
	if frameOptions != "ALLOW-FROM https://example.com" {
		t.Errorf("Expected X-Frame-Options to be 'ALLOW-FROM https://example.com', got '%s'", frameOptions)
	}
}
