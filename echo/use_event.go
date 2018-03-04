package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler_Blank() (Response, error) {
	return Response{
		Message: "Go Serverless v1.0! Your function executed successfully!",
	}, nil
}

func makeResponse(status int, body interface{}) events.APIGatewayProxyResponse {
	data, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: status,
	}
}

func Handler_APIGateway(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handlerAPIGateway_GET(request)
	case "POST":
		return handlerAPIGateway_POST(request)
	default:
		return handlerAPIGateway_POST(request)
	}
}

func handlerAPIGateway_GET(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var resp EchoResponse
	msg, ok := request.QueryStringParameters["q"]
	if ok {
		resp = NewEchoResponse(msg, false)
	} else {
		resp = NewEchoResponse("", true)
	}

	return makeResponse(200, resp), nil
}

func handlerAPIGateway_POST(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := EchoResponse{
		Message: request.Body,
		Blank:   true,
	}

	return makeResponse(200, resp), nil
}

func mainWithEvent() {
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		// LAMBDA_TASK_ROOT=/var/task
		lambda.Start(Handler_APIGateway)
	} else {
		fmt.Println("dev mode")
	}
}
