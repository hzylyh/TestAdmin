package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type TInterfaceInfo struct {
	conf.Model
	ProjectId   string `json:"projectId"`
	InterfaceId string `json:"interfaceId"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Type        string `json:"type"`
	Desc        string `json:"desc"`
}
