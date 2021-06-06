package interfaces

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/infrastructure/auth"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	userUUID := uuid.New()
	// mock UserApp methods
	userAppMock.SaveFn = func(u *entity.User) (*entity.User, error) {
		return &entity.User{
			ID:          userUUID,
			Firstname:   "prateek",
			Middlename:  "",
			Lastname:    "reddy",
			Mobile:      "9876543210",
			Email:       "dummy@dummy.com",
			Countrycode: "+91",
			Password:    "dummydummydummy",
		}, nil
	}
	// mock createToken interface
	tokenMock.CreateTokenFn = func(u uuid.UUID) (*auth.TokenDetails, error) {
		return &auth.TokenDetails{
			AccessToken:  "uniquetoken",
			RefreshToken: "refreshtoken",
			TokenUuid:    "abcd",
			RefreshUuid:  "efgh",
			AtExpires:    1234,
			RtExpires:    3456,
		}, nil
	}

	// mock createAuth method
	authMock.CreateAuthFn = func(u uuid.UUID, td *auth.TokenDetails) error {
		return nil
	}

	router := mux.NewRouter()
	router.HandleFunc("/register", au.Register).Methods(http.MethodPut)

	inputJSON := `{
		"firstname":"prateek",
		"lastname": "reddy",
		"mobile":"9876543210",
		"email":"dummy@dummy.com",
		"password":"dummydummydummy"
	}`
	req, err := http.NewRequest(http.MethodPut, "/register", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("internal test error Register user handler: %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	userData := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &userData)

	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "uniquetoken", userData["access_token"])
	assert.Equal(t, "refreshtoken", userData["refresh_token"])
	assert.Equal(t, userUUID.String(), userData["id"])
	assert.Equal(t, "prateek", userData["first_name"])
	assert.Equal(t, "reddy", userData["last_name"])
	assert.Equal(t, "", userData["middle_name"])
}

func TestLoginUser_Success(t *testing.T) {
	// mock data
	userUUID := uuid.New()
	userAppMock.GetByEmailPasswordFn = func(us *entity.User) (*entity.User, error) {
		return &entity.User{
			ID:          userUUID,
			Firstname:   "prateek",
			Middlename:  "",
			Lastname:    "reddy",
			Mobile:      "9876543210",
			Email:       "dummy@dummy.com",
			Countrycode: "+91",
			Password:    "dummydummydummy",
		}, nil
	}

	tokenMock.CreateTokenFn = func(u uuid.UUID) (*auth.TokenDetails, error) {
		return &auth.TokenDetails{
			AccessToken:  "uniquetoken",
			RefreshToken: "refreshtoken",
			TokenUuid:    "abcd",
			RefreshUuid:  "efgh",
			AtExpires:    1234,
			RtExpires:    3456,
		}, nil
	}

	authMock.CreateAuthFn = func(u uuid.UUID, td *auth.TokenDetails) error {
		return nil
	}

	inputJSON := `{
		"email":"dummy@dummy.com",
		"password":"dummydummydummy"
	}`

	router := mux.NewRouter()
	router.HandleFunc("/users/login", au.Login).Methods(http.MethodPost)
	req, err := http.NewRequest(http.MethodPost, "/users/login", bytes.NewBufferString(inputJSON))
	// bearer := "bearer " + "uniquetoken"
	// req.Header.Add("Authorization", bearer)
	if err != nil {
		t.Errorf("Unit test LoginHandler failed with some internal error %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	loginResp := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &loginResp)

	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "uniquetoken", loginResp["access_token"])
	assert.Equal(t, "refreshtoken", loginResp["refresh_token"])
	assert.Equal(t, userUUID.String(), loginResp["id"])
	assert.Equal(t, "prateek", loginResp["first_name"])
	assert.Equal(t, "", loginResp["middle_name"])
	assert.Equal(t, "reddy", loginResp["last_name"])
}
