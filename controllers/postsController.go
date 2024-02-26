package controllers

import (
	"belajarweb/initializers"
	"belajarweb/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct{
		Body string
		Title string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	//Create a new post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.JSON(400, gin.H{
			"error": "Failed to create post",
		})
		return
	}
	//Return the new post

	c.JSON(200, gin.H{
		"message": "Post created successfully",
	})
}

func PostsIndex(c *gin.Context) {
	//Get all posts
	var posts []models.Post
	// Get all the posts from the db
	initializers.DB.Find(&posts)

	//Return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from url
	var post models.Post
	//Get the post from the db
	initializers.DB.First(&post, c.Param("id"))

	//Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get id from url
	var post models.Post
	//Get the post from the db
	initializers.DB.First(&post, c.Param("id"))

	//Get data off req body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	//Update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//Return the updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get id from url
	id:= c.Param("id")
	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)


	//Return the deleted post
	c.Status(200)
}

