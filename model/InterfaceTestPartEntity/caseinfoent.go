package InterfaceTestPartEntity

import "github.com/jinzhu/gorm"

type ItfCaseInfo struct {
	gorm.Model
	CaseId   string `json:"caseId"`
	ModuleId string `json:"moduleId"`
	CaseName string `json:"caseName"`
	CaseDesc string `json:"caseDesc"`
}
