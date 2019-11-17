package InterfaceTestPartEntity

import "github.com/jinzhu/gorm"

type ItfCaseInfo struct {
	gorm.Model
	SuitId   int    `json:"suitId"`
	CaseName string `json:"caseName"`
	ReqData  string `json:"reqData" gorm:"type:text"`
	ExpRes   string `json:"expRes"  gorm:"type:text"`
	ItfId    int    `json:"interfaceId"`
}
