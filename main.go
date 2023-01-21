package main

import (
	controllers "github.com/Muhammad-Konain/GO-Gin-Demo/Controllers"
	initializers "github.com/Muhammad-Konain/GO-Gin-Demo/Initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/CreatePost", controllers.CreatePosts)
	r.PUT("/UpdatePost/:id", controllers.UpdatePost)
	r.GET("/GetAllPosts", controllers.GetPosts)
	r.GET("/GetAllPosts/:id", controllers.GetPostById)
	r.DELETE("/DeletePost/:id", controllers.DeletePost)

	r.Run()
}
