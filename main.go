package main

import (
	"go-resto/app/helpers"
	"go-resto/app/routes"
	database "go-resto/db"
	"go-resto/exception"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	//setup env
	var env *viper.Viper = viper.New()
	env.SetConfigFile(".env")
	env.AddConfigPath(".")
	env.AutomaticEnv()

	//setup DB
	err := database.NewDB(env)
	if err != nil {
		panic(err)
	}

	err = env.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Prefork:      true,
		ErrorHandler: exception.ErrorHandler,
	})

	routes.Routes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return helpers.ErrResponse(c, "Not Found", fiber.StatusNotFound) // => 404 "Not Found"
	})

	err = app.Listen(env.GetString("APP_ADDRESS"))
	log.Fatal(err)
}
