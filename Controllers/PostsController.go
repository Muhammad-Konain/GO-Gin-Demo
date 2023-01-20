package controllers

import (
	"log"

	initializers "github.com/Muhammad-Konain/GO-Gin-Demo/Initializers"
	models "github.com/Muhammad-Konain/GO-Gin-Demo/Models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {

	var request struct {
		Title string
		Body  string
	}

	c.Bind(&request)

	post := models.Post{
		Title: request.Title,
		Body:  request.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		log.Fatal(result.Error)
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result != nil {
		c.JSON(201, gin.H{
			"posts": posts,
		})
	}
}
