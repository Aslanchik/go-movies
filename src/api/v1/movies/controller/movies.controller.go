package movies_controller

import (
	handler "go-movies/src/api/v1/movies/handler"

	"github.com/gofiber/fiber/v2"
)

func FetchAll(ctx *fiber.Ctx) error {
	return handler.FetchAll(ctx)
}

func FetchById(ctx *fiber.Ctx) error {
	return handler.FetchById(ctx)
}

func Insert(ctx *fiber.Ctx) error {
	return handler.Insert(ctx)
}

func UpdateById(ctx *fiber.Ctx) error {
	return handler.UpdateById(ctx)
}

func DeleteById(ctx *fiber.Ctx) error {
	return handler.DeleteById(ctx)
}
