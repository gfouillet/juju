# UTC Date Audit Documentation

This directory contains the results of a comprehensive audit of UTC date handling in the Juju domain layer.

## Quick Start

1. **Start here:** Read [UTC_AUDIT_QUICK_REFERENCE.md](./UTC_AUDIT_QUICK_REFERENCE.md) for a high-level overview
2. **For details:** Read [UTC_DATE_AUDIT_REPORT.md](./UTC_DATE_AUDIT_REPORT.md) for complete analysis
3. **For tracking:** Use the CSV files to track remediation progress

## File Guide

### 📄 Main Reports

#### [UTC_DATE_AUDIT_REPORT.md](./UTC_DATE_AUDIT_REPORT.md)
**Size:** 23KB | **Lines:** 691

The comprehensive audit report containing:
- Executive summary with key statistics
- Common patterns for UTC date handling (correct and incorrect)
- Detailed analysis of SQL schema date fields
- Complete list of code violations with examples
- Good examples for reference
- Risk assessment
- Recommendations prioritized by urgency
- Domain-by-domain status breakdown

**Best for:** Understanding the full scope of the issue and detailed technical context.

---

#### [UTC_AUDIT_QUICK_REFERENCE.md](./UTC_AUDIT_QUICK_REFERENCE.md)
**Size:** 5KB | **Lines:** 179

A condensed reference guide containing:
- Summary statistics
- Critical violations by domain
- Common pattern examples (correct vs incorrect)
- Quick fix templates
- Prioritized recommendations
- Testing strategy

**Best for:** Quick lookup during development and code reviews.

---

### 📊 Tracking Files (CSV)

#### [UTC_AUDIT_VIOLATIONS.csv](./UTC_AUDIT_VIOLATIONS.csv)
**Records:** 28 violations

Columns:
- `File Path` - Exact file location
- `Line Context` - Function/context where violation occurs
- `Violation Type` - Type of missing UTC conversion
- `Severity` - HIGH or MEDIUM
- `Current Code` - Current violating code
- `Fixed Code` - Corrected code with .UTC()
- `Domain` - Juju domain (application, machine, etc.)
- `Layer` - state or service

**Best for:** Tracking remediation progress and creating fix PRs.

---

#### [UTC_AUDIT_SCHEMA_FIELDS.csv](./UTC_AUDIT_SCHEMA_FIELDS.csv)
**Records:** 58 date fields

Columns:
- `Schema File` - SQL schema file path
- `Table Name` - Database table
- `Field Name` - Column name
- `Field Type` - DATETIME, TIMESTAMP, or DATE
- `Has UTC Default` - YES or NO
- `Code Layer Verification` - Code enforcement status
- `Status` - OK, VIOLATION, or NEEDS_VERIFICATION
- `Notes` - Additional context

**Best for:** Understanding database schema coverage and planning schema improvements.

---

## Audit Scope

The audit covered:
- ✅ 24 SQL schema files in `domain/schema`
- ✅ 58 date/timestamp field definitions
- ✅ 360 Go files in state layer (`domain/*/state`)
- ✅ 91 Go files with time handling
- ✅ Service layer time patterns (`domain/*/service`)

## Key Findings

### Critical Statistics

| Metric | Count | Severity |
|--------|-------|----------|
| State layer: `clock.Now()` without UTC | 19 | HIGH |
| State layer: `time.Now()` without UTC | 6 | HIGH |
| Service layer violations | 5 | HIGH |
| Schema fields without UTC default | 45 | MEDIUM |
| **Total violations** | **75** | - |

### Domains Requiring Fixes (Prioritized)

| Priority | Domain | Violations | Impact |
|----------|--------|------------|--------|
| **HIGH** | resource | 6 | Resource tracking timestamps |
| **HIGH** | application | 4 | Application lifecycle timestamps |
| **HIGH** | machine | 2 | Machine agent timestamps |
| **HIGH** | status | 2 | Status "last seen" tracking |
| **HIGH** | access | 1 | User login tracking |
| MEDIUM | cloudimagemetadata | 3 | Image metadata timestamps |
| MEDIUM | relation | 2 | Relation update tracking |
| MEDIUM | changestream | 1 | Change stream timestamps |
| MEDIUM | crossmodelrelation | 1 | Cross-model updates |

