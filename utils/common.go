package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetJsonBody(c *gin.Context) map[string]interface{} {
	var (
		reqInfo map[string]interface{}
		body    []byte
		err     error
	)
	if body, err = ioutil.ReadAll(c.Request.Body); err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &reqInfo)
	return reqInfo
}
