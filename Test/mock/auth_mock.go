package mock

import (
	"danglingmind.com/ddd/infrastructure/auth"
	"github.com/google/uuid"
)

type AuthInterface struct {
	CreateAuthFn    func(uuid.UUID, *auth.TokenDetails) error
	FetchAuthFn     func(string) (uuid.UUID, error)
	DeleteRefreshFn func(string) error
	DeleteTokensFn  func(*auth.AccessDetails) error
}

func (a *AuthInterface) CreateAuth(userid uuid.UUID, td *auth.TokenDetails) error {
	return a.CreateAuthFn(userid, td)
}

func (a *AuthInterface) FetchAuth(tokenUuid string) (uuid.UUID, error) {
	return a.FetchAuthFn(tokenUuid)
}
func (a *AuthInterface) DeleteRefresh(refreshUuid string) error {
	return a.DeleteRefreshFn(refreshUuid)
}
func (a *AuthInterface) DeleteTokens(authD *auth.AccessDetails) error {
	return a.DeleteTokensFn(authD)
}
