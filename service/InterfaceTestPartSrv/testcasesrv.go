package InterfaceTestPartSrv

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"salotto/model"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/model/vo"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/hassert"
	"salotto/utils/parse"
	"salotto/utils/qjson"
	"salotto/utils/requests"
	"salotto/utils/stime"
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
		beginTime     string
		transNodeInfo InterfaceTestPartEntity.TNodeInfo
	)
	nodeList := qj.GetArray("nodeList")
	beginTime = stime.GetCurTime()

	for _, nodeInfo := range nodeList {
		mapstructure.Decode(nodeInfo, &transNodeInfo)
		if transNodeInfo.NodeType == "用例" {
			fmt.Println(transNodeInfo)
			runCase(transNodeInfo.NodeId, beginTime)
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

func (tcs testCaseService) GetCaseTreeNode(q *qjson.QJson) (nodeInfo *vo.CaseTreeInfoVO, err error) {
	projectId := q.GetString("projectId")
	var (
		allNodeInfo   []*InterfaceTestPartEntity.TNodeInfo
		allNodeInfoVO []*vo.CaseTreeInfoVO
		resNodeInfo   InterfaceTestPartEntity.TNodeInfo
		resNodeInfoVO vo.CaseTreeInfoVO
	)
	if err := service.DB.Model(InterfaceTestPartEntity.TItfCaseInfo{}).Where("project_id = ?", projectId).Find(&allNodeInfo).Error; err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		for _, nodeInfo := range allNodeInfo {
			newNode := &vo.CaseTreeInfoVO{
				ProjectId: nodeInfo.ProjectId,
				NodeId:    nodeInfo.NodeId,
				PNodeId:   nodeInfo.PNodeId,
				NodeName:  nodeInfo.NodeName,
				NodeType:  nodeInfo.NodeType,
				NodeDesc:  nodeInfo.NodeDesc,
				Children:  nil,
			}
			allNodeInfoVO = append(allNodeInfoVO, newNode)
		}
		service.DB.Where("project_id = ? and p_node_id = ?", projectId, "-1").First(&resNodeInfo)
		resNodeInfoVO.ProjectId = projectId
		resNodeInfoVO.NodeId = resNodeInfo.NodeId
		resNodeInfoVO.PNodeId = resNodeInfo.PNodeId
		resNodeInfoVO.NodeName = resNodeInfo.NodeName
		makeTree(allNodeInfoVO, &resNodeInfoVO)
	}

	return &resNodeInfoVO, nil
}

func (tcs testCaseService) InitTree(q *qjson.QJson) error {
	projectId := q.GetString("projectId")
	projectName := q.GetString("projectName")
	projectTreeRoot := InterfaceTestPartEntity.TNodeInfo{
		ProjectId: projectId,
		NodeId:    "0",
		PNodeId:   "-1",
		NodeNum:   "",
		NodeName:  projectName,
		NodeType:  "fold",
		NodeDesc:  "",
	}
	if err := service.DB.Create(&projectTreeRoot).Error; err != nil {
		return err
	}
	return nil
}

func (tcs testCaseService) AddNode(ct *vo.CaseTreeInfoVO) error {
	fmt.Println(ct)
	newNodeInfo := InterfaceTestPartEntity.TNodeInfo{
		ProjectId: ct.ProjectId,
		NodeId:    utils.GenerateUUID(),
		PNodeId:   ct.PNodeId,
		NodeNum:   "",
		NodeName:  ct.NodeName,
		NodeType:  ct.NodeType,
		NodeDesc:  ct.NodeDesc,
	}
	if err := service.DB.Create(&newNodeInfo).Error; err != nil {
		return err
	}
	return nil
}

func (tcs testCaseService) DelNode(q *qjson.QJson) error {
	nodeId := q.GetString("nodeId")
	if err := service.DB.Where("node_id = ?", nodeId).Delete(&InterfaceTestPartEntity.TNodeInfo{}).Error; err != nil {
		return err
	}
	return nil
}

func (tcs testCaseService) EditNode(v *vo.CaseTreeInfoVO) error {
	newNodeInfo := InterfaceTestPartEntity.TNodeInfo{
		ProjectId: v.ProjectId,
		NodeId:    v.NodeId,
		NodeName:  v.NodeName,
		NodeDesc:  v.NodeDesc,
	}
	if err := service.DB.Model(InterfaceTestPartEntity.TNodeInfo{}).Where("node_id = ?", v.NodeId).Update(&newNodeInfo).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tcs testCaseService) GetSingleNodeInfo(q *qjson.QJson) (*InterfaceTestPartEntity.TNodeInfo, error) {
	var (
		nodeId string
		ret    InterfaceTestPartEntity.TNodeInfo
	)
	nodeId = q.GetString("nodeId")

	if err := service.DB.Where("node_id = ?", nodeId).First(&ret).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &ret, nil
}

func runCase(caseId, beginTime string) {
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
		hassert.MulAssertNew(act, assertInfos, stepInfo, beginTime)
	}
	//exp := `{"name": "houzheyu", "age": 33, "list": ["a", "b"]}`
	//act := `{"name": "houzheyu", "age": 33}`
	//verify := []string{"name", "age", "list"}
}

func makeTree(allNodeInfo []*vo.CaseTreeInfoVO, nodeInfo *vo.CaseTreeInfoVO) {
	var (
		children []*vo.CaseTreeInfoVO
	)
	children = haveChild(allNodeInfo, nodeInfo)
	if len(children) > 0 {
		nodeInfo.Children = append(nodeInfo.Children, children...)
		for _, v := range children {
			innerChildren := haveChild(allNodeInfo, v)
			if len(innerChildren) > 0 {
				makeTree(allNodeInfo, v)
			}
		}
	}
}

func haveChild(allNodeInfo []*vo.CaseTreeInfoVO, nodeInfo *vo.CaseTreeInfoVO) (children []*vo.CaseTreeInfoVO) {
	for _, nodeInfoInAll := range allNodeInfo {
		if nodeInfo.NodeId == nodeInfoInAll.PNodeId {
			children = append(children, nodeInfoInAll)
		}
	}
	return children
}
