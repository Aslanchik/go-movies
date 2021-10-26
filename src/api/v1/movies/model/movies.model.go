package movies_model

import (
	schema "go-movies/src/api/v1/movies/schema"
	mongodb "go-movies/src/db/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchAll(ctx *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var movies []schema.Movie = make([]schema.Movie, 0)

	if err := cursor.All(ctx.Context(), &movies); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(200).JSON(movies)
}

func FetchById(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	collection := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME)

	query := bson.D{{Key: "_id", Value: id}}

	movie := new(schema.Movie)

	error := collection.FindOne(ctx.Context(), query).Decode(&movie)

	if error != nil {
		if error == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(error)
		}
		return ctx.Status(500).JSON(error)
	}

	return ctx.Status(200).JSON(movie)
}

func Insert(ctx *fiber.Ctx) error {
	collection := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME)

	movie := new(schema.Movie)

	if err := ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	movie.ID = ""

	res, err := collection.InsertOne(ctx.Context(), movie)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: res.InsertedID}}
	createdDoc := collection.FindOne(ctx.Context(), filter)

	createdMovie := &schema.Movie{}
	createdDoc.Decode(createdMovie)

	return ctx.Status(201).JSON(createdMovie)
}

func UpdateById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	movieId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	movie := new(schema.Movie)

	if err := ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: movieId}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: movie.Title},
			{Key: "year", Value: movie.Year},
			{Key: "director", Value: movie.Director},
			{Key: "genre", Value: movie.Genre},
		},
	}}

	err = mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).FindOneAndUpdate(ctx.Context(), query, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return ctx.Status(404).SendString(err.Error())
		}
		return ctx.Status(500).SendString(err.Error())
	}

	movie.ID = id
	return ctx.Status(200).JSON(movie)
}

func Upsert(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	movieId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	movie := new(schema.Movie)

	if err := ctx.BodyParser(movie); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	// set the query to upsert if no movie exists
	opts := options.Update().SetUpsert(true)
	query := bson.D{{Key: "_id", Value: movieId}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "title", Value: movie.Title},
			{Key: "year", Value: movie.Year},
			{Key: "director", Value: movie.Director},
			{Key: "genre", Value: movie.Genre},
		},
	}}

	res, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).UpdateOne(ctx.Context(), query, update, opts)

	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(200).JSON(res)

}

func DeleteById(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(
		ctx.Params("id"),
	)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// find and delete the movie with the given ID
	query := bson.D{{Key: "_id", Value: id}}
	result, err := mongodb.Instance.Db.Collection(schema.SCHEMA_NAME).DeleteOne(ctx.Context(), &query)

	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// the movie might not exist
	if result.DeletedCount < 1 {
		return ctx.Status(404).Send([]byte("Movie with given id does not exist."))
	}

	// the movie was deleted
	return ctx.SendStatus(204)
}
