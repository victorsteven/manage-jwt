package model

import (
	"errors"
	"github.com/badoux/checkmail"
)

type User struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email string `gorm:"size:255;not null;unique" json:"email"`
}

func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return  errors.New("required email")
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return  errors.New("invalid email")
		}
	}
	return nil
}

func (s *Server) CreateUser(user *User) (*User, error) {
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return nil, emailErr
	}
	err := s.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) GetUserByID(id uint64) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
