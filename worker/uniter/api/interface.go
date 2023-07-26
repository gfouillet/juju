// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"time"

	"github.com/juju/charm/v11"
	"github.com/juju/names/v4"

	"github.com/juju/juju/api/agent/uniter"
	"github.com/juju/juju/core/application"
	"github.com/juju/juju/core/life"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/relation"
	"github.com/juju/juju/core/secrets"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/rpc/params"
)

// TODO(wallyworld) - mockgen breaks on WatchRelationUnits method due to generics.
// The generated mock file needed to be edited manually to fix the error(s).
// Typo below is deliberate to avoid mock linting from failing.
// //go:gen-erate go run go.uber.org/mock/mockgen -package api -destination uniter_mocks.go github.com/juju/juju/worker/uniter/api UniterClient

//go:generate go run go.uber.org/mock/mockgen -package api -destination domain_mocks.go github.com/juju/juju/worker/uniter/api Unit,Relation,RelationUnit,Application,Charm
//go:generate go run go.uber.org/mock/mockgen -package api -destination secrets_mocks.go github.com/juju/juju/worker/uniter/api SecretsClient,SecretsBackend

// ProviderIDGetter defines the API to get provider ID.
type ProviderIDGetter interface {
	ProviderID() string
	Refresh() error
	Name() string
}

// Unit defines the methods on uniter.api.Unit.
type Unit interface {
	ProviderIDGetter
	Life() life.Value
	Refresh() error
	ApplicationTag() names.ApplicationTag
	EnsureDead() error
	ClearResolved() error
	DestroyAllSubordinates() error
	HasSubordinates() (bool, error)
	LXDProfileName() (string, error)
	CanApplyLXDProfile() (bool, error)
	CharmURL() (string, error)
	Watch() (watcher.NotifyWatcher, error)

	// Used by runner.context.

	ApplicationName() string
	ConfigSettings() (charm.Settings, error)
	LogActionMessage(names.ActionTag, string) error
	Name() string
	NetworkInfo(bindings []string, relationId *int) (map[string]params.NetworkInfoResult, error)
	RequestReboot() error
	SetUnitStatus(unitStatus status.Status, info string, data map[string]interface{}) error
	SetAgentStatus(agentStatus status.Status, info string, data map[string]interface{}) error
	State() (params.UnitStateResult, error)
	SetState(unitState params.SetUnitStateArg) error
	Tag() names.UnitTag
	UnitStatus() (params.StatusResult, error)
	CommitHookChanges(params.CommitHookChangesArgs) error
	PublicAddress() (string, error)
	PrincipalName() (string, bool, error)
	AssignedMachine() (names.MachineTag, error)
	AvailabilityZone() (string, error)
	MeterStatus() (statusCode, statusInfo string, rErr error)
	PrivateAddress() (string, error)
	Resolved() params.ResolvedMode

	// Used by remotestate watcher.

	WatchConfigSettingsHash() (watcher.StringsWatcher, error)
	WatchTrustConfigSettingsHash() (watcher.StringsWatcher, error)
	WatchRelations() (watcher.StringsWatcher, error)
	WatchAddressesHash() (watcher.StringsWatcher, error)
	WatchUpgradeSeriesNotifications() (watcher.NotifyWatcher, error)
	WatchActionNotifications() (watcher.StringsWatcher, error)
	WatchStorage() (watcher.StringsWatcher, error)
	WatchInstanceData() (watcher.NotifyWatcher, error)
	UpgradeSeriesStatus() (model.UpgradeSeriesStatus, string, error)

	// Used by relationer.

	Application() (Application, error)
	RelationsStatus() ([]uniter.RelationStatus, error)
	Destroy() error

	// Used by operation.Callbacks.

	SetUpgradeSeriesStatus(upgradeSeriesStatus model.UpgradeSeriesStatus, reason string) error
	SetCharmURL(curl string) error
}

// Application defines the methods on uniter.api.Application.
type Application interface {
	Life() life.Value
	Tag() names.ApplicationTag
	Status(unitName string) (params.ApplicationStatusResult, error)
	SetStatus(unitName string, appStatus status.Status, info string, data map[string]interface{}) error
	CharmModifiedVersion() (int, error)
	CharmURL() (string, bool, error)

	// Used by remotestate watcher.

	WatchLeadershipSettings() (watcher.NotifyWatcher, error)
	Watch() (watcher.NotifyWatcher, error)
	Refresh() error
}

