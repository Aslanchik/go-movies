package movie

import (
	"context"
	"fmt"
	"go-movies/src/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	CreateMovie(movie *entities.Movie) (*entities.Movie, error)
	FetchMovie(query bson.D) (*[]entities.Movie, error)
	UpdateMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error)
	UpsertMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error)
	DeleteMovie(query bson.D) error
}

type repository struct {
	Collection *mongo.Collection
}

// creates a single instance of the repo
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateMovie(movie *entities.Movie) (*entities.Movie, error) {
	movie.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *repository) FetchMovie(query bson.D) (*[]entities.Movie, error) {
	var movies []entities.Movie

	cursor, err := r.Collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var movie entities.Movie
		_ = cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	fmt.Println(movies)
	return &movies, nil
}

func (r *repository) UpdateMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error) {
	update := bson.M{"set": movie}

	_, err := r.Collection.UpdateOne(context.Background(), query, update)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *repository) UpsertMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error) {
	opts := options.Update().SetUpsert(true)
	update := bson.M{"set": movie}

	_, err := r.Collection.UpdateOne(context.Background(), query, update, opts)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *repository) DeleteMovie(query bson.D) error {
	_, err := r.Collection.DeleteOne(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
