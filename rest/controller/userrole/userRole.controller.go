package userrole

import (
	"net/http"

	"github.com/gorilla/mux"
	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"

	service "squareby.com/admin/cloudspacemanager/rest/service/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type UserRoleController struct {
	logging         *customLog.Logging
	userRoleService *service.UserRoleService
	errorCustom     *customError.ErrorCustom
}

func NewUserRoleController() *UserRoleController {
	return &UserRoleController{
		logging:         customLog.NewLogging(),
		userRoleService: service.NewUserRoleService(),
		errorCustom:     customError.NewErrorCustom(),
	}
}

func (_urc *UserRoleController) CreateUserRole(rw http.ResponseWriter, r *http.Request) {

	userrole_request := r.Context().Value(models.KeyUserRole{}).(dto.UserRoleRequest)

	_, err := _urc.userRoleService.CreateUserRole(&userrole_request)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Role created successfully"))

} //End of the function CreateUserRole

func (_urc *UserRoleController) UpdateUserRole(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	userrole_request := r.Context().Value(models.KeyUserRole{}).(dto.UserRoleRequest)

	err := _urc.userRoleService.UpdateUserRole(id, &userrole_request)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Role updated successfully"))

} //End of the function UpdateUserRole

func (_urc *UserRoleController) GetUserRoleById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userrole_response, err := _urc.userRoleService.GetUserRoleById(id)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = userrole_response.ToJSON(rw)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetBrandById

func (_urc *UserRoleController) GetActiveUserRoleList(rw http.ResponseWriter, r *http.Request) {

	userrole_response_list, err := _urc.userRoleService.GetActiveUserRoleList()

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = userrole_response_list.ToJSON(rw)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveUserRoleList

func (_urc *UserRoleController) GetAllUserRoleList(rw http.ResponseWriter, r *http.Request) {

	userrole_response_list, err := _urc.userRoleService.GetAllUserRoleList()

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = userrole_response_list.ToJSON(rw)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllUserRoleList

func (_urc *UserRoleController) EnableDisableUserRole(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	userrole_enable_request := r.Context().Value(models.KeyUserRole{}).(dto.UserRoleEnableRequest)

	err := _urc.userRoleService.EnableDisableUserRole(id, &userrole_enable_request)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Role updated successfully"))

} //End of the function EnableDisableUserRole

func (_urc *UserRoleController) DeleteUserRole(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	err := _urc.userRoleService.DeleteUserRole(id)

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		_urc.logging.ErrorLog("", err)
		http.Error(rw, _urc.errorCustom.GetErrorData(err).ClientMessage, _urc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Role deleted successfully"))

} //End of the function DeleteUserRole
