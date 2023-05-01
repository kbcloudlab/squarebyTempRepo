package userrole

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"squareby.com/admin/cloudspacemanager/configs"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type UserRoleRepo struct {
	collection *mongo.Collection
	logging    *customLog.Logging
}

func NewUserRoleRepo() *UserRoleRepo {
	return &UserRoleRepo{
		collection: configs.DB().Collection(models.UserRole_CollectionName),
		logging:    customLog.NewLogging(),
	}
}

func (_urr *UserRoleRepo) CreateUserRole(_cache *models.UserRole) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	inserted, err := _urr.collection.InsertOne(ctx, _cache)

	if ctx.Err() == context.DeadlineExceeded {
		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	} else if err != nil {
		system_err := customError.NewSystemErr("DB insertion failed:"+err.Error()+" ::<<userRole.repo.go -> CreateUserRoleRepo()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return "", system_err
	}

	_urr.logging.MessageLog("User Role created successfully")

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil

} //End of the method CreateUserRole

func (_urr *UserRoleRepo) UpdateUserRole(_id string, _userRole *models.UserRole) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<userRole.repo.go -> UpdateUserRole()>>", "", 0)
		return system_err
	}

	userrole_fields := (&models.UserRole{}).Fields()

	filter := bson.M{userrole_fields.Col_Id: id}

	update := bson.M{"$set": _userRole}

	_, err = _urr.collection.UpdateOne(ctx, filter, update)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("DB update failed:"+err.Error()+" ::<<userRole.repository.go -> UpdateUserRole()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err

	}

	_urr.logging.MessageLog("User Role updated successfully")

	return nil
} //End of the function UpdateUserRole


func (_urr *UserRoleRepo) GetUserRoleById(_id string) (*models.UserRole, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<userRole.repository.go -> GetUserRoleById()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err
	}

	userrole_fields := (&models.UserRole{}).Fields()

	filter := bson.M{userrole_fields.Col_Id: id}

	userrole := &models.UserRole{}

	err = _urr.collection.FindOne(ctx, filter).Decode(userrole)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("Cache with id '"+_id+"' not found:: <<userRole.repository.go -> GetUserRoleById()>>", "UserRole not found", 0)
		_urr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return nil, notfound_err

	}

	_urr.logging.MessageLog("User Role by id: " + _id + " retrieved successfully")

	return userrole, err

} //End of the function GetCacheById

func (_urr *UserRoleRepo) GetActiveUserRoleList() ([]*models.UserRole, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	userrole_fields := (&models.UserRole{}).Fields()
	filter := bson.M{userrole_fields.Col_Enabled: true}

	userrole_list := []*models.UserRole{}

	cursor, err := _urr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get UserRole list failed:"+err.Error()+" ::<<userRole.repository.go -> GetActiveUserRoleList()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		userrole := &models.UserRole{}
		if err := cursor.Decode(userrole); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to UserRole model failed:"+err.Error()+" ::<<userRole.repository.go -> GetActiveUserRoleList()>>", "", 0)
			return nil, system_err

		}

		userrole_list = append(userrole_list, userrole)
	}

	defer cursor.Close(context.Background())

	_urr.logging.MessageLog("User Role list retrieved successfully")

	return userrole_list, err

} //End of the function GetActiveUserRoleList

func (_urr *UserRoleRepo) GetAllUserRoleList() ([]*models.UserRole, error) {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	// userrole_fields := (&models.UserRole{}).Fields()
	// filter := bson.M{userrole_fields.Col_Enabled: true}

	filter := bson.M{}

	userrole_list := []*models.UserRole{}

	cursor, err := _urr.collection.Find(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	} else if err != nil {

		system_err := customError.NewSystemErr("get UserRole list failed:"+err.Error()+" ::<<userRole.repository.go -> GetAllUserRoleList()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return nil, system_err

	}

	for cursor.Next(ctx) {
		userrole := &models.UserRole{}
		if err := cursor.Decode(userrole); err != nil {

			system_err := customError.NewSystemErr("decoding raw DB data to UserRole model failed:"+err.Error()+" ::<<userRole.repository.go -> GetAllUserRoleList()>>", "", 0)
			return nil, system_err

		}

		userrole_list = append(userrole_list, userrole)
	}

	defer cursor.Close(context.Background())

	_urr.logging.MessageLog("User Role list retrieved successfully")

	return userrole_list, err

} //End of the function GetAllUserRoleList

func (_urr *UserRoleRepo) DeleteUserRole(_id string) error {

	ctx, cancel := context.WithTimeout(context.Background(), configs.EnvVariables.DBRequestTimeoutSmall)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(_id)

	if err != nil {
		//Root error
		system_err := customError.NewSystemErr("converting _id of string to Hex format failed:"+err.Error()+" ::<<userRole.repository.go -> DeleteUserRole()>>", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	userrole_fields := (&models.UserRole{}).Fields()
	filter := bson.M{userrole_fields.Col_Id: id}

	_, err = _urr.collection.DeleteOne(ctx, filter)

	if ctx.Err() == context.DeadlineExceeded {

		system_err := customError.NewSystemErr("database timeout: ", "", 0)
		_urr.logging.ErrorLog(system_err.ErrorData.Message, ctx.Err())

		return system_err

	} else if err != nil {
		//Root error
		notfound_err := customError.NewNotFoundErr("Cache deletion failed :: <<cloudspauserRoleceUserRole.repository.go -> DeleteUserRole()>>", "UserRole not found", 0)
		_urr.logging.ErrorLog(notfound_err.ErrorData.Message, notfound_err)
		return notfound_err

	}

	_urr.logging.MessageLog("User Role deleted successfully")

	return nil

} //End of the function DeleteUserRole
