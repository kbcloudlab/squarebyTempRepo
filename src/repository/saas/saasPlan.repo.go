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

type SaasPlanRepo struct {
	collection *mongo.Collection
	logging    *customLog.Logging
}

func NewSaasPlanRepo() *SaasPlanRepo {
	return &SaasPlanRepo{
		collection: configs.DB().Collection(models.SaasPlan_CollectionName),
		logging:    customLog.NewLogging(),
	}
}

func (_spr *SaasPlanRepo) CreateSaasPlan(_cache *models.SaasPlan) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	inserted, err := _spr.collection.InsertOne(ctx, _cache)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	} else if err != nil {
		system_err := customError.NewSystemErr("DB insertion failed:"+err.Error()+" ::<<saasPlan.repo.go -> CreateSaasPlan()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	}

	_spr.logging.MessageLog("SaasPlan created successfully")

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil

} //End of the method CreateSaasPlan

func (_spr *SaasPlanRepo) UpdateSaasPlan(_id string, _saasPlan *models.SaasPlan) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlan.repo.go -> UpdateSaasPlan()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	saasplan_fields := (&models.SaasPlan{}).Fields()

	filter := bson.M{saasplan_fields.Col_Id: id}

	update := bson.M{"$set": _saasPlan.NewDuplicateModel()}

	_, err = _spr.collection.UpdateOne(ctx, filter, update)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("DB update failed:"+err.Error()+" ::<<saasPlan.repository.go -> UpdateSaasPlan()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	_spr.logging.MessageLog("SaasPlan updated successfully")

	return nil
} //End of the function UpdateSaasPlan

func (_spr *SaasPlanRepo) GetSaasPlanById(_id string) (*models.SaasPlan, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlan.repository.go -> GetSaasPlanById()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err
	}

	saasplan_fields := (&models.SaasPlan{}).Fields()

	filter := bson.M{saasplan_fields.Col_Id: id}

	var saas_plan *models.SaasPlan

	err = _spr.collection.FindOne(ctx, filter).Decode(&saas_plan)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("SaasPlan with id '"+_id+"' not found:: <<cloudspaceSaasPlan.repository.go -> GetSaasPlanById()>>", "SaasPlan not found", 0)
		_spr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return nil, notfound_err

	}

	_spr.logging.MessageLog("SaasPlan by id: " + _id + " retrieved successfully")

	return saas_plan, err

} //End of the function GetSaasPlanById

func (_spr *SaasPlanRepo) GetActiveSaasPlanList() ([]*models.SaasPlan, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	saasplan_fields := (&models.SaasPlan{}).Fields()

	filter := bson.M{saasplan_fields.Col_Enabled: true}

	saasplan_list := []*models.SaasPlan{}

	cursor, err := _spr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get SaasPlan list failed:"+err.Error()+" ::<<saasPlan.repository.go -> GetActiveSaasPlanList()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		var saas_plan *models.SaasPlan
		if err := cursor.Decode(&saas_plan); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to SaasPlan model failed:"+err.Error()+" ::<<saasPlan.repository.go -> GetActiveSaasPlanList()>>", "", 0)
			return nil, system_err

		}

		if saas_plan != nil {
			saasplan_list = append(saasplan_list, saas_plan)
		}
	}

	defer cursor.Close(context.Background())

	_spr.logging.MessageLog("SaasPlan list retrieved successfully")

	return saasplan_list, err

} //End of the function GetActiveSaasPlanList

func (_spr *SaasPlanRepo) GetAllSaasPlanList() ([]*models.SaasPlan, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	filter := bson.M{}

	saasplan_list := []*models.SaasPlan{}

	cursor, err := _spr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get SaasPlan list failed:"+err.Error()+" ::<<saasPlan.repository.go -> GetAllSaasPlanList()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		var saas_plan *models.SaasPlan
		if err := cursor.Decode(&saas_plan); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to SaasPlan model failed:"+err.Error()+" ::<<saasPlan.repository.go -> GetAllSaasPlanList()>>", "", 0)
			return nil, system_err

		}

		if saas_plan != nil {
			saasplan_list = append(saasplan_list, saas_plan)
		}
	}

	defer cursor.Close(context.Background())

	_spr.logging.MessageLog("SaasPlan list retrieved successfully")

	return saasplan_list, err

} //End of the function GetAllSaasPlanList

func (_spr *SaasPlanRepo) DeleteSaasPlan(_id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<saasPlan.repository.go -> DeleteSaasPlan()>>", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	userrole_fields := (&models.SaasPlan{}).Fields()
	filter := bson.M{userrole_fields.Col_Id: id}

	_, err = _spr.collection.DeleteOne(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_spr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())

		return system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("SaasPlan deletion failed :: <<saasPlan.repository.go -> DeleteSaasPlan()>>", "SaasPlan not found", 0)
		_spr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return notfound_err

	}

	_spr.logging.MessageLog("SaasPlan deleted successfully")

	return nil

} //End of the function DeleteSaasPlan
