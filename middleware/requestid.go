package middleware

import (
	"github.com/donnigundala/dg-core/ctxutil"
	"github.com/donnigundala/dg-core/logging"
	"github.com/gin-gonic/gin"
)

// RequestID is a Gin middleware that injects a unique request ID and context-aware logger into each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the request ID from the header or create a new one
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = ctxutil.NewRequestID()
		}

		// Add the ID to the response header so the client can track it
		c.Header("X-Request-ID", requestID)

		// Store the request ID in the Gin context
		c.Set("request_id", requestID)

		// Create a child logger with the request_id field
		logger := logging.Default().With("request_id", requestID)

		// Store the logger and request ID in the standard context
		ctx := c.Request.Context()
		ctx = ctxutil.WithLogger(ctx, logger.Underlying())
		ctx = ctxutil.WithRequestID(ctx, requestID)

		// Update the request with the new context
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// RequestIDWithDefault returns the RequestID middleware with default configuration.
func RequestIDWithDefault() gin.HandlerFunc {
	return RequestID()
}
