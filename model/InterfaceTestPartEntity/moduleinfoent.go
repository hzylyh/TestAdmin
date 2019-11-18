package InterfaceTestPartEntity

import "salotto/conf"

type ProjectModule struct {
	conf.Model
	ProjectId  int    `json:"projectId"`
	ModuleNum  string `json:"moduleNum"`
	ModuleName string `json:"moduleName"`
	ModuleDesc string `json:"moduleDesc"`
}
