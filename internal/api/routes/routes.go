package routes

import (
	"test_openapi_go/internal/api/handlers"
	"test_openapi_go/internal/api/middleware"

	_ "test_openapi_go/internal/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes(r *gin.Engine) {
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", handlers.Ping)

	api := r.Group("/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		todos := api.Group("/todos")
		todos.Use(middleware.AuthMiddleware())
		{
			todos.GET("", handlers.GetTodos)
			todos.POST("", handlers.CreateTodo)
			todos.GET("/:id", handlers.GetTodoByID)
			todos.PUT("/:id", handlers.UpdateTodo)
			todos.DELETE("/:id", handlers.DeleteTodo)
		}
	}
}
