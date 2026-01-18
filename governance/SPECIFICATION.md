# dg-http — Specification

**Module:** dg-http  
**Document Type:** Specification  
**Governance Level:** Kernel-Adjacent  
**Status:** Authoritative  
**Change Policy:** Requires Re-certification  
**Effective Date:** 2026-01-19  
**Last Reviewed:** 2026-01-19  
**Governing Authority:** dg Governance Authority

---

## Overview
dg-HTTP provides an **Abstracted Transport Contract** for the DG ecosystem. It defines how modules register routes and handle requests without depending on a specific HTTP engine or server implementation.

While dg-http is contract-first, it MAY expose declarative capability descriptors (e.g., providers or metadata)
that describe how an HTTP capability can be wired into the application container.
These descriptors MUST remain non-executable and MUST NOT instantiate or select any concrete HTTP engine.

## Contract Category

**Type:** Type A – Cross-cutting Transport Capability

dg-http represents an optional, cross-cutting transport abstraction.
Applications MAY run without an HTTP transport (e.g., CLI tools, workers, batch jobs).
Therefore, the contract defines expressive capability only and does not require a concrete runtime implementation.

---

## Authority Lock
> [!IMPORTANT]
> dg-http is a **contract-first package**.
> It MAY expose declarative capability descriptors, but it MUST NOT act as a runtime authority.
> It MUST NOT:
> - Initialize the Gin engine (e.g., `gin.New()`).
> - Start network listeners (e.g., `s.ListenAndServe()`).
> - Manage server timeouts or TLS configuration.
> - Register global middleware autonomously.
> - Call `os.Exit` or handle signals.
>
> Responsibility for the server lifecycle and engine initialization lies solely with the **Application Layer (Skeleton)**.
>
> Declarative descriptors provided by dg-http (such as providers or metadata)
> MUST NOT:
> - Select a concrete HTTP engine or adapter.
> - Instantiate router or server implementations.
> - Trigger side effects during registration.

## Non-Goals

dg-http explicitly does NOT aim to:

- Provide a production-ready HTTP server
- Select or recommend a routing engine
- Manage middleware ordering or execution policy
- Handle TLS, timeouts, or graceful shutdown
- Own or influence application lifecycle

---

## Components

1. **Router & RouteGroup Interfaces**: Generic wrappers that hide the underlying HTTP engine.
2. **Controller Contract**: A standard interface for modules to register their endpoints.
3. **Middleware Contract**: A standard way to intercept requests.
4. **Abstract Context**: A wrapper to prevent controllers from depending on `*gin.Context`.

---

## Implementation Guidelines
1. The **Application Layer** provides the concrete implementation (e.g., a GORM/Gin adapter).
2. The **Application Layer** owns the `main()` loop and server startup.
3. Interfaces must remain engine-agnostic to support future transport swaps (e.g., moving to Fiber or standard `net/http`).
4. Any HTTP engine adapter (Gin, net/http, Fiber, etc.) MUST live outside the contracts module.
Adapters MUST NOT influence contract evolution, provider behavior, or capability descriptors.
5. **Adapter Middleware Pattern**: Adapters MAY provide engine-specific middleware (e.g., `RequestID`, `Logger`) as convenience helpers. These MUST resolve cross-cutting dependencies (Logger, Tracer) strictly via the system container.

---

## Compliance
- Must contain zero implementation logic.
- Must have zero background goroutines.
- Must pass the **DG HTTP Compliance Checklist (COMPLIANCE.md)**.
