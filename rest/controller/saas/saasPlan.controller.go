package saas

import (
	"net/http"

	"github.com/gorilla/mux"
	dto "squareby.com/admin/cloudspacemanager/rest/dto/saas"
	service "squareby.com/admin/cloudspacemanager/rest/service/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanController struct {
	logging         *customLog.Logging
	saasPlanService *service.SaasPlanService
	errorCustom     *customError.ErrorCustom
}

func NewSaasPlanController() *SaasPlanController {
	return &SaasPlanController{
		logging:         customLog.NewLogging(),
		saasPlanService: service.NewSaasPlanService(),
		errorCustom:     customError.NewErrorCustom(),
	}
}

func (_spc *SaasPlanController) CreateSaasPlan(rw http.ResponseWriter, r *http.Request) {

	saasplan_request := r.Context().Value(models.KeySaasPlan{}).(dto.SaasPlanRequest)

	_, err := _spc.saasPlanService.CreateSaasPlan(&saasplan_request)

	if err != nil {
		_spc.logging.ErrorLog("", err)
		http.Error(rw, _spc.errorCustom.GetErrorData(err).ClientMessage, _spc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	if err != nil {
		_spc.logging.ErrorLog("", err)
		http.Error(rw, _spc.errorCustom.GetErrorData(err).ClientMessage, _spc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("SaasPlan successfully"))

} //End of the function CreateSaasPlan

func (_sfc *SaasPlanController) UpdateSaasPlan(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saasplan_request := r.Context().Value(models.KeySaasPlan{}).(dto.SaasPlanRequest)

	err := _sfc.saasPlanService.UpdateSaasPlan(id, &saasplan_request)

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
	rw.Write([]byte("SaasPlan updated successfully"))

} //End of the function UpdateSaasPlan

func (_sfc *SaasPlanController) EnableDisableSaasPlan(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saasplan_enable_request := r.Context().Value(models.KeySaasPlan{}).(dto.SaasPlanEnableRequest)

	err := _sfc.saasPlanService.EnableDisableSaasPlan(id, &saasplan_enable_request)

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
	rw.Write([]byte("SaasPlan updated successfully"))

} //End of the function EnableDisableSaasPlan

func (_sfc *SaasPlanController) GetSaasPlanById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saas_plan, err := _sfc.saasPlanService.GetSaasPlanById(id)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_plan.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSaasPlanById

func (_sfc *SaasPlanController) GetActiveSaasPlanList(rw http.ResponseWriter, r *http.Request) {

	saasplan_reslist, err := _sfc.saasPlanService.GetActiveSaasPlanList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasplan_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveSaasPlanList

func (_sfc *SaasPlanController) GetAllSaasPlanList(rw http.ResponseWriter, r *http.Request) {

	saasplan_reslist, err := _sfc.saasPlanService.GetAllSaasPlanList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasplan_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllSaasPlanList

func (_sfc *SaasPlanController) DeleteSaasPlan(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	err := _sfc.saasPlanService.DeleteSaasPlan(id)

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
	rw.Write([]byte("SaasPlan deleted successfully"))

} //End of the function DeleteSaasPlan
