package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"manage-jwt/auth"
	"testing"
)

//Auth is created when a user signin.
//The actual signing in of the user is not handled in the model. the model business is simply:
//Give me a data that conform to the types i have established, then i will help you save it in the database.
//Testing of signing in will be done in the controller.
func TestCreateAuth_Success(t *testing.T) {
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
	//lets see the database and get that user and use his id:
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	newAuth, err := server.CreateAuth(user.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, newAuth.ID, 1)
	assert.EqualValues(t, newAuth.UserID, user.ID)
	//since, a random uuid will be created, lets asset that the uuid is not nil:
	assert.NotNil(t, newAuth.AuthUUID)
}

func TestFetchAuth_Success(t *testing.T) {
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
	//lets see the database and get that auth:
	au, err := seedOneAuth()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	fetchAuth := &auth.AuthDetails{
		AuthUuid: "43b78a87-6bcf-439a-ab2e-940d50c4dc33",
		UserId:   1,
	}
	gotAuth, err := server.FetchAuth(fetchAuth)
	assert.Nil(t, err)
	assert.EqualValues(t, gotAuth.ID, 1)
	assert.EqualValues(t, gotAuth.UserID, au.UserID)
	assert.EqualValues(t, gotAuth.AuthUUID, au.AuthUUID)
}

func TestDeleteAuth_Success(t *testing.T) {
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
	//lets see the database and get that auth:
	au, err := seedOneAuth()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	fetchAuth := &auth.AuthDetails{
		AuthUuid: au.AuthUUID,
		UserId:   au.UserID,
	}
	err = server.DeleteAuth(fetchAuth)
	assert.Nil(t, err)
}
