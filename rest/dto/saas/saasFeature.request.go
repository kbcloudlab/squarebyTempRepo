package saas

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasFeatureRequest struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
}

func (_sfr *SaasFeatureRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasFeatureRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_sfr *SaasFeatureRequest) NewUserRoleFromRequest(_mapId bool) *models.SaasFeature {

	saas_feature := &models.SaasFeature{
		Title: _sfr.Title,
	}

	if _mapId {
		saas_feature.Id = _sfr.Id
	}

	return saas_feature

}

type SaasFeatureEnableRequest struct {
	Id      uint64 `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_sfr *SaasFeatureEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasFeatureEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
