package main

import (
	"fmt"

	"covid-19-api/cmd/build"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/sirupsen/logrus"
)

func main() {
	app, appConfig := SetupApp()

	logrus.Fatal(app.Listen(fmt.Sprintf(":%v", appConfig.Config.APP_PORT)))
}

func SetupApp() (*fiber.App, *build.ApplicationImp) {
	// config.PrintConfig()

	appConfig, err := build.Wire(".env")
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{})

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recover.New())

	InitRouter(app, appConfig.Config.APP_PREFIX, appConfig)

	return app, appConfig
}

func InitRouter(app *fiber.App, apiPrefix string, appConfig *build.ApplicationImp) {
	var router fiber.Router

	if prefix := apiPrefix; len(prefix) != 0 {
		router = app.Group(prefix)
	} else {
		router = app
	}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("up")
	})

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Error 404: Not Found",
			"status":  false,
		})
	})
}
