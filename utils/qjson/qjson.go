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
