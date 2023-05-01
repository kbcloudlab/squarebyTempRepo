package sector

import (
	"encoding/json"
	"io"

	"github.com/lib/pq"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SpaceCategoryRequest struct {
	Id         uint64 `json:"id"`
	SectorCode string `json:"sectorCode"` //space category can be edited from one sector to another (ex: from service sector to product sector)
	Name       string `json:"name"`
}

func (_sfr *SpaceCategoryRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SpaceCategoryRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_sfr *SpaceCategoryRequest) NewSpaceCategoryFromRequest(_mapId bool) *models.SpaceCategory {

	space_category := &models.SpaceCategory{
		SectorCode: _sfr.SectorCode,
		Name:       _sfr.Name,
	}

	if _mapId {
		space_category.Id = _sfr.Id
	}

	return space_category

}

type SpaceCategoryEnableRequest struct {
	Id      uint64 `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_sfr *SpaceCategoryEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SpaceCategoryEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

type CategoryProductContainerRequest struct {
	Id                     uint64        `json:"id"`
	ProductContainerIdList pq.Int64Array `json:"productContainerIdList"`
}

func (_sfr *CategoryProductContainerRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_sfr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to CategoryProductContainerRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
