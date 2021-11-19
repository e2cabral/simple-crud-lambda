package main

import (
	"aws-lambda-simple-example/data/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	repository := repositories.NewBookRepository()
	result := repository.Get()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
