package routes

import (
	"go-movies/src/pkg/entities"
	"go-movies/src/pkg/movie"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MovieRouter(app fiber.Router, service movie.Service) {
	app.Get("/", fetchMovie(service))
	app.Post("/", createMovie(service))
	app.Put("/:id", updateMovie(service))
	app.Post("/upsert/:id", upsertMovie(service))
	app.Delete("/id", deleteMovie(service))
}

func createMovie(service movie.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request entities.Movie
		err := ctx.BodyParser(&request)
		if err != nil {
			_ = ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.CreateMovie(&request)
		return ctx.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func fetchMovie(service movie.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request bson.D
		err := ctx.BodyParser(&request)
		if err != nil {
			_ = ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		movies, err := service.FetchMovie(request)
		if err != nil {
			_ = ctx.JSON(&fiber.Map{
				"status": false,
				"error":  err.Error(),
			})
		}
		return ctx.JSON(&fiber.Map{
			"status": true,
			"books":  &movies,
		})
	}
}

func updateMovie(service movie.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request entities.Movie
		err := ctx.BodyParser(&request)
		if err != nil {
			_ = ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		query := bson.D{{Key: "_id", Value: request.ID}}
		result, dberr := service.UpdateMovie(&request, query)
		if dberr != nil {
			_ = ctx.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return ctx.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func upsertMovie(service movie.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var request entities.Movie
		err := ctx.BodyParser(&request)
		if err != nil {
			_ = ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		query := bson.D{{Key: "_id", Value: request.ID}}
		result, dberr := service.UpsertMovie(&request, query)
		if dberr != nil {
			_ = ctx.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return ctx.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func deleteMovie(service movie.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// get id from req.params and check if is a valid object id
		id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
		if err != nil {
			return ctx.Status(400).JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		// initialize query object
		query := bson.D{{Key: "_id", Value: id}}
		dberr := service.DeleteMovie(query)
		if dberr != nil {
			_ = ctx.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}

		return ctx.JSON(&fiber.Map{
			"status":  false,
			"message": "Updated successfully",
		})
	}
}
