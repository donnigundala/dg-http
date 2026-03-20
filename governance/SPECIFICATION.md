# dg-http — Constitutional Specification

**Module:** dg-http  
**Document Type:** Constitutional Specification  
**Governance Level:** Kernel-Adjacent  
**Status:** Authoritative (v1.8.0 Aligned)  
**Effective Date:** 2026-03-11  
**Governing Authority:** dg Governance Authority

---

## 1. Overview
dg-http provides the **Abstracted Transport Contract** for the dg ecosystem. It defines the "Law" of HTTP interaction—how modules register routes and handle requests—without owning any implementation authority.

### Contract Classification: **Pure Plugin (Type A)**
dg-http is a **Pure Contract package**. It contains zero implementation logic, zero background goroutines, and zero server lifecycle ownership.

---

## 2. Authority & Prohibitions
> [!IMPORTANT]
> To preserve Kernel supremacy and Application sovereignty, dg-http is strictly prohibited from:
> - Containing a `foundation.ServiceProvider`.
> - Including engine-specific adapters (e.g., Gin, Fiber) in its root or contracts.
> - Managing network listeners or server startup.
> - Loading, parsing, or managing configuration (ENV, YAML, etc.).
> - Depending on third-party HTTP libraries (Standard Library only).

---

## 3. Core Contracts

### 3.1 Controller interface
Modules must implement the `Controller` interface to bind endpoints to the system.
```go
type Controller interface {
    RegisterRoutes(router Router)
}
```

### 3.2 Abstract Context (v1.8.0)
The `Context` interface provides a request-scoped boundary that includes first-class observability slots. handlers MUST use these slots for logging and tracing to ensure ecosystem-wide correlation.
- `ctx.Logger()`: Returns the request-scoped logger.
- `ctx.Tracer()`: Returns the request-scoped tracer.

---

## 4. Middleware Policy
Middleware in the dg-ecosystem follows the principle of **Adapter Sovereignty**:
- **Bridge Middleware**: (Allowed) Statelessly connects runtime signals (RequestID, Tracing) to Core capabilities.
- **Policy Middleware**: (Allowed) Enforces transport-level behavior (CORS, Limits). MUST be explicitly opt-in.
- **Implementation**: Middleware is **NOT portable** across adapters. Switching engines (e.g., Gin -> Fiber) requires re-implementation.

---

## 5. Configuration Strategy
dg-http assumes all configuration is **already resolved** by the Application (Skeleton). 
- Adapters MUST receive configuration explicitly (e.g., `Port`, `Timeout`).
- Adapters MUST NOT infer behavior from the environment or global state.

---

## 6. Implementation Responsibilities
The **Application Layer (Skeleton)** or an **Implementation Provider** (e.g., `dg-provider-gin`) is responsible for:
1. Instantiating the concrete HTTP engine.
2. Satisfying the `dg-http` capability surface via a Service Provider.
3. Managing the asynchronous server lifecycle.
