package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)
//READ THIS TO UNDERSTAND THE TEST CASES BELOW:
//Remember what the program emphasizes, we didnt specify how long a token will last while creating it. Which means, a token can last forever. So, the token used  below is a valid one, except you alter it.
//The way we invalidate a token is to create a new jwt with a different uuid, thereby rendering the formerly created valid token invalid.
//We ran all test with a valid token. If have time, you can add test cases and use a random token, then assert for the errors. As an example, i altered the token in this test "TestToken_Invalid", as error was the result

func TestCreateToken(t *testing.T) {
	au := AuthDetails{
		AuthUuid: "43b78a87-6bcf-439a-ab2e-940d50c4dc33", //this can be anything
		UserId:   1,
	}
	token, err := CreateToken(au)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestVerifyToken(t *testing.T) {
	//In order to generate a request, let use the logout endpoint
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkduc"

	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)

	jwtAns, err := VerifyToken(req)

	assert.Nil(t, err)
	assert.NotNil(t, jwtAns) //this is of type *jwt.Token
}

func TestExtractToken(t *testing.T) {
	//In order to generate a request, let use the logout endpoint
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkduc"

	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)

	result := ExtractToken(req)
	assert.NotNil(t, result)
	assert.EqualValues(t, result, token)
}

//Check the auth details from the token:
func TestExtractTokenAuth(t *testing.T) {
	//In order to generate a request, let use the logout endpoint
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkduc"

	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)

	result, err := ExtractTokenAuth(req)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.UserId)
	assert.NotNil(t, result.AuthUuid)
}


func TestTokenValid(t *testing.T) {
	//In order to generate a request, let use the logout endpoint
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkduc"

	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)

	errToken := TokenValid(req)
	assert.Nil(t, errToken)
}

//i added garbage to the token, so is not valid
func TestToken_Invalid(t *testing.T) {
	//In order to generate a request, let use the logout endpoint
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX3V1aWQiOiJjMmUxYjBjMy00ZGRjLTQ0NjUtYWVkNC1iNGE2NDM5NzI4M2MiLCJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX2lkIjoxfQ.FWbfdhEJeK7mjZ-lWvs9scuyUrSKPrC4xafUoEqkducxx"

	tokenString := fmt.Sprintf("Bearer %v", token)
	req.Header.Set("Authorization", tokenString)

	errToken := TokenValid(req)
	assert.NotNil(t, errToken)
	assert.EqualValues(t, "illegal base64 data at input byte 45", errToken.Error())
}