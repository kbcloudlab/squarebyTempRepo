package sector

import (
	service "squareby.com/admin/cloudspacemanager/sidechain/service/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
)

type SectorMasterBridge struct {
	sectorMasterService *service.SectorMasterService
}

func NewSectorMasterBridge() *SectorMasterBridge {
	return &SectorMasterBridge{
		sectorMasterService: service.NewSectorMasterService(),
	}
}

func (_sfb *SectorMasterBridge) GetSectorMasterByCode(_code string) (*models.SectorMaster, error) {

	return _sfb.sectorMasterService.GetSectorMasterByCode(_code)

} //End of the method GetSectorMasterByCode
