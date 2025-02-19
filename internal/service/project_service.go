package service

import (
	"project-porto/domain"
	"project-porto/dto"
	"project-porto/internal/repository"
)

	type ProjectService struct {
		Repo *repository.ProjectRepository
	}

	func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
		return &ProjectService{Repo: repo}
	}

	func (s *ProjectService) CreateProject(req dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
		project := domain.Project{
			Title:       req.Title,
			Description: req.Description,
			Link:        req.Link,
			LinkImg:     req.LinkImg,
		}

		err := s.Repo.Create(&project)
		if err != nil {
			return nil, err
		}

		savedProject, err := s.Repo.GetByID(project.ID)
		if err != nil {
		return nil, err
	}

	return &dto.ProjectResponse{
		ID:          savedProject.ID,
		Title:       savedProject.Title,
		Description: savedProject.Description,
		Link:        savedProject.Link,
		LinkImg:     savedProject.LinkImg,   
		CreatedAt:   savedProject.CreatedAt, 
	}, nil

	}

	func (s *ProjectService) GetAllProjects() ([]dto.ProjectResponse, error) {
		projects, err := s.Repo.GetAll()
		if err != nil {
			return nil, err
		}

		var responses []dto.ProjectResponse
		for _, project := range projects {
			responses = append(responses, dto.ProjectResponse{
				ID:          project.ID,
				Title:       project.Title,
				Description: project.Description,
				Link:        project.Link,
				LinkImg:     project.LinkImg,
				CreatedAt:   project.CreatedAt, 
			})
		}

		return responses, nil
	}
