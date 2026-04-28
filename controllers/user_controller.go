package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApplicationUser struct {
	ID    string
	Name  string
	Email string
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	usersGroup := rg.Group("/users")
	{
		usersGroup.GET("/", getUsers)
		usersGroup.GET("/:id", getUserById)
		usersGroup.POST("/", createUser)
	}
}

// Should retrun all the user that we have
func getUsers(c *gin.Context) {
	fmt.Println("Getting a specific user")
}

func createUser(c *gin.Context) {
	fmt.Println("Creating a new user")
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}