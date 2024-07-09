package handlers

import (
	"context"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/XATAB1CH/achievement-holder/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Signup(c *gin.Context) {
	var (
		errHash error
	)

	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm-password")

	if password != confirmPassword {
		c.JSON(400, gin.H{"error": "passwords are not equal"})
		return
	}

	password, errHash = utils.GenerateHashPassword(password)
	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	id := postgresql.InsertUser(conn, name, email, password)

	if id == 0 {
		c.JSON(500, gin.H{"error": "user is already created"})
		return
	}

	c.JSON(200, gin.H{"id": id, "name": name, "password": password, "email": email})
}

func Login(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	userDB, err := postgresql.GetUserByName(conn, name)
	if err != nil {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}

	expectedPassword := userDB.Password

	if !utils.CompareHashPassword(password, expectedPassword) {
		c.JSON(401, gin.H{"error": "wrong password"})
		return
	}

	// Create the JWT claims
	claims := &models.Claims{
		Name: userDB.Name,
	}

	tokenString, err := utils.GenerateJWTToken(*claims)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	// Create the cookie
	c.SetCookie("token", tokenString, 3600, "/", "127.0.0.1", false, true)

	c.JSON(200, gin.H{"name": name, "password": password})
}
