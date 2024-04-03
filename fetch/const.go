package fetch

var HEADERS = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
	"accept":          "application/json, text/plain, */*",
	"accept-encoding": "gzip, deflate, br",
	"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"origin":          "https://www.acfun.cn",
	"referer":         "",
	"content-type":    "application/x-www-form-urlencoded",
}

const (
	IndexUrl          = "https://www.acfun.cn/"
	BangumiIndexUrl   = "https://www.acfun.cn/v/list155/index.htm"
	AnimateIndexUrl   = "https://www.acfun.cn/v/list1/index.htm"
	MusicIndexUrl     = "https://www.acfun.cn/v/list58/index.htm"
	EntertainIndexUrl = "https://www.acfun.cn/v/list60/index.htm"
	DanceIndexUrl     = "https://www.acfun.cn/v/list123/index.htm"
	GameIndexUrl      = "https://www.acfun.cn/v/list59/index.htm"
	TechIndexUrl      = "https://www.acfun.cn/v/list70/index.htm"
	MovieIndexUrl     = "https://www.acfun.cn/v/list68/index.htm"
	SportIndexUrl     = "https://www.acfun.cn/v/list69/index.htm"
	FishPondIndexUrl  = "https://www.acfun.cn/v/list125/index.htm"
	ArticleIndexUrl   = "https://www.acfun.cn/v/list63/index.htm"

	BillboardApiUrl = "https://www.acfun.cn/rest/pc-direct/rank/channel"
	ApiSrcUrl       = "https://www.acfun.cn/?pagelets=pagelet_header,pagelet_game,pagelet_douga,pagelet_amusement,pagelet_bangumi_list,pagelet_life,pagelet_tech,pagelet_dance,pagelet_music,pagelet_film,pagelet_fishpond,pagelet_sport&reqID=0&ajaxpipe=1"
)

const (
	BillbordItem_Type_Day  = "day"
	BillbordItem_Type_3Day = "3-day"
	BillbordItem_Type_Week = "week"
)
