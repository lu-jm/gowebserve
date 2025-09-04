package main

import (
	"log"
	"test/database"
	"test/handlers"
	"test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitGormWithConfig(); err != nil {
		log.Fatal("db conntect fail", err)
	}
	db := database.GetDB()
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/user", handlers.GreateUser)
	r.GET("/user/:id", handlers.GetUserById)

	log.Println("server on 8080")
	r.Run(":8080")
}
