package dto

import "time"

type CreateFeedbackRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Feedback string `json:"feedback"`
}

type FeedbackResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Feedback  string    `json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
}
