package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
	"github.com/go-crud/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostIndexId)
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.PostUpdate)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
