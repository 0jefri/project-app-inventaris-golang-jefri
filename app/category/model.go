package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Category) TableName() string {
	return "category"
}
