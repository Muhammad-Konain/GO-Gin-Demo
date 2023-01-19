package main

import (
	"log"

	initializers "github.com/Muhammad-Konain/GO-Gin-Demo/Initializers"
	models "github.com/Muhammad-Konain/GO-Gin-Demo/Models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatal(err.Error())
	}
}
