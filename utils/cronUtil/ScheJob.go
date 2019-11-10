package cronUtil

import "fmt"

type ScheduleJob struct {
	taskChan  chan int
	cmdList   []string
	taskInfos []TaskInfo
}

type TaskInfo struct {
	TaskId  int
	Command string
	Depends string
}

func (sj *ScheduleJob) Run() {
	//for i, cmd := range sj.cmdList {
	//	sj.taskChan <- i
	//	go ShellExecutor(cmd, sj.taskChan)
	//	//<- myCh
	//}

	fmt.Println(sj.taskInfos)
}
