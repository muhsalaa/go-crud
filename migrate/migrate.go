package main

import (
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// this models instance of post, with the nil value of its corresponding struct
	initializers.DB.AutoMigrate(&models.Post{})
}
