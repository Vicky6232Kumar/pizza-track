package main

import (
	"log/slog"
	"net/http"
	"pizza-tracking/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
}

type OrderRequest struct {
	Name    string
	Phone   string
	Address string
	Items   []OrderItemRequest
}

type OrderItemRequest struct {
	Size        string
	Pizza       string
	Quantity    int
	Instruction string
}

// This handler if for the customer to place a new order
func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(form.Items))
	for i := range orderItems {
		orderItems[i].Size = form.Items[i].Size
		orderItems[i].Pizza = form.Items[i].Pizza
		orderItems[i].Quantity = form.Items[i].Quantity
		orderItems[i].Instruction = form.Items[i].Instruction
	}

	order :=
		models.Order{
			CustomerName: form.Name,
			Phone:        form.Phone,
			Address:      form.Address,
			Status:       models.OrderStatuses[0],
			Item:         orderItems,
		}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Failed to create order")
		return
	}

	slog.Info("Order created successfully", "orderId", order.ID)

	h.notificationManager.Notify("admin:new_orders:", "new_order")
	// send a generic order data
	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "orderId": order.ID})

}

func (h *Handler) HandleOrderGet(c *gin.Context) {
	orderId := c.Param("id")

	if orderId == "" {
		c.String(http.StatusBadRequest, "OrderId not found")
		return
	}

	order, err := h.orders.GetOrder(orderId)
	if err != nil {
		c.String(http.StatusNotFound, "Order not found", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}
