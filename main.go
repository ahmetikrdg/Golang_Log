package main

import (
	"Log/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type Customer struct {
	Name     string `json:"full_name"`
	Username string `json:"username"`
	Phone    string `json:"phone_number"`
}

var customers = []Customer{ //fake data
	{Name: "Ahmet Karadag", Username: "ahmetikrdg", Phone: "05433434343"},
	{Name: "Tony Stark", Username: "IronMan", Phone: "3423423423"},
}

func main() {
	app := fiber.New()

	app.Get("/customers", GetAllCustomer)
	app.Get("/customer/:username", GetByUsernameWithData)

	logger.Info("application will listen")

	app.Listen(":8000")
}

func GetAllCustomer(c *fiber.Ctx) error {
	logger.Info("listening is successful", zap.Bool("situation", true))
	return c.Status(http.StatusOK).JSON(customers)
}

func GetByUsernameWithData(c *fiber.Ctx) error {
	username := c.Params("username")

	for _, element := range customers {
		if username == element.Username {
			logger.Info("user found")
			return c.Status(http.StatusOK).JSON(element)
		}
	}

	logger.Error("user not found")
	return c.SendString("Username you were looking for was not found :(")
}
