package swagger

import (
	"encoding/json"
	"fmt"
	"salotto/utils/requests"
)

func main() {
	SwaggerParse("http://localhost:9879/v2/api-docs")
}

func SwaggerParse(url string) map[string]interface{} {
	var resMap map[string]interface{}
	res := requests.Post(url, "")
	if err := json.Unmarshal(res, &resMap); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resMap["paths"])
	if pathMap, ok := resMap["paths"].(map[string]interface{}); ok {
		return pathMap
	}
	return nil
}
