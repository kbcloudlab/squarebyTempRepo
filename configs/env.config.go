package configs

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

var lock = &sync.Mutex{}

type envVariablesStruct struct {
	IsProduction          string
	DatabaseURL           string
	DatabaseName          string
	SQLDatabaseURL        string
	SQLDatabaseUsername   string
	SQLDbPassword         string
	SQLDatabaseName       string
	SQLDatabasePort       string
	ServerPortNumber      string
	RPCimServerPortNumber string
	// RPCAuthServerPortNumber     string
	RPCProductServicePortNumber string
	IdleTimeout                 int
	ReadTimeout                 int
	WriteTimeout                int
	DBRequestTimeoutSmall       time.Duration
	DBRequestTimeoutMedium      time.Duration
	DBRequestTimeoutLarge       time.Duration
	AllowedMethods              string
	AllowedOrigins              string
}

var EnvVariables *envVariablesStruct

func NewEnvVariables(_log *customLog.Logging) *envVariablesStruct {

	//load .env file from given path
	load_err := godotenv.Load(".env")

	if load_err != nil {
		_log.ErrorLogPanic("Error loading .env file: ", load_err)
		return nil
	}

	if EnvVariables == nil {
		lock.Lock()
		defer lock.Unlock()
		if EnvVariables == nil {
			EnvVariables = createEnvVariables(_log)
		}
	}

	return EnvVariables
}

