package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
	"salotto/utils/swagger"
)

var ItfTestSrv = &itfTestService{}

type itfTestService struct {
}

func (its *itfTestService) AddInterface(interfaceInfo *InterfaceTestPartEntity.InterfaceInfo) {
	if err := service.DB.Create(interfaceInfo).Error; err != nil {
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

func (its *itfTestService) ImportSwagger(qj *qjson.QJson) {
	var (
		res map[string]interface{}
	)
	res = swagger.SwaggerParse(qj.GetString("swaggerUrl"))
	for url, allType := range res {
		if allTypeInfo, ok := allType.(map[string]interface{}); ok {
			for reqType, itfInfo := range allTypeInfo {
				if itfSth, ok := itfInfo.(map[string]interface{}); ok {
					interfaceInfo := &InterfaceTestPartEntity.InterfaceInfo{
						Name: "",
						Url:  url,
						Type: reqType,
						Desc: itfSth["summary"].(string),
					}
					if err := service.DB.Create(interfaceInfo).Error; err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
