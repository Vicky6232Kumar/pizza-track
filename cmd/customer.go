package cmd

type OrderFormData struct {
}

type OrderRequest struct {
	Name    string             `form:"name" binding:"required; max=100, min=2"`
	Phone   string             `form:"phone" binding:"required; max=15, min=10"`
	Address string             `form:"address" binding:"required; max=255, min=5"`
	Items   []OrderItemRequest `form:"items" binding:"required"`
}

type OrderItemRequest struct {
	Size        string `form:"size" binding:"required"`
	Pizza       string `form:"pizza" binding:"required"`
	Quantity    int    `form:"quantity" binding:"required"`
	Instruction string `form:"instruction" binding:"required"`
}


// This handler if for the customer to place a new order
func (h *Handler) HandleNewOrderPost(c *gin.Context){
	var form OrderRequest

	if err := c.ShouldBind(&form); err :=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	order : models.Order{
		CustomerName: form.Name,
		phone: form.Phone,
		Address: form.Address,
		Status: models.order[0],
		Item: form.Items
	}

	if err := h.orders.CreateOrder(&order); err != nil{
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Failed to create order")
		return
	}

	slog.Info("Order created successfully", "orderId", order.ID)

	// send a generic order data
	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "orderId": order.ID})

}