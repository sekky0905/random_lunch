package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sekky0905/random_lunch/src/application"
	"github.com/sekky0905/random_lunch/src/functions/common"
)


func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token := os.Getenv("VERIFICATION_TOKEN")
	values, err := url.ParseQuery(req.Body)
	if err != nil {
		log.Printf("failed to choose shop. original error = %+v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil
	}

	slackReq := common.GenerateSlackRequestFromValues(values)

	if slackReq.Token != token {
		log.Printf("invalid token. token = %+v", token)
		return events.APIGatewayProxyResponse{
			Body:       "Invalid token.",
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	a := application.ShopApplicationService{}
	params, err := a.ChooseShops(slackReq)
	if err != nil {
		log.Printf("failed to choose shop. original error = %+v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	jsonBytes, err := json.Marshal(params)
	if err != nil {
		log.Printf("failed to choose shop. original error = %+v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
