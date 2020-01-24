package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateTodo_Success(t *testing.T) {
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
	//Todo created
	todo := &Todo{Title: "Todo title", UserID: user.ID}
	newTodo, err := server.CreateTodo(todo)
	assert.Nil(t, err)
	assert.EqualValues(t, newTodo.ID, 1)
	assert.EqualValues(t, newTodo.Title, todo.Title)
}

//For the following tests, it is not the job of the model to check if the userId is valid or not. That must have checked in the controllers. What the tests do is, just to ensure that data is inserted in the database, to that affect, once we pass in a valid uint64 for the userId, irrespective of what the number is, our test will pass.
func TestCreateTodo_Empty_Todo(t *testing.T) {
	//We will not be hitting the database since the execution will stop when the title is empty
	todo := &Todo{Title: "", UserID: 1}
	newTodo, err := server.CreateTodo(todo)
	assert.NotNil(t, err)
	assert.Nil(t, newTodo)
	assert.EqualValues(t, err.Error(), "please provide a valid title")
}

func TestCreateTodo_No_UserID(t *testing.T) {
	//We will not be hitting the database since the execution will stop when the userId is less than or equal to zero
	todo := &Todo{Title: "the title", UserID: 0}
	newTodo, err := server.CreateTodo(todo)
	assert.NotNil(t, err)
	assert.Nil(t, newTodo)
	assert.EqualValues(t, err.Error(), "a valid user id is required")
}
