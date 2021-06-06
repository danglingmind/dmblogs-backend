package interfaces

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetById_Success(t *testing.T) {
	userUUID := uuid.New()
	// mock response for GetById UserAppInterface
	userAppMock.GetByIdFn = func(u uuid.UUID) (*entity.User, error) {
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

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", us.GetUserById).Methods(http.MethodGet)

	req, err := http.NewRequest(http.MethodGet, "/users/"+userUUID.String(), nil)
	if err != nil {
		t.Errorf("test failed getById: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	var user *entity.User

	err = json.Unmarshal(rr.Body.Bytes(), &user)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, rr.Code)
	assert.EqualValues(t, "prateek", user.Firstname)
	assert.EqualValues(t, "reddy", user.Lastname)
}

func TestGetAllUsers(t *testing.T) {
	uId1 := uuid.New()
	uId2 := uuid.New()
	// mock response for GetAll from UserApp
	userAppMock.GetAllFn = func() ([]entity.User, error) {
		return []entity.User{
			{
				ID:          uId1,
				Firstname:   "prateek",
				Middlename:  "",
				Lastname:    "reddy",
				Mobile:      "9876543210",
				Email:       "jiriaya@dummy.com",
				Countrycode: "+90",
				Password:    "dummydummydummy",
			},
			{
				ID:          uId2,
				Firstname:   "jiraiya",
				Middlename:  "",
				Lastname:    "sama",
				Mobile:      "9876543211",
				Email:       "dummy@dummy.com",
				Countrycode: "+91",
				Password:    "dummydummydummy",
			},
		}, nil
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", us.GetAllUsers).Methods(http.MethodGet)

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Errorf("test failed handler: GetAllUsers: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	var users []entity.User

	err = json.Unmarshal(rr.Body.Bytes(), &users)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, rr.Code)
	assert.EqualValues(t, uId1, users[0].ID)
	assert.EqualValues(t, "prateek", users[0].Firstname)
	assert.EqualValues(t, "reddy", users[0].Lastname)
	assert.EqualValues(t, uId2, users[1].ID)
	assert.EqualValues(t, "jiraiya", users[1].Firstname)
	assert.EqualValues(t, "sama", users[1].Lastname)
}
