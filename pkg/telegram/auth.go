package telegram

import (
	"context"
	"fmt"

	"github.com/pushkariov/pocket/pkg/storage"
)

func (t *PocketTelegramBot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := t.generateRedirectURL(chatID)

	requestToken, err := t.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", fmt.Errorf("can not get request token %s", err)
	}

	if err := t.tokenStorage.SaveToken(chatID, requestToken, storage.RequestTokens); err != nil {
		return "", err
	}

	return t.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (t *PocketTelegramBot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", t.redirectURL, chatID)
}
