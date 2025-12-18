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
import (
    "github.com/donnigundala/dg-http"
    "github.com/donnigundala/dg-core/foundation"
)

func main() {
    app := foundation.New(".")
    
    // Register the provider
    app.Register(&http.HttpServiceProvider{})
    
    app.Boot()
    
    // Resolve the engine if needed
    router := app.Make("router").(*gin.Engine)
    router.Run(":8080")
}
```

### 2. Manual Initialization
If you prefer full control, you can still initialize the components manually:

```go
router.Use(http.LoggerWithDefault())
router.Use(http.RecoveryWithDefault())
router.Use(http.CORSWithDefault())
```

## Contributing
This package is part of the `dg-framework` monorepo. Please refer to the main repository for contributing guidelines.
