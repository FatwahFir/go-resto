package routes

import (
	CategoryController "go-resto/app/controller/category"
	MenuController "go-resto/app/controller/menu"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	category := app.Group("/category")
	category.Get("/", CategoryController.Index)
	category.Post("/", CategoryController.Create)
	category.Put("/:id", CategoryController.Update)
	category.Get("/:id", CategoryController.Show)
	category.Delete("/:id", CategoryController.Delete)

	menu := app.Group("/menu")
	menu.Get("/", MenuController.Index)
	menu.Get("/:id", MenuController.Show)
	menu.Post("/", MenuController.Create)
	menu.Put("/:id", MenuController.Update)
	menu.Delete("/:id", MenuController.Delete)
}
