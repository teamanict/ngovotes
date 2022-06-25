package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teamanict/ngovotes/controllers"
)

func RegisterAPIRoutes(app fiber.Router) {
	// Register Candidate
	app.Post("/candidate/register", controllers.RegisterCandidate)
	// Login
	// Vote
}