package main

import (
	"fmt"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/fetch/dougaInfo"
)

type ApiResp struct {
	Container string   `json:"container"`
	Id        string   `json:"id"`
	HTML      string   `json:"html"`
	Style     []string `json:"styles"`
	Scripts   []string `json:"scripts"`
	Mode      string   `json:"mode"`
	Js        []string `json:"js"`
	Css       []string `json:"css"`
}

func main() {
	common.InitConfDriver()
	common.InitLogger()
	// err := timeseries.InitClient()
	// if err != nil {
	// 	common.Log.Error("init timeseries driver faild: " + err.Error())
	// 	return
	// }

	// poll.Poller()
	dougaInfo.InitGrpcClient()
	a, err := dougaInfo.GetVideoInfo("4741173")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(a)
}
