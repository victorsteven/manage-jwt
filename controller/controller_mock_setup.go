package controller

import (
	"github.com/jinzhu/gorm"
	"manage-jwt/auth"
	"manage-jwt/model"
)

var (
	createUserModel func(*model.User) (*model.User, error)
	createAuthModel func(uint64) (*model.Auth, error)
	signIn func(auth.AuthDetails) (string, error)
	getUserByEmail func(string) (*model.User, error)
	fetchAuth func(*auth.AuthDetails) (*model.Auth, error)
	createTodoModel func(*model.Todo) (*model.Todo, error)
	deleteAuth func(*auth.AuthDetails) error
)

type fakeServer struct {}
type fakeSignin struct {}

//Since this methods are under the modelInterface, we must define all method there, but observe that the ones we dont need just have "panic("implement me")", while the ones we need to mock have contents.
func (fs *fakeServer) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	panic("implement me")
}
func (fs *fakeServer) ValidateEmail(string) error {
	panic("implement me")
}

func (fs *fakeServer) GetUserByEmail(email string) (*model.User, error) {
	return getUserByEmail(email)
}
func (fs *fakeServer) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	return createTodoModel(todo)
}
func (fs *fakeServer) FetchAuth(au *auth.AuthDetails) (*model.Auth, error) {
	return fetchAuth(au)
}
func (fs *fakeServer) DeleteAuth(au *auth.AuthDetails) error {
	return deleteAuth(au)
}

func (fs *fakeServer) CreateUser(user *model.User) (*model.User, error) {
	return createUserModel(user)
}
func (fs *fakeServer) CreateAuth(userId uint64) (*model.Auth, error) {
	return createAuthModel(userId)
}

func (fs *fakeSignin) SignIn(authD auth.AuthDetails) (string, error) {
	return signIn(authD)
}
