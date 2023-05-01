package userrole

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dto "squareby.com/admin/cloudspacemanager/rest/dto/userrole"
	service "squareby.com/admin/cloudspacemanager/rest/service/userrole"
	models "squareby.com/admin/cloudspacemanager/src/models/userrole"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type PrivilegeActionController struct {
	logging                *customLog.Logging
	privilegeActionService *service.PrivilegeActionService
	errorCustom            *customError.ErrorCustom
}

func NewPrivilegeActionController() *PrivilegeActionController {
	return &PrivilegeActionController{
		logging:                customLog.NewLogging(),
		privilegeActionService: service.NewPrivilegeActionService(),
		errorCustom:            customError.NewErrorCustom(),
	}
}

//This method is called by the privilegeAction.routes.go file. This runs when the program starts. It inserts all the default values if the data is not already availabe in the table
func (_sfc *PrivilegeActionController) CreatePrivilegeAction() {

	_ = _sfc.privilegeActionService.CreatePrivilegeAction()

} //End of the function CreatePrivilegeAction

func (_sfc *PrivilegeActionController) UpdatePrivilegeAction(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	saasfeature_request := r.Context().Value(models.KeyPrivilegeAction{}).(dto.PrivilegeActionRequest)

	err = _sfc.privilegeActionService.UpdatePrivilegeAction(id_int, &saasfeature_request)

	if err != nil {
		// _sfc.logging.ErrorLog("", err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		// _sfc.logging.ErrorLog("", err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("PrivilegeAction updated successfully"))

} //End of the function UpdatePrivilegeAction

func (_pac *PrivilegeActionController) EnableDisablePrivilegeAction(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_pac.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _pac.errorCustom.GetErrorData(invalid_err).ClientMessage, _pac.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	action_enable_request := r.Context().Value(models.KeyPrivilegeAction{}).(dto.PrivilegeActionEnableRequest)

	err = _pac.privilegeActionService.EnableDisablePrivilegeAction(id_int, &action_enable_request)

	if err != nil {
		// _sfc.logging.ErrorLog("", err)
		http.Error(rw, _pac.errorCustom.GetErrorData(err).ClientMessage, _pac.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		// _sfc.logging.ErrorLog("", err)
		http.Error(rw, _pac.errorCustom.GetErrorData(err).ClientMessage, _pac.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("PrivilegeAction updated successfully"))

} //End of the function EnableDisableSaasPlanAds

func (_sfc *PrivilegeActionController) GetPrivilegeActionById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	saas_feature, err := _sfc.privilegeActionService.GetPrivilegeActionById(id_int)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_feature.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSectorCategoryList

func (_sfc *PrivilegeActionController) GetPrivilegeActionByCode(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	code := vars["code"]

	saas_feature, err := _sfc.privilegeActionService.GetPrivilegeActionByCode(code)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_feature.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetPrivilegeActionByCode

func (_sfc *PrivilegeActionController) GetPrivilegeActionByCodeNumber(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	code_number := vars["codeNumber"]

	code_number_int, err := strconv.ParseUint(code_number, 10, 16)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	// userrole_enable_request := r.Context().Value(models.KeyUserRole{}).(dto.UserRoleEnableRequest)

	saas_feature, err := _sfc.privilegeActionService.GetPrivilegeActionByCodeNumber(uint16(code_number_int))

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_feature.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetPrivilegeActionByCodeNumber

func (_sfc *PrivilegeActionController) GetActivePrivilegeActionList(rw http.ResponseWriter, r *http.Request) {

	saasfeature_reslist, err := _sfc.privilegeActionService.GetActivePrivilegeActionList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasfeature_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActivePrivilegeActionList

func (_sfc *PrivilegeActionController) GetAllPrivilegeActionList(rw http.ResponseWriter, r *http.Request) {

	saasfeature_reslist, err := _sfc.privilegeActionService.GetAllPrivilegeActionList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasfeature_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllPrivilegeActionList
