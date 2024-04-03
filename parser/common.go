package parser

import (
	"fmt"
	"sokwva/acfun/billboard/common"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CommonSubPart(raw string) (result []string, err error) {
	result = make([]string, 10)
	document, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resultTypes := document.Find("ul")
	resultTypes.Each(func(i int, s *goquery.Selection) {
		resultItems := s.Find("li")
		resultItems.Each(func(i int, s *goquery.Selection) {
			name := s.Find("a")
			href, ok := name.Attr("href")
			acid := strings.ReplaceAll(href, "https://m.acfun.cn/v/?ac=", "")
			if ok {
				result[i] = acid
			}
			common.Log.Debug("parser.CommonSubPart: ", acid, name.Text())
		})
	})
	return
}
