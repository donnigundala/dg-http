package dghttp_test

import (
	"context"
	"mime/multipart"
	"testing"

	"github.com/dgframe/core/logging"
	"github.com/dgframe/core/observability"
	dghttp "github.com/dgframe/dg-http"
	"github.com/stretchr/testify/assert"
)

// mockContext implements contracts.Context for testing.
type mockContext struct {
	dghttp.Context
	ctx    context.Context
	logger logging.Logger
	tracer observability.Tracer
}

func (m *mockContext) Request() context.Context     { return m.ctx }
func (m *mockContext) Logger() logging.Logger       { return m.logger }
func (m *mockContext) Tracer() observability.Tracer { return m.tracer }

// Stub out other methods
func (m *mockContext) Param(key string) string                             { return "" }
func (m *mockContext) Query(key string) string                             { return "" }
func (m *mockContext) Bind(obj interface{}) error                          { return nil }
func (m *mockContext) Status(code int)                                     {}
func (m *mockContext) JSON(code int, obj interface{})                      {}
func (m *mockContext) FormFile(name string) (*multipart.FileHeader, error) { return nil, nil }
func (m *mockContext) MultipartForm() (*multipart.Form, error)             { return nil, nil }
func (m *mockContext) Next()                                               {}
func (m *mockContext) Abort()                                              {}
func (m *mockContext) AbortWithStatusJSON(code int, obj interface{})       {}

func TestMiddlewareObservability(t *testing.T) {
	// 1. Setup mock observability (Simulating Kernel-provided slots)
	mockLogger := logging.NewNoop()
	mockTracer := observability.NoopTracer{}

	ctx := &mockContext{
		ctx:    context.Background(),
		logger: mockLogger,
		tracer: mockTracer,
	}

	// 2. Perform the "Ghost Test" - Use observability in a generic HTTP context
	performObservabilityWork(ctx)

	assert.NotNil(t, ctx.Logger())
	assert.NotNil(t, ctx.Tracer())
}

// performObservabilityWork is a "Ghost" function that doesn't know about OTel or slog.
// It only knows about the dg-http contracts and core interfaces.
func performObservabilityWork(ctx dghttp.Context) {
	// Use the logger
	ctx.Logger().Info("Handling request")

	// Start a span
	_, span := ctx.Tracer().Start(ctx.Request(), "http.handle")
	defer span.End()

	// Add an attribute
	span.SetAttributes(observability.StringAttr("http.method", "GET"))
}
