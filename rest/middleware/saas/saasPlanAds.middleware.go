package saas

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanAdsMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewSaasPlanAdsMiddleware() *SaasPlanAdsMiddleware {
	return &SaasPlanAdsMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *SaasPlanAdsMiddleware) CreateUpdateSaasPlanAds(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		saasplan_request := dto.SaasPlanAdsRequest{}

		err := saasplan_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySaasPlanAds{}, saasplan_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method CreateUpdateSaasPlanAds

func (_pfc *SaasPlanAdsMiddleware) EnableDisableSaasPlanAds(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		saasplan_request := dto.SaasPlanAdsEnableRequest{}

		err := saasplan_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySaasPlanAds{}, saasplan_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method EnableDisableSaasPlanAds

func (_pfc *SaasPlanAdsMiddleware) PublishSaasPlanAds(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		_next.ServeHTTP(rw, r)
	})
} //End of the method PublishSaasPlanAds
