# Complete Prompt: Fix UTC Date/Time Violations in Juju Domain Layer

## Mission
Fix UTC enforcement violations in the Juju codebase where dates are stored in the database without explicit UTC conversion. This ensures consistent timezone handling across distributed deployments.

## Context & Problem
Juju is a distributed orchestration system that runs across multiple timezones. The database schema contains 58 date/timestamp fields, and a comprehensive audit found violations where `clock.Now()` and `time.Now()` are used without `.UTC()` conversion before storing in the database.

**Impact of not fixing:**
- Time comparison failures across deployments
- Incorrect expiration/rotation logic
- Debugging issues in distributed systems
- Data corruption when server timezone changes

## What to Fix

### Pattern to Apply
**WRONG (Before):**
```go
UpdatedAt: st.clock.Now()        // Stores in local timezone
lastSeen := time.Now()            // Stores in local timezone
now := ptr(st.clock.Now())       // Stores in local timezone
```

**CORRECT (After):**
```go
UpdatedAt: st.clock.Now().UTC()      // Always stores in UTC
lastSeen := time.Now().UTC()          // Always stores in UTC
now := ptr(st.clock.Now().UTC())     // Always stores in UTC
```

## Files to Modify (HIGH Priority - 25 violations across 15 files)

### 1. Application Domain State (5 violations)

**File: `domain/application/state/peer_relation.go`**
- **Line ~184:** `UpdatedAt: st.clock.Now(),`
- **Fix:** `UpdatedAt: st.clock.Now().UTC(),`
- **Context:** `insertNewRelationStatus` function, relation status timestamp

**File: `domain/application/state/unit.go`**
- **Line ~879:** `now := ptr(st.clock.Now())`
- **Fix:** `now := ptr(st.clock.Now().UTC())`
- **Context:** CAAS unit creation, unit status timestamp

**File: `domain/application/state/resource.go`**
- **Line ~47:** `now := st.clock.Now()`
- **Fix:** `now := st.clock.Now().UTC()`
- **Context:** `buildResourcesToAdd` function, resource creation timestamp
- **Line ~280:** `potentialResources[i].CreatedAt = time.Now()`
- **Fix:** `potentialResources[i].CreatedAt = time.Now().UTC()`
- **Context:** Converting pending to potential resources

**File: `domain/application/state/unitstate.go`**
- **Line ~1227:** `statusTime := time.Now()`
- **Fix:** `statusTime := time.Now().UTC()`
- **Context:** Filesystem status pending
- **Line ~1361:** `statusTime := time.Now()`
- **Fix:** `statusTime := time.Now().UTC()`
- **Context:** Volume status pending

### 2. Application Domain Service (2 violations)

**File: `domain/application/service/unit.go`**
- **Line ~346:** `now := ptr(s.clock.Now())`
- **Fix:** `now := ptr(s.clock.Now().UTC())`
- **Context:** `makeUnitStatusArgs` function, unit status timestamps

**File: `domain/application/service/service.go`**
- **Line ~139:** `Since: ptr(s.clock.Now()),`
- **Fix:** `Since: ptr(s.clock.Now().UTC()),`
- **Context:** `recordCreateMachineStatusHistory` function, machine status

### 3. Machine Domain State (2 violations)

**File: `domain/machine/state/state.go`**
- **Line ~885:** `AgentStartedAt: st.clock.Now(),`
- **Fix:** `AgentStartedAt: st.clock.Now().UTC(),`
- **Context:** `SetMachineCloudInstance` function, agent started timestamp

**File: `domain/machine/state/placement.go`**
- **Line ~208:** `now := clock.Now()`
- **Fix:** `now := clock.Now().UTC()`
- **Context:** `insertMachine` function, machine placement timestamp

### 4. Machine Domain Service (1 violation)

**File: `domain/machine/service/service.go`**
- **Line ~620:** `Since: ptr(clock.Now()),`
- **Fix:** `Since: ptr(clock.Now().UTC()),`
- **Context:** `recordCreateMachineStatusHistory` function, status history

### 5. Resource Domain State (6 violations)

**File: `domain/resource/state/resource.go`**
- **Line ~994:** `AddedAt: st.clock.Now(),`
- **Fix:** `AddedAt: st.clock.Now().UTC(),`
- **Context:** Unit resource link creation

- **Line ~1389:** `now := st.clock.Now()`
- **Fix:** `now := st.clock.Now().UTC()`
- **Context:** `buildResourcesToAdd` function

- **Line ~1450:** `CreatedAt: st.clock.Now(),`
- **Fix:** `CreatedAt: st.clock.Now().UTC(),`
- **Context:** Upload resource creation

- **Line ~1624:** `CreatedAt: st.clock.Now(),`
- **Fix:** `CreatedAt: st.clock.Now().UTC(),`
- **Context:** Store resource creation

- **Line ~2055:** `createdAt = st.clock.Now()`
- **Fix:** `createdAt = st.clock.Now().UTC()`
- **Context:** Resource timestamp fallback

### 6. Status Domain State (2 violations)

**File: `domain/status/state/model/modelstate.go`**
- **Line ~1116:** `LastSeen: st.clock.Now(),`
- **Fix:** `LastSeen: st.clock.Now().UTC(),`
- **Context:** `RecordUnitPresence` function, unit last seen

- **Line ~1200:** `LastSeen: st.clock.Now(),`
- **Fix:** `LastSeen: st.clock.Now().UTC(),`
- **Context:** `RecordMachinePresence` function, machine last seen

### 7. Access Domain State (1 violation)

**File: `domain/access/state/user.go`**
- **Line ~770:** `CreatedAt: time.Now(),`
- **Fix:** `CreatedAt: time.Now().UTC(),`
- **Context:** `addUser` function, user creation timestamp

