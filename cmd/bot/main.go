package main

import (
	"os"

	"github.com/boltdb/bolt"
	"github.com/pushkariov/pocket/pkg/server"
	"github.com/pushkariov/pocket/pkg/storage/boltdb"
	"github.com/pushkariov/pocket/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
)

var telegramToken, pocketKey = os.Getenv("TELEGRAM_KEY"), os.Getenv("POCKET_KEY")

func main() {
	pocketClient, err := pocket.NewClient(pocketKey)
	if err != nil {
		panic(err)
	}

	db, err := bolt.Open("bot.db", 8600, nil)
	if err != nil {
		panic(err)
	}

	tokenStorage, err := boltdb.NewTokenStorage(db)
	if err != nil {
		panic(err)
	}

	telegeamBot := telegram.NewPocketTelegramBot(telegramToken, pocketClient, tokenStorage, "http://localhost")

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenStorage, "https://t.me/newpocketapi_bot")

	go func() {
		if err := telegeamBot.StartBot(); err != nil {
			panic(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		panic(err)
	}
}
