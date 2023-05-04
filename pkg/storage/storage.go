package storage

// Bucket.
type Bucket string

// AccessTokens.
const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

// TokenStorage storage for token.
type TokenStorage interface {
	SaveToken(chatID int64, token string, bucket Bucket) error
	GetToken(chatID int64, bucket Bucket) (string, error)
}
