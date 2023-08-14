package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render(
			"index",
			fiber.Map{},
			"layouts/main",
		)
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render(
			"htmx/div",
			fiber.Map{
				"Text": time.Now().Format("2006-01-02 15:04:05"),
			},
		)
	})

	app.Get("login", func(c *fiber.Ctx) error {
		return c.Render(
			"login",
			fiber.Map{},
			"layouts/main",
		)
	})

	app.Get("register", func(c *fiber.Ctx) error {
		return c.Render(
			"register",
			fiber.Map{},
			"layouts/main",
		)
	})

	log.Fatal(app.Listen(":3000"))
}
