// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/juju/errors"
	"github.com/juju/names/v4"

	"github.com/juju/juju/apiserver/common/credentialcommon"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/core/changestream"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/domain/credential"
	credentialerrors "github.com/juju/juju/domain/credential/errors"
	"github.com/juju/juju/domain/model"
)

// WatcherFactory instances return a watcher for a specified credential UUID,
type WatcherFactory interface {
	NewValueWatcher(
		namespace, uuid string, changeMask changestream.ChangeType,
	) (watcher.NotifyWatcher, error)
}

// State describes retrieval and persistence methods for credentials.
type State interface {
	// UpsertCloudCredential adds or updates a cloud credential with the given name, cloud, owner.
	// If the credential already exists, the existing credential's value of Invalid is returned.
	UpsertCloudCredential(ctx context.Context, name, cloudName, owner string, credential cloud.Credential) (*bool, error)

	// InvalidateCloudCredential marks the cloud credential for the given name, cloud, owner as invalid.
	InvalidateCloudCredential(ctx context.Context, name, cloudName, owner, reason string) error

	// CloudCredentialsForOwner returns the owner's cloud credentials for a given cloud,
	// keyed by credential name.
	CloudCredentialsForOwner(ctx context.Context, owner, cloudName string) (map[string]cloud.Credential, error)

	// CloudCredential returns the cloud credential for the given name, cloud, owner.
	CloudCredential(ctx context.Context, name, cloudName, owner string) (cloud.Credential, error)

	// AllCloudCredentialsForOwner returns all cloud credentials stored on the controller
	// for a given owner.
	AllCloudCredentialsForOwner(ctx context.Context, owner string) ([]credential.CloudCredential, error)

	// RemoveCloudCredential removes a cloud credential with the given name, cloud, owner.
	RemoveCloudCredential(ctx context.Context, name, cloudName, owner string) error

	// WatchCredential returns a new NotifyWatcher watching for changes to the specified credential.
	WatchCredential(
		ctx context.Context,
		getWatcher func(string, string, changestream.ChangeType) (watcher.NotifyWatcher, error),
		name, cloudName, owner string,
	) (watcher.NotifyWatcher, error)

	// ModelsUsingCloudCredential returns a map of uuid->name for models which use the credential.
	ModelsUsingCloudCredential(ctx context.Context, id credential.ID) (map[model.UUID]string, error)
}

// ValidationContextGetter returns the artefacts for a specified model, used to make credential validation calls.
type ValidationContextGetter func(modelUUID model.UUID) (credentialcommon.CredentialValidationContext, error)

// Logger facilitates emitting log messages.
type Logger interface {
	Debugf(string, ...interface{})
}

// Service provides the API for working with external controllers.
type Service struct {
	st             State
	watcherFactory WatcherFactory
	logger         Logger

	// These are set via options after the service is created.

	validationContextGetter ValidationContextGetter
	validator               credentialcommon.CredentialValidator

	// TODO(wallyworld) - remove when models are out of mongo
	legacyUpdater func(tag names.CloudCredentialTag) error
	legacyRemover func(tag names.CloudCredentialTag) error
}

// NewService returns a new service reference wrapping the input state.
func NewService(st State, watcherFactory WatcherFactory, logger Logger) *Service {
	return &Service{
		st:             st,
		watcherFactory: watcherFactory,
		logger:         logger,
		validator:      noopValidator{},
	}
}

// noopValidator always validates any credential.
// TODO(wallyworld) - this is a placeholder.
type noopValidator struct {
}

func (noopValidator) Validate(
	_ credentialcommon.CredentialValidationContext,
	_ credential.ID,
	_ *cloud.Credential,
	_ bool,
) (machineErrors []error, err error) {
	return nil, nil
}

// WithValidationContextGetter configures the service to use the specified function
// to get a context used to validate a credential for a specified model.
// TODO(wallyworld) - remove when models are out of mongo
func (s *Service) WithValidationContextGetter(validationContextGetter ValidationContextGetter) *Service {
	s.validationContextGetter = validationContextGetter
	return s
}

// WithCredentialValidator configures the service to use the specified
// credential validator.
func (s *Service) WithCredentialValidator(validator credentialcommon.CredentialValidator) *Service {
	s.validator = validator
	return s
}

// WithLegacyUpdater configures the service to use the specified function
// to update credential details in mongo.
// TODO(wallyworld) - remove when models are out of mongo
func (s *Service) WithLegacyUpdater(updater func(tag names.CloudCredentialTag) error) *Service {
	s.legacyUpdater = updater
	return s
}

