package main

import (
	"log"

	"github.com/PatrikOlin/haberdashery/pkg/db"
	"github.com/PatrikOlin/haberdashery/pkg/garment"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	_, err := db.Open()
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
}

func main() {

	app := fiber.New()
	app.Use(cors.New())
	setupRoutes(app)

	app.Listen(":4040")
}

func setupRoutes(app *fiber.App) {

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Tjänare världen!")
	})

	v1.Get("/garments", garment.GetAllGarments)
	v1.Post("/garments/:id", garment.CreateGarment)
	v1.Put("/garments/:id", garment.UpdateGarment)
	v1.Patch("/garments/:id/increment", garment.IncrementGarment)
}
