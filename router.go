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
		itfTestCaseApi   *gin.RouterGroup
		overviewApi      *gin.RouterGroup
	)
	router = gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	router.LoadHTMLGlob("web/dist/*.html")          // 添加入口index.html
	router.LoadHTMLFiles("web/static/*/*")          // 添加资源路径
	router.Static("/static", "./web/dist/static")   // 添加资源路径
	router.StaticFile("/", "./web/dist/index.html") //前端接口

	// 创建根路由
	apiRoot = router.Group("/api")

	authApi = apiRoot.Group("/user")
	authApi.POST("/login", controller.Login)

	// 总览部分
	overviewApi = apiRoot.Group("/overview")
	overviewApi.POST("/getProjectList", controller.GetProjectList)
	overviewApi.POST("/addProject", controller.AddProject)

	// 定时任务部分
	scheduleApi = apiRoot.Group("/schedule")
	scheduleApi.POST("/run", SchedulePartCtl.RunSchedule)
	scheduleApi.POST("/stop", SchedulePartCtl.StopSchedule)
	scheduleApi.POST("/getScheduleList", SchedulePartCtl.GetScheduleList)

	// 接口测试部分
	interfaceTestApi = apiRoot.Group("/itfPart")
	interfaceTestApi.POST("/test", InterfaceTestPartCtl.Test)
	// 用例部分
	itfTestCaseApi = interfaceTestApi.Group("/case")
	itfTestCaseApi.POST("/run", InterfaceTestPartCtl.RunCase)

	return router
}
