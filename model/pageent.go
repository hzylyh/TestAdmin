package model

type PageInfo struct {
	PageNum  float64     `json:"pageNum"`
	PageSize float64     `json:"pageSize"`
	Total    float64     `json:"total"`
	List     interface{} `json:"list"`
}
