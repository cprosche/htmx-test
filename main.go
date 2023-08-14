package main

import (
	"log"
	"time"

	"github.com/cprosche/htmx-test/store"
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

	db, err := store.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", authenticate, func(c *fiber.Ctx) error {
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

		return c.Redirect("/")
	})

	app.Post("register", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		user := store.User{
			Username: username,
			Password: password,
		}

		err = user.Create(db)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3000"))
}

func authenticate(c *fiber.Ctx) error {
	authenticated := false

	if !authenticated {
		return c.Redirect("/login")
	}

	return c.Next()
}
