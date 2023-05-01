package saas

import (
	"encoding/json"
	"io"

	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanRequest struct {
	Id                   string              `json:"id"`
	PlanName             string              `json:"planName"`
	ValidityDuration     uint16              `json:"validityDuration"`
	Price                float32             `json:"price"`
	IsRenewable          bool                `json:"isRenewable"`
	FeatureCountTypeList []*countTypeRequest `bson:"featureCountTypeList"`
	FeatureBoolTypeList  []*boolTypeRequest  `bson:"featureBoolTypeList"`
}

type countTypeRequest struct { //Ex: "Number of Cloudspaces per account" = 2
	SaasFeatureCode string `json:"saasFeatureCode"`
	Count           int16  `json:"count"`
}

type boolTypeRequest struct { //Ex: "Allow order management" = false
	SaasFeatureCode string `json:"saasFeatureCode"`
	Allow           bool   `json:"allow"`
}

func (_spr *SaasPlanRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasPlanRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spr *SaasPlanRequest) NewSaasPlanFromRequest(_mapId bool) *models.SaasPlan {

	feature_counttype_list := []*models.CountType{}
	for _, counttype_req := range _spr.FeatureCountTypeList {
		feature_counttype_list = append(feature_counttype_list, &models.CountType{
			SaasFeatureCode: counttype_req.SaasFeatureCode,
			Count:           counttype_req.Count,
		})
	}

	feature_booltype_list := []*models.BoolType{}
	for _, booltype_req := range _spr.FeatureBoolTypeList {
		feature_booltype_list = append(feature_booltype_list, &models.BoolType{
			SaasFeatureCode: booltype_req.SaasFeatureCode,
			Allow:           booltype_req.Allow,
		})
	}

	saas_pan := &models.SaasPlan{
		PlanName:             _spr.PlanName,
		ValidityDuration:     _spr.ValidityDuration,
		Price:                _spr.Price,
		IsRenewable:          _spr.IsRenewable,
		FeatureCountTypeList: feature_counttype_list,
		FeatureBoolTypeList:  feature_booltype_list,
	}

	if _mapId {
		saas_pan.Id = _spr.Id
	}

	return saas_pan

}

type SaasPlanEnableRequest struct {
	Id      string `json:"id"`
	Enabled bool   `json:"enabled"`
}

func (_spr *SaasPlanEnableRequest) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not parse from JSON data received from http request to SaasPlanEnableRequest", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}
