package cronUtil

import (
	"database/sql"
	"fmt"
	"github.com/robfig/cron"
)

func OneCronMulCommand(jobId int, parallel int) *cron.Cron {
	var (
		myCh      chan int
		waitChan  chan TaskInfo
		myCron    *cron.Cron
		newJob    *ScheduleJob
		cronExp   string
		taskInfos []TaskInfo
	)

	cronExp = getCron(jobId)
	taskInfos = getTaskInfos(jobId)

	myCh = make(chan int, parallel)
	waitChan = make(chan TaskInfo, len(taskInfos))
	myCron = cron.New()

	newJob = &ScheduleJob{
		taskChan: myCh,
		waitChan: waitChan,
		//cmdList:  cmdList,
		taskInfos: taskInfos,
	}

	myCron.AddJob(cronExp, newJob)
	myCron.Start()

	return myCron
}

func getCron(jobId int) string {
	var (
		cron string
		row  *sql.Row
	)
	row = CronDB.Table("schedules").Where("id = ?", jobId).Select("cron").Row()
	row.Scan(&cron)
	return cron
}

func getTaskInfos(jobId int) (taskInfos []TaskInfo) {
	var (
		//commands []string
		rows *sql.Rows
		err  error
	)

	if rows, err = CronDB.Table("schedule_jobs").Where("job_id = ?", jobId).Select("task_id, command, depends").Rows(); err != nil {
		fmt.Println("查询任务出错，请排查", err)
	}
	defer rows.Close()
	for rows.Next() {
		//var v string
		//rows.Scan(&v)
		//commands = append(commands, v)
		var taskInfo TaskInfo
		CronDB.ScanRows(rows, &taskInfo)
		taskInfos = append(taskInfos, taskInfo)
	}
	return taskInfos
}

func CheckDepends(taskIds []int) bool {
	var (
		status int
		row    *sql.Row
	)
	for _, taskId := range taskIds {
		row = CronDB.Table("schedule_jobs").Where("task_id = ?", taskId).Select("status").Row()
		row.Scan(&status)
		if status != 0 {
			return false
		}
	}

	return true
}

func Run(jobId int) *cron.Cron {

	c := OneCronMulCommand(jobId, 1)

	return c
}
