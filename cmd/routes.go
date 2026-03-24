package main

import "github.com/gin-gonic/gin"

func setUpRoutes(router *gin.Engine, h *Handler) {
	// router.GET("/", )
	router.POST("/new-order", h.HandleNewOrderPost)
	router.GET("/order/:id", h.HandleOrderGet)

}
