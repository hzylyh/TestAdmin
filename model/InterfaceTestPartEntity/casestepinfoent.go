package InterfaceTestPartEntity

import "salotto/conf"

type ItfCaseStepInfo struct {
	conf.Model
	CaseId    string `json:"caseId"`
	ItfId     string `json:"interfaceId"`
	StepNum   int    `json:"stepNum"`
	StepName  string `json:"stepName"`
	ReqData   string `json:"reqData" gorm:"type:text"`
	ExpRes    string `json:"expRes"  gorm:"type:text"`
	Variables string `json:"variables" gorm:"type:text"`
}
