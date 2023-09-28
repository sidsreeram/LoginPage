package token

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwt_token string
var Big_secret []byte

func init() {
	godotenv.Load("globals.env")
	jwt_token = os.Getenv("SECRET_KEY")
	Big_secret = []byte(jwt_token)
}

func Generatejwt(username string, password string) (string, error) {
	expiresat := time.Now().Add(48 * time.Hour)
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      expiresat.Unix(),
	}

	tokenjwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := tokenjwt.SignedString(Big_secret)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

func ValidateToken(jwt_token string) bool {
	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Big_secret, nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
