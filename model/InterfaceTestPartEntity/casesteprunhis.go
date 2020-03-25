package InterfaceTestPartEntity

import "time"

type TItfCaseStepRunHis struct {
	//conf.Model
	StepHisId  string    `json:"stepHisId"`
	BeginTime  string    `json:"beginTime"`
	CaseId     string    `json:"caseId"`
	StepId     string    `json:"stepId"`
	StepStatus string    `json:"stepStatus"`
	StepLog    string    `json:"stepLog"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
