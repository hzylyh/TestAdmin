package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type TNodeInfo struct {
	conf.Model
	ProjectId string
	NodeId    string
	PNodeId   string
	NodeNum   string
	NodeName  string
	NodeType  string
	NodeDesc  string
}
