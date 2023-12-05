package CategoryController

import (
	"go-resto/app/helpers"
	"go-resto/app/models"
	database "go-resto/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// var (
// 	validate *validator.Validate = validator.New()
// )

func Index(c *fiber.Ctx) error {
	var categories []models.Category
	if err := database.DB.Preload("Menus").Find(&categories).Error; err != nil {
		log.Errorf("category index: %v", err)
		return err
	}
	return helpers.OkResponse(c, "Get data successfully", categories)
}

func Create(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		log.Errorf("category create: %v", err)
		return err
	}

	if err := helpers.Validate.Struct(category); err != nil {
		log.Errorf("category create: %v", err)
		return err
	}

	if err := database.DB.Create(&category).Error; err != nil {
		log.Errorf("category create: %v", err)
		return err
	}

	return helpers.OkResponse(c, "Create data successfully", category)
}

func Update(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		log.Errorf("category update: %v", err)
		return err
	}
	if err := helpers.Validate.Struct(category); err != nil {
		log.Errorf("category update: %v", err)
		return err
	}
	id := c.Params("id")
	err := database.DB.Model(&models.Category{}).Where("id = ?", id).Update("name", category.Name).Error
	if err != nil {
		log.Errorf("category update: %v", err)
		return err
	}

	return helpers.OkResponse(c, "Update data successfully", nil)
}

func Show(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := database.DB.Take(&category, "id = ?", c.Params("id")).Error; err != nil {
		log.Errorf("category show: %v", err)
		return err
	}
	return helpers.OkResponse(c, "Get data successfully", category)
}

func Delete(c *fiber.Ctx) error {
	if err := database.DB.Delete(&models.Category{}, "id = ?", c.Params("id")).Error; err != nil {
		log.Errorf("category create: %v", err)
		return err
	}
	return helpers.OkResponse(c, "Delete data successfully", nil)
}
