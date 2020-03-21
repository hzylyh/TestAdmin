package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"salotto/controller"
	"salotto/controller/InterfaceTestPartCtl"
	"salotto/controller/SchedulePartCtl"
	"salotto/middleware"
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
		itfManageApi     *gin.RouterGroup
		projectModuleApi *gin.RouterGroup
		itfCaseStepApi   *gin.RouterGroup
	)
	router = gin.New()

	// 日志中间件
	router.Use(middleware.LoggerToFile())

	router.Use(gin.Recovery())

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
	apiRoot.Use()

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

	// 用例部分
	itfTestCaseApi = interfaceTestApi.Group("/case")
	itfTestCaseApi.POST("/run", InterfaceTestPartCtl.RunCase)
	itfTestCaseApi.POST("/add", InterfaceTestPartCtl.AddCase)
	itfTestCaseApi.POST("/get", InterfaceTestPartCtl.GetCase)
	itfTestCaseApi.POST("/edit", InterfaceTestPartCtl.EditCase)
	itfTestCaseApi.POST("/delete", InterfaceTestPartCtl.DelCase)
	itfTestCaseApi.POST("/getList", InterfaceTestPartCtl.GetCaseList)
	itfTestCaseApi.POST("/getTree", InterfaceTestPartCtl.GetCaseTree)
	itfTestCaseApi.POST("/test", InterfaceTestPartCtl.TestSth)

	// 用例步骤
	itfCaseStepApi = itfTestCaseApi.Group("/step")
	itfCaseStepApi.POST("/add", InterfaceTestPartCtl.AddCaseStep)
	itfCaseStepApi.POST("/get", InterfaceTestPartCtl.GetCaseStepDetail)
	itfCaseStepApi.POST("/edit", InterfaceTestPartCtl.EditCaseStep)
	itfCaseStepApi.POST("/delete", InterfaceTestPartCtl.DelCaseStep)
	itfCaseStepApi.POST("/getList", InterfaceTestPartCtl.GetStepList)

	// 模块管理部分
	projectModuleApi = interfaceTestApi.Group("/module")
	projectModuleApi.POST("/add", InterfaceTestPartCtl.AddProjectModule)
	projectModuleApi.POST("/getList", InterfaceTestPartCtl.GetProjectModuleList)

	// 接口管理部分
	itfManageApi = interfaceTestApi.Group("/interface")
	itfManageApi.POST("/add", InterfaceTestPartCtl.AddInterface)
	itfManageApi.POST("/edit", InterfaceTestPartCtl.EditInterface)
	itfManageApi.POST("/get", InterfaceTestPartCtl.GetSingleInterfaceInfo)
	itfManageApi.POST("/delete", InterfaceTestPartCtl.DelInterface)
	itfManageApi.POST("/getList", InterfaceTestPartCtl.GetInterfaceList)
	itfManageApi.POST("/getSelectOptions", InterfaceTestPartCtl.GetInterfaceSelectOptions)
	itfManageApi.POST("/importSwagger", InterfaceTestPartCtl.ImportSwagger)

	return router
}
