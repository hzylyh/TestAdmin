package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"salotto/conf"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//func ResponseWithJson(w http.ResponseWriter, code int, payload interface{})  {
//	response, _ := json.Marshal(payload)
//	fmt.Println(string(response))
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//	w.Write(response)
//}

func ResponseOk(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Code: conf.OkStatus, Msg: conf.OkMessage, Data: data})
}

func ResponseOkWithMsg(c *gin.Context, msg string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Code: conf.OkStatus, Msg: msg, Data: data})
}

func ResponseFailOther(c *gin.Context, code, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Code: code, Msg: message})
}
