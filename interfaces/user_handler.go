package interfaces

import (
	"encoding/json"
	"net/http"

	"danglingmind.com/ddd/v1/application"
	"danglingmind.com/ddd/v1/domain/entity"
)

// TODO: include auth and token layers as well
type User struct {
	us application.UserAppInterface
}

func NewUser(u application.UserAppInterface) *User {
	return &User{
		us: u,
	}
}

// Define all the handlers for your REST APIs
func (u *User) GetUserById(w http.ResponseWriter, r *http.Request) {
	user, err := u.us.GetById(r.Context(), 123) // TODO: fetch id from url
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "no user found")
		return
	}
	JSON(w, http.StatusOK, user)
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.us.GetAll(r.Context())
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	JSON(w, http.StatusOK, users)
}

func (u *User) Save(w http.ResponseWriter, r *http.Request) {
	user := entity.NewEmptyUser()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	// save the user
	err = u.us.Save(r.Context(), user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	Respond(w, http.StatusOK, "user saved")
}
