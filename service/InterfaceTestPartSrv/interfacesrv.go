package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

var ItfTestSrv = &itfTestService{}

type itfTestService struct {
}

func (its *itfTestService) AddInterface(interfaceInfo *InterfaceTestPartEntity.InterfaceInfo) {
	if err := service.DB.Create(interfaceInfo); err != nil {
		fmt.Println(err)
	}
}

func (its *itfTestService) GetInterfaceList(qj *qjson.QJson) *model.PageInfo {
	var (
		ret      []*InterfaceTestPartEntity.InterfaceInfo
		pageInfo *model.PageInfo
	)

	pageInfo = utils.Pagination(&ret, qj)
	return pageInfo
}
