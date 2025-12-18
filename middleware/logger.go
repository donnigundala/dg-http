package middleware

import (
	"time"

	"github.com/donnigundala/dg-core/logging"
	"github.com/gin-gonic/gin"
)

// LoggerConfig defines the configuration for the logger middleware
type LoggerConfig struct {
	Logger        *logging.Logger
	SkipPaths     []string
	LogLatency    bool
	LogClientIP   bool
	LogMethod     bool
	LogPath       bool
	LogStatusCode bool
}

// DefaultLoggerConfig returns the default logger configuration
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Logger:        logging.Default(),
		SkipPaths:     []string{},
		LogLatency:    true,
		LogClientIP:   true,
		LogMethod:     true,
		LogPath:       true,
		LogStatusCode: true,
	}
}

// Logger returns a middleware that logs HTTP requests
func Logger(config LoggerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip logging for certain paths
		path := c.Request.URL.Path
		for _, skip := range config.SkipPaths {
			if path == skip {
				c.Next()
				return
			}
		}

		start := time.Now()
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Build log attributes
		attrs := make([]interface{}, 0, 10)

		if config.LogStatusCode {
			attrs = append(attrs, "status", c.Writer.Status())
		}

		if config.LogMethod {
			attrs = append(attrs, "method", c.Request.Method)
		}

		if config.LogPath {
			fullPath := path
			if raw != "" {
				fullPath = path + "?" + raw
			}
			attrs = append(attrs, "path", fullPath)
		}

		if config.LogClientIP {
			attrs = append(attrs, "ip", c.ClientIP())
		}

		if config.LogLatency {
			latency := time.Since(start)
			attrs = append(attrs, "latency", latency)
		}

		// Log the request
		// Only log errors or info based on status needed?
		// For now consistent with previous behavior: Info level.
		config.Logger.Info("Request", attrs...)
	}
}

// LoggerWithDefault returns a logger middleware with default configuration
func LoggerWithDefault() gin.HandlerFunc {
	return Logger(DefaultLoggerConfig())
}
