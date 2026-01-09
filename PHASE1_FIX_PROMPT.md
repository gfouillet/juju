# Agent Prompt: Fix HIGH Priority UTC Violations (Phase 1)

## Objective

Fix all HIGH severity UTC violations in the Juju domain layer by adding `.UTC()` to all `clock.Now()` and `time.Now()` calls that store dates in the database.

## Context

A comprehensive UTC audit identified 75 violations where dates may be stored without UTC conversion. Phase 1 focuses on the 25 HIGH priority violations in the state and service layers that directly impact database storage.

**Why this matters:**
- Juju runs in distributed environments across different timezones
- Without UTC conversion, dates are stored in the server's local timezone
- This causes: time comparison failures, incorrect expiration/rotation logic, and debugging issues

## Files to Fix

### HIGH Priority Violations (25 total)

#### State Layer (20 violations)

**Application Domain (5 violations):**
1. `domain/application/state/peer_relation.go` - Line ~137
   - Current: `UpdatedAt: st.clock.Now()`
   - Fixed: `UpdatedAt: st.clock.Now().UTC()`

2. `domain/application/state/unit.go`
   - Current: `now := ptr(st.clock.Now())`
   - Fixed: `now := ptr(st.clock.Now().UTC())`

3. `domain/application/state/resource.go`
   - Current: `now := st.clock.Now()`
   - Fixed: `now := st.clock.Now().UTC()`

4. `domain/application/state/unitstate.go` (2 instances)
   - Current: `= time.Now()`
   - Fixed: `= time.Now().UTC()`

**Machine Domain (2 violations):**
5. `domain/machine/state/state.go`
   - Current: `AgentStartedAt: st.clock.Now()`
   - Fixed: `AgentStartedAt: st.clock.Now().UTC()`

6. `domain/machine/state/placement.go`
   - Current: `now := clock.Now()`
   - Fixed: `now := clock.Now().UTC()`

**Resource Domain (5 violations):**
7-10. `domain/resource/state/resource.go` (4 clock.Now() violations)
    - Current: `AddedAt: st.clock.Now()`
    - Fixed: `AddedAt: st.clock.Now().UTC()`
    
    - Current: `now := st.clock.Now()`
    - Fixed: `now := st.clock.Now().UTC()`
    
    - Current: `CreatedAt: st.clock.Now()` (2 instances)
    - Fixed: `CreatedAt: st.clock.Now().UTC()`

11. `domain/resource/state/resource.go` (1 time.Now() violation)
    - Current: `time.Now()`
    - Fixed: `time.Now().UTC()`

**Status Domain (2 violations):**
12-13. `domain/status/state/model/modelstate.go` (2 instances)
    - Current: `LastSeen: st.clock.Now()`
    - Fixed: `LastSeen: st.clock.Now().UTC()`

**Access Domain (1 violation):**
14. `domain/access/state/user.go`
    - Current: `time.Now()`
    - Fixed: `time.Now().UTC()`

**Operation Domain (1 violation):**
15. `domain/operation/state/prune.go`
    - Current: `expires{At: time.Now().Add(-age)}`
    - Fixed: `expires{At: time.Now().UTC().Add(-age)}`

#### Service Layer (5 violations)

**Access Service (1 violation):**
16. `domain/access/service/user.go`
    - Current: `time.Now()`
    - Fixed: `time.Now().UTC()`

**Application Service (2 violations):**
17. `domain/application/service/unit.go`
    - Current: `now := ptr(s.clock.Now())`
    - Fixed: `now := ptr(s.clock.Now().UTC())`

18. `domain/application/service/service.go`
    - Current: `Since: ptr(s.clock.Now())`
    - Fixed: `Since: ptr(s.clock.Now().UTC())`

**Machine Service (1 violation):**
19. `domain/machine/service/service.go`
    - Current: `Since: ptr(clock.Now())`
    - Fixed: `Since: ptr(clock.Now().UTC())`

**Note:** There's also 1 MEDIUM priority violation in crossmodelrelation/service that could optionally be included.

## Correct Patterns to Follow

### Pattern 1: Direct UTC conversion
```go
// CORRECT
EnqueuedAt: time.Now().UTC()
UpdatedAt: st.clock.Now().UTC()
```

### Pattern 2: Store UTC for reuse
```go
// CORRECT
now := time.Now().UTC()
createdAt := now
updatedAt := now
```

