package gin

import (
	"net/http"
	"time"

	"github.com/donnigundala/dg-core/contracts/foundation"
	dgobs "github.com/donnigundala/dg-observability"
	obscontracts "github.com/donnigundala/dg-observability/contracts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ⚖️ GOVERNANCE: Adapter Middleware Pattern
//
// These middleware functions are COMPOSITION HELPERS, not authority.
// They MUST NOT:
//   - Introduce new semantics beyond dg-http and dg-observability contracts
//   - Implement hidden retries or policy decisions
//   - Parse configuration or manage lifecycle
//   - Create circular dependencies
//
// They exist solely to bridge Gin's execution model to sovereign contracts.

// RequestID adds a unique request ID to each request.
// This is a stateless helper with no external dependencies.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// Logger creates a logging middleware that resolves the logger from the container.
// Uses TYPED BINDINGS to prevent runtime breakage.
func Logger(app foundation.Application) gin.HandlerFunc {
	loggerRaw, err := app.Make(dgobs.LoggerBinding)
	if err != nil {
		// Fallback to no-op if logger not configured
		return func(c *gin.Context) { c.Next() }
	}
	logger := loggerRaw.(obscontracts.Logger)

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		requestID := c.GetString("RequestID")

		logger.Info("HTTP Request",
			obscontracts.StringAttr("method", c.Request.Method),
			obscontracts.StringAttr("path", path),
			obscontracts.StringAttr("query", query),
			obscontracts.Int64Attr("status", int64(status)),
			obscontracts.Int64Attr("latency_ms", latency.Milliseconds()),
			obscontracts.StringAttr("request_id", requestID),
		)
	}
}

// Tracing creates an OTel tracing middleware.
// Uses TYPED BINDINGS to resolve the tracer from the container.
func Tracing(app foundation.Application, serviceName string) gin.HandlerFunc {
	tracerRaw, err := app.Make(dgobs.TracerBinding)
	if err != nil {
		// Fallback to no-op if tracer not configured
		return func(c *gin.Context) { c.Next() }
	}
	tracer := tracerRaw.(obscontracts.Tracer)

	return func(c *gin.Context) {
		// Start span without attributes first, then set them via SetAttributes
		ctx, span := tracer.Start(c.Request.Context(), "http.request")
		defer span.End()

		span.SetAttributes(
			obscontracts.StringAttr("http.method", c.Request.Method),
			obscontracts.StringAttr("http.path", c.Request.URL.Path),
			obscontracts.StringAttr("service.name", serviceName),
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		span.SetAttributes(obscontracts.Int64Attr("http.status_code", int64(c.Writer.Status())))
	}
}

// Recovery provides panic recovery with logging.
func Recovery(app foundation.Application) gin.HandlerFunc {
	loggerRaw, err := app.Make(dgobs.LoggerBinding)
	if err != nil {
		// Fallback to gin's default recovery
		return gin.Recovery()
	}
	logger := loggerRaw.(obscontracts.Logger)

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var errMsg string
				switch v := r.(type) {
				case error:
					errMsg = v.Error()
				case string:
					errMsg = v
				default:
					errMsg = "unknown panic"
				}

				logger.Error("Panic recovered",
					obscontracts.StringAttr("error", errMsg),
					obscontracts.StringAttr("path", c.Request.URL.Path),
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}

// CORS provides basic CORS support.
// This is a stateless helper with no external dependencies.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization, X-Request-ID")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// SecurityHeaders adds basic security headers.
// This is a stateless helper with no external dependencies.
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	}
}

// BodySizeLimit limits the size of the request body.
func BodySizeLimit(limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
		c.Next()
	}
}
