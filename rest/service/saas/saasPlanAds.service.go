package saas

import (
	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"

	repo "squareby.com/admin/cloudspacemanager/src/repository/saas"
)

type SaasPlanAdsService struct {
	saasPlanAdsRepo *repo.SaasPlanAdsRepo
}

func NewSaasPlanAdsService() *SaasPlanAdsService {
	return &SaasPlanAdsService{
		saasPlanAdsRepo: repo.NewSaasPlanAdsRepo(),
	}
}

func (_sps *SaasPlanAdsService) CreateSaasPlanAds(_saasPlanAdsRequest *dto.SaasPlanAdsRequest) (string, error) {

	saasplan_model := _saasPlanAdsRequest.NewSaasPlanAdsFromRequest(false)

	return _sps.saasPlanAdsRepo.CreateSaasPlanAds(saasplan_model)

} //End of the method CreateSaasPlanAds

func (_sps *SaasPlanAdsService) UpdateSaasPlanAds(_id string, _saasPlanAdsRequest *dto.SaasPlanAdsRequest) error {

	exists_saas_plan, err := _sps.saasPlanAdsRepo.GetSaasPlanAdsById(_id)
	if err != nil {
		return err
	}

	saasplan_model := _saasPlanAdsRequest.NewSaasPlanAdsFromRequest(false)
	saasplan_model.Enabled = exists_saas_plan.Enabled

	return _sps.saasPlanAdsRepo.UpdateSaasPlanAds(_id, saasplan_model)

} //End of the method UpdateSaasPlanAds

func (_sps *SaasPlanAdsService) PublishSaasPlanAds(_id string) error {

	saas_plan_ads, err := _sps.saasPlanAdsRepo.GetSaasPlanAdsById(_id)
	if err != nil {
		return err
	}

	saas_plan_ads.Published = true

	return _sps.saasPlanAdsRepo.UpdateSaasPlanAds(_id, saas_plan_ads.NewDuplicateModel())

} //End of the method EnableDisableSaasPlanAds

func (_sps *SaasPlanAdsService) EnableDisableSaasPlanAds(_id string, _saasPlanEnableRequest *dto.SaasPlanAdsEnableRequest) error {

	return _sps.saasPlanAdsRepo.EnableDisableSaasPlanAds(_id, _saasPlanEnableRequest.Enabled)

} //End of the method EnableDisableSaasPlanAds

func (_sps *SaasPlanAdsService) GetSaasPlanAdsById(_id string) (*dto.SaasPlanAdsResponse, error) {

	saasplan_model, err := _sps.saasPlanAdsRepo.GetSaasPlanAdsById(_id)

	if err != nil {
		return nil, err
	}

	saas_plan_res, err := dto.NewSaasPlanAdsResponse(saasplan_model)

	if err != nil {
		return nil, err
	}

	return saas_plan_res, nil

} //End of the method GetSaasPlanAdsById

func (_sps *SaasPlanAdsService) GetActiveSaasPlanAdsList() (dto.SaasPlanAdsResponse1List, error) {

	saasplan_model_list, err := _sps.saasPlanAdsRepo.GetActiveSaasPlanAdsList()

	if err != nil {
		return nil, err
	}

	saas_plan_res_list := dto.SaasPlanAdsResponse1List{}
	for _, saas_plan_model := range saasplan_model_list {
		saas_plan_res := dto.NewSaasPlanAdsResponse1(saas_plan_model)

		if err == nil {
			saas_plan_res_list = append(saas_plan_res_list, saas_plan_res)
		}
	}

	return saas_plan_res_list, nil

} //End of the method GetActiveSaasPlanAdsList

func (_sps *SaasPlanAdsService) GetPublishedSaasPlanAdsList() (dto.SaasPlanAdsResponse1List, error) {

	saasplan_model_list, err := _sps.saasPlanAdsRepo.GetPublishedSaasPlanAdsList()

	if err != nil {
		return nil, err
	}

	saas_plan_res_list := dto.SaasPlanAdsResponse1List{}
	for _, saas_plan_model := range saasplan_model_list {
		saas_plan_res := dto.NewSaasPlanAdsResponse1(saas_plan_model)

		if err == nil {
			saas_plan_res_list = append(saas_plan_res_list, saas_plan_res)
		}
	}

	return saas_plan_res_list, nil

} //End of the method GetPublishedSaasPlanAdsList

func (_sps *SaasPlanAdsService) GetAllSaasPlanAdsList() (dto.SaasPlanAdsResponse1List, error) {

	saasplan_model_list, err := _sps.saasPlanAdsRepo.GetAllSaasPlanAdsList()

	if err != nil {
		return nil, err
	}

	saas_plan_res_list := dto.SaasPlanAdsResponse1List{}
	for _, saas_plan_model := range saasplan_model_list {
		saas_plan_res := dto.NewSaasPlanAdsResponse1(saas_plan_model)

		if err == nil {
			saas_plan_res_list = append(saas_plan_res_list, saas_plan_res)
		}
	}

	return saas_plan_res_list, nil

} //End of the method GetAllSaasPlanAdsList

func (_sps *SaasPlanAdsService) DeleteSaasPlanAds(_id string) error {

	return _sps.saasPlanAdsRepo.DeleteSaasPlanAds(_id)

} //End of the method DeleteSaasPlanAds