### Exemplary Domains ✓

These domains demonstrate correct UTC handling:
- **operation** - Consistent `time.Now().UTC()` throughout
- **removal** - All removal services use `clock.Now().UTC()`
- **secret** - Proper UTC handling in state layer

## Pattern Reference

### ✓ CORRECT Patterns

```go
// Pattern 1: Direct UTC conversion
EnqueuedAt: time.Now().UTC()

// Pattern 2: Using injected clock with UTC
Since: ptr(st.clock.Now().UTC())

// Pattern 3: Store UTC for reuse
now := time.Now().UTC()
createdAt := now
updatedAt := now
```

### ⚠ VIOLATION Patterns

```go
// Missing .UTC() with clock - WRONG
UpdatedAt: st.clock.Now()  // FIX: Add .UTC()

// Missing .UTC() with time - WRONG
lastSeen := time.Now()  // FIX: Add .UTC()

// Missing .UTC() in pointer - WRONG
now := ptr(st.clock.Now())  // FIX: Add .UTC()
```

## Recommendations

### Phase 1: Immediate Fixes (2-3 days)
- Fix all 25 state layer violations
- Fix all 5 service layer violations
- Target: 30 line changes across ~20 files

### Phase 2: Infrastructure (1-2 days)
- Add linting rules to catch missing `.UTC()`
- Add to CI/CD pipeline
- Create test utilities for UTC verification

### Phase 3: Best Practices (0.5 days)
- Update `CODING.md` with UTC requirements
- Add code review checklist item
- Document exceptions (if any)

### Phase 4: Long-term (Ongoing)
- Add UTC defaults to new schema fields
- Monitor for new violations
- Train team on UTC requirements

## Risk Assessment

**Current Risk Level:** MEDIUM to HIGH

**Why this matters:**
- Juju is a distributed system that may run across different timezones
- Inconsistent date storage leads to:
  - Failed time comparisons
  - Incorrect time-based operations (rotation, expiration, scheduling)
  - Debugging difficulties
  - Data corruption if server timezone changes

**Mitigation:**
1. Fix violations (Phase 1)
2. Add automated checks (Phase 2)
3. Document requirements (Phase 3)
4. Ensure UTC server deployment (operational requirement)

## How to Use These Files

### For Developers
1. Review [UTC_AUDIT_QUICK_REFERENCE.md](./UTC_AUDIT_QUICK_REFERENCE.md) before writing date-related code
2. Check patterns when touching files in [UTC_AUDIT_VIOLATIONS.csv](./UTC_AUDIT_VIOLATIONS.csv)
3. Use fix templates from quick reference

### For Code Reviewers
1. Reference quick guide for correct patterns
2. Verify all `clock.Now()` and `time.Now()` calls include `.UTC()`
3. Check schema fields have UTC defaults where appropriate

### For Project Managers
1. Review executive summary in main report
2. Use violation CSV to track remediation progress
3. Prioritize HIGH severity domains first

### For Architects
1. Review complete report for architectural implications
2. Consider linting rules and automated enforcement
3. Plan schema improvements for new tables

## Next Steps

1. ✅ Audit complete - Documentation created
2. ⏭️ Review audit findings with team
3. ⏭️ Create fix PRs for HIGH priority violations
4. ⏭️ Implement linting rules
5. ⏭️ Update coding guidelines
6. ⏭️ Add UTC verification tests

## Questions?

- See section 1 in main report for pattern explanations
- See section 3 in main report for detailed violation examples
- See section 4 in main report for good examples to follow
- See section 6 in main report for implementation recommendations

---

**Audit Date:** 2026-01-05  
**Auditor:** GitHub Copilot Agent  
**Scope:** Juju domain layer (domain/*)
