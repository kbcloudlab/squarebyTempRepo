package sector

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SectorMasterResponse struct {
	Id       uint64 `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
	Enabled  bool   `json:"enabled"`
}

type SectorMasterResponseList []*SectorMasterResponse

func NewSectorMasterResponse(_sectorMaster *models.SectorMaster) *SectorMasterResponse {

	return &SectorMasterResponse{
		Id:       _sectorMaster.Id,
		Code:     _sectorMaster.Code,
		Name:     _sectorMaster.Name,
		ImageURL: _sectorMaster.ImageURL,
		Enabled:  _sectorMaster.Enabled,
	}
}

func (_csr *SectorMasterResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SectorMasterResponse to JSON:: <<sectorMasterResponse.response.go -> (_csrl *SectorMasterResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_csrl *SectorMasterResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SectorMasterResponseList to JSON:: <<sectorMasterResponse.response.go -> (_csrl *SectorMasterResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
