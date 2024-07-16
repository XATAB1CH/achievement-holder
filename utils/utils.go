package utils

import (
	"strings"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtKey = []byte("secret")
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWTToken(claims models.Claims) (string, error) {
	signKey := jwtKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signKey)
}

func IsRightFormat(dsn string) bool {
	str := strings.Split(dsn, ".")
	if len(str) != 2 {
		return false
	}

	format := str[1]
	if format == "png" || format == "jpg" || format == "jpeg" || format == "gif" {
		return true
	}

	return false
}
