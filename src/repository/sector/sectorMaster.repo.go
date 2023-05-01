package sector

import (
	"context"

	"gorm.io/gorm"
	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SectorMasterRepo struct {
	sqlDb   *gorm.DB
	logging *customLog.Logging
}

func NewSectorMasterRepo() *SectorMasterRepo {
	return &SectorMasterRepo{
		sqlDb:   configs.SQLDB(),
		logging: customLog.NewLogging(),
	}
}

func (_sfr *SectorMasterRepo) CreateSectorMaster(_sectorMaster *models.SectorMaster) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	_sfr.sqlDb.WithContext(ctx).Create(&_sectorMaster)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())
		return 0, system_err

	} else if _sectorMaster.Id <= 0 {

		system_err := customError.NewSystemErr("Saas feature creation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	}

	_sfr.logging.MessageLog("Saas Feature created successfully")

	return _sectorMaster.Id, nil

} //End of the method CreateSectorMaster

func (_sfr *SectorMasterRepo) UpdateSectorMaster(_id uint64, _sectorMaster *models.SectorMaster) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	exist_saasfeature, err := _sfr.GetSectorMasterById(_id)
	if err != nil {
		return err
	}

	_sectorMaster.Id = exist_saasfeature.Id
	_sfr.sqlDb.WithContext(ctx).Save(&_sectorMaster)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if _sectorMaster.Id == 0 {
		system_err := customError.NewSystemErr("SectorMaster updation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	return nil

} //End of the method UpdateSectorMaster

func (_sfr *SectorMasterRepo) GetSectorMasterByCode(_code string) (*models.SectorMaster, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SectorMaster{}).Fields()

	// sector_master := &models.SectorMaster{}
	var sector_master *models.SectorMaster

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Code+"=?", _code).Find(&sector_master) //

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return sector_master, system_err

	} else if sector_master.Id == 0 {

		notfound_err := customError.NewNotFoundErr("SectorMaster not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return sector_master, notfound_err

	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return sector_master, nil

} //End of the method GetSectorMasterByCode

func (_sfr *SectorMasterRepo) GetSectorMasterById(_id uint64) (*models.SectorMaster, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SectorMaster{}).Fields()

	// sector_master := &models.SectorMaster{}
	var sector_master *models.SectorMaster

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Id+"=?", _id).Find(&sector_master) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return sector_master, system_err
	} else if sector_master == nil {
		notfound_err := customError.NewNotFoundErr("SectorMaster not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return sector_master, notfound_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return sector_master, nil

} //End of the method GetSectorMasterById

func (_sfr *SectorMasterRepo) GetActiveSectorMasterList() ([]*models.SectorMaster, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SectorMaster{}).Fields()

	saasfeature_list := []*models.SectorMaster{}
	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Enabled + "=true").Find(&saasfeature_list) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saasfeature_list, system_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return saasfeature_list, nil

} //End of the method GetActiveSectorMasterList

func (_sfr *SectorMasterRepo) GetAllSectorMasterList() ([]*models.SectorMaster, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasfeature_list := []*models.SectorMaster{}
	_sfr.sqlDb.WithContext(ctx).Find(&saasfeature_list) //

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saasfeature_list, system_err

	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return saasfeature_list, nil

} //End of the method GetAllSectorMasterList
