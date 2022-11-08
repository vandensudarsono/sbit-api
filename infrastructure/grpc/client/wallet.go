package client

import (
	pb "sbit/infrastructure/grpc/proto/wallet"
	logging "sbit/infrastructure/log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var walletClient pb.SbitServiceClient

func NewWalletClient(optionts ...string) {
	opts := []grpc.DialOption{}
	srvAddr := viper.GetString("walletService.address")

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(srvAddr, opts...)
	if err != nil {
		logging.Fatalf("fail to dial wallet client service: %v", err)
	}

	walletClient = pb.NewSbitServiceClient(conn)

}

func GetWalletClient() pb.SbitServiceClient {
	return walletClient
}
