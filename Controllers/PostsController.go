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

func GetPostById(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result != nil {
		c.JSON(201, gin.H{
			"posts": post,
		})
	}
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var request struct {
		Title string
		Body  string
	}

	c.Bind(&request)

	var post models.Post
	result := initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: request.Title,
		Body:  request.Body,
	})

	if result != nil {
		c.JSON(201, gin.H{
			"posts": post,
		})
	}
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		log.Fatal("Error in deleting post " + id)
	}

	c.Status(100)
}
