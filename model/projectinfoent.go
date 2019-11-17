package model

import "salotto/conf"

type ProjectInfo struct {
	conf.Model
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	Creator     string `json:"creator"`
	Modifier    string `json:"modifier"`
}
