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

func InitGrpcClient() error {
	var authSet []grpc.DialOption
	if common.ConfHandle.RPC.DougaInfo.UserName == "" {
		authSet = append(authSet, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	rpcClient, err := grpc.Dial(common.ConfHandle.RPC.DougaInfo.Addr+":"+common.ConfHandle.RPC.DougaInfo.Port, authSet...)
	if err != nil {
		return err
	}
	RPCConn = rpcproto.NewGetClient(rpcClient)
	return nil
}

func CloseGrpcClient() {
	rpcClient.Close()
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
