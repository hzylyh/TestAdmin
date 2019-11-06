package cronUtil

import (
	"github.com/robfig/cron"
)

func OneCronMulCommand(cronExp string, cmdList []string, parallel int) *cron.Cron {
	myCh := make(chan int, parallel)

	c := cron.New()

	c.AddFunc(cronExp, func() {
		for i, cmd := range cmdList {
			myCh <- i
			go ShellExecutor(cmd, myCh)
			//<- myCh
		}
	})
	c.Start()

	return c
}

func Run(jobId int) *cron.Cron {
	var (
		cron     string
		commands []string
	)
	row := CronDB.Table("schedules").Where("id = ?", jobId).Select("cron").Row()
	row.Scan(&cron)

	rows, _ := CronDB.Table("schedule_jobs").Where("job_id = ?", jobId).Select("command").Rows()
	defer rows.Close()
	for rows.Next() {
		var v string
		rows.Scan(&v)
		commands = append(commands, v)
	}

	c := OneCronMulCommand(cron, commands, 3)

	return c
}
