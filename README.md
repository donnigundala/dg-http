# dg-http (Constitutional Contracts)

[![Compliance: ✅](https://img.shields.io/badge/Compliance-✅_PASS-green)](governance/CERTIFICATION.md)

`dg-http` is a **Pure Contract Plugin (Type A)** providing the authoritative HTTP transport interfaces for the dg ecosystem. It defines the "Universal Protocol" for request handling, routing, and observability integration without owning any infrastructure.

---

## 🏛️ Sovereign Architecture

To enforce perfect isolation between authority and infrastructure, `dg-http` contains **zero implementation logic**.

- **`contracts/`**: The authoritative root. Defines interfaces for `Router`, `Context`, `Middleware`, and `Controller`. 
- **`governance/`**: Contains the [Constitutional Specification](governance/SPECIFICATION.md) and [Certification](governance/CERTIFICATION.md).
- **Zero-Dependency**: Depends strictly on `dg-core` and the Go Standard Library.

---

## 🚀 v1.8.0 Alignment: Observability Slots

In the v1.8.0 hardening, the `Context` interface was upgraded to include first-class observability slots. This ensures that any HTTP handler can perform logging and tracing that is automatically correlated by the underlying engine.

```go
func (c *UserController) GetProfile(ctx contracts.Context) {
    // Zero-casting access to request-scoped observability
    ctx.Logger().Info("Fetching profile")
    
    _, span := ctx.Tracer().Start(ctx.Request(), "db.query")
    defer span.End()
    
    // ... logic
}
```

---

## 🛠️ Usage Tier: Pure Plugin

`dg-http` is a **Type A (Cross-cutting) Capability**. It is strictly optional and provides a silent `Noop` fallback.

1. **Modules**: Depend only on `github.com/dgframe/dg-http/contracts`.
2. **Implementation**: Provided by sovereign adapters like `dg-provider-gin` or `dg-provider-fiber`.
3. **Skeleton**: Binds a concrete implementation to the system during bootstrap.

---

## 📜 Governance

- **Zero Infrastructure**: This module MUST NOT import any HTTP engine (Gin, Fiber, etc.).
- **Adapter Independence**: Implementation details (timeouts, server limits, TLS) belong to the implementation provider, not the contract.
- **Sealed Status**: `dg-http` is considered a sealed transport contract. Changes require explicit re-certification by the Governance Authority.

---
**Standard**: [Sovereign Plugin Governance Blueprint](../../core/docs/05-GOVERNANCE/GOVERNANCE_MODEL.md)
