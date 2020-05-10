package services

import (
	"github.com/maei/shared_utils_go/rest_errors"
	_ "github.com/maei/shared_utils_go/rest_errors"
)

var PingService pingServiceInterface = &pingService{}

type pingServiceInterface interface {
	Ping() (string, rest_errors.RestErr)
}

type pingService struct{}

func (*pingService) Ping() (string, rest_errors.RestErr) {
	return "Hello World", nil
}
