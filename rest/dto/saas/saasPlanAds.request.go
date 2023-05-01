package saas

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanAdsRequest struct {
	Id                 string   `json:"id"`
	SaasPlanId         string   `json:"saasPlanId"`
	PlanName           string   `json:"planName"`
	FeatureDescription []string `json:"featureDescription"`
	// DisplayIndex       int16    `json:"displayIndex"` //This field will be added in "Sorting" request

}

func (_spar *SaasPlanAdsRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_spar)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasPlanAdsRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spar *SaasPlanAdsRequest) NewSaasPlanAdsFromRequest(_mapId bool) *models.SaasPlanAds {

	saas_plan_ads := &models.SaasPlanAds{
		SaasPlanId:         _spar.SaasPlanId,
		PlanName:           _spar.PlanName,
		FeatureDescription: _spar.FeatureDescription,
	}

	if _mapId {
		saas_plan_ads.Id = _spar.Id
	}

	return saas_plan_ads

}

type SaasPlanAdsEnableRequest struct {
	Id      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_spar *SaasPlanAdsEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_spar)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasPlanAdsEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
