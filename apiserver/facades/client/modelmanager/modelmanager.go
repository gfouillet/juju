// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelmanager

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/juju/description/v5"
	"github.com/juju/errors"
	"github.com/juju/loggo/v2"
	"github.com/juju/names/v5"
	jujutxn "github.com/juju/txn/v3"
	"github.com/juju/version/v2"

	"github.com/juju/juju/apiserver/authentication"
	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/credentialcommon"
	commonsecrets "github.com/juju/juju/apiserver/common/secrets"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/caas"
	jujucloud "github.com/juju/juju/cloud"
	"github.com/juju/juju/controller/modelmanager"
	"github.com/juju/juju/core/credential"
	"github.com/juju/juju/core/life"
	coremodel "github.com/juju/juju/core/model"
	"github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/core/permission"
	modelerrors "github.com/juju/juju/domain/model/errors"
	"github.com/juju/juju/environs"
	environscloudspec "github.com/juju/juju/environs/cloudspec"
	"github.com/juju/juju/environs/config"
	environsContext "github.com/juju/juju/environs/envcontext"
	"github.com/juju/juju/environs/space"
	"github.com/juju/juju/internal/tools"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/stateenvirons"
)

var (
	logger = loggo.GetLogger("juju.apiserver.modelmanager")

	// Overridden by tests.
	supportedFeaturesGetter = stateenvirons.SupportedFeatures
)

type newCaasBrokerFunc func(_ context.Context, args environs.OpenParams) (caas.Broker, error)

// ModelService defines a interface for interacting with the underlying state.
type ModelService interface {
	DeleteModel(context.Context, coremodel.UUID) error
}

// ModelExporter defines a interface for exporting models.
type ModelExporter interface {
	ExportModelPartial(ctx context.Context, cfg state.ExportConfig, store objectstore.ObjectStore) (description.Model, error)
}

// CloudService provides access to clouds.
type CloudService interface {
	common.CloudService
	ListAll(ctx context.Context) ([]jujucloud.Cloud, error)
}

// CredentialService exposes State methods needed by credential manager.
type CredentialService interface {
	CloudCredential(ctx context.Context, id credential.ID) (jujucloud.Credential, error)
	InvalidateCredential(ctx context.Context, id credential.ID, reason string) error
}

// StateBackend represents the mongo backend.
type StateBackend interface {
	common.ModelManagerBackend
	InvalidateModelCredential(string) error
}

// ModelManagerAPI implements the model manager interface and is
// the concrete implementation of the api end point.
type ModelManagerAPI struct {
	*common.ModelStatusAPI
	modelService       ModelService
	state              StateBackend
	modelExporter      ModelExporter
	ctlrState          common.ModelManagerBackend
	cloudService       CloudService
	credentialService  CredentialService
	store              objectstore.ObjectStore
	configSchemaSource config.ConfigSchemaSourceGetter
	check              common.BlockCheckerInterface
	authorizer         facade.Authorizer
	toolsFinder        common.ToolsFinder
	apiUser            names.UserTag
	isAdmin            bool
	model              common.Model
	getBroker          newCaasBrokerFunc
}

// NewModelManagerAPI creates a new api server endpoint for managing
// models.
func NewModelManagerAPI(
	st StateBackend,
	modelExporter ModelExporter,
	ctlrSt common.ModelManagerBackend,
	cloudService CloudService,
	credentialService CredentialService,
	modelService ModelService,
	store objectstore.ObjectStore,
	configSchemaSource config.ConfigSchemaSourceGetter,
	toolsFinder common.ToolsFinder,
	getBroker newCaasBrokerFunc,
	blockChecker common.BlockCheckerInterface,
	authorizer facade.Authorizer,
	m common.Model,
) (*ModelManagerAPI, error) {
	if !authorizer.AuthClient() {
		return nil, apiservererrors.ErrPerm
	}
	// Since we know this is a user tag (because AuthClient is true),
	// we just do the type assertion to the UserTag.
	apiUser, _ := authorizer.GetAuthTag().(names.UserTag)
	// Pretty much all of the user manager methods have special casing for admin
	// users, so look once when we start and remember if the user is an admin.
	err := authorizer.HasPermission(permission.SuperuserAccess, st.ControllerTag())
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return nil, errors.Trace(err)
	}
	isAdmin := err == nil

	return &ModelManagerAPI{
		ModelStatusAPI:     common.NewModelStatusAPI(st, authorizer, apiUser),
		state:              st,
		modelExporter:      modelExporter,
		ctlrState:          ctlrSt,
		cloudService:       cloudService,
		credentialService:  credentialService,
		store:              store,
		configSchemaSource: configSchemaSource,
		getBroker:          getBroker,
		check:              blockChecker,
		authorizer:         authorizer,
		toolsFinder:        toolsFinder,
		apiUser:            apiUser,
		isAdmin:            isAdmin,
		model:              m,
		modelService:       modelService,
	}, nil
}

// authCheck checks if the user is acting on their own behalf, or if they
// are an administrator acting on behalf of another user.
func (m *ModelManagerAPI) authCheck(user names.UserTag) error {
	if m.isAdmin {
		logger.Tracef("%q is a controller admin", m.apiUser.Id())
		return nil
	}

	// We can't just compare the UserTags themselves as the provider part
	// may be unset, and gets replaced with 'local'. We must compare against
	// the Canonical value of the user tag.
	if m.apiUser == user {
		return nil
	}
	return apiservererrors.ErrPerm
}

func (m *ModelManagerAPI) hasWriteAccess(modelTag names.ModelTag) (bool, error) {
	err := m.authorizer.HasPermission(permission.WriteAccess, modelTag)
	return err == nil, err
}

// ConfigSource describes a type that is able to provide config.
// Abstracted primarily for testing.
type ConfigSource interface {
	Config() (*config.Config, error)
}

