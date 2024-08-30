package main

import (
	l "new_ip_data_api/config/logger"
	db "new_ip_data_api/conn"
	"new_ip_data_api/controller"
	"new_ip_data_api/repository"
	"new_ip_data_api/route"
	"new_ip_data_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	l.LoggerInit()
	defer l.Logger.Sync()
	l.Logger.Info("Aplicação iniciada")

	server := gin.Default()

	dbConnection, err := db.Init()
	if err != nil {
		panic(err)
	}

	IpDataRepository := repository.NewIpDataRepository(dbConnection)
	IpDataUsecase := usecase.NewIpDataUsecase(IpDataRepository)
	ipDataController := controller.NewIpDataController(IpDataUsecase)
	route.Endpoints(server, &ipDataController)
}
