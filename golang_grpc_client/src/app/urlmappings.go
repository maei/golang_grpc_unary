package app

import "github.com/maei/golang_grpc_sum_client/src/controller"

func publicRoutes() {
	router.GET("/ping", controller.PingController.Ping)
	router.POST("/calculate", controller.CalculatorController.Calculate)

}

func mapUrls() {
	publicRoutes()
}
