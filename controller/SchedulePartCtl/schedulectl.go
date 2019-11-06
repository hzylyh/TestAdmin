package SchedulePartCtl

import (
	"github.com/gin-gonic/gin"
	"salotto/conf"
	"salotto/service/SchedulePartSrv"
	"salotto/utils"
	"salotto/utils/cronUtil"
)

func RunSchedule(c *gin.Context) {
	mj := cronUtil.New(1)
	go mj.RunJob()
	utils.ResponseOk(c, "duile")
}

func StopSchedule(c *gin.Context) {
	job := conf.JobList["a"]
	job.Stop()
	utils.ResponseOk(c, "duile")
}

func GetScheduleList(c *gin.Context) {
	res := SchedulePartSrv.Schedule.GetScheduleList()
	utils.ResponseOk(c, res)
}
