package main

import (
	"os"
	"project-porto/domain"
	"project-porto/internal/api"
	"project-porto/internal/config"
	"project-porto/internal/connection"
	"project-porto/internal/repository"
	"project-porto/internal/routes"
	"project-porto/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Tambahkan Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	routes.SetupRoutes(app, projectHandler, certificateHandler, feedbackHandler)

	port := os.Getenv("PORT") 
	if port == "" {
		port = "9000"
	}
	app.Listen(":" + port)
}