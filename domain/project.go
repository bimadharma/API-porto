package domain

import "time"

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Link        string
	LinkImg     string    `json:"link_img"` 
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

// Custom nama tabel di database
func (Project) TableName() string {
	return "projects"
}
