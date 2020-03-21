package InterfaceTestPartSrv

import (
	"database/sql"
	"errors"
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
		rows *sql.Rows
		err  error
	)
	moduleMap := qj.GetMap("cases")
	// 暂时不对module进行排序
	for k, v := range moduleMap {
		fmt.Println(k)
		caseIdList, ok := v.([]interface{})
		if !ok {
			return errors.New("转化错误")
		}

		// 如果前台传过来的module对应的caseList为空，则根据moduleId查询case，全部执行
		if len(caseIdList) == 0 {
			if rows, err = service.DB.Model(&InterfaceTestPartEntity.TItfCaseInfo{}).Select("case_id").Where("module_id = ?", k).Rows(); err != nil {
				fmt.Println(err)
				return err
			}
			defer rows.Close()
			for rows.Next() {
				var caseId string
				rows.Scan(&caseId)
				caseIdList = append(caseIdList, caseId)
			}
		}

		for _, caseId := range caseIdList {
			runCase(caseId.(string))
		}
	}

	return nil
}

func (tcs testCaseService) DelCase(q *qjson.QJson) error {
	caseId := q.GetString("caseId")
	// 目前为物理删除，且只删除了用例步骤表
	if err := service.DB.Where("case_id = ?", caseId).Delete(&InterfaceTestPartEntity.TItfCaseInfo{}).Error; err != nil {
		return err
	}
	return nil
}

func (tcs testCaseService) GetCase(q *qjson.QJson) (InterfaceTestPartEntity.TItfCaseInfo, error) {
	var (
		caseInfo = InterfaceTestPartEntity.TItfCaseInfo{}
		err      error
	)
	caseId := q.GetString("caseId")
	// 目前为物理删除，且只删除了用例步骤表
	if err = service.DB.Where("case_id = ?", caseId).First(&caseInfo).Error; err != nil {
		fmt.Println(err)
	}
	return caseInfo, nil
}

func (tcs testCaseService) EditCase(itfCaseInfo *InterfaceTestPartEntity.TItfCaseInfo) error {
	if err := service.DB.Model(InterfaceTestPartEntity.TItfCaseInfo{}).Where("case_id = ?", itfCaseInfo.CaseId).Update(itfCaseInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func runCase(caseId string) {
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
	service.DB.Where(map[string]interface{}{"case_id": caseId}).Order("step_num").Find(&stepInfos)
	myRequests := requests.NewRequests()
	// 挨个步骤请求
	for _, stepInfo := range stepInfos {
		var (
			itfInfo        InterfaceTestPartEntity.TInterfaceInfo
			itfHeaderInfos []*InterfaceTestPartEntity.TInterfaceHeadersInfo
			variables      []InterfaceTestPartEntity.TCaseStepVarInfo
			assertInfos    []InterfaceTestPartEntity.TAssertInfo
		)
		replaceReqData := parser.Parse(stepInfo.ReqData) // 请求体变量替换，将请求种的${xxx}替换
		service.DB.Where("interface_id = ?", stepInfo.ItfId).First(&itfInfo)

		// 查询请求头信息，进行替换
		if err := service.DB.Where("interface_id = ?", stepInfo.ItfId).Find(&itfHeaderInfos).Error; err != nil {
			return
		} else {
			for _, headerInfo := range itfHeaderInfos {
				headerInfo.HeaderValue = parser.Parse(headerInfo.HeaderValue)
			}
		}

		act := string(myRequests.Post(itfInfo.Url, replaceReqData, itfHeaderInfos)) // 发送请求
		// 根据用例步骤id获取步骤变量
		service.DB.Where("step_id = ?", stepInfo.StepId).Find(&variables)
		for _, eachStepVar := range variables {
			collectCol := eachStepVar.CollectCol
			collectColAlias := eachStepVar.CollectColAlias
			properties[collectCol] = gjson.Get(act, collectCol).String()      // 将变量字段作key，存入map
			properties[collectColAlias] = gjson.Get(act, collectCol).String() // 将变量别名作key，存入map，后续考虑分两个map
		}
		//verify := []string{"resultEntity.textToSpeechContent", "resultMessage", "resultEntity.timeoutInstructionList"}
		//hassert.MulAssert(stepInfo.ExpRes, act, verify)
		service.DB.Where("step_id = ?", stepInfo.StepId).Find(&assertInfos)
		hassert.MulAssertNew(act, assertInfos, stepInfo)
	}
	//exp := `{"name": "houzheyu", "age": 33, "list": ["a", "b"]}`
	//act := `{"name": "houzheyu", "age": 33}`
	//verify := []string{"name", "age", "list"}
}
