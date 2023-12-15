package service

import (
	"github.com/gin-gonic/gin"
	"github.com/renatompf/dtos"
	"github.com/renatompf/initializers"
	"github.com/renatompf/models"
	"log"
	"net/http"
)

func GetAllPosts(c *gin.Context) {
	var allPosts []models.Post

	initializers.DB.Find(&allPosts)

	c.IndentedJSON(http.StatusOK, allPosts)
}

func GetPostById(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.IndentedJSON(http.StatusOK, post)
}

func CreateNewPost(c *gin.Context) {
	var newPostDTO dtos.NewPostDTO

	err := c.BindJSON(&newPostDTO)
	if err != nil {
		log.Fatal("Error while mapping body into the DTO")
		return
	}

	post := models.Post{Title: newPostDTO.Title, Body: newPostDTO.Body}

	initializers.DB.Create(&post)

	c.IndentedJSON(http.StatusCreated, post)
}

func DeletePostByID(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.IndentedJSON(http.StatusNoContent, nil)
}

func UpdatePostById(c *gin.Context) {
	var newPostDTO dtos.NewPostDTO

	err := c.BindJSON(&newPostDTO)
	if err != nil {
		log.Fatal("Error while mapping body into the DTO")
		return
	}

	var post models.Post

	id := c.Param("id")
	initializers.DB.First(&post, id)

	post.Title = newPostDTO.Title
	post.Body = newPostDTO.Body

	initializers.DB.Updates(&post)

	c.IndentedJSON(http.StatusOK, post)
}
