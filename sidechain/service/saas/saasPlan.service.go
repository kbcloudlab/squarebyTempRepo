package saas

import (
	models "squareby.com/admin/cloudspacemanager/src/models/saas"

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

func (_sfs *SaasPlanService) GetSaasPlanById(_id string) (*models.SaasPlan, error) {

	return _sfs.saasPlanRepo.GetSaasPlanById(_id)

} //End of the method GetSaasPlanById
