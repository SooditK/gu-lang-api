package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sooditk/gu-lang-tatti/controllers"
	"github.com/sooditk/gu-lang-tatti/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsersWithCreditCards)
	r.POST("/users2", controllers.CreateUser2)
	r.GET("/users2", controllers.GetAllUsersWithCreditCards2)
	r.POST("/users3", controllers.CreateUser3)
	r.GET("/users3", controllers.GetAllUsers)
	r.GET("/languages", controllers.GetAllLanguages)

	r.Run()
}
