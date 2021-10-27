package movies_controller

import (
	handler "go-movies/src/api/v1/movies/handler"
	schema "go-movies/src/api/v1/movies/schema"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAll(ctx *fiber.Ctx) error {
	query := bson.D{{}}

	data, err := handler.FetchAll(ctx, query)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(200).JSON(data)
}

func FetchById(ctx *fiber.Ctx) error {
	return handler.FetchById(ctx)
}

func Insert(ctx *fiber.Ctx) error {
	movie := new(schema.Movie)
	movie.ID = ""
	if err := ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	data, err := handler.Insert(ctx, movie)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	return ctx.Status(201).JSON(data)
}

func Upsert(ctx *fiber.Ctx) error {
	return handler.Upsert(ctx)
}

func UpdateById(ctx *fiber.Ctx) error {
	return handler.UpdateById(ctx)
}

func DeleteById(ctx *fiber.Ctx) error {
	return handler.DeleteById(ctx)
}
