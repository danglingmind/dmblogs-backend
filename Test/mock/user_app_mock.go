package mock

import (
	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
)

// UserAppInterface is a mock for actual user app interface
type UserAppInterface struct {
	SaveFn               func(*entity.User) (*entity.User, error)
	GetByIdFn            func(u uuid.UUID) (*entity.User, error)
	GetAllFn             func() ([]entity.User, error)
	GetByEmailPasswordFn func(us *entity.User) (*entity.User, error)
}

func (u *UserAppInterface) Save(user *entity.User) (*entity.User, error) {
	return u.SaveFn(user)
}

func (u *UserAppInterface) GetById(id uuid.UUID) (*entity.User, error) {
	return u.GetByIdFn(id)
}

func (u *UserAppInterface) GetAll() ([]entity.User, error) {
	return u.GetAllFn()
}

func (u *UserAppInterface) GetByEmailPassword(us *entity.User) (*entity.User, error) {
	return u.GetByEmailPasswordFn(us)
}
