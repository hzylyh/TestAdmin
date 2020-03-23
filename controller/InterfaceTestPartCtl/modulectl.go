package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddProjectModule(c *gin.Context) {
	var projectModuleInfo InterfaceTestPartEntity.TProjectModule
	c.ShouldBind(&projectModuleInfo)
	if err := InterfaceTestPartSrv.ModuleSrv.AddProjectModule(&projectModuleInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
	}
}

func DelProjectModule(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if err := InterfaceTestPartSrv.ModuleSrv.DelProjectModule(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "删除失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "删除成功", nil)
	}
}

func GetProjectModuleList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.ModuleSrv.GetModuleList(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "查询成功", res)
	}
}
