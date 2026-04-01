package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBModel struct {
	DB    *gorm.DB
	Order OrderModel
	Users UserModel
}

func InitDB(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Order{}, &OrderItem{}, &User{}, &Restaurants{}, &RestaurantsMenu{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

	return &DBModel{
		DB:    db,
		Order: OrderModel{DB: db},
		Users: UserModel{DB: db},
	}, nil
}
