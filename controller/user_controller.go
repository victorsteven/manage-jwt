package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manage-jwt/auth"
	"manage-jwt/model"
	"net/http"
)

var (
	userDB = model.UserDB{}
)

func CreateUser(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	user, err := userDB.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	authData, err := authDB.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.UserUuid = authData.UserUUID

	//Login the user:
	token, loginErr := signIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try to login later")
		return
	}
	fmt.Println("the user: ", user)
	c.JSON(http.StatusCreated, token)
}
