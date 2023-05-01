package userrole

import (
	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	repo "squareby.com/admin/cloudspacemanager/src/repository/userrole"
)

type UserRoleService struct {
	userRoleRepo *repo.UserRoleRepo
}

func NewUserRoleService() *UserRoleService {
	return &UserRoleService{
		userRoleRepo: repo.NewUserRoleRepo(),
	}
}

func (_urs *UserRoleService) CreateUserRole(_userRoleRequest *dto.UserRoleRequest) (string, error) {

	user_role := _userRoleRequest.NewUserRoleFromRequest(false)
	user_role.Enabled = true

	return _urs.userRoleRepo.CreateUserRole(user_role)

} //End of the method CreateUserRole

func (_urs *UserRoleService) UpdateUserRole(_id string, _userRoleRequest *dto.UserRoleRequest) error {

	exists_userrole, err := _urs.userRoleRepo.GetUserRoleById(_id)
	if err != nil {
		return err
	}

	user_role := _userRoleRequest.NewUserRoleFromRequest(false)
	user_role.Enabled = exists_userrole.Enabled

	return _urs.userRoleRepo.UpdateUserRole(_id, user_role.NewDuplicateModel())

} //End of the method UpdateUserRole

func (_urs *UserRoleService) GetUserRoleById(_id string) (*dto.UserRoleResponse, error) {

	user_role, err := _urs.userRoleRepo.GetUserRoleById(_id)

	if err != nil {
		return nil, err
	}

	return dto.NewUserRoleResponse(user_role), nil

} //End of the method GetUserRoleById

func (_urs *UserRoleService) GetActiveUserRoleList() (dto.UserRoleResponseList, error) {

	user_role_list, err := _urs.userRoleRepo.GetActiveUserRoleList()

	if err != nil {
		return nil, err
	}

	userrole_res_list := dto.UserRoleResponseList{}

	for _, userrole := range user_role_list {
		userrole_res_list = append(userrole_res_list, dto.NewUserRoleResponse(userrole))
	}

	return userrole_res_list, nil

} //End of the method GetActiveUserRoleList

func (_urs *UserRoleService) GetAllUserRoleList() (dto.UserRoleResponseList, error) {

	user_role_list, err := _urs.userRoleRepo.GetAllUserRoleList()

	if err != nil {
		return nil, err
	}

	userrole_res_list := dto.UserRoleResponseList{}

	for _, userrole := range user_role_list {
		userrole_res_list = append(userrole_res_list, dto.NewUserRoleResponse(userrole))
	}

	return userrole_res_list, nil

} //End of the method GetAllUserRoleList

func (_urs *UserRoleService) EnableDisableUserRole(_id string, _userRoleEnableRequest *dto.UserRoleEnableRequest) error {

	userrole, err := _urs.userRoleRepo.GetUserRoleById(_id)

	if err != nil {
		return err
	}

	userrole.Enabled = _userRoleEnableRequest.Enabled

	return _urs.userRoleRepo.UpdateUserRole(_id, userrole.NewDuplicateModel())

} //End of the method EnableDisableUserRole

func (_urs *UserRoleService) DeleteUserRole(_id string) error {

	return _urs.userRoleRepo.DeleteUserRole(_id)

} //End of the method DeleteUserRole
