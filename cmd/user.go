package main

import (
	"net/http"
	"pizza-tracking/internal/models"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserName string `json:"userName"`
	Password string
}

type UserLoginData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserSignUpData struct {
	UserLoginData
	Name string `json:"name" binding:"required"`
}

func (h *Handler) UserRegistration(c *gin.Context) {
	var signUpData UserSignUpData

	if err := c.ShouldBind(&signUpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userData := models.User{
		Email:    signUpData.Email,
		Password: signUpData.Password,
		Name:     signUpData.Name,
	}

	mail, id, err := h.users.CreateUser(&userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// session
	SetSessionValue(c, "userId", id)
	SetSessionValue(c, "email", mail)

	// trigger the email service

	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully created",
		"name":    userData.Name,
	})
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
	// SetSessionValue(c, "userName", user.UserName)

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "orderId": user.ID})
}

func (h *Handler) HandleUserLogout(c *gin.Context) {
	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

}
