package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddCaseStep(c *gin.Context) {
	var caseStepInfo InterfaceTestPartEntity.ItfCaseStepInfo
	c.ShouldBind(&caseStepInfo)
	if err := InterfaceTestPartSrv.CaseStepSrv.AddCaseStep(&caseStepInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
	}
}

func GetStepList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.CaseStepSrv.GetStepList(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", res)
	}
}
