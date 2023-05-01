package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/sector"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/sector"
)

type SectorMasterRoutes struct {
	ServeMuxRouter         *(mux.Router)
	sectorMasterController *controller.SectorMasterController
	sectorMasterMiddleware *middleware.SectorMasterMiddleware
}

func NewSectorMasterRoutes(_sm *(mux.Router)) *SectorMasterRoutes {
	return &SectorMasterRoutes{
		ServeMuxRouter:         _sm,
		sectorMasterController: controller.NewSectorMasterController(),
		sectorMasterMiddleware: middleware.NewSectorMasterMiddleware(),
	}
}

func (_pcr *SectorMasterRoutes) CreateSectorMaster() {
	_pcr.sectorMasterController.CreateSectorMaster()
}

func (_pcr *SectorMasterRoutes) UpdateSectorMaster() {
	update_sectorMaster_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_sectorMaster_router.HandleFunc("/admin/update-sector-master/{id:[0-9]+}", _pcr.sectorMasterController.UpdateSectorMaster)

	//Validation with Middleware
	update_sectorMaster_router.Use(_pcr.sectorMasterMiddleware.UpdateSectorMaster)
}

func (_pcr *SectorMasterRoutes) EnableDisableSectorMaster() {
	update_sectorMaster_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_sectorMaster_router.HandleFunc("/admin/enabledisable-sector-master/{id:[0-9]+}", _pcr.sectorMasterController.EnableDisableSectorMaster)

	//Validation with Middleware 
	update_sectorMaster_router.Use(_pcr.sectorMasterMiddleware.EnableDisableSectorMaster)
}

func (_pcr *SectorMasterRoutes) GetSectorMasterById() {
	get_sectorMaster_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_sectorMaster_byid_router.HandleFunc("/sector-master/{id:[0-9]+}", _pcr.sectorMasterController.GetSectorMasterById)
}

func (_pcr *SectorMasterRoutes) GetSectorMasterByCode() {
	get_sectorMaster_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_sectorMaster_byid_router.HandleFunc("/sector-master-by-code/{id:[0-9a-zA-Z]+}", _pcr.sectorMasterController.GetSectorMasterByCode)
}

func (_pcr *SectorMasterRoutes) GetActiveSectorMasterList() {
	get_sectorMaster_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_sectorMaster_router.HandleFunc("/active-sector-master-list", _pcr.sectorMasterController.GetActiveSectorMasterList)
}

func (_pcr *SectorMasterRoutes) GetAllSectorMasterList() {
	get_sectorMaster_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_sectorMaster_router.HandleFunc("/all-sector-master-list", _pcr.sectorMasterController.GetAllSectorMasterList)
}
