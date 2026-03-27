package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
}

func (h *Handler) HandlerOrderPut(c *gin.Context) {
	orderId := c.Param("id")
	newStatus := c.PostForm("status")

	if err := h.orders.UpdateOrderStatus(orderId, newStatus); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	h.notificationManager.Notify("order:"+orderId, "order_updated")

	c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated the status",
	})
}

func (h *Handler) HandleOrderDelete(c *gin.Context) {
	orderId := c.Param("id")

	if err := h.orders.DeleteOrder(orderId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully deleted the order",
	})
}
