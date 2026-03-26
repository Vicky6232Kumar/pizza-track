package main

import (
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (h *Handler) notificationHandler(c *gin.Context) {
	orderId := c.Query("orderId")

	if orderId == "" {
		c.String(400, "Invalid orderId")
	}

	_, err := h.orders.GetOrder(orderId)
	if err != nil {
		c.String(404, "Order Not found")
		return
	}

	key := "order:" + orderId
	client := make(chan string, 10)
	h.notificationManager.AddClient(key, client)

	defer func() {
		h.notificationManager.RemoveClient(key, client)
		slog.Info("Customer client disconnected", "orderId", orderId)
	}()

	h.streamSSE(c, client)

}

func (h *Handler) streamSSE(c *gin.Context, client chan string) {
	c.Header("Content-Type", "event/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-client; ok {
			c.SSEvent("message", msg)
			return true
		}
		return true
	})

}

func (h *Handler) adminNotificationHandler(c *gin.Context) {
	key := "admin:new_orders"
	client := make(chan string, 10)

	h.notificationManager.AddClient(key, client)

	defer func() {
		h.notificationManager.RemoveClient(key, client)
		slog.Info("Admin Client disconnected")
	}()

	h.streamSSE(c, client)
}
