package saas

import (
	service "squareby.com/admin/cloudspacemanager/sidechain/service/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
)

type SaasPlanBridge struct {
	saasPlanService *service.SaasPlanService
}

func NewSaasPlanBridge() *SaasPlanBridge {
	return &SaasPlanBridge{
		saasPlanService: service.NewSaasPlanService(),
	}
}

func (_sfb *SaasPlanBridge) GetSaasPlanById(_id string) (*models.SaasPlan, error) {

	return _sfb.saasPlanService.GetSaasPlanById(_id)

} //End of the method GetSaasPlanById
