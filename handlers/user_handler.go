package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test/database"
	"test/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error user info",
		})
		return
	}
	db := database.GetDB()
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("user not found, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "create user error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "get userId error",
		})
		return
	}

	db := database.GetDB()
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
