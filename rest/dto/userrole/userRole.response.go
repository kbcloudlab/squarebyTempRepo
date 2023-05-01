package userrole

import (
	"encoding/json"
	"io"

	userRoleBridge "squareby.com/admin/cloudspacemanager/sidechain/bridge/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type UserRoleResponse struct {
	Id                  string               `json:"id"`
	RoleName            string               `json:"roleName"`
	ActionBoolTypeList  []*boolTypeResponse  `json:"actionBoolTypeList"`
	ActionCountTypeList []*countTypeResponse `json:"actionCountTypeList"`
	Enabled             bool                 `json:"enabled"`
}

type boolTypeResponse struct { //Ex: "Allow order management" = false
	PrivilegeActionResponse *PrivilegeActionResponse `json:"privilegeActionResponse"`
	Allow                   bool                     `json:"allow"`
}

type countTypeResponse struct { //Ex: "Number of Cloudspaces per account" = 2
	PrivilegeActionResponse *PrivilegeActionResponse `json:"privilegeActionResponse"`
	Count                   int16                    `json:"count"`
}

type UserRoleResponseList []*UserRoleResponse

func newBoolTypeResponse(_bridge *userRoleBridge.PrivilegeActionBridge, _boolTypeModel *models.BoolType) (*boolTypeResponse, error) {

	privilege_action, err := _bridge.GetPrivilegeActionByCodeNumber(_boolTypeModel.ActionCodeNumber)
	if err != nil {
		return nil, err
	}

	return &boolTypeResponse{
		PrivilegeActionResponse: NewPrivilegeActionResponse(privilege_action),
		Allow:                   _boolTypeModel.Allow,
	}, nil

}

func newCountTypeResponse(_bridge *userRoleBridge.PrivilegeActionBridge, _countTypeModel *models.CountType) (*countTypeResponse, error) {

	privilege_action, err := _bridge.GetPrivilegeActionByCodeNumber(_countTypeModel.ActionCodeNumber)
	if err != nil {
		return nil, err
	}

	return &countTypeResponse{
		PrivilegeActionResponse: NewPrivilegeActionResponse(privilege_action),
		Count:                   _countTypeModel.Count,
	}, nil

}

func NewUserRoleResponse(_userRole *models.UserRole) *UserRoleResponse {

	privilegeaction_bridge := userRoleBridge.NewPrivilegeActionBridge()

	booltype_res_list := []*boolTypeResponse{}
	for _, booltype_model := range _userRole.ActionBoolTypeList {
		for _, privilege_action := range models.PrivilegeActionDataList {
			if booltype_model.ActionCodeNumber == privilege_action.CodeNumber {
				booltype_res, err := newBoolTypeResponse(privilegeaction_bridge, booltype_model)
				if err == nil {
					booltype_res_list = append(booltype_res_list, booltype_res)
				}
			}
		}
	}

	counttype_res_list := []*countTypeResponse{}
	for _, counttype_model := range _userRole.ActionCountTypeList {
		for _, privilege_action := range models.PrivilegeActionDataList {
			if counttype_model.ActionCodeNumber == privilege_action.CodeNumber {
				counttype_res, err := newCountTypeResponse(privilegeaction_bridge, counttype_model)
				if err == nil {
					counttype_res_list = append(counttype_res_list, counttype_res)
				}
			}
		}
	}

	return &UserRoleResponse{
		Id:                  _userRole.Id,
		RoleName:            _userRole.RoleName,
		ActionBoolTypeList:  booltype_res_list,
		ActionCountTypeList: counttype_res_list,
		Enabled:             _userRole.Enabled,
	}
}

func (_csr *UserRoleResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert UserRoleResponse to JSON:: <<cloudspaceUserRole.response.go -> (_csrl *UserRoleResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_csrl *UserRoleResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert UserRoleResponseList to JSON:: <<cloudspaceUserRole.response.go -> (_csrl *UserRoleResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
