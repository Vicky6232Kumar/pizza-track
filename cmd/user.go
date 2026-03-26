package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserName string `json:"userName"`
	Password string
}

func (h *Handler) UserRegistration(c *gin.Context) {

}

func (h *Handler) HandleUserLogin(c *gin.Context) {
	var formData UserData

	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.users.AuthenticateUser(formData.UserName, formData.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SetSessionValue(c, "userId", user.ID)
	SetSessionValue(c, "userName", user.UserName)

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "orderId": user.ID})
}

func (h *Handler) HandleUserLogout(c *gin.Context) {
	if err := clearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

}
