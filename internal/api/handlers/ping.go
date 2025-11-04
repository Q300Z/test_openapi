package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
//
// @Summary      Ping the server
// @ID pingGet
// @Description  Checks the health and responsiveness of the API server.
// @Tags         ping
// @Produce      json
// @Success      200   {string}  string "pong"
// @Failure      500   {object}  models.ErrorResponse "Internal Server Error"
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