func (m *ModelManagerAPI) newModelConfig(
	ctx context.Context,
	cloudSpec environscloudspec.CloudSpec,
	args params.ModelCreateArgs,
	source ConfigSource,
) (*config.Config, error) {
	// For now, we just smash to the two maps together as we store
	// the account values and the model config together in the
	// *config.Config instance.
	joint := make(map[string]interface{})
	for key, value := range args.Config {
		joint[key] = value
	}
	if _, ok := joint[config.UUIDKey]; ok {
		return nil, errors.New("uuid is generated, you cannot specify one")
	}
	if args.Name == "" {
		return nil, errors.NewNotValid(nil, "Name must be specified")
	}
	if _, ok := joint[config.NameKey]; ok {
		return nil, errors.New("name must not be specified in config")
	}
	joint[config.NameKey] = args.Name

	baseConfig, err := source.Config()
	if err != nil {
		return nil, errors.Trace(err)
	}

	regionSpec := &environscloudspec.CloudRegionSpec{Cloud: cloudSpec.Name, Region: cloudSpec.Region}
	if joint, err = m.state.ComposeNewModelConfig(m.configSchemaSource, joint, regionSpec); err != nil {
		return nil, errors.Trace(err)
	}

	creator := modelmanager.ModelConfigCreator{
		Provider: environs.Provider,
		FindTools: func(n version.Number) (tools.List, error) {
			if jujucloud.CloudTypeIsCAAS(cloudSpec.Type) {
				return tools.List{&tools.Tools{Version: version.Binary{Number: n}}}, nil
			}
			toolsList, err := m.toolsFinder.FindAgents(ctx, common.FindAgentsParams{
				Number: n,
			})
			if err != nil {
				return nil, errors.Trace(err)
			}
			return toolsList, nil
		},
	}
	return creator.NewModelConfig(cloudSpec, baseConfig, joint)
}

func (m *ModelManagerAPI) checkAddModelPermission(cloud string, userTag names.UserTag) (bool, error) {
	perm, err := m.ctlrState.GetCloudAccess(cloud, userTag)
	if err != nil && !errors.Is(err, errors.NotFound) {
		return false, errors.Trace(err)
	}
	if !perm.EqualOrGreaterCloudAccessThan(permission.AddModelAccess) {
		return false, nil
	}
	return true, nil
}

// CreateModel creates a new model using the account and
// model config specified in the args.
func (m *ModelManagerAPI) CreateModel(ctx context.Context, args params.ModelCreateArgs) (params.ModelInfo, error) {
	result := params.ModelInfo{}

	// Get the controller model first. We need it both for the state
	// server owner and the ability to get the config.
	controllerModel, err := m.ctlrState.Model()
	if err != nil {
		return result, errors.Trace(err)
	}

	var cloudTag names.CloudTag
	cloudRegionName := args.CloudRegion
	if args.CloudTag != "" {
		var err error
		cloudTag, err = names.ParseCloudTag(args.CloudTag)
		if err != nil {
			return result, errors.Trace(err)
		}
	} else {
		cloudTag = names.NewCloudTag(controllerModel.CloudName())
	}
	if cloudRegionName == "" && cloudTag.Id() == controllerModel.CloudName() {
		cloudRegionName = controllerModel.CloudRegion()
	}

	err = m.authorizer.HasPermission(permission.SuperuserAccess, m.state.ControllerTag())
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return result, errors.Trace(err)
	}
	if err != nil {
		canAddModel, err := m.checkAddModelPermission(cloudTag.Id(), m.apiUser)
		if err != nil {
			return result, errors.Trace(err)
		}
		if !canAddModel {
			return result, apiservererrors.ErrPerm
		}
	}

	ownerTag, err := names.ParseUserTag(args.OwnerTag)
	if err != nil {
		return result, errors.Trace(err)
	}

	// a special case of ErrPerm will happen if the user has add-model permission but is trying to
	// create a model for another person, which is not yet supported.
	if !m.isAdmin && ownerTag != m.apiUser {
		return result, errors.Annotatef(apiservererrors.ErrPerm, "%q permission does not permit creation of models for different owners", permission.AddModelAccess)
	}

	cloud, err := m.cloudService.Cloud(ctx, cloudTag.Id())
	if err != nil {
		if errors.Is(err, errors.NotFound) && args.CloudTag != "" {
			// A cloud was specified, and it was not found.
			// Annotate the error with the supported clouds.
			clouds, err := m.cloudService.ListAll(ctx)
			if err != nil {
				return result, errors.Trace(err)
			}
			cloudNames := make([]string, 0, len(clouds))
			for _, cld := range clouds {
				cloudNames = append(cloudNames, cld.Name)
			}
			sort.Strings(cloudNames)
			return result, errors.NewNotFound(err, fmt.Sprintf(
				"cloud %q not found, expected one of %q",
				cloudTag.Id(), cloudNames,
			))
		}
		return result, errors.Annotate(err, "getting cloud definition")
	}

	var cloudCredentialTag names.CloudCredentialTag
	if args.CloudCredentialTag != "" {
		var err error
		cloudCredentialTag, err = names.ParseCloudCredentialTag(args.CloudCredentialTag)
		if err != nil {
			return result, errors.Trace(err)
		}
	} else {
		if ownerTag == controllerModel.Owner() {
			cloudCredentialTag, _ = controllerModel.CloudCredentialTag()
		} else {
			// TODO(axw) check if the user has one and only one
			// cloud credential, and if so, use it? For now, we
			// require the user to specify a credential unless
			// the cloud does not require one.
			var hasEmpty bool
			for _, authType := range cloud.AuthTypes {
				if authType != jujucloud.EmptyAuthType {
					continue
				}
				hasEmpty = true
				break
			}
			if !hasEmpty {
				return result, errors.NewNotValid(nil, "no credential specified")
			}
		}
	}

	var cred *jujucloud.Credential
	if cloudCredentialTag != (names.CloudCredentialTag{}) {
		credentialValue, err := m.credentialService.CloudCredential(ctx, credential.IdFromTag(cloudCredentialTag))
		if err != nil {
			return result, errors.Annotate(err, "getting credential")
		}
		cloudCredential := jujucloud.NewNamedCredential(
			credentialValue.Label,
			credentialValue.AuthType(),
			credentialValue.Attributes(),
			credentialValue.Revoked,
		)
		cred = &cloudCredential
	}

	cloudSpec, err := environscloudspec.MakeCloudSpec(*cloud, cloudRegionName, cred)
	if err != nil {
		return result, errors.Trace(err)
	}

	var createdModel common.Model
	if jujucloud.CloudIsCAAS(*cloud) {
		createdModel, err = m.newCAASModel(
			ctx,
			cloudSpec,
			args,
			controllerModel,
			cloudTag,
			cloudRegionName,
			cloudCredentialTag,
			ownerTag,
		)
	} else {
		createdModel, err = m.newModel(
			ctx,
			cloudSpec,
			args,
			controllerModel,
			cloudTag,
			cloudRegionName,
			cloudCredentialTag,
			ownerTag,
		)
	}
	if err != nil {
		return result, errors.Trace(err)
	}

	return m.getModelInfo(ctx, createdModel.ModelTag(), false)
}

