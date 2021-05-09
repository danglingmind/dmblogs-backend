// Domain entity and services are combined in entity layer to have better(dot) representation of services
package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"danglingmind.com/ddd/infrastructure/security"
	"github.com/badoux/checkmail"
)

type User struct {
	ID          int       `json:"id" type:"int"`
	Firstname   string    `json:"firstname" type:"string"`
	Middlename  string    `json:"middlename" type:"string"`
	Lastname    string    `json:"lastname" type:"string"`
	Countrycode string    `json:"countrycode" type:"string"`
	Mobile      string    `json:"mobile" type:"string"`
	Email       string    `json:"email" type:"string"`
	Password    string    `json:"password" type:"string"`
	Created     time.Time `json:"created" type:"time"`
	Modified    time.Time `json:"modified" type:"time"`
	Active      bool      `json:"active" type:"bool"`
}

func NewEmptyUser() User {
	return User{}
}

func (u *User) PrepareToSave() (err error) {
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

	return nil
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
