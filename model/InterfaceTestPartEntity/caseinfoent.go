package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type TItfCaseInfo struct {
	conf.Model
	CaseId   string `json:"caseId"`
	ModuleId string `json:"moduleId"`
	CaseName string `json:"caseName"`
	CaseDesc string `json:"caseDesc"`
}
