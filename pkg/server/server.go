package server

import (
	"net/http"
	"strconv"

	"github.com/pushkariov/pocket/pkg/storage"
	"github.com/zhashkevych/go-pocket-sdk"
)

// AuthorizationServer.
type AuthorizationServer struct {
	server       *http.Server
	pocketClient *pocket.Client
	tokenStorage storage.TokenStorage
	redirectURL  string
}

// NewAuthorizationServer.
func NewAuthorizationServer(
	pocketClient *pocket.Client,
	tokenStorage storage.TokenStorage,
	redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{
		pocketClient: pocketClient,
		tokenStorage: tokenStorage,
		redirectURL:  redirectURL,
	}
}

// Start the authorization server.
func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":80",
		Handler: s,
	}

	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := s.tokenStorage.GetToken(chatID, storage.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authRest, err := s.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.tokenStorage.SaveToken(chatID, authRest.AccessToken, storage.AccessTokens)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
