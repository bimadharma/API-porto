package repository

import (
	"project-porto/domain"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) Create(project *domain.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) GetAll() ([]domain.Project, error) {
	var projects []domain.Project
	err := r.DB.Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) GetByID(id uint) (*domain.Project, error) {
	var project domain.Project
	err := r.DB.First(&project, id).Error
	return &project, err
}

