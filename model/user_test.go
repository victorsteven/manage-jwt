package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestValidateEmail_Success(t *testing.T) {
	//Correct email
	email := "stevensunflash@gmail.com"
	err := server.ValidateEmail(email)
	assert.Nil(t, err)
}

//Using table test to check the two failures at once
func TestValidateEmail_Failure(t *testing.T) {
	samples := []struct {
		email            string
		errMsgInvalid    string
		errMsgEmptyEmail string
	}{
		{
			//Invalid email
			email:         "stevensunflash.com",
			errMsgInvalid: "invalid email",
		},
		{
			//Empty email
			email:            "",
			errMsgEmptyEmail: "required email",
		},
	}
	for _, v := range samples {
		err := server.ValidateEmail(v.email)

		assert.NotNil(t, err) //there must be an error in either case

		if err != nil && v.errMsgInvalid != "" {
			assert.EqualValues(t, v.errMsgInvalid, "invalid email")
		}
		if err != nil && v.errMsgEmptyEmail != "" {
			assert.EqualValues(t, v.errMsgEmptyEmail, "required email")
		}
	}
}

func TestCreateUser_Success(t *testing.T) {
	//Initialize DB:
	var err error
	server.DB, err = server.database()
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}
	defer server.DB.Close()
	err = refreshUserTable()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	//User created
	user := &User{Email: "stevensunflash@gmail.com"}
	u, err := server.CreateUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.ID, 1)
	assert.EqualValues(t, u.Email, "stevensunflash@gmail.com")
}

func TestCreateUser_Duplicate_Email(t *testing.T) {
	//Initialize DB:
	var err error
	server.DB, err = server.database()
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}
	defer server.DB.Close()
	err = refreshUserTable()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	_, err = seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed user: %v", err)
	}
	//remember we have seeded this user, so we want to insert him again, it should fail
	userRequest := &User{Email: "frank@gmail.com"}
	u, err := server.CreateUser(userRequest)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Contains(t, err.Error(), "duplicate")
}

//We will test only for success here, you can write failure cases if you have time, and also to improve ur code coverage
func TestGetUserByEmail_Success(t *testing.T) {
	//Initialize DB:
	var err error
	server.DB, err = server.database()
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}
	defer server.DB.Close()

	err = refreshUserTable()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed user: %v", err)
	}
	email := "frank@gmail.com"
	getUser, err := server.GetUserByEmail(email)
	assert.Nil(t, err)
	assert.EqualValues(t, getUser.Email, user.Email)
}

//We will test only for success here, you can write failure cases if you have time, and also to improve ur code coverage
func TestGetUserByID_Success(t *testing.T) {
	//Initialize DB:
	var err error
	server.DB, err = server.database()
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}
	defer server.DB.Close()
	err = refreshUserTable()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed user: %v", err)
	}
	userId := uint64(1) //convert int to uint64
	getUser, err := server.GetUserByID(userId)
	assert.Nil(t, err)
	assert.EqualValues(t, getUser.ID, user.ID)
}
