package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Blank   bool   `json:"blank"`
	Message string `json:"message"`
}

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
		resp = EchoResponse{
			Message: msg,
			Blank:   false,
		}
	} else {
		resp = EchoResponse{
			Message: "",
			Blank:   true,
		}
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

func main() {
	lambda.Start(Handler_APIGateway)
}
