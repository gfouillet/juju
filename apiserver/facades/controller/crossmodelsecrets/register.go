// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package crossmodelsecrets

import (
	"context"
	"fmt"
	"reflect"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/crossmodel"
	"github.com/juju/juju/apiserver/facade"
	corelogger "github.com/juju/juju/core/logger"
	coremodel "github.com/juju/juju/core/model"
	secretservice "github.com/juju/juju/domain/secret/service"
	secretbackendservice "github.com/juju/juju/domain/secretbackend/service"
	"github.com/juju/juju/internal/secrets/provider"
)

// Register is called to expose a package of facades onto a given registry.
func Register(registry facade.FacadeRegistry) {
	registry.MustRegisterForMultiModel("CrossModelSecrets", 1, func(stdCtx context.Context, ctx facade.MultiModelContext) (facade.Facade, error) {
		api, err := makeStateCrossModelSecretsAPI(stdCtx, ctx)
		if err != nil {
			return nil, fmt.Errorf("creating CrossModelSecrets facade: %w", err)
		}
		return api, nil
	}, reflect.TypeOf((*CrossModelSecretsAPI)(nil)))
}

// makeStateCrossModelSecretsAPI creates a new server-side CrossModelSecrets API facade
// backed by global state.
func makeStateCrossModelSecretsAPI(stdCtx context.Context, ctx facade.MultiModelContext) (*CrossModelSecretsAPI, error) {
	authCtxt := ctx.Resources().Get("offerAccessAuthContext").(*common.ValueResource).Value
	serviceFactory := ctx.ServiceFactory()

	backendService := serviceFactory.SecretBackend()
	secretInfoGetter := func(modelUUID string) SecretService {
		secretBackendAdminConfigGetter := func(stdCtx context.Context) (*provider.ModelBackendConfigInfo, error) {
			return backendService.GetSecretBackendConfigForAdmin(stdCtx, coremodel.UUID(modelUUID))
		}
		backendUserSecretConfigGetter := func(
			stdCtx context.Context, gsg secretservice.GrantedSecretsGetter, accessor secretservice.SecretAccessor,
		) (*provider.ModelBackendConfigInfo, error) {
			return backendService.BackendConfigInfo(stdCtx, secretbackendservice.BackendConfigParams{
				GrantedSecretsGetter: gsg,
				Accessor:             accessor,
				ModelUUID:            coremodel.UUID(modelUUID),
				SameController:       true,
			})
		}
		return ctx.ServiceFactoryForModel(coremodel.UUID(modelUUID)).Secret(secretBackendAdminConfigGetter, backendUserSecretConfigGetter)
	}

	modelInfo, err := serviceFactory.ModelInfo().GetModelInfo(stdCtx)
	if err != nil {
		return nil, fmt.Errorf("retrieving model info: %w", err)
	}

	st := ctx.State()
	return NewCrossModelSecretsAPI(
		ctx.Resources(),
		authCtxt.(*crossmodel.AuthContext),
		st.ControllerUUID(),
		modelInfo.UUID,
		secretInfoGetter,
		backendService,
		&crossModelShim{st.RemoteEntities()},
		&stateBackendShim{st},
		ctx.Logger().Child("crossmodelsecrets", corelogger.SECRETS),
	)
}
