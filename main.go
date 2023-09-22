package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	r := gin.Default()

	dsn := "host=localhost user=postgres password=2585 dbname=gormdemo port=5432  sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	// Auto-migrate the Task model
	db.AutoMigrate(&Task{})

	// Define API routes
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	// Start the Gin server
	r.Run(":8080")
}

// Get all tasks
func getTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Get a task by ID
func getTask(c *gin.Context) {
	var task Task
	id := c.Param("id")
	db.First(&task, id)
	c.JSON(http.StatusOK, task)
}

// Create a new task
func createTask(c *gin.Context) {
	var task Task
	c.BindJSON(&task)
	db.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// Update a task by ID
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	db.First(&task, id)
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Delete a task by ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	db.First(&task, id)
	db.Delete(&task)
	c.Status(http.StatusNoContent)
}
