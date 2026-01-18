package dghttp

import (
	"github.com/donnigundala/dg-core/contracts/foundation"
	"github.com/donnigundala/dg-http/contracts"
)

// HttpServiceProvider implements a sovereign http capability surface.
//
// DESIGN NOTE: This provider is purely declarative. It should be configured
// via functional setters (WithRouter) before registration. It is NOT intended
// to be resolved from the container directly.
type HttpServiceProvider struct {
	router contracts.Router
}

// NewHttpServiceProvider creates a new instance of the provider.
func NewHttpServiceProvider() *HttpServiceProvider {
	return &HttpServiceProvider{
		router: contracts.NewNoopRouter(),
	}
}

// WithRouter injects a concrete router implementation into the capability surface.
// NOTE: Must be called before Register().
func (p *HttpServiceProvider) WithRouter(router contracts.Router) *HttpServiceProvider {
	p.router = router
	return p
}

func (p *HttpServiceProvider) Name() string    { return Binding }
func (p *HttpServiceProvider) Version() string { return Version }

func (p *HttpServiceProvider) Dependencies() []string {
	return []string{}
}

// Register installs the http router into the container.
func (p *HttpServiceProvider) Register(app foundation.Application) error {
	app.Singleton(RouterBinding, func() (interface{}, error) {
		return p.router, nil
	})
	return nil
}

func (p *HttpServiceProvider) Boot(app foundation.Application) error     { return nil }
func (p *HttpServiceProvider) Shutdown(app foundation.Application) error { return nil }
