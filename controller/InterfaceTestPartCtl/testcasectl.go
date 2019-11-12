package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/service/InterfaceTestPartSrv"
)

func RunCase(c *gin.Context) {
	InterfaceTestPartSrv.TestCaseSrv.RunCase()
}
