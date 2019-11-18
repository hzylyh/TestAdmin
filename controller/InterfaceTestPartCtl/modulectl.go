package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
)

func AddProjectModule(c *gin.Context) {
	var projectModuleInfo InterfaceTestPartEntity.ProjectModule
	c.ShouldBind(&projectModuleInfo)
	if err := InterfaceTestPartSrv.ModuleSrv.AddProjectModule(&projectModuleInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", err.Error())
	} else {
		utils.ResponseOkWithMsg(c, "新增成功", nil)
	}
}
