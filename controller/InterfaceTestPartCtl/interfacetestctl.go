package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddInterface(c *gin.Context) {
	var interfaceInfo InterfaceTestPartEntity.InterfaceInfo
	c.ShouldBind(&interfaceInfo)
	InterfaceTestPartSrv.ItfTestSrv.AddInterface(&interfaceInfo)
	utils.ResponseOkWithMsg(c, "新增成功", nil)
}

func GetInterfaceList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ItfTestSrv.GetInterfaceList(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}

func ImportSwagger(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	InterfaceTestPartSrv.ItfTestSrv.ImportSwagger(&reqInfo)
	utils.ResponseOkWithMsg(c, "导入成功", nil)
}

func GetInterfaceSelectOptions(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ItfTestSrv.GetInterfaceSelectOptions(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}
