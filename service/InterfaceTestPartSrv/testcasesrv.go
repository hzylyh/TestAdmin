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

func (tcs testCaseService) AddCase(itfCaseInfo *InterfaceTestPartEntity.ItfCaseInfo) error {
	itfCaseInfo.CaseId = utils.GenerateUUID()
	if err := service.DB.Create(itfCaseInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tcs testCaseService) GetCaseList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.ItfCaseInfo
	)

	//if pageInfo, err = utils.Pagination(&ret, qj); err != nil {
	if pageInfo, err = utils.PaginationWithDB(service.DB.Where(map[string]interface{}{"module_id": qj.GetString("moduleId")}), &ret, qj); err != nil {
		return nil, err
	} else {
		return pageInfo, nil
	}
}

func (tcs testCaseService) GetCaseTree(qj *qjson.QJson) (caseInfos []map[string]interface{}, err error) {
	var (
		ret []*InterfaceTestPartEntity.ItfCaseInfo
	)

	//if pageInfo, err = utils.Pagination(&ret, qj); err != nil {
	if err = service.DB.Where(map[string]interface{}{"module_id": qj.GetString("moduleId")}).Find(&ret).Error; err != nil {
		return nil, err
	}
	for _, eachRes := range ret {
		eachMap := map[string]interface{}{
			"id":    eachRes.CaseId,
			"value": eachRes.CaseId,
			"label": eachRes.CaseName,
		}
		caseInfos = append(caseInfos, eachMap)
	}
	return caseInfos, nil
}

func (tcs testCaseService) RunCase() error {
	exp := `{"name": "houzheyu", "age": 33, "list": ["a", "b"]}`
	act := `{"name": "houzheyu", "age": 33}`
	verify := []string{"name", "age", "list"}
	utils.MulAssert(exp, act, verify)
	return nil
}
