package main

import (
	"aws-lambda-simple-example/data/repositories"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := event.PathParameters["id"]

	repository := repositories.NewBookRepository()
	result := repository.Delete(id)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
