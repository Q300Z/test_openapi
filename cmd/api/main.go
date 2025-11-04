package main

import (
	"test_openapi/internal/api/routes"

	"test_openapi/pkg/config"

	"github.com/gin-gonic/gin"
)

// @title TodoList API
// @version 1.0
// @description This is a sample server for a todo list.
// @host localhost:8081
// @schemes http

func main() {

	config.LoadConfig()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8081")

}
