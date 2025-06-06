package main

import (
	"github.com/DarkSamus86/TelegramBotGetCurRate-/config"
	"github.com/DarkSamus86/TelegramBotGetCurRate-/internal/delivery"
	"github.com/DarkSamus86/TelegramBotGetCurRate-/internal/infrastructure"
	"github.com/DarkSamus86/TelegramBotGetCurRate-/internal/usecase"
)

func main() {
	config.LoadEnv()
	token := config.GetTelegramToken()

	telegram := infrastructure.NewTelegramApi(token)
	msgUC := usecase.NewMessageUseCase(telegram)
	handler := delivery.NewTelegramHandler(msgUC)

	handler.Run()
}
