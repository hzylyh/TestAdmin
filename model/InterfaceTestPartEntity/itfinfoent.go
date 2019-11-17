package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type InterfaceInfo struct {
	conf.Model
	Name string
	Url  string
	Type string
	Desc string
}
