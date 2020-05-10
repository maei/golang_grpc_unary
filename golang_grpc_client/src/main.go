package main

import (
	"github.com/maei/golang_grpc_sum_client/src/app"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info("starting calculator-service")
	app.StartApplication()
}
