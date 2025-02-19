package service

import (
	"project-porto/domain"
	"project-porto/dto"
	"project-porto/internal/repository"
)

type CertificateService struct {
	Repo *repository.CertificateRepository
}

func NewCertificateService(repo *repository.CertificateRepository) *CertificateService {
	return &CertificateService{Repo: repo}
}

func (s *CertificateService) CreateCertificate(req dto.CreateCertificateRequest) (*dto.CertificateResponse, error) {
	certificate := domain.Certificate{
		Title:       req.Title,
		Description: req.Description,
		LinkImg:     req.LinkImg,
	}

	err := s.Repo.Create(&certificate)
	if err != nil {
		return nil, err
	}

	return &dto.CertificateResponse{
		ID:          certificate.ID,
		Title:       certificate.Title,
		Description: certificate.Description,
		LinkImg:     certificate.LinkImg,
		CreatedAt:   certificate.CreatedAt,
	}, nil
}

func (s *CertificateService) GetAllCertificates() ([]dto.CertificateResponse, error) {
	certificates, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.CertificateResponse
	for _, certificate := range certificates {
		responses = append(responses, dto.CertificateResponse{
			ID:          certificate.ID,
			Title:       certificate.Title,
			Description: certificate.Description,
			LinkImg:     certificate.LinkImg,
			CreatedAt:   certificate.CreatedAt,
		})
	}

	return responses, nil
}