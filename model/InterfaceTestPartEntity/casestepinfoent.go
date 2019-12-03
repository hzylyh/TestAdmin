package InterfaceTestPartEntity

import "salotto/conf"

type ItfCaseStepInfo struct {
	conf.Model
	CaseId      string       `json:"caseId"`
	ItfId       string       `json:"interfaceId"`
	StepId      string       `json:"stepId"`
	StepNum     int          `json:"stepNum"`
	StepName    string       `json:"stepName"`
	StepDesc    string       `json:"stepDesc"`
	ReqData     string       `json:"reqData" gorm:"type:text"`
	ExpRes      string       `json:"expRes"  gorm:"type:text"`
	AssertInfos []AssertInfo `json:"assertInfos" gorm:"foreignkey:StepId;association_foreignkey:StepId"`
	//Variables string `json:"variables" gorm:"type:text"`
	Variables []CaseStepVarInfo `json:"variables" gorm:"foreignkey:StepId;association_foreignkey:StepId""`
}

type CaseStepVarInfo struct {
	conf.Model
	StepId          string
	CollectCol      string
	CollectColAlias string
}

type AssertInfo struct {
	conf.Model
	StepId    string
	AssertCol string `json:"assertCol"`
	Method    string `json:"method"`
	ExpValue  string `json:"expValue"`
}
