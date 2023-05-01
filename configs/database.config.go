package configs

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	saasModel "squareby.com/admin/cloudspacemanager/src/models/saas"
	sectorModel "squareby.com/admin/cloudspacemanager/src/models/sector"
	userroleModel "squareby.com/admin/cloudspacemanager/src/models/userrole"

	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

var db *mongo.Database
var db_sql *gorm.DB

func DBConnect() {

	logging := customLog.NewLogging()

	/*** MongoDB connection ***/
	clientOption := options.Client().ApplyURI(EnvVariables.DatabaseURL)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		logging.ErrorLogPanic("could not connect to database ", err)
	}

	db = client.Database(EnvVariables.DatabaseName)

	logging.MessageLog("MongoDB connection successfull")
	/*** MongoDB connection ends ***/

	/*** SQL DB connection ***/
	dsn := "host=" + EnvVariables.SQLDatabaseURL + " user=" + EnvVariables.SQLDatabaseUsername + " dbname=" + EnvVariables.SQLDatabaseName + " port=" + EnvVariables.SQLDatabasePort

	if EnvVariables.IsProduction == "1" {
		dsn = "host=" + EnvVariables.SQLDatabaseURL + " user=" + EnvVariables.SQLDatabaseUsername + " dbname=" + EnvVariables.SQLDatabaseName + " port=" + EnvVariables.SQLDatabasePort + " password=" + EnvVariables.SQLDbPassword
	}

	sql_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logging.ErrorLogPanic("could not connect to sql database ", err)
	}

	db_sql = sql_conn

	logging.MessageLog("SQL DB connection successfull")

	// db_sql.AutoMigrate(&masterentryModel.ProductContainer{})
	db_sql.AutoMigrate(&saasModel.SaasFeature{})
	db_sql.AutoMigrate(&userroleModel.PrivilegeAction{})
	db_sql.AutoMigrate(&sectorModel.SectorMaster{})
	db_sql.AutoMigrate(&sectorModel.SpaceCategory{})

	/*** SQL DB connection ends***/

}

func DB() *mongo.Database {
	return db
}

func SQLDB() *gorm.DB {
	return db_sql
}
