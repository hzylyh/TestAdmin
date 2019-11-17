package qjson

type QJson struct {
	ReqInfo map[string]interface{}
}

func (qj *QJson) GetString(key string) string {
	if res, ok := qj.ReqInfo[key].(string); ok {
		return res
	}
	return ""
}

func (qj *QJson) GetInt(key string) float64 {
	if res, ok := qj.ReqInfo[key].(float64); ok {
		return res
	}
	return 0
}
