package vo

type CaseTreeInfoVO struct {
	ProjectId string            `json:"projectId"`
	NodeId    string            `json:"nodeId"`
	PNodeId   string            `json:"pNodeId"`
	NodeName  string            `json:"nodeName"`
	NodeType  string            `json:"nodeType"`
	NodeDesc  string            `json:"nodeDesc"`
	Children  []*CaseTreeInfoVO `json:"children"`
}
