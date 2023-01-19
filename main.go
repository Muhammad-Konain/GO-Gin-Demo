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

	r.GET("/", controllers.CreatePosts)

	r.Run()
}
