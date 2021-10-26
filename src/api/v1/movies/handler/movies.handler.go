package movies_handler

import (
	model "go-movies/src/api/v1/movies/model"

	"github.com/gofiber/fiber/v2"
)

func FetchAll(ctx *fiber.Ctx) error {
	return model.FetchAll(ctx)
}

func FetchById(ctx *fiber.Ctx) error {
	return ctx.JSON("Fetch movie")
}

func Insert(ctx *fiber.Ctx) error {
	return ctx.JSON("Insert new movie")
}

func UpdateById(ctx *fiber.Ctx) error {
	return ctx.JSON("Update movie")
}

func DeleteById(ctx *fiber.Ctx) error {
	return ctx.JSON("Delete movie")
}
