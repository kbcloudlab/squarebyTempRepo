package userrole

import (
	"context"

	"gorm.io/gorm"
	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type PrivilegeActionRepo struct {
	sqlDb       *gorm.DB
	logging     *customLog.Logging
	customError *customError.ErrorCustom
}

func NewPrivilegeActionRepo() *PrivilegeActionRepo {
	return &PrivilegeActionRepo{
		sqlDb:       configs.SQLDB(),
		logging:     customLog.NewLogging(),
		customError: customError.NewErrorCustom(),
	}
}

func (_sfr *PrivilegeActionRepo) CreatePrivilegeAction(_privilegeAction *models.PrivilegeAction) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	_sfr.sqlDb.WithContext(ctx).Create(&_privilegeAction)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())
		return 0, system_err

	} else if _privilegeAction.Id <= 0 {

		system_err := customError.NewSystemErr("Saas feature creation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	}

	_sfr.logging.MessageLog("Saas Feature created successfully")

	return _privilegeAction.Id, nil

} //End of the method CreatePrivilegeAction

func (_sfr *PrivilegeActionRepo) UpdatePrivilegeAction(_id uint64, _privilegeAction *models.PrivilegeAction) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	exist_saasfeature, err := _sfr.GetPrivilegeActionById(_id)
	if err != nil {
		return err
	}

	_privilegeAction.Id = exist_saasfeature.Id
	_sfr.sqlDb.WithContext(ctx).Save(&_privilegeAction)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if _privilegeAction.Id <= 0 {
		system_err := customError.NewSystemErr("PrivilegeAction updation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil

} //End of the method UpdatePrivilegeAction

func (_sfr *PrivilegeActionRepo) GetPrivilegeActionByCode(_code string) (*models.PrivilegeAction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.PrivilegeAction{}).Fields()

	// saas_feature := &models.PrivilegeAction{}
	var saas_feature *models.PrivilegeAction

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Code+"=?", _code).Find(&saas_feature) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saas_feature, system_err
	} else if saas_feature.Id == 0 {
		notfound_err := customError.NewNotFoundErr("PrivilegeAction not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return saas_feature, notfound_err
	}

	return saas_feature, nil

} //End of the method GetPrivilegeActionByCode

func (_sfr *PrivilegeActionRepo) GetPrivilegeActionByCodeNumber(_codeNumber uint16) (*models.PrivilegeAction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.PrivilegeAction{}).Fields()

	// saas_feature := &models.PrivilegeAction{}
	var saas_feature *models.PrivilegeAction

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_CodeNumber+"=?", _codeNumber).Find(&saas_feature) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saas_feature, system_err
	} else if saas_feature.Id == 0 {
		notfound_err := customError.NewNotFoundErr("PrivilegeAction not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return saas_feature, notfound_err
	}

	return saas_feature, nil

} //End of the method GetPrivilegeActionByCodeNumber

func (_sfr *PrivilegeActionRepo) GetPrivilegeActionById(_id uint64) (*models.PrivilegeAction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.PrivilegeAction{}).Fields()

	// saas_feature := &models.PrivilegeAction{}
	var saas_feature *models.PrivilegeAction

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Id+"=?", _id).Find(&saas_feature) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saas_feature, system_err
	} else if saas_feature == nil {
		notfound_err := customError.NewNotFoundErr("PrivilegeAction not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return saas_feature, notfound_err
	}

	return saas_feature, nil

} //End of the method GetPrivilegeActionById

func (_sfr *PrivilegeActionRepo) GetActivePrivilegeActionList() ([]*models.PrivilegeAction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.PrivilegeAction{}).Fields()

	saasfeature_list := []*models.PrivilegeAction{}
	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Enabled + "=true").Find(&saasfeature_list) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saasfeature_list, system_err
	}

	return saasfeature_list, nil

} //End of the method GetActivePrivilegeActionList

func (_sfr *PrivilegeActionRepo) GetAllPrivilegeActionList() ([]*models.PrivilegeAction, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasfeature_list := []*models.PrivilegeAction{}
	_sfr.sqlDb.WithContext(ctx).Find(&saasfeature_list) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saasfeature_list, system_err
	}

	return saasfeature_list, nil

} //End of the method GetAllPrivilegeActionList
