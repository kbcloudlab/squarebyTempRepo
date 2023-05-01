package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "squareby.com/admin/cloudspacemanager/rest/controller/userrole"
	middleware "squareby.com/admin/cloudspacemanager/rest/middleware/userrole"
)

type UserRoleRoutes struct {
	ServeMuxRouter     *(mux.Router)
	userRoleController *controller.UserRoleController
	userRoleMiddleware *middleware.UserRoleMiddleware
}

func NewUserRoleRoutes(_sm *(mux.Router)) *UserRoleRoutes {
	return &UserRoleRoutes{
		ServeMuxRouter:     _sm,
		userRoleController: controller.NewUserRoleController(),
		userRoleMiddleware: middleware.NewUserRoleMiddleware(),
	}
}

func (_pcr *UserRoleRoutes) CreateUserRole() {
	userRole_router := _pcr.ServeMuxRouter.Methods(http.MethodPost).Subrouter()
	userRole_router.HandleFunc("/admin/create-user-role", _pcr.userRoleController.CreateUserRole)

	//Middleware
	userRole_router.Use(_pcr.userRoleMiddleware.CreateUpdateUserRole)
}

func (_pcr *UserRoleRoutes) UpdateUserRole() {
	update_userRole_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_userRole_router.HandleFunc("/admin/update-user-role/{id:[0-9a-fA-F]+}", _pcr.userRoleController.UpdateUserRole)

	//Validation with Middleware
	update_userRole_router.Use(_pcr.userRoleMiddleware.CreateUpdateUserRole)
}

func (_pcr *UserRoleRoutes) GetActiveUserRoleList() {
	get_userRole_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_userRole_router.HandleFunc("/active-user-role-list", _pcr.userRoleController.GetActiveUserRoleList)
}

func (_pcr *UserRoleRoutes) GetAllUserRoleList() {
	get_userRole_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_userRole_router.HandleFunc("/all-user-role-list", _pcr.userRoleController.GetAllUserRoleList)
}

func (_pcr *UserRoleRoutes) GetUserRoleById() {
	get_userRole_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodGet).Subrouter()
	get_userRole_byid_router.HandleFunc("/user-role/{id:[0-9a-fA-F]+}", _pcr.userRoleController.GetUserRoleById)
}

func (_pcr *UserRoleRoutes) EnableDisableUserRole() {
	update_userRole_router := _pcr.ServeMuxRouter.Methods(http.MethodPut).Subrouter()
	update_userRole_router.HandleFunc("/admin/enabledisable-user-role/{id:[0-9a-fA-F]+}", _pcr.userRoleController.EnableDisableUserRole)

	//Validation with Middleware
	update_userRole_router.Use(_pcr.userRoleMiddleware.EnableDisableUserRole)
}

func (_pcr *UserRoleRoutes) DeleteUserRole() {
	get_userRole_byid_router := _pcr.ServeMuxRouter.Methods(http.MethodDelete).Subrouter()
	get_userRole_byid_router.HandleFunc("/admin/delete-user-role/{id:[0-9a-fA-F]+}", _pcr.userRoleController.DeleteUserRole)
}
