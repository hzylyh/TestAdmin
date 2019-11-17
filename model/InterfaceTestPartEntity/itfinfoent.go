package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type InterfaceInfo struct {
	conf.Model
	ProjectId int    `json:"projectId"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Type      string `json:"type"`
	Desc      string `json:"desc"`
}
