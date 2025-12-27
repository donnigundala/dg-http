# dg-http

> The official HTTP plugin for the dg-framework, providing a production-ready, Gin-native web stack.

This package was decoupled from `dg-core` to enable a more modular architecture. It provides the HTTP Kernel implementation, production middlewares, and utilities for handling requests and responses.

## Features

- 🚀 **Gin-Native Router**: High-performance routing with full access to the Gin ecosystem.
- 🛡️ **Production Middlewares**: 
  - CORS, Logger, Recovery
  - Security Headers, Rate Limiting
  - Request ID, Body Size Limit, Gzip Compression
- ✅ **Request/Response Helpers**: Standardized JSON responses and request binding utilities.
- 📁 **File Uploads**: Optimized multipart file upload handlers with validation.
- 🏥 **Health Checks**: Integrated health monitoring endpoints.

## Installation

```bash
go get github.com/donnigundala/dg-http
```

## Usage

### 1. Register as a Service Provider
The easiest way to use `dg-http` is by registering its service provider. This automatically binds the router and kernel to the container.

```go
package main

import (
    "github.com/donnigundala/dg-core/foundation"
    "github.com/donnigundala/dg-http"
)

func main() {
    app := foundation.New(".")
    
    // 1. Register the HTTP provider
    app.Register(dghttp.NewHttpServiceProvider())
    
    // 2. Start (starts the HTTP server)
    app.Start() 
}
```

### Integration via InfrastructureSuite
In your `bootstrap/app.go`, the HTTP provider is typically part of the `FrameworkSuite` or registered directly:

```go
func (a *Application) registerProviders() error {
	return a.foundation.Register(
		providers.NewFrameworkServiceProvider(), // Often contains dghttp
		providers.NewAppServiceProvider(),
	)
}
```

## Configuration

The plugin uses the `http` key in your configuration file.

### Configuration Mapping (YAML vs ENV)

| YAML Key | Environment Variable | Default | Description |
| :--- | :--- | :--- | :--- |
| `http.enabled` | `HTTP_ENABLED` | `true` | Enable HTTP server |
| `http.addr` | `HTTP_ADDR` | `:8080` | Server address |
| `http.read_timeout` | `HTTP_READ_TIMEOUT` | `30s` | Read timeout |
| `http.write_timeout` | `HTTP_WRITE_TIMEOUT` | `30s` | Write timeout |
| `http.idle_timeout` | `HTTP_IDLE_TIMEOUT` | `60s` | Connection idle timeout |

### Middleware Configuration

Middlewares are typically configured via the `HttpKernel` in your application.

```go
func (k *Kernel) Middlewares() []dghttp.Middleware {
    return []dghttp.Middleware{
        dghttp.CORSMiddleware(),
        dghttp.GzipMiddleware(),
    }
}
```

## 📊 Observability

`dg-http` is instrumented with OpenTelemetry metrics. If `dg-observability` is registered and enabled, the following metrics are automatically emitted:

*   `http_server_request_count_total`: Counter (labels: `http_method`, `http_route`, `http_status_code`)
*   `http_server_request_duration_milliseconds`: Histogram (labels: `http_method`, `http_route`, `http_status_code`)

To enable observability, ensure the `dg-observability` plugin is registered and configured:

```yaml
observability:
  enabled: true
  service_name: "my-app"
```

The metrics are collected via a middleware that is automatically applied when using the standard HTTP Kernel or can be applied manually:

```go
router.Use(dghttp.Observability())
```

## Contributing
This package is part of the `dg-framework` monorepo. Please refer to the main repository for contributing guidelines.
