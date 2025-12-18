package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/donnigundala/dg-core/errors"
)

// RateLimitConfig defines the configuration for rate limiting middleware.
type RateLimitConfig struct {
	// RequestsPerSecond is the number of requests allowed per second.
	// Default: 10
	RequestsPerSecond float64

	// BurstSize is the maximum burst size.
	// Default: 20
	BurstSize int

	// KeyFunc is a function to extract the key for rate limiting (e.g., IP address).
	// Default: uses client IP from Gin
	KeyFunc func(*gin.Context) string
}

// DefaultRateLimitConfig returns the default rate limit configuration.
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		RequestsPerSecond: 10,
		BurstSize:         20,
		KeyFunc:           func(c *gin.Context) string { return c.ClientIP() },
	}
}

// RateLimit returns a middleware that limits requests per client.
func RateLimit(config RateLimitConfig) gin.HandlerFunc {
	// Apply defaults
	if config.RequestsPerSecond == 0 {
		config.RequestsPerSecond = 10
	}
	if config.BurstSize == 0 {
		config.BurstSize = 20
	}
	if config.KeyFunc == nil {
		config.KeyFunc = func(c *gin.Context) string { return c.ClientIP() }
	}

	// Create a map to store limiters per client
	limiters := &sync.Map{}

	return func(c *gin.Context) {
		// Get client key
		key := config.KeyFunc(c)

		// Get or create limiter for this client
		limiterInterface, _ := limiters.LoadOrStore(key, rate.NewLimiter(
			rate.Limit(config.RequestsPerSecond),
			config.BurstSize,
		))
		limiter := limiterInterface.(*rate.Limiter)

		// Check if request is allowed
		if !limiter.Allow() {
			err := errors.ErrTooManyRequests.
				WithField("client", key)

			// Use standard error response format
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": gin.H{
					"code":    err.Code,
					"message": err.Message,
				},
			})
			return
		}

		c.Next()
	}
}

// RateLimitWithDefault returns a RateLimit middleware with default configuration.
func RateLimitWithDefault() gin.HandlerFunc {
	return RateLimit(DefaultRateLimitConfig())
}
