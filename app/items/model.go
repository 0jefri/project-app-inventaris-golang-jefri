package items

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	IdCategory   string    `json:"category_id"`
	Name         string    `json:"name"`
	PhotoURL     string    `json:"photo_url"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
}

func (c Items) TableName() string {
	return "items"
}
