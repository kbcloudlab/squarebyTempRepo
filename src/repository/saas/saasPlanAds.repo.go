package saas

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanAdsRepo struct {
	collection *mongo.Collection
	logging    *customLog.Logging
}

func NewSaasPlanAdsRepo() *SaasPlanAdsRepo {
	return &SaasPlanAdsRepo{
		collection: configs.DB().Collection(models.SaasPlanAds_CollectionName),
		logging:    customLog.NewLogging(),
	}
}

func (_spr *SaasPlanAdsRepo) CreateSaasPlanAds(_cache *models.SaasPlanAds) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	inserted, err := _spr.collection.InsertOne(ctx, _cache)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	} else if err != nil {
		system_err := customError.NewSystemErr("DB insertion failed:"+err.Error()+" ::<<saasPlanAds.repo.go -> CreateSaasPlanAds()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	}

	_spr.logging.MessageLog("SaasPlanAds created successfully")

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil

} //End of the method CreateSaasPlanAds

func (_spr *SaasPlanAdsRepo) UpdateSaasPlanAds(_id string, _saasPlanAds *models.SaasPlanAds) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlanAds.repo.go -> UpdateSaasPlanAds()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	saasplan_fields := (&models.SaasPlanAds{}).Fields()

	filter := bson.M{saasplan_fields.Col_Id: id}

	update := bson.M{"$set": _saasPlanAds.NewDuplicateModel()}

	_, err = _spr.collection.UpdateOne(ctx, filter, update)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("DB update failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> UpdateSaasPlanAds()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	_spr.logging.MessageLog("SaasPlanAds updated successfully")

	return nil
} //End of the function UpdateSaasPlanAds

func (_spr *SaasPlanAdsRepo) EnableDisableSaasPlanAds(_id string, _flag bool) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlanAds.repo.go -> EnableDisableSaasPlanAds()>>", "", 0)
		return system_err
	}

	saas_plan, err := _spr.GetSaasPlanAdsById(_id)

	if err != nil {
		return err
	}

	saasplan_fields := (&models.SaasPlanAds{}).Fields()

	filter := bson.M{saasplan_fields.Col_Id: id}

	saas_plan.Enabled = _flag
	update := bson.M{"$set": saas_plan.NewDuplicateModel()}

	_, err = _spr.collection.UpdateOne(ctx, filter, update)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("DB update failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> EnableDisableSaasPlanAds()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	_spr.logging.MessageLog("SaasPlanAds updated successfully")

	return nil
} //End of the function EnableDisableSaasPlanAds

func (_spr *SaasPlanAdsRepo) GetSaasPlanAdsById(_id string) (*models.SaasPlanAds, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetSaasPlanAdsById()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err
	}

	saasplan_fields := (&models.SaasPlanAds{}).Fields()

	filter := bson.M{saasplan_fields.Col_Id: id}

	var saas_plan *models.SaasPlanAds

	err = _spr.collection.FindOne(ctx, filter).Decode(&saas_plan)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("SaasPlanAds with id '"+_id+"' not found:: <<cloudspaceSaasPlanAds.repository.go -> GetSaasPlanAdsById()>>", "SaasPlanAds not found", 0)
		_spr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return nil, notfound_err

	}

	_spr.logging.MessageLog("SaasPlanAds by id: " + _id + " retrieved successfully")

	return saas_plan, err

} //End of the function GetSaasPlanAdsById

func (_spr *SaasPlanAdsRepo) GetActiveSaasPlanAdsList() ([]*models.SaasPlanAds, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasplan_fields := (&models.SaasPlanAds{}).Fields()

	filter := bson.M{saasplan_fields.Col_Enabled: true}

	saasplan_list := []*models.SaasPlanAds{}

	cursor, err := _spr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get SaasPlanAds list failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetActiveSaasPlanAdsList()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		var saas_plan *models.SaasPlanAds
		if err := cursor.Decode(&saas_plan); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to SaasPlanAds model failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetActiveSaasPlanAdsList()>>", "", 0)
			return nil, system_err

		}

		if saas_plan != nil {
			saasplan_list = append(saasplan_list, saas_plan)
		}
	}

	defer cursor.Close(context.Background())

	_spr.logging.MessageLog("SaasPlanAds list retrieved successfully")

	return saasplan_list, err

} //End of the function GetActiveSaasPlanAdsList

func (_spr *SaasPlanAdsRepo) GetPublishedSaasPlanAdsList() ([]*models.SaasPlanAds, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasplan_fields := (&models.SaasPlanAds{}).Fields()

	filter := bson.M{saasplan_fields.Col_Published: true, saasplan_fields.Col_Enabled: true}

	saasplan_list := []*models.SaasPlanAds{}

	cursor, err := _spr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get SaasPlanAds list failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetActiveSaasPlanAdsList()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		var saas_plan *models.SaasPlanAds
		if err := cursor.Decode(&saas_plan); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to SaasPlanAds model failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetActiveSaasPlanAdsList()>>", "", 0)
			return nil, system_err

		}

		if saas_plan != nil {
			saasplan_list = append(saasplan_list, saas_plan)
		}
	}

	defer cursor.Close(context.Background())

	_spr.logging.MessageLog("SaasPlanAds list retrieved successfully")

	return saasplan_list, err

} //End of the function GetPublishedSaasPlanAdsList

func (_spr *SaasPlanAdsRepo) GetAllSaasPlanAdsList() ([]*models.SaasPlanAds, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	filter := bson.M{}

	saasplan_list := []*models.SaasPlanAds{}

	cursor, err := _spr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get SaasPlanAds list failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetAllSaasPlanAdsList()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		var saas_plan *models.SaasPlanAds
		if err := cursor.Decode(&saas_plan); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to SaasPlanAds model failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> GetAllSaasPlanAdsList()>>", "", 0)
			return nil, system_err

		}

		if saas_plan != nil {
			saasplan_list = append(saasplan_list, saas_plan)
		}
	}

	defer cursor.Close(context.Background())

	_spr.logging.MessageLog("SaasPlanAds list retrieved successfully")

	return saasplan_list, err

} //End of the function GetAllSaasPlanAdsList

func (_spr *SaasPlanAdsRepo) DeleteSaasPlanAds(_id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlanAds.repository.go -> DeleteSaasPlanAds()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	userrole_fields := (&models.SaasPlanAds{}).Fields()
	filter := bson.M{userrole_fields.Col_Id: id}

	_, err = _spr.collection.DeleteOne(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())

		return system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("SaasPlanAds deletion failed :: <<saasPlanAds.repository.go -> DeleteSaasPlanAds()>>", "SaasPlanAds not found", 0)
		_spr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return notfound_err

	}

	_spr.logging.MessageLog("SaasPlanAds deleted successfully")

	return nil

} //End of the function DeleteSaasPlanAds
