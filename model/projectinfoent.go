package model

type ProjectInfo struct {
	Model
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	Creator     string `json:"creator"`
	Modifier    string `json:"modifier"`
}
