package common

import (
	"net/url"

	"github.com/sekky0905/random_lunch/src/models"
)

func GenerateSlackRequestFromValues(values url.Values) *models.SlackReq {
	token := values.Get("token")
	channelID := values.Get("channel_id")
	text := values.Get("text")
	return models.NewSlackSlackReq(token, channelID, text)
}
