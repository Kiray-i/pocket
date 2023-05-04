package boltdb

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pushkariov/pocket/pkg/storage"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (t *TokenRepository) Save(chatID int64, token string, bucket storage.Bucket) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

func (t *TokenRepository) Get(chatID int64, bucket storage.Bucket) (string, error) {
	var token string

	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToBytes(chatID))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("can not get token from storage: %s", err)
	}

	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
