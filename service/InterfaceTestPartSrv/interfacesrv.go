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

func (its *itfTestService) AddInterface(interfaceInfo *InterfaceTestPartEntity.TInterfaceInfo) {
	interfaceInfo.InterfaceId = utils.GenerateUUID()
	if err := service.DB.Create(interfaceInfo).Error; err != nil {
		fmt.Println(err)
		return
	}
}

func (its *itfTestService) GetInterfaceList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.TInterfaceInfo
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
					interfaceInfo := &InterfaceTestPartEntity.TInterfaceInfo{
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

func (its *itfTestService) GetInterfaceSelectOptions(json *qjson.QJson) (ret []*InterfaceTestPartEntity.TInterfaceInfo, err error) {
	json.GetString("projectId")
	if service.DB.Find(&ret); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}

func (its *itfTestService) GetSingleInterfaceInfo(q *qjson.QJson) (InterfaceTestPartEntity.TInterfaceInfo, error) {
	var (
		ret InterfaceTestPartEntity.TInterfaceInfo
		err error
	)
	itfId := q.GetNum("id")
	if err = service.DB.Where("id = ?", itfId).First(&ret).Error; err != nil {
		fmt.Println(err)
	}
	return ret, err
}

func (its *itfTestService) EditInterface(itfInfo *InterfaceTestPartEntity.TInterfaceInfo) error {
	fmt.Println(itfInfo)
	if err := service.DB.Model(&InterfaceTestPartEntity.TInterfaceInfo{}).Where("id = ?", itfInfo.ID).Update(itfInfo).Error; err != nil {
		return err
	}
	return nil
}

func (its *itfTestService) DelInterface(q *qjson.QJson) error {
	itfId := q.GetNum("id")
	if err := service.DB.Model(&InterfaceTestPartEntity.TInterfaceInfo{}).Where("id = ?", itfId).Delete(&InterfaceTestPartEntity.TInterfaceInfo{}).Error; err != nil {
		return err
	}
	return nil
}
