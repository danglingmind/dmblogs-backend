package mock

import (
	"net/http"

	"danglingmind.com/ddd/infrastructure/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenInterface struct {
	CreateTokenFn          func(uuid.UUID) (*auth.TokenDetails, error)
	ExtractTokenMetadataFn func(*http.Request) (*auth.AccessDetails, error)
	TokenValidFn           func(*http.Request) error
	VerifyTokenFn          func(*http.Request) (*jwt.Token, error)
	ExtractTokenFn         func(r *http.Request) string
}

func (t *TokenInterface) CreateToken(userid uuid.UUID) (*auth.TokenDetails, error) {
	return t.CreateTokenFn(userid)
}

func (t *TokenInterface) ExtractTokenMetadata(r *http.Request) (*auth.AccessDetails, error) {
	return t.ExtractTokenMetadataFn(r)
}

func (t *TokenInterface) TokenValid(r *http.Request) error {
	return t.TokenValidFn(r)
}

func (t *TokenInterface) VerifyToken(r *http.Request) (*jwt.Token, error) {
	return t.VerifyTokenFn(r)
}

func (t *TokenInterface) ExtractToken(r *http.Request) string {
	return t.ExtractTokenFn(r)
}
