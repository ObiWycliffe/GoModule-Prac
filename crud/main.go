package main

import (
	"log"

	"github.com/ObiWycliffe/GoModule/crud/handler"
	"github.com/ObiWycliffe/GoModule/crud/task"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// define database
	dsn := "obi:root@tcp(127.0.0.1:3306)/Go_CRUD?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// to migrate db schema based on defined entities
	db.AutoMigrate(&task.Task{})

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// task routes
	api.GET("/task", taskHandler.Index)
	api.POST("/task", taskHandler.Store)
	api.GET("/task/:id", taskHandler.Show)
	api.PUT("/task/:id", taskHandler.Update)
	api.DELETE("/task/:id", taskHandler.Destroy)

	router.Run(":9090")
}
