package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Index
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /index [get]
func (h *Handler) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index page",
	})
}
