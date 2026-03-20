package dghttp

import (
	"context"
	"mime/multipart"

	"github.com/dgframe/core/logging"
	"github.com/dgframe/core/observability"
)

// NewNoopRouter creates a silent, inert router.
// NewNoopRouter creates a fail-fast placeholder router that panics if used without a bound HTTP capability.
func NewNoopRouter() Router {
	return &noopRouter{}
}

type noopRouter struct{}

func (n *noopRouter) IsNoop() bool { return true }

func (n *noopRouter) Group(prefix string) RouteGroup {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) Use(middleware ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) Handle(method, path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) GET(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) POST(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) PUT(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) DELETE(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouter) PATCH(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

type noopRouteGroup struct{}

func (n *noopRouteGroup) Handle(method, path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) Group(prefix string) RouteGroup {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) Use(middleware ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) GET(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) POST(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) PUT(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) DELETE(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

func (n *noopRouteGroup) PATCH(path string, handlers ...Middleware) {
	panic("dg-http: Router capability not provided (Type B violation)")
}

// NewNoopContext creates an inert context.
func NewNoopContext() Context {
	return &noopContext{}
}

// noopContext is a no-op implementation of Context.
type noopContext struct{}

func (n *noopContext) Request() context.Context     { return context.Background() }
func (n *noopContext) Logger() logging.Logger       { return logging.NewNoop() }
func (n *noopContext) Tracer() observability.Tracer { return observability.NoopTracer{} }
func (n *noopContext) Param(key string) string      { return "" }

func (n *noopContext) Query(key string) string {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) Bind(obj interface{}) error {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) FormFile(name string) (*multipart.FileHeader, error) {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) MultipartForm() (*multipart.Form, error) {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) JSON(code int, obj interface{}) {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) Status(code int) {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) Next() {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) Abort() {
	panic("dg-http: Context capability not provided (Type B violation)")
}

func (n *noopContext) AbortWithStatusJSON(code int, obj interface{}) {
	panic("dg-http: Context capability not provided (Type B violation)")
}
