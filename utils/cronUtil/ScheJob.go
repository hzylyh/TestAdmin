package cronUtil

import (
	"encoding/json"
	"fmt"
)

type ScheduleJob struct {
	taskChan  chan int
	waitChan  chan TaskInfo
	cmdList   []string
	taskInfos []TaskInfo
}

type TaskInfo struct {
	TaskId  int
	Command string
	Depends string
}

// 实现Run接口，自定义运行方法
func (sj *ScheduleJob) Run() {
	//for i, cmd := range sj.cmdList {
	//	sj.taskChan <- i
	//	go ShellExecutor(cmd, sj.taskChan)
	//	//<- myCh
	//}

	//go func() {
	//	for {
	//		select {
	//		case taskId := <- sj.waitChan:
	//			fmt.Println("等待执行",taskId)
	//		}
	//	}
	//}()

	taskInfos := sj.taskInfos

LOOPTASK:
	//fmt.Println(sj.taskInfos)
	for i, taskInfo := range taskInfos {
		var dependsTask []int
		err := json.Unmarshal([]byte(taskInfo.Depends), &dependsTask)
		if err != nil {
			fmt.Printf("unmarshal err=%v\n", err)
		}
		if CheckDepends(dependsTask) { // 检查依赖是否运行完成
			sj.taskChan <- i
			go ShellExecutor(taskInfo.Command, sj.taskChan)
		} else {
			sj.waitChan <- taskInfo
		}

		if len(taskInfos) > 0 && i == len(taskInfos)-1 {
			fmt.Println("运行完成")
			taskInfos = taskInfos[0:0]
			for {
				select {
				case chanTask := <-sj.waitChan:
					taskInfos = append(taskInfos, chanTask)
				default:
					goto LOOPTASK
				}
			}
		}
		//fmt.Printf("反序列化后 slice=%v\n", slice)

	}

}
