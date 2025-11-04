package handlers

import (
	"net/http"
	"test_openapi_go/internal/api/auth"
	"test_openapi_go/internal/api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string][]byte{}

// Register godoc
//
// @Summary      Register a new user
// @ID authRegister
// @Description  Registers a new user with a unique username and password. The password will be hashed before storage.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User registration details"
// @Success      201   {object}  map[string]string "User created successfully"
// @Failure      400   {object}  models.ErrorResponse "Bad Request - Invalid input or missing fields"
// @Failure      500   {object}  models.ErrorResponse "Internal Server Error - Failed to hash password"
// @Router       /v1/auth/register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	users[user.Username] = hashedPassword

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// Login godoc
//
// @Summary      Login a user
// @ID authLogin
// @Description  Authenticates a user with username and password, returning a JWT token upon successful login.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User login credentials"
// @Success      200   {object}  map[string]string "Successful authentication, returns JWT token"
// @Failure      400   {object}  models.ErrorResponse "Bad Request - Invalid input or missing fields"
// @Failure      401   {object}  models.ErrorResponse "Unauthorized - Invalid credentials"
// @Failure      500   {object}  models.ErrorResponse "Internal Server Error - Failed to generate token"
// @Router       /v1/auth/login [post]
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, ok := users[user.Username]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
