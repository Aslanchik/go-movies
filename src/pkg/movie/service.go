package movie

import (
	"go-movies/src/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	CreateMovie(movie *entities.Movie) (*entities.Movie, error)
	FetchMovie(query bson.D) (*[]entities.Movie, error)
	UpdateMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error)
	UpsertMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error)
	DeleteMovie(query bson.D) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateMovie(movie *entities.Movie) (*entities.Movie, error) {
	return s.repository.CreateMovie(movie)
}

func (s *service) FetchMovie(query bson.D) (*[]entities.Movie, error) {
	return s.repository.FetchMovie(query)
}

func (s *service) UpdateMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error) {
	return s.repository.UpdateMovie(movie, query)
}

func (s *service) UpsertMovie(movie *entities.Movie, query bson.D) (*entities.Movie, error) {
	return s.repository.UpsertMovie(movie, query)
}

func (s *service) DeleteMovie(query bson.D) error {
	return s.repository.DeleteMovie(query)
}
