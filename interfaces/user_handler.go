package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"danglingmind.com/ddd/application"
	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

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
	defer func() {
		if err := recover(); err != nil {
			Error(w, http.StatusBadRequest, fmt.Errorf("panic while converting userid"), "provide a valid user id")
		}
	}()

	vars := mux.Vars(r)
	id := uuid.MustParse(vars["id"])

	user, err := u.us.GetById(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	JSON(w, http.StatusOK, user)
	return
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.us.GetAll()
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	JSON(w, http.StatusOK, users)
}

func (u *User) Save(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	// save the user
	_, err = u.us.Save(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	Respond(w, http.StatusOK, "user saved")
}
