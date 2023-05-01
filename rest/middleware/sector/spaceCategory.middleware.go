package sector

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SpaceCategoryMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewSpaceCategoryMiddleware() *SpaceCategoryMiddleware {
	return &SpaceCategoryMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *SpaceCategoryMiddleware) UpdateSpaceCategory(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		spacecategory_request := dto.SpaceCategoryRequest{}

		err := spacecategory_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySpaceCategory{}, spacecategory_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method UpdateSpaceCategory

func (_pfc *SpaceCategoryMiddleware) AddRemoveProductContainer(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		category_container_request := dto.CategoryProductContainerRequest{}

		err := category_container_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySpaceCategory{}, category_container_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method AddRemoveProductContainer

func (_pfc *SpaceCategoryMiddleware) EnableDisableSpaceCategory(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		spacecategory_enable_request := dto.SpaceCategoryEnableRequest{}

		err := spacecategory_enable_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySpaceCategory{}, spacecategory_enable_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method EnableDisableSpaceCategory
