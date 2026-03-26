package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.Engine, h *Handler, store sessions.Store) {
	router.Use(sessions.Sessions("pizza-tracker", store))
	// router.GET("/", )

	router.POST("/login", h.HandleUserLogin)

	userProtectedRoute := router.Group("")
	userProtectedRoute.Use(h.AuthMiddleware())
	{
		router.POST("/new-order", h.HandleNewOrderPost)
		router.GET("/order/:id", h.HandleOrderGet)
		router.POST("/logout", h.HandleUserLogout)
	}

}
