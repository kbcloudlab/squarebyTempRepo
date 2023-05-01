package saas

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasFeatureMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewSaasFeatureMiddleware() *SaasFeatureMiddleware {
	return &SaasFeatureMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *SaasFeatureMiddleware) UpdateSaasFeature(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		saasfeature_request := dto.SaasFeatureRequest{}

		err := saasfeature_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySaasFeature{}, saasfeature_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method UpdateSaasFeature

func (_pfc *SaasFeatureMiddleware) EnableDisableSaasFeature(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		saasfeature_enable_request := dto.SaasFeatureEnableRequest{}

		err := saasfeature_enable_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySaasFeature{}, saasfeature_enable_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method EnableDisableSaasFeature
