package main

import (
	"aws-lambda-simple-example/data/repositories"
	"aws-lambda-simple-example/domain"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var book domain.Book

	id := event.PathParameters["id"]

	if err := json.Unmarshal([]byte(event.Body), &book); err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewBookRepository()
	result := repository.Update(id, book)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
