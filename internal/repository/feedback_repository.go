package repository

import (
	"project-porto/domain"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	DB *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{DB: db}
}

func (r *FeedbackRepository) Create(feedback *domain.Feedback) error {
	return r.DB.Create(feedback).Error
}

func (r *FeedbackRepository) GetAll() ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	err := r.DB.Find(&feedbacks).Error
	return feedbacks, err
}
