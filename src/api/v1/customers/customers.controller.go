package customers

import "github.com/gofiber/fiber/v2"

func FetchAll(ctx *fiber.Ctx) error {
	return ctx.JSON("Fetch all customers")
}

func FetchById(ctx *fiber.Ctx) error {
	return ctx.JSON("Fetch customer")
}

func Insert(ctx *fiber.Ctx) error {
	return ctx.JSON("Insert new customer")
}

func UpdateById(ctx *fiber.Ctx) error {
	return ctx.JSON("Update customer")
}

func DeleteById(ctx *fiber.Ctx) error {
	return ctx.JSON("Delete customer")
}
