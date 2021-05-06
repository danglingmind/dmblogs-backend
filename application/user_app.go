package application

import (
	"context"

	"danglingmind.com/ddd/v1/domain/entity"
	"danglingmind.com/ddd/v1/domain/repository"
)

// this is the only api which is exposed to the interfaces to communicate with business elements
// like entity or any other business service
type UserAppInterface interface {
	GetAll(context.Context) ([]entity.User, error)
	GetById(context.Context, uint64) (*entity.User, error)
	Save(context.Context, entity.User) error
	Update(ctx context.Context, id uint64, values map[string]interface{}) error
	GetByEmailPassword(ctx context.Context, us *entity.User) (*entity.User, error)
}

type UserApp struct {
	user repository.UserRepository // connect with the infrastructure layer
}

var _ UserAppInterface = &UserApp{}

// TODO: have the correct context here
func (u *UserApp) GetAll(ctx context.Context) ([]entity.User, error) {
	return u.user.GetAll(ctx)
}

func (u *UserApp) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	return u.user.GetById(ctx, id)
}

func (u *UserApp) Save(ctx context.Context, user entity.User) error {
	return u.user.Save(ctx, user)
}

func (u *UserApp) Update(ctx context.Context, id uint64, values map[string]interface{}) error {
	return nil
}

func (u *UserApp) GetByEmailPassword(ctx context.Context, us *entity.User) (*entity.User, error) {
	return u.user.GetByEmailPassword(ctx, us)
}
