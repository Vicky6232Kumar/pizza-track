package models

import (
	"time"

	"gorm.io/gorm"
)

type ResturantModel struct {
	DB *gorm.DB
}

type Resturants struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string `gorm:"size:200"`
	PhoneNo   int
	Email     string           `gorm:"size:60"`
	Location  string           `gorm:"size:200"`
	IsOpen    bool             `gorm:"default:true"`
	Menu      []ResturantModel `gorm:"foreignKey:OrderID" json:"menu"`
	CreatedAt time.Time        `json:"createdAt"`
}

// [Action] : other important details are also important to take for scalability

// [Warn] : the menu list should ideally present in the nosql
type ResturantsMenu struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ResturantId string `gorm:"index;not null" json:"resturantId"`
	Name        string `gorm:"not null"`
	Description string `gorm:"size:500"`
	Price       int    `gorm:"not null"`
	IsAvailable bool   `gorm:"default:false"`
}
