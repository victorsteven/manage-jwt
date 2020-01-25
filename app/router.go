package app

import (
	"manage-jwt/controller"
	"manage-jwt/middlewares"
)

func route() {
	router.GET("/", controller.Index)
	router.POST("/user", controller.CreateUser)
	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/login", controller.Login)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
