package boltdb

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pushkariov/pocket/pkg/storage"
)

// TokenStorage is a BoltDB token storage.
type TokenStorage struct {
	db *bolt.DB
}

// NewTokenStorage creates a BoltDB token storage.
func NewTokenStorage(db *bolt.DB) *TokenStorage {
	return &TokenStorage{
		db: db,
	}
}

// SaveToken saves token in the storage.
func (t *TokenStorage) SaveToken(chatID int64, token string, bucket storage.Bucket) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(chatID), []byte(token))
	})
}

// GetToken get token from storage.
func (t *TokenStorage) GetToken(chatID int64, bucket storage.Bucket) (string, error) {
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
