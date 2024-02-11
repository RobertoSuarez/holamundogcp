package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.SendString("Hola, Mundo! / hola")
	})

	log.Println(app.Listen(":3000"))
}