// WithLegacyRemover configures the service to use the specified function
// to remove credential details from mongo.
// TODO(wallyworld) - remove when models are out of mongo
func (s *Service) WithLegacyRemover(remover func(tag names.CloudCredentialTag) error) *Service {
	s.legacyRemover = remover
	return s
}

// CloudCredential returns the cloud credential for the given tag.
func (s *Service) CloudCredential(ctx context.Context, tag names.CloudCredentialTag) (cloud.Credential, error) {
	return s.st.CloudCredential(ctx, tag.Name(), tag.Cloud().Id(), tag.Owner().Id())
}

// AllCloudCredentialsForOwner returns all cloud credentials stored on the controller
// for a given owner.
func (s *Service) AllCloudCredentialsForOwner(ctx context.Context, owner string) ([]credential.CloudCredential, error) {
	creds, err := s.st.AllCloudCredentialsForOwner(ctx, owner)
	if err != nil {
		return nil, errors.Trace(err)
	}
	result := make([]credential.CloudCredential, len(creds))
	for i, c := range creds {
		result[i] = credential.CloudCredential{Credential: c.Credential, CloudName: c.CloudName}
	}
	return result, nil
}

// CloudCredentialsForOwner returns the owner's cloud credentials for a given cloud,
// keyed by credential name.
func (s *Service) CloudCredentialsForOwner(ctx context.Context, owner, cloudName string) (map[string]cloud.Credential, error) {
	return s.st.CloudCredentialsForOwner(ctx, owner, cloudName)
}

// UpdateCloudCredential adds or updates a cloud credential with the given tag.
func (s *Service) UpdateCloudCredential(ctx context.Context, tag names.CloudCredentialTag, cred cloud.Credential) error {
	_, err := s.st.UpsertCloudCredential(ctx, tag.Name(), tag.Cloud().Id(), tag.Owner().Id(), cred)
	return err
}

// RemoveCloudCredential removes a cloud credential with the given tag.
func (s *Service) RemoveCloudCredential(ctx context.Context, tag names.CloudCredentialTag) error {
	return s.st.RemoveCloudCredential(ctx, tag.Name(), tag.Cloud().Id(), tag.Owner().Id())
}

// InvalidateCredential marks the cloud credential for the given name, cloud, owner as invalid.
func (s *Service) InvalidateCredential(ctx context.Context, tag names.CloudCredentialTag, reason string) error {
	return s.st.InvalidateCloudCredential(ctx, tag.Name(), tag.Cloud().Id(), tag.Owner().Id(), reason)
}

// WatchCredential returns a watcher that observes changes to the specified credential.
func (s *Service) WatchCredential(ctx context.Context, tag names.CloudCredentialTag) (watcher.NotifyWatcher, error) {
	if s.watcherFactory != nil {
		return s.st.WatchCredential(ctx, s.watcherFactory.NewValueWatcher, tag.Name(), tag.Cloud().Id(), tag.Owner().Id())
	}
	return nil, errors.NotYetAvailablef("credential watcher")
}

func (s *Service) modelsUsingCredential(ctx context.Context, id credential.ID) (map[model.UUID]string, error) {
	models, err := s.st.ModelsUsingCloudCredential(ctx, id)
	if err != nil && !errors.Is(err, errors.NotFound) {
		return nil, errors.Trace(err)
	}
	return models, nil
}

func (s *Service) validateCredentialForModel(modelUUID model.UUID, id credential.ID, cred *cloud.Credential) ([]error, error) {
	if s.validator == nil || s.validationContextGetter == nil {
		return nil, errors.Errorf("missing validation helpers")
	}
	ctx, err := s.validationContextGetter(modelUUID)
	if err != nil {
		return []error{errors.Trace(err)}, nil
	}

	modelErrors, err := s.validator.Validate(ctx, id, cred, false)
	if err != nil {
		return []error{errors.Trace(err)}, nil
	}
	return modelErrors, nil
}

// UpdateCredentialModelResult holds details of a model
// which was affected by a credential update, and any
// errors encountered validating the credential.
type UpdateCredentialModelResult struct {
	// ModelUUID contains model's UUID.
	ModelUUID model.UUID

	// ModelName contains model name.
	ModelName string

	// Errors contains the errors accumulated while trying to update a credential.
	Errors []error
}

