package common

import (
	"gin-vue-bookStore/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_creat")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseTocken(user model.User)(string,error){
	expirationTime := time.Now().Add(7*24* time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "Csc0001",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

func ParseToken(tokenString string)(*jwt.Token,*Claims,error){
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})
	return token,claims,err
}