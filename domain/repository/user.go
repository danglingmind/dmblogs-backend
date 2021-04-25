package repository

import (
	"context"

	"danglingmind.com/ddd/v1/domain/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id uint64) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, id uint64, values map[string]interface{}) error
	Save(ctx context.Context, user entity.User) error
}
