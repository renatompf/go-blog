package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renatompf/controllers"
	"github.com/renatompf/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToPostgres()
}

func main() {
	r := gin.Default()
	group := r.Group("/posts")
	group.GET("", controllers.GetAllPosts)
	group.GET("/:id", controllers.GetPostById)
	group.POST("", controllers.CreateNewPost)
	group.DELETE(":id", controllers.DeletePostByID)
	group.PUT("/:id", controllers.UpdatePostById)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
