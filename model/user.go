package model

import (
	"errors"
	"fmt"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDB struct {
	DB map[uint64]*User
}

//seed one user, just to make sure we always have a user to test with:
func (s *UserDB) SeedUser() {
	user := &User{}
	user.ID = 1
	user.FirstName = "Frank"
	user.LastName = "Abdul"
	user.Email = "frank@gmail.com"
	s.DB = make(map[uint64]*User)
	s.DB[user.ID] = user
	fmt.Println("SEEDER IS CALLED")
}

func (s *UserDB) CreateUser(user *User) (*User, error) {
	//Seed the database first, so that the user created now will have an id of 2 instead of 1
	s.SeedUser()

	user.ID = uint64(len(s.DB) + 1)
	//if no record have been inserted the map yet
	if s.DB == nil {
		s.DB = make(map[uint64]*User)
		s.DB[user.ID] = user
	} else {
		s.DB[user.ID] = user
	}
	return user, nil
}

func (s *UserDB) GetUserByEmail(email string) (*User, error) {
	//Seed the user, then find him
	s.SeedUser()

	foundUser := &User{}
	if s.DB != nil {
		for _, v := range s.DB {
			if v.Email == email {
				foundUser = v
			}
		}
		return foundUser, nil
	} else {
		return nil, errors.New("no account found")
	}
}

func (s *UserDB) GetUserByID(id uint64) (*User, error) {
	foundUser := &User{}
	if s.DB != nil {
		for _, v := range s.DB {
			if v.ID == id {
				foundUser = v
			}
		}
		return foundUser, nil
	} else {
		return nil, errors.New("no account found by the id")
	}
}
