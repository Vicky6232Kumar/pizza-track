package main

import "pizza-tracking/internal/models"

type Handler struct {
	orders              *models.OrderModel
	users               *models.UserModel
	notificationManager *NotificationManager
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		orders:              &dbModel.Order,
		users:               &dbModel.Users,
		notificationManager: NewNotificationManager(),
	}
}
