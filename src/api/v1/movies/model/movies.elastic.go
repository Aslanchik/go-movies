package movies_model

import (
	"encoding/json"
	schema "go-movies/src/api/v1/movies/schema"
	"go-movies/src/db/elastic"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const INDEX = "movies"

func InsertElastic(ctx *fiber.Ctx, movie *schema.Movie, movieId interface{}) error {
	es := *elastic.Instance.Client

	id := movieId.(primitive.ObjectID).Hex()
	body, _ := json.Marshal(movie)

	res, err := es.Index(INDEX, strings.NewReader(string(body)), es.Index.WithDocumentID(id), es.Index.WithRefresh("true"))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
