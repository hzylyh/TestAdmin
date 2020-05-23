package InterfaceTestPartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/service/InterfaceTestPartSrv"
	"salotto/utils"
	"salotto/utils/qjson"
)

/**
获取用例运行情况数据
*/
func GetCaseRunInfo(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if res, err := InterfaceTestPartSrv.DashboardSrv.GetCaseRunInfo(&reqInfo); err != nil {
		utils.ResponseOkWithMsg(c, "查询失败", nil)
	} else {
		utils.ResponseOk(c, res)
	}
}
