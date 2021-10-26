package app

import (
	"go-movies/src/api/v1"
	"go-movies/src/db"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	db.InitDatabases()

	api.SetupRoutes(app)

	app.Listen(":3000")
}
