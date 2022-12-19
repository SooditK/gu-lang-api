package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sooditk/gu-lang-tatti/models"
)

type CreditCard struct {
	Number string
}

type User struct {
	Name       string
	CreditCard CreditCard
}

func CreateUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	creditCard := models.CreditCard{Number: user.CreditCard.Number}
	models.DB.Create(&creditCard)

	models.DB.Create(&models.User{Name: user.Name, CreditCard: creditCard})

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetAllUsersWithCreditCards(c *gin.Context) {
	var users []models.User
	err := models.DB.Preload("CreditCard").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
