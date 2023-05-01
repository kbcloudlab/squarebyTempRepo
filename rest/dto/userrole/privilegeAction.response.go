package userrole

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type PrivilegeActionResponse struct {
	Id         uint64 `json:"id"`
	Code       string `json:"code"`
	CodeNumber uint16 `json:"codeNumber"`
	Title      string `json:"title"`
	ActionType int16  `json:"actionType"`
	Enabled    bool   `json:"enabled"`
}

type PrivilegeActionResponseList []*PrivilegeActionResponse

func NewPrivilegeActionResponse(_privilegeAction *models.PrivilegeAction) *PrivilegeActionResponse {

	return &PrivilegeActionResponse{
		Id:         _privilegeAction.Id,
		Code:       _privilegeAction.Code,
		CodeNumber: _privilegeAction.CodeNumber,
		Title:      _privilegeAction.Title,
		ActionType: _privilegeAction.ActionType,
		Enabled:    _privilegeAction.Enabled,
	}
}

func (_csr *PrivilegeActionResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert PrivilegeActionResponse to JSON:: <<privilegeActionResponse.response.go -> (_csrl *PrivilegeActionResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_csrl *PrivilegeActionResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert PrivilegeActionResponseList to JSON:: <<privilegeActionResponse.response.go -> (_csrl *PrivilegeActionResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