func (m *ModelManagerAPI) newCAASModel(
	ctx context.Context,
	cloudSpec environscloudspec.CloudSpec,
	createArgs params.ModelCreateArgs,
	controllerModel common.Model,
	cloudTag names.CloudTag,
	cloudRegionName string,
	cloudCredentialTag names.CloudCredentialTag,
	ownerTag names.UserTag,
) (_ common.Model, err error) {
	newConfig, err := m.newModelConfig(ctx, cloudSpec, createArgs, controllerModel)
	if err != nil {
		return nil, errors.Annotate(err, "failed to create config")
	}
	controllerConfig, err := m.state.ControllerConfig()
	if err != nil {
		return nil, errors.Annotate(err, "getting controller config")
	}

	defer func() {
		// Retain the error stack but with a better message.
		if errors.Is(err, errors.AlreadyExists) {
			err = errors.Wrap(err, errors.NewAlreadyExists(nil,
				`
the model cannot be created because a namespace with the proposed
model name already exists in the k8s cluster.
Please choose a different model name.
`[1:],
			))
		}
	}()

	broker, err := m.getBroker(ctx, environs.OpenParams{
		ControllerUUID: controllerConfig.ControllerUUID(),
		Cloud:          cloudSpec,
		Config:         newConfig,
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to open kubernetes client")
	}

	callCtx := environsContext.WithoutCredentialInvalidator(ctx)
	if err = broker.Create(
		callCtx,
		environs.CreateParams{ControllerUUID: controllerConfig.ControllerUUID()},
	); err != nil {
		return nil, errors.Annotatef(err, "creating namespace %q", createArgs.Name)
	}

	storageProviderRegistry := stateenvirons.NewStorageProviderRegistry(broker)

	model, st, err := m.state.NewModel(state.ModelArgs{
		Type:                    state.ModelTypeCAAS,
		CloudName:               cloudTag.Id(),
		CloudRegion:             cloudRegionName,
		CloudCredential:         cloudCredentialTag,
		Config:                  newConfig,
		Owner:                   ownerTag,
		StorageProviderRegistry: storageProviderRegistry,
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to create new model")
	}
	defer st.Close()

	return model, nil
}

func (m *ModelManagerAPI) newModel(
	ctx context.Context,
	cloudSpec environscloudspec.CloudSpec,
	createArgs params.ModelCreateArgs,
	controllerModel common.Model,
	cloudTag names.CloudTag,
	cloudRegionName string,
	cloudCredentialTag names.CloudCredentialTag,
	ownerTag names.UserTag,
) (common.Model, error) {
	newConfig, err := m.newModelConfig(ctx, cloudSpec, createArgs, controllerModel)
	if err != nil {
		return nil, errors.Annotate(err, "failed to create config")
	}

	controllerCfg, err := m.state.ControllerConfig()
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Create the Environ.
	env, err := environs.New(ctx, environs.OpenParams{
		ControllerUUID: controllerCfg.ControllerUUID(),
		Cloud:          cloudSpec,
		Config:         newConfig,
	})
	if err != nil {
		return nil, errors.Annotate(err, "failed to open environ")
	}

	callCtx := environsContext.WithoutCredentialInvalidator(ctx)
	err = env.Create(
		callCtx,
		environs.CreateParams{
			ControllerUUID: controllerCfg.ControllerUUID(),
		},
	)
	if err != nil {
		return nil, errors.Annotate(err, "failed to create environ")
	}
	storageProviderRegistry := stateenvirons.NewStorageProviderRegistry(env)

	// NOTE: check the agent-version of the config, and if it is > the current
	// version, it is not supported, also check existing tools, and if we don't
	// have tools for that version, also die.
	model, st, err := m.state.NewModel(state.ModelArgs{
		Type:                    state.ModelTypeIAAS,
		CloudName:               cloudTag.Id(),
		CloudRegion:             cloudRegionName,
		CloudCredential:         cloudCredentialTag,
		Config:                  newConfig,
		Owner:                   ownerTag,
		StorageProviderRegistry: storageProviderRegistry,
		EnvironVersion:          env.Provider().Version(),
	})
	if err != nil {
		// Clean up the environ.
		if e := env.Destroy(callCtx); e != nil {
			logger.Warningf("failed to destroy environ, error %v", e)
		}
		return nil, errors.Annotate(err, "failed to create new model")
	}
	defer st.Close()

	if err = model.AutoConfigureContainerNetworking(env, m.configSchemaSource); err != nil {
		if errors.Is(err, errors.NotSupported) {
			logger.Debugf("Not performing container networking autoconfiguration on a non-networking environment")
		} else {
			return nil, errors.Annotate(err, "Failed to perform container networking autoconfiguration")
		}
	}

	invalidatorFuncGetter := credentialcommon.ModelCredentialInvalidatorGetter(m.credentialService, credentialStateShim{st.(StateBackend)})
	invalidatorFunc, err := invalidatorFuncGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	callCtx = environsContext.WithCredentialInvalidator(ctx, invalidatorFunc)
	if err = space.ReloadSpaces(callCtx, spaceStateShim{
		ModelManagerBackend: st,
	}, env); err != nil {
		if errors.Is(err, errors.NotSupported) {
			logger.Debugf("Not performing spaces load on a non-networking environment")
		} else {
			return nil, errors.Annotate(err, "Failed to perform spaces discovery")
		}
	}
	return model, nil
}

func (m *ModelManagerAPI) dumpModel(ctx context.Context, args params.Entity, simplified bool) ([]byte, error) {
	modelTag, err := names.ParseModelTag(args.Tag)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if !m.isAdmin {
		if err := m.authorizer.HasPermission(permission.AdminAccess, modelTag); err != nil {
			return nil, err
		}
	}

	_, release, err := m.state.GetBackend(modelTag.Id())
	if err != nil {
		if errors.Is(err, errors.NotFound) {
			return nil, errors.Trace(apiservererrors.ErrBadId)
		}
		return nil, errors.Trace(err)
	}
	defer release()

	exportConfig := state.ExportConfig{IgnoreIncompleteModel: true}
	if simplified {
		exportConfig.SkipActions = true
		exportConfig.SkipAnnotations = true
		exportConfig.SkipCloudImageMetadata = true
		exportConfig.SkipCredentials = true
		exportConfig.SkipIPAddresses = true
		exportConfig.SkipSettings = true
		exportConfig.SkipSSHHostKeys = true
		exportConfig.SkipStatusHistory = true
		exportConfig.SkipLinkLayerDevices = true
	}

	model, err := m.modelExporter.ExportModelPartial(ctx, exportConfig, m.store)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bytes, err := description.Serialize(model)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return bytes, nil
}

func (m *ModelManagerAPI) dumpModelDB(args params.Entity) (map[string]interface{}, error) {
	modelTag, err := names.ParseModelTag(args.Tag)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if !m.isAdmin {
		if err := m.authorizer.HasPermission(permission.AdminAccess, modelTag); err != nil {
			return nil, err
		}
	}

	type dumper interface {
		DumpAll() (map[string]interface{}, error)
		ModelTag() names.ModelTag
	}

	var st dumper = m.state
	if st.ModelTag() != modelTag {
		newSt, release, err := m.state.GetBackend(modelTag.Id())
		if errors.Is(err, errors.NotFound) {
			return nil, errors.Trace(apiservererrors.ErrBadId)
		} else if err != nil {
			return nil, errors.Trace(err)
		}
		defer release()
		st = newSt
	}

	return st.DumpAll()
}

// DumpModels will export the models into the database agnostic
// representation. The user needs to either be a controller admin, or have
// admin privileges on the model itself.
func (m *ModelManagerAPI) DumpModels(ctx context.Context, args params.DumpModelRequest) params.StringResults {
	results := params.StringResults{
		Results: make([]params.StringResult, len(args.Entities)),
	}
	for i, entity := range args.Entities {
		bytes, err := m.dumpModel(ctx, entity, args.Simplified)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		// We know here that the bytes are valid YAML.
		results.Results[i].Result = string(bytes)
	}
	return results
}

// DumpModelsDB will gather all documents from all model collections
// for the specified model. The map result contains a map of collection
// names to lists of documents represented as maps.
func (m *ModelManagerAPI) DumpModelsDB(ctx context.Context, args params.Entities) params.MapResults {
	results := params.MapResults{
		Results: make([]params.MapResult, len(args.Entities)),
	}
	for i, entity := range args.Entities {
		dumped, err := m.dumpModelDB(entity)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results.Results[i].Result = dumped
	}
	return results
}

// ListModelSummaries returns models that the specified user
// has access to in the current server.  Controller admins (superuser)
// can list models for any user.  Other users
// can only ask about their own models.
func (m *ModelManagerAPI) ListModelSummaries(ctx context.Context, req params.ModelSummariesRequest) (params.ModelSummaryResults, error) {
	result := params.ModelSummaryResults{}

	userTag, err := names.ParseUserTag(req.UserTag)
	if err != nil {
		return result, errors.Trace(err)
	}

	err = m.authCheck(userTag)
	if err != nil {
		return result, errors.Trace(err)
	}

	modelInfos, err := m.state.ModelSummariesForUser(userTag, req.All && m.isAdmin)
	if err != nil {
		return result, errors.Trace(err)
	}
	err = m.fillInStatusBasedOnCloudCredentialValidity(ctx, modelInfos)
	if err != nil {
		return result, errors.Trace(err)
	}

	for _, mi := range modelInfos {
		summary := m.makeModelSummary(mi)
		result.Results = append(result.Results, params.ModelSummaryResult{Result: summary})
	}
	return result, nil

}

func (m *ModelManagerAPI) makeModelSummary(mi state.ModelSummary) *params.ModelSummary {
	summary := &params.ModelSummary{
		Name:           mi.Name,
		UUID:           mi.UUID,
		Type:           string(mi.Type),
		OwnerTag:       names.NewUserTag(mi.Owner).String(),
		ControllerUUID: mi.ControllerUUID,
		IsController:   mi.IsController,
		Life:           life.Value(mi.Life.String()),

		CloudTag:    mi.CloudTag,
		CloudRegion: mi.CloudRegion,

		CloudCredentialTag: mi.CloudCredentialTag,

		ProviderType: mi.ProviderType,
		AgentVersion: mi.AgentVersion,

		Status:             common.EntityStatusFromState(mi.Status),
		Counts:             []params.ModelEntityCount{},
		UserLastConnection: mi.UserLastConnection,
	}
	if mi.MachineCount > 0 {
		summary.Counts = append(summary.Counts, params.ModelEntityCount{params.Machines, mi.MachineCount})
	}

	if mi.CoreCount > 0 {
		summary.Counts = append(summary.Counts, params.ModelEntityCount{params.Cores, mi.CoreCount})
	}

	if mi.UnitCount > 0 {
		summary.Counts = append(summary.Counts, params.ModelEntityCount{params.Units, mi.UnitCount})
	}

	access, err := common.StateToParamsUserAccessPermission(mi.Access)
	if err == nil {
		summary.UserAccess = access
	}
	if mi.Migration != nil {
		migration := mi.Migration
		startTime := migration.StartTime()
		endTime := new(time.Time)
		*endTime = migration.EndTime()
		var zero time.Time
		if *endTime == zero {
			endTime = nil
		}

		summary.Migration = &params.ModelMigrationStatus{
			Status: migration.StatusMessage(),
			Start:  &startTime,
			End:    endTime,
		}
	}
	return summary
}

// fillInStatusBasedOnCloudCredentialValidity fills in the Status on every model (if credential is invalid).
func (m *ModelManagerAPI) fillInStatusBasedOnCloudCredentialValidity(ctx context.Context, summaries []state.ModelSummary) error {
	credentialModels := map[names.CloudCredentialTag][]string{}
	indexByUUID := make(map[string]int)
	for i, model := range summaries {
		if model.CloudCredentialTag == "" {
			continue
		}
		indexByUUID[model.UUID] = i
		tag, err := names.ParseCloudCredentialTag(model.CloudCredentialTag)
		if err != nil {
			logger.Warningf("could not parse cloud credential tag %v for model%v: %v", model.CloudCredentialTag, model.UUID, err)
			// Don't stop the rest of the models
			continue
		}
		summaries, ok := credentialModels[tag]
		if !ok {
			summaries = []string{}
		}
		credentialModels[tag] = append(summaries, model.UUID)
	}
	if len(credentialModels) == 0 {
		return nil
	}

	// TODO(wallyworld) - bulk query
	for tag := range credentialModels {
		cred, err := m.credentialService.CloudCredential(ctx, credential.IdFromTag(tag))
		if err != nil {
			return errors.Trace(err)
		}
		if cred.Invalid {
			for _, uuid := range credentialModels[tag] {
				idx, ok := indexByUUID[uuid]
				if !ok {
					continue
				}
				details := &summaries[idx]
				details.Status = state.ModelStatusInvalidCredential(cred.InvalidReason)
			}
		}
	}
	return nil
}

// ListModels returns the models that the specified user
// has access to in the current server.  Controller admins (superuser)
// can list models for any user.  Other users
// can only ask about their own models.
func (m *ModelManagerAPI) ListModels(ctx context.Context, user params.Entity) (params.UserModelList, error) {
	result := params.UserModelList{}

	userTag, err := names.ParseUserTag(user.Tag)
	if err != nil {
		return result, errors.Trace(err)
	}

	err = m.authCheck(userTag)
	if err != nil {
		return result, errors.Trace(err)
	}

	modelInfos, err := m.state.ModelBasicInfoForUser(userTag, m.isAdmin)
	if err != nil {
		return result, errors.Trace(err)
	}

	for _, mi := range modelInfos {
		var ownerTag names.UserTag
		if names.IsValidUser(mi.Owner) {
			ownerTag = names.NewUserTag(mi.Owner)
		} else {
			// no reason to fail the request here, as it wasn't the users fault
			logger.Warningf("for model %v, got an invalid owner: %q", mi.UUID, mi.Owner)
		}
		lastConnection := mi.LastConnection
		result.UserModels = append(result.UserModels, params.UserModel{
			Model: params.Model{
				Name:     mi.Name,
				UUID:     mi.UUID,
				Type:     string(mi.Type),
				OwnerTag: ownerTag.String(),
			},
			LastConnection: &lastConnection,
		})
	}

	return result, nil
}

// DestroyModels will try to destroy the specified models.
// If there is a block on destruction, this method will return an error.
// From ModelManager v7 onwards, DestroyModels gains 'force' and 'max-wait' parameters.
func (m *ModelManagerAPI) DestroyModels(ctx context.Context, args params.DestroyModelsParams) (params.ErrorResults, error) {
	results := params.ErrorResults{
		Results: make([]params.ErrorResult, len(args.Models)),
	}

	destroyModel := func(modelUUID string, destroyStorage, force *bool, maxWait *time.Duration, timeout *time.Duration) error {
		st, releaseSt, err := m.state.GetBackend(modelUUID)
		if err != nil {
			return errors.Trace(err)
		}
		defer releaseSt()

		stModel, err := st.Model()
		if err != nil {
			return errors.Trace(err)
		}
		if !m.isAdmin {
			if err := m.authorizer.HasPermission(permission.AdminAccess, stModel.ModelTag()); err != nil {
				return err
			}
		}

		err = errors.Trace(common.DestroyModel(ctx, st, destroyStorage, force, maxWait, timeout))
		if err != nil {
			return errors.Trace(err)
		}

		// TODO (stickupkid): There are consequences to this failing after the
		// model has been deleted. Although in it's current guise this shouldn't
		// cause too much fallout. If we're unable to delete the model from the
		// database, then we won't be able to create a new model with the same
		// model uuid as there is a UNIQUE constraint on the model uuid column.
		err = m.modelService.DeleteModel(ctx, coremodel.UUID(stModel.UUID()))
		if err != nil && errors.Is(err, modelerrors.NotFound) {
			return nil
		}
		return errors.Trace(err)
	}

	for i, arg := range args.Models {
		tag, err := names.ParseModelTag(arg.ModelTag)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		if err := destroyModel(tag.Id(), arg.DestroyStorage, arg.Force, arg.MaxWait, arg.Timeout); err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
	}
	return results, nil
}

// ModelInfo returns information about the specified models.
func (m *ModelManagerAPI) ModelInfo(ctx context.Context, args params.Entities) (params.ModelInfoResults, error) {
	results := params.ModelInfoResults{
		Results: make([]params.ModelInfoResult, len(args.Entities)),
	}

	getModelInfo := func(arg params.Entity) (params.ModelInfo, error) {
		tag, err := names.ParseModelTag(arg.Tag)
		if err != nil {
			return params.ModelInfo{}, errors.Trace(err)
		}
		modelInfo, err := m.getModelInfo(ctx, tag, true)
		if err != nil {
			return params.ModelInfo{}, errors.Trace(err)
		}
		if modelInfo.CloudCredentialTag != "" {
			credentialTag, err := names.ParseCloudCredentialTag(modelInfo.CloudCredentialTag)
			if err != nil {
				return params.ModelInfo{}, errors.Trace(err)
			}
			cred, err := m.credentialService.CloudCredential(ctx, credential.IdFromTag(credentialTag))
			if err != nil {
				return params.ModelInfo{}, errors.Trace(err)
			}
			valid := !cred.Invalid
			modelInfo.CloudCredentialValidity = &valid
		}
		return modelInfo, nil
	}

	for i, arg := range args.Entities {
		modelInfo, err := getModelInfo(arg)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results.Results[i].Result = &modelInfo
	}
	return results, nil
}

func (m *ModelManagerAPI) getModelInfo(ctx context.Context, tag names.ModelTag, withSecrets bool) (params.ModelInfo, error) {
	st, release, err := m.state.GetBackend(tag.Id())
	if errors.Is(err, errors.NotFound) {
		return params.ModelInfo{}, errors.Trace(apiservererrors.ErrPerm)
	} else if err != nil {
		return params.ModelInfo{}, errors.Trace(err)
	}
	defer release()

	model, err := st.Model()
	if errors.Is(err, errors.NotFound) {
		return params.ModelInfo{}, errors.Trace(apiservererrors.ErrPerm)
	} else if err != nil {
		return params.ModelInfo{}, errors.Trace(err)
	}

	info := params.ModelInfo{
		Name:           model.Name(),
		Type:           string(model.Type()),
		UUID:           model.UUID(),
		ControllerUUID: model.ControllerUUID(),
		IsController:   st.IsController(),
		OwnerTag:       model.Owner().String(),
		Life:           life.Value(model.Life().String()),
		CloudTag:       names.NewCloudTag(model.CloudName()).String(),
		CloudRegion:    model.CloudRegion(),
	}

	if cloudCredentialTag, ok := model.CloudCredentialTag(); ok {
		info.CloudCredentialTag = cloudCredentialTag.String()
	}

	// If model is not alive - dying or dead - or if it is being imported,
	// there is no guarantee that the rest of the call will succeed.
	// For these models we can ignore NotFound errors coming from persistence layer.
	// However, for Alive models, these errors are genuine and cannot be ignored.
	ignoreNotFoundError := model.Life() != state.Alive || model.MigrationMode() == state.MigrationModeImporting

	// If we received an error and cannot ignore it, we should consider it fatal and surface it.
	// We should do the same if we can ignore NotFound errors but the given error is of some other type.
	shouldErr := func(thisErr error) bool {
		if thisErr == nil {
			return false
		}
		return !ignoreNotFoundError || !errors.Is(thisErr, errors.NotFound)
	}
	cfg, err := model.Config()
	if shouldErr(err) {
		return params.ModelInfo{}, errors.Trace(err)
	}
	if err == nil {
		info.ProviderType = cfg.Type()

		if agentVersion, exists := cfg.AgentVersion(); exists {
			info.AgentVersion = &agentVersion
		}
	}

	status, err := model.Status()
	if shouldErr(err) {
		return params.ModelInfo{}, errors.Trace(err)
	}
	if err == nil {
		entityStatus := common.EntityStatusFromState(status)
		info.Status = entityStatus
	}

	// If the user is a controller superuser, they are considered a model
	// admin.
	modelAdmin := m.isAdmin
	if !m.isAdmin {
		err = m.authorizer.HasPermission(permission.AdminAccess, model.ModelTag())
		modelAdmin = err == nil
	}

	users, err := model.Users()
	if shouldErr(err) {
		return params.ModelInfo{}, errors.Trace(err)
	}
	if err == nil {
		for _, user := range users {
			if !modelAdmin && m.authCheck(user.UserTag) != nil {
				// The authenticated user is neither the controller
				// superuser, a model administrator, nor the model user, so
				// has no business knowing about the model user.
				continue
			}

			userInfo, err := common.ModelUserInfo(user, model)
			if err != nil {
				return params.ModelInfo{}, errors.Trace(err)
			}
			info.Users = append(info.Users, userInfo)
		}

		if len(info.Users) == 0 {
			// No users, which means the authenticated user doesn't
			// have access to the model.
			return params.ModelInfo{}, errors.Trace(apiservererrors.ErrPerm)
		}
	}

	canSeeMachinesAndSecrets := modelAdmin
	if !canSeeMachinesAndSecrets {
		canSeeMachinesAndSecrets, err = m.hasWriteAccess(tag)
		if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
			return params.ModelInfo{}, errors.Trace(err)
		}
	}
	if canSeeMachinesAndSecrets {
		if info.Machines, err = common.ModelMachineInfo(st); shouldErr(err) {
			return params.ModelInfo{}, err
		}
	}
	if withSecrets && canSeeMachinesAndSecrets {
		if info.SecretBackends, err = commonsecrets.BackendSummaryInfo(
			m.state, st, st, st.ControllerUUID(), false, commonsecrets.BackendFilter{},
		); shouldErr(err) {
			return params.ModelInfo{}, err
		}
		// Don't expose the id.
		for i := range info.SecretBackends {
			info.SecretBackends[i].ID = ""
		}
	}

	migration, err := st.LatestMigration()
	if err != nil && !errors.Is(err, errors.NotFound) {
		return params.ModelInfo{}, errors.Trace(err)
	}
	if err == nil {
		startTime := migration.StartTime()
		endTime := new(time.Time)
		*endTime = migration.EndTime()
		var zero time.Time
		if *endTime == zero {
			endTime = nil
		}
		info.Migration = &params.ModelMigrationStatus{
			Status: migration.StatusMessage(),
			Start:  &startTime,
			End:    endTime,
		}
	}

	fs, err := supportedFeaturesGetter(model, m.cloudService, m.credentialService)
	if err != nil {
		return params.ModelInfo{}, err
	}
	for _, feat := range fs.AsList() {
		mappedFeat := params.SupportedFeature{
			Name:        feat.Name,
			Description: feat.Description,
		}

		if feat.Version != nil {
			mappedFeat.Version = feat.Version.String()
		}

		info.SupportedFeatures = append(info.SupportedFeatures, mappedFeat)
	}
	return info, nil
}

// ModifyModelAccess changes the model access granted to users.
func (m *ModelManagerAPI) ModifyModelAccess(ctx context.Context, args params.ModifyModelAccessRequest) (result params.ErrorResults, _ error) {
	result = params.ErrorResults{
		Results: make([]params.ErrorResult, len(args.Changes)),
	}

	err := m.authorizer.HasPermission(permission.SuperuserAccess, m.state.ControllerTag())
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return result, errors.Trace(err)
	}
	canModifyController := err == nil

	if len(args.Changes) == 0 {
		return result, nil
	}

	for i, arg := range args.Changes {
		modelAccess := permission.Access(arg.Access)
		if err := permission.ValidateModelAccess(modelAccess); err != nil {
			err = errors.Annotate(err, "could not modify model access")
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		modelTag, err := names.ParseModelTag(arg.ModelTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(errors.Annotate(err, "could not modify model access"))
			continue
		}
		err = m.authorizer.HasPermission(permission.AdminAccess, modelTag)
		if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
			return result, errors.Trace(err)
		}
		canModify := err == nil || canModifyController

		if !canModify {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
			continue
		}

		targetUserTag, err := names.ParseUserTag(arg.UserTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(errors.Annotate(err, "could not modify model access"))
			continue
		}

		result.Results[i].Error = apiservererrors.ServerError(
			changeModelAccess(m.state, modelTag, m.apiUser, targetUserTag, arg.Action, modelAccess, m.isAdmin))
	}
	return result, nil
}

