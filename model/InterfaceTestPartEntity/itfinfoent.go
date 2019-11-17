package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type InterfaceInfo struct {
	conf.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}
