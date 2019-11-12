package InterfaceTestPartCtl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"salotto/utils"
)

func Test(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))
	utils.ResponseOk(c, "hahaha")
}
