package saas

import (
	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"

	repo "squareby.com/admin/cloudspacemanager/src/repository/saas"
)

type SaasPlanService struct {
	saasPlanRepo *repo.SaasPlanRepo
}

func NewSaasPlanService() *SaasPlanService {
	return &SaasPlanService{
		saasPlanRepo: repo.NewSaasPlanRepo(),
	}
}

func (_sps *SaasPlanService) CreateSaasPlan(_saasPlanRequest *dto.SaasPlanRequest) (string, error) {

	saasplan_model := _saasPlanRequest.NewSaasPlanFromRequest(false)
	saasplan_model.Enabled = true

	return _sps.saasPlanRepo.CreateSaasPlan(saasplan_model)

} //End of the method CreateSaasPlan

func (_sps *SaasPlanService) UpdateSaasPlan(_id string, _saasPlanRequest *dto.SaasPlanRequest) error {

	exists_saas_plan, err := _sps.saasPlanRepo.GetSaasPlanById(_id)
	if err != nil {
		return err
	}

	saasplan_model := _saasPlanRequest.NewSaasPlanFromRequest(false)
	saasplan_model.Enabled = exists_saas_plan.Enabled

	return _sps.saasPlanRepo.UpdateSaasPlan(_id, saasplan_model)

} //End of the method UpdateSaasPlan

func (_sps *SaasPlanService) EnableDisableSaasPlan(_id string, _saasPlanEnableReq *dto.SaasPlanEnableRequest) error {

	saas_plan, err := _sps.saasPlanRepo.GetSaasPlanById(_id)
	if err != nil {
		return err
	}

	saas_plan.Enabled = _saasPlanEnableReq.Enabled

	return _sps.saasPlanRepo.UpdateSaasPlan(_id, saas_plan)

} //End of the method EnableDisableSaasPlan

func (_sps *SaasPlanService) GetSaasPlanById(_id string) (*dto.SaasPlanResponse, error) {

	saasplan_model, err := _sps.saasPlanRepo.GetSaasPlanById(_id)

	if err != nil {
		return nil, err
	}

	saas_plan_res, err := dto.NewSaasPlanResponse(saasplan_model)

	if err != nil {
		return nil, err
	}

	return saas_plan_res, nil

} //End of the method GetSaasPlanById

func (_sps *SaasPlanService) GetActiveSaasPlanList() (dto.SaasPlanResponse1List, error) {

	saasplan_model_list, err := _sps.saasPlanRepo.GetActiveSaasPlanList()

	if err != nil {
		return nil, err
	}

	saas_plan_res_list := dto.SaasPlanResponse1List{}
	for _, saas_plan_model := range saasplan_model_list {
		saas_plan_res := dto.NewSaasPlanResponse1(saas_plan_model)

		if err == nil {
			saas_plan_res_list = append(saas_plan_res_list, saas_plan_res)
		}
	}

	return saas_plan_res_list, nil

} //End of the method GetActiveSaasPlanList

func (_sps *SaasPlanService) GetAllSaasPlanList() (dto.SaasPlanResponse1List, error) {

	saasplan_model_list, err := _sps.saasPlanRepo.GetAllSaasPlanList()

	if err != nil {
		return nil, err
	}

	saas_plan_res_list := dto.SaasPlanResponse1List{}
	for _, saas_plan_model := range saasplan_model_list {
		saas_plan_res := dto.NewSaasPlanResponse1(saas_plan_model)

		if err == nil {
			saas_plan_res_list = append(saas_plan_res_list, saas_plan_res)
		}
	}

	return saas_plan_res_list, nil

} //End of the method GetAllSaasPlanList

func (_sps *SaasPlanService) DeleteSaasPlan(_id string) error {

	return _sps.saasPlanRepo.DeleteSaasPlan(_id)

} //End of the method DeleteSaasPlan
