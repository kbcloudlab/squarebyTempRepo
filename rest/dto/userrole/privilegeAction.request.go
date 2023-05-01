package userrole

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type PrivilegeActionRequest struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
}

func (_sfr *PrivilegeActionRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to PrivilegeActionRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_sfr *PrivilegeActionRequest) NewUserRoleFromRequest(_mapId bool) *models.PrivilegeAction {

	saas_feature := &models.PrivilegeAction{
		Title: _sfr.Title,
	}

	if _mapId {
		saas_feature.Id = _sfr.Id
	}

	return saas_feature

}

type PrivilegeActionEnableRequest struct {
	Id      uint64 `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_sfr *PrivilegeActionEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to PrivilegeActionEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
