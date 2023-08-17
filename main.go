package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// db, err := store.ConnectToDb()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	users := map[string]string{}

	app.Get("/", func(c *fiber.Ctx) error {
		// set cookie
		// c.Cookie(&fiber.Cookie{
		// 	Name:     "username",
		// 	Value:    "john",
		// 	Expires:  time.Now().Add(time.Hour * 24),
		// 	HTTPOnly: true,
		// })
		c.Set("HX-Redirect", "/success")

		// get cookie
		cookie := c.Cookies("username")
		if cookie == "" {
			return c.Render(
				"htmx/login",
				fiber.Map{},
				"layouts/main",
			)
		}

		return c.Render(
			"htmx/homepage",
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
		usernameCookie := c.Cookies("username")
		usernameForm := c.FormValue("username")
		passwordForm := c.FormValue("password")

		if usernameCookie != "" {
			// set header
			return c.Render(
				"htmx/homepage",
				fiber.Map{},
			)
		}

		if usernameCookie == "" && usernameForm != "" && passwordForm != "" {
			c.Cookie(&fiber.Cookie{
				Name:     "username",
				Value:    usernameForm,
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
		username := c.FormValue("username")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		users[username] = password

		return c.Render(
			"htmx/success",
			fiber.Map{},
		)
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "username",
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
