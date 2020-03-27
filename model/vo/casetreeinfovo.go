package vo

type CaseTreeInfoVO struct {
	ProjectId string            `json:"projectId"`
	NodeId    string            `json:"nodeId"`
	PNodeId   string            `json:"pNodeId"`
	NodeName  string            `json:"nodeName"`
	NodeDesc  string            `json:"nodeDesc"`
	Children  []*CaseTreeInfoVO `json:"children"`
}
