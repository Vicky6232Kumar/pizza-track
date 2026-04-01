package models

import (
	"time"

	"gorm.io/gorm"
)

type RestaurantModel struct {
	DB *gorm.DB
}

type Restaurants struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string `gorm:"size:200"`
	PhoneNo   string
	Email     string            `gorm:"size:60"`
	Location  string            `gorm:"size:200"`
	IsOpen    bool              `gorm:"default:true"`
	Menu      []RestaurantsMenu `gorm:"foreignKey:RestaurantId" json:"menu"`
	CreatedAt time.Time         `json:"createdAt"`
}

// [Action] : other important details are also important to take for scalability

// [Warn] : the menu list should ideally present in the nosql
type RestaurantsMenu struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	RestaurantId string `gorm:"index;not null" json:"restaurantId"`
	Name         string `gorm:"not null"`
	Description  string `gorm:"size:500"`
	Price        int    `gorm:"not null"`
	IsAvailable  bool   `gorm:"default:false"`
}
