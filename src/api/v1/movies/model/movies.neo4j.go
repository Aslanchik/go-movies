package movies_model

import (
	"fmt"
	schema "go-movies/src/api/v1/movies/schema"
	"go-movies/src/db/neo4j"

	"github.com/gofiber/fiber/v2"
)

func InsertNeo(ctx *fiber.Ctx, movie *schema.Movie, movieId interface{}) error {
	query := fmt.Sprintf(`CREATE (n:Movie {_id: '%s',title: '%s', year: '%s', director: '%s', genre: '%s'})`, movieId, movie.Title, movie.Year, movie.Director, movie.Genre)

	_, err := neo4j.Instance.Session.Run(query, nil)
	if err != nil {
		return err
	}
	return nil
}
