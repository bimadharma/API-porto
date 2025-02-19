package routes

import (
	"project-porto/internal/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, projectHandler *api.ProjectHandler, certificateHandler *api.CertificateHandler, feedbackHandler *api.FeedbackHandler) {
	api := app.Group("/api")

	// Endpoint Project
	api.Post("/projects", projectHandler.CreateProject)
	api.Get("/projects", projectHandler.GetAllProjects)

	// Endpoint Certificate
	api.Post("/certificates", certificateHandler.CreateCertificate)
	api.Get("/certificates", certificateHandler.GetAllCertificates)

	// Endpoint Feedback
	api.Post("/feedbacks", feedbackHandler.CreateFeedback)
	api.Get("/feedbacks", feedbackHandler.GetAllFeedbacks)
}