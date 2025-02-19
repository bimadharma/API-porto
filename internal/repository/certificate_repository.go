package repository

import (
	"project-porto/domain"

	"gorm.io/gorm"
)

type CertificateRepository struct {
	DB *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) *CertificateRepository {
	return &CertificateRepository{DB: db}
}

func (r *CertificateRepository) Create(certificate *domain.Certificate) error {
	return r.DB.Create(certificate).Error
}

func (r *CertificateRepository) GetAll() ([]domain.Certificate, error) {
	var certificates []domain.Certificate
	err := r.DB.Find(&certificates).Error
	return certificates, err
}