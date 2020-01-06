package InterfaceTestPartEntity

import "github.com/jinzhu/gorm"

type TItfCaseInfo struct {
	gorm.Model
	CaseId   string `json:"caseId"`
	ModuleId string `json:"moduleId"`
	CaseName string `json:"caseName"`
	CaseDesc string `json:"caseDesc"`
}
