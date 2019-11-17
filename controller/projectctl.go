package controller

import (
	"github.com/gin-gonic/gin"
	"salotto/model"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddProject(c *gin.Context) {
	var projectInfo model.ProjectInfo
	c.ShouldBind(&projectInfo)
	service.ProjectSrv.AddProject(&projectInfo)
	utils.ResponseOk(c, "新增成功")
}

func GetProjectList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	res := service.ProjectSrv.GetProjectList(&reqInfo)
	utils.ResponseOk(c, res)
}
