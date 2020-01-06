package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/vo"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddCaseStep(c *gin.Context) {
	var caseStepInfo vo.CaseStepInfoVO
	c.ShouldBind(&caseStepInfo)
	if err := InterfaceTestPartSrv.CaseStepSrv.AddCaseStep(&caseStepInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
	}
}

func GetCaseStepDetail(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if stepDetail, err := InterfaceTestPartSrv.CaseStepSrv.GetCaseStepDetail(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", stepDetail)
	}
}

func EditCaseStep(c *gin.Context) {
	var caseStepInfo vo.CaseStepInfoVO
	c.ShouldBind(&caseStepInfo)
	//reqInfo := qjson.QJson{
	//	ReqInfo: utils.GetJsonBody(c),
	//}
	if err := InterfaceTestPartSrv.CaseStepSrv.EditCaseStep(&caseStepInfo); err != nil {
		utils.ResponseOkWithMsg(c, "修改失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "修改成功", nil)
	}

}

func DelCaseStep(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if err := InterfaceTestPartSrv.CaseStepSrv.DelCase(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "删除失败", nil)
	} else {
		utils.ResponseOkWithMsg(c, "删除成功", nil)
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
