package api

import (
	"go-movies/src/api/v1/cinemas"
	"go-movies/src/api/v1/customers"
	"go-movies/src/api/v1/movies"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	movies.InitRoutes(app)
	customers.InitRoutes(app)
	cinemas.InitRoutes(app)
}
