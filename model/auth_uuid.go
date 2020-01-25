package model

import (
	"github.com/twinj/uuid"
	"manage-jwt/auth"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

func (s *Server) FetchAuth(authD *auth.AuthDetails) (*Auth, error) {
	au := &Auth{}
	err := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

//Once a user row in the auth table
func (s *Server) DeleteAuth(authD *auth.AuthDetails) error {
	au := &Auth{}
	db := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Once the user signup/login, create a row in the auth table, with a new uuid
func (s *Server) CreateAuth(userId uint64) (*Auth, error) {
	au := &Auth{}
	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.UserID = userId
	err := s.DB.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}
