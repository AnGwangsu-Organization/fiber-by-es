package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your/repo/database"
	"log"
)

func main() {
	app := fiber.New()

	database.Init()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Ping Pong by fiber!!!\n")
	})

	log.Fatal(app.Listen(":3000"))
}
