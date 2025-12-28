package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

// TracingMiddleware injects a span into the request context.
type TracingMiddleware struct {
	ServiceName string
	DebugMode   bool
}

// NewTracingMiddleware creates a new tracing middleware.
func NewTracingMiddleware(serviceName string, debugMode bool) *TracingMiddleware {
	return &TracingMiddleware{
		ServiceName: serviceName,
		DebugMode:   debugMode,
	}
}

// Handle implements the Gin middleware handler.
func (m *TracingMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.Tracer(m.ServiceName)
		savedCtx := c.Request.Context()
		defer func() {
			c.Request = c.Request.WithContext(savedCtx)
		}()

		// 1. Extract context from headers (Propagation)
		ctx := otel.GetTextMapPropagator().Extract(savedCtx, propagation.HeaderCarrier(c.Request.Header))

		// 2. Start Span
		opts := []trace.SpanStartOption{
			trace.WithAttributes(semconv.HTTPMethodKey.String(c.Request.Method)),
			trace.WithAttributes(semconv.HTTPRouteKey.String(c.FullPath())),
			trace.WithAttributes(semconv.HTTPUserAgentKey.String(c.Request.UserAgent())),
			trace.WithSpanKind(trace.SpanKindServer),
		}

		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.FullPath())
		if c.FullPath() == "" {
			spanName = fmt.Sprintf("%s %s", c.Request.Method, "unknown_route")
		}

		ctx, span := tracer.Start(ctx, spanName, opts...)
		defer span.End()

		// 3. Update Request Context
		c.Request = c.Request.WithContext(ctx)

		// 4. Inject Trace ID into response headers (if DebugMode is enabled)
		if m.DebugMode {
			sc := span.SpanContext()
			if sc.IsValid() {
				c.Writer.Header().Set("X-Trace-Id", sc.TraceID().String())
				c.Writer.Header().Set("X-Span-Id", sc.SpanID().String())
			}
		}

		c.Next()

		// 5. Record Status Code
		status := c.Writer.Status()
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(status))

		if status >= 500 {
			span.RecordError(fmt.Errorf("HTTP %d", status))
		}
	}
}
