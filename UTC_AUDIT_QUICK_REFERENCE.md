# UTC Date Audit - Quick Reference

## Summary

This document provides a quick reference for the UTC date audit findings. For full details, see [UTC_DATE_AUDIT_REPORT.md](./UTC_DATE_AUDIT_REPORT.md).

## Key Statistics

- **Total date fields in schemas:** 58
- **Fields with UTC defaults:** 13 (22%)
- **Fields without UTC defaults:** 45 (78%)
- **State layer violations:** 25 files
  - `clock.Now()` without UTC: 19 instances
  - `time.Now()` without UTC: 6 instances
- **Service layer violations:** 5 files

## Critical Violations by Domain

| Domain | Files | Violations | Priority |
|--------|-------|------------|----------|
| application | 5 | 4 | HIGH |
| resource | 1 | 6 | HIGH |
| machine | 2 | 2 | HIGH |
| status | 1 | 2 | HIGH |
| access | 2 | 1 | HIGH |
| relation | 2 | 2 | MEDIUM |
| cloudimagemetadata | 2 | 3 | MEDIUM |
| changestream | 2 | 1 | MEDIUM |
| crossmodelrelation | 1 | 1 | MEDIUM |

## Common Patterns

### ✓ CORRECT Patterns

```go
// Pattern 1: Direct UTC conversion
EnqueuedAt: time.Now().UTC()

// Pattern 2: Using injected clock with UTC
Since: ptr(st.clock.Now().UTC())

// Pattern 3: Store UTC for reuse
now := time.Now().UTC()
```

### ⚠ VIOLATION Patterns

```go
// Missing .UTC() with clock
UpdatedAt: st.clock.Now()  // WRONG

// Missing .UTC() with time
lastSeen := time.Now()  // WRONG

// Missing .UTC() with clock in pointer
now := ptr(st.clock.Now())  // WRONG
```

## Files Requiring Fixes

### State Layer (19 violations)

1. `domain/application/state/peer_relation.go` - 1 violation
2. `domain/application/state/unit.go` - 1 violation
3. `domain/application/state/resource.go` - 1 violation
4. `domain/changestream/state/state.go` - 1 violation
5. `domain/cloudimagemetadata/state/state.go` - 3 violations
6. `domain/crossmodelrelation/state/model/remoteapplication.go` - 1 violation
7. `domain/machine/state/state.go` - 1 violation
8. `domain/machine/state/placement.go` - 1 violation
9. `domain/relation/state/subordinateunit.go` - 1 violation
10. `domain/relation/state/relation.go` - 1 violation
11. `domain/resource/state/resource.go` - 5 violations
12. `domain/status/state/model/modelstate.go` - 2 violations

### Service Layer (5 violations)

1. `domain/access/service/user.go` - 1 violation
2. `domain/application/service/unit.go` - 1 violation
3. `domain/application/service/service.go` - 1 violation
4. `domain/crossmodelrelation/service/remoteapplication.go` - 1 violation
5. `domain/machine/service/service.go` - 1 violation

### Additional violations (time.Now())

1. `domain/access/state/user.go` - 1 violation
2. `domain/application/state/unitstate.go` - 2 violations
3. `domain/application/state/resource.go` - 1 violation
4. `domain/operation/state/prune.go` - 1 violation

## Exemplary Code (For Reference)

### Operation Domain ✓
- `domain/operation/state/start.go` - Perfect UTC usage
- `domain/operation/state/task.go` - Consistent UTC throughout

### Removal Service ✓
- All files in `domain/removal/service/*` - Consistent UTC usage

### Secret State ✓
- `domain/secret/state/state.go` - Proper UTC handling

## Quick Fix Template

### For clock.Now() violations:

```diff
- UpdatedAt: st.clock.Now(),
+ UpdatedAt: st.clock.Now().UTC(),
```

### For time.Now() violations:

```diff
- lastLogin := time.Now()
+ lastLogin := time.Now().UTC()
```

### For pointer violations:

```diff
- now := ptr(st.clock.Now())
+ now := ptr(st.clock.Now().UTC())
```

## Recommendations

### Immediate (HIGH Priority)
1. Fix all 19 `clock.Now()` violations in state layer
2. Fix all 6 `time.Now()` violations in state layer
3. Fix all 5 violations in service layer

### Short-term (MEDIUM Priority)
4. Add UTC defaults to schema fields where applicable
5. Document intentional non-UTC fields

### Long-term (Best Practices)
6. Add linting rule to catch missing `.UTC()`
7. Add tests to verify UTC storage
8. Update coding guidelines with UTC requirement

## Testing Verification

After fixes, verify with:

```go
// Test that timestamps are stored in UTC
func TestTimestampIsUTC(t *testing.T) {
    // Set non-UTC timezone
    os.Setenv("TZ", "America/New_York")
    
    // Create record with timestamp
    // ...
    
    // Verify stored time is UTC
    assert.Equal(t, "UTC", storedTime.Location().String())
}
```

## Risk Assessment

**Risk Level:** MEDIUM to HIGH

**Impact if not fixed:**
- Inconsistent date storage across deployments
- Time comparison failures
- Incorrect time-based operations (expiration, rotation)
- Debugging difficulties in distributed systems

## Estimated Effort

- **Fixes:** 2-3 days (25-30 line changes across 20 files)
- **Linting rules:** 1 day
- **Guidelines update:** 0.5 days
- **Total:** ~4 days

---

For detailed analysis, code examples, and complete findings, see [UTC_DATE_AUDIT_REPORT.md](./UTC_DATE_AUDIT_REPORT.md).
