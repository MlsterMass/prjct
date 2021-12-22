package handler

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app"
	"net/http"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

// @Summary SignUp
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
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")

	if password1 == password2 {
		if err := c.ShouldBind(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid input body")
			return
		}

		_, err := h.services.Authorization.CreateUser(input)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusFound, "/index")
	}
}

type signInInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	input.Username = c.PostForm("username")
	input.Password = generatePasswordHash(c.PostForm("password"))

	if err := c.ShouldBind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		c.HTML(http.StatusOK, "sign-in.html", gin.H{
			"err": "User doesn't exist",
		})
	}

	c.Redirect(http.StatusFound, "/index")

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
