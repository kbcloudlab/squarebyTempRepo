package saas

import (
	"encoding/json"
	"io"

	saasBridge "squareby.com/admin/cloudspacemanager/sidechain/bridge/saas"
	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanAdsResponse struct {
	Id string `json:"id"`
	// SaasPlanId         string   `json:"saasPlanId"`
	SaasPlanResponse   *saasPlanResponse `json:"saasPlanResponse"`
	DisplayIndex       int16             `json:"displayIndex"` //Display order
	PlanName           string            `json:"planName"`
	FeatureDescription []string          `json:"featureDescription"`
	Published          bool              `json:"published"`
	Enabled            bool              `json:"enabled"`
}

type saasPlanResponse struct {
	Id       string  `json:"id"`
	PlanName string  `json:"planName"`
	Price    float32 `json:"price"`
}

type SaasPlanAdsResponseList []*SaasPlanAdsResponse

func NewSaasPlanAdsResponse(_saasPlanAds *models.SaasPlanAds) (*SaasPlanAdsResponse, error) {

	saas_plan_bridge := saasBridge.NewSaasPlanBridge()
	saas_plan, err := saas_plan_bridge.GetSaasPlanById(_saasPlanAds.SaasPlanId)

	var saas_plan_res *saasPlanResponse
	if err == nil {
		saas_plan_res = &saasPlanResponse{
			Id:       saas_plan.Id,
			PlanName: saas_plan.PlanName,
			Price:    saas_plan.Price,
		}
	}

	return &SaasPlanAdsResponse{
		Id:                 _saasPlanAds.Id,
		SaasPlanResponse:   saas_plan_res,
		DisplayIndex:       _saasPlanAds.DisplayIndex,
		PlanName:           _saasPlanAds.PlanName,
		FeatureDescription: _saasPlanAds.FeatureDescription,
		Published:          _saasPlanAds.Published,
		Enabled:            _saasPlanAds.Enabled,
	}, nil
}

func (_spr *SaasPlanAdsResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanAdsResponse to JSON:: <<saasPlanAdsResponse.response.go -> (_spr *SaasPlanAdsResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spr *SaasPlanAdsResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanAdsResponseList to JSON:: <<saasPlanAdsResponse.response.go -> (_spr *SaasPlanAdsResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

type SaasPlanAdsResponse1 struct {
	Id               string            `json:"id"`
	SaasPlanResponse *saasPlanResponse `json:"saasPlanResponse"`
	DisplayIndex     int16             `json:"displayIndex"` //Display order
	PlanName         string            `json:"planName"`
	Published        bool              `json:"published"`
	Enabled          bool              `json:"enabled"`
}

type SaasPlanAdsResponse1List []*SaasPlanAdsResponse1

func NewSaasPlanAdsResponse1(_saasPlanAds *models.SaasPlanAds) *SaasPlanAdsResponse1 {

	saas_plan_bridge := saasBridge.NewSaasPlanBridge()
	saas_plan, err := saas_plan_bridge.GetSaasPlanById(_saasPlanAds.SaasPlanId)

	var saas_plan_res *saasPlanResponse
	if err == nil {
		saas_plan_res = &saasPlanResponse{
			Id:       saas_plan.Id,
			PlanName: saas_plan.PlanName,
		}
	}

	return &SaasPlanAdsResponse1{
		Id:               _saasPlanAds.Id,
		SaasPlanResponse: saas_plan_res,
		DisplayIndex:     _saasPlanAds.DisplayIndex,
		PlanName:         _saasPlanAds.PlanName,
		Published:        _saasPlanAds.Published,
		Enabled:          _saasPlanAds.Enabled,
	}
}

func (_spr *SaasPlanAdsResponse1) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanAdsResponse1 to JSON:: <<saasPlanAdsResponse.response.go -> (_spr *SaasPlanAdsResponse1) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spr *SaasPlanAdsResponse1List) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanAdsResponse1List to JSON:: <<saasPlanAdsResponse.response.go -> (_spr *SaasPlanAdsResponse1List) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
