package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"manage-jwt/auth"
	"manage-jwt/model"
	"manage-jwt/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

//NOTE WE ARE PERFORMING UNIT TESTS ON THE LOGIN FUNCTION, SO WE MOCKED ALL FUNCTIONS/METHODS THAT THE LOGIN DEPEND. CHECK OUT THE FILE "controller_mock_setup.go" FILE TO SEE HOW THE MOCK IS CREATED AND USED HERE

func TestLogin_Success(t *testing.T) {
	model.Model = &fakeServer{}
	service.Authorize = &fakeSignin{}

	getUserByEmail = func(email string) (*model.User, error) {
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
	req, err := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/user/login", Login)
	r.ServeHTTP(rr, req)

	var token string
	err = json.Unmarshal(rr.Body.Bytes(), &token)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiI4M2IwOTYxMi05ZGZjLTRjMWQtOGY3ZC1hNTg5YWNlYzcwODEiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjo1fQ.otegNS-W9OE8RsqGtiyJRCB-H0YXBygNXP91qeCPdF8", token)
}

//An example is when the email is not found in the database.
//We only mock according to demand. In the test below, we mocked only the GetUserEmail method, since execution will stop there, because we told it to return not found
func TestLogin_Not_Found_User(t *testing.T) {
	model.Model = &fakeServer{}
	service.Authorize = &fakeSignin{}

	getUserByEmail = func(email string) (*model.User, error) {
		return nil, errors.New("email not found")
	}
	u := model.User{
		Email: "vicsdfddt@gmail.com",
	}
	byteSlice, err := json.Marshal(&u)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/user/login", Login)
	r.ServeHTTP(rr, req)

	var errString string
	err = json.Unmarshal(rr.Body.Bytes(), &errString)
	assert.Nil(t, err)
	assert.NotNil(t, errString)
	assert.EqualValues(t, http.StatusNotFound, rr.Code)
	assert.EqualValues(t, "email not found", errString)
}