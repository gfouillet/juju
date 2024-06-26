// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package apiserver

import (
	"context"
	"net/http"

	"github.com/go-macaroon-bakery/macaroon-bakery/v3/bakery"
	"github.com/go-macaroon-bakery/macaroon-bakery/v3/bakery/checkers"
	"github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"
	"github.com/juju/errors"

	"github.com/juju/juju/apiserver/apiserverhttp"
	"github.com/juju/juju/apiserver/common/crossmodel"
	"github.com/juju/juju/core/macaroon"
	"github.com/juju/juju/state"
)

const (
	localOfferAccessLocationPath = "/offeraccess"
)

type localOfferAuthHandler struct {
	authCtx *crossmodel.AuthContext
}

func addOfferAuthHandlers(offerAuthCtxt *crossmodel.AuthContext, mux *apiserverhttp.Mux) {
	appOfferHandler := &localOfferAuthHandler{authCtx: offerAuthCtxt}
	appOfferDischargeMux := http.NewServeMux()

	discharger := httpbakery.NewDischarger(
		httpbakery.DischargerParams{
			Key:     offerAuthCtxt.OfferThirdPartyKey(),
			Checker: httpbakery.ThirdPartyCaveatCheckerFunc(appOfferHandler.checkThirdPartyCaveat),
		})
	discharger.AddMuxHandlers(appOfferDischargeMux, localOfferAccessLocationPath)

	_ = mux.AddHandler("POST", localOfferAccessLocationPath+"/discharge", appOfferDischargeMux)
	_ = mux.AddHandler("GET", localOfferAccessLocationPath+"/publickey", appOfferDischargeMux)
}

func newOfferAuthcontext(pool *state.StatePool) (*crossmodel.AuthContext, error) {
	// Create a bakery service for discharging third-party caveats for
	// local offer access authentication. This service does not persist keys;
	// its macaroons should be very short-lived.
	st, err := pool.SystemState()
	if err != nil {
		return nil, errors.Trace(err)
	}
	location := "juju model " + st.ModelUUID()
	checker := checkers.New(macaroon.MacaroonNamespace)

	// Create a bakery service for local offer access authentication. This service
	// persists keys into MongoDB in a TTL collection.
	store, err := st.NewBakeryStorage()
	if err != nil {
		return nil, errors.Trace(err)
	}
	bakeryConfig := st.NewBakeryConfig()
	key, err := bakeryConfig.GetOffersThirdPartyKey()
	if err != nil {
		return nil, errors.Trace(err)
	}

	controllerConfig, err := st.ControllerConfig()
	if err != nil {
		return nil, errors.Annotate(err, "unable to get controller config")
	}
	loginTokenRefreshURL := controllerConfig.LoginTokenRefreshURL()
	if loginTokenRefreshURL != "" {
		offerBakery, err := crossmodel.NewJaaSOfferBakery(
			loginTokenRefreshURL, location, bakeryConfig, store, checker,
		)
		if err != nil {
			return nil, errors.Trace(err)
		}
		return crossmodel.NewAuthContext(crossmodel.GetBackend(st), key, offerBakery)
	}
	offerBakery, err := crossmodel.NewLocalOfferBakery(location, bakeryConfig, store, checker)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return crossmodel.NewAuthContext(crossmodel.GetBackend(st), key, offerBakery)
}

func (h *localOfferAuthHandler) checkThirdPartyCaveat(stdctx context.Context, req *http.Request, cavInfo *bakery.ThirdPartyCaveatInfo, _ *httpbakery.DischargeToken) ([]checkers.Caveat, error) {
	logger.Debugf("check offer third party caveat %q", cavInfo.Condition)
	details, err := h.authCtx.CheckOfferAccessCaveat(string(cavInfo.Condition))
	if err != nil {
		return nil, errors.Trace(err)
	}

	firstPartyCaveats, err := h.authCtx.CheckLocalAccessRequest(details)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return firstPartyCaveats, nil
}
