package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateTodo_Success(t *testing.T) {
	//Initialize DB:
	err := server.database()
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}
	err = refreshUserTable()
	if err != nil {
		log.Fatalf("cannot refresh db tables: %v", err)
	}
	//Todo created
	todo := &Todo{Title: "Todo title"}
	newTodo, err := server.CreateTodo(todo)
	assert.Nil(t, err)
	assert.EqualValues(t, newTodo.ID, 1)
	assert.EqualValues(t, newTodo.Title, todo.Title)
}

func TestCreateTodo_Empty_Todo(t *testing.T) {
	//We will not be hitting the database since the execution will stop when the todo is empty
	todo := &Todo{Title: ""}
	newTodo, err := server.CreateTodo(todo)
	assert.NotNil(t, err)
	assert.Nil(t, newTodo)
	assert.EqualValues(t, err.Error(), "please provide a valid title")
}