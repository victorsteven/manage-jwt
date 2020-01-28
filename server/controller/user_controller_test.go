package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"jwt-across-platforms/server/model"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCreateUser_Success(t *testing.T) {

	model.Model = &fakeServer{} //this is where the swapping of the real method with the fake one

	createUserModel = func(*model.User) (*model.User, error) {
		return &model.User{
			ID:    1,
			Email: "sunflash@gmail.com",
		}, nil
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

	var user model.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusCreated, rr.Code)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "sunflash@gmail.com", user.Email)
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