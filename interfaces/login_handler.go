package interfaces

import (
	"encoding/json"
	"net/http"

	"danglingmind.com/ddd/application"
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/infrastructure/auth"
)

type Authenticate struct {
	userApp        application.UserAppInterface
	authInterface  auth.AuthInterface
	tokenInterface auth.TokenInterface
}

func NewAuthenticator(uApp application.UserAppInterface, au auth.AuthInterface, tk auth.TokenInterface) *Authenticate {
	return &Authenticate{
		userApp:        uApp,
		authInterface:  au,
		tokenInterface: tk,
	}
}

func (au *Authenticate) Login(w http.ResponseWriter, r *http.Request) {
	// get the user
	u := entity.NewEmptyUser()
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "bad request body")
		return
	}
	// validate user json
	if ok, err := u.Validate(); !ok {
		Error(w, http.StatusMethodNotAllowed, err, err.Error())
		return
	}
	// validate credentials
	user, err := au.userApp.GetByEmailPassword(r.Context(), &u)
	if err != nil {
		Error(w, http.StatusUnauthorized, err, err.Error())
		return
	}

	// create new token
	tk, err := au.tokenInterface.CreateToken(uint64(user.ID))
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	err = au.authInterface.CreateAuth(uint64(user.ID), tk)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	userData := make(map[string]interface{})
	userData["access_token"] = tk.AccessToken
	userData["refresh_token"] = tk.RefreshToken
	userData["id"] = user.ID
	userData["first_name"] = user.Firstname
	userData["last_name"] = user.Lastname
	userData["middle_name"] = user.Middlename
	JSON(w, http.StatusOK, userData)
}

func (au *Authenticate) Logout(w http.ResponseWriter, r *http.Request) {
	metadata, err := au.tokenInterface.ExtractTokenMetadata(r)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	err = au.authInterface.DeleteTokens(metadata)
	if err != nil {
		Error(w, http.StatusUnauthorized, err, err.Error())
		return
	}
	JSON(w, http.StatusOK, "logged out")
}

// TODO: implement this method
func (au *Authenticate) Refresh(w http.ResponseWriter, r *http.Request) {

}
