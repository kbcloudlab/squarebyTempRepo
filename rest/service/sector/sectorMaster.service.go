package sector

import (
	"errors"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
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

func (_sfs *SectorMasterService) CreateSectorMaster() error {

	for _, sector_master_data := range models.SectorMasterDataList {

		_, err := _sfs.sectorMasterRepo.GetSectorMasterByCode(sector_master_data.Code)

		var notFoundErr *customError.NotFoundErr = &customError.NotFoundErr{}

		if errors.As(err, &notFoundErr) {
			// sector_master_model := &models.SectorMaster{
			// 	Code:     sector_master_data.Code,
			// 	Name:     sector_master_data.Name,
			// 	ImageURL: sector_master_data.ImageURL,
			// 	Enabled:  true,
			// }
			_, _ = _sfs.sectorMasterRepo.CreateSectorMaster(sector_master_data)

		} else if err != nil {
			return err
		}

	}

	return nil

} //End of the method CreateSectorMaster

func (_sfs *SectorMasterService) UpdateSectorMaster(_id uint64, _sectorMasterRequest *dto.SectorMasterRequest) error {

	sector_master := _sectorMasterRequest.NewUserRoleFromRequest(false)

	exist_sectormaster, err := _sfs.sectorMasterRepo.GetSectorMasterById(_id)

	if err != nil {
		return err
	}

	exist_sectormaster.Name = sector_master.Name
	exist_sectormaster.ImageURL = sector_master.ImageURL

	return _sfs.sectorMasterRepo.UpdateSectorMaster(_id, exist_sectormaster)

} //End of the method UpdateSectorMaster

func (_sfs *SectorMasterService) EnableDisableSectorMaster(_id uint64, _sectorMasterRequest *dto.SectorMasterEnableRequest) error {

	sector_master, err := _sfs.sectorMasterRepo.GetSectorMasterById(_id)
	if err != nil {
		return err
	}

	sector_master.Enabled = _sectorMasterRequest.Enabled

	return _sfs.sectorMasterRepo.UpdateSectorMaster(_id, sector_master)

} //End of the method EnableDisableSectorMaster

func (_sfs *SectorMasterService) GetSectorMasterByCode(_code string) (*dto.SectorMasterResponse, error) {

	sector_master, err := _sfs.sectorMasterRepo.GetSectorMasterByCode(_code)

	if err != nil {
		return nil, err
	}

	return dto.NewSectorMasterResponse(sector_master), nil

} //End of the method GetSectorMasterByCode

func (_sfs *SectorMasterService) GetSectorMasterById(_id uint64) (*dto.SectorMasterResponse, error) {

	sector_master, err := _sfs.sectorMasterRepo.GetSectorMasterById(_id)

	if err != nil {
		return nil, err
	}

	return dto.NewSectorMasterResponse(sector_master), nil

} //End of the method GetSectorMasterById

func (_sfs *SectorMasterService) GetActiveSectorMasterList() (dto.SectorMasterResponseList, error) {

	sector_master_list, err := _sfs.sectorMasterRepo.GetActiveSectorMasterList()

	if err != nil {
		return nil, err
	}

	sectormaster_res_list := dto.SectorMasterResponseList{}
	for _, sectormaster := range sector_master_list {
		sectormaster_res_list = append(sectormaster_res_list, dto.NewSectorMasterResponse(sectormaster))
	}

	return sectormaster_res_list, nil

} //End of the method GetActiveSectorMasterList

func (_sfs *SectorMasterService) GetAllSectorMasterList() (dto.SectorMasterResponseList, error) {

	sector_master_list, err := _sfs.sectorMasterRepo.GetAllSectorMasterList()

	if err != nil {
		return nil, err
	}

	sectormaster_res_list := dto.SectorMasterResponseList{}
	for _, sectormaster := range sector_master_list {
		sectormaster_res_list = append(sectormaster_res_list, dto.NewSectorMasterResponse(sectormaster))
	}

	return sectormaster_res_list, nil

} //End of the method GetAllSectorMasterList
