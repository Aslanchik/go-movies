package movies_handler

import (
	model "go-movies/src/api/v1/movies/model"
	schema "go-movies/src/api/v1/movies/schema"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAll(ctx *fiber.Ctx, query bson.D) (*[]schema.Movie, error) {
	data, err := model.FetchAll(ctx, query)
	return data, err
}

func FetchById(ctx *fiber.Ctx) error {
	return model.FetchById(ctx)
}

func Insert(ctx *fiber.Ctx, movie *schema.Movie) (interface{}, error) {
	id, err := model.Insert(ctx, movie)
	model.InsertNeo(ctx, movie, id) // need to add error handling here
	return id, err
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
