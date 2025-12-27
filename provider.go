package dghttp

import (
	"reflect"

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

			var loggerInstance Logger
			// Try to resolve logger from container
			if log, err := app.Make("logger"); err == nil {
				// Adapt the logger to our Logger interface
				if adapted, ok := log.(interface {
					Debug(msg string, args ...interface{})
					Info(msg string, args ...interface{})
					Warn(msg string, args ...interface{})
					Error(msg string, args ...interface{})
				}); ok {
					loggerInstance = &loggerAdapter{logger: adapted}
				}
			}

			return NewHTTPServer(p.Config, routerInterface.(*gin.Engine), WithHTTPLogger(loggerInstance)), nil
		})
	}

	return nil
}

// loggerAdapter adapts a generic logger to http.Logger interface.
type loggerAdapter struct {
	logger interface {
		Debug(msg string, args ...interface{})
		Info(msg string, args ...interface{})
		Warn(msg string, args ...interface{})
		Error(msg string, args ...interface{})
	}
}

func (l *loggerAdapter) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *loggerAdapter) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *loggerAdapter) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *loggerAdapter) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l *loggerAdapter) With(args ...interface{}) Logger {
	// Try to call With(args...) via reflection to support different return types
	v := reflect.ValueOf(l.logger)
	m := v.MethodByName("With")
	if m.IsValid() {
		valArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			valArgs[i] = reflect.ValueOf(arg)
		}
		results := m.Call(valArgs)
		if len(results) == 1 {
			if nextLogger, ok := results[0].Interface().(interface {
				Debug(msg string, args ...interface{})
				Info(msg string, args ...interface{})
				Warn(msg string, args ...interface{})
				Error(msg string, args ...interface{})
			}); ok {
				return &loggerAdapter{logger: nextLogger}
			}
		}
	}
	return l
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
