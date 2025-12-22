package health

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsHandler returns a Gin handler for Prometheus metrics.
func MetricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// LivenessHandler returns a Gin handler for liveness probes.
// Liveness probes check if the application is running.
func LivenessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "alive",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	}
}

// ReadinessHandler returns a Gin handler for readiness probes.
// Readiness probes check if the application is ready to serve traffic.
func (m *Manager) ReadinessHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		results := m.CheckAll(ctx)

		// Check if all checks passed
		allHealthy := true
		for _, err := range results {
			if err != nil {
				allHealthy = false
				break
			}
		}

		if allHealthy {
			c.JSON(200, gin.H{
				"status":    "ready",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		} else {
			c.JSON(503, gin.H{
				"status":    "not_ready",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		}
	}
}

// HealthHandler returns a detailed health check handler.
// This provides more detailed information about each health check.
func (m *Manager) HealthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		results := m.CheckAll(ctx)

		// Build response
		allHealthy := true
		checks := make(map[string]gin.H)

		for name, err := range results {
			if err != nil {
				allHealthy = false
				checks[name] = gin.H{
					"status": "unhealthy",
					"error":  err.Error(),
				}
			} else {
				checks[name] = gin.H{
					"status": "healthy",
				}
			}
		}

		status := "healthy"
		if !allHealthy {
			status = "unhealthy"
		}

		statusCode := 200
		if !allHealthy {
			statusCode = 503
		}

		c.JSON(statusCode, gin.H{
			"status":    status,
			"timestamp": time.Now().Format(time.RFC3339),
			"checks":    checks,
		})
	}
}
