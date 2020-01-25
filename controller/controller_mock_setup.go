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

)

type fakeServer struct {}

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
func (fs *fakeServer) GetUserByID(uint64) (*model.User, error) {
	panic("implement me")
}
func (fs *fakeServer) CreateTodo(*model.Todo) (*model.Todo, error) {
	panic("implement me")
}
func (fs *fakeServer) FetchAuth(*auth.AuthDetails) (*model.Auth, error) {
	panic("implement me")
}
func (fs *fakeServer) DeleteAuth(*auth.AuthDetails) error {
	panic("implement me")
}

func (fs *fakeServer) CreateUser(user *model.User) (*model.User, error) {
	return createUserModel(user)
}
func (fs *fakeServer) CreateAuth(userId uint64) (*model.Auth, error) {
	return createAuthModel(userId)
}


type fakeSignin struct {}

func (fs *fakeSignin) SignIn(authD auth.AuthDetails) (string, error) {
	return signIn(authD)
}
