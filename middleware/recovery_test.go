package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donnigundala/dg-http/middleware"
	"github.com/donnigundala/dg-core/logging"
	"github.com/gin-gonic/gin"
)

func setupRecoveryRouter(m gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(m)
	return r
}

// TestRecovery_PanicHandling tests that panics are recovered
func TestRecovery_PanicHandling(t *testing.T) {
	logger := logging.Default()

	r := setupRecoveryRouter(middleware.Recovery(logger))
	r.GET("/test", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

// TestRecovery_NormalRequest tests that normal requests pass through
func TestRecovery_NormalRequest(t *testing.T) {
	logger := logging.Default()

	r := setupRecoveryRouter(middleware.Recovery(logger))
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

// TestRecoveryWithDefault tests the default recovery middleware
func TestRecoveryWithDefault(t *testing.T) {
	r := setupRecoveryRouter(middleware.RecoveryWithDefault())
	r.GET("/test", func(c *gin.Context) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", w.Code)
	}
}

// TestRecovery_DifferentPanicTypes tests recovery with different panic types
func TestRecovery_DifferentPanicTypes(t *testing.T) {
	logger := logging.Default()

	testCases := []struct {
		name     string
		panicVal interface{}
	}{
		{"String", "string panic"},
		{"Error", http.ErrAbortHandler},
		{"Int", 42},
		{"Nil", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := setupRecoveryRouter(middleware.Recovery(logger))
			r.GET("/test", func(c *gin.Context) {
				panic(tc.panicVal)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			r.ServeHTTP(w, req)

			if w.Code != http.StatusInternalServerError {
				t.Errorf("Expected status 500, got %d", w.Code)
			}
		})
	}
}
