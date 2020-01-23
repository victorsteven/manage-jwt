package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manage-jwt/auth"
	"manage-jwt/model"
	"net/http"
)

var (
	postDB = model.PostDB{}
)

func CreatePost(c *gin.Context) {
	tokenID, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	foundUserUUID, err := authDB.FetchAuth(tokenID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	fmt.Println("the found user UUID: ", foundUserUUID)
	var p model.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	p.UserID = foundUserUUID.UserID
	post, err := postDB.Create(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, post)
}
