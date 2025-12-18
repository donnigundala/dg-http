package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/donnigundala/dg-http/middleware"
	"github.com/donnigundala/dg-core/logging"
	"github.com/gin-gonic/gin"
)

// Helper to setup router with middleware
func setupLoggerRouter(m gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(m)
	return r
}

// TestLogger_BasicRequest tests basic request logging
func TestLogger_BasicRequest(t *testing.T) {
	logger := logging.Default()
	config := middleware.DefaultLoggerConfig()
	config.Logger = logger

	r := setupLoggerRouter(middleware.Logger(config))
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "OK" {
		t.Errorf("Expected 'OK', got '%s'", w.Body.String())
	}
}

// TestLogger_WithQueryParams tests logging with query parameters
func TestLogger_WithQueryParams(t *testing.T) {
	logger := logging.Default()
	config := middleware.DefaultLoggerConfig()
	config.Logger = logger

	r := setupLoggerRouter(middleware.Logger(config))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test?foo=bar&baz=qux", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestLogger_SkipPaths tests that certain paths are skipped
func TestLogger_SkipPaths(t *testing.T) {
	logger := logging.Default()
	config := middleware.DefaultLoggerConfig()
	config.Logger = logger
	config.SkipPaths = []string{"/health", "/metrics"}

	r := setupLoggerRouter(middleware.Logger(config))
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/api/users", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Test skipped path
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Test non-skipped path
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/users", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestLogger_ClientIP tests client IP detection
func TestLogger_ClientIP_XForwardedFor(t *testing.T) {
	logger := logging.Default()
	config := middleware.DefaultLoggerConfig()
	config.Logger = logger
	config.LogClientIP = true

	r := setupLoggerRouter(middleware.Logger(config))
	r.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Forwarded-For", "203.0.113.1")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestLogger_StatusCodes tests logging different status codes
func TestLogger_StatusCodes(t *testing.T) {
	logger := logging.Default()
	config := middleware.DefaultLoggerConfig()
	config.Logger = logger

	testCases := []struct {
		name       string
		statusCode int
	}{
		{"OK", http.StatusOK},
		{"Created", http.StatusCreated},
		{"BadRequest", http.StatusBadRequest},
		{"NotFound", http.StatusNotFound},
		{"InternalError", http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := setupLoggerRouter(middleware.Logger(config))
			r.GET("/test", func(c *gin.Context) {
				c.Status(tc.statusCode)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			r.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Errorf("Expected status %d, got %d", tc.statusCode, w.Code)
			}
		})
	}
}

// TestLoggerWithDefault tests the default logger middleware
func TestLoggerWithDefault(t *testing.T) {
	r := setupLoggerRouter(middleware.LoggerWithDefault())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// TestLogger_ConfigOptions tests various configuration options
func TestLogger_ConfigOptions(t *testing.T) {
	logger := logging.Default()
	config := middleware.LoggerConfig{
		Logger:        logger,
		SkipPaths:     []string{"/skip"},
		LogLatency:    false,
		LogClientIP:   false,
		LogMethod:     true,
		LogPath:       true,
		LogStatusCode: true,
	}

	r := setupLoggerRouter(middleware.Logger(config))
	r.POST("/api/users", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(`{"name":"test"}`))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
