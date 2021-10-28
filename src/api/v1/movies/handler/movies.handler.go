package movies_handler

import (
	model "go-movies/src/api/v1/movies/model"
	schema "go-movies/src/api/v1/movies/schema"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchAll(ctx *fiber.Ctx, query bson.D) (*[]schema.Movie, error) {
	return model.FetchAll(ctx, query)
}

func FetchById(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	return model.FetchById(ctx, query, movie)
}

func Insert(ctx *fiber.Ctx, movie *schema.Movie) (interface{}, error) {
	id, err := model.Insert(ctx, movie)
	model.InsertNeo(ctx, movie, id)     // need to add error handling here
	model.InsertElastic(ctx, movie, id) // need to add error handling here
	return id, err
}

func Upsert(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	return model.Upsert(ctx, query, movie)
}

func UpdateById(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	return model.UpdateById(ctx, query, movie)
}

func DeleteById(ctx *fiber.Ctx, query bson.D) (*mongo.DeleteResult, error) {
	return model.DeleteById(ctx, query)
}
