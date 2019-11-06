package cronUtil

import (
	"fmt"
	"github.com/robfig/cron"
	"os/exec"
	"salotto/conf"
)

//type CronJob interface {
//	RunJob()
//}

type MyJob struct {
	JobId int
}

func New(jobId int) *MyJob {
	return &MyJob{
		JobId: jobId,
	}
}

func (mj *MyJob) RunJob() {
	c := Run(mj.JobId)
	conf.JobList["a"] = c
}

func (mj *MyJob) StopJob(c *cron.Cron) {
	fmt.Println(mj.JobId)
}

func ShellExecutor(cmd string, ch chan int) (err error) {
	var (
		res    *exec.Cmd
		output []byte
	)
	res = exec.Command("G:/cygwin64/bin/bash.exe", "-c", cmd)
	output, err = res.Output()
	<-ch
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	return
}
