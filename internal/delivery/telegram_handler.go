package delivery

import (
	"fmt"
	"github.com/DarkSamus86/TelegramBotGetCurRate-/config"
	"github.com/DarkSamus86/TelegramBotGetCurRate-/internal/usecase"
)

type TelegramHandler struct {
	usecase *usecase.MessageUseCase
}

func NewTelegramHandler(uc *usecase.MessageUseCase) *TelegramHandler {
	return &TelegramHandler{usecase: uc}
}

func (h *TelegramHandler) Run() {
	chatID := config.GetTelegramChatID()
	text := "Hi from server"

	err := h.usecase.SendMessage(chatID, text)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Success")
	}
}
