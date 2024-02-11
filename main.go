package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

func main() {

	viper.SetDefault("PORT", 3000)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! con fiber, viper y go")
	})

	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.SendString("Hola, Mundo! / hola")
	})

	log.Println(app.Listen(":" + viper.GetString("PORT")))
}
