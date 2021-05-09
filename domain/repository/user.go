package repository

import (
	"danglingmind.com/ddd/domain/entity"
)

type UserRepository interface {
	GetById(id uint64) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user *entity.User) (*entity.User, error)
	GetByEmailPassword(us *entity.User) (*entity.User, error)
}
