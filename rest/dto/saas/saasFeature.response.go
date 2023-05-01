package saas

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasFeatureResponse struct {
	Id          uint64 `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	FeatureType int16  `json:"featureType"`
	Enabled     bool   `json:"enabled"`
}

type SaasFeatureResponseList []*SaasFeatureResponse

func NewSaasFeatureResponse(_saasFeature *models.SaasFeature) *SaasFeatureResponse {

	return &SaasFeatureResponse{
		Id:          _saasFeature.Id,
		Code:        _saasFeature.Code,
		Title:       _saasFeature.Title,
		FeatureType: _saasFeature.FeatureType,
		Enabled:     _saasFeature.Enabled,
	}
}

func (_csr *SaasFeatureResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasFeatureResponse to JSON:: <<saasFeatureResponse.response.go -> (_csrl *SaasFeatureResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_csrl *SaasFeatureResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_csrl)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasFeatureResponseList to JSON:: <<saasFeatureResponse.response.go -> (_csrl *SaasFeatureResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
