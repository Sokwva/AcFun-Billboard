package poll

import (
	"slices"
	"sokwva/acfun/billboard/common"
	"time"
)

var (
	lastSuccessResp map[string]interface{} = make(map[string]interface{})
	allDone         chan bool              = make(chan bool)

	departDone chan string = make(chan string)
)

type TaskContainer struct {
	TargetUrl string
	Trigger   func(Url string, taskName string)
}

func Poller() {
	common.Log.Debug("poll.Poller start")
	for {
		timer := time.NewTimer(time.Duration(common.ConfHandle.Poller.Interval) * time.Minute)
		<-timer.C
		common.Log.Debug("poll.Poller start action goroutine")
		go action()
	}
}

func action() {
	go depTaskCheck()
	for i, v := range tasks {
		go v.Trigger(v.TargetUrl, i)
	}
	<-allDone
	common.Log.Debug("poll.depTaskCheck all done")
}

// 总分区任务情况检查
func depTaskCheck() {
	common.Log.Debug("poll.depTaskCheck start")
	var workList []string = []string{}
	for i := range tasks {
		workList = append(workList, i)
	}
	for {
		if len(workList) == 0 {
			allDone <- true
			return
		}
		item := <-departDone
		common.Log.Debug("poll.depTaskCheck " + item + " done")
		workList = slices.DeleteFunc(workList, func(s string) bool {
			return s == item
		})
	}
}

// 分区内排行榜任务情况检查
func taskCheck(workListSrc []string, allLocalDone chan bool, perTaskDone chan string) {
	var workList []string
	copy(workList, workListSrc)
	for {
		if len(workList) == 0 {
			allLocalDone <- true
			return
		}
		item := <-perTaskDone
		workList = slices.DeleteFunc(workList, func(s string) bool {
			return s == item
		})
	}
}
