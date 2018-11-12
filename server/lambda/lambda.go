package main

import (
	"log"
	"os"

	"template.github.com/server/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

var initialized = false
var muxLambda *gorillamux.GorillaMuxAdapter

func init() {
	rootRoute := mux.NewRouter()
	api1.Init(rootRoute)
	app.App.Config.DatabaseConfig = &app.DatabaseConfig{JdbcUrl: os.Getenv("JDBC_URL")}
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Mux cold start")
	muxLambda = gorillamux.New(rootRoute)
}

//RequestHandler Handler that process APIGatewayProxyRequest
func RequestHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return muxLambda.Proxy(req)
}

func main() {

	lambda.Start(RequestHandler)
}