func userAuthorizedToChangeAccess(st common.ModelManagerBackend, userIsAdmin bool, userTag names.UserTag) error {
	if userIsAdmin {
		// Just confirm that the model that has been given is a valid model.
		_, err := st.Model()
		if err != nil {
			return errors.Trace(err)
		}
		return nil
	}

	// Get the current user's ModelUser for the Model to see if the user has
	// permission to grant or revoke permissions on the model.
	currentUser, err := st.UserAccess(userTag, st.ModelTag())
	if err != nil {
		if errors.Is(err, errors.NotFound) {
			// No, this user doesn't have permission.
			return apiservererrors.ErrPerm
		}
		return errors.Annotate(err, "could not retrieve user")
	}
	if currentUser.Access != permission.AdminAccess {
		return apiservererrors.ErrPerm
	}
	return nil
}

// changeModelAccess performs the requested access grant or revoke action for the
// specified user on the specified model.
func changeModelAccess(accessor common.ModelManagerBackend, modelTag names.ModelTag, apiUser, targetUserTag names.UserTag, action params.ModelAction, access permission.Access, userIsAdmin bool) error {
	st, release, err := accessor.GetBackend(modelTag.Id())
	if err != nil {
		return errors.Annotate(err, "could not lookup model")
	}
	defer release()

	if err := userAuthorizedToChangeAccess(st, userIsAdmin, apiUser); err != nil {
		return errors.Trace(err)
	}

	model, err := st.Model()
	if err != nil {
		return errors.Trace(err)
	}

	switch action {
	case params.GrantModelAccess:
		_, err = model.AddUser(state.UserAccessSpec{User: targetUserTag, CreatedBy: apiUser, Access: access})
		if errors.Is(err, errors.AlreadyExists) {
			modelUser, err := st.UserAccess(targetUserTag, modelTag)
			if errors.Is(err, errors.NotFound) {
				// Conflicts with prior check, must be inconsistent state.
				err = jujutxn.ErrExcessiveContention
			}
			if err != nil {
				return errors.Annotate(err, "could not look up model access for user")
			}

			// Only set access if greater access is being granted.
			if modelUser.Access.EqualOrGreaterModelAccessThan(access) {
				return errors.Errorf("user already has %q access or greater", access)
			}
			if _, err = st.SetUserAccess(modelUser.UserTag, modelUser.Object, access); err != nil {
				return errors.Annotate(err, "could not set model access for user")
			}
			return nil
		}
		return errors.Annotate(err, "could not grant model access")

	case params.RevokeModelAccess:
		switch access {
		case permission.ReadAccess:
			// Revoking read access removes all access.
			err := st.RemoveUserAccess(targetUserTag, modelTag)
			return errors.Annotate(err, "could not revoke model access")
		case permission.WriteAccess:
			// Revoking write access sets read-only.
			modelUser, err := st.UserAccess(targetUserTag, modelTag)
			if err != nil {
				return errors.Annotate(err, "could not look up model access for user")
			}
			_, err = st.SetUserAccess(modelUser.UserTag, modelUser.Object, permission.ReadAccess)
			return errors.Annotate(err, "could not set model access to read-only")
		case permission.AdminAccess:
			// Revoking admin access sets read-write.
			modelUser, err := st.UserAccess(targetUserTag, modelTag)
			if err != nil {
				return errors.Annotate(err, "could not look up model access for user")
			}
			_, err = st.SetUserAccess(modelUser.UserTag, modelUser.Object, permission.WriteAccess)
			return errors.Annotate(err, "could not set model access to read-write")

		default:
			return errors.Errorf("don't know how to revoke %q access", access)
		}

	default:
		return errors.Errorf("unknown action %q", action)
	}
}

