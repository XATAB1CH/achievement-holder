package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Name         string        `json:"name"`
	Achievements []Achievement `json:"achievements"`
}
