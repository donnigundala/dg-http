package dghttp

import (
	"github.com/donnigundala/dg-core/contracts/foundation"
	"github.com/gin-gonic/gin"
)

// HttpServiceProvider implements the PluginProvider interface.
type HttpServiceProvider struct {
	// Config is auto-injected by dg-core if using config tags
	// or manually provided.
	Config Config `config:"http"`
}

// NewHttpServiceProvider creates a new HTTP service provider.
func NewHttpServiceProvider() *HttpServiceProvider {
	return &HttpServiceProvider{}
}

// Name returns the plugin name.
func (p *HttpServiceProvider) Name() string {
	return Binding
}

// Version returns the plugin version.
func (p *HttpServiceProvider) Version() string {
	return Version
}

// Dependencies returns the list of dependencies.
func (p *HttpServiceProvider) Dependencies() []string {
	return []string{}
}

// Register registers the HTTP services into the container.
func (p *HttpServiceProvider) Register(app foundation.Application) error {
	// 1. Register a default Router (Gin Engine) IF NOT already present.
	if _, err := app.Make("router"); err != nil {
		app.Singleton("router", func() (interface{}, error) {
			return NewRouter(), nil
		})
	}

	// 2. Register the Server as the main plugin instance IF enabled.
	if p.Config.Enabled {
		app.Singleton(Binding, func() (interface{}, error) {
			routerInterface, err := app.Make("router")
			if err != nil {
				return nil, err
			}

			// Use the application logger (slog.Logger) wrapped in our defaultLogger
			loggerInstance := &defaultLogger{Logger: app.Log()}

			return NewHTTPServer(p.Config, routerInterface.(*gin.Engine), WithHTTPLogger(loggerInstance)), nil
		})
	}

	return nil
}

// Boot boots the HTTP services.
func (p *HttpServiceProvider) Boot(app foundation.Application) error {
	// Try to resolve the kernel and bootstrap it
	if instance, err := app.Make("kernel"); err == nil {
		if kernel, ok := instance.(interface{ Bootstrap() error }); ok {
			if err := kernel.Bootstrap(); err != nil {
				return err
			}
		}
	}
	return nil
}
