package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load the .env file:", err)
	}
}

func GetTelegramToken() string {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN environment variable not set")
	}
	return token
}

func GetTelegramChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		fmt.Println("CHAT_ID environment variable not set")
	}
	return chatID
}
