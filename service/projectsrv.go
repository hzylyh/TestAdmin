package service

import (
	"fmt"
	"salotto/model"
	"salotto/utils"
	"salotto/utils/qjson"
)

var ProjectSrv = &projectService{}

type projectService struct {
}

func (ps *projectService) AddProject(projectInfo *model.TProjectInfo) (string, error) {
	projectInfo.ProjectId = utils.GenerateUUID()
	if err := DB.Create(projectInfo).Error; err != nil {
		fmt.Println(err)
		return "", err
	}
	return projectInfo.ProjectId, nil
}

func (ps *projectService) GetProjectList(qj *qjson.QJson) []*model.TProjectInfo {
	var ret []*model.TProjectInfo
	if err := DB.Order("created_at desc").Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}

func (ps *projectService) DelProject(q *qjson.QJson) error {
	projectId := q.GetString("projectId")
	if err := DB.Where("project_id = ?", projectId).Delete(&model.TProjectInfo{}).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
