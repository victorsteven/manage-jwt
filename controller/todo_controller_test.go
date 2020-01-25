package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"manage-jwt/auth"
	"manage-jwt/model"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCreateTodo_Success(t *testing.T) {
	//Now exchange the real implementation with our mock
	model.Model = &fakeServer{}

	fetchAuth = func(*auth.AuthDetails) (*model.Auth, error) {
		return &model.Auth{
			ID:    1,
			UserID: 1,
			AuthUUID: "83b09612-9dfc-4c1d-8f7d-a589acec7081",
		}, nil
	}
	createTodoModel = func(*model.Todo) (*model.Todo, error) {
		return &model.Todo{
			ID:     1,
			UserID: 1,
			Title:  "the title",
		}, nil
	}
	todo := model.Todo{
		Title: "the title",
	}
	byteSlice, err := json.Marshal(&todo)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	//It is an authenticated user can create a todo, so, lets pass a token to our request headers
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkduc"
	tokenString := fmt.Sprintf("Bearer %v", tk)
	req.Header.Set("Authorization", tokenString)

	rr := httptest.NewRecorder()
	r.POST("/todo", CreateTodo)
	r.ServeHTTP(rr, req)

	var newTodo model.Todo
	err = json.Unmarshal(rr.Body.Bytes(), &newTodo)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusCreated, rr.Code)
	assert.EqualValues(t, 1, newTodo.UserID)
	assert.EqualValues(t, "the title", newTodo.Title)
}

//Anything from empty or wrong token returns unauthorized. From the example, we used a wrong token. we added wrong letters at the end.
//Since for sure a todo will not be created, we avoided mocking the fetchAuth and the createTodoModel methods.
func TestCreateTodo_Unauthorized_User(t *testing.T) {
	todo := model.Todo{
		Title: "the title",
	}
	byteSlice, err := json.Marshal(&todo)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	//It is an authenticated user can create a todo, so, lets pass a token to our request headers
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkducxx"
	tokenString := fmt.Sprintf("Bearer %v", tk)
	req.Header.Set("Authorization", tokenString)

	rr := httptest.NewRecorder()
	r.POST("/todo", CreateTodo)
	r.ServeHTTP(rr, req)

	var errMsg string
	err = json.Unmarshal(rr.Body.Bytes(), &errMsg)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, rr.Code)
	assert.EqualValues(t, "unauthorized", errMsg)
}

//When wrong input is supplied. Here also, we wont mock any external methods
func TestCreateTodo_Invalid_Input(t *testing.T) {
	invalidTitle := 12345 //using an integer instead of a string
	byteSlice, err := json.Marshal(&invalidTitle)
	if err != nil {
		t.Error("Cannot marshal to json")
	}
	r := gin.Default()
	req, err := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(byteSlice))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/todo", CreateTodo)
	r.ServeHTTP(rr, req)

	var msg string
	err = json.Unmarshal(rr.Body.Bytes(), &msg) //since we outputted the error as string in the controller
	assert.Nil(t, err) //we can unmarshall without issues
	assert.EqualValues(t, "invalid json", msg)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}