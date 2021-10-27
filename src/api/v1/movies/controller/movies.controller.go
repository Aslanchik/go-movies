package movies_controller

import (
	handler "go-movies/src/api/v1/movies/handler"
	schema "go-movies/src/api/v1/movies/schema"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	// get id from req.params and check if is a valid object id
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// initialize query object
	query := bson.D{{Key: "_id", Value: id}}

	// assign memory address for movie struct
	movie := new(schema.Movie)

	error := handler.FetchById(ctx, query, movie)
	if error != nil {
		if error == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(error)
		}
		return ctx.Status(400).JSON(error)
	}
	return ctx.Status(200).JSON(movie)
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
	// get id from req.params and check if is a valid object id
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// initialize query object
	query := bson.D{{Key: "_id", Value: id}}

	movie := new(schema.Movie)

	if err := ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	err = handler.Upsert(ctx, query, movie)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}
	return ctx.SendStatus(200)
}

func UpdateById(ctx *fiber.Ctx) error {
	// get id from req.params and check if is a valid object id
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// initialize query object
	query := bson.D{{Key: "_id", Value: id}}

	movie := new(schema.Movie)

	if err = ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	err = handler.UpdateById(ctx, query, movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(err)
		}
		return ctx.Status(400).JSON(err)
	}
	return ctx.SendStatus(200)
}

func DeleteById(ctx *fiber.Ctx) error {
	// get id from req.params and check if is a valid object id
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// initialize query object
	query := bson.D{{Key: "_id", Value: id}}

	result, err := handler.DeleteById(ctx, query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	if result.DeletedCount < 1 {
		return ctx.Status(404).Send([]byte("Movie with given id does not exist."))
	}
	return ctx.SendStatus(204)
}
