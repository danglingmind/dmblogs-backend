// Domain entity and services are combined in entity layer to have better(dot) representation of services
package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"danglingmind.com/ddd/v1/infrastructure/security"
)

type User struct {
	ID         int       `json:"id" type:"int"`
	Firstname  string    `json:"firstname" type:"string"`
	Middlename string    `json:"middlename" type:"string"`
	Lastname   string    `json:"lastname" type:"string"`
	Mobile     string    `json:"mobile" type:"string"`
	Email      string    `json:"email" type:"string"`
	Password   string    `json:"password" type:"string"`
	Created    time.Time `json:"created" type:"time"`
	Modified   time.Time `json:"modified" type:"time"`
	Active     bool      `json:"active" type:"bool"`
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

func (u *User) LoginWithEmailPassword(email, password string) (bool, error) {

	if html.EscapeString(strings.TrimSpace(u.Email)) == u.Email {
		err := security.VerifyPassword(u.Password, password)
		if err != nil {
			return false, errors.New("password didn't match")
		}
		return true, nil
	}
	return false, errors.New("email not found")
}
