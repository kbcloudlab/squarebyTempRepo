package userrole

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type UserRoleRequest struct {
	Id                  string              `json:"id"`
	RoleName            string              `json:"roleName"`
	ActionBoolTypeList  []*boolTypeRequest  `json:"actionBoolTypeList"`
	ActionCountTypeList []*countTypeRequest `json:"actionCountTypeList"`
}

type countTypeRequest struct { //Ex: "Number of Cloudspaces per account" = 2
	ActionCodeNumber uint16 `json:"actionCodeNumber"`
	Count            int16  `json:"count"`
}

type boolTypeRequest struct { //Ex: "Allow order management" = false
	ActionCodeNumber uint16 `json:"actionCodeNumber"`
	Allow            bool   `json:"allow"`
}

func (_urr *UserRoleRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_urr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to UserRoleRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_urr *UserRoleRequest) NewUserRoleFromRequest(_mapId bool) *models.UserRole {

	action_counttype_list := []*models.CountType{}
	for _, counttype_req := range _urr.ActionCountTypeList {
		action_counttype_list = append(action_counttype_list, &models.CountType{
			ActionCodeNumber: counttype_req.ActionCodeNumber,
			Count:            counttype_req.Count,
		})
	}

	action_booltype_list := []*models.BoolType{}
	for _, booltype_req := range _urr.ActionBoolTypeList {
		action_booltype_list = append(action_booltype_list, &models.BoolType{
			ActionCodeNumber: booltype_req.ActionCodeNumber,
			Allow:            booltype_req.Allow,
		})
	}

	userrole := &models.UserRole{
		RoleName:            _urr.RoleName,
		ActionBoolTypeList:  action_booltype_list,
		ActionCountTypeList: action_counttype_list,
	}

	if _mapId {
		userrole.Id = _urr.Id
	}

	return userrole

}

type UserRoleEnableRequest struct {
	Id      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_urr *UserRoleEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_urr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to UserRoleEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
