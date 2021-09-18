package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}

func handler(ctx context.Context, event MyEvent) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello World. key1=%s, key2=%s, key3=%s", event.Key1, event.Key2, event.Key3),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
