package app

import (
	"fmt"
	"go-movies/src/api/v1"
	mongo "go-movies/src/db/mongodb"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	err := mongo.ConfigureAndConnect()

	if err == nil {
		fmt.Println("DB Connection Established!")
	}

	api.SetupRoutes(app)

	app.Listen(":3000")
}
