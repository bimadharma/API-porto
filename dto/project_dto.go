package dto

import "time"

	type CreateProjectRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		LinkImg     string `json:"link_img"` 
	}

	type ProjectResponse struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		LinkImg     string    `json:"link_img"`
		CreatedAt   time.Time `json:"created_at"`
	}
