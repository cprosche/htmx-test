package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
		},
	))

	// db, err := store.ConnectToDb()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	users := map[string]string{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render(
			"layouts/main",
			fiber.Map{},
		)
	})

	app.Get("login", func(c *fiber.Ctx) error {
		emailCookie := c.Cookies("email")
		if emailCookie != "" {
			return c.Render(
				"htmx/homepage",
				fiber.Map{},
			)
		}

		return c.Render(
			"htmx/login",
			fiber.Map{},
		)
	})

	app.Get("register", func(c *fiber.Ctx) error {
		return c.Render(
			"htmx/register",
			fiber.Map{},
		)
	})

	app.Post("login", func(c *fiber.Ctx) error {
		time.Sleep(500 * time.Millisecond)
		emailCookie := c.Cookies("email")
		emailForm := c.FormValue("email")
		passwordForm := c.FormValue("password")

		if emailCookie != "" {
			return c.Render(
				"htmx/homepage",
				fiber.Map{},
			)
		}

		if emailCookie == "" && emailForm != "" && passwordForm != "" {
			c.Cookie(&fiber.Cookie{
				Name:     "email",
				Value:    emailForm,
				Expires:  time.Now().Add(time.Hour * 24),
				HTTPOnly: true,
			})
			return c.Render(
				"htmx/homepage",
				fiber.Map{},
			)
		}

		return c.SendStatus(fiber.StatusBadRequest)
	})

	app.Post("register", func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		if email == "" || password == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		users[email] = password

		return c.Render(
			"htmx/success",
			fiber.Map{},
		)
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "email",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})

		return c.Render(
			"htmx/login",
			fiber.Map{},
		)
	})

	log.Fatal(app.Listen(":3000"))
}

// func authenticate(c *fiber.Ctx) error {
// 	authenticated := false

// 	if !authenticated {
// 		return c.Redirect("/login")
// 	}

// 	return c.Next()
// }

// TODO: Add authentication middleware
