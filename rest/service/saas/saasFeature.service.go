package saas

import (
	"errors"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	repo "squareby.com/admin/cloudspacemanager/src/repository/saas"
)

type SaasFeatureService struct {
	saasFeatureRepo *repo.SaasFeatureRepo
}

func NewSaasFeatureService() *SaasFeatureService {
	return &SaasFeatureService{
		saasFeatureRepo: repo.NewSaasFeatureRepo(),
	}
}

func (_sfs *SaasFeatureService) CreateSaasFeature() error {

	for _, saas_feature_data := range models.SaasFeatureDataList {

		_, err := _sfs.saasFeatureRepo.GetSaasFeatureByCode(saas_feature_data.Code)

		var notFoundErr *customError.NotFoundErr = &customError.NotFoundErr{}

		if errors.As(err, &notFoundErr) {
			saas_feature_model := &models.SaasFeature{
				Code:        saas_feature_data.Code,
				Title:       saas_feature_data.Title,
				FeatureType: saas_feature_data.FeatureType,
				Enabled:     true,
			}
			_, _ = _sfs.saasFeatureRepo.CreateSaasFeature(saas_feature_model)

		} else if err != nil {
			return err
		}

	}

	return nil

} //End of the method CreateSaasFeature

func (_sfs *SaasFeatureService) UpdateSaasFeature(_id uint64, _saasFeatureRequest *dto.SaasFeatureRequest) error {

	saas_feature := _saasFeatureRequest.NewUserRoleFromRequest(false)

	exist_saasfeature, err := _sfs.saasFeatureRepo.GetSaasFeatureById(_id)

	if err != nil {
		return err
	}

	saas_feature.Id = exist_saasfeature.Id
	saas_feature.FeatureType = exist_saasfeature.FeatureType
	saas_feature.Code = exist_saasfeature.Code
	saas_feature.Enabled = exist_saasfeature.Enabled

	return _sfs.saasFeatureRepo.UpdateSaasFeature(_id, saas_feature)

} //End of the method UpdateSaasFeature

func (_sfs *SaasFeatureService) EnableDisableSaasFeature(_id uint64, _saasFeatureRequest *dto.SaasFeatureEnableRequest) error {

	saas_feature, err := _sfs.saasFeatureRepo.GetSaasFeatureById(_id)
	if err != nil {
		return err
	}

	saas_feature.Enabled = _saasFeatureRequest.Enabled

	return _sfs.saasFeatureRepo.UpdateSaasFeature(_id, saas_feature)

} //End of the method EnableDisableSaasFeature

func (_sfs *SaasFeatureService) GetSaasFeatureByCode(_code string) (*dto.SaasFeatureResponse, error) {

	saas_feature, err := _sfs.saasFeatureRepo.GetSaasFeatureByCode(_code)

	if err != nil {
		return nil, err
	}

	return dto.NewSaasFeatureResponse(saas_feature), nil

} //End of the method GetSaasFeatureByCode

func (_sfs *SaasFeatureService) GetSaasFeatureById(_id uint64) (*dto.SaasFeatureResponse, error) {

	saas_feature, err := _sfs.saasFeatureRepo.GetSaasFeatureById(_id)

	if err != nil {
		return nil, err
	}

	return dto.NewSaasFeatureResponse(saas_feature), nil

} //End of the method GetSaasFeatureById

func (_sfs *SaasFeatureService) GetActiveSaasFeatureList() (dto.SaasFeatureResponseList, error) {

	saas_feature_list, err := _sfs.saasFeatureRepo.GetActiveSaasFeatureList()

	if err != nil {
		return nil, err
	}

	saasfeature_res_list := dto.SaasFeatureResponseList{}
	for _, saasfeature := range saas_feature_list {
		saasfeature_res_list = append(saasfeature_res_list, dto.NewSaasFeatureResponse(saasfeature))
	}

	return saasfeature_res_list, nil

} //End of the method GetActiveSaasFeatureList

func (_sfs *SaasFeatureService) GetAllSaasFeatureList() (dto.SaasFeatureResponseList, error) {

	saas_feature_list, err := _sfs.saasFeatureRepo.GetAllSaasFeatureList()

	if err != nil {
		return nil, err
	}

	saasfeature_res_list := dto.SaasFeatureResponseList{}
	for _, saasfeature := range saas_feature_list {
		saasfeature_res_list = append(saasfeature_res_list, dto.NewSaasFeatureResponse(saasfeature))
	}

	return saasfeature_res_list, nil

} //End of the method GetAllSaasFeatureList
