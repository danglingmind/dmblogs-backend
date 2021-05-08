package persistence

import (
	"context"

	"danglingmind.com/ddd/domain/repository"
)

type Repositories struct {
	User repository.UserRepository
}

// initialize all the domain repositories with respective DBs
func NewRepository(host, username, password, dbname string, port int) (*Repositories, error) {
	// initialize Mysql repository
	msCtx := context.Background()
	msStore := NewMySqlStore()
	err := msStore.Open(msCtx, host, username, password, dbname, port)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: NewUserRepository(msStore),
	}, nil

}
