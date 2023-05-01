package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/userrole"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/userrole"
)

type PrivilegeActionRoutes struct {
	ServeMuxRouter            *(mux.Router)
	privilegeActionController *controller.PrivilegeActionController
	privilegeActionMiddleware *middleware.PrivilegeActionMiddleware
}

func NewPrivilegeActionRoutes(_sm *(mux.Router)) *PrivilegeActionRoutes {
	return &PrivilegeActionRoutes{
		ServeMuxRouter:            _sm,
		privilegeActionController: controller.NewPrivilegeActionController(),
		privilegeActionMiddleware: middleware.NewPrivilegeActionMiddleware(),
	}
}

func (_pcr *PrivilegeActionRoutes) CreatePrivilegeAction() {
	_pcr.privilegeActionController.CreatePrivilegeAction()
}

func (_pcr *PrivilegeActionRoutes) UpdatePrivilegeAction() {
	update_privilegeAction_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_privilegeAction_router.HandleFunc("/admin/update-privilege-action/{id:[0-9]+}", _pcr.privilegeActionController.UpdatePrivilegeAction)

	//Validation with Middleware
	update_privilegeAction_router.Use(_pcr.privilegeActionMiddleware.UpdatePrivilegeAction)
}

func (_pcr *PrivilegeActionRoutes) EnableDisablePrivilegeAction() {
	update_privilegeAction_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_privilegeAction_router.HandleFunc("/admin/enabledisable-privilege-action/{id:[0-9]+}", _pcr.privilegeActionController.EnableDisablePrivilegeAction)

	//Validation with Middleware
	update_privilegeAction_router.Use(_pcr.privilegeActionMiddleware.EnableDisablePrivilegeAction)
}

func (_pcr *PrivilegeActionRoutes) GetPrivilegeActionById() {
	get_privilegeAction_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_privilegeAction_byid_router.HandleFunc("/privilege-action/{id:[0-9]+}", _pcr.privilegeActionController.GetPrivilegeActionById)
}

func (_pcr *PrivilegeActionRoutes) GetPrivilegeActionByCode() {
	get_privilegeAction_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_privilegeAction_byid_router.HandleFunc("/privilege-action-by-code/{code:[0-9a-zA-Z]+}", _pcr.privilegeActionController.GetPrivilegeActionByCode)
}

func (_pcr *PrivilegeActionRoutes) GetPrivilegeActionByCodeNumber() {
	get_privilegeAction_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_privilegeAction_byid_router.HandleFunc("/privilege-action-by-codenumber/{codeNumber:[0-9]+}", _pcr.privilegeActionController.GetPrivilegeActionByCodeNumber)
}

func (_pcr *PrivilegeActionRoutes) GetActivePrivilegeActionList() {
	get_privilegeAction_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_privilegeAction_router.HandleFunc("/active-privilege-action-list", _pcr.privilegeActionController.GetActivePrivilegeActionList)
}

func (_pcr *PrivilegeActionRoutes) GetAllPrivilegeActionList() {
	get_privilegeAction_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_privilegeAction_router.HandleFunc("/all-privilege-action-list", _pcr.privilegeActionController.GetAllPrivilegeActionList)
}
