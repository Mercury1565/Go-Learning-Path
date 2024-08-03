package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var blogs = []Blog{
	{ID: "1", Name: "Hermon", Title: "Hermon's journal", Body: "Hello there, this is hermon"},
	{ID: "2", Name: "Fenu", Title: "Fenu's journal", Body: "Hello there, this is fenu"},
	{ID: "3", Name: "Mami", Title: "Mami's journal", Body: "Hello there, this is mami"},
	{ID: "4", Name: "Zebe", Title: "Zebe's journal", Body: "Hello there, this is zebe"},
}

func getAllBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs)
}

func getBlogById(id string) (*Blog, error) {
	for i, blog := range blogs {
		if blog.ID == id {
			return &blogs[i], nil
		}
	}
	return nil, errors.New("Blog not found")
}

func getBlog(c *gin.Context) {
	id := c.Param("id")

	blog, err := getBlogById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, blog)
}

func addBlog(c *gin.Context) {
	var newBlog Blog

	err := c.BindJSON(&newBlog)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	blogs = append(blogs, newBlog)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Blog added successfully ! "})
}

func editBlog(c *gin.Context) {
	// Bind inputed JSON
	var editedBlog Blog
	err := c.BindJSON(&editedBlog)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	// Edit blog of a specific ID
	id := c.Param("id")
	blog, err := getBlogById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
		return
	}

	blog.Title = editedBlog.Title
	blog.Body = editedBlog.Body

	c.IndentedJSON(http.StatusOK, blog)
}

func deleteBlog(c *gin.Context) {
	id := c.Param("id")

	for i, blog := range blogs {
		if blog.ID == id {
			blogs = append(blogs[:i], blogs[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted successsfully!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
}

func main() {
	router := gin.Default()

	router.GET("/blogs", getAllBlogs)
	router.GET("/blogs/:id", getBlog)
	router.POST("/blogs", addBlog)
	router.PUT("/blogs/:id", editBlog)
	router.DELETE("/blogs/:id", deleteBlog)

	router.Run("localhost:4000")
}
