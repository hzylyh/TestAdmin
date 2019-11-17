package InterfaceTestPartCtl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"salotto/model/InterfaceTestPartEntity"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

func AddInterface(c *gin.Context) {
	var interfaceInfo InterfaceTestPartEntity.InterfaceInfo
	c.ShouldBind(&interfaceInfo)
	InterfaceTestPartSrv.ItfTestSrv.AddInterface(&interfaceInfo)
	utils.ResponseOkWithMsg(c, "新增成功", nil)
}

func GetInterfaceList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	fmt.Println(reqInfo)
	res := InterfaceTestPartSrv.ItfTestSrv.GetInterfaceList(&reqInfo)
	utils.ResponseOk(c, res)
	//utils.ResponseOk(c, )
}
