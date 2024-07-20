package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserID       int           `json:"user_id"`
	Name         string        `json:"name"`
	Achievements []Achievement `json:"achievements"`
	Feedbacks    []Feedback    `json:"feedbacks"`
}

func (c *Claims) AddAchievements(achievements []Achievement) {
	c.Achievements = achievements
}
