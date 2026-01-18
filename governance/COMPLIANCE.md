# dg-http — Compliance

**Module:** dg-http  
**Document Type:** Compliance Declaration  
**Governance Level:** Kernel-Adjacent  
**Status:** Compliant (Path A – Adapter-Scoped Middleware)  
**Change Policy:** Informational (Derived from Specification)  
**Effective Date:** 2026-01-19  
**Last Reviewed:** 2026-01-19  
**Reviewed By:** Antigravity (AI Auditor)  
**Governing Authority:** dg Governance Authority

---

## Authority Statement
`dg-http` is a **contract-first transport module** with declarative capability descriptors.
It contains zero server implementation code, zero background listeners,
and zero dependency on a concrete HTTP engine within the contract surface.

## Verdict: ✅ PASS
The package is officially admitted as a kernel-compliant transport contract
with adapter-scoped, non-portable middleware permitted by specification.

---

## Audit Checkpoints

| Checkpoint | Status | Notes |
| :--- | :--- | :--- |
| **Kernel Supremacy** | ✅ PASS | Declarative provider exists; lifecycle and runtime remain owned by Skeleton. |
| **No Abstraction Leak** | ✅ PASS | Gin is fully isolated within adapter; contracts expose only abstract Router. |
| **No Autonomous Logic** | ✅ PASS | All background server startup code deleted. |
| **Transport Boundary** | ✅ PASS | Contracts are engine-agnostic; adapters implement concrete behavior. |
| **Middleware Portability** | ✅ PASS | Adapter middleware is explicitly non-portable and opt-in (Path A). |

---

> [!NOTE]
> Adapter middleware included in dg-http adapters is **non-portable by design**.
> Behavioral equivalence across different HTTP engines is not guaranteed.
> See SPECIFICATION.md and INJECTION_CONTRACT.md for binding and authority rules.