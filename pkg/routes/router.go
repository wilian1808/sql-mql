package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wilian1808/sqlmql/pkg/helpers"
	"github.com/wilian1808/sqlmql/pkg/models"
	"github.com/wilian1808/sqlmql/pkg/validate"
)

var ResponseMQL = ""

// IndexPage func .
func IndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"data": "hola",
	})
}

// GetQuery func .
func GetQuery(c *fiber.Ctx) error {
	query := models.Query{}
	if err := c.BodyParser(&query); err != nil {
		return err
	}

	// obtenemos los valores de la consulta
	arr, err := helpers.FormatData(query.Data)
	if err != nil {
		return err
	}

	// tipos recibidos
	helpers.GetTypes(arr)
	fmt.Printf("\n\n")

	// para ver los tipos de dato
	// helpers.GetTypes(arr)
	result, err := validate.RedirectTransform(arr)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", result)
	ResponseMQL = result
	return nil
}

func ResultMQL(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": ResponseMQL,
	})
}