// CheckAndUpdateCredential updates the credential after first checking that any models which use the credential
// can still access the cloud resources. If force is true, update the credential even if there are issues
// validating the credential.
// TODO(wallyworld) - the validation getter can be set during service construction once dqlite is used everywhere.
// Note - it is expected that `WithValidationContextGetter` is called to set up the service to have a non-nil
// validationContextGetter prior to calling this function, or else an error will be returned.
// TODO(wallyworld) - we need a strategy to handle changes which occur after the affected models have been read
// but before validation can complete.
func (s *Service) CheckAndUpdateCredential(ctx context.Context, id credential.ID, cred cloud.Credential, force bool) ([]UpdateCredentialModelResult, error) {
	if s.validationContextGetter == nil {
		return nil, errors.New("cannot validate credential with nil context getter")
	}

	models, err := s.modelsUsingCredential(ctx, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var (
		modelsErred  bool
		modelsResult []UpdateCredentialModelResult
	)
	for uuid, name := range models {
		result := UpdateCredentialModelResult{
			ModelUUID: uuid,
			ModelName: name,
		}
		result.Errors, err = s.validateCredentialForModel(uuid, id, &cred)
		if err != nil {
			return nil, errors.Trace(err)
		}
		modelsResult = append(modelsResult, result)
		if len(result.Errors) > 0 {
			modelsErred = true
		}
	}
	// Since we get a map above, for consistency ensure that models are added
	// sorted by model uuid.
	sort.Slice(modelsResult, func(i, j int) bool {
		return modelsResult[i].ModelUUID < modelsResult[j].ModelUUID
	})

	if modelsErred && !force {
		return modelsResult, credentialerrors.CredentialModelValidation
	}

	existingInvalid, err := s.st.UpsertCloudCredential(ctx, id.Name, id.Cloud, id.Owner, cred)
	if err != nil {
		if errors.Is(err, errors.NotFound) {
			err = fmt.Errorf("%w %q for credential %q", credentialerrors.UnknownCloud, id.Name, id.Cloud)
		}
		return nil, errors.Trace(err)
	}
	if s.legacyUpdater == nil || cred.Invalid {
		return modelsResult, nil
	}

	// Credential is valid - revoke the suspended status of any relevant models.

	// TODO(wallyworld) - we still manage models in mongo.
	// This can be removed after models are in dqlite.
	// Existing credential will become valid after this call, and
	// the model status of all models that use it will be reverted.
	if existingInvalid != nil && *existingInvalid {
		tag, err := id.Tag()
		if err != nil {
			return nil, errors.Trace(err)
		}
		if err := s.legacyUpdater(tag); err != nil {
			return nil, errors.Trace(err)
		}
	}
	return modelsResult, nil
}

// CheckAndRevokeCredential removes the credential after first checking that any models which use the credential
// can still access the cloud resources. If force is true, update the credential even if there are issues
// validating the credential.
// TODO(wallyworld) - we need a strategy to handle changes which occur after the affected models have been read
// but before validation can complete.
func (s *Service) CheckAndRevokeCredential(ctx context.Context, id credential.ID, force bool) error {
	models, err := s.modelsUsingCredential(ctx, id)
	if err != nil {
		return errors.Trace(err)
	}
	if len(models) != 0 {
		opMessage := "cannot be deleted as"
		if force {
			opMessage = "will be deleted but"
		}
		s.logger.Debugf("credential %v %v it is used by model%v",
			id,
			opMessage,
			modelsPretty(models),
		)
		if !force {
			// Some models still use this credential - do not delete this credential...
			return errors.Errorf("cannot revoke credential %v: it is still used by %d model%v", id, len(models), plural(len(models)))
		}
	}
	err = s.st.RemoveCloudCredential(ctx, id.Name, id.Cloud, id.Owner)
	if err != nil || s.legacyRemover == nil {
		return errors.Trace(err)
	} else {
		// If credential was successfully removed, we also want to clear all references to it from the models.
		tag, err := id.Tag()
		if err != nil {
			return errors.Trace(err)
		}
		if err := s.legacyRemover(tag); err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

func plural(length int) string {
	if length == 1 {
		return ""
	}
	return "s"
}

func modelsPretty(in map[model.UUID]string) string {
	// map keys are notoriously randomly ordered
	uuids := []string{}
	for uuid := range in {
		uuids = append(uuids, string(uuid))
	}
	sort.Strings(uuids)

	firstLine := ":\n- "
	if len(uuids) == 1 {
		firstLine = " "
	}

	return fmt.Sprintf("%v%v%v",
		plural(len(in)),
		firstLine,
		strings.Join(uuids, "\n- "),
	)
}