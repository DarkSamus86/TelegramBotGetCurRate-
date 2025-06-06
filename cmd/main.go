package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	telegramAPIURL = "https://api.telegram.org/bot"
)

type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(token, text string, chatID int64) error {
	url := fmt.Sprintf("%s%s/sendMessage", telegramAPIURL, token)

	body := SendMessageRequest{
		ChatID: chatID,
		Text:   text,
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
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить .env файл:", err)
		return
	}
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("TELEGRAM_TOKEN environment variable not set")
	}

	chatID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		fmt.Println("CHAT_ID environment variable not set")
	}

	err = sendMessage(token, "Ответ телеграм бота", chatID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Сообщение отправлено")
}
