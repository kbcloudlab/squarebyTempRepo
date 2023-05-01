package userrole

import (
	"errors"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	repo "squareby.com/admin/cloudspacemanager/src/repository/userrole"
)

type PrivilegeActionService struct {
	privilegeActionRepo *repo.PrivilegeActionRepo
}

func NewPrivilegeActionService() *PrivilegeActionService {
	return &PrivilegeActionService{
		privilegeActionRepo: repo.NewPrivilegeActionRepo(),
	}
}

func (_sfs *PrivilegeActionService) CreatePrivilegeAction() error {

	for _, action_data := range models.PrivilegeActionDataList {

		_, err := _sfs.privilegeActionRepo.GetPrivilegeActionByCode(action_data.Code)

		var notFoundErr *customError.NotFoundErr = &customError.NotFoundErr{}

		if errors.As(err, &notFoundErr) {
			saas_feature_model := &models.PrivilegeAction{
				Code:       action_data.Code,
				CodeNumber: action_data.CodeNumber,
				Title:      action_data.Title,
				ActionType: action_data.ActionType,
				Enabled:    true,
			}
			_, _ = _sfs.privilegeActionRepo.CreatePrivilegeAction(saas_feature_model)

		} else if err != nil {
			return err
		}

	}

	return nil

} //End of the method CreatePrivilegeAction

func (_sfs *PrivilegeActionService) UpdatePrivilegeAction(_id uint64, _privilegeActionRequest *dto.PrivilegeActionRequest) error {

	privileged_action := _privilegeActionRequest.NewUserRoleFromRequest(false)

	exist_action, err := _sfs.privilegeActionRepo.GetPrivilegeActionById(_id)

	if err != nil {
		return err
	}

	// privileged_action.Id = exist_action.Id
	// privileged_action.ActionType = exist_action.ActionType
	// privileged_action.Code = exist_action.Code
	// privileged_action.Enabled = exist_action.Enabled

	//Only Title can be changed
	exist_action.Title = privileged_action.Title

	return _sfs.privilegeActionRepo.UpdatePrivilegeAction(_id, exist_action)

} //End of the method UpdatePrivilegeAction

func (_pas *PrivilegeActionService) EnableDisablePrivilegeAction(_id uint64, _privilegeActionRequest *dto.PrivilegeActionEnableRequest) error {

	privilege_action, err := _pas.privilegeActionRepo.GetPrivilegeActionById(_id)

	if err != nil {
		return err
	}

	privilege_action.Enabled = _privilegeActionRequest.Enabled

	return _pas.privilegeActionRepo.UpdatePrivilegeAction(_id, privilege_action)

} //End of the method EnableDisablePrivilegeAction

func (_sfs *PrivilegeActionService) GetPrivilegeActionByCode(_code string) (*dto.PrivilegeActionResponse, error) {

	privileged_action, err := _sfs.privilegeActionRepo.GetPrivilegeActionByCode(_code)

	if err != nil {
		return nil, err
	}

	return dto.NewPrivilegeActionResponse(privileged_action), nil

} //End of the method GetPrivilegeActionByCode

func (_sfs *PrivilegeActionService) GetPrivilegeActionByCodeNumber(_codeNumber uint16) (*dto.PrivilegeActionResponse, error) {

	privileged_action, err := _sfs.privilegeActionRepo.GetPrivilegeActionByCodeNumber(_codeNumber)

	if err != nil {
		return nil, err
	}

	return dto.NewPrivilegeActionResponse(privileged_action), nil

} //End of the method GetPrivilegeActionByCodeNumber

func (_sfs *PrivilegeActionService) GetPrivilegeActionById(_id uint64) (*dto.PrivilegeActionResponse, error) {

	privileged_action, err := _sfs.privilegeActionRepo.GetPrivilegeActionById(_id)

	if err != nil {
		return nil, err
	}

	// return &dto.PrivilegeActionResponse{
	// 	Id:         privileged_action.Id,
	// 	ActionType: privileged_action.ActionType,
	// 	Code:       privileged_action.Code,
	// 	Title:      privileged_action.Title,
	// }, nil

	return dto.NewPrivilegeActionResponse(privileged_action), nil

} //End of the method GetPrivilegeActionById

func (_sfs *PrivilegeActionService) GetActivePrivilegeActionList() (dto.PrivilegeActionResponseList, error) {

	action_list, err := _sfs.privilegeActionRepo.GetActivePrivilegeActionList()

	if err != nil {
		return nil, err
	}

	action_res_list := dto.PrivilegeActionResponseList{}
	for _, privileged_action := range action_list {
		action_res_list = append(action_res_list, dto.NewPrivilegeActionResponse(privileged_action))
	}

	return action_res_list, nil

} //End of the method GetActivePrivilegeActionList

func (_sfs *PrivilegeActionService) GetAllPrivilegeActionList() (dto.PrivilegeActionResponseList, error) {

	action_list, err := _sfs.privilegeActionRepo.GetAllPrivilegeActionList()

	if err != nil {
		return nil, err
	}

	action_res_list := dto.PrivilegeActionResponseList{}
	for _, privileged_action := range action_list {
		action_res_list = append(action_res_list, dto.NewPrivilegeActionResponse(privileged_action))
	}

	return action_res_list, nil

} //End of the method GetAllPrivilegeActionList
