package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pushkariov/pocket/pkg/storage"
	"github.com/zhashkevych/go-pocket-sdk"
)

// PocketTelegramBot is a telegram bot for pocket.
type PocketTelegramBot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	tokenStorage storage.TokenStorage
	redirectURL  string
}

// NewPocketTelegramBot  creates telegram bot.
func NewPocketTelegramBot(
	telegramToken string,
	pocketClient *pocket.Client,
	tokenStorage storage.TokenStorage,
	redirectURL string) *PocketTelegramBot {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	return &PocketTelegramBot{
		bot:          bot,
		pocketClient: pocketClient,
		tokenStorage: tokenStorage,
		redirectURL:  redirectURL,
	}
}

// StartBot start bot.
func (t *PocketTelegramBot) StartBot() error {
	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	updates := t.initUpdatesChannel()

	if err := t.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (t *PocketTelegramBot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
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

func (t *PocketTelegramBot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return t.bot.GetUpdatesChan(updateConfig)
}
