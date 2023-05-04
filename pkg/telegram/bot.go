package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

// TelegramBot
type TelegramBot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	redirectURL  string
}

// NewTelegramBot creates telegram bot.
func NewTelegramBot(telegramToken string, pocketClient *pocket.Client, redirectURL string) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	return &TelegramBot{
		bot:          bot,
		pocketClient: pocketClient,
		redirectURL:  redirectURL,
	}
}

// StartBot.
func (t *TelegramBot) StartBot() error {
	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	updates := t.initUpdatesChannel()

	if err := t.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (t *TelegramBot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := t.handleCommand(update.Message); err != nil {
				return err
			}
			continue
		}

		t.handleMessage(update.Message)
	}

	return nil
}

func (t *TelegramBot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return t.bot.GetUpdatesChan(updateConfig)
}
