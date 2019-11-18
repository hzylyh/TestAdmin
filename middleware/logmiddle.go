package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s\t|\t%s|\t%s|\t%s|\t%s|\t%d|\t%s|\t%s|\t%s \n",
			//客户端IP
			param.ClientIP,
			//时间格式
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			//http请求方式 get post等
			param.Method,
			//客户端请求的路径
			param.Path,
			//http请求协议版本
			param.Request.Proto,
			//http请求状态码
			param.StatusCode,
			//耗时
			param.Latency,
			//http请求代理头
			param.Request.UserAgent(),
			//处理请求错误时设置错误消息
			param.ErrorMessage,
		)
	})
}
