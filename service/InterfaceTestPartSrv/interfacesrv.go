package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
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
		pageNum  = qj.GetInt("pageNum")
		pageSize = qj.GetInt("pageSize")
		total    float64
	)

	if err := service.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&ret); err != nil {
		fmt.Println(err)
	}

	// 获取总条数
	service.DB.Model(&ret).Count(&total)

	pageInfo = &model.PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    total,
		List:     ret,
	}
	return pageInfo
}
