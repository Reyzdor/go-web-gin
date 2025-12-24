package handlers

import (
	"Application/database"
	"Application/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminPage(c *gin.Context) {
	userID := 1

	user, err := repository.GetUserByID(userID)
	if err != nil || user.Role != "admin" {
		c.HTML(http.StatusForbidden, "error.html", gin.H{
			"error": "You're not admin!",
		})
		return
	}

	c.HTML(http.StatusOK, "admin.html", nil)
}

func CreatePost(c *gin.Context) {
	userID := 1

	user, err := repository.GetUserByID(userID)
	if err != nil || user.Role != "admin" {
		c.HTML(http.StatusForbidden, "error.html", gin.H{
			"error": "Only for admins!",
		})

		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	_, err = database.DB.Exec(
		"INSERT INTO posts (title, content, user_id) VALUES(?,?,?)",
		title, content, userID,
	)

	if err != nil {
		c.String(500, "Error: %v", err)
		return
	}

	c.Redirect(302, "/admin")
}
