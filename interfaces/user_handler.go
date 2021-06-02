package interfaces

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"danglingmind.com/ddd/application"
	"danglingmind.com/ddd/domain/entity"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	} else {
		user, err := u.us.GetById(uint64(id))
		if err != nil {
			Error(w, http.StatusInternalServerError, err, err.Error())
			return
		}
		JSON(w, http.StatusOK, user)
		return
	}

}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.us.GetAll()
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	log.Println("testing logger from log package")
	logrus.Info("testing from logrus")
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
