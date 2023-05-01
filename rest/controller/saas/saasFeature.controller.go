package saas

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"
	service "squareby.com/admin/cloudspacemanager/rest/service/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasFeatureController struct {
	logging            *customLog.Logging
	saasFeatureService *service.SaasFeatureService
	errorCustom        *customError.ErrorCustom
}

func NewSaasFeatureController() *SaasFeatureController {
	return &SaasFeatureController{
		logging:            customLog.NewLogging(),
		saasFeatureService: service.NewSaasFeatureService(),
		errorCustom:        customError.NewErrorCustom(),
	}
}

//This method is called by the saasFeature.routes.go file. This runs when the program starts. It inserts all the default values if the data is not already availabe in the table
func (_sfc *SaasFeatureController) CreateSaasFeature() {

	_ = _sfc.saasFeatureService.CreateSaasFeature()

} //End of the function CreateSaasFeature

func (_sfc *SaasFeatureController) UpdateSaasFeature(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	saasfeature_request := r.Context().Value(models.KeySaasFeature{}).(dto.SaasFeatureRequest)

	err = _sfc.saasFeatureService.UpdateSaasFeature(id_int, &saasfeature_request)

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
	rw.Write([]byte("SaasFeature updated successfully"))

} //End of the function UpdateSaasFeature

func (_sfc *SaasFeatureController) EnableDisableSaasFeature(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	saasfeature_enable_request := r.Context().Value(models.KeySaasFeature{}).(dto.SaasFeatureEnableRequest)

	err = _sfc.saasFeatureService.EnableDisableSaasFeature(id_int, &saasfeature_enable_request)

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
	rw.Write([]byte("SaasFeature updated successfully"))

} //End of the function EnableDisableSaasFeature

func (_sfc *SaasFeatureController) GetSaasFeatureById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	saas_feature, err := _sfc.saasFeatureService.GetSaasFeatureById(id_int)

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

func (_sfc *SaasFeatureController) GetSaasFeatureByCode(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	code := vars["code"]

	saas_feature, err := _sfc.saasFeatureService.GetSaasFeatureByCode(code)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_feature.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSaasFeatureByCode

func (_sfc *SaasFeatureController) GetActiveSaasFeatureList(rw http.ResponseWriter, r *http.Request) {

	saasfeature_reslist, err := _sfc.saasFeatureService.GetActiveSaasFeatureList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasfeature_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveSaasFeatureList

func (_sfc *SaasFeatureController) GetAllSaasFeatureList(rw http.ResponseWriter, r *http.Request) {

	saasfeature_reslist, err := _sfc.saasFeatureService.GetAllSaasFeatureList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasfeature_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllSaasFeatureList
