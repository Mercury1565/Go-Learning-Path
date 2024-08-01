package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles = []Article{
	{ID: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{ID: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	// define a new_article and attach the entered JSON data on it
	var newArticle Article
	json.NewDecoder(r.Body).Decode(&newArticle)

	// Add new_article to articles
	articles = append(articles, newArticle)
}

func getArticleByID(id int) (*Article, error) {
	for i, article := range articles {
		if article.ID == id {
			return &articles[i], nil
		}
	}
	return nil, errors.New("article not found")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	// split the entered URL path to access the ID entered
	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathParts[len(pathParts)-1])

	// Handle error if entered ID is not valid
	if err != nil {
		http.Error(w, "Please enter valid id", http.StatusBadRequest)
		return
	}

	// Get the article with an ID 'id'
	article, err := getArticleByID(id)

	// Article not found
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func editArticle(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathParts[len(pathParts)-1])

	// Handle error if entered ID is not valid
	if err != nil {
		http.Error(w, "Please enter valid id", http.StatusBadRequest)
	}

	// Get the article with ID 'id'
	article, err := getArticleByID(id)

	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
	}

	// Get the JSON data from the request
	var editedArticle Article
	json.NewDecoder(r.Body).Decode(&editedArticle)

	// Edit fields of the target article
	article.Title = editedArticle.Title
	article.Desc = editedArticle.Desc
	article.Content = editedArticle.Content
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathParts[len(pathParts)-1])

	// Handle error if entered ID is not valid
	if err != nil {
		http.Error(w, "Please enter valid id", http.StatusBadRequest)
	}

	// Find the article with ID 'id' and remove it
	for i, article := range articles {
		if article.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
			return
		}
	}

	// Article not found
	http.Error(w, "Article not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", homePage)

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			returnAllArticles(w, r)
		} else if r.Method == http.MethodPost {
			addArticle(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			returnArticle(w, r)
		} else if r.Method == http.MethodPut {
			editArticle(w, r)
		} else if r.Method == http.MethodDelete {
			deleteArticle(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
