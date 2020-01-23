package controller

import (
	"github.com/gin-gonic/gin"
	"manage-jwt/auth"
	"manage-jwt/model"
	"net/http"
)

var (
	authDB = model.AuthDB{}
)

func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//check if the user exist:
	user, err := userDB.GetUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice
	authData, err := authDB.CreateAuth(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var authD auth.AuthDetails
	authD.UserId = authData.UserID
	authD.UserUuid = authData.UserUUID

	token, loginErr := signIn(authD)
	if loginErr != nil {
		c.JSON(http.StatusForbidden, "Please try to login later")
		return
	}
	c.JSON(http.StatusOK, token)
}

func LogOut(c *gin.Context) {
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//foundUserUUID, err := authDB.FetchAuth(tokenID)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, err.Error())
	//	return
	//}

	//if  found the UserUUID, delete it, else, return error
	delErr := authDB.DeleteAuth(tokenID)
	if delErr != nil {
		c.JSON(http.StatusUnauthorized, delErr.Error())
		return
	}
}

func signIn(userUuid auth.AuthDetails) (string, error) {
	token, err := auth.CreateToken(userUuid)
	if err != nil {
		return "", err
	}
	return token, nil
}
