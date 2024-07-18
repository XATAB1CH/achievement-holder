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
	str := strings.Split(id, ":")
	userID, _ := strconv.Atoi(str[1])
	user, err := postgresql.GetUserByID(conn, userID)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	achievements, err := postgresql.GetAchievementsByUserID(conn, userID)
	if err != nil {
		c.HTML(http.StatusNotFound, "search_error.html", nil)
	}

	claims := models.Claims{
		Name:         user.Name,
		Achievements: achievements,
	}

	c.HTML(http.StatusOK, "index.html", claims)
}
