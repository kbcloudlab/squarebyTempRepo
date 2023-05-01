package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/saas"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/saas"
)

type SaasPlanRoutes struct {
	ServeMuxRouter     *(mux.Router)
	saasPlanController *controller.SaasPlanController
	saasPlanMiddleware *middleware.SaasPlanMiddleware
}

func NewSaasPlanRoutes(_sm *(mux.Router)) *SaasPlanRoutes {
	return &SaasPlanRoutes{
		ServeMuxRouter:     _sm,
		saasPlanController: controller.NewSaasPlanController(),
		saasPlanMiddleware: middleware.NewSaasPlanMiddleware(),
	}
}

func (_pcr *SaasPlanRoutes) CreateSaasPlan() {
	create_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodPost).Subrouter()
	create_saasPlan_router.HandleFunc("/admin/create-saas-plan", _pcr.saasPlanController.CreateSaasPlan)

	//Validation with Middleware
	create_saasPlan_router.Use(_pcr.saasPlanMiddleware.CreateUpdateSaasPlan)
}

func (_pcr *SaasPlanRoutes) UpdateSaasPlan() {
	update_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasPlan_router.HandleFunc("/admin/update-saas-plan/{id:[0-9a-fA-F]+}", _pcr.saasPlanController.UpdateSaasPlan)

	//Validation with Middleware
	update_saasPlan_router.Use(_pcr.saasPlanMiddleware.CreateUpdateSaasPlan)
}

func (_pcr *SaasPlanRoutes) EnableDisableSaasPlan() {
	update_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasPlan_router.HandleFunc("/admin/enabledisable-saas-plan/{id:[0-9a-fA-F]+}", _pcr.saasPlanController.EnableDisableSaasPlan)

	//Validation with Middleware
	update_saasPlan_router.Use(_pcr.saasPlanMiddleware.EnableDisableSaasPlan)
}

func (_pcr *SaasPlanRoutes) GetSaasPlanById() {
	get_saasPlan_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlan_byid_router.HandleFunc("/saas-plan/{id:[0-9a-fA-F]+}", _pcr.saasPlanController.GetSaasPlanById)
}

func (_pcr *SaasPlanRoutes) GetActiveSaasPlanList() {
	get_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlan_router.HandleFunc("/active-saas-plan-list", _pcr.saasPlanController.GetActiveSaasPlanList)
}

func (_pcr *SaasPlanRoutes) GetAllSaasPlanList() {
	get_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlan_router.HandleFunc("/all-saas-plan-list", _pcr.saasPlanController.GetAllSaasPlanList)
}

func (_pcr *SaasPlanRoutes) DeleteSaasPlan() {
	update_saasPlan_router := _pcr.ServeMuxRouter.Methods(http.MethodDelete).Subrouter()
	update_saasPlan_router.HandleFunc("/admin/delete-saas-plan/{id:[0-9a-fA-F]+}", _pcr.saasPlanController.DeleteSaasPlan)

}
