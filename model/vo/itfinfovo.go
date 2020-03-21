package vo

import "salotto/model/InterfaceTestPartEntity"

type InterfaceInfoVO struct {
	Id          string                                          `json:"id"`
	ProjectId   string                                          `json:"projectId"`
	InterfaceId string                                          `json:"interfaceId"`
	Name        string                                          `json:"name"`
	Url         string                                          `json:"url"`
	Type        string                                          `json:"type"`
	Headers     []InterfaceTestPartEntity.TInterfaceHeadersInfo `json:"headers"`
	Desc        string                                          `json:"desc"`
}
