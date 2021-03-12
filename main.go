package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wilian1808/sqlmql/pkg/helpers"
	"github.com/wilian1808/sqlmql/pkg/models"
	"github.com/wilian1808/sqlmql/pkg/mongo"
	"github.com/wilian1808/sqlmql/pkg/sintaxsql"
)

func main() {
	// template
	// engine := pug.New("./views", ".pug")
	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })

	// app.Static("/", "./public")
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// }))

	// app.Get("/", index)
	// app.Post("/query", query)
	// app.Listen(":5555")

	// evaluamos la cadena de entrada damos formato e procedemos a designar
	// s := "select * from users"
	// s := "insert into users (username, paternal, maternal, email, age) VALUES (marco, perez, perez, marco@gmail.com, 25)"
	// s := "UPDATE users SET username = juan, paternal = perez, maternal = perez, email = perez@gmail.com, age = 25 WHERE username = pablo"
	s := "FROM users WHERE username = juan"
	data, _ := helpers.FormatData(s)
	// podemos ver los tipos de dato
	// helpers.GetTypes(data)
	//fmt.Println(data)
	designar(data)
	// tenemos que designar dependiendo con que palabra empieza SELECT, INSERT, DELETE, UPDATE
}

// evaluador de las funciones
func designar(data []interface{}) {
	switch data[0] {
	case sintaxsql.SelectSQL:
		res := mongo.TransformSelect(data)
		fmt.Println(res)
	case sintaxsql.InsertSQL:
		res := mongo.TransformInsert(data)
		fmt.Println(res)
	case sintaxsql.UpdateSQL:
		res := mongo.TransformUpdate(data)
		fmt.Println(res)
	case sintaxsql.DeleteSQL:
		res := mongo.TransformDelete(data)
		fmt.Println(res)
	default:
		res := fmt.Errorf("Error en la sintaxis: %s", data[0])
		fmt.Println(res)
	}
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"data": "db.users.find()",
	})
}

func query(c *fiber.Ctx) error {
	query := models.Query{}
	if err := c.BodyParser(&query); err != nil {
		return err
	}

	// obtenemos los valores de la consulta
	arr, err := helpers.FormatData(query.Data)
	if err != nil {
		return err
	}

	helpers.GetTypes(arr)

	return nil
}
