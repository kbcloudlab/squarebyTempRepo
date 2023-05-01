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

type SaasPlanAdsController struct {
	logging            *customLog.Logging
	saasPlanAdsService *service.SaasPlanAdsService
	errorCustom        *customError.ErrorCustom
}

func NewSaasPlanAdsController() *SaasPlanAdsController {
	return &SaasPlanAdsController{
		logging:            customLog.NewLogging(),
		saasPlanAdsService: service.NewSaasPlanAdsService(),
		errorCustom:        customError.NewErrorCustom(),
	}
}

func (_spc *SaasPlanAdsController) CreateSaasPlanAds(rw http.ResponseWriter, r *http.Request) {

	saasplan_request := r.Context().Value(models.KeySaasPlanAds{}).(dto.SaasPlanAdsRequest)

	_, err := _spc.saasPlanAdsService.CreateSaasPlanAds(&saasplan_request)

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
	rw.Write([]byte("SaasPlanAds successfully"))

} //End of the function CreateSaasPlanAds

func (_sfc *SaasPlanAdsController) UpdateSaasPlanAds(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saasplan_request := r.Context().Value(models.KeySaasPlanAds{}).(dto.SaasPlanAdsRequest)

	err := _sfc.saasPlanAdsService.UpdateSaasPlanAds(id, &saasplan_request)

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
	rw.Write([]byte("SaasPlanAds updated successfully"))

} //End of the function UpdateSaasPlanAds

func (_sfc *SaasPlanAdsController) EnableDisableSaasPlanAds(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saasplan_request := r.Context().Value(models.KeySaasPlanAds{}).(dto.SaasPlanAdsEnableRequest)

	err := _sfc.saasPlanAdsService.EnableDisableSaasPlanAds(id, &saasplan_request)

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
	rw.Write([]byte("SaasPlanAds updated successfully"))

} //End of the function EnableDisableSaasPlanAds

func (_sfc *SaasPlanAdsController) PublishSaasPlanAds(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	err := _sfc.saasPlanAdsService.PublishSaasPlanAds(id)

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
	rw.Write([]byte("SaasPlanAds published successfully"))

} //End of the function PublishSaasPlanAds

func (_sfc *SaasPlanAdsController) GetSaasPlanAdsById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	saas_plan, err := _sfc.saasPlanAdsService.GetSaasPlanAdsById(id)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saas_plan.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSaasPlanAdsById

func (_sfc *SaasPlanAdsController) GetActiveSaasPlanAdsList(rw http.ResponseWriter, r *http.Request) {

	saasplan_reslist, err := _sfc.saasPlanAdsService.GetActiveSaasPlanAdsList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasplan_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveSaasPlanAdsList

func (_sfc *SaasPlanAdsController) GetPublishedSaasPlanAdsList(rw http.ResponseWriter, r *http.Request) {

	saasplan_reslist, err := _sfc.saasPlanAdsService.GetPublishedSaasPlanAdsList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasplan_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetPublishedSaasPlanAdsList

func (_sfc *SaasPlanAdsController) GetAllSaasPlanAdsList(rw http.ResponseWriter, r *http.Request) {

	saasplan_reslist, err := _sfc.saasPlanAdsService.GetAllSaasPlanAdsList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = saasplan_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllSaasPlanAdsList

func (_sfc *SaasPlanAdsController) DeleteSaasPlanAds(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	err := _sfc.saasPlanAdsService.DeleteSaasPlanAds(id)

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
	rw.Write([]byte("SaasPlanAds deleted successfully"))

} //End of the function DeleteSaasPlanAds
