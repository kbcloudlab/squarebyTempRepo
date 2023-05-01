package userrole

import (
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"

	repo "squareby.com/admin/cloudspacemanager/src/repository/userrole"
)

type PrivilegeActionService struct {
	saasFeatureRepo *repo.PrivilegeActionRepo
}

func NewPrivilegeActionService() *PrivilegeActionService {
	return &PrivilegeActionService{
		saasFeatureRepo: repo.NewPrivilegeActionRepo(),
	}
}

func (_sfs *PrivilegeActionService) GetPrivilegeActionByCodeNumber(_codeNumber uint16) (*models.PrivilegeAction, error) {

	return _sfs.saasFeatureRepo.GetPrivilegeActionByCodeNumber(_codeNumber)

} //End of the method GetPrivilegeActionByCodeNumber
