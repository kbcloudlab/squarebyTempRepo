package userrole

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type PrivilegeActionMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewPrivilegeActionMiddleware() *PrivilegeActionMiddleware {
	return &PrivilegeActionMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *PrivilegeActionMiddleware) UpdatePrivilegeAction(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		action_request := dto.PrivilegeActionRequest{}

		err := action_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeyPrivilegeAction{}, action_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method UpdatePrivilegeAction

func (_pfc *PrivilegeActionMiddleware) EnableDisablePrivilegeAction(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		action_enable_request := dto.PrivilegeActionEnableRequest{}

		err := action_enable_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeyPrivilegeAction{}, action_enable_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method UpdatePrivilegeAction
