package app

import (
	"manage-jwt/controller"
	"manage-jwt/middlewares"
)

func route() {
	router.POST("/user", controller.CreateUser)
	router.POST("/post", middlewares.TokenAuthMiddleware(), controller.CreatePost)
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
