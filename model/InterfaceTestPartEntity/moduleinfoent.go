package InterfaceTestPartEntity

import "salotto/conf"

type ProjectModule struct {
	conf.Model
	ProjectId  int    `json:"projectId"`
	PModuleId  int    `json:"pModuleId"`
	ModuleNum  string `json:"moduleNum"`
	ModuleName string `json:"moduleName"`
	ModuleDesc string `json:"moduleDesc"`
}
