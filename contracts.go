package dghttp

import (
	"context"
	"mime/multipart"

	"github.com/dgframe/core/logging"
	"github.com/dgframe/core/observability"
)

// Router defines the authoritative contract for an HTTP router.
// Implementation is owned by the application layer.
type Router interface {
	NoopMarker
	RouteGroup
}

// NoopMarker allows the system to detect if a capability is a no-op implementation.
type NoopMarker interface {
	IsNoop() bool
}

// RouteGroup defines a scoped set of routes.
type RouteGroup interface {
	// Handle registers a new route with the given method and path.
	Handle(method, path string, handlers ...Middleware)

	// Group creates a nested route group.
	Group(prefix string) RouteGroup

	// Use adds middleware to the group.
	Use(middleware ...Middleware)

	// Convenience methods
	GET(path string, handlers ...Middleware)
	POST(path string, handlers ...Middleware)
	PUT(path string, handlers ...Middleware)
	DELETE(path string, handlers ...Middleware)
	PATCH(path string, handlers ...Middleware)
}

// Controller defines the interface for module-level route registration.
type Controller interface {
	// RegisterRoutes is called by the application to bind routes to the router.
	RegisterRoutes(router Router)
}

// Middleware defines the contract for request interception.
type Middleware func(ctx Context)

// Context defines the abstracted request/response boundary.
// This interface prevents controllers from depending on concrete engines like *gin.Context.
type Context interface {
	// Request returns the underlying context.Context.
	Request() context.Context

	// Logger returns the request-scoped logger.
	Logger() logging.Logger

	// Tracer returns the request-scoped tracer.
	Tracer() observability.Tracer

	// Params returns route parameters.
	Param(key string) string

	// Query returns query string parameters.
	Query(key string) string

	// Bind binds the request body to the given object.
	Bind(obj interface{}) error

	// FormFile returns the first file for the provided form key.
	FormFile(name string) (*multipart.FileHeader, error)

	// MultipartForm returns the parsed multipart form.
	MultipartForm() (*multipart.Form, error)

	// JSON sends a JSON response.
	JSON(code int, obj interface{})

	// Status sets the HTTP status code.
	Status(code int)

	// Next executes the next handler in the chain.
	Next()

	// Abort interrupts the request chain.
	Abort()

	// AbortWithStatusJSON sends an error response and interrupts the chain.
	AbortWithStatusJSON(code int, obj interface{})
}
