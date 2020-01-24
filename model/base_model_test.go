package model

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}


func (s *Server) database() error {
	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USER_TEST")
	password := os.Getenv("DB_PASSWORD_TEST")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME_TEST")
	port := os.Getenv("DB_PORT")

	err := s.Initialize(dbDriver, username, password, port, host, database)

	return err
}

//Drop test db data if exist:
func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&User{}, &Todo{}, &Auth{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&User{}, &Todo{}, &Auth{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUser() (*User, error) {
	user := &User{
		Email: "frank@gmail.com",
	}
	err := server.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}