package dailybananaboard

import (
	"fmt"
	"sokwva/acfun/billboard/fetch"
	"sokwva/acfun/billboard/parser"
	"strings"
)

func BananaBoard() (resultList []string, err error) {
	raw, err := fetch.Do(fetch.IndexUrl, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	headerSlice := strings.Split(raw, "为什么只能投5蕉</span></div>")
	tailSlice := strings.Split(headerSlice[1], "<!--香蕉榜广告位-->")
	targetSlice := tailSlice[0]

	targetSlice = strings.ReplaceAll(targetSlice, "\\n", "")
	targetSlice = strings.ReplaceAll(targetSlice, "\\", "")

	resultList, err = parser.BananaBoard(targetSlice)
	return
}
