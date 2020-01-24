package controller

import (
	"github.com/gin-gonic/gin"
	"manage-jwt/auth"
	"manage-jwt/model"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	errEmail := model.Model.ValidateEmail(u.Email)
	if errEmail != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid email given")
		return
	}
	user, err := model.Model.CreateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	authData, err := model.Model.CreateAuth(user.ID)
	if err != nil {
		//since we are dealing with only email, the common error we be "email already exist, if you have more field, please dont hard this this error message as i do below:
		c.JSON(http.StatusInternalServerError, "email already taken")
		return
	}

	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.AuthUuid = authData.AuthUUID

	//Login the user:
	token, loginErr := signIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "please try to login later")
		return
	}
	c.JSON(http.StatusCreated, token)
}
