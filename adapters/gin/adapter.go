package gin

import (
	"context"
	"mime/multipart"

	"github.com/donnigundala/dg-http/contracts"
	"github.com/gin-gonic/gin"
)

// Router implements dghttp.Router using Gin.
type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) contracts.Router {
	return &Router{engine: engine}
}

func (r *Router) Group(prefix string) contracts.RouteGroup {
	return &RouteGroup{group: r.engine.Group(prefix)}
}

func (r *Router) Use(middleware ...contracts.Middleware) {
	for _, m := range middleware {
		r.engine.Use(wrapMiddleware(m))
	}
}

// RouteGroup implements dghttp.RouteGroup using Gin.
type RouteGroup struct {
	group *gin.RouterGroup
}

func (g *RouteGroup) Group(prefix string) contracts.RouteGroup {
	return &RouteGroup{group: g.group.Group(prefix)}
}

func (g *RouteGroup) Use(middleware ...contracts.Middleware) {
	for _, m := range middleware {
		g.group.Use(wrapMiddleware(m))
	}
}

func (g *RouteGroup) Handle(method, path string, handlers ...contracts.Middleware) {
	ginHandlers := make([]gin.HandlerFunc, len(handlers))
	for i, h := range handlers {
		ginHandlers[i] = wrapMiddleware(h)
	}
	g.group.Handle(method, path, ginHandlers...)
}

func (g *RouteGroup) GET(path string, handlers ...contracts.Middleware) {
	g.Handle("GET", path, handlers...)
}
func (g *RouteGroup) POST(path string, handlers ...contracts.Middleware) {
	g.Handle("POST", path, handlers...)
}
func (g *RouteGroup) PUT(path string, handlers ...contracts.Middleware) {
	g.Handle("PUT", path, handlers...)
}
func (g *RouteGroup) DELETE(path string, handlers ...contracts.Middleware) {
	g.Handle("DELETE", path, handlers...)
}
func (g *RouteGroup) PATCH(path string, handlers ...contracts.Middleware) {
	g.Handle("PATCH", path, handlers...)
}

// Context implements dghttp.Context using Gin.
type Context struct {
	ginCtx *gin.Context
}

func (c *Context) Request() context.Context   { return c.ginCtx.Request.Context() }
func (c *Context) Param(key string) string    { return c.ginCtx.Param(key) }
func (c *Context) Query(key string) string    { return c.ginCtx.Query(key) }
func (c *Context) Bind(obj interface{}) error { return c.ginCtx.ShouldBind(obj) }
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return c.ginCtx.FormFile(name)
}
func (c *Context) MultipartForm() (*multipart.Form, error) { return c.ginCtx.MultipartForm() }
func (c *Context) JSON(code int, obj interface{})          { c.ginCtx.JSON(code, obj) }
func (c *Context) Status(code int)                         { c.ginCtx.Status(code) }
func (c *Context) Next()                                   { c.ginCtx.Next() }
func (c *Context) Abort()                                  { c.ginCtx.Abort() }
func (c *Context) AbortWithStatusJSON(code int, obj interface{}) {
	c.ginCtx.AbortWithStatusJSON(code, obj)
}

func wrapMiddleware(m contracts.Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		m(&Context{ginCtx: c})
	}
}
