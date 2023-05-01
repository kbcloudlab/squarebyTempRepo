package sector

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dto "squareby.com/admin/cloudspacemanager/rest/dto/sector"
	service "squareby.com/admin/cloudspacemanager/rest/service/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SectorMasterController struct {
	logging             *customLog.Logging
	sectorMasterService *service.SectorMasterService
	errorCustom         *customError.ErrorCustom
}

func NewSectorMasterController() *SectorMasterController {
	return &SectorMasterController{
		logging:             customLog.NewLogging(),
		sectorMasterService: service.NewSectorMasterService(),
		errorCustom:         customError.NewErrorCustom(),
	}
}

//This method is called by the sectorMaster.routes.go file. This runs when the program starts. It inserts all the default values if the data is not already availabe in the table
func (_sfc *SectorMasterController) CreateSectorMaster() {

	_ = _sfc.sectorMasterService.CreateSectorMaster()

} //End of the function CreateSectorMaster

func (_sfc *SectorMasterController) UpdateSectorMaster(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	sectormaster_request := r.Context().Value(models.KeySectorMaster{}).(dto.SectorMasterRequest)

	err = _sfc.sectorMasterService.UpdateSectorMaster(id_int, &sectormaster_request)

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
	rw.Write([]byte("SectorMaster updated successfully"))

} //End of the function UpdateSectorMaster

func (_sfc *SectorMasterController) EnableDisableSectorMaster(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	sectormaster_enable_request := r.Context().Value(models.KeySectorMaster{}).(dto.SectorMasterEnableRequest)

	err = _sfc.sectorMasterService.EnableDisableSectorMaster(id_int, &sectormaster_enable_request)

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
	rw.Write([]byte("SectorMaster updated successfully"))

} //End of the function EnableDisableSectorMaster

func (_sfc *SectorMasterController) GetSectorMasterById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	sector_master, err := _sfc.sectorMasterService.GetSectorMasterById(id_int)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = sector_master.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSectorCategoryList

func (_sfc *SectorMasterController) GetSectorMasterByCode(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	code := vars["code"]

	sector_master, err := _sfc.sectorMasterService.GetSectorMasterByCode(code)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = sector_master.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSectorMasterByCode

func (_sfc *SectorMasterController) GetActiveSectorMasterList(rw http.ResponseWriter, r *http.Request) {

	sectormaster_reslist, err := _sfc.sectorMasterService.GetActiveSectorMasterList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = sectormaster_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveSectorMasterList

func (_sfc *SectorMasterController) GetAllSectorMasterList(rw http.ResponseWriter, r *http.Request) {

	sectormaster_reslist, err := _sfc.sectorMasterService.GetAllSectorMasterList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = sectormaster_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllSectorMasterList
