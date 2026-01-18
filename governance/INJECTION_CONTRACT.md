# dg-http — Injection Contract

**Module:** dg-http  
**Document Type:** Injection Contract  
**Governance Level:** Kernel-Adjacent  
**Status:** Authoritative  
**Change Policy:** Requires Re-certification  
**Effective Date:** 2026-01-19  
**Last Reviewed:** 2026-01-19  
**Governing Authority:** dg Governance Authority

---

## Contract Category

**Type:** Type A – Cross-cutting Transport Capability

dg-http represents an optional transport abstraction.
Applications MAY operate without an HTTP transport (e.g., CLI tools, workers).
This contract defines expressive capability only and does not require a concrete runtime implementation.

---

## 1. Purpose
This document defines the interaction boundary for HTTP transport. 

**Authority:** The kernel/application governs the server lifecycle. `dg-http` defines the **interfaces only**.

---

## 2. Core Interfaces

### 2.1 Route Registration
Modules must implement the `Controller` interface to register their routes.

```go
type Controller interface {
    RegisterRoutes(router Router)
}
```

### 2.2 Router Abstraction
The implementer (Skeleton) provides an object satisfying the `Router` interface.

```go
type Router interface {
    Group(prefix string) Router
    Handle(method, path string, handlers ...Middleware)
}
```

---

## 3. Explicit Prohibitions

- ❌ **No Provider**: The package must not contain a `foundation.ServiceProvider`.
- ❌ **No Engine Leakage**: Do not expose `*gin.Engine` in the contract.
- ❌ **No Auto-Startup**: The plugin must not start any server listeners.
- ❌ **No Global State**: No package-level router instances.

---

## 4. Implementation Responsibilities
The **Application Layer (Skeleton)** is responsible for:
1. Creating the actual Gin engine.
2. Wrapping Gin to satisfy the `dghttp.Router` interface.
3. Starting the server asynchronously and managing its lifecycle.
4. Injecting the `Router` into modules during their registration phase.
5. Skeleton adapters MUST NOT extend, redefine, or influence dg-http contract semantics.

---

## 5. Compliance
Failure to adhere to this contract results in immediate rejection by the **Compliance Gauntlet**.
