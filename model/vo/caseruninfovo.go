package vo

/**

 */
type CaseRunInfoVO struct {
	BeginTime []string `json:"beginTime"`
	Success   []int64  `json:"success"`
	Fail      []int64  `json:"fail"`
	Total     []int64  `json:"total"`
}
