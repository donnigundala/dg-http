package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/donnigundala/dg-core/errors"
	"github.com/gin-gonic/gin"
)

// Timeout returns a middleware that adds a timeout to the context.
// Note: This implements cooperative cancellation. Handlers must respect c.Request.Context().
// For strict HTTP timeout (creating 504 even if handler is blocked), a buffering middleware is needed,
// but that has performance implications. This middleware ensures the context is cancelled.
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Update request with new context
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		// Check if context timed out
		if ctx.Err() == context.DeadlineExceeded {
			// Only write error if headers haven't been written
			if !c.Writer.Written() {
				err := errors.New("request timeout").
					WithCode("REQUEST_TIMEOUT").
					WithStatusCode(http.StatusRequestTimeout).
					WithField("timeout", timeout.String())

				c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{
					"error": gin.H{
						"code":    err.Code,
						"message": err.Message,
					},
				})
			}
		}
	}
}
