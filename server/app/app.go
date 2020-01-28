package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"jwt-across-platforms/server/model"
	"os"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
}
var router = gin.Default()

func  StartApp() {
	//var conn Connect
	dbdriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	_, err := model.Model.Initialize(dbdriver, username, password, db_port, host, database)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	route()

	port := os.Getenv("PORT") //using heroku host
	if port == "" {
		port = "8888" //localhost
	}
	log.Fatal(router.Run(":"+port))
}
