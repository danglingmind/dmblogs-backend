package repository

import (
	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user *entity.User) (*entity.User, error)
	GetByEmailPassword(us *entity.User) (*entity.User, error)
}
