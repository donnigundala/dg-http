package http

import (
	"net/http"

	"github.com/donnigundala/dg-core/contracts/foundation"
	httpContract "github.com/donnigundala/dg-core/contracts/http"
	"github.com/gin-gonic/gin"
)

var _ httpContract.Kernel = (*Kernel)(nil)

// Kernel is the concrete implementation of the HTTP kernel.
type Kernel struct {
	app        foundation.Application
	engine     *gin.Engine
	middleware []gin.HandlerFunc
}

// NewKernel creates a new Kernel instance.
func NewKernel(app foundation.Application, engine *gin.Engine) *Kernel {
	return &Kernel{
		app:        app,
		engine:     engine,
		middleware: []gin.HandlerFunc{
			// Default global middleware can be added here
		},
	}
}

// GetMiddleware returns the global middleware stack.
func (k *Kernel) GetMiddleware() []gin.HandlerFunc {
	return k.middleware
}

// Handle handles the incoming HTTP request.
func (k *Kernel) Handle(w http.ResponseWriter, r *http.Request) {
	// Serve the request through Gin engine
	k.engine.ServeHTTP(w, r)
}

// ServeHTTP satisfies the http.Handler interface.
func (k *Kernel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k.Handle(w, r)
}
