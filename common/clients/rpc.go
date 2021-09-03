package clients

import (
	accountRpc "github.com/zbwang163/ad_account_overpass"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_info_platform/common/consts"
	"google.golang.org/grpc"
	"log"
)

var (
	AccountClient accountRpc.AccountServiceClient
	ContentClient contentRpc.ContentServiceClient
	rpcConfig     = map[string]string{
		"ad.info.account_server": "account-service:50001",
		"ad.info.content_server": "ad.info.content_server:50002",
	}
)

func init() {
	InitRpc(consts.AccountServerPSM)
	InitRpc(consts.ContentServerPSM)
}

func InitRpc(psm string) {
	conn, err := grpc.Dial(rpcConfig[psm], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	switch psm {
	case "ad.info.account_server":
		AccountClient = accountRpc.NewAccountServiceClient(conn)
	case "ad.info.content_server":
		ContentClient = contentRpc.NewContentServiceClient(conn)
	}
	_ = AccountClient
	_ = ContentClient
}