// ModelDefaultsForClouds returns the default config values for the specified
// clouds.
func (m *ModelManagerAPI) ModelDefaultsForClouds(ctx context.Context, args params.Entities) (params.ModelDefaultsResults, error) {
	result := params.ModelDefaultsResults{}
	if !m.isAdmin {
		return result, apiservererrors.ErrPerm
	}
	result.Results = make([]params.ModelDefaultsResult, len(args.Entities))
	for i, entity := range args.Entities {
		cloudTag, err := names.ParseCloudTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		result.Results[i] = m.modelDefaults(cloudTag.Id())
	}
	return result, nil
}

func (m *ModelManagerAPI) modelDefaults(cloud string) params.ModelDefaultsResult {
	result := params.ModelDefaultsResult{}
	values, err := m.ctlrState.ModelConfigDefaultValues(cloud)
	if err != nil {
		result.Error = apiservererrors.ServerError(err)
		return result
	}
	result.Config = make(map[string]params.ModelDefaults)
	for attr, val := range values {
		settings := params.ModelDefaults{
			Controller: val.Controller,
			Default:    val.Default,
		}
		for _, v := range val.Regions {
			settings.Regions = append(
				settings.Regions, params.RegionDefaults{
					RegionName: v.Name,
					Value:      v.Value})
		}
		result.Config[attr] = settings
	}
	return result
}

