package usecase

import (
	"github.com/DarkSamus86/TelegramBotGetCurRate-/internal/domain"
)

type MessageUseCase struct {
	Sender domain.MessageSender
}

func NewMessageUseCase(sender domain.MessageSender) *MessageUseCase {
	return &MessageUseCase{Sender: sender}
}

func (uc *MessageUseCase) SendMessage(chatId int64, message string) error {
	return uc.Sender.SendMessage(chatId, message)
}
