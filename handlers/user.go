package handlers

import (
	"net/http"

	"github.com/Iagobarros211256/voluryashop/models"
	"github.com/Iagobarros211256/voluryashop/repository"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	repository.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
