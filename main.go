package main

import (
	"project-porto/domain"
	"project-porto/internal/api"
	"project-porto/internal/config"
	"project-porto/internal/connection"
	"project-porto/internal/repository"
	"project-porto/internal/routes"
	"project-porto/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	connection.ConnectDatabase()

	db := connection.DB
	db.AutoMigrate(&domain.Project{}, &domain.Certificate{}, &domain.Feedback{})

	// Project
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := api.NewProjectHandler(projectService)

	// Certificate
	certificateRepo := repository.NewCertificateRepository(db)
	certificateService := service.NewCertificateService(certificateRepo)
	certificateHandler := api.NewCertificateHandler(certificateService)

	// Feedback
	feedbackRepo := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepo)
	feedbackHandler := api.NewFeedbackHandler(feedbackService)

	app := fiber.New()
	routes.SetupRoutes(app, projectHandler, certificateHandler, feedbackHandler)

	app.Listen(":9000")
}