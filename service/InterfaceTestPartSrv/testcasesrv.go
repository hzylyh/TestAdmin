package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

var TestCaseSrv = &testCaseService{}

type testCaseService struct {
}

func (tcs testCaseService) AddCase(itfCaseInfo *InterfaceTestPartEntity.ItfCaseInfo) {
	if err := service.DB.Create(itfCaseInfo).Error; err != nil {
		fmt.Println(err)
	}
}

func (tcs testCaseService) GetCaseList(qj *qjson.QJson) interface{} {
	var (
		ret      []*InterfaceTestPartEntity.ItfCaseInfo
		pageInfo *model.PageInfo
	)

	pageInfo = utils.Pagination(&ret, qj)
	return pageInfo
}
