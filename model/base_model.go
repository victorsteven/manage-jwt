package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"manage-jwt/auth"
)

type Server struct {
	DB *gorm.DB
}

var (
	//Server now implements the modelInterface, so he can define its methods
	Model modelInterface = &Server{}
)

type modelInterface interface {
	Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error)

	//user methods
	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)
	GetUserByID(uint64) (*User, error)

	//todo methods:
	CreateTodo(*Todo) (*Todo, error)


	//auth methods:
	FetchAuth(*auth.AuthDetails) (*Auth, error)
	DeleteAuth(*auth.AuthDetails) error
	CreateAuth(uint64) (*Auth, error)
}

func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	s.DB.Debug().AutoMigrate(
		&User{},
		&Auth{},
		&Todo{},
	)
	return s.DB, nil
}
