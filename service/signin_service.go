package service

import "manage-jwt/auth"

//because i will mock signin in the future while testing, it will be defined in an interface:
type sigInInterface interface {
	SignIn(auth.AuthDetails) (string, error)
}
type signInStruct struct {}

//let expose this package:
var (
	Authorize sigInInterface = &signInStruct{}
)

func (ss signInStruct) SignIn(authD auth.AuthDetails) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}