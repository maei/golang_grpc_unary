package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_grpc_sum_client/src/services"
	"github.com/maei/shared_utils_go/rest_errors"
	"io/ioutil"
	"net/http"
)

var CalculatorController calculatorControllerInterface = &calculatorController{}

type calculatorControllerInterface interface {
	Calculate(c echo.Context) error
}

type calculatorController struct{}

func (*calculatorController) Calculate(c echo.Context) error {
	defer c.Request().Body.Close()
	bodyJSON := c.Request().Body

	bs, err := ioutil.ReadAll(bodyJSON)
	if err != nil {
		return rest_errors.NewInternalServerError("failed reading json body", err)
	}

	result, calcErr := services.CalculatorService.Calculation(bs)
	if calcErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]rest_errors.RestErr{
			"message": calcErr,
		})
	}

	return c.JSON(http.StatusOK, map[string]int32{
		"message": result,
	})
}
