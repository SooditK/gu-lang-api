package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sooditk/gu-lang-tatti/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User3
	if err := models.DB.Preload("Languages").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Fetching Users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetAllLanguages(c *gin.Context) {
	var languages []models.Language
	if err := models.DB.Find(&languages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Fetching Languages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"languages": languages})
}

func CreateUser3(c *gin.Context) {
	var input struct {
		Name          string
		LanguageNames []struct {
			Name string
		}
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Error Parsing Input"})
		return
	}

	user := models.User3{Name: input.Name}
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Creating User"})
		return
	}

	var languages []models.Language
	for _, languageName := range input.LanguageNames {
		var language models.Language

		if err := models.DB.Where("name = ?", languageName.Name).FirstOrCreate(&language, models.Language{Name: languageName.Name}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Creating Language"})
			return
		}
		languages = append(languages, language)
	}

	if err := models.DB.Model(&user).Association("Languages").Replace(languages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Error Associating Languages with User"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":      user,
		"languages": languages,
	})
}
