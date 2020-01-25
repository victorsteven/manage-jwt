package app

import (
	"manage-jwt/controller"
	"manage-jwt/middlewares"
)

func route() {
	router.GET("/", controller.Index)
	router.POST("/user", controller.CreateUser)
	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/user/login", controller.Login)
	router.POST("/user/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
