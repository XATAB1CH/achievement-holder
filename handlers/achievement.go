package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/XATAB1CH/achievement-holder/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Create(c *gin.Context) {

	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	file, err := c.FormFile("file")
	if err != nil {
		text := "Нужно прикрепить файл!"
		c.HTML(http.StatusNotFound, "creation_error.html", text)
		return
	}

	dsn := "assets/images/" + file.Filename
	// Проверяем, что файл нужного формата
	if !utils.IsRightFormat(dsn) {
		text := "Недопустимый формат файла!"
		c.HTML(http.StatusNotFound, "creation_error.html", text)
		return
	}
	c.SaveUploadedFile(file, dsn)

	title := c.PostForm("title")
	info := c.PostForm("info")

	claims, _ := c.Get("claims")
	userID := claims.(*models.Claims).UserID

	id := postgresql.InsertAchievement(conn, title, dsn, info, userID)

	if id == 0 {
		text := "Это достижение уже опубликовано!"
		c.HTML(http.StatusNotFound, "creation_error.html", text)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func Information(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id")) // Получаем id достижения
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	achievement, err := postgresql.GetAchievementByID(conn, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	val, _ := c.Get("auth")
	auth := val.(string)

	switch auth {
	case "true":
		c.HTML(http.StatusOK, "information_auth.html", achievement)
	case "false":
		c.HTML(http.StatusOK, "information.html", achievement)
	}

}

func Search(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	name := c.PostForm("name")

	user, err := postgresql.GetUserByName(conn, name)
	userID := strconv.Itoa(user.ID)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	c.Redirect(http.StatusFound, "/:"+userID)
}

func Delete(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id")) // Получаем id достижения
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	_ = postgresql.DeleteAchievement(conn, id)

	c.Redirect(http.StatusFound, "/")
}
