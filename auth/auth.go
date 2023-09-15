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


func init()  {
	godotenv.Load("globals.env")
	jwt_token=os.Getenv("SECRET_KEY")
	Big_secret=[]byte(jwt_token)
}


type Payload struct{
	Username string
	Password string
	jwt.StandardClaims
}


func Generatejwt(username string,password string) string {
	
	expiresat:=time.Now().Add(48 * time.Hour)
	claims:=&Payload{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresat.Unix(),
		},
	}
	
	tokenjwt:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//specifies the secret key using which the jwt token should be signed with
	tokenstring,_:=tokenjwt.SignedString(Big_secret)
	return tokenstring
}

//for validating the jwt token coming with each request
func ValidateToken(jwt_token string) bool {

	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return Big_secret, nil
    })
	//checks if the token is expired or not
	if err == nil && token.Valid {
        return true
    } else {
        return false
    }
}