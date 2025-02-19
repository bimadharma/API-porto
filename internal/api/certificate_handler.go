package api

import (
	"project-porto/dto"
	"project-porto/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CertificateHandler struct {
	Service *service.CertificateService
}

func NewCertificateHandler(service *service.CertificateService) *CertificateHandler {
	return &CertificateHandler{Service: service}
}

func (h *CertificateHandler) CreateCertificate(c *fiber.Ctx) error {
	var req dto.CreateCertificateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	certificate, err := h.Service.CreateCertificate(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create certificate"})
	}

	return c.Status(201).JSON(certificate)
}

func (h *CertificateHandler) GetAllCertificates(c *fiber.Ctx) error {
	certificates, err := h.Service.GetAllCertificates()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch certificates"})
	}

	return c.JSON(certificates)
}
