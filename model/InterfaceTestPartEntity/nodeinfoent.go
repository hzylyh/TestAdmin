package InterfaceTestPartEntity

import (
	"salotto/conf"
)

type TNodeInfo struct {
	conf.Model
	ProjectId string `json:"projectId"`
	NodeId    string `json:"nodeId"`
	PNodeId   string `json:"pNodeId"`
	NodeNum   string `json:"nodeNum"`
	NodeName  string `json:"nodeName"`
	NodeType  string `json:"nodeType"`
	NodeDesc  string `json:"nodeDesc"`
}
