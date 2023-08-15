package main

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/cprosche/htmx-test/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, err := store.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}

	users := map[string]string{}

	app.Get("/", basicauth.New(basicauth.Config{
		Users: users,
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Redirect("/login")
		},
	}), func(c *fiber.Ctx) error {
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

	app.Post("login", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		user := store.User{
			Username: username,
			Password: password,
		}

		err = user.Login(db)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := username + ":" + password
		token = base64.StdEncoding.EncodeToString([]byte(token))

		c.Response().Header.Add("HX-Redirect", "/")

		return c.SendString("success")
	})

	app.Post("register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		users[username] = password

		// user := store.User{
		// 	Username: username,
		// 	Password: password,
		// }

		// err = user.Create(db)
		// if err != nil {
		// 	log.Println(err)
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }

		return c.Render(
			"htmx/success",
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
