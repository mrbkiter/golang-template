package main

import (
	"log"

	"template.github.com/server/config"
	"template.github.com/server/repo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

var muxLambda *gorillamux.GorillaMuxAdapter

func init() {
	rootRoute := mux.NewRouter()
	api1.Init(rootRoute)
	dbConfig := &config.DatabaseConfig{}
	dbConfig.Driver = "postgres"
	dbConfig.JdbcUrl = `host=localhost port=5432 sslmode=disable 
	dbname=testgo user=dbapplication_user password=dbapplication_user`
	config.GetConfig().DatabaseConfig = dbConfig

	//initialize repo
	repo.Init()
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Mux start")
	muxLambda = gorillamux.New(rootRoute)
}

//RequestHandler Handler that process APIGatewayProxyRequest
func RequestHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return muxLambda.Proxy(req)
}

func main() {

	lambda.Start(RequestHandler)
}
