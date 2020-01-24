package model

import (
	"github.com/jinzhu/gorm"
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

func (s *Server) database() (*gorm.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("DB_USER_TEST")
	password := os.Getenv("DB_PASSWORD_TEST")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME_TEST")
	port := os.Getenv("DB_PORT")

	//var err error
	return s.Initialize(dbDriver, username, password, port, host, database)
	//return s.DB, err
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

func seedOneAuth() (*Auth, error) {
	au := &Auth{
		AuthUUID: "43b78a87-6bcf-439a-ab2e-940d50c4dc33",
		UserID:   1,
	}
	err := server.DB.Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}
