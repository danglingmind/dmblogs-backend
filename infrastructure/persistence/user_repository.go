package persistence

import (
	"errors"
	"fmt"
	"strings"

	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
	"danglingmind.com/ddd/infrastructure/security"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *gorm.DB
}

// UserRepo implements user repository
var _ repository.UserRepository = &UserRepo{}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) GetById(id uint64) (*entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (u *UserRepo) Save(user *entity.User) (*entity.User, error) {
	err := user.PrepareToSave()
	if err != nil {
		return nil, err
	}
	err = u.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, fmt.Errorf("email or mobile already taken")
		}
		//any other db error
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetByEmailPassword(us *entity.User) (*entity.User, error) {
	var user entity.User
	err := u.db.Debug().Where("email = ?", us.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, us.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password")
	}
	return &user, nil
}
