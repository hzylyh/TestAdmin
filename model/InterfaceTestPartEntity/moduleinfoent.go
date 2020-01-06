package InterfaceTestPartEntity

import "salotto/conf"

type TProjectModule struct {
	conf.Model
	ProjectId  string `json:"projectId"`
	ModuleId   string `json:"moduleId"`
	PModuleId  string `json:"pModuleId"`
	ModuleNum  string `json:"moduleNum"`
	ModuleName string `json:"moduleName"`
	ModuleDesc string `json:"moduleDesc"`
}
