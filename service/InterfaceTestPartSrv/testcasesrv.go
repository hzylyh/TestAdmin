package InterfaceTestPartSrv

import (
	"fmt"
	"github.com/tidwall/gjson"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/parse"
	"salotto/utils/qjson"
	"salotto/utils/requests"
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

func (tcs testCaseService) RunCase(qj *qjson.QJson) error {
	var (
		stepInfos  []InterfaceTestPartEntity.ItfCaseStepInfo
		properties = make(map[string]string)
		Ihandler   parse.TokenHandler
	)

	//var properties = map[string]string{"name": "houzheyu", "age": "2222"}
	Ihandler = &parse.VariableTokenHandler{
		Variables: properties,
	}
	parser := parse.GenericTokenParser{
		OpenToken:  "${",
		CloseToken: "}",
		Handler:    Ihandler,
	}

	// 查询对应用例步骤
	service.DB.Where(map[string]interface{}{"case_id": qj.GetString("caseId")}).Preload("Variables").Find(&stepInfos)

	// 挨个步骤请求
	for _, stepInfo := range stepInfos {
		replaceReqData := parser.Parse(stepInfo.ReqData) // 变量替换，将请求种的${xxx}替换
		fmt.Println(replaceReqData)
		act := string(requests.Post("http://localhost:8089/api/itfPart/case/test", ""))
		for _, eachStepVar := range stepInfo.Variables {
			collectCol := eachStepVar.CollectCol
			properties[collectCol] = gjson.Get(act, collectCol).String() // 将变量字段作key，存入map
			properties[collectCol] = gjson.Get(act, collectCol).String() // 将变量别名作key，存入map，后续考虑分两个map
			fmt.Println(properties)
		}
		//verify := []string{"name", "age", "list"}
		//utils.MulAssert(stepInfo.ExpRes, act, verify)
	}
	//exp := `{"name": "houzheyu", "age": 33, "list": ["a", "b"]}`
	//act := `{"name": "houzheyu", "age": 33}`
	//verify := []string{"name", "age", "list"}

	return nil
}
