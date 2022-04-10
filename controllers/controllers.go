package controllers

import (
	"time"

	"github.com/amirnilofari/go-chitchat/models"
	"github.com/gin-gonic/gin"
)

// Post controllers
func AllPost(c *gin.Context) {
	p1 := models.Post{
		Title:       "sport",
		Description: "I like sport",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	c.JSON(200, p1)
}

func CreatePost(c *gin.Context) {}

func DeletePost(c *gin.Context) {}
