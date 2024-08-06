package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Task represents a task with its properties.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(target_id string) (*Task, error, int) {
	for i, task := range tasks {
		if task.ID == target_id {
			return &tasks[i], nil, i
		}
	}
	return nil, errors.New("task not found"), -1
}

func getTask(c *gin.Context) {
	target_id := c.Param("id")
	task, err, _ := getTaskByID(target_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	target_id := c.Param("id")
	task, err, _ := getTaskByID(target_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	task.Title = newTask.Title
	task.Description = newTask.Description

	c.IndentedJSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	target_id := c.Param("id")
	_, err, idx := getTaskByID(target_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
}

func addTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "task added"})
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks", addTask)

	router.Run("localhost:8080")
}
