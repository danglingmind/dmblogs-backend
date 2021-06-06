// Domain entity and services are combined in entity layer to have better(dot) representation of services
package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"danglingmind.com/ddd/infrastructure/security"
	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Firstname   string    `json:"firstname" gorm:"size:100;not null;"`
	Middlename  string    `json:"middlename" gorm:"size:100;"`
	Lastname    string    `json:"lastname" gorm:"size:100;"`
	Countrycode string    `json:"countrycode" type:"string"`
	Mobile      string    `json:"mobile" gorm:"size:10;"`
	Email       string    `json:"email" gorm:"size:100;not null;unique"`
	Password    string    `json:"password" gorm:"size:100;not null;"`
	Created     time.Time `json:"created" gorm:"default:CURRENT_TIMESTAMP"`
	Modified    time.Time `json:"modified" gorm:"default:CURRENT_TIMESTAMP"`
	Active      bool      `json:"active" gorm:"default:1;"`
}

func NewEmptyUser() User {
	return User{}
}

// gorm hooks
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	u.Active = true // Set default active to active (true)
	u.Created = time.Now()
	u.Modified = time.Now()
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Middlename = html.EscapeString(strings.TrimSpace(u.Middlename))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	// Process password
	passEncrypted, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(passEncrypted)

	_, err = u.Validate()

	return
}

func (u *User) Validate() (bool, error) {

	if u.Password == "" {
		return false, fmt.Errorf("password is required")
	}
	if u.Email == "" {
		return false, fmt.Errorf("email is required")
	}
	if u.Email != "" {
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return false, fmt.Errorf("please provide a valid email")
		}
	}
	return true, nil
}
