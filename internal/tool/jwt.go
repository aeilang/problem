package tool

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Email string `json:"name"`
	jwt.RegisteredClaims
}

func GenJwt(email string) (string, error) {
	claims := &jwtCustomClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signString := os.Getenv("SECRET")
	
	t, err := token.SignedString([]byte(signString))
	log.Println(err)
	return t, err
}

func VerifyToken(tokenString string) (string, error) {
	tokenString = strings.TrimSpace(tokenString)
	signString := os.Getenv("SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signString), nil
	})

	if err != nil {
		return "", err
	}
	claims := token.Claims.(*jwtCustomClaims)
	return claims.Email, nil
}
