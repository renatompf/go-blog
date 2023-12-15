package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/renatompf/service"
)

func RegisterPostRoutes(r *gin.Engine) {
	postGroup := r.Group("/posts")

	postGroup.GET("", service.GetAllPosts)
	postGroup.GET("/:id", service.GetPostById)
	postGroup.POST("", service.CreateNewPost)
	postGroup.DELETE(":id", service.DeletePostByID)
	postGroup.PUT("/:id", service.UpdatePostById)
}
