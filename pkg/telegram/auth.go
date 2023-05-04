package telegram

import (
	"context"
	"fmt"
)

func (t *TelegramBot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := t.generateRedirectURL(chatID)

	requestToken, err := t.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", fmt.Errorf("can not get request token %s", err)
	}

	return t.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (t *TelegramBot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", t.redirectURL, chatID)
}
