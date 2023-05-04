package main

import (
	"os"

	"github.com/pushkariov/pocket/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
)

var telegramToken = os.Getenv("TELEGRAM_KEY")

func main() {
	pocketClient, err := pocket.NewClient("106603-d34dd30dc8bf1509e1fbbd1")
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
