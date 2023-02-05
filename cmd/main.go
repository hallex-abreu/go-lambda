package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	QuantityPersons   int `json:"quantity_persons"`
	QuantityKids      int `json:"quantity_kids"`
	QuantityConectors int `json:"quantity_conectors"`
	QuantityVisitors  int `json:"quantity_visitors"`
}

type LambdaResponse struct {
	Message string
}

func LambdaHandler(event LambdaEvent) (LambdaResponse, error) {
	return LambdaResponse{
		Message: fmt.Sprintf("quantity persons %d, quantity kids %d.", event.QuantityPersons, event.QuantityKids),
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
