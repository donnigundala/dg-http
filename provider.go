package http

import (
	"net/http"

	"github.com/donnigundala/dg-core/contracts/foundation"
	"github.com/gin-gonic/gin"
)

// HttpServiceProvider implements the PluginProvider interface.
type HttpServiceProvider struct {
	// Config is auto-injected by dg-core if using config tags
	// or manually provided.
	Config Config `config:"http"`
}

// Name returns the plugin name.
func (p *HttpServiceProvider) Name() string {
	return "http"
}

// Version returns the plugin version.
func (p *HttpServiceProvider) Version() string {
	return "1.0.0"
}

// Dependencies returns the list of dependencies.
func (p *HttpServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the HTTP services into the container.
func (p *HttpServiceProvider) Register(app foundation.Application) error {
	// 1. Register the Router (Gin Engine) IF NOT already present
	if _, err := app.Make("router"); err != nil {
		app.Singleton("router", func() (interface{}, error) {
			return NewRouter(), nil
		})
	}

	// 2. Register the Kernel IF NOT already present
	if _, err := app.Make("kernel"); err != nil {
		app.Singleton("kernel", func() (interface{}, error) {
			// Resolve the router/engine we just registered (or that was already there)
			routerInterface, err := app.Make("router")
			if err != nil {
				return nil, err
			}

			engine := routerInterface.(http.Handler)
			return NewKernel(app, engine.(*gin.Engine)), nil
		})
	}

	// 3. Register the Server as the main plugin instance
	// This allows the framework to call Start() and Stop() automatically
	app.Singleton("http", func() (interface{}, error) {
		routerInterface, err := app.Make("router")
		if err != nil {
			return nil, err
		}

		handler := routerInterface.(http.Handler)
		return NewHTTPServer(p.Config, handler), nil
	})

	return nil
}

// Boot boots the HTTP services.
func (p *HttpServiceProvider) Boot(app foundation.Application) error {
	return nil
}
