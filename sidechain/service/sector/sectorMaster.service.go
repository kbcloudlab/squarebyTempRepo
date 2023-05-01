package sector

import (
	models "squareby.com/admin/cloudspacemanager/src/models/sector"

	repo "squareby.com/admin/cloudspacemanager/src/repository/sector"
)

type SectorMasterService struct {
	sectorMasterRepo *repo.SectorMasterRepo
}

func NewSectorMasterService() *SectorMasterService {
	return &SectorMasterService{
		sectorMasterRepo: repo.NewSectorMasterRepo(),
	}
}

func (_sfs *SectorMasterService) GetSectorMasterByCode(_code string) (*models.SectorMaster, error) {

	return _sfs.sectorMasterRepo.GetSectorMasterByCode(_code)

} //End of the method GetSectorMasterByCode
