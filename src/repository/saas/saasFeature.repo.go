package saas

import (
	"context"

	"gorm.io/gorm"
	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasFeatureRepo struct {
	sqlDb       *gorm.DB
	logging     *customLog.Logging
	customError *customError.ErrorCustom
}

func NewSaasFeatureRepo() *SaasFeatureRepo {
	return &SaasFeatureRepo{
		sqlDb:       configs.SQLDB(),
		logging:     customLog.NewLogging(),
		customError: customError.NewErrorCustom(),
	}
}

func (_sfr *SaasFeatureRepo) CreateSaasFeature(_saasFeature *models.SaasFeature) (uint64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	_sfr.sqlDb.WithContext(ctx).Create(&_saasFeature)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())
		return 0, system_err

	} else if _saasFeature.Id <= 0 {

		system_err := customError.NewSystemErr("Saas feature creation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return 0, system_err

	}

	_sfr.logging.MessageLog("Saas Feature created successfully")

	return _saasFeature.Id, nil

} //End of the method CreateSaasFeature

func (_sfr *SaasFeatureRepo) UpdateSaasFeature(_id uint64, _saasFeature *models.SaasFeature) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	exist_saasfeature, err := _sfr.GetSaasFeatureById(_id)
	if err != nil {
		return err
	}

	_saasFeature.Id = exist_saasfeature.Id
	_sfr.sqlDb.WithContext(ctx).Save(&_saasFeature)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if _saasFeature.Id == 0 {
		system_err := customError.NewSystemErr("SaasFeature updation failed", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	return nil

} //End of the method UpdateSaasFeature

func (_sfr *SaasFeatureRepo) GetSaasFeatureByCode(_code string) (*models.SaasFeature, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SaasFeature{}).Fields()

	// saas_feature := &models.SaasFeature{}
	var saas_feature *models.SaasFeature

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Code+"=?", _code).Find(&saas_feature) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saas_feature, system_err
	} else if saas_feature.Id == 0 {
		notfound_err := customError.NewNotFoundErr("SaasFeature not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return saas_feature, notfound_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return saas_feature, nil

} //End of the method GetSaasFeatureByCode

func (_sfr *SaasFeatureRepo) GetSaasFeatureById(_id uint64) (*models.SaasFeature, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SaasFeature{}).Fields()

	// saas_feature := &models.SaasFeature{}
	var saas_feature *models.SaasFeature

	_sfr.sqlDb.WithContext(ctx).Where(saasfeature_fields.Col_Id+"=?", _id).Find(&saas_feature) //

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return saas_feature, system_err
	} else if saas_feature == nil {
		notfound_err := customError.NewNotFoundErr("SaasFeature not found", "", 0)
		_sfr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return saas_feature, notfound_err
	} else if ctx.Err() != nil {

		system_err := customError.NewSystemErr("Something went wrong", "", 0)
		_sfr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	return saas_feature, nil

} //End of the method GetSaasFeatureById

func (_sfr *SaasFeatureRepo) GetActiveSaasFeatureList() ([]*models.SaasFeature, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	//Get the fields
	saasfeature_fields := (&models.SaasFeature{}).Fields()

	saasfeature_list := []*models.SaasFeature{}
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

} //End of the method GetActiveSaasFeatureList

func (_sfr *SaasFeatureRepo) GetAllSaasFeatureList() ([]*models.SaasFeature, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasfeature_list := []*models.SaasFeature{}
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

} //End of the method GetAllSaasFeatureList
