package models

import (
	"bytes"
	"fmt"
)

const UserName = "ランチのお店決めてくれる君"

// SlackParams は、Slack のパラメータを表す。
type SlackParams struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// SlackReq は、Slack のリクエストを表す。
type SlackReq struct {
	Token     string
	ChannelID string
	Text      string
}

// NewSlackSlackReq は、SlackSlackReq を生成し、返す。
func NewSlackSlackReq(token, channelID, Text string) *SlackReq {
	return &SlackReq{
		Token:     token,
		ChannelID: channelID,
		Text:      Text,
	}
}

// BuildChoiceText は、選択用のテキストを作成する。
func BuildChoiceText(shops []Shop) string {
	buf := bytes.Buffer{}

	buf.WriteString("お店が決まりました!\n")
	for i, shop := range shops {
		str := fmt.Sprintf("%d:【店名】%s 【URL】%s 【メモ】%s\n", i+1, shop.Name, shop.URL, shop.Memo)
		buf.WriteString(str)
	}
	return buf.String()
}

// BuildCreateText は、生成用のテキストを作成する。
func BuildCreateText(shop *Shop) string {
	return fmt.Sprintf("お店を登録しました!\n【店名】%s 【URL】%s 【メモ】%s\n", shop.Name, shop.URL, shop.Memo)
}

// BuildListText は、一覧取得用のテキストを作成する。
func BuildListText(shops []Shop) string {
	buf := bytes.Buffer{}

	buf.WriteString("登録されているお店一覧です!\n")
	for i, shop := range shops {
		str := fmt.Sprintf("%d:【店名】%s 【URL】%s 【メモ】%s\n", i+1, shop.Name, shop.URL, shop.Memo)
		buf.WriteString(str)
	}
	return buf.String()
}
