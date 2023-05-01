package saas

import (
	models "squareby.com/admin/cloudspacemanager/src/models/saas"

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

func (_sfs *SaasFeatureService) GetSaasFeatureByCode(_featureCode string) (*models.SaasFeature, error) {

	return _sfs.saasFeatureRepo.GetSaasFeatureByCode(_featureCode)

} //End of the method GetSaasFeatureById
