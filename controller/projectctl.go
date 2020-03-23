package controller

import (
	"github.com/gin-gonic/gin"
	"salotto/model"
	"salotto/service"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddProject(c *gin.Context) {
	var projectInfo model.TProjectInfo
	c.ShouldBind(&projectInfo)
	if projectId, err := service.ProjectSrv.AddProject(&projectInfo); err != nil {
		utils.ResponseOkWithMsg(c, "新增失败", err)
	} else {
		res := make(map[string]string)
		res["projectId"] = projectId
		utils.ResponseOkWithMsg(c, "新增成功", res)
	}

}

func GetProjectList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	res := service.ProjectSrv.GetProjectList(&reqInfo)
	utils.ResponseOk(c, res)
}
