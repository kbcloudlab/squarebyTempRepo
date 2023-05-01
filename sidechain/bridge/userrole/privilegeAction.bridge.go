package userrole

import (
	service "squareby.com/admin/cloudspacemanager/sidechain/service/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
)

type PrivilegeActionBridge struct {
	saasFeatureService *service.PrivilegeActionService
}

func NewPrivilegeActionBridge() *PrivilegeActionBridge {
	return &PrivilegeActionBridge{
		saasFeatureService: service.NewPrivilegeActionService(),
	}
}

func (_sfb *PrivilegeActionBridge) GetPrivilegeActionByCodeNumber(_codeNumber uint16) (*models.PrivilegeAction, error) {

	return _sfb.saasFeatureService.GetPrivilegeActionByCodeNumber(_codeNumber)

} //End of the method GetPrivilegeActionById
