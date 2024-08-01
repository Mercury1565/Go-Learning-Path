package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	// Capitalizing the field names here is crucial
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author`
	Quantity int    `json:"quantity`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)

	if err != nil {
		return
	}

	// Successsful binding
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// id query not found
	if ok == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query not found"})
		return
	}

	book, err := getBookById(id)

	// Book not found
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book ot found"})
		return
	}

	// Book not available
	if book.Quantity == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book unavailable"})
		return
	}

	// Book found, decrement quantity
	book.Quantity -= 1
	c.JSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// id query not found
	if ok == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query not found"})
		return
	}

	book, err := getBookById(id)

	// Book not found
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book ot found"})
		return
	}

	// return book, increment quantity
	book.Quantity += 1
	c.JSON(http.StatusOK, book)

}

func main() {
	// Setup router
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080")
}
