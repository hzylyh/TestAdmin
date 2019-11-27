package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func RunCase(c *gin.Context) {
	InterfaceTestPartSrv.TestCaseSrv.RunCase()
}

func AddCase(c *gin.Context) {
	var itfTestCaseInfo InterfaceTestPartEntity.ItfCaseInfo
	c.ShouldBind(&itfTestCaseInfo)
	if err := InterfaceTestPartSrv.TestCaseSrv.AddCase(&itfTestCaseInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
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
