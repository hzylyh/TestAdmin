package vo

import "salotto/model/InterfaceTestPartEntity"

type CaseStepInfoVO struct {
	CaseId     string                                     `json:"caseId"`
	ItfId      string                                     `json:"interfaceId"`
	StepId     string                                     `json:"stepId"`
	StepStatus string                                     `json:"stepStatus"`
	StepLog    string                                     `json:"stepLog"`
	StepNum    int                                        `json:"stepNum"`
	StepName   string                                     `json:"stepName"`
	StepDesc   string                                     `json:"stepDesc"`
	ReqData    string                                     `json:"reqData"`
	ExpRes     string                                     `json:"expRes"`
	Variables  []InterfaceTestPartEntity.TCaseStepVarInfo `json:"variables"`
	//CollectCol      string `json:"collectCol"`
	//CollectColAlias string `json:"collectColAlias"`
	AssertInfos []InterfaceTestPartEntity.TAssertInfo `json:"assertInfos"`
	//AssertCol       string `json:"assertCol"`
	//Method          string `json:"method"`
	//ExpValue        string `json:"expValue"`
}
