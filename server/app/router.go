package app

import (
	"github.com/gin-gonic/contrib/static"
	"jwt-across-platforms/server/controller"
	"jwt-across-platforms/server/middlewares"
)

func route() {
	router.Use(static.Serve("/", static.LocalFile("./web", true))) //for the vue app

	router.GET("/", controller.Index)
	router.POST("/user", controller.CreateUser)
	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/login", controller.Login)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}
