package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/pug"
	"github.com/wilian1808/sqlmql/pkg/routes"
)

func main() {
	// template
	engine := pug.New("./views", ".pug")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/", routes.IndexPage)
	app.Post("/query", routes.GetQuery)
	app.Get("/result", routes.ResultMQL)
	app.Listen(":5555")

	// s := "select * from users"
	// s := "insert into users (username, paternal, maternal, email, age) VALUES (marco, perez, perez, marco@gmail.com, 25)"
	// s := "UPDATE users SET username = juan, paternal = perez, maternal = perez, email = perez@gmail.com, age = 25 WHERE username = pablo"
	// s := "DELETE * FROM users WHERE username = juan"
}
