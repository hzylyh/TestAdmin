package InterfaceTestPartEntity

import "salotto/conf"

type TInterfaceHeadersInfo struct {
	conf.Model
	HeaderId    string `json:"headerId"`
	InterfaceId string `json:"interfaceId"`
	HeaderName  string `json:"headerName"`
	HeaderValue string `json:"headerValue"`
}
