> **Sealed:** dg-http is governance-complete. All core contracts, provider, and governance documents are certified and frozen. Changes require explicit re-certification.
# dg-http (Sovereign Transport)

[![Compliance: ✅](https://img.shields.io/badge/Compliance-✅_PASS-green)](governance/COMPLIANCE.md)

`dg-http` is a **Certified Sovereign Plugin** providing the authoritative HTTP transport capability for the DG Framework. It follows the **Capability Surface Pattern**, decoupling the application container from concrete HTTP engines.

## 🏛️ Sovereign Architecture

`dg-http` is structured as a multi-module repository to enforce perfect isolation between authority and infrastructure.

- **`contracts/` (Module)**: The authoritative root. Defines interfaces for `Router`, `Context`, and `Middleware`. 100% dependency-free.
- **`adapters/`**: Houses concrete bridge implementations.
    - **`adapters/gin` (Module)**: The official Gin-to-DG adapter.
- **`provider.go`**: The Capability Surface. A declarative Service Provider that registers the router slot into the container using **Typed Binding Constants**.

---

## 🛠️ Capability Category: Type A

`dg-http` is a **Type A (Cross-cutting) Capability**. It provides expressive transport power but is strictly optional for the system to function.

- **Silent No-op**: Includes a `NewNoopRouter()` in the `contracts` module.
- **Non-Invasive**: CLI tasks, batch jobs, and workers can exclude the HTTP stack entirely without code changes or registration panics.

---

## 🚀 Usage

### ⚙️ 1. Application Layer (Bootstrap)

The application layer (Skeleton) owns the engine and binds it to the sovereign surface.

```go
import (
    dghttp "github.com/donnigundala/dg-http"
    dggin "github.com/donnigundala/dg-http/adapters/gin"
    "github.com/gin-gonic/gin"
)

func Register(app foundation.Application) {
    // 1. Initialize the engine (Application Layer owns this)
    engine := gin.Default()

    // 2. Configure the Sovereign Surface
    provider := dghttp.NewHttpServiceProvider().
        WithRouter(dggin.NewRouter(engine))

    // 3. Register
    app.Register(provider)
}
```

### 📦 2. Module Layer (Consumption)

Modules depend ONLY on the `contracts` module, keeping them vendor-blind.

```go
import "github.com/donnigundala/dg-http/contracts"

func (c *UserController) RegisterRoutes(router contracts.Router) {
    router.GET("/users", c.List)
}
```

### 🛠️ 3. Adapter Middleware (Helpers)

The Gin adapter provides common middleware helpers that are **sovereign-aware** (resolve observability from the container).

```go
import dggin "github.com/donnigundala/dg-http/adapters/gin"

func (k *Kernel) globalMiddleware() []gin.HandlerFunc {
    return []gin.HandlerFunc{
        dggin.RequestID(),                    // Stateless helper
        dggin.Tracing(k.app, "my-service"),   // Resolves Tracer from container
        dggin.Logger(k.app),                   // Resolves Logger from container
    }
}
```

---

## 📜 Governance

- **Zero Infrastructure in Root**: The root `dg-http` module MUST NOT import `gin` or any other engine.
- **Adapter Isolation**: All engine-specific code belongs in `adapters/`.
- **Typed Bindings**: Use `dghttp.RouterBinding` for container resolution.

> **Note:** dg-http is governance-complete and should be treated as a sealed transport contract.
> Any changes to the core contracts, provider, or governance documents require explicit re-certification.

---
**Standard**: [Sovereign Plugin Governance Blueprint](../../dg-core/docs/GOVERNANCE_BLUEPRINT.md)
