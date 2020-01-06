package model

import "salotto/conf"

type TProjectInfo struct {
	conf.Model
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	Creator     string `json:"creator"`
	Modifier    string `json:"modifier"`
}
