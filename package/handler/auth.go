package handler

import (
	"net/http"

	todo "github.com/Sent-Malin/zhaskGO"
	"github.com/gin-gonic/gin"
)

//signUp осуществляет регистрацию пользователей
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//signInInput нужна для парсинга токена
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"PASSWORD" binding:"required"`
}

//signIn осуществляет вход пользователя выдавая токен
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
