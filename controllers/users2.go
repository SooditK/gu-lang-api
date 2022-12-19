package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sooditk/gu-lang-tatti/models"
)

func CreateUser2(c *gin.Context) {
	var input struct {
		Name       string
		CreditCard []struct {
			Number string
		}
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Error Binding JSON"})
		return
	}

	user := models.User2{Name: input.Name}
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Creating User"})
		return
	}

	for _, creditCard := range input.CreditCard {
		card := models.CreditCard2{Number: creditCard.Number, UserID: user.ID}

		if err := models.DB.Create(&card).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Creating Credit Card"})
			return
		}

		user.CreditCards2 = append(user.CreditCards2, card)
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetAllUsersWithCreditCards2(c *gin.Context) {
	var users []models.User2
	if err := models.DB.Preload("CreditCards2").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Fetching Users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
