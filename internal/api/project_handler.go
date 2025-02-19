package api

import (
	"project-porto/internal/service"

	"project-porto/dto"

	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	Service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{Service: service}
}

func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	var req dto.CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	project, err := h.Service.CreateProject(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create project"})
	}

	return c.Status(201).JSON(project)
}

func (h *ProjectHandler) GetAllProjects(c *fiber.Ctx) error {
	projects, err := h.Service.GetAllProjects()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch projects"})
	}

	return c.JSON(projects)
}