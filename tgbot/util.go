package onessentialstg

import (
	"gopkg.in/telegram-bot-api.v4"
)

func NewReply(msg *tgbotapi.Message, text string) (reply tgbotapi.MessageConfig) {
	reply = tgbotapi.NewMessage(msg.Chat.ID, text)
	reply.ReplyToMessageID = msg.MessageID

	return
}
