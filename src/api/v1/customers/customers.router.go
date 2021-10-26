package customers

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	api := app.Group("/api/v1/customers")

	api.Get("/", FetchAll)
	api.Get("/:id", FetchById)
	api.Post("/", Insert)
	api.Put("/:id", UpdateById)
	api.Delete("/:id", DeleteById)
}