// Relation defines the methods on uniter.api.Relation.
type Relation interface {
	Endpoint() (*uniter.Endpoint, error)
	Id() int
	Life() life.Value
	OtherApplication() string
	Refresh() error
	SetStatus(status2 relation.Status) error
	String() string
	Suspended() bool
	Tag() names.RelationTag
	Unit(names.UnitTag) (RelationUnit, error)
	UpdateSuspended(bool)
}

// RelationUnit defines the methods on uniter.api.RelationUnit.
type RelationUnit interface {
	ApplicationSettings() (*uniter.Settings, error)
	Endpoint() uniter.Endpoint
	EnterScope() error
	LeaveScope() error
	Relation() Relation
	ReadSettings(name string) (params.Settings, error)
	Settings() (*uniter.Settings, error)
}

// Charm defines the methods on uniter.api.Charm.
type Charm interface {
	String() string
	LXDProfileRequired() (bool, error)
	ArchiveSha256() (string, error)
}

// SecretsAccessor is used by the hook context to access the secrets backend.
type SecretsAccessor interface {
	CreateSecretURIs(int) ([]*secrets.URI, error)
	SecretMetadata() ([]secrets.SecretOwnerMetadata, error)
	SecretRotated(uri string, oldRevision int) error
}

// SecretsWatcher is used by the remote state watcher.
type SecretsWatcher interface {
	WatchConsumedSecretsChanges(unitName string) (watcher.StringsWatcher, error)
	GetConsumerSecretsRevisionInfo(string, []string) (map[string]secrets.SecretRevisionInfo, error)
	WatchObsolete(ownerTags ...names.Tag) (watcher.StringsWatcher, error)
}

// SecretsBackend provides access to a secrets backend.
type SecretsBackend interface {
	GetContent(uri *secrets.URI, label string, refresh, peek bool) (secrets.SecretValue, error)
	SaveContent(uri *secrets.URI, revision int, value secrets.SecretValue) (secrets.ValueRef, error)
	DeleteContent(uri *secrets.URI, revision int) error
	DeleteExternalContent(ref secrets.ValueRef) error
}

// SecretsClient provides access to the secrets manager facade.
type SecretsClient interface {
	SecretsWatcher
	SecretsAccessor
}

// StorageAccessor is an interface for accessing information about
// storage attachments.
type StorageAccessor interface {
	StorageAttachment(names.StorageTag, names.UnitTag) (params.StorageAttachment, error)
	UnitStorageAttachments(names.UnitTag) ([]params.StorageAttachmentId, error)
	DestroyUnitStorageAttachments(names.UnitTag) error
	RemoveStorageAttachment(names.StorageTag, names.UnitTag) error
}

// UniterClient provides methods used by the uniter api facade client.
type UniterClient interface {
	StorageAccessor
	Charm(curl *charm.URL) (Charm, error)
	Unit(tag names.UnitTag) (Unit, error)
	Action(tag names.ActionTag) (*uniter.Action, error)
	Application(tag names.ApplicationTag) (Application, error)
	ActionStatus(tag names.ActionTag) (string, error)
	Relation(tag names.RelationTag) (Relation, error)
	RelationById(int) (Relation, error)
	Model() (*model.Model, error)
	ModelConfig() (*config.Config, error)
	UnitStorageAttachments(unitTag names.UnitTag) ([]params.StorageAttachmentId, error)
	StorageAttachment(storageTag names.StorageTag, unitTag names.UnitTag) (params.StorageAttachment, error)
	GoalState() (application.GoalState, error)
	CloudSpec() (*params.CloudSpec, error)
	ActionBegin(tag names.ActionTag) error
	ActionFinish(tag names.ActionTag, status string, results map[string]interface{}, message string) error
	UnitWorkloadVersion(tag names.UnitTag) (string, error)
	SetUnitWorkloadVersion(tag names.UnitTag, version string) error
	OpenedMachinePortRangesByEndpoint(machineTag names.MachineTag) (map[names.UnitTag]network.GroupedPortRanges, error)
	OpenedPortRangesByEndpoint() (map[names.UnitTag]network.GroupedPortRanges, error)
	LeadershipSettings() uniter.LeadershipSettingsAccessor
	SLALevel() (string, error)
	CloudAPIVersion() (string, error)
	APIAddresses() ([]string, error)
	WatchRelationUnits(names.RelationTag, names.UnitTag) (watcher.RelationUnitsWatcher, error)
	WatchStorageAttachment(names.StorageTag, names.UnitTag) (watcher.NotifyWatcher, error)
	WatchUpdateStatusHookInterval() (watcher.NotifyWatcher, error)
	UpdateStatusHookInterval() (time.Duration, error)
	StorageAttachmentLife([]params.StorageAttachmentId) ([]params.LifeResult, error)
}