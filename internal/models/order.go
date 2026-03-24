package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{"Order Placed", "Preparing", "Baking", "Quality Check", "Out for Delivery", "Delivered"}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID           string      `gorm:"primaryKey;size:14" json:"id"`
	Status       string      `grom:"not null" json:"status"`
	CustomerName string      `grom:"not null" json:"customerName"`
	Phone        string      `grom:"not null" json:"phone"`
	Address      string      `grom:"not null" json:"address"`
	Item         []OrderItem `gorm:"foreignKey:OrderID" json:"item"`
	CreatedAt    time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ID          string `gorm:"primaryKey;size:14" json:"id"`
	OrderID     string `gorm:"index;size:14;not null" json:"orderId"`
	Size        string `gorm:"not null" json:"size"`
	Pizza       string `gorm:"not null" json:"pizza"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	Instruction string `gorm:"not null" json:"instruction"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (o *OrderModel) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}

func (o *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Item").First(&order, "id = ?", id).Error
	return &order, err
}
