package poll_test

import (
	"fmt"
	"reflect"
	"slices"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/fetch"
	dailyboard "sokwva/acfun/billboard/fetch/subPart"
	"sokwva/acfun/billboard/parser"
	"strconv"
	"testing"

	"github.com/BurntSushi/toml"
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

var (
	tasks map[string]TaskContainer = map[string]TaskContainer{
		"dailyBanana": {
			TargetUrl: "",
			Trigger: func(Url string, taskName string) {
				var fetchResp []string = []string{}
				detailInfo, err := dailyboard.BananaBoard()
				if err != nil {
					fmt.Println("Here")
					//HTML获取
					str, err := dailyboard.BananaBoardHTML()
					if err == nil {
						fetchResp, err = parser.BananaBoard(str)
						if err != nil {
							departDone <- taskName
							return
						}
					} else {
						//获取失败则直接使用上次的acid任务列表
						if reflect.TypeOf(lastSuccessResp[taskName]).Kind() == reflect.Slice {
							fetchResp = (lastSuccessResp[taskName]).([]string)
						}
						common.Log.Info("poller.commonTask: call " + taskName + " faild,use last fetch result." + strconv.Itoa(len(fetchResp)))
					}

					//执行任务：从上面逐个获取的稿件acid列表取出acid，然后逐个获取稿件信息
					for _, v := range fetchResp {
						common.Log.Debug("poller.commonTask: task[" + v + "] ready to call fetch.FetchInfoAndSaveToTSDB")
						fmt.Println(taskName, v)
					}
				} else {
					for _, v := range detailInfo.RankList {
						common.Log.Debug("poller.commonTask: direct task[" + v.DougaID + "] ready to call fetch.FetchInfoAndSaveToTSDB")
						tags := map[string]string{
							"acid": v.DougaID,
						}
						fields := map[string]interface{}{
							"commentCount": v.CommentCountRealValue,
							"stowCount":    v.StowCount,
							"likeCount":    v.LikeCount,
							"shareCount":   v.ShareCount,
							"danmuCount":   v.DanmakuCount,
							"viewCount":    v.ViewCount,
							"bananaCount":  v.BananaCount,
						}
						fmt.Println(taskName, tags, fields)
						// timeseries.SaveTSRecord(taskName, tags, fields)
					}
				}
				lastSuccessResp[taskName] = fetchResp
				departDone <- taskName
			},
		},
		"animation": {
			TargetUrl: fetch.AnimateIndexUrl,
			Trigger:   commonTask(),
		},
		"music": {
			TargetUrl: fetch.MusicIndexUrl,
			Trigger:   commonTask(),
		},
		"entertain": {
			TargetUrl: fetch.EntertainIndexUrl,
			Trigger:   commonTask(),
		},
		"dance": {
			TargetUrl: fetch.DanceIndexUrl,
			Trigger:   commonTask(),
		},
		"game": {
			TargetUrl: fetch.GameIndexUrl,
			Trigger:   commonTask(),
		},
		"tech": {
			TargetUrl: fetch.TechIndexUrl,
			Trigger:   commonTask(),
		},
		"movie": {
			TargetUrl: fetch.MovieIndexUrl,
			Trigger:   commonTask(),
		},
		"sport": {
			TargetUrl: fetch.SportIndexUrl,
			Trigger:   commonTask(),
		},
		"fishpond": {
			TargetUrl: fetch.FishPondIndexUrl,
			Trigger:   commonTask(),
		},
		"article": {
			TargetUrl: fetch.ArticleIndexUrl,
			Trigger:   commonTask(),
		},
	}
)

func commonTask() func(Url, taskName string) {
	return func(Url string, taskName string) {
		//任务acid列表
		var fetchResp []string
		detailInfo, err := dailyboard.SubPartStr(Url)
		fmt.Println(err.Error())
		if err != nil {
			//获取失败就退回HTML获取方式
			str, err := dailyboard.SubPartStrHTML(Url)
			if err == nil {
				fetchResp, err = parser.CommonSubPart(str)
				if err != nil {
					departDone <- taskName
					return
				}
			} else {
				//获取失败则直接使用上次的acid任务列表
				if reflect.TypeOf(lastSuccessResp[taskName]).Kind() == reflect.Slice {
					fetchResp = (lastSuccessResp[taskName]).([]string)
				}
				common.Log.Info("poller.commonTask: call " + taskName + " faild,use last fetch result." + strconv.Itoa(len(fetchResp)))
			}
			//执行任务：从上面逐个获取的稿件acid列表取出acid，然后逐个获取稿件信息
			for _, v := range fetchResp {
				common.Log.Debug("poller.commonTask: task[" + v + "] ready to call fetch.FetchInfoAndSaveToTSDB")
				fmt.Println(taskName, v)
			}
		} else {
			//直接从接口获取成功
			for _, v := range detailInfo.RankList {
				common.Log.Debug("poller.commonTask: direct task[" + v.DougaID + "] ready to call fetch.FetchInfoAndSaveToTSDB")
				tags := map[string]string{
					"acid": v.DougaID,
				}
				fields := map[string]interface{}{
					"commentCount": v.CommentCountRealValue,
					"stowCount":    v.StowCount,
					"likeCount":    v.LikeCount,
					"shareCount":   v.ShareCount,
					"danmuCount":   v.DanmakuCount,
					"viewCount":    v.ViewCount,
					"bananaCount":  v.BananaCount,
				}
				fmt.Println(taskName, tags, fields)

				// timeseries.SaveTSRecord(taskName, tags, fields)

			}
		}
		lastSuccessResp[taskName] = fetchResp
		departDone <- taskName
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

var ConfHandle common.Conf

func Test_poller(t *testing.T) {
	if _, err := toml.DecodeFile("../conf.toml", &ConfHandle); err != nil {
		panic(err)
	}
	common.InitLogger()
	common.Log.Debug("poll.Poller start")
	common.Log.Debug("poll.Poller start action goroutine")
	action()
}
