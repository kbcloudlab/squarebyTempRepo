package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/saas"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/saas"
)

type SaasPlanAdsRoutes struct {
	ServeMuxRouter        *(mux.Router)
	saasPlanAdsController *controller.SaasPlanAdsController
	saasPlanAdsMiddleware *middleware.SaasPlanAdsMiddleware
}

func NewSaasPlanAdsRoutes(_sm *(mux.Router)) *SaasPlanAdsRoutes {
	return &SaasPlanAdsRoutes{
		ServeMuxRouter:        _sm,
		saasPlanAdsController: controller.NewSaasPlanAdsController(),
		saasPlanAdsMiddleware: middleware.NewSaasPlanAdsMiddleware(),
	}
}

func (_pcr *SaasPlanAdsRoutes) CreateSaasPlanAds() {
	create_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodPost).Subrouter()
	create_saasPlanAds_router.HandleFunc("/admin/create-saas-plan-ads", _pcr.saasPlanAdsController.CreateSaasPlanAds)

	//Validation with Middleware
	create_saasPlanAds_router.Use(_pcr.saasPlanAdsMiddleware.CreateUpdateSaasPlanAds)
}

func (_pcr *SaasPlanAdsRoutes) UpdateSaasPlanAds() {
	update_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasPlanAds_router.HandleFunc("/admin/update-saas-plan-ads/{id:[0-9a-fA-F]+}", _pcr.saasPlanAdsController.UpdateSaasPlanAds)

	//Validation with Middleware
	update_saasPlanAds_router.Use(_pcr.saasPlanAdsMiddleware.CreateUpdateSaasPlanAds)
}

func (_pcr *SaasPlanAdsRoutes) EnableDisableSaasPlanAds() {
	update_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasPlanAds_router.HandleFunc("/admin/enabledisable-saas-plan-ads/{id:[0-9a-fA-F]+}", _pcr.saasPlanAdsController.EnableDisableSaasPlanAds)

	//Validation with Middleware
	update_saasPlanAds_router.Use(_pcr.saasPlanAdsMiddleware.EnableDisableSaasPlanAds)
}

func (_pcr *SaasPlanAdsRoutes) PublishSaasPlanAds() {
	update_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_saasPlanAds_router.HandleFunc("/admin/publish-saas-plan-ads/{id:[0-9a-fA-F]+}", _pcr.saasPlanAdsController.PublishSaasPlanAds)

	//Validation with Middleware
	update_saasPlanAds_router.Use(_pcr.saasPlanAdsMiddleware.PublishSaasPlanAds)
}

func (_pcr *SaasPlanAdsRoutes) GetSaasPlanAdsById() {
	get_saasPlanAds_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlanAds_byid_router.HandleFunc("/saas-plan-ads/{id:[0-9a-fA-F]+}", _pcr.saasPlanAdsController.GetSaasPlanAdsById)
}

func (_pcr *SaasPlanAdsRoutes) GetActiveSaasPlanAdsList() {
	get_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlanAds_router.HandleFunc("/active-saas-plan-ads-list", _pcr.saasPlanAdsController.GetActiveSaasPlanAdsList)
}

func (_pcr *SaasPlanAdsRoutes) GetPublishedSaasPlanAdsList() {
	get_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlanAds_router.HandleFunc("/published-saas-plan-ads-list", _pcr.saasPlanAdsController.GetPublishedSaasPlanAdsList)
}

func (_pcr *SaasPlanAdsRoutes) GetAllSaasPlanAdsList() {
	get_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_saasPlanAds_router.HandleFunc("/all-saas-plan-ads-list", _pcr.saasPlanAdsController.GetAllSaasPlanAdsList)
}

func (_pcr *SaasPlanAdsRoutes) DeleteSaasPlanAds() {
	update_saasPlanAds_router := _pcr.ServeMuxRouter.Methods(http.MethodDelete).Subrouter()
	update_saasPlanAds_router.HandleFunc("/admin/delete-saas-plan-ads/{id:[0-9a-fA-F]+}", _pcr.saasPlanAdsController.DeleteSaasPlanAds)

}
