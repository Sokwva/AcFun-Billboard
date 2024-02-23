package fetch

import (
	"fmt"
	"sokwva/acfun/billboard/fetch/dougaInfo"
)

func FetchInfoAndSaveToTSDB(acid string, done chan string) {
	info, err := dougaInfo.GetVideoInfo(acid)
	if err != nil {
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
	// timeseries.SaveTSRecord(tags, fields)
	fmt.Println(tags, fields)
	done <- acid
}
