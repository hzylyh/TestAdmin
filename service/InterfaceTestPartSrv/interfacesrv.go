package InterfaceTestPartSrv

import (
	"fmt"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/model/vo"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
	"salotto/utils/swagger"
)

var ItfTestSrv = &itfTestService{}

type itfTestService struct {
}

//func (its *itfTestService) AddInterface(interfaceInfo *InterfaceTestPartEntity.TInterfaceInfo) {
//	interfaceInfo.InterfaceId = utils.GenerateUUID()
//	if err := service.DB.Create(interfaceInfo).Error; err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func (its *itfTestService) AddInterface(interfaceInfo *vo.InterfaceInfoVO) {

	interfaceInfo.InterfaceId = utils.GenerateUUID()
	newItfInfo := &InterfaceTestPartEntity.TInterfaceInfo{
		ProjectId:   interfaceInfo.ProjectId,
		InterfaceId: interfaceInfo.InterfaceId,
		Name:        interfaceInfo.Name,
		Url:         interfaceInfo.Url,
		Type:        interfaceInfo.Type,
		Desc:        interfaceInfo.Desc,
	}

	tx := service.DB.Begin()

	if err := tx.Create(newItfInfo).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	for _, headerInfo := range interfaceInfo.Headers {
		newHeaderInfo := &InterfaceTestPartEntity.TInterfaceHeadersInfo{
			HeaderId:    utils.GenerateUUID(),
			InterfaceId: interfaceInfo.InterfaceId,
			HeaderName:  headerInfo.HeaderName,
			HeaderValue: headerInfo.HeaderValue,
		}
		if err := tx.Create(newHeaderInfo).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return
		}

	}

	tx.Commit()

	//interfaceInfo.InterfaceId = utils.GenerateUUID()
	//if err := service.DB.Create(interfaceInfo).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
}

func (its *itfTestService) GetInterfaceList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.TInterfaceInfo
	)

	projectId := qj.GetString("projectId")

	if pageInfo, err = utils.PaginationWithDB(service.DB.Where("project_id = ?", projectId), &ret, qj); err != nil {
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

func (its *itfTestService) GetSingleInterfaceInfo(q *qjson.QJson) (*vo.InterfaceInfoVO, error) {
	var (
		ret     = InterfaceTestPartEntity.TInterfaceInfo{}
		headers = []InterfaceTestPartEntity.TInterfaceHeadersInfo{}
		err     error
	)
	itfId := q.GetString("interfaceId")
	if err = service.DB.Where("interface_id = ?", itfId).First(&ret).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = service.DB.Where("interface_id = ?", itfId).Find(&headers).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &vo.InterfaceInfoVO{
		ProjectId:   ret.ProjectId,
		InterfaceId: ret.InterfaceId,
		Name:        ret.Name,
		Url:         ret.Url,
		Type:        ret.Type,
		Headers:     headers,
		Desc:        ret.Desc,
	}, err
}

func (its *itfTestService) EditInterface(interfaceInfo *vo.InterfaceInfoVO) error {

	var (
		total float64
	)

	newItfInfo := &InterfaceTestPartEntity.TInterfaceInfo{
		ProjectId:   interfaceInfo.ProjectId,
		InterfaceId: interfaceInfo.InterfaceId,
		Name:        interfaceInfo.Name,
		Url:         interfaceInfo.Url,
		Type:        interfaceInfo.Type,
		Desc:        interfaceInfo.Desc,
	}

	tx := service.DB.Begin()

	if err := tx.Model(&InterfaceTestPartEntity.TInterfaceInfo{}).Where("interface_id = ?", interfaceInfo.InterfaceId).Update(newItfInfo).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, headerInfo := range interfaceInfo.Headers {
		newHeaderInfo := &InterfaceTestPartEntity.TInterfaceHeadersInfo{
			HeaderId:    headerInfo.HeaderId,
			InterfaceId: interfaceInfo.InterfaceId,
			HeaderName:  headerInfo.HeaderName,
			HeaderValue: headerInfo.HeaderValue,
		}
		if err := tx.Model(&InterfaceTestPartEntity.TInterfaceHeadersInfo{}).Where("interface_id = ?", interfaceInfo.InterfaceId).Count(&total).Error; err != nil {
			return nil
		} else {
			if total == 0 {
				newHeaderInfo.HeaderId = utils.GenerateUUID()
				if err := tx.Create(newHeaderInfo).Error; err != nil {
					fmt.Println(err)
					tx.Rollback()
					return nil
				}
			} else {
				if err := tx.Model(&InterfaceTestPartEntity.TInterfaceHeadersInfo{}).Where("header_id = ?", headerInfo.HeaderId).Update(newHeaderInfo).Error; err != nil {
					fmt.Println(err)
					tx.Rollback()
					return nil
				}
			}
		}

	}

	tx.Commit()
	return nil
}

func (its *itfTestService) DelInterface(q *qjson.QJson) error {
	itfId := q.GetNum("id")
	if err := service.DB.Model(&InterfaceTestPartEntity.TInterfaceInfo{}).Where("id = ?", itfId).Delete(&InterfaceTestPartEntity.TInterfaceInfo{}).Error; err != nil {
		return err
	}
	return nil
}
