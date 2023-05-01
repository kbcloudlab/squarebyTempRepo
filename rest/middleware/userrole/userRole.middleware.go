package userrole

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type UserRoleMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewUserRoleMiddleware() *UserRoleMiddleware {
	return &UserRoleMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *UserRoleMiddleware) CreateUpdateUserRole(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		userRole_request := dto.UserRoleRequest{}

		err := userRole_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeyUserRole{}, userRole_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method CreateUpdateUserRole

func (_pfc *UserRoleMiddleware) EnableDisableUserRole(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		userRole_enable_request := dto.UserRoleEnableRequest{}

		err := userRole_enable_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeyUserRole{}, userRole_enable_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method EnableDisableUserRole
