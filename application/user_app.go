package application

import (
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
	"github.com/google/uuid"
)

// this is the only api which is exposed to the interfaces to communicate with business elements
// like entity or any other business service
type UserAppInterface interface {
	GetById(id uuid.UUID) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user *entity.User) (*entity.User, error)
	GetByEmailPassword(us *entity.User) (*entity.User, error)
}

type UserApp struct {
	user repository.UserRepository // connect with the infrastructure layer
}

var _ UserAppInterface = &UserApp{}

func (u *UserApp) GetAll() ([]entity.User, error) {
	return u.user.GetAll()
}

func (u *UserApp) GetById(id uuid.UUID) (*entity.User, error) {
	return u.user.GetById(id)
}

func (u *UserApp) Save(user *entity.User) (*entity.User, error) {
	return u.user.Save(user)
}

func (u *UserApp) GetByEmailPassword(us *entity.User) (*entity.User, error) {
	return u.user.GetByEmailPassword(us)
}
