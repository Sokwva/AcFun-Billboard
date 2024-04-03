package poll

import (
	"reflect"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/db/timeseries"
	"sokwva/acfun/billboard/fetch"
	dailyboard "sokwva/acfun/billboard/fetch/subPart"
	"sokwva/acfun/billboard/parser"
	"strconv"
)

func commonTask() func(Url, taskName string) {
	return func(Url string, taskName string) {
		var perTaskDone chan string = make(chan string)
		var localTaskDone chan bool = make(chan bool)
		//任务acid列表
		var fetchResp []string
		detailInfo, err := dailyboard.SubPartStr(Url)
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
			//启动任务检查协程
			go taskCheck(fetchResp, localTaskDone, perTaskDone)
			//执行任务：从上面逐个获取的稿件acid列表取出acid，然后逐个获取稿件信息
			for _, v := range fetchResp {
				common.Log.Debug("poller.commonTask: task[" + v + "] ready to call fetch.FetchInfoAndSaveToTSDB")
				go fetch.FetchInfoAndSaveToTSDB(taskName, v, perTaskDone)
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
				timeseries.SaveTSRecord(taskName, tags, fields)
			}
		}
		//等待完成
		<-localTaskDone
		lastSuccessResp[taskName] = fetchResp
		departDone <- taskName
	}
}

var (
	tasks map[string]TaskContainer = map[string]TaskContainer{
		"dailyBanana": {
			TargetUrl: "",
			Trigger: func(Url string, taskName string) {
				var perTaskDone chan string = make(chan string)
				var localTaskDone chan bool = make(chan bool)
				var fetchResp []string
				detailInfo, err := dailyboard.BananaBoard()
				if err != nil {
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
					//启动任务检查协程
					go taskCheck(fetchResp, localTaskDone, perTaskDone)
					//执行任务：从上面逐个获取的稿件acid列表取出acid，然后逐个获取稿件信息
					for _, v := range fetchResp {
						common.Log.Debug("poller.commonTask: task[" + v + "] ready to call fetch.FetchInfoAndSaveToTSDB")
						go fetch.FetchInfoAndSaveToTSDB(taskName, v, perTaskDone)
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
						timeseries.SaveTSRecord(taskName, tags, fields)
					}
				}
				//等待完成
				<-localTaskDone
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
			Trigger: func(Url, taskName string) {
				var perTaskDone chan string = make(chan string)
				var localTaskDone chan bool = make(chan bool)
				var fetchResp []string
				detailInfo, err := dailyboard.ArticleSubPart()
				if err != nil {
					str, err := dailyboard.ArticleSubPartStr()
					if err == nil {
						fetchResp, err = parser.CommonSubPart(str)
						if err != nil {
							departDone <- taskName
							return
						}
					} else {
						if reflect.TypeOf(lastSuccessResp[taskName]).Kind() == reflect.Slice {
							fetchResp = (lastSuccessResp[taskName]).([]string)
						}
						common.Log.Info("poller.commonTask: call " + taskName + " faild,use last fetch result." + strconv.Itoa(len(fetchResp)))
					}
					go taskCheck(fetchResp, localTaskDone, perTaskDone)
					for _, v := range fetchResp {
						common.Log.Debug("poller.commonTask: task[" + v + "] ready to call fetch.FetchInfoAndSaveToTSDB")
						go fetch.FetchInfoAndSaveToTSDB(taskName, v, perTaskDone)
					}
				} else {
					for _, v := range detailInfo.RankList {
						common.Log.Debug("poller.commonTask: direct task[" + strconv.Itoa(v.ResourceID) + "] ready to call fetch.FetchInfoAndSaveToTSDB")
						tags := map[string]string{
							"acid": strconv.Itoa(v.ResourceID),
						}
						fields := map[string]interface{}{
							"commentCount": v.CommentCount,
							"viewCount":    v.ViewCount,
							"bananaCount":  v.BananaCount,
						}
						timeseries.SaveTSRecord(taskName, tags, fields)
					}
				}
				//等待完成
				<-localTaskDone
				lastSuccessResp[taskName] = fetchResp
				departDone <- taskName
			},
		},
	}
)
