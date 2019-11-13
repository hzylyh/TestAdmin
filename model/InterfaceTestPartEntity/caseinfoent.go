package InterfaceTestPartEntity

import "github.com/jinzhu/gorm"

type CaseInfoEntity struct {
	gorm.Model
	SuitId   int    `json:"suit_id"`
	CaseName string `json:"case_name"`
	ReqData  string `json:"req_data"`
	ItfId    int    `json:"interface_id"`
}
