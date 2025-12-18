package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/donnigundala/dg-core/errors"
	"github.com/donnigundala/dg-core/logging"
	"github.com/gin-gonic/gin"
)

// Recovery returns a middleware that recovers from panics.
func Recovery(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace (internal logging)
				stack := string(debug.Stack())

				logger.ErrorContext(c.Request.Context(), "Panic recovered",
					"error", err,
					"stack", stack,
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
				)

				// Create error response
				wrappedErr := errors.New(fmt.Sprintf("internal server error: %v", err)).
					WithCode("PANIC_RECOVERED").
					WithStatusCode(http.StatusInternalServerError)

				// Report to external service (if configured)
				errors.Report(wrappedErr, map[string]interface{}{
					"method": c.Request.Method,
					"path":   c.Request.URL.Path,
					"stack":  stack,
				})

				// Write error response
				// errors.WriteHTTPError would write to w, but we should use Gin
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"code":    wrappedErr.Code,
						"message": "Internal Server Error",
					},
				})
			}
		}()

		c.Next()
	}
}

// RecoveryWithDefault returns a Recovery middleware using the default logger.
func RecoveryWithDefault() gin.HandlerFunc {
	return Recovery(logging.Default())
}
