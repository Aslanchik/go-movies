package movies_model

import (
	schema "go-movies/src/api/v1/movies/schema"
	mongodb "go-movies/src/db/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchAll(ctx *fiber.Ctx, query bson.D) (*[]schema.Movie, error) {
	cursor, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).Find(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	var movies []schema.Movie = make([]schema.Movie, 0)

	if err := cursor.All(ctx.Context(), &movies); err != nil {
		return nil, err
	}

	return &movies, err
}

func FetchById(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	collection := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME)

	error := collection.FindOne(ctx.Context(), query).Decode(&movie)

	if error != nil {
		return error
	}

	return nil
}

func Insert(ctx *fiber.Ctx, movie *schema.Movie) (interface{}, error) {
	collection := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME)

	res, err := collection.InsertOne(ctx.Context(), movie)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func UpdateById(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: movie.Title},
			{Key: "year", Value: movie.Year},
			{Key: "director", Value: movie.Director},
			{Key: "genre", Value: movie.Genre},
		},
	}}

	err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).FindOneAndUpdate(ctx.Context(), query, update).Err()

	if err != nil {
		return err
	}

	return nil
}

func Upsert(ctx *fiber.Ctx, query bson.D, movie *schema.Movie) error {
	// set the query to upsert if no movie exists
	opts := options.Update().SetUpsert(true)
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: movie.Title},
			{Key: "year", Value: movie.Year},
			{Key: "director", Value: movie.Director},
			{Key: "genre", Value: movie.Genre},
		},
	}}

	_, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).UpdateOne(ctx.Context(), query, update, opts)

	if err != nil {
		return err
	}

	return nil

}

func DeleteById(ctx *fiber.Ctx, query bson.D) (*mongo.DeleteResult, error) {
	result, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).DeleteOne(ctx.Context(), &query)
	if err != nil {
		return nil, err
	}

	// the movie was deleted
	return result, nil
}
