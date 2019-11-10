package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/utils"
)

func Test(c *gin.Context) {
	utils.ResponseOk(c, "hahaha")
}