// SetModelDefaults writes new values for the specified default model settings.
func (m *ModelManagerAPI) SetModelDefaults(ctx context.Context, args params.SetModelDefaults) (params.ErrorResults, error) {
	results := params.ErrorResults{Results: make([]params.ErrorResult, len(args.Config))}
	if err := m.check.ChangeAllowed(ctx); err != nil {
		return results, errors.Trace(err)
	}
	for i, arg := range args.Config {
		results.Results[i].Error = apiservererrors.ServerError(
			m.setModelDefaults(ctx, arg),
		)
	}
	return results, nil
}

func (m *ModelManagerAPI) setModelDefaults(ctx context.Context, args params.ModelDefaultValues) error {
	if !m.isAdmin {
		return apiservererrors.ErrPerm
	}

	if err := m.check.ChangeAllowed(ctx); err != nil {
		return errors.Trace(err)
	}
	// Make sure we don't allow changing agent-version.
	if _, found := args.Config["agent-version"]; found {
		return errors.New("agent-version cannot have a default value")
	}

	var rspec *environscloudspec.CloudRegionSpec
	if args.CloudTag != "" {
		spec, err := m.makeRegionSpec(args.CloudTag, args.CloudRegion)
		if err != nil {
			return errors.Trace(err)
		}
		rspec = spec
	}
	return m.ctlrState.UpdateModelConfigDefaultValues(args.Config, nil, rspec)
}

