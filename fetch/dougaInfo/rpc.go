package dougaInfo

import (
	"context"
	"encoding/json"
	"sokwva/acfun/billboard/common"
	rpcproto "sokwva/acfun/dougaInfo/server/protoLib"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var rpcClient *grpc.ClientConn
var RPCConn rpcproto.GetClient

func InitGrpcClient(userName string, target string, port string) error {
	var authSet []grpc.DialOption
	if userName == "" {
		authSet = append(authSet, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	rpcClient, err := grpc.NewClient(target+":"+port, authSet...)
	common.Log.Debug("dougaInfo.InitGrpcClient: grpc.NewClient ready to call")
	if err != nil {
		return err
	}
	common.Log.Debug("dougaInfo.InitGrpcClient: grpc.NewClient ok")
	RPCConn = rpcproto.NewGetClient(rpcClient)
	common.Log.Debug("dougaInfo.InitGrpcClient: rpcproto.NewGetClient create client")
	return nil
}

func CloseGrpcClient() {
	if rpcClient != nil {
		rpcClient.Close()
	}
}

func GetVideoInfo(acid string) (DougaInfo, error) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()
	resp, err := RPCConn.GetDougaInfo(ctx, &rpcproto.Acid{Acid: acid})
	if err != nil {
		return DougaInfo{}, err
	}
	var result DougaInfo
	bytesRes, err := json.Marshal(resp)
	if err != nil {
		return DougaInfo{}, err
	}
	err = json.Unmarshal(bytesRes, &result)
	if err != nil {
		return DougaInfo{}, err
	}
	return result, nil
}
