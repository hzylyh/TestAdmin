package InterfaceTestPartEntity

import "salotto/conf"

type ItfCaseStepInfo struct {
	conf.Model
	StepNum  int    `json:"stepNum"`
	StepName string `json:"stepName"`
	ReqData  string `json:"reqData" gorm:"type:text"`
	ExpRes   string `json:"expRes"  gorm:"type:text"`
	CaseId   string `json:"caseId"`
	ItfId    string `json:"interfaceId"`
}