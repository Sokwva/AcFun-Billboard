package saveDougaInfoToDb

import (
	"fmt"
	"sokwva/acfun/billboard/common"
	"sokwva/acfun/billboard/db/timeseries"
	"sokwva/acfun/billboard/fetch/dougaInfo"
	"testing"
)

func Test_mongoConn(t *testing.T) {
	common.InitConfDriver("../../conf.toml")
	common.InitLogger()
	err := timeseries.InitClient()
	if err != nil {
		common.Log.Error("init timeseries driver faild: " + err.Error())
		return
	}
	err = dougaInfo.InitGrpcClient(common.ConfHandle.RPC.DougaInfo.UserName, common.ConfHandle.RPC.DougaInfo.Addr, common.ConfHandle.RPC.DougaInfo.Port)
	if err != nil {
		common.Log.Error("grpc client init faild")
		return
	}
	defer dougaInfo.CloseGrpcClient()

	err = InitClient()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(common.ConfHandle.DougaInfoSave)
	fmt.Println(CheckACIDExist("31506965"))
}
