package main

import (
	"fmt"

	"github.com/debojyoti-majumder/go-limekey/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("This is a simple server")

	r := gin.Default()
	v1Api := r.Group("/api")

	controllers.RegisterUserRoutes(v1Api)

	fmt.Println("Running auth service")
	r.Run(":8084")
}
