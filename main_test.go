package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"test_openapi_go/internal/api/models"
	"test_openapi_go/internal/api/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	routes.SetupRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestAuthRoutes(t *testing.T) {
	router := gin.Default()
	routes.SetupRoutes(router)

	// Register
	w := httptest.NewRecorder()
	userJSON := `{"username":"testuser","password":"testpassword"}`
	req, _ := http.NewRequest("POST", "/v1/auth/register", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Login
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/auth/login", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestTodoRoutes(t *testing.T) {
	router := gin.Default()
	routes.SetupRoutes(router)

	// 1. Register and Login to get a token
	userJSON := `{"username":"todouser","password":"todopassword"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/auth/register", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/auth/login", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse struct {
		Token string `json:"token"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)
	token := loginResponse.Token
	assert.NotEmpty(t, token)

	authHeader := "Bearer " + token

	// 2. Create Todo (POST /v1/todos)
	newTodo := models.Todo{Title: "Buy groceries", Done: false}
	newTodoJSON, _ := json.Marshal(newTodo)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/v1/todos", strings.NewReader(string(newTodoJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdTodo models.Todo
	err = json.Unmarshal(w.Body.Bytes(), &createdTodo)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdTodo.ID)
	assert.Equal(t, newTodo.Title, createdTodo.Title)
	assert.Equal(t, newTodo.Done, createdTodo.Done)

	todoID := createdTodo.ID

	// 3. Get All Todos (GET /v1/todos)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/todos", nil)
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var todos []models.Todo
	err = json.Unmarshal(w.Body.Bytes(), &todos)
	assert.NoError(t, err)
	assert.Len(t, todos, 1)
	assert.Equal(t, createdTodo.ID, todos[0].ID)

	// 4. Get Single Todo (GET /v1/todos/{id})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/todos/"+todoID, nil)
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var fetchedTodo models.Todo
	err = json.Unmarshal(w.Body.Bytes(), &fetchedTodo)
	assert.NoError(t, err)
	assert.Equal(t, createdTodo.ID, fetchedTodo.ID)
	assert.Equal(t, createdTodo.Title, fetchedTodo.Title)

	// 5. Update Todo (PUT /v1/todos/{id})
	updatedTodo := models.Todo{ID: todoID, Title: "Buy groceries and milk", Done: true}
	updatedTodoJSON, _ := json.Marshal(updatedTodo)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/v1/todos/"+todoID, strings.NewReader(string(updatedTodoJSON)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseUpdatedTodo models.Todo
	err = json.Unmarshal(w.Body.Bytes(), &responseUpdatedTodo)
	assert.NoError(t, err)
	assert.Equal(t, updatedTodo.Title, responseUpdatedTodo.Title)
	assert.Equal(t, updatedTodo.Done, responseUpdatedTodo.Done)

	// 6. Delete Todo (DELETE /v1/todos/{id})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/v1/todos/"+todoID, nil)
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	// 7. Verify Deletion (GET /v1/todos)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/v1/todos", nil)
	req.Header.Set("Authorization", authHeader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var remainingTodos []models.Todo
	err = json.Unmarshal(w.Body.Bytes(), &remainingTodos)
	assert.NoError(t, err)
	assert.Len(t, remainingTodos, 0)
}