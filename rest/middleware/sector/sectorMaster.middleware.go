package sector

import (
	"context"
	"net/http"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SectorMasterMiddleware struct {
	logging     *customLog.Logging
	errorCustom *customError.ErrorCustom
}

func NewSectorMasterMiddleware() *SectorMasterMiddleware {
	return &SectorMasterMiddleware{
		logging:     customLog.NewLogging(),
		errorCustom: customError.NewErrorCustom(),
	}
}

func (_pfc *SectorMasterMiddleware) UpdateSectorMaster(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		sectormaster_request := dto.SectorMasterRequest{}

		err := sectormaster_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySectorMaster{}, sectormaster_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method UpdateSectorMaster

func (_pfc *SectorMasterMiddleware) EnableDisableSectorMaster(_next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		sectormaster_enable_request := dto.SectorMasterEnableRequest{}

		err := sectormaster_enable_request.FromJSON(r.Body)

		if err != nil {
			_pfc.logging.ErrorLog("", err)
			http.Error(rw, _pfc.errorCustom.GetErrorData(err).ClientMessage, _pfc.errorCustom.GetErrorData(err).StatusCode)
			return
		}

		ctx := context.WithValue(r.Context(), models.KeySectorMaster{}, sectormaster_enable_request)
		r = r.WithContext(ctx)

		_next.ServeHTTP(rw, r)
	})
} //End of the method EnableDisableSectorMaster
