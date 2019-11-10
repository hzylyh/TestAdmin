package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"salotto/controller"
	"salotto/controller/InterfaceTestPartCtl"
	"salotto/controller/SchedulePartCtl"
)

/**
路由控制
*/
func MapRoutes() *gin.Engine {
	var (
		router           *gin.Engine
		apiRoot          *gin.RouterGroup
		authApi          *gin.RouterGroup
		scheduleApi      *gin.RouterGroup
		interfaceTestApi *gin.RouterGroup
	)
	router = gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	// 创建根路由
	apiRoot = router.Group("/salotto")

	authApi = apiRoot.Group("/user")
	authApi.POST("/login", controller.Login)

	// 定时任务部分
	scheduleApi = apiRoot.Group("/schedule")
	scheduleApi.POST("/run", SchedulePartCtl.RunSchedule)
	scheduleApi.POST("/stop", SchedulePartCtl.StopSchedule)
	scheduleApi.POST("/getScheduleList", SchedulePartCtl.GetScheduleList)

	// 接口测试部分
	interfaceTestApi = apiRoot.Group("/itfPart")
	interfaceTestApi.POST("/test", InterfaceTestPartCtl.Test)

	return router
}
