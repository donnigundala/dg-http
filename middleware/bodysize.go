package middleware

import (
	"net/http"

	"github.com/donnigundala/dg-core/errors"
	"github.com/gin-gonic/gin"
)

// BodySizeLimit returns a middleware that limits the request body size.
func BodySizeLimit(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only apply to methods that can have a body
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			// Limit the request body size
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		}

		c.Next()
	}
}

// BodySizeLimitWithError returns a middleware that limits body size and returns a proper error.
func BodySizeLimitWithError(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only apply to methods that can have a body
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			// Check Content-Length header first
			if c.Request.ContentLength > maxBytes {
				err := errors.New("request body too large").
					WithCode("BODY_TOO_LARGE").
					WithStatusCode(http.StatusRequestEntityTooLarge).
					WithField("max_bytes", maxBytes).
					WithField("content_length", c.Request.ContentLength)

				c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
					"error": gin.H{
						"code":    err.Code,
						"message": err.Message,
					},
				})
				return
			}

			// Limit the request body size
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		}

		c.Next()
	}
}