func createEnvVariables(_log *customLog.Logging) *envVariablesStruct {

	is_production, found := os.LookupEnv("IS_PRODUCTION")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'IS_PRODUCTION' not found", "", 0))
	}

	db_url, found := os.LookupEnv("DB_URL")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'DB_URL' not found", "", 0))
	}

	db_name, found := os.LookupEnv("DB_NAME")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'DB_NAME' not found", "", 0))
	}

	sql_db_url, found := os.LookupEnv("SQL_DB_URL")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SQL_DB_URL' not found", "", 0))
	}

	sql_db_username, found := os.LookupEnv("SQL_DB_USERNAME")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SQL_DB_USERNAME' not found", "", 0))
	}

	sql_db_password, found := os.LookupEnv("SQL_DB_PASSWORD")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SQL_DB_PASSWORD' not found", "", 0))
	}

	sql_dbname, found := os.LookupEnv("SQL_DB_NAME")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SQL_DB_NAME' not found", "", 0))
	}

	sql_db_port, found := os.LookupEnv("SQL_DB_PORT")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SQL_DB_PORT' not found", "", 0))
	} else {
		_, err := strconv.Atoi(sql_db_port)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid SQL_DB_PORT value. Should be an integer value", "", 0))
		}
	}

	server_port, found := os.LookupEnv("SERVER_PORT")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'SERVER_PORT' not found", "", 0))
	} else {
		_, err := strconv.Atoi(server_port)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid SERVER_PORT value. Should be an integer value", "", 0))
		}
	}

	rpc_im_server_port, found := os.LookupEnv("RPC_IM_SERVER_PORT")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'RPC_IM_SERVER_PORT' not found", "", 0))
	} else {
		_, err := strconv.Atoi(rpc_im_server_port)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid RPC_IM_SERVER_PORT value. Should be an integer value", "", 0))
		}
	}

	// rpc_auther_server_port, found := os.LookupEnv("RPC_AUTH_SERVER_PORT")
	// if !found {
	// 	_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'RPC_AUTH_SERVER_PORT' not found", "", 0))
	// } else {
	// 	_, err := strconv.Atoi(rpc_auther_server_port)
	// 	if err != nil {
	// 		_log.ErrorLogPanic("", customError.NewSystemErr("Invalid RPC_AUTH_SERVER_PORT value. Should be an integer value", "", 0))
	// 	}
	// }

	rpc_productservice_port, found := os.LookupEnv("RPC_PRODUCTSERVICE_PORT")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'RPC_PRODUCTSERVICE_PORT' not found", "", 0))
	} else {
		_, err := strconv.Atoi(rpc_productservice_port)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid RPC_PRODUCTSERVICE_PORT value. Should be an integer value", "", 0))
		}
	}

	idle_timeout, found := os.LookupEnv("IDLE_TIMEOUT")
	idle_timeout_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'IDLE_TIMEOUT' not found", "", 0))
	} else {
		idle_timeout, err := strconv.Atoi(idle_timeout)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid IDLE_TIMEOUT value. Should be an integer value", "", 0))
		}
		idle_timeout_int = idle_timeout
	}

	read_timeout, found := os.LookupEnv("READ_TIMEOUT")
	read_timeout_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'READ_TIMEOUT' not found", "", 0))
	} else {
		read_timeout, err := strconv.Atoi(read_timeout)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid READ_TIMEOUT value. Should be an integer value", "", 0))
		}
		read_timeout_int = read_timeout
	}

	write_timeout, found := os.LookupEnv("WRITE_TIMEOUT")
	write_timeout_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'WRITE_TIMEOUT' not found", "", 0))
	} else {
		write_timeout, err := strconv.Atoi(write_timeout)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid WRITE_TIMEOUT value. Should be an integer value", "", 0))
		}
		write_timeout_int = write_timeout
	}

	db_timeout_small, found := os.LookupEnv("DB_TIMEOUT_SMALL")
	db_timeout_small_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'DB_TIMEOUT_SMALL' not found", "", 0))
	} else {
		db_timeout_small, err := strconv.Atoi(db_timeout_small)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid DB_TIMEOUT_SMALL value. Should be an integer value", "", 0))
		}
		db_timeout_small_int = db_timeout_small
	}

	db_timeout_medium, found := os.LookupEnv("DB_TIMEOUT_MEDIUM")
	db_timeout_medium_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'DB_TIMEOUT_MEDIUM' not found", "", 0))
	} else {
		db_timeout_medium, err := strconv.Atoi(db_timeout_medium)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid DB_TIMEOUT_MEDIUM value. Should be an integer value", "", 0))
		}
		db_timeout_medium_int = db_timeout_medium
	}

	db_timeout_large, found := os.LookupEnv("DB_TIMEOUT_LARGE")
	db_timeout_large_int := 0
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'DB_TIMEOUT_LARGE' not found", "", 0))
	} else {
		db_timeout_large, err := strconv.Atoi(db_timeout_large)
		if err != nil {
			_log.ErrorLogPanic("", customError.NewSystemErr("Invalid DB_TIMEOUT_LARGE value. Should be an integer value", "", 0))
		}
		db_timeout_large_int = db_timeout_large
	}

	allowed_methods, found := os.LookupEnv("ALLOWED_METHODS")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'ALLOWED_METHODS' not found", "", 0))
	}

	allowed_origins, found := os.LookupEnv("ALLOWED_ORIGINS")
	if !found {
		_log.ErrorLogPanic("", customError.NewSystemErr("Environment variable 'ALLOWED_ORIGINS' not found", "", 0))
	}

	return &envVariablesStruct{
		IsProduction:          is_production,
		DatabaseURL:           db_url,
		DatabaseName:          db_name,
		SQLDatabaseURL:        sql_db_url,
		SQLDatabaseName:       sql_dbname,
		SQLDatabasePort:       sql_db_port,
		SQLDatabaseUsername:   sql_db_username,
		SQLDbPassword:         sql_db_password,
		ServerPortNumber:      server_port,
		RPCimServerPortNumber: rpc_im_server_port,
		// RPCAuthServerPortNumber:     rpc_auther_server_port,
		RPCProductServicePortNumber: rpc_productservice_port,
		IdleTimeout:                 idle_timeout_int,
		ReadTimeout:                 read_timeout_int,
		WriteTimeout:                write_timeout_int,
		DBRequestTimeoutSmall:       time.Duration(db_timeout_small_int) * time.Second,
		DBRequestTimeoutMedium:      time.Duration(db_timeout_medium_int) * time.Second,
		DBRequestTimeoutLarge:       time.Duration(db_timeout_large_int) * time.Second,
		AllowedMethods:              allowed_methods,
		AllowedOrigins:              allowed_origins,
	}

} //
