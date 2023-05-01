package sector

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SectorMasterRequest struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (_sfr *SectorMasterRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SectorMasterRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_sfr *SectorMasterRequest) NewUserRoleFromRequest(_mapId bool) *models.SectorMaster {

	sector_master := &models.SectorMaster{
		Name: _sfr.Name,
	}

	if _mapId {
		sector_master.Id = _sfr.Id
	}

	return sector_master

}

type SectorMasterEnableRequest struct {
	Id      uint64 `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_sfr *SectorMasterEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SectorMasterEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
