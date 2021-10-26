package cinemas

import "github.com/gofiber/fiber/v2"

func FetchAll(ctx *fiber.Ctx) error {
	return ctx.JSON("Fetch all cinemas")
}

func FetchById(ctx *fiber.Ctx) error {
	return ctx.JSON("Fetch cinema")
}

func Insert(ctx *fiber.Ctx) error {
	return ctx.JSON("Insert new cinema")
}

func UpdateById(ctx *fiber.Ctx) error {
	return ctx.JSON("Update cinema")
}

func DeleteById(ctx *fiber.Ctx) error {
	return ctx.JSON("Delete cinema")
}
