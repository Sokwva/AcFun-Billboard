package poll

import (
	"reflect"
	"slices"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/fetch"
	dailybananaboard "sokwva/acfun/billboard/fetch/banana"
	"strconv"
	"time"
)

var lastSuccessResponse map[string]interface{} = make(map[string]interface{})

func Poller() {
	for {
		timer := time.NewTimer(time.Duration(common.ConfHandle.Poller.Interval) * time.Minute)
		<-timer.C
		go action()
	}
}

func action() {
	var dailyBananaResult []string
	dailyBananaResult, err := dailybananaboard.BananaBoard()
	if err != nil {
		if reflect.TypeOf(lastSuccessResponse["dailyBanana"]).Kind() == reflect.Slice {
			dailyBananaResult = (lastSuccessResponse["dailyBanana"]).([]string)
		}
		common.Log.Info("poller.action: call dailybananaboard.BananaBoard faild,use last fetch result." + strconv.Itoa(len(dailyBananaResult)))
		return
	}
	var done chan string = make(chan string)
	var allDone chan bool = make(chan bool)
	go taskCheck(dailyBananaResult, allDone, done)
	for _, v := range dailyBananaResult {
		go fetch.FetchInfoAndSaveToTSDB(v, done)
	}
	<-allDone
	lastSuccessResponse["dailyBanana"] = dailyBananaResult
}

func taskCheck(dailyBananaResult []string, allDone chan bool, done chan string) {
	var workList []string
	copy(workList, dailyBananaResult)
	for {
		item := <-done
		if len(workList) == 0 {
			allDone <- true
			return
		}
		workList = slices.DeleteFunc(workList, func(s string) bool {
			return s == item
		})
	}
}
