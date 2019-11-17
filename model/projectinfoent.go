package model

type ProjectInfo struct {
	Model
	ProjectName string `json:"project_name"`
	ProjectDesc string `json:"project_desc"`
	Creator     string `json:"creator"`
	Modifier    string `json:"modifier"`
}
