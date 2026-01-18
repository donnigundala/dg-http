

# dg-http — Middleware Policy

**Module:** dg-http  
**Document Type:** Middleware Policy  
**Governance Level:** Kernel-Adjacent  
**Status:** Authoritative  
**Change Policy:** Requires Re-certification  
**Effective Date:** 2026-01-19  
**Last Reviewed:** 2026-01-19  
**Governing Authority:** dg Governance Authority

---

## Purpose

This document defines the rules, scope, and limitations of middleware provided by dg-http adapters.
Its goal is to preserve **transport sovereignty**, **adapter isolation**, and **explicit application control**.

Middleware policy is intentionally strict to prevent abstraction drift and hidden authority transfer.

---

## Core Principle

> **Middleware is adapter-scoped, non-portable, and opt-in by design.**

dg-http standardizes **contracts and intent**, not middleware implementations.

---

## Classification of Middleware

### 1. Bridge Middleware (Allowed)

Bridge middleware connects runtime signals to dg-core capabilities.

Examples:
- Request ID propagation
- Structured logging
- Distributed tracing
- Panic recovery (runtime safety only)

Characteristics:
- Stateless
- No policy decisions
- Resolves dependencies via the application container
- Does not modify business semantics

Bridge middleware MAY exist inside adapters.

---

### 2. Policy Middleware (Allowed, with constraints)

Policy middleware enforces HTTP-level behavior.

Examples:
- CORS
- Security headers
- Body size limits
- Rate limiting (transport-level only)

Rules:
- MUST be explicitly opt-in
- MUST NOT be auto-wired
- MUST NOT assume cross-adapter portability
- MUST NOT encode application-specific policy

Policy middleware MAY exist inside adapters as convenience helpers.

---

### 3. Forbidden Middleware

The following middleware types are NOT allowed inside dg-http adapters:

- Authentication or authorization logic
- Business rules
- Application-specific routing decisions
- Persistence or network side effects
- Configuration parsing from environment or files
- Lifecycle ownership (startup, shutdown, signals)

Such middleware belongs to the **Application Layer (Skeleton)** or higher-level modules.

---

## Portability Rule (Critical)

> **Middleware implementations are NOT portable across HTTP adapters.**

Switching HTTP engines (e.g., Gin → Fiber) REQUIRES:
- Reimplementation of middleware
- Explicit reassessment of behavior

dg-http makes **no guarantee** of behavioral equivalence across adapters.

---

## Adapter Author Responsibilities

Adapter authors MUST:
- Treat middleware as engine-specific
- Keep middleware fully contained within the adapter
- Avoid leaking engine semantics through contracts
- Document middleware behavior clearly

Adapter authors MUST NOT:
- Share middleware across adapters
- Introduce hidden defaults
- Assume execution order guarantees outside the adapter

---

## Application Responsibilities

Applications MUST:
- Explicitly opt into adapter middleware
- Own middleware ordering
- Own policy selection and configuration

dg-http does not manage middleware composition.

---

## Rationale

Attempting to create portable middleware across HTTP engines leads to:
- Leaky abstractions
- Lowest-common-denominator behavior
- Hidden coupling to execution models

dg-http intentionally avoids this trap.

---

## Related Governance Documents

- SPECIFICATION.md
- INJECTION_CONTRACT.md
- COMPLIANCE.md
- CERTIFICATION_REPORT.md