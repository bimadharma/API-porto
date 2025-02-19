package domain

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Feedback  string         `json:"feedback"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
