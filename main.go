package main

import (
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/db/timeseries"
	"sokwva/acfun/billboard/fetch/dougaInfo"
	"sokwva/acfun/billboard/poll"
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
	err := timeseries.InitClient()
	if err != nil {
		common.Log.Error("init timeseries driver faild: " + err.Error())
		return
	}
	err = dougaInfo.InitGrpcClient()
	if err != nil {
		common.Log.Error("grpc client init faild")
		return
	}
	defer dougaInfo.CloseGrpcClient()

	poll.Poller()
}
