package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/renatompf/services"
)

func RegisterPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/posts")

	postGroup.GET("", services.GetAllPosts)
	postGroup.GET("/:id", services.GetPostById)
	postGroup.POST("", services.CreateNewPost)
	postGroup.DELETE(":id", services.DeletePostByID)
	postGroup.PUT("/:id", services.UpdatePostById)
}
