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

type SpaceCategoryController struct {
	logging              *customLog.Logging
	spaceCategoryService *service.SpaceCategoryService
	errorCustom          *customError.ErrorCustom
}

func NewSpaceCategoryController() *SpaceCategoryController {
	return &SpaceCategoryController{
		logging:              customLog.NewLogging(),
		spaceCategoryService: service.NewSpaceCategoryService(),
		errorCustom:          customError.NewErrorCustom(),
	}
}

//This method is called by the spaceCategory.routes.go file. This runs when the program starts. It inserts all the default values if the data is not already availabe in the table
func (_sfc *SpaceCategoryController) CreateSpaceCategory() {

	_ = _sfc.spaceCategoryService.CreateSpaceCategory()

} //End of the function CreateSpaceCategory

func (_sfc *SpaceCategoryController) UpdateSpaceCategory(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	spacecategory_request := r.Context().Value(models.KeySpaceCategory{}).(dto.SpaceCategoryRequest)

	err = _sfc.spaceCategoryService.UpdateSpaceCategory(id_int, &spacecategory_request)

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
	rw.Write([]byte("SpaceCategory updated successfully"))

} //End of the function UpdateSpaceCategory

func (_sfc *SpaceCategoryController) AddRemoveProductContainer(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	categorycontainer_request := r.Context().Value(models.KeySpaceCategory{}).(dto.CategoryProductContainerRequest)

	err = _sfc.spaceCategoryService.AddRemoveProductContainer(id_int, &categorycontainer_request)

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
	rw.Write([]byte("SpaceCategory updated successfully"))

} //End of the function AddRemoveProductContainer

func (_sfc *SpaceCategoryController) EnableDisableSpaceCategory(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	spacecategory_enable_request := r.Context().Value(models.KeySpaceCategory{}).(dto.SpaceCategoryEnableRequest)

	err = _sfc.spaceCategoryService.EnableDisableSpaceCategory(id_int, &spacecategory_enable_request)

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
	rw.Write([]byte("SpaceCategory updated successfully"))

} //End of the function EnableDisableSpaceCategory

func (_sfc *SpaceCategoryController) GetSpaceCategoryById(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		invalid_err := customError.NewInvalidFieldErr("Id is not number or may be buffer overflow number", "Invalid Id", 0)
		_sfc.logging.ErrorLog(invalid_err.ErrorData.Message, err)
		http.Error(rw, _sfc.errorCustom.GetErrorData(invalid_err).ClientMessage, _sfc.errorCustom.GetErrorData(invalid_err).StatusCode)
	}

	space_category, err := _sfc.spaceCategoryService.GetSpaceCategoryById(id_int)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = space_category.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSectorCategoryList

func (_sfc *SpaceCategoryController) GetSpaceCategoryByCode(rw http.ResponseWriter, r *http.Request) {

	//Extract params from request
	vars := mux.Vars(r)
	code := vars["code"]

	space_category, err := _sfc.spaceCategoryService.GetSpaceCategoryByCode(code)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = space_category.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetSpaceCategoryByCode

func (_sfc *SpaceCategoryController) GetActiveSpaceCategoryList(rw http.ResponseWriter, r *http.Request) {

	spacecategory_reslist, err := _sfc.spaceCategoryService.GetActiveSpaceCategoryList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = spacecategory_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetActiveSpaceCategoryList

func (_sfc *SpaceCategoryController) GetAllSpaceCategoryList(rw http.ResponseWriter, r *http.Request) {

	spacecategory_reslist, err := _sfc.spaceCategoryService.GetAllSpaceCategoryList()

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

	err = spacecategory_reslist.ToJSON(rw)

	if err != nil {
		http.Error(rw, _sfc.errorCustom.GetErrorData(err).ClientMessage, _sfc.errorCustom.GetErrorData(err).StatusCode)
		return
	}

} //End of the function GetAllSpaceCategoryList
