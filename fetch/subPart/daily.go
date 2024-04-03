package dailyboard

import (
	"encoding/json"
	"fmt"
	"sokwva/acfun/billboard/fetch"
	"strings"
)

func SubPartStr(targetUrl string) (detailInfo *ApiResp, err error) {
	detailInfo = &ApiResp{}
	raw, err := fetch.Post(fetch.BillboardApiUrl, "rankPeriod=DAY&channelId="+fetch.GetPartIdFromUrl(targetUrl)+"&rankLimit=10", true)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(raw), detailInfo)
	if err != nil {
		return
	}
	return
}

func BananaBoard() (detailInfo *ApiResp, err error) {
	raw, err := fetch.Get(fetch.BillboardApiUrl, true)
	detailInfo = &ApiResp{}
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(raw), detailInfo)
	return
}

func BananaBoardHTML() (targetSlice string, err error) {
	raw, err := fetch.Get(fetch.IndexUrl, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	headerSlice := strings.Split(raw, "为什么只能投5蕉</span></div>")
	tailSlice := strings.Split(headerSlice[1], "<!--香蕉榜广告位-->")
	targetSlice = tailSlice[0]

	targetSlice = strings.ReplaceAll(targetSlice, "\\n", "")
	targetSlice = strings.ReplaceAll(targetSlice, "\\", "")

	return
}

func SubPartStrHTML(targetUrl string) (targetSlice string, err error) {
	raw, err := fetch.Get(targetUrl, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	headerSlice := strings.Split(raw, "class=\"rank-list-main\"")
	tailSlice := strings.Split(headerSlice[1], "<div class=\"footer\"")
	targetSlice = tailSlice[0]

	// 去除 "data-v-08099b04>"
	targetSlice = targetSlice[17:]

	targetSlice = strings.ReplaceAll(targetSlice, "\\n", "")
	targetSlice = strings.ReplaceAll(targetSlice, "\\", "")

	return
}

func ArticleSubPart() (detailInfo *ArticleApi, err error) {
	detailInfo = &ArticleApi{}
	raw, err := fetch.Post(fetch.BillboardApiUrl, "rankPeriod=DAY&channelId="+fetch.GetPartIdFromUrl(fetch.ArticleIndexUrl)+"&rankLimit=10", true)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(raw), detailInfo)
	return
}

func ArticleSubPartStr() (targetSlice string, err error) {
	targetSlice, err = fetch.Get(fetch.ArticleIndexUrl, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
