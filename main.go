package main

import (
	"agus/backend-api/controllers"
	"agus/backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/posts", controllers.StorePost)
	router.GET("/api/posts/:id", controllers.FindPostById)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)

	router.Run(":3000")
}