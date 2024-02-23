package parser

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func BananaBoard(raw string) (result []string, err error) {
	result = make([]string, 10)
	document, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//三类榜单：日榜 三日榜 周榜
	resultTypes := document.Find("div.banana-list")
	resultTypes.Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")
		//但是只记录今天的日榜
		if class != "banana-list day-list" {
			return
		}

		resultItems := s.Find("div.banana-video.log-item")
		resultItems.Each(func(i int, s *goquery.Selection) {
			name := s.Find("a.banana-video-title")
			href, ok := name.Attr("href")
			acid := strings.ReplaceAll(href, "/v/ac", "")
			if ok {
				result[i] = acid
			}
			fmt.Println(acid, name.Text())
		})
	})
	return
}
