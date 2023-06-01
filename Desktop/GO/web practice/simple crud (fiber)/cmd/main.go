package main

import (
	"log"

	"github.com/basel2053/simple-crud-fiber-/pkg/books"
	"github.com/basel2053/simple-crud-fiber-/pkg/common/config"
	"github.com/basel2053/simple-crud-fiber-/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c,err := config.LoadConfig()

	if err!=nil{
		log.Fatalln("Failed at config",err)
	}

    app := fiber.New()
		db:=db.Init(c.DBURL)
		

		books.RegisterRoutes(app,db)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(c.Port)
}