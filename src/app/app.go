package app

import (
	"go-movies/src/api/v2/routes"
	"go-movies/src/db"
	"go-movies/src/db/mongodb"
	"go-movies/src/pkg/movie"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApp() {
	db.InitDatabases()

	movieCollection := mongodb.Instance.Db.Collection("movies")
	movieRepo := movie.NewRepo(movieCollection)
	movieService := movie.NewService(movieRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome!"))
	})
	api := app.Group("/api/v1/movies")
	routes.MovieRouter(api, movieService)

	app.Listen(":3000")
}
