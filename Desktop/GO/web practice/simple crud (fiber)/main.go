package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

		app.Post("/todos",func(c *fiber.Ctx) error {
			result:= c.Body()
			fmt.Println(result)
			return c.SendString("Hey ")
		})
		
		app.Get("/todos/:id",func(c *fiber.Ctx) error {
			var id = c.Params("id")
			return c.Status(fiber.StatusOK).SendString(id);
		})

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen("localhost:3000")
}