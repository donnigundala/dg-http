# dg-http — Configuration Policy

**Module:** dg-http  
**Document Type:** Configuration Policy  
**Governance Level:** Kernel-Adjacent  
**Status:** Authoritative  
**Change Policy:** Requires Re-certification  
**Effective Date:** 2026-01-19  
**Last Reviewed:** 2026-01-19  
**Governing Authority:** dg Governance Authority

---

## Purpose

This document defines how configuration is handled in relation to dg-http.
Its goal is to preserve **kernel authority**, **application ownership**, and
**adapter neutrality** by explicitly forbidding configuration ownership inside dg-http.

---

## Core Rule

> **dg-http MUST NOT own, load, parse, or manage configuration.**

All configuration belongs to the **Application Layer (Skeleton)** or higher-level modules.

---

## Scope of dg-http Configuration

dg-http itself:
- Does NOT read environment variables
- Does NOT read files
- Does NOT parse configuration formats
- Does NOT define configuration schemas
- Does NOT apply defaults

dg-http MAY:
- Accept already-resolved values via constructor arguments
- Accept configuration through dependency injection
- Forward configuration values to adapters without interpretation

---

## Adapter Configuration Rules

Adapters MUST follow these rules:

- Adapters MUST NOT load configuration from environment variables or files
- Adapters MUST NOT define global configuration state
- Adapters MUST receive configuration explicitly from the application
- Adapters MUST treat configuration as immutable input

Allowed examples:
- HTTP address passed as a string
- Timeouts passed as `time.Duration`
- Boolean feature flags passed explicitly

Forbidden examples:
- Reading `PORT`, `HOST`, or `GIN_MODE`
- Loading `.env` or `.yaml` files
- Applying hidden defaults
- Inferring behavior from runtime environment

---

## Middleware Configuration

Middleware configuration:
- MUST be explicit
- MUST be passed by the application
- MUST NOT be auto-derived from environment or global state

Middleware MUST NOT:
- Read environment variables
- Perform configuration parsing
- Change behavior implicitly based on runtime context

---

## Application Responsibilities

The application (Skeleton) is responsible for:

- Loading configuration
- Parsing configuration formats
- Validating configuration values
- Injecting configuration into dg-http adapters
- Owning configuration lifecycle and mutation

dg-http assumes configuration is **already resolved**.

---

## Rationale

Configuration ownership is a form of authority.

Allowing dg-http or its adapters to load configuration would:
- Transfer authority away from the application
- Create hidden coupling
- Break deterministic behavior
- Complicate testing and certification

dg-http therefore enforces strict configuration neutrality.

---

## Related Governance Documents

- SPECIFICATION.md
- INJECTION_CONTRACT.md
- COMPLIANCE.md
- MIDDLEWARE_POLICY.md
- CERTIFICATION.md