### 8. Access Domain Service (1 violation)

**File: `domain/access/service/user.go`**
- **Line ~378:** `if err := s.st.UpdateLastModelLogin(ctx, name, modelUUID, time.Now()); err != nil {`
- **Fix:** `if err := s.st.UpdateLastModelLogin(ctx, name, modelUUID, time.Now().UTC()); err != nil {`
- **Context:** `UpdateLastModelLogin` function, login timestamp

### 9. Operation Domain State (1 violation)

**File: `domain/operation/state/prune.go`**
- **Line ~88:** `expiresAt := expires{At: time.Now().Add(-age)}`
- **Fix:** `expiresAt := expires{At: time.Now().UTC().Add(-age)}`
- **Context:** `PruneOperations` function, expiration calculation

## Implementation Steps

1. **For each file listed above:**
   - Locate the exact line with the violation
   - Add `.UTC()` to the `clock.Now()` or `time.Now()` call
   - Verify the change maintains existing code structure

2. **Verification commands:**
   ```bash
   # Check no violations remain (should return nothing)
   grep -r "clock\.Now()" domain/application/state domain/machine/state domain/resource/state domain/status/state domain/access/state domain/operation/state --include="*.go" | grep -v "\.UTC()" | grep -v "_test.go"
   
   grep -r "time\.Now()" domain/application/state domain/access/state domain/operation/state --include="*.go" | grep -v "\.UTC()" | grep -v "_test.go"
   
   grep -r "clock\.Now()" domain/application/service domain/machine/service --include="*.go" | grep -v "\.UTC()" | grep -v "_test.go"
   ```

3. **Build verification:**
   ```bash
   # Build all modified packages
   go build ./domain/application/state/...
   go build ./domain/application/service/...
   go build ./domain/machine/state/...
   go build ./domain/machine/service/...
   go build ./domain/resource/state/...
   go build ./domain/status/state/...
   go build ./domain/access/state/...
   go build ./domain/access/service/...
   go build ./domain/operation/state/...
   ```

4. **Test verification (optional but recommended):**
   ```bash
   # Run tests for modified packages
   go test ./domain/application/state/...
   go test ./domain/machine/state/...
   go test ./domain/resource/state/...
   # etc.
   ```

## Commit Strategy

Group changes by domain for easier review:

```bash
# Example commit messages:
"fix(application): add UTC conversion to timestamp fields in state and service layers"
"fix(machine): add UTC conversion to timestamp fields in state and service layers"
"fix(resource): add UTC conversion to all timestamp fields in state layer"
"fix(status): add UTC conversion to LastSeen timestamps"
"fix(access): add UTC conversion to user timestamp fields"
"fix(operation): add UTC conversion to operation expiration calculation"
```

Or as a single commit:
```bash
"fix(domain): add UTC conversion to all HIGH priority timestamp fields

Fix 25 HIGH priority UTC violations across state and service layers:
- Application domain: 7 fixes (5 state + 2 service)
- Machine domain: 3 fixes (2 state + 1 service)
- Resource domain: 6 fixes (all state)
- Status domain: 2 fixes (state)
- Access domain: 2 fixes (1 state + 1 service)
- Operation domain: 1 fix (state)

All clock.Now() and time.Now() calls now include .UTC() to ensure
dates are stored in UTC regardless of server timezone."
```

## Important Notes

- **DO NOT modify test files** (*_test.go) - they are excluded from this audit
- **DO NOT change** files with existing correct UTC usage
- The fix is mechanical: simply add `.UTC()` to each violation
- Changes are backward compatible - if server is already in UTC, behavior is unchanged
- Risk: LOW - adding .UTC() is safe for database storage

## Success Criteria

✅ All 25 violations fixed (21 lines changed across 14 files)
✅ All modified packages build successfully  
✅ No new violations introduced
✅ Code follows existing patterns in the repository

## Expected Outcome

- 14 files modified
- 21 lines changed (adding `.UTC()`)
- All dates now stored in UTC regardless of server timezone
- Prevents time-related bugs in distributed deployments

## Reference Examples

**Good patterns (from exemplary code in the repo):**

From `domain/operation/state/start.go`:
```go
now := time.Now().UTC()
err = st.insertOperation(ctx, tx, insertOperation{
    EnqueuedAt: now,
    // ...
})
```

From `domain/removal/service/*.go`:
```go
s.clock.Now().UTC().Add(wait)
```

**Before/After Example:**

```go
// BEFORE (domain/application/state/peer_relation.go)
func (st *State) insertNewRelationStatus(ctx context.Context, tx *sqlair.TX, uuid corerelation.UUID) error {
    status := setRelationStatus{
        RelationUUID: uuid,
        Status:       corestatus.Joining,
        UpdatedAt:    st.clock.Now(),  // ⚠ MISSING .UTC()
    }
    // ...
}

// AFTER
func (st *State) insertNewRelationStatus(ctx context.Context, tx *sqlair.TX, uuid corerelation.UUID) error {
    status := setRelationStatus{
        RelationUUID: uuid,
        Status:       corestatus.Joining,
        UpdatedAt:    st.clock.Now().UTC(),  // ✓ FIXED
    }
    // ...
}
```

## Summary

This is a straightforward mechanical fix to ensure timezone consistency. Simply locate each line listed above and add `.UTC()` to the time call. The changes are safe, backward compatible, and critical for correct operation in distributed environments.

**Total effort:** 2-3 hours for careful implementation and verification
**Files:** 14 files
**Lines changed:** 21 lines
**Risk:** Low (additive change, backward compatible)
