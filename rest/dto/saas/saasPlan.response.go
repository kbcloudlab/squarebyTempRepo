package saas

import (
	"encoding/json"
	"io"

	saasBridge "squareby.com/admin/cloudspacemanager/sidechain/bridge/saas"

	models "squareby.com/admin/cloudspacemanager/src/models/saas"
	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	customLog "squareby.com/admin/cloudspacemanager/src/programfiles/logging"
)

type SaasPlanResponse struct {
	Id                   string               `json:"id"`
	PlanName             string               `json:"planName"`
	ValidityDuration     uint16               `json:"validityDuration"`
	Price                float32              `json:"price"`
	IsRenewable          bool                 `json:"isRenewable"`
	FeatureCountTypeList []*countTypeResponse `json:"featureCountTypeList"`
	FeatureBoolTypeList  []*boolTypeResponse  `json:"featureBoolTypeList"`
}

type countTypeResponse struct { //Ex: "Number of Cloudspaces per account" = 2
	SaasFeatureResponse *SaasFeatureResponse `json:"saasFeatureResponse"`
	Count               int16                `json:"count"`
}

type boolTypeResponse struct { //Ex: "Allow order management" = false
	SaasFeatureResponse *SaasFeatureResponse `json:"saasFeatureResponse"`
	Allow               bool                 `json:"allow"`
}

type SaasPlanResponseList []*SaasPlanResponse

func newCountTypeResponse(_countTypeModel *models.CountType, _featureBridge *saasBridge.SaasFeatureBridge) (*countTypeResponse, error) {

	saas_feature, err := _featureBridge.GetSaasFeatureByCode(_countTypeModel.SaasFeatureCode)

	if err != nil {
		return nil, err
	}

	return &countTypeResponse{
		SaasFeatureResponse: NewSaasFeatureResponse(saas_feature),
		Count:               _countTypeModel.Count,
	}, nil

}

func newBoolTypeResponse(_boolTypeModel *models.BoolType, _featureBridge *saasBridge.SaasFeatureBridge) (*boolTypeResponse, error) {

	saas_feature, err := _featureBridge.GetSaasFeatureByCode(_boolTypeModel.SaasFeatureCode)

	if err != nil {
		return nil, err
	}

	return &boolTypeResponse{
		SaasFeatureResponse: NewSaasFeatureResponse(saas_feature),
		Allow:               _boolTypeModel.Allow,
	}, nil

}

func NewSaasPlanResponse(_saasPlan *models.SaasPlan) (*SaasPlanResponse, error) {

	feature_bridge := saasBridge.NewSaasFeatureBridge()

	counttype_res_list := []*countTypeResponse{}
	for _, counttype := range _saasPlan.FeatureCountTypeList {
		counttype_res, err := newCountTypeResponse(counttype, feature_bridge)
		if err != nil {
			return nil, err
		}
		counttype_res_list = append(counttype_res_list, counttype_res)
	}

	booltype_res_list := []*boolTypeResponse{}
	for _, booltype := range _saasPlan.FeatureBoolTypeList {
		booltype_res, err := newBoolTypeResponse(booltype, feature_bridge)
		if err != nil {
			return nil, err
		}
		booltype_res_list = append(booltype_res_list, booltype_res)
	}

	return &SaasPlanResponse{
		Id:                   _saasPlan.Id,
		PlanName:             _saasPlan.PlanName,
		ValidityDuration:     _saasPlan.ValidityDuration,
		Price:                _saasPlan.Price,
		IsRenewable:          _saasPlan.IsRenewable,
		FeatureCountTypeList: counttype_res_list,
		FeatureBoolTypeList:  booltype_res_list,
	}, nil
}

func (_spr *SaasPlanResponse) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanResponse to JSON:: <<saasPlanResponse.response.go -> (_spr *SaasPlanResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spr *SaasPlanResponseList) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanResponseList to JSON:: <<saasPlanResponse.response.go -> (_spr *SaasPlanResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

type SaasPlanResponse1 struct {
	Id               string  `json:"id"`
	PlanName         string  `json:"planName"`
	ValidityDuration uint16  `json:"validityDuration"`
	Price            float32 `json:"price"`
	IsRenewable      bool    `json:"isRenewable"`
	Enabled          bool    `json:"enabled"`
}

type SaasPlanResponse1List []*SaasPlanResponse1

func NewSaasPlanResponse1(_saasPlan *models.SaasPlan) *SaasPlanResponse1 {
	return &SaasPlanResponse1{
		Id:               _saasPlan.Id,
		PlanName:         _saasPlan.PlanName,
		ValidityDuration: _saasPlan.ValidityDuration,
		Price:            _saasPlan.Price,
		IsRenewable:      _saasPlan.IsRenewable,
		Enabled:          _saasPlan.Enabled,
	}
}

func (_spr *SaasPlanResponse1) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanResponse1 to JSON:: <<saasPlanResponse.response.go -> (_spr *SaasPlanResponse) ToJSON()", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

func (_spr *SaasPlanResponse1List) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(_spr)

	if err != nil {
		//[SystemErr]
		system_err := customError.NewSystemErr("could not convert SaasPlanResponse1List to JSON:: <<saasPlanResponse.response.go -> (_spr *SaasPlanResponseList) ToJSON()>>", "", 0)
		customLog.NewLogging().ErrorLog(system_err.ErrorData.Message, system_err)
		return system_err
	}

	return nil
}

// type SaasPlanResponse2 struct {
// 	Id       string `json:"id"`
// 	PlanName string `json:"planName"`
// }

// func NewSaasPlanResponse2(_saasPlan *models.SaasPlan) *SaasPlanResponse2 {
// 	return &SaasPlanResponse2{
// 		Id:       _saasPlan.Id,
// 		PlanName: _saasPlan.PlanName,
// 	}
// }
