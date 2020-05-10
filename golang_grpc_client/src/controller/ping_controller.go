package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/maei/golang_grpc_sum_client/src/services"
	"net/http"
)

var PingController pingControllerInterface = &pingController{}

type pingControllerInterface interface {
	Ping(c echo.Context) error
}

type pingController struct{}

func (*pingController) Ping(c echo.Context) error {
	result, err := services.PingService.Ping()
	if err != nil {
		log.Error("ping failed", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": result,
	})
}
