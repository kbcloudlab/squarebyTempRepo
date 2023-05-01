package saas

import (
	service "squareby.com/admin/cloudspacemanager/sidechain/service/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
)

type SaasFeatureBridge struct {
	saasFeatureService *service.SaasFeatureService
}

func NewSaasFeatureBridge() *SaasFeatureBridge {
	return &SaasFeatureBridge{
		saasFeatureService: service.NewSaasFeatureService(),
	}
}

func (_sfb *SaasFeatureBridge) GetSaasFeatureByCode(_featureCode string) (*models.SaasFeature, error) {

	return _sfb.saasFeatureService.GetSaasFeatureByCode(_featureCode)

} //End of the method GetSaasFeatureById
