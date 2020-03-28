package InterfaceTestPartCtl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/model/vo"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func RunCase(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}

	// goroutine先在此控制
	go InterfaceTestPartSrv.TestCaseSrv.RunCase(&reqInfo)
	utils.ResponseOkWithMsg(c, "执行成功", nil)
}

func AddCase(c *gin.Context) {
	var itfTestCaseInfo InterfaceTestPartEntity.TItfCaseInfo
	c.ShouldBind(&itfTestCaseInfo)
	if err := InterfaceTestPartSrv.TestCaseSrv.AddCase(&itfTestCaseInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
	}

}

func EditCase(c *gin.Context) {
	var itfTestCaseInfo InterfaceTestPartEntity.TItfCaseInfo
	c.ShouldBind(&itfTestCaseInfo)
	if err := InterfaceTestPartSrv.TestCaseSrv.EditCase(&itfTestCaseInfo); err != nil {
		utils.ResponseOkWithMsg(c, "修改失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "修改成功", nil)
	}

}

func GetCase(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if caseInfo, err := InterfaceTestPartSrv.TestCaseSrv.GetCase(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", caseInfo)
	}
}

func DelCase(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if err := InterfaceTestPartSrv.TestCaseSrv.DelCase(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "删除失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "删除成功", nil)
	}
}

func GetCaseList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.TestCaseSrv.GetCaseList(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", res)
	}
}

func GetCaseTree(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.TestCaseSrv.GetCaseTree(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", res)
	}
}

func TestSth(c *gin.Context) {
	var resMap map[string]interface{}
	json.Unmarshal([]byte(`{"name": "houzheyu", "age": 33}`), &resMap)
	c.JSON(200, resMap)
}

func GetCaseTreeNode(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.TestCaseSrv.GetCaseTreeNode(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", res)
	}
}

func InitTree(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if err := InterfaceTestPartSrv.TestCaseSrv.InitTree(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "成功", nil)
	}
}

func AddNode(c *gin.Context) {
	var nodeInfoVO vo.CaseTreeInfoVO
	c.ShouldBind(&nodeInfoVO)
	if err := InterfaceTestPartSrv.TestCaseSrv.AddNode(&nodeInfoVO); err != nil {
		utils.ResponseOkWithMsg(c, "失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "成功", nil)
	}
}