### Pattern 3: UTC with pointer
```go
// CORRECT
now := ptr(st.clock.Now().UTC())
```

### Pattern 4: UTC with operations
```go
// CORRECT
expires := time.Now().UTC().Add(-age)
```

## Examples from Exemplary Code

Reference these files for correct UTC patterns:
- `domain/operation/state/start.go` - Consistent `time.Now().UTC()` throughout
- `domain/operation/state/task.go` - Multiple correct UTC usages
- `domain/removal/service/*.go` - All removal services use UTC correctly

## Instructions

1. **Find and fix each violation:**
   - For each file listed above, locate the exact line with the violation
   - Add `.UTC()` to the `clock.Now()` or `time.Now()` call
   - Ensure the fix maintains the existing code structure

2. **Verify the changes:**
   - Build the affected packages to ensure no compilation errors
   - Run existing tests for the modified files
   - Use `grep -r "clock\.Now()" domain/*/state --include="*.go" | grep -v "\.UTC()"` to verify no new violations

3. **Testing strategy:**
   - Run unit tests for each modified package
   - Focus on state and service layer tests
   - Example: `go test ./domain/application/state/...`

4. **DO NOT change:**
   - Test files (*_test.go) - these are excluded from the audit
   - Files with existing correct UTC usage
   - MEDIUM priority violations (unless specifically requested)

5. **Commit strategy:**
   - Group changes by domain for easier review
   - Example commit messages:
     - `fix(application): add UTC conversion to timestamp fields in state layer`
     - `fix(machine): add UTC conversion to timestamp fields in state and service layers`
     - `fix(resource): add UTC conversion to all timestamp fields`
     - `fix(status): add UTC conversion to LastSeen timestamps`
     - `fix(access): add UTC conversion to user timestamps`

## Reference Files

For detailed information, see:
- **Complete audit report:** `UTC_DATE_AUDIT_REPORT.md`
- **Quick reference:** `UTC_AUDIT_QUICK_REFERENCE.md`
- **Violations CSV:** `UTC_AUDIT_VIOLATIONS.csv` (rows 2-24 are HIGH priority)

## Success Criteria

- All 25 HIGH priority violations fixed
- All modified packages build successfully
- All existing tests pass
- No new violations introduced
- Code follows existing patterns in the repository

## Estimated Effort

- Time: 2-3 days
- Changes: ~30 lines across 15 files
- Risk: LOW (adding .UTC() is backward compatible when storing in UTC)

## Example Fix

**Before:**
```go
// domain/application/state/peer_relation.go
func (st *State) insertNewRelationStatus(ctx context.Context, tx *sqlair.TX, uuid corerelation.UUID) error {
    status := setRelationStatus{
        RelationUUID: uuid,
        Status:       corestatus.Joining,
        UpdatedAt:    st.clock.Now(),  // ⚠ MISSING .UTC()
    }
    // ... rest of function
}
```

**After:**
```go
// domain/application/state/peer_relation.go
func (st *State) insertNewRelationStatus(ctx context.Context, tx *sqlair.TX, uuid corerelation.UUID) error {
    status := setRelationStatus{
        RelationUUID: uuid,
        Status:       corestatus.Joining,
        UpdatedAt:    st.clock.Now().UTC(),  // ✓ FIXED
    }
    // ... rest of function
}
```

## Notes

- The fix is mechanical: simply add `.UTC()` to each violation
- This ensures all dates are stored in UTC regardless of server timezone
- The change is backward compatible - if the server is already in UTC, behavior is unchanged
- If the server is in a different timezone, dates will now correctly be converted to UTC before storage

## Commands to Run

```bash
# Find all HIGH priority violations that need fixing
grep -E "(domain/application|domain/machine|domain/resource|domain/status|domain/access|domain/operation)" UTC_AUDIT_VIOLATIONS.csv | grep "HIGH"

# After fixes, verify no violations remain
grep -r "clock\.Now()" domain/application/state domain/machine/state domain/resource/state domain/status/state domain/access/state domain/operation/state --include="*.go" | grep -v "\.UTC()" | grep -v "_test.go"

# Run tests for modified packages
go test ./domain/application/state/...
go test ./domain/application/service/...
go test ./domain/machine/state/...
go test ./domain/machine/service/...
go test ./domain/resource/state/...
go test ./domain/status/state/...
go test ./domain/access/state/...
go test ./domain/access/service/...
go test ./domain/operation/state/...
```
