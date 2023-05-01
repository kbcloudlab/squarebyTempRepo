package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/sector"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/sector"
)

type SpaceCategoryRoutes struct {
	ServeMuxRouter          *(mux.Router)
	spaceCategoryController *controller.SpaceCategoryController
	spaceCategoryMiddleware *middleware.SpaceCategoryMiddleware
}

func NewSpaceCategoryRoutes(_sm *(mux.Router)) *SpaceCategoryRoutes {
	return &SpaceCategoryRoutes{
		ServeMuxRouter:          _sm,
		spaceCategoryController: controller.NewSpaceCategoryController(),
		spaceCategoryMiddleware: middleware.NewSpaceCategoryMiddleware(),
	}
}

func (_pcr *SpaceCategoryRoutes) CreateSpaceCategory() {
	_pcr.spaceCategoryController.CreateSpaceCategory()
}

func (_pcr *SpaceCategoryRoutes) UpdateSpaceCategory() {
	update_spaceCategory_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_spaceCategory_router.HandleFunc("/admin/update-space-category/{id:[0-9]+}", _pcr.spaceCategoryController.UpdateSpaceCategory)

	//Validation with Middleware
	update_spaceCategory_router.Use(_pcr.spaceCategoryMiddleware.UpdateSpaceCategory)
}

func (_pcr *SpaceCategoryRoutes) AddRemoveProductContainer() {
	update_spaceCategory_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_spaceCategory_router.HandleFunc("/admin/addremove-productcontainer-spacecategory/{id:[0-9]+}", _pcr.spaceCategoryController.AddRemoveProductContainer)

	//Validation with Middleware
	update_spaceCategory_router.Use(_pcr.spaceCategoryMiddleware.AddRemoveProductContainer)
}

func (_pcr *SpaceCategoryRoutes) EnableDisableSpaceCategory() {
	update_spaceCategory_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_spaceCategory_router.HandleFunc("/admin/enabledisable-space-category/{id:[0-9]+}", _pcr.spaceCategoryController.EnableDisableSpaceCategory)

	//Validation with Middleware
	update_spaceCategory_router.Use(_pcr.spaceCategoryMiddleware.EnableDisableSpaceCategory)
}

func (_pcr *SpaceCategoryRoutes) GetSpaceCategoryById() {
	get_spaceCategory_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_spaceCategory_byid_router.HandleFunc("/space-category/{id:[0-9]+}", _pcr.spaceCategoryController.GetSpaceCategoryById)
}

func (_pcr *SpaceCategoryRoutes) GetSpaceCategoryByCode() {
	get_spaceCategory_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_spaceCategory_byid_router.HandleFunc("/space-category-by-code/{code:[^\n]+}", _pcr.spaceCategoryController.GetSpaceCategoryByCode)
}

func (_pcr *SpaceCategoryRoutes) GetActiveSpaceCategoryList() {
	get_spaceCategory_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_spaceCategory_router.HandleFunc("/active-space-category-list", _pcr.spaceCategoryController.GetActiveSpaceCategoryList)
}

func (_pcr *SpaceCategoryRoutes) GetAllSpaceCategoryList() {
	get_spaceCategory_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_spaceCategory_router.HandleFunc("/all-space-category-list", _pcr.spaceCategoryController.GetAllSpaceCategoryList)
}
