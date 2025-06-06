package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramApi struct {
	Token string
}

type sendMessageRequest struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramApi(token string) *TelegramApi {
	return &TelegramApi{Token: token}
}

func (api *TelegramApi) SendMessage(chatId int64, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", api.Token)

	body := sendMessageRequest{
		ChatId: chatId,
		Text:   message,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram api response status code %d", resp.StatusCode)
	}

	return nil
}
