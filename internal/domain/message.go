package domain

type MessageSender interface {
	SendMessage(chatId int64, message string) error
}
