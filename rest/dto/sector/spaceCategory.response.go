package sector

import (
	"encoding/json"
	"io"

	"github.com/lib/pq"
	sectorBridge "squareby.com/admin/cloudspacemanager/sidechain/bridge/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

/** The GRPC call will just make the codes complicated unecessarily.
	Simply respond to client the ProductContainerIdList and then from the client retrieve the ProductContainerResponse list using the Id list
	by connecting to the ProductService server RESTfully
**/

type SpaceCategoryResponse struct {
	Id                     uint64                `json:"id"`
	SectorMasterResponse   *SectorMasterResponse `json:"sectorMasterResponse"`
	Code                   string                `json:"code"`
	Name                   string                `json:"name"`
	ImageURL               string                `json:"imageURL"`
	ProductContainerIdList pq.Int64Array         `json:"productContainerIdList"`
	Enabled                bool                  `json:"enabled"`
}

type SpaceCategoryResponseList []*SpaceCategoryResponse

func NewSpaceCategoryResponse(_spaceCategory *models.SpaceCategory) (*SpaceCategoryResponse, error) {

	sectormaster_bridge := sectorBridge.NewSectorMasterBridge()
	sectormaster, err := sectormaster_bridge.GetSectorMasterByCode(_spaceCategory.SectorCode)
	if err != nil {
		return nil, err
	}

	return &SpaceCategoryResponse{
		Id:                     _spaceCategory.Id,
		SectorMasterResponse:   NewSectorMasterResponse(sectormaster),
		Code:                   _spaceCategory.Code,
		Name:                   _spaceCategory.Name,
		ImageURL:               _spaceCategory.ImageURL,
		ProductContainerIdList: _spaceCategory.ProductContainerIdList,
		Enabled:                _spaceCategory.Enabled,
	}, nil
}

func (_csr *SpaceCategoryResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SpaceCategoryResponse to JSON:: <<spaceCategoryResponse.response.go -> (_csrl *SpaceCategoryResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_csrl *SpaceCategoryResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SpaceCategoryResponseList to JSON:: <<spaceCategoryResponse.response.go -> (_csrl *SpaceCategoryResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

type SpaceCategoryResponse1 struct {
	Id                     uint64                `json:"id"`
	SectorMasterResponse   *SectorMasterResponse `json:"sectorMasterResponse"`
	Code                   string                `json:"code"`
	Name                   string                `json:"name"`
	ImageURL               string                `json:"imageURL"`
	Enabled                bool                  `json:"enabled"`
}

type SpaceCategoryResponse1List []*SpaceCategoryResponse1

func NewSpaceCategoryResponse1(_spaceCategory *models.SpaceCategory) (*SpaceCategoryResponse1, error) {

	sectormaster_bridge := sectorBridge.NewSectorMasterBridge()
	sectormaster, err := sectormaster_bridge.GetSectorMasterByCode(_spaceCategory.SectorCode)
	if err != nil {
		return nil, err
	}

	return &SpaceCategoryResponse1{
		Id:                     _spaceCategory.Id,
		SectorMasterResponse:   NewSectorMasterResponse(sectormaster),
		Code:                   _spaceCategory.Code,
		Name:                   _spaceCategory.Name,
		ImageURL:               _spaceCategory.ImageURL,
		Enabled:                _spaceCategory.Enabled,
	}, nil
}

func (_csrl *SpaceCategoryResponse1List) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SpaceCategoryResponse1List to JSON:: <<spaceCategoryResponse.response.go -> (_csrl *SpaceCategoryResponse1List) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