// UnsetModelDefaults removes the specified default model settings.
func (m *ModelManagerAPI) UnsetModelDefaults(ctx context.Context, args params.UnsetModelDefaults) (params.ErrorResults, error) {
	results := params.ErrorResults{Results: make([]params.ErrorResult, len(args.Keys))}
	if !m.isAdmin {
		return results, apiservererrors.ErrPerm
	}

	if err := m.check.ChangeAllowed(ctx); err != nil {
		return results, errors.Trace(err)
	}

	for i, arg := range args.Keys {
		var rspec *environscloudspec.CloudRegionSpec
		if arg.CloudTag != "" {
			spec, err := m.makeRegionSpec(arg.CloudTag, arg.CloudRegion)
			if err != nil {
				results.Results[i].Error = apiservererrors.ServerError(
					errors.Trace(err))
				continue
			}
			rspec = spec
		}
		results.Results[i].Error = apiservererrors.ServerError(
			m.ctlrState.UpdateModelConfigDefaultValues(nil, arg.Keys, rspec),
		)
	}
	return results, nil
}

// makeRegionSpec is a helper method for methods that call
// state.UpdateModelConfigDefaultValues.
func (m *ModelManagerAPI) makeRegionSpec(cloudTag, r string) (*environscloudspec.CloudRegionSpec, error) {
	cTag, err := names.ParseCloudTag(cloudTag)
	if err != nil {
		return nil, errors.Trace(err)
	}
	rspec, err := environscloudspec.NewCloudRegionSpec(cTag.Id(), r)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return rspec, nil
}

