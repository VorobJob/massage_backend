package main

import (
	"fmt"
	"log"
	"massage_app/initializers"
	"massage_app/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	dsn := os.Getenv("DB_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.Worker{})

	r := gin.Default()

	api := r.Group("/api/users")
	{
		api.GET("/", getWorkers)
		api.GET("/:id", getWorker)
		api.POST("/", createWorker)
		api.PUT("/:id", updateWorker)
		api.DELETE("/:id", deleteWorker)
	}
	r.Run()
}

func getWorkers(c *gin.Context) {
	var workers []models.Worker
	db.Find(&workers)
	c.JSON(http.StatusOK, gin.H{"workers": workers})
}

func getWorker(c *gin.Context) {
	id := c.Param("id")
	var worker models.Worker
	if err := db.First(&worker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"worker": worker})
}

func createWorker(c *gin.Context) {
	fmt.Println("Hello")
	var newWorker models.Worker
	if err := c.ShouldBindJSON(&newWorker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(db.Create(&newWorker))
	c.JSON(http.StatusCreated, gin.H{"worker": newWorker})
}

func updateWorker(c *gin.Context) {
	id := c.Param("id")
	var worker models.Worker
	if err := db.First(&worker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		return
	}

	var updatedWorker models.Worker
	if err := c.ShouldBindJSON(&updatedWorker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&worker).Updates(&updatedWorker)
	c.JSON(http.StatusOK, gin.H{"worker": updatedWorker})
}

func deleteWorker(c *gin.Context) {
	id := c.Param("id")
	var worker models.Worker
	if err := db.First(&worker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Worker not found"})
		return
	}
	db.Delete(&worker)
	c.JSON(http.StatusOK, gin.H{"message": "Worker deleted successfully"})
}
