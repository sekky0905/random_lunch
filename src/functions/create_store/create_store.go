package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/sekky0905/random_lunch/src/application"
	"github.com/sekky0905/random_lunch/src/functions/common"
	"github.com/sekky0905/random_lunch/src/models"
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
	shop, err := translateTextToShop(slackReq.Text)
	if err != nil {
		log.Printf("invalid command text. original error = %+v", err)
		return events.APIGatewayProxyResponse{
			Body:       "コマンドが不正です",
			StatusCode: http.StatusOK,
		}, nil
	}

	if slackReq.Token != token {
		log.Printf("invalid token. token = %+v", token)
		return events.APIGatewayProxyResponse{
			Body:       "Invalid token.",
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	a := &application.ShopApplicationService{}
	params, err := a.CreateShop(slackReq, shop)
	if err != nil {
		log.Printf("failed to unmerchasl json bytes. body = %+v, original error = %+v", req.Body, err)
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
		StatusCode: http.StatusCreated,
	}, nil
}

func translateTextToShop(text string) (*models.Shop, error) {
	const minLength = 2

	// name url memo
	s := strings.Split(text, " ")
	if len(s) < minLength  {
		return nil, errors.Errorf("command text should be over 2. text = %s", text)
	}

	shop := &models.Shop{
		Name: s[0],
		URL:  s[1],
	}

	if len(s) == 3 {
		shop.Memo = s[2]
	}

	return shop, nil
}

func main() {
	lambda.Start(handler)
}
