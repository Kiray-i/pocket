package main

import (
	"os"

	"github.com/pushkariov/pocket/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
)

var telegramToken, pocketKey = os.Getenv("TELEGRAM_KEY"), os.Getenv("POCKET_KEY")

func main() {
	pocketClient, err := pocket.NewClient(pocketKey)
	if err != nil {
		panic(err)
	}

	// db, err := bolt.Open("bot.db", 8600, nil)
	// if err != nil {
	// 	panic(err)
	// }

	telegeamBot := telegram.NewPocketTelegramBot(telegramToken, pocketClient, "http://localhost")
	if err := telegeamBot.StartBot(); err != nil {
		panic(err)
	}
}
