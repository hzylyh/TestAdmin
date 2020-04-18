package InterfaceTestPartEntity

import "salotto/conf"

type TItfCaseStepInfo struct {
	conf.Model
	CaseId   string `json:"caseId"`
	ItfId    string `json:"interfaceId"`
	StepId   string `json:"stepId"`
	StepNum  int    `json:"stepNum"`
	StepName string `json:"stepName"`
	StepDesc string `json:"stepDesc"`
	ReqData  string `json:"reqData" gorm:"type:text"`
	ExpRes   string `json:"expRes"  gorm:"type:text"`
	//AssertInfos []TAssertInfo `json:"assertInfos" gorm:"foreignkey:StepId;association_foreignkey:StepId"`
	////Variables string `json:"variables" gorm:"type:text"`
	//Variables []TCaseStepVarInfo `json:"variables" gorm:"foreignkey:StepId;association_foreignkey:StepId""`
}

type TCaseStepVarInfo struct {
	conf.Model
	VariableId      string `json:"variableId"`
	StepId          string `json:"stepId"`
	CollectCol      string `json:"collectCol"`
	CollectColAlias string `json:"collectColAlias"`
}

type TAssertInfo struct {
	conf.Model
	AssertId  string `json:"assertId"`
	StepId    string `json:"stepId"`
	AssertCol string `json:"assertCol"`
	Method    string `json:"method"`
	ExpValue  string `json:"expValue"`
}