// ChangeModelCredential changes cloud credential reference for models.
// These new cloud credentials must already exist on the controller.
func (m *ModelManagerAPI) ChangeModelCredential(ctx context.Context, args params.ChangeModelCredentialsParams) (params.ErrorResults, error) {
	if err := m.check.ChangeAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}

	err := m.authorizer.HasPermission(permission.SuperuserAccess, m.state.ControllerTag())
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return params.ErrorResults{}, errors.Trace(err)
	}
	controllerAdmin := err == nil
	// Only controller or model admin can change cloud credential on a model.
	checkModelAccess := func(tag names.ModelTag) error {
		if controllerAdmin {
			return nil
		}
		return m.authorizer.HasPermission(permission.AdminAccess, tag)
	}

	replaceModelCredential := func(arg params.ChangeModelCredentialParams) error {
		modelTag, err := names.ParseModelTag(arg.ModelTag)
		if err != nil {
			return errors.Trace(err)
		}
		if err := checkModelAccess(modelTag); err != nil {
			return errors.Trace(err)
		}
		credentialTag, err := names.ParseCloudCredentialTag(arg.CloudCredentialTag)
		if err != nil {
			return errors.Trace(err)
		}
		model, releaser, err := m.state.GetModel(modelTag.Id())
		if err != nil {
			return errors.Trace(err)
		}
		defer releaser()

		updated, err := model.SetCloudCredential(credentialTag)
		if err != nil {
			return errors.Trace(err)
		}
		if !updated {
			return errors.Errorf("model %v already uses credential %v", modelTag.Id(), credentialTag.Id())
		}
		return nil
	}

	results := make([]params.ErrorResult, len(args.Models))
	for i, arg := range args.Models {
		if err := replaceModelCredential(arg); err != nil {
			results[i].Error = apiservererrors.ServerError(err)
		}
	}
	return params.ErrorResults{Results: results}, nil
}
