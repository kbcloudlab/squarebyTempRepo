package main

import (
	"squareby.com/admin/cloudspacemanager/configs"
	"squareby.com/admin/cloudspacemanager/rest"

	// "squareby.com/admin/cloudspacemanager/rpc"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

func main() {
	logging := customLog.NewLogging()

	// initialization(logging)
	configs.NewEnvVariables(logging)

	configs.DBConnect()

	//RPC server must be run before the REST server, because if the REST server is run first, then the codes get blocked there itself.
	// rpc.StartServer(configs.EnvVariables.RPCimServerPortNumber, logging)
	rest.StartServer(configs.EnvVariables.ServerPortNumber, logging)

} //END


