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

	// evaluamos la cadena de entrada damos formato e procedemos a evaluar
	data, _ := helpers.FormatData("select username,paternal, maternal,email,age from users")
	// podemos ver los tipos de dato
	// helpers.GetTypes(data)
	//fmt.Println(data)
	evaluar(data)
	// tenemos que evaluar dependiendo con que palabra empieza SELECT, INSERT, DELETE, UPDATE
}

// evaluador de las funciones
func evaluar(data []interface{}) {
	switch data[0] {
	case sintaxsql.SelectSQL:
		res := mongo.TransformSelect(data)
		fmt.Println(res)
	case sintaxsql.InsertSQL:
		fmt.Println("la consulta es un insert")
	case sintaxsql.UpdateSQL:
		fmt.Println("la consulta es un update")
	case sintaxsql.DeleteSQL:
		fmt.Println("la consulta es un delete")
	default:
		fmt.Println("Error en la sintaxis")
	}
}

//
//
//
//
//
//
//
//
//
//
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
