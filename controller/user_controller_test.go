package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"manage-jwt/auth"
	"manage-jwt/model"
	"manage-jwt/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

//We will mock the domain methods so to as to achieve unit test in our controllers
var (
	createUserModel func(*model.User) (*model.User, error)
	createAuthModel func(uint64) (*model.Auth, error)
	signIn func(auth.AuthDetails) (string, error)
)

type fakeServer struct {}

func (fs *fakeServer) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	panic("implement me")
}
func (fs *fakeServer) ValidateEmail(string) error {
	panic("implement me")
}
func (fs *fakeServer) GetUserByEmail(string) (*model.User, error) {
	panic("implement me")
}
func (fs *fakeServer) GetUserByID(uint64) (*model.User, error) {
	panic("implement me")
}
func (fs *fakeServer) CreateTodo(*model.Todo) (*model.Todo, error) {
	panic("implement me")
}
func (fs *fakeServer) FetchAuth(*auth.AuthDetails) (*model.Auth, error) {
	panic("implement me")
}
func (fs *fakeServer) DeleteAuth(*auth.AuthDetails) error {
	panic("implement me")
}

func (fs *fakeServer) CreateUser(user *model.User) (*model.User, error) {
	return createUserModel(user)
}
func (fs *fakeServer) CreateAuth(userId uint64) (*model.Auth, error) {
	return createAuthModel(userId)
}


type fakeSignin struct {}

func (fs *fakeSignin) SignIn(authD auth.AuthDetails) (string, error) {
	return signIn(authD)
}


func TestCreateUser_Success(t *testing.T) {
	model.Model = &fakeServer{}
	service.Authorize = &fakeSignin{}

	createUserModel = func(*model.User) (*model.User, error) {
		return &model.User{
			ID:    1,
			Email: "sunflash@gmail.com",
		}, nil
	}
	createAuthModel  = func(uint64) (*model.Auth, error) {
		return &model.Auth{
			ID:       1,
			UserID:   1,
			AuthUUID: "83b09612-9dfc-4c1d-8f7d-a589acec7081",
		}, nil
	}
	signIn  = func(auth.AuthDetails) (string, error) {
		return  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiI4M2IwOTYxMi05ZGZjLTRjMWQtOGY3ZC1hNTg5YWNlYzcwODEiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo1fQ.otegNS-W9OE8RsqGtiyJRCB-H0YXBygNXP91qeCPdF8", nil
	}

	//Now let test only the controller implementation,  void of external methods. Remember, the end result when the function runs to to return a JWT. And that JWT that will be returned is the one we have defined above.
	u := model.User{
		Email: "vicsdfddt@gmail.com",
	}
	byteSlice, err := json.Marshal(&u)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/user", CreateUser)
	r.ServeHTTP(rr, req)

	var token string
	err = json.Unmarshal(rr.Body.Bytes(), &token)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusCreated, rr.Code)
	assert.EqualValues(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiI4M2IwOTYxMi05ZGZjLTRjMWQtOGY3ZC1hNTg5YWNlYzcwODEiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo1fQ.otegNS-W9OE8RsqGtiyJRCB-H0YXBygNXP91qeCPdF8", token)
}

//We dont need to mock anything here, since our execution will never call the external methods
//Now use an integer instead of a string for the input email
func TestCreateUser_Invalid_Input(t *testing.T) {
	invalidEmail := 12345 //using an integer instead of a string
	byteSlice, err := json.Marshal(&invalidEmail)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/user", CreateUser)
	r.ServeHTTP(rr, req)

	var msg string
	err = json.Unmarshal(rr.Body.Bytes(), &msg) //since we outputted the error as string in the controller
	assert.Nil(t, err) //we can unmarshall without issues
	assert.EqualValues(t, "invalid json", msg)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}