package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

var initialized = false
var muxLambda *gorillamux.GorillaMuxAdapter

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		rootRoute := mux.NewRouter()
		api1.Init(rootRoute)
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Mux cold start")
		muxLambda = gorillamux.New(rootRoute)

		initialized = true
	}

	// If no name is provided in the HTTP request body, throw an error
	return muxLambda.Proxy(req)
}

func main() {

	lambda.Start(Handler)
}