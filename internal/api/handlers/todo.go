package handlers

import (
	"net/http"
	"test_openapi_go/internal/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var todos []models.Todo = []models.Todo{}

// GetTodos godoc
//
// @Summary      Get all todos
// @ID todosGet
// @Description  Retrieves a list of all todo items.
// @Tags         todos
// @Produce      json
// @Success      200  {array}   models.Todo "List of todo items"
// @Failure      500  {object}  models.ErrorResponse "Internal Server Error"
// @Security     BearerAuth
// @Router       /v1/todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// CreateTodo godoc
//
// @Summary      Create a new todo
// @ID todosCreate
// @Description  Creates a new todo item with the provided details.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.Todo  true  "Todo item details"
// @Success      201   {object}  models.Todo "Newly created todo item"
// @Failure      400   {object}  models.ErrorResponse "Bad Request - Invalid input or missing fields"
// @Failure      500   {object}  models.ErrorResponse "Internal Server Error"
// @Security     BearerAuth
// @Router       /v1/todos [post]
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = uuid.New().String() // Generate a unique ID
	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

// GetTodoByID godoc
//
// @Summary      Get a todo by ID
// @ID todosGetByID
// @Description  Retrieves a single todo item by its ID.
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  models.Todo "Todo item found"
// @Failure      404  {object}  models.ErrorResponse "Not Found - Todo item not found"
// @Failure      500  {object}  models.ErrorResponse "Internal Server Error"
// @Security     BearerAuth
// @Router       /v1/todos/{id} [get]
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Todo not found"})
}

// UpdateTodo godoc
//
// @Summary      Update an existing todo
// @ID todosUpdate
// @Description  Updates an existing todo item with the provided details.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "Todo ID"
// @Param        todo  body      models.Todo  true  "Todo item details to update"
// @Success      200   {object}  models.Todo "Updated todo item"
// @Failure      400   {object}  models.ErrorResponse "Bad Request - Invalid input or missing fields"
// @Failure      404   {object}  models.ErrorResponse "Not Found - Todo item not found"
// @Failure      500   {object}  models.ErrorResponse "Internal Server Error"
// @Security     BearerAuth
// @Router       /v1/todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			// Update fields
			todos[i].Title = updatedTodo.Title
			todos[i].Done = updatedTodo.Done
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Todo not found"})
}

// DeleteTodo godoc
//
// @Summary      Delete a todo
// @ID todosDelete
// @Description  Deletes a todo item by its ID.
// @Tags         todos
// @Produce      json
// @Param        id   path      string  true  "Todo ID"
// @Success      204  {string}  string "Todo item deleted"
// @Failure      404  {object}  models.ErrorResponse "Not Found - Todo item not found"
// @Failure      500  {object}  models.ErrorResponse "Internal Server Error"
// @Security     BearerAuth
// @Router       /v1/todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Todo not found"})
}
