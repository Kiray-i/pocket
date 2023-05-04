package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart       = "start"
	replyStartTemplate = "Приет! Чтобы сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
)

func (t *TelegramBot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return t.handleStartCommand(message)
	default:
		return t.handleUnknownCommand(message)
	}
}

func (t *TelegramBot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	t.bot.Send(msg)
}

func (t *TelegramBot) handleStartCommand(message *tgbotapi.Message) error {
	authLink, err := t.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(replyStartTemplate, authLink))

	_, err = t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("cannot send message:%s", err)
	}

	return err
}

func (t *TelegramBot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я незнаю такой команды")
	_, err := t.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("cannot send message:%s", err)
	}
	return err
}
