package api

import (
	"project-porto/dto"
	"project-porto/internal/service"

	"github.com/gofiber/fiber/v2"
)

type FeedbackHandler struct {
	Service *service.FeedbackService
}

func NewFeedbackHandler(service *service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{Service: service}
}

func (h *FeedbackHandler) CreateFeedback(c *fiber.Ctx) error {
	var req dto.CreateFeedbackRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	feedback, err := h.Service.CreateFeedback(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create feedback"})
	}

	return c.Status(201).JSON(feedback)
}

func (h *FeedbackHandler) GetAllFeedbacks(c *fiber.Ctx) error {
	feedbacks, err := h.Service.GetAllFeedbacks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch feedbacks"})
	}

	return c.JSON(feedbacks)
}
