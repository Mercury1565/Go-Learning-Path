package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"duedate"`
	Status      string `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now().Format("2006-01-02"), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2).Format("2006-01-02"), Status: "Completed"},
}

func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(id string) (*Task, error) {
	for idx, task := range tasks {
		if task.ID == id {
			return &tasks[idx], nil
		}
	}
	return nil, errors.New("task not found")
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	var updated_task Task
	id := c.Param("id")

	task, err := getTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	e := c.BindJSON(&updated_task)

	if e != nil {
		return
	}

	// Now, update the task
	if updated_task.Title != "" {
		task.Title = updated_task.Title
	}
	if updated_task.Description != "" {
		task.Description = updated_task.Description
	}

	c.IndentedJSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {
		if id == task.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, tasks)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func createTask(c *gin.Context) {
	var newTask Task

	err := c.BindJSON(&newTask)

	if err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusOK, tasks)

}

func main() {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks", createTask)

	router.Run("localhost:3000")
}
