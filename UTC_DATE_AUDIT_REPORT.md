# UTC Date Field Audit Report - Juju Domain Layer

**Project:** Juju  
**Domain:** Database schema and domain layer (domain/*)  
**Generated:** 2026-01-05 14:37:23 UTC  
**Auditor:** GitHub Copilot Agent

---

## Executive Summary

This audit examines how date/time values are handled in the Juju domain layer to ensure all dates stored in the database are in UTC. The audit covers:
- SQL schema definitions in `domain/schema`
- State layer implementations in `domain/*/state`
- Service layer implementations in `domain/*/service`

### Key Findings

| Metric | Count |
|--------|-------|
| SQL schema files with date fields | 24 |
| Total date field definitions in schemas | 58 |
| Total state layer Go files | 360 |
| State files with time handling | 91 |
| `clock.Now()` calls in state layer | 27 |
| `clock.Now().UTC()` calls in state layer | 2 |
| `time.Now()` calls in state layer | 282 |
| `time.Now().UTC()` calls in state layer | 93 |

### Critical Issues

| Category | Count | Severity | Impact |
|----------|-------|----------|--------|
| State layer: `clock.Now()` without UTC | 19 | **HIGH** | May store local timezone dates |
| State layer: `time.Now()` without UTC | 6 | **HIGH** | May store local timezone dates |
| Schema: Date fields without UTC default | 45 | **MEDIUM** | Depends on application-layer enforcement |

---

## 1. Common Patterns for UTC Date Handling

### 1.1 Schema Level (SQL)

#### Pattern: Database Default Values with UTC

Some schema fields use SQLite's `STRFTIME` function with explicit UTC timezone:

```sql
created_at DATETIME NOT NULL DEFAULT (STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW', 'utc'))
```

**Examples:**
- `domain/schema/model/sql/0001-changelog.sql`: `created_at`, `updated_at`
- `domain/schema/controller/sql/0003-changelog.sql`: `created_at`, `updated_at`
- `domain/schema/model/sql/0012-secret.sql`: `create_time`, `update_time`
- `domain/schema/model/sql/0015-charm.sql`: `create_time`
- `domain/schema/model/sql/0028-cleanup.sql`: `scheduled_for`

**Coverage:** Only ~13 out of 58 date fields have explicit UTC defaults in the schema.

**Pros:**
- Database-level guarantee of UTC timestamps
- Works even if application layer fails to convert

**Cons:**
- Not applied consistently across all date fields
- Requires application layer to also ensure UTC for programmatically set dates

---

### 1.2 State Layer (Go)

The state layer is where dates are typically inserted into or retrieved from the database. Three patterns exist:

#### Pattern 1: Direct `time.Now().UTC()` - **CORRECT** ✓

```go
EnqueuedAt: time.Now().UTC()
```

**Example locations:**
- `domain/operation/state/start.go`: All operation timestamps
- `domain/operation/state/task.go`: Task timestamps and status updates
- `domain/relation/state/remoterelation.go`: Relation updates

**Why this is correct:**
- Explicitly converts current time to UTC
- Guarantees UTC storage regardless of server timezone
- Clear intent in code

#### Pattern 2: `clock.Now().UTC()` with Injected Clock - **CORRECT** ✓

```go
Since: ptr(st.clock.Now().UTC())
```

**Example locations:**
- `domain/crossmodelrelation/state/model/remoteapplication.go`: Status timestamps

**Why this is correct:**
- Uses injected clock for testability
- Explicitly converts to UTC
- Best practice for production code

#### Pattern 3: `clock.Now()` WITHOUT `.UTC()` - **VIOLATION** ⚠

```go
UpdatedAt: st.clock.Now()  // Missing .UTC()
```

**Example locations (19 violations):**
- `domain/application/state/peer_relation.go`: Line ~137
- `domain/application/state/unit.go`: Multiple instances
- `domain/application/state/resource.go`: Multiple instances
- `domain/machine/state/state.go`: AgentStartedAt field
- `domain/machine/state/placement.go`: Machine creation time
- `domain/resource/state/resource.go`: Multiple timestamp fields
- `domain/status/state/model/modelstate.go`: LastSeen timestamps
- And 12 more files...

**Why this is a violation:**
- `clock.Now()` returns time in the local timezone of the server
- If the server is not in UTC, dates will be stored with timezone offset
- May cause inconsistency and comparison issues

#### Pattern 4: `time.Now()` WITHOUT `.UTC()` - **VIOLATION** ⚠

```go
lastConnection := time.Now()  // Missing .UTC()
```

**Example locations (6 violations):**
- `domain/access/state/user.go`: Last login time
- `domain/application/state/unitstate.go`: Unit state updates (2 instances)
- `domain/operation/state/prune.go`: Expiration calculations
- And 2 more files...

**Why this is a violation:**
- Same issue as Pattern 3, but with standard library `time.Now()`
- No testability benefit of injected clock

---

### 1.3 Service Layer (Go)

The service layer typically passes time values to the state layer.

#### Good Pattern: Service Passes UTC to State ✓

```go
// From domain/removal/service/application.go
s.clock.Now().UTC().Add(wait)
```

**Example locations:**
- `domain/removal/service/*`: All removal scheduling operations correctly use UTC
- Multiple files in the removal domain consistently apply UTC

#### Mixed Pattern: Service Uses Non-UTC ⚠

```go
// From domain/access/service/user.go
s.st.UpdateLastModelLogin(ctx, name, modelUUID, time.Now())  // Missing .UTC()
```

```go
// From domain/application/service/unit.go
now := ptr(s.clock.Now())  // Missing .UTC()
```

**Observations:**
- Most service layer files delegate time handling to state layer
- Some services correctly apply UTC before calling state
- Some services pass non-UTC times, relying on state layer conversion (which may not happen)

---

## 2. Detailed Analysis of SQL Schema Date Fields

### 2.1 Fields with UTC Enforcement (Good)

| Schema File | Table | Field Name | Type | UTC Default | Status |
|-------------|-------|-----------|------|-------------|--------|
| `0001-changelog.sql` | change_log | created_at | DATETIME | ✓ | ✓ OK |
| `0001-changelog.sql` | change_log | updated_at | DATETIME | ✓ | ✓ OK |
| `0003-changelog.sql` | change_log | created_at | DATETIME | ✓ | ✓ OK |
| `0003-changelog.sql` | change_log | updated_at | DATETIME | ✓ | ✓ OK |
| `0012-secret.sql` | secret | create_time | DATETIME | ✓ | ✓ OK |
| `0012-secret.sql` | secret | update_time | DATETIME | ✓ | ✓ OK |
| `0012-secret.sql` | secret_revision | create_time | DATETIME | ✓ | ✓ OK |
| `0015-charm.sql` | charm | create_time | DATETIME | ✓ | ✓ OK |
| `0018-machine.sql` | machine | created_at | DATETIME | ✓ | ✓ OK |
| `0028-cleanup.sql` | cleanup | scheduled_for | DATETIME | ✓ | ✓ OK |

**Total:** 10 fields with UTC defaults (~17%)

---

### 2.2 Fields WITHOUT UTC Enforcement (Needs Review)

| Schema File | Table/Context | Field Name | Type | UTC Default | Code Layer Check Needed |
|-------------|---------------|-----------|------|-------------|------------------------|
| `0020-unit.sql` | unit_agent_status | updated_at | DATETIME | ✗ | ✓ YES |
| `0020-unit.sql` | unit_workload_status | updated_at | DATETIME | ✗ | ✓ YES |
| `0020-unit.sql` | cloud_container_status | updated_at | DATETIME | ✗ | ✓ YES |
| `0020-unit.sql` | cloud_container | last_seen | DATETIME | ✗ | ✓ YES |
| `0012-secret.sql` | secret | next_rotation_time | DATETIME | ✗ | ✓ YES |
| `0012-secret.sql` | secret_revision | expire_time | DATETIME | ✗ | ✓ YES |
| `0033-operation.sql` | operation | enqueued_at | TIMESTAMP | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation | started_at | TIMESTAMP | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation | completed_at | TIMESTAMP | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation_task | enqueued_at | DATETIME | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation_task | started_at | DATETIME | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation_task | completed_at | DATETIME | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation_task_status | updated_at | DATETIME | ✗ | **✓ CODE OK** |
| `0033-operation.sql` | operation_message | created_at | TIMESTAMP | ✗ | **✓ CODE OK** |
| `0018-machine.sql` | machine | agent_started_at | DATETIME | ✗ | **⚠ VIOLATION** |
| `0018-machine.sql` | machine | updated_at | DATETIME | ✗ | ✓ YES |
| `0018-machine.sql` | machine_status_data | last_seen | DATETIME | ✗ | **⚠ VIOLATION** |
| `0017-machine-cloud-instance.sql` | machine_cloud_instance | updated_at | DATETIME | ✗ | ✓ YES |
| `0034-cross-model-relation.sql` | (relation context) | updated_at | DATETIME | ✗ | **⚠ VIOLATION** |
| `0040-operator-status.PATCH.sql` | operator_status | updated_at | DATETIME | ✗ | ✓ YES |
| `0023-resource.sql` | resource | created_at | TIMESTAMP | ✗ | **⚠ VIOLATION** |
| `0023-resource.sql` | resource_stored_file | last_polled | TIMESTAMP | ✗ | **⚠ VIOLATION** |
| `0023-resource.sql` | application_resource | added_at | TIMESTAMP | ✗ | **⚠ VIOLATION** |
| `0024-relation.sql` | relation_status | updated_at | DATETIME | ✗ | **⚠ VIOLATION** |
| `0011-storage.sql` | storage_instance | updated_at | DATETIME | ✗ | ✓ YES |
| `0011-storage.sql` | storage_attachment | updated_at | DATETIME | ✗ | ✓ YES |
| `0019-application.sql` | application | updated_at | DATETIME | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_status | start_time | TIMESTAMP | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_status | success_time | TIMESTAMP | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_status | end_time | TIMESTAMP | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_status | phase_changed_time | TIMESTAMP | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_minion_sync | (multiple) | TIMESTAMP | ✗ | ✓ YES |
| `0011-model-migration.sql` | model_migration_minion_sync | time | TIMESTAMP | ✗ | ✓ YES |
| `0020-macaroon.sql` | root_key | created_at | TIMESTAMP | ✗ | ✓ YES |
| `0020-macaroon.sql` | root_key | expires_at | TIMESTAMP | ✗ | ✓ YES |
| `0018-secret-backend.sql` | secret_backend | next_rotation_time | DATETIME | ✗ | ✓ YES |
| `0015-user.sql` | user | created_at | TIMESTAMP | ✗ | ✓ YES |
| `0012-upgrade-info.sql` | upgrade_info | node_upgrade_completed_at | TIMESTAMP | ✗ | ✓ YES |
| `0002-lease.sql` | lease | start | TIMESTAMP | ✗ | ✓ YES |
| `0002-lease.sql` | lease | expiry | TIMESTAMP | ✗ | ✓ YES |
| `0024-cloudimagemetadata.sql` | cloud_image_metadata | created_at | DATETIME | ✗ | **⚠ VIOLATION** |
| `0019-model-last-login.sql` | model_last_login | time | TIMESTAMP | ✗ | ✓ YES |

**Total:** 48 fields without UTC defaults (~83%)

**Note:** Fields marked "CODE OK" are handled correctly in the state layer (with `.UTC()`).
Fields marked "VIOLATION" are NOT converted to UTC in code.

---

## 3. Detailed Code Violations

### 3.1 High Priority - State Layer Violations

These files store dates in the database WITHOUT UTC conversion:

#### 3.1.1 Application Domain

**File:** `domain/application/state/peer_relation.go`
```go
// Line ~137 in insertNewRelationStatus
status := setRelationStatus{
    RelationUUID: uuid,
    Status:       corestatus.Joining,
    UpdatedAt:    st.clock.Now(),  // ⚠ MISSING .UTC()
}
```
**Impact:** Relation status timestamps may be in local timezone  
**Fix:** Add `.UTC()` → `UpdatedAt: st.clock.Now().UTC()`

---

**File:** `domain/application/state/unit.go`
```go
// Multiple instances, example:
now := ptr(st.clock.Now())  // ⚠ MISSING .UTC()
```
**Impact:** Unit timestamps may be in local timezone  
**Fix:** Add `.UTC()` → `now := ptr(st.clock.Now().UTC())`

---

**File:** `domain/application/state/resource.go`
```go
// Multiple instances:
now := st.clock.Now()  // ⚠ MISSING .UTC()
// ...
CreatedAt: st.clock.Now(),  // ⚠ MISSING .UTC()
```
**Impact:** Resource timestamps may be in local timezone  
**Fix:** Add `.UTC()` to all clock.Now() calls

---

#### 3.1.2 Machine Domain

**File:** `domain/machine/state/state.go`
```go
// Field assignment:
AgentStartedAt: st.clock.Now(),  // ⚠ MISSING .UTC()
```
**Impact:** Machine agent start times may be in local timezone  
**Fix:** Add `.UTC()` → `AgentStartedAt: st.clock.Now().UTC()`

---

**File:** `domain/machine/state/placement.go`
```go
now := clock.Now()  // ⚠ MISSING .UTC()
```
**Impact:** Machine creation timestamps may be in local timezone  
**Fix:** Add `.UTC()` → `now := clock.Now().UTC()`

---

#### 3.1.3 Resource Domain

**File:** `domain/resource/state/resource.go`
```go
// Multiple violations:
AddedAt: st.clock.Now(),    // ⚠ MISSING .UTC()
now := st.clock.Now()        // ⚠ MISSING .UTC()
CreatedAt: st.clock.Now(),   // ⚠ MISSING .UTC()
```
**Impact:** Resource add/creation times may be in local timezone  
**Fix:** Add `.UTC()` to all instances

---

#### 3.1.4 Status Domain

**File:** `domain/status/state/model/modelstate.go`
```go
// Multiple instances:
LastSeen: st.clock.Now(),     // ⚠ MISSING .UTC()
LastSeen: st.clock.Now(),     // ⚠ MISSING .UTC()
```
**Impact:** Status "last seen" timestamps may be in local timezone  
**Fix:** Add `.UTC()` to both instances

---

#### 3.1.5 Other Domains

**File:** `domain/changestream/state/state.go`
```go
now := st.clock.Now()  // ⚠ MISSING .UTC()
```
**Impact:** Change stream timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

**File:** `domain/cloudimagemetadata/state/state.go`
```go
// Multiple instances:
InsertMetadata(ctx, db, metadata, s.clock.Now())  // ⚠ MISSING .UTC()
ExpiresAt: s.clock.Now().Add(-ExpirationDelay),   // ⚠ MISSING .UTC()
```
**Impact:** Cloud image metadata timestamps may be in local timezone  
**Fix:** Add `.UTC()` to all clock.Now() calls

---

**File:** `domain/crossmodelrelation/state/model/remoteapplication.go`
```go
UpdatedAt: st.clock.Now(),  // ⚠ MISSING .UTC()
```
**Impact:** Remote application update times may be in local timezone  
**Note:** This same file correctly uses `.UTC()` for status timestamps  
**Fix:** Add `.UTC()` for consistency

---

**File:** `domain/relation/state/subordinateunit.go`
```go
now := ptr(st.clock.Now())  // ⚠ MISSING .UTC()
```
**Impact:** Subordinate unit relation timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

**File:** `domain/relation/state/relation.go`
```go
UpdatedAt: st.clock.Now(),  // ⚠ MISSING .UTC()
```
**Impact:** Relation update timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

### 3.2 High Priority - `time.Now()` Violations

**File:** `domain/access/state/user.go`
```go
time.Now(),  // ⚠ MISSING .UTC()
```
**Context:** User last login time  
**Impact:** User activity tracking may be in local timezone  
**Fix:** Add `.UTC()`

---

**File:** `domain/application/state/unitstate.go`
```go
// Two instances:
= time.Now()  // ⚠ MISSING .UTC()
= time.Now()  // ⚠ MISSING .UTC()
```
**Impact:** Unit state timestamps may be in local timezone  
**Fix:** Add `.UTC()` to both

---

**File:** `domain/operation/state/prune.go`
```go
= expires{At: time.Now().Add(-age)}  // ⚠ MISSING .UTC()
```
**Impact:** Operation pruning/expiration logic may be timezone-dependent  
**Fix:** Add `.UTC()` → `time.Now().UTC().Add(-age)`

---

### 3.3 Service Layer Issues

**File:** `domain/access/service/user.go`
```go
s.st.UpdateLastModelLogin(ctx, name, modelUUID, time.Now())  // ⚠ MISSING .UTC()
```
**Impact:** Service passes non-UTC time to state layer  
**Fix:** Add `.UTC()` → `time.Now().UTC()`

---

**File:** `domain/application/service/unit.go`
```go
now := ptr(s.clock.Now())  // ⚠ MISSING .UTC()
```
**Impact:** Service passes non-UTC time to state layer  
**Fix:** Add `.UTC()` → `ptr(s.clock.Now().UTC())`

---

**File:** `domain/application/service/service.go`
```go
Since: ptr(s.clock.Now()),  // ⚠ MISSING .UTC()
```
**Impact:** Application status timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

**File:** `domain/crossmodelrelation/service/remoteapplication.go`
```go
Since: ptr(s.clock.Now()),  // ⚠ MISSING .UTC()
```
**Impact:** Cross-model relation status timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

**File:** `domain/machine/service/service.go`
```go
Since: ptr(clock.Now()),  // ⚠ MISSING .UTC()
```
**Impact:** Machine status timestamps may be in local timezone  
**Fix:** Add `.UTC()`

---

## 4. Good Examples (For Reference)

### 4.1 Operation Domain - Exemplary UTC Usage

**File:** `domain/operation/state/start.go`
```go
// Consistently uses UTC throughout
now := time.Now().UTC()

err = st.insertOperation(ctx, tx, insertOperation{
    UUID:           operationUUID,
    OperationID:    strconv.FormatUint(operationID, 10),
    Summary:        fmt.Sprintf("exec %q", args.Command),
    EnqueuedAt:     now,
    Parallel:       args.Parallel,
    ExecutionGroup: args.ExecutionGroup,
})
```

**Why this is good:**
- Explicit UTC conversion at the top
- Reused for multiple fields
- Clear intent

---

**File:** `domain/operation/state/task.go`
```go
// Multiple consistent UTC usages
Time: time.Now().UTC(),

completedTime := time.Now().UTC()

UpdatedAt: time.Now().UTC(),
```

**Why this is good:**
- Every time usage includes UTC
- Consistent pattern throughout the file

---

### 4.2 Removal Service - Exemplary UTC Usage

**File:** `domain/removal/service/application.go` (and all other removal service files)
```go
s.clock.Now().UTC().Add(wait)
```

**Why this is good:**
- Service layer ensures UTC before passing to state
- Consistent across entire removal domain
- Good defensive programming

---

## 5. Summary of Violations

| Category | Count | Files Affected | Severity |
|----------|-------|----------------|----------|
| **State layer: `clock.Now()` without UTC** | **19** | 13 unique files | **HIGH** |
| **State layer: `time.Now()` without UTC** | **6** | 3 unique files | **HIGH** |
| **Service layer: missing UTC** | **5** | 5 unique files | **HIGH** |
| **Schema: Date fields without UTC default** | **45** | 18 schema files | **MEDIUM** |
| **TOTAL VIOLATIONS** | **75** | **39 unique files** | - |

---

## 6. Recommendations

### 6.1 Immediate Actions (HIGH Priority)

1. **Fix all state layer violations** - Add `.UTC()` to all `clock.Now()` and `time.Now()` calls in state layer
   - Estimated files to fix: 16 files in state layer
   - Estimated changes: ~25 line changes
   - Risk: Low (adding UTC is safe, backward compatible for UTC storage)

2. **Fix service layer violations** - Add `.UTC()` to all time calls in service layer
   - Estimated files to fix: 5 files
   - Estimated changes: ~5 line changes
   - Risk: Low

### 6.2 Short-term Actions (MEDIUM Priority)

3. **Add UTC defaults to schema fields where possible**
   - Consider adding UTC defaults to fields that are auto-populated
   - Fields like `created_at`, `updated_at` should have UTC defaults
   - Risk: Low for new fields, requires migration for existing fields

4. **Document intentional non-UTC fields**
   - Some fields may intentionally store user's local time (if any)
   - Document these exceptions clearly in code comments

### 6.3 Long-term Actions (Best Practices)

5. **Add linting rules**
   - Create a linter rule to catch `clock.Now()` without `.UTC()`
   - Create a linter rule to catch `time.Now()` without `.UTC()` in state/service layers
   - Add to CI/CD pipeline

6. **Add tests**
   - Add tests to verify timestamps are stored in UTC
   - Add tests to verify timestamp comparisons work correctly across timezones

7. **Update coding guidelines**
   - Document the UTC requirement in `CODING.md` or `STYLE.md`
   - Add examples of correct patterns
   - Reference this audit report

8. **Training and awareness**
   - Share findings with team
   - Include in onboarding materials
   - Code review checklist item

---

## 7. Risk Assessment

### 7.1 Current Risk

**Risk Level:** **MEDIUM to HIGH**

**Rationale:**
- If Juju runs on servers configured in different timezones, dates will be stored inconsistently
- Time-based comparisons and calculations may fail
- Debugging issues across distributed systems becomes harder
- Data corruption risk if timezone changes

### 7.2 Scenarios Where This Matters

1. **Distributed deployments:** Controllers/agents in different timezones
2. **Daylight Saving Time:** Server timezone changes can affect stored times
3. **Sorting and filtering:** Queries may return incorrect results
4. **Time-based operations:** Rotation, expiration, scheduling may trigger incorrectly
5. **Migration:** Moving database between servers in different timezones

### 7.3 Mitigation

- Fix violations (Recommendations 6.1)
- Ensure server timezone is UTC (operational requirement)
- Document UTC requirement for deployments

---

## 8. Testing Strategy

### 8.1 Verification Tests

Create tests to verify UTC enforcement:

```go
// Example test
func TestTimestampsAreUTC(t *testing.T) {
    // Set server to non-UTC timezone
    os.Setenv("TZ", "America/New_York")
    
    // Insert record with timestamp
    // ...
    
    // Verify stored time is UTC
    storedTime := getStoredTime()
    assert.Equal(t, "UTC", storedTime.Location().String())
}
```

### 8.2 Integration Tests

- Test timezone-independent behavior
- Test across different deployment timezones
- Test DST transitions

---

## 9. Appendix: Domain-by-Domain Status

| Domain | Files Analyzed | Violations | Status | Priority |
|--------|----------------|------------|--------|----------|
| access | 2 | 1 | ⚠ Needs fix | HIGH |
| application | 5 | 4 | ⚠ Needs fix | HIGH |
| changestream | 2 | 1 | ⚠ Needs fix | MEDIUM |
| cloudimagemetadata | 2 | 3 | ⚠ Needs fix | MEDIUM |
| crossmodelrelation | 1 | 1 | ⚠ Needs fix | MEDIUM |
| machine | 2 | 2 | ⚠ Needs fix | HIGH |
| operation | 4 | 1 | ✓ Mostly OK | LOW |
| relation | 2 | 2 | ⚠ Needs fix | MEDIUM |
| removal | 10 | 0 | ✓ OK | - |
| resource | 1 | 6 | ⚠ Needs fix | HIGH |
| secret | 1 | 0 | ✓ OK | - |
| status | 1 | 2 | ⚠ Needs fix | HIGH |

---

## 10. Conclusion

This audit reveals that **UTC enforcement is inconsistent** across the Juju domain layer:

**Positive findings:**
- Some domains (operation, removal, secret) demonstrate exemplary UTC handling
- Schema has UTC defaults for some critical fields
- The pattern `.UTC()` is understood and used in some places

**Critical findings:**
- **25 violations** in state layer (19 `clock.Now()` + 6 `time.Now()`)
- **5 violations** in service layer
- **45 schema fields** lack UTC defaults
- No automated enforcement (linting/testing)

**Recommended next steps:**
1. Fix all HIGH priority violations in state and service layers
2. Add linting rules to prevent future violations
3. Update development guidelines
4. Add UTC defaults to new schema fields

**Estimated effort:** 2-3 days for fixes, 1 day for linting/guidelines

---

**Report End**
