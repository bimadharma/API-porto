package dto

import "time"

type CreateCertificateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkImg     string `json:"link_img"`
}

type CertificateResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	LinkImg     string    `json:"link_img"`
	CreatedAt   time.Time `json:"created_at"`
}