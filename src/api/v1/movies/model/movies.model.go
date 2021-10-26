package movies_model

import (
	schema "go-movies/src/api/v1/movies/schema"
	mongo "go-movies/src/db/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func FetchAll(ctx *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := mongo.Instance.Db.Collection(schema.SCHEMA_NAME).Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	var movies []schema.Movie = make([]schema.Movie, 0)

	if err := cursor.All(ctx.Context(), &movies); err != nil {
		return ctx.Status(500).JSON(err)
	}

	return ctx.Status(200).JSON(movies)
}
