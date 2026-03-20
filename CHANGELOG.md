# Changelog

All notable changes to this project will be documented in this file.

## 2026-03-11

### Added
- **v1.8.0 Alignment**: Fully aligned with the hardened `dg-core v1.8.0` kernel boundaries.
- **Observability Integration**: Added `Logger()` and `Tracer()` slots to the `Context` interface.
- **Ghost Test**: Verified zero-leakage integration for observability across core boundaries.
- **Consolidated Governance**: Merged `INJECTION_CONTRACT.md`, `MIDDLEWARE_POLICY.md`, and `CONFIGURATION.md` into `SPECIFICATION.md`.

### Changed
- **Pure Plugin Transformation**: Officially designated as a **Type A (Pure Contract)** plugin.
- **Purification**: Removed all internal implementations and adapters to enforce perfect isolation.
- **Documentation**: Simplified governance into `SPECIFICATION.md` and `CERTIFICATION.md`.

## 2026-01-07

### Added
- **Compliance Milestone**: Fully compliant with `dg-core` v1.2.0 Kernel Authority.
- **Abstracted Transport**: Introduced `Router`, `RouteGroup`, and `Context` interfaces to hide engine implementation.
- **Compliance Documentation**: Added `SPECIFICATION.md`, `INJECTION_CONTRACT.md`, and `COMPLIANCE.md`.

### Changed
- **Architectural Shift**: Transformed into an **Absolute Zero Contract** package.
- **Provider Removal**: Deleted `provider.go`. No longer participates in lifecycle or container registration.
- **The Great Purge**: Removed 20+ implementation files (Gin engine management, server runners, middleware logic).
- **Semantic Alignment**: Transport errors aligned with `dg-core` kernel errors.
