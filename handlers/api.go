package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/XATAB1CH/achievement-holder/models"
	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Demo(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error", nil)
	}

	id := c.Param("id")

	// странная передача данных
	id = strings.TrimPrefix(id, ":")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	user, err := postgresql.GetUserByID(conn, userID)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	achievements, err := postgresql.GetAchievementsByUserID(conn, userID)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	feedbacks, _ := postgresql.GetFeedbacks(conn)

	claims := models.Claims{
		Name:         user.Name,
		Achievements: achievements,
		Feedbacks:    feedbacks,
	}

	c.HTML(http.StatusOK, "index.html", claims)
}

func Feedback(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback.html", nil)
}

func FeedbackForm(c *gin.Context) {
	conn, err := pgx.Connect(context.Background(), postgresql.GetDSN())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	name := c.PostForm("name")
	feedback := c.PostForm("feedback")

	id := postgresql.InsertFeedback(conn, name, feedback)

	if id == 0 {
		text := "Вы уже опубликовали отзыв!"
		c.HTML(http.StatusNotFound, "feedback_error.html", text)
		return
	}

	c.Redirect(http.StatusFound, "/")
}
