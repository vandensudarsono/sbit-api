package action

import (
	"sbit/adapter/controller"
	"sbit/infrastructure/rest"
	"sbit/infrastructure/rest/routers"
	v1 "sbit/infrastructure/rest/routers/v1"
)

func RunRest() {
	//instance server
	server := rest.NewHttpServer()

	//controller
	depositController := controller.NewDepositController()
	walletController := controller.NewWalletController()

	//setup router v1
	v1Group := routers.InitV1(server.App)
	v1.InitDepositGroup(v1Group, depositController)
	v1.InitWalletGroup(v1Group, walletController)

	server.Run()
}
