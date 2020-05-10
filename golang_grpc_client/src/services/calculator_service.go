package services

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/maei/golang_grpc_sum_client/src/domain/calculatorpb"
	"github.com/maei/golang_grpc_sum_client/src/grpc_client"
	"github.com/maei/shared_utils_go/rest_errors"
	"log"
)

var CalculatorService calculatorServiceInterface = &calculatorService{}

type calculatorServiceInterface interface {
	Calculation(bs []byte) (int32, rest_errors.RestErr)
}

type calculatorService struct{}

func (*calculatorService) Calculation(bs []byte) (int32, rest_errors.RestErr) {
	var cal = &calculatropb.Calculator{}
	err := jsonpb.UnmarshalString(string(bs), cal)
	if err != nil {
		log.Println()
		return 0, rest_errors.NewInternalServerError("error while string to proto", err)
	}

	conn, err := grpc_client.GRPCClient.SetClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := calculatropb.NewCalculatorServiceClient(conn)

	req := &calculatropb.CalculatorRequest{Calculation: cal}

	res, calErr := c.Calc(context.Background(), req)
	if calErr != nil {
		log.Println(calErr)
		return 0, rest_errors.NewInternalServerError("error while calculating", calErr)
	}

	return res.Result, nil
}
