package service

import (
	"project-porto/domain"
	"project-porto/dto"
	"project-porto/internal/repository"
)

type FeedbackService struct {
	Repo *repository.FeedbackRepository
}

func NewFeedbackService(repo *repository.FeedbackRepository) *FeedbackService {
	return &FeedbackService{Repo: repo}
}

func (s *FeedbackService) CreateFeedback(req dto.CreateFeedbackRequest) (*dto.FeedbackResponse, error) {
	feedback := domain.Feedback{
		Name:     req.Name,
		Email:    req.Email,
		Feedback: req.Feedback,
	}

	err := s.Repo.Create(&feedback)
	if err != nil {
		return nil, err
	}

	return &dto.FeedbackResponse{
		ID:        feedback.ID,
		Name:      feedback.Name,
		Email:     feedback.Email,
		Feedback:  feedback.Feedback,
		CreatedAt: feedback.CreatedAt,
	}, nil
}

func (s *FeedbackService) GetAllFeedbacks() ([]dto.FeedbackResponse, error) {
	feedbacks, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.FeedbackResponse
	for _, feedback := range feedbacks {
		responses = append(responses, dto.FeedbackResponse{
			ID:        feedback.ID,
			Name:      feedback.Name,
			Email:     feedback.Email,
			Feedback:  feedback.Feedback,
			CreatedAt: feedback.CreatedAt,
		})
	}

	return responses, nil
}
