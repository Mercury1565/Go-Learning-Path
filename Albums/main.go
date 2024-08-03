package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define data
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addNewAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(id string) (*album, error) {
	for idx, album := range albums {
		if album.ID == id {
			return &albums[idx], nil
		}
	}
	return nil, errors.New("album not found")
}

func albumById(c *gin.Context) {
	id := c.Param("id")

	album, err := getAlbumById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("albums/:id", albumById)
	router.POST("/albums", addNewAlbum)

	router.Run("localhost:8080")
}
