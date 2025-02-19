package domain

import "time"

type Certificate struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	LinkImg     string    `json:"link_img"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}