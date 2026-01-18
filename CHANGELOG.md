# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-07

### Added
- **Compliance Milestone**: Fully compliant with `dg-core` v1.2.0 Kernel Authority.
- **Abstracted Transport**: Introduced `Router`, `RouteGroup`, and `Context` interfaces to hide engine implementation.
- **Compliance Documentation**: Added `SPECIFICATION.md`, `INJECTION_CONTRACT.md`, and `COMPLIANCE.md`.

### Changed
- **Architectural Shift**: Transformed into an **Absolute Zero Contract** package.
- **Provider Removal**: Deleted `provider.go`. No longer participates in lifecycle or container registration.
- **The Great Purge**: Removed 20+ implementation files (Gin engine management, server runners, middleware logic).
- **Semantic Alignment**: Transport errors aligned with `dg-core` kernel errors.

### [Legacy State] - 2025-12-27 (Pre-Gauntlet)
- Initial implementation with internal Gin management (deprecated by Gauntlet 3).

[1.0.0]: https://github.com/donnigundala/dg-http/releases/tag/v1.0.0
