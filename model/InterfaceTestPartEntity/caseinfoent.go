package InterfaceTestPartEntity

import "github.com/jinzhu/gorm"

type ItfCaseInfo struct {
	gorm.Model
	ModuleId int    `json:"moduleId"`
	CaseName string `json:"caseName"`
	CaseDesc string `json:"caseDesc"`
}
