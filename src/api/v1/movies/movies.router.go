package movies

import (
	controller "go-movies/src/api/v1/movies/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	api := app.Group("/api/v1/movies")

	api.Get("/", controller.FetchAll)
	api.Get("/:id", controller.FetchById)
	api.Post("/", controller.Insert)
	api.Put("/:id", controller.UpdateById)
	api.Post("/upsert/:id", controller.Upsert)
	api.Delete("/:id", controller.DeleteById)
}
