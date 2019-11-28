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
	interfaceInfo.InterfaceId = utils.GenerateUUID()
	if err := service.DB.Create(interfaceInfo).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func (its *itfTestService) GetInterfaceList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.InterfaceInfo
	)

	if pageInfo, err = utils.PaginationWithDB(service.DB, &ret, qj); err != nil {
		return nil, err
	} else {
		return pageInfo, nil
	}
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
						return
					}
				}
			}
		}
	}
}
