package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemId      uint           `json:"itemId" gorm:"primary_key"`
	ItemCode    string         `json:"itemCode"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	OrderId     uint           `json:"orderId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
