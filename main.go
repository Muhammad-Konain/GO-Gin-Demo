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
	r.GET("/GetAllPosts", controllers.GetPosts)

	r.Run()
}
