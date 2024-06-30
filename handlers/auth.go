package handlers

import (
	"github.com/gin-gonic/gin"
)

// var (
// 	jwtKey       = []byte("golden_horse")
// 	user         models.User
// 	existingUser models.User
// 	errHash      error
// )

func Signup(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm-password")

	// if password != confirmPassword {
	// 	c.JSON(400, gin.H{"error": "passwords are not equal"})
	// 	return
	// }

	// user := models.User{
	// 	Name:     name,
	// 	Email:    email,
	// 	Password: password,
	// }

	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(400, gin.H{"error": "shouldBindJSON"})
	// 	return
	// }

	// models.DB.Where("email = ?", user.Email).First(&existingUser)

	// if existingUser.ID != 0 {
	// 	c.JSON(400, gin.H{"error": "user already exists"})
	// 	return
	// }

	// user.Password, errHash = utils.GenerateHashPassword(user.Password)

	// if errHash != nil {
	// 	c.JSON(500, gin.H{"error": "could not generate password hash"})
	// 	return
	// }

	// добавляем пользоваетеля в базу данных

	c.JSON(200, gin.H{"name": name, "password": password, "confirmPassword": confirmPassword, "email": email})
}

func Login(c *gin.Context) { // проверяем существует ли пользователь в базе данных

	name := c.PostForm("name")
	password := c.PostForm("password")

	// user := models.User{
	// 	Name:     name,
	// 	Password: password,
	// }

	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// проверяем существование пользователя

	// errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	// if !errHash {
	// 	c.JSON(400, gin.H{"error": "invalid password"})
	// 	return
	// }

	// expirationTime := time.Now().Add(5 * time.Minute)

	// claims := &models.Claims{
	// 	Role: existingUser.Role,
	// 	StandardClaims: jwt.StandardClaims{
	// 		Subject:   existingUser.Email,
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tokenString, err := token.SignedString(jwtKey)

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "could not generate token"})
	// 	return
	// }

	// c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	// c.JSON(200, gin.H{
	// 	"status":   "ok",
	// 	"username": user.Name,
	// 	"password": user.Password,
	// })

	c.JSON(200, gin.H{"name": name, "password": password})
}

// func Home(c *gin.Context) {

// 	cookie, err := c.Cookie("token")

// 	if err != nil {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	claims, err := utils.ParseToken(cookie)

// 	if err != nil {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	if claims.Role != "user" && claims.Role != "admin" {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
// }

// func Premium(c *gin.Context) {

// 	cookie, err := c.Cookie("token")

// 	if err != nil {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	claims, err := utils.ParseToken(cookie)

// 	if err != nil {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	if claims.Role != "admin" {
// 		c.JSON(401, gin.H{"error": "unUserized"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
// }

// func Logout(c *gin.Context) {
// 	c.SetCookie("token", "", -1, "/", "localhost", false, true)
// 	c.JSON(200, gin.H{"success": "user logged out"})
// }
