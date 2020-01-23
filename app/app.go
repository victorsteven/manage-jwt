package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApp() {

	route()

	log.Fatal(router.Run(":8080"))

}
