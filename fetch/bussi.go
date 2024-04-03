package fetch

import (
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/db/timeseries"
	"sokwva/acfun/billboard/fetch/dougaInfo"
	"strings"
)

func FetchInfoAndSaveToTSDB(measurement string, acid string, done chan string) {
	info, err := dougaInfo.GetVideoInfo(acid)
	if err != nil {
		common.Log.Error(err.Error())
		return
	}
	tags := map[string]string{
		"acid": acid,
	}
	fields := map[string]interface{}{
		"commentCount": info.CommentCountRealValue,
		"stowCount":    info.StowCount,
		"likeCount":    info.LikeCount,
		"shareCount":   info.ShareCount,
		"danmuCount":   info.DanmakuCount,
		"viewCount":    info.ViewCount,
		"bananaCount":  info.BananaCount,
	}
	timeseries.SaveTSRecord(measurement, tags, fields)
	common.Log.Debug("write data to tsdb:", acid, fields)
	done <- acid
}

func GetPartIdFromUrl(targetUrl string) string {
	x := strings.Replace(targetUrl, "https://www.acfun.cn/v/list", "", -1)
	return strings.Replace(x, "/index.htm", "", -1)
}
