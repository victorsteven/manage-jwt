package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthDetails struct {
	AuthUuid string
	UserId   uint64
}

func CreateToken(authD AuthDetails) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_uuid"] = authD.AuthUuid
	claims["user_id"] = authD.UserId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}


//func GenerateTokenPair() (map[string]string, error) {
//	// Create token
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// Set claims
//	// This is the information which frontend can use
//	// The backend can also decode the token and get admin etc.
//	claims := token.Claims.(jwt.MapClaims)
//	claims["sub"] = 1
//	claims["name"] = "Jon Doe"
//	claims["admin"] = true
//	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
//
//	// Generate encoded token and send it as response.
//	// The signing string should be secret (a generated UUID works too)
//	t, err := token.SignedString([]byte("secret"))
//	if err != nil {
//		return nil, err
//	}
//
//	refreshToken := jwt.New(jwt.SigningMethodHS256)
//	rtClaims := refreshToken.Claims.(jwt.MapClaims)
//	rtClaims["sub"] = 1
//	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	rt, err := refreshToken.SignedString([]byte("secret"))
//	if err != nil {
//		return nil, err
//	}
//
//	return map[string]string{
//		"access_token":  t,
//		"refresh_token": rt,
//	}, nil
//}


//func CreateToken(authD AuthDetails) (map[string]string, error) {
//	//generate 15 min token
//	claimsShort := jwt.MapClaims{}
//	claimsShort["authorized"] = true
//	claimsShort["auth_uuid"] = authD.AuthUuid
//	claimsShort["user_id"] = authD.UserId
//	claimsShort["exp"] = time.Now().Add(time.Minute * 15).Unix()
//	tokenShort := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsShort)
//	short, err := tokenShort.SignedString([]byte(os.Getenv("API_SECRET")))
//	if err != nil {
//		fmt.Println("Error generating short token: ", err)
//		return nil, err
//	}
//
//	//generate refresh token 24hours long
//	claimsLong := jwt.MapClaims{}
//	claimsLong["exp"] = time.Now().Add(time.Hour * 24).Unix()
//	tokenLong := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsLong)
//	long, err :=  tokenLong.SignedString([]byte(os.Getenv("API_SECRET")))
//	if err != nil {
//		fmt.Println("Error generating long token: ", err)
//		return nil, err
//	}
//
//	//refreshToken := jwt.New(jwt.SigningMethodHS256)
//	//rtClaims := refreshToken.Claims.(jwt.MapClaims)
//	//rtClaims["sub"] = 1
//	//rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//	//rt, err := refreshToken.SignedString([]byte("secret"))
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	return map[string]string{
//		"access_token": short,
//		"refresh_token": long,
//	}, nil
//}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (*AuthDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		authUuid, ok := claims["auth_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AuthDetails{
			AuthUuid: authUuid,
			UserId:   userId,
		}, nil
	}
	return nil, err
}
