package model

import (
	"errors"
	"github.com/twinj/uuid"
)

type Auth struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	UserUUID string `json:"user_uuid"`
}

type AuthDB struct {
	DB map[string]*Auth
}

func SeedAuth() {
	auth := Auth{}
	auth.UserID = 1
	auth.UserUUID = uuid.NewV4().String() //create a new uuid
}

func (au AuthDB) FetchAuth(userUuid string) (*Auth, error) {
	SeedAuth() //let it be that we have a record in the database
	foundAuth := &Auth{}
	if au.DB != nil {
		for _, v := range au.DB {
			if v.UserUUID == userUuid {
				foundAuth = v
			}
		}
	} else {
		return nil, errors.New("you are not authorized")
	}
	if foundAuth.UserID == 0 {
		return nil, errors.New("you are not authorized, token not valid")
	}
	return foundAuth, nil
}

//Once a user is logged out, delete the uuid of that particular user
func (au AuthDB) DeleteAuth(userUuid string) error {
	foundAuth := &Auth{}
	if au.DB != nil {
		for _, v := range au.DB {
			if v.UserUUID == userUuid {
				foundAuth = v
			}
		}
		if foundAuth != nil {
			delete(au.DB, foundAuth.UserUUID) //remove from the database
		}
	} else {
		return errors.New("you are not authorized")
	}
	return nil
}

//once the user signup/login, create a row in the auth table, with a new uuid
func (au *AuthDB) CreateAuth(userId uint64) (*Auth, error) {
	SeedAuth() //it be that we have a user earlier registered

	auth := &Auth{}
	auth.ID = uint64(len(au.DB) + 1) //increment the number
	auth.UserID = userId
	auth.UserUUID = uuid.NewV4().String() //create a new uuid

	//if no record have been inserted the map yet
	if au.DB == nil {
		au.DB = make(map[string]*Auth)
		au.DB[auth.UserUUID] = auth
	} else {
		au.DB[auth.UserUUID] = auth
	}
	return auth, nil
}
