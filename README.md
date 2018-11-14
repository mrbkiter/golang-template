# golang-template
golang template project for lambda and server side. The project contains following packages: 

- api1 (API v1)
- api1/model (model package of APIv1)
- model (internal model)
- app (services)
- repo
- utils
- web (contains context)

For pacakges related versions (API1) it should not be used in other packages. This approach wold help isolate API package versions 
and internal ones. 

How to build and deploy.
1. Install golang dep. 
2. Go to template.github.com/server, running ```dep init```. It would initialize project.
3. Run ```dep ensure``` to download dependencies. 

Prepare postgres db:
1. If you look into server.go, database configuration is declared there. Please update it properly. 
2. Create dbname as what you inputed. 
3. Manually run script.sql in repo directory. Now you can start up your app. 

- For normal app server

Go to /server/command package, run ```go run server.go```. It would start server at port 8000

- For lambda function

Go to /server/lambda, run following command:

- ```env GOOS=linux GOARCH=amd64 go build```
- ```zip lambda.zip lambda```

Then deploy lambda.zip to your lambda function. After that you need to bind an API Gateway that proxy to your function. 

How to test API:

There are 2 endpoints: 

- GET /api/v1/candidates/{id} 
- POST /api/v1/candidates 
Content-Type: application/json 
Body: 
{
	"id": "123",
	"firstName": "1234"
}

