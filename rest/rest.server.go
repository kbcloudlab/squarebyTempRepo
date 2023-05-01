package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"squareby.com/admin/cloudspacemanager/configs"
	routes "squareby.com/admin/cloudspacemanager/rest/routes"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

func StartServer(_port string, _log *customLog.Logging) {

	//Creating new serve mux and register the controllers
	sm := mux.NewRouter()

	sectormaster_routes := routes.NewSectorMasterRoutes(sm)
	sectormaster_routes.CreateSectorMaster() //Created automatically is data not already present in the database
	sectormaster_routes.UpdateSectorMaster()
	sectormaster_routes.EnableDisableSectorMaster()
	sectormaster_routes.GetSectorMasterById()
	sectormaster_routes.GetSectorMasterByCode()
	sectormaster_routes.GetActiveSectorMasterList()
	sectormaster_routes.GetAllSectorMasterList()

	spacecategory_routes := routes.NewSpaceCategoryRoutes(sm)
	spacecategory_routes.CreateSpaceCategory() //Created automatically is data not already present in the database
	spacecategory_routes.UpdateSpaceCategory()
	spacecategory_routes.AddRemoveProductContainer()
	spacecategory_routes.EnableDisableSpaceCategory()
	spacecategory_routes.GetSpaceCategoryById()
	spacecategory_routes.GetSpaceCategoryByCode()
	spacecategory_routes.GetActiveSpaceCategoryList()
	spacecategory_routes.GetAllSpaceCategoryList()

	saasfeature_routes := routes.NewSaasFeatureRoutes(sm)
	saasfeature_routes.CreateSaasFeature() //Created automatically is data not already present in the database
	saasfeature_routes.UpdateSaasFeature()
	saasfeature_routes.EnableDisableSaasFeature()
	saasfeature_routes.GetSaasFeatureById()
	saasfeature_routes.GetSaasFeatureByCode()
	saasfeature_routes.GetActiveSaasFeatureList()
	saasfeature_routes.GetAllSaasFeatureList()

	privilegeaction_routes := routes.NewPrivilegeActionRoutes(sm)
	privilegeaction_routes.CreatePrivilegeAction() //Created automatically is data not already present in the database
	privilegeaction_routes.UpdatePrivilegeAction()
	privilegeaction_routes.EnableDisablePrivilegeAction()
	privilegeaction_routes.GetPrivilegeActionById()
	privilegeaction_routes.GetPrivilegeActionByCode()
	privilegeaction_routes.GetPrivilegeActionByCodeNumber()
	privilegeaction_routes.GetActivePrivilegeActionList()
	privilegeaction_routes.GetAllPrivilegeActionList()

	userrole_routes := routes.NewUserRoleRoutes(sm)
	userrole_routes.CreateUserRole()
	userrole_routes.UpdateUserRole()
	userrole_routes.GetUserRoleById()
	userrole_routes.GetActiveUserRoleList()
	userrole_routes.GetAllUserRoleList()
	userrole_routes.DeleteUserRole()
	userrole_routes.EnableDisableUserRole()

	saasplan_routes := routes.NewSaasPlanRoutes(sm)
	saasplan_routes.CreateSaasPlan()
	saasplan_routes.UpdateSaasPlan()
	saasplan_routes.EnableDisableSaasPlan()
	saasplan_routes.GetSaasPlanById()
	saasplan_routes.GetActiveSaasPlanList()
	saasplan_routes.GetAllSaasPlanList()
	saasplan_routes.DeleteSaasPlan()

	saasplan_ads_routes := routes.NewSaasPlanAdsRoutes(sm)
	saasplan_ads_routes.CreateSaasPlanAds()
	saasplan_ads_routes.UpdateSaasPlanAds()
	saasplan_ads_routes.GetSaasPlanAdsById()
	saasplan_ads_routes.GetActiveSaasPlanAdsList()
	saasplan_ads_routes.GetPublishedSaasPlanAdsList()
	saasplan_ads_routes.GetAllSaasPlanAdsList()
	saasplan_ads_routes.DeleteSaasPlanAds()
	saasplan_ads_routes.EnableDisableSaasPlanAds()
	saasplan_ads_routes.PublishSaasPlanAds()

	serverInit(_port, sm, _log)
}

func serverInit(_port string, _mux *mux.Router, _log *customLog.Logging) {

	//CORS
	c := configureCORS()

	server := &http.Server{
		Addr:         ":" + configs.EnvVariables.ServerPortNumber,
		Handler:      c.Handler(_mux),
		IdleTimeout:  time.Duration(configs.EnvVariables.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(configs.EnvVariables.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(configs.EnvVariables.WriteTimeout) * time.Second,
	}

	_log.MessageLog("Starting server...")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			_log.ErrorLogPanic("", err)
		}

	}()
	_log.MessageLog("Started server on port: " + configs.EnvVariables.ServerPortNumber)

	//Graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	_log.MessageLog("Recieved terminate, graceful shutdown" + sig.String())

	tc, err1 := context.WithTimeout(context.Background(), 30*time.Second)

	if err1 != nil {
		_log.MessageLog("Error gracefully shutting down")
	}
	server.Shutdown(tc)

} //End of the function ServerInit

func configureCORS() *cors.Cors {

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	return c

} //End of the function configureCORS
