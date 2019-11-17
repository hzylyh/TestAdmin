package service

import (
	"fmt"
	"salotto/model"
	"salotto/utils/qjson"
)

var ProjectSrv = &projectService{}

type projectService struct {
}

func (ps *projectService) AddProject(projectInfo *model.ProjectInfo) {
	if err := DB.Create(projectInfo); err != nil {
		fmt.Println(err)
	}

}

func (ps *projectService) GetProjectList(qj *qjson.QJson) []*model.ProjectInfo {
	var ret []*model.ProjectInfo
	if err := DB.Find(&ret); err != nil {
		fmt.Println(err)
	}
	return ret
}
