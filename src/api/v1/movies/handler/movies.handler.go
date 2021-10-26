package movies_handler

import (
	model "go-movies/src/api/v1/movies/model"

	"github.com/gofiber/fiber/v2"
)

func FetchAll(ctx *fiber.Ctx) error {
	return model.FetchAll(ctx)
}

func FetchById(ctx *fiber.Ctx) error {
	return model.FetchById(ctx)
}

func Insert(ctx *fiber.Ctx) error {
	return model.Insert(ctx)
}

func Upsert(ctx *fiber.Ctx) error {
	return model.Upsert(ctx)
}

func UpdateById(ctx *fiber.Ctx) error {
	return model.UpdateById(ctx)
}

func DeleteById(ctx *fiber.Ctx) error {
	return model.DeleteById(ctx)
}
