package InterfaceTestPartSrv

import (
	"fmt"
	"github.com/tidwall/gjson"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/hassert"
	"salotto/utils/parse"
	"salotto/utils/qjson"
	"salotto/utils/requests"
)

var TestCaseSrv = &testCaseService{}

type testCaseService struct {
}

func (tcs testCaseService) AddCase(itfCaseInfo *InterfaceTestPartEntity.TItfCaseInfo) error {
	itfCaseInfo.CaseId = utils.GenerateUUID()
	if err := service.DB.Create(itfCaseInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tcs testCaseService) GetCaseList(qj *qjson.QJson) (pageInfo *model.PageInfo, err error) {
	var (
		ret []*InterfaceTestPartEntity.TItfCaseInfo
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
		ret []*InterfaceTestPartEntity.TItfCaseInfo
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
		stepInfos  []InterfaceTestPartEntity.TItfCaseStepInfo
		properties = make(map[string]string)
		Ihandler   parse.TokenHandler
	)

	Ihandler = &parse.VariableTokenHandler{
		Variables: properties,
	}
	parser := parse.GenericTokenParser{
		OpenToken:  "${",
		CloseToken: "}",
		Handler:    Ihandler,
	}

	// 查询对应用例步骤
	service.DB.Where(map[string]interface{}{"case_id": qj.GetString("caseId")}).Order("step_num").Find(&stepInfos)
	myRequests := requests.NewRequests()
	// 挨个步骤请求
	for _, stepInfo := range stepInfos {
		var (
			itfInfo   InterfaceTestPartEntity.TInterfaceInfo
			variables []InterfaceTestPartEntity.TCaseStepVarInfo
		)
		replaceReqData := parser.Parse(stepInfo.ReqData) // 变量替换，将请求种的${xxx}替换
		service.DB.Where("interface_id = ?", stepInfo.ItfId).First(&itfInfo)
		act := string(myRequests.Post(itfInfo.Url, replaceReqData))
		// 根据用例步骤id获取步骤变量
		service.DB.Where("step_id = ?", stepInfo.StepId).Find(&variables)
		for _, eachStepVar := range variables {
			collectCol := eachStepVar.CollectCol
			collectColAlias := eachStepVar.CollectColAlias
			properties[collectCol] = gjson.Get(act, collectCol).String()      // 将变量字段作key，存入map
			properties[collectColAlias] = gjson.Get(act, collectCol).String() // 将变量别名作key，存入map，后续考虑分两个map
		}
		verify := []string{"resultEntity.textToSpeechContent", "resultEntity.timeoutInstructionList"}
		hassert.MulAssert(stepInfo.ExpRes, act, verify)
	}
	//exp := `{"name": "houzheyu", "age": 33, "list": ["a", "b"]}`
	//act := `{"name": "houzheyu", "age": 33}`
	//verify := []string{"name", "age", "list"}

	return nil
}
