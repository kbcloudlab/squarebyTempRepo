package sector

import (
	"context"

	"gorm.io/gorm"
	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SpaceCategoryRepo struct {
	sqlDb   *gorm.DB
	logging *customLog.Logging
}

func NewSpaceCategoryRepo() *SpaceCategoryRepo {
	return &SpaceCategoryRepo{
		sqlDb:   configs.SQLDB(),
		logging: customLog.NewLogging(),
	}
}

func (_sfr *SpaceCategoryRepo) CreateSpaceCategory(_spaceCategory *models.SpaceCategory) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	_sfr.sqlDb.WithContext(ctx).Create(&_spaceCategory)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())
		return 0, system_err

	} else if _spaceCategory.Id <= 0 {

		system_err := customError.NewSystemErr("Saas feature creation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	}

	_sfr.logging.MessageLog("Saas Feature created successfully")

	return _spaceCategory.Id, nil

} //End of the method CreateSpaceCategory

func (_sfr *SpaceCategoryRepo) UpdateSpaceCategory(_id uint64, _spaceCategory *models.SpaceCategory) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	exist_saasfeature, err := _sfr.GetSpaceCategoryById(_id)
	if err != nil {
		return err
	}

	_spaceCategory.Id = exist_saasfeature.Id
	_sfr.sqlDb.WithContext(ctx).Save(&_spaceCategory)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if _spaceCategory.Id == 0 {
		system_err := customError.NewSystemErr("SpaceCategory updation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	return nil

} //End of the method UpdateSpaceCategory

func (_sfr *SpaceCategoryRepo) GetSpaceCategoryByCode(_code string) (*models.SpaceCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SpaceCategory{}).Fields()

	// space_category := &models.SpaceCategory{}
	var space_category *models.SpaceCategory

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Code+"=?", _code).Find(&space_category) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return space_category, system_err
	} else if space_category.Id == 0 {
		notfound_err := customError.NewNotFoundErr("SpaceCategory not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return space_category, notfound_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return space_category, nil

} //End of the method GetSpaceCategoryByCode

func (_sfr *SpaceCategoryRepo) GetSpaceCategoryById(_id uint64) (*models.SpaceCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SpaceCategory{}).Fields()

	// space_category := &models.SpaceCategory{}
	var space_category *models.SpaceCategory

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Id+"=?", _id).Find(&space_category) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return space_category, system_err
	} else if space_category.Id == 0 {
		notfound_err := customError.NewNotFoundErr("SpaceCategory not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return space_category, notfound_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return space_category, nil

} //End of the method GetSpaceCategoryById

func (_sfr *SpaceCategoryRepo) GetActiveSpaceCategoryList() ([]*models.SpaceCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SpaceCategory{}).Fields()

	saasfeature_list := []*models.SpaceCategory{}
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

} //End of the method GetActiveSpaceCategoryList

func (_sfr *SpaceCategoryRepo) GetAllSpaceCategoryList() ([]*models.SpaceCategory, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasfeature_list := []*models.SpaceCategory{}
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

} //End of the method GetAllSpaceCategoryList
