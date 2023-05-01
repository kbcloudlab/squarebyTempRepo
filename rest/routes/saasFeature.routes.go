package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/saas"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/saas"
)

type SaasFeatureRoutes struct {
	ServeMuxRouter        *(mux.Router)
	saasFeatureController *controller.SaasFeatureController
	saasFeatureMiddleware *middleware.SaasFeatureMiddleware
}

func NewSaasFeatureRoutes(_sm *(mux.Router)) *SaasFeatureRoutes {
	return &SaasFeatureRoutes{
		ServeMuxRouter:        _sm,
		saasFeatureController: controller.NewSaasFeatureController(),
		saasFeatureMiddleware: middleware.NewSaasFeatureMiddleware(),
	}
}

func (_pcr *SaasFeatureRoutes) CreateSaasFeature() {
	_pcr.saasFeatureController.CreateSaasFeature()
}

func (_pcr *SaasFeatureRoutes) UpdateSaasFeature() {
	update_saasFeature_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasFeature_router.HandleFunc("/admin/update-saas-feature/{id:[0-9]+}", _pcr.saasFeatureController.UpdateSaasFeature)

	//Validation with Middleware
	update_saasFeature_router.Use(_pcr.saasFeatureMiddleware.UpdateSaasFeature)
}

func (_pcr *SaasFeatureRoutes) EnableDisableSaasFeature() {
	update_saasFeature_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasFeature_router.HandleFunc("/admin/enabledisable-saas-feature/{id:[0-9]+}", _pcr.saasFeatureController.EnableDisableSaasFeature)

	//Validation with Middleware
	update_saasFeature_router.Use(_pcr.saasFeatureMiddleware.EnableDisableSaasFeature)
}

func (_pcr *SaasFeatureRoutes) GetSaasFeatureById() {
	get_saasFeature_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasFeature_byid_router.HandleFunc("/saas-feature/{id:[0-9]+}", _pcr.saasFeatureController.GetSaasFeatureById)
}

func (_pcr *SaasFeatureRoutes) GetSaasFeatureByCode() {
	get_saasFeature_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasFeature_byid_router.HandleFunc("/saas-feature-by-code/{id:[0-9a-zA-Z]+}", _pcr.saasFeatureController.GetSaasFeatureByCode)
}

func (_pcr *SaasFeatureRoutes) GetActiveSaasFeatureList() {
	get_saasFeature_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasFeature_router.HandleFunc("/active-saas-feature-list", _pcr.saasFeatureController.GetActiveSaasFeatureList)
}

func (_pcr *SaasFeatureRoutes) GetAllSaasFeatureList() {
	get_saasFeature_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasFeature_router.HandleFunc("/all-saas-feature-list", _pcr.saasFeatureController.GetAllSaasFeatureList)
}
