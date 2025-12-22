package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	instrumentationName = "github.com/donnigundala/dg-http"
)

// ObservabilityConfig configures the observability middleware.
type ObservabilityConfig struct {
	ServiceName string
}

// Observability returns a middleware that records HTTP metrics.
func Observability(config ObservabilityConfig) gin.HandlerFunc {
	meter := otel.GetMeterProvider().Meter(instrumentationName)

	requestCounter, err := meter.Int64Counter(
		"http.server.request.count",
		metric.WithDescription("Total number of HTTP requests processed"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		// Log error but don't panic? For now, we assume OTel won't fail easily.
		// In a real plugin we might want a way to report this initialization error.
	}

	durationHistogram, err := meter.Float64Histogram(
		"http.server.request.duration",
		metric.WithDescription("Duration of HTTP requests"),
		metric.WithUnit("ms"),
	)
	if err != nil {
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := float64(time.Since(start).Milliseconds())

		status := c.Writer.Status()
		method := c.Request.Method
		route := c.FullPath()
		if route == "" {
			route = "unknown"
		}

		attrs := metric.WithAttributes(
			attribute.String("http.method", method),
			attribute.Int("http.status_code", status),
			attribute.String("http.route", route),
		)

		requestCounter.Add(c.Request.Context(), 1, attrs)
		durationHistogram.Record(c.Request.Context(), duration, attrs)
	}
}
