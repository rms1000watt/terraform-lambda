package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func Handler(request events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return res, ErrNameNotProvided
	}

	res.Body = "Hello " + request.Body
	res.StatusCode = 200

	return
}

func main() {
	lambda.Start(Handler)
}
