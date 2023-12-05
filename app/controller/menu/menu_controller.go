package MenuController

import (
	"go-resto/app/helpers"
	"go-resto/app/models"
	database "go-resto/db"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// var (
// 	validate *validator.Validate = validator.New()
// )

func Index(c *fiber.Ctx) error {
	var menus []models.Menu
	if err := database.DB.Joins("Category").Find(&menus).Error; err != nil {
		log.Errorf("menu index: %v", err)
		return err
	}
	return helpers.OkResponse(c, "Get data successfully", menus)
}

func Create(c *fiber.Ctx) error {
	menu := new(models.Menu)
	if err := c.BodyParser(menu); err != nil {
		log.Errorf("menu create: %v", err)
		return err
	}

	file, err := c.FormFile("image")
	menu.IsAvailable = false
	menu.CategoryID = c.FormValue("category_id")
	menu.PriceModal, err = strconv.Atoi(c.FormValue("price_modal"))

	if err != nil {
		log.Errorf("menu create: %v", err)
		return err
	}

	filename := time.Now().Format("20060102150405") + "-" + file.Filename

	menu.Image = filename

	if err := helpers.Validate.Struct(menu); err != nil {
		log.Errorf("menu create: %v", err)
		return err
	}

	c.SaveFile(file, "./storage/menu-images/"+filename)

	if err := database.DB.Create(&menu).Error; err != nil {
		log.Errorf("menu create: %v", err)
		return err
	}

	return helpers.OkResponse(c, "Create data successfully", menu)
}

func Update(c *fiber.Ctx) error {

	menu := new(models.Menu)

	oldImage := database.DB.Select("image").Take(&models.Menu{}, "id = ?", c.Params("id")).Error

	if err := c.BodyParser(menu); err != nil {
		log.Errorf("menu update: %v", err)
		return err
	}

	filename := oldImage
	menu.CategoryID = c.FormValue("category_id")
	menu.PriceModal, _ = strconv.Atoi(c.FormValue("price_modal"))
	if err := helpers.Validate.Struct(menu); err != nil {
		log.Errorf("menu update: %v", err)
		return err
	}

	if file, _ := c.FormFile("image"); file != nil {
		// fmt.Println(file)
		filename = time.Now().Format("20060102150405") + "-" + file.Filename
		c.SaveFile(file, "./storage/menu-images/"+filename)
	}
	menu.Image = filename
	// if err != nil {
	// 	log.Errorf("menu create: %v", err)
	// 	return err
	// }

	// id := c.Params("id")
	// err := database.DB.Model(&models.Category{}).Where("id = ?", id).Update("name", menu.Name).Error
	// if err != nil {
	// 	log.Errorf("menu update: %v", err)
	// 	return err
	// }

	return helpers.OkResponse(c, "Update data successfully", nil)
}

func Show(c *fiber.Ctx) error {
	menu := new(models.Menu)
	if err := database.DB.Take(&menu, "id = ?", c.Params("id")).Error; err != nil {
		log.Errorf("menu show: %v", err)
		return err
	}

	return helpers.OkResponse(c, "Get Data Successfully", menu)
}

func Delete(c *fiber.Ctx) error {
	menu := new(models.Menu)
	if err := database.DB.Take(&menu, "id = ?", c.Params("id")).Error; err != nil {
		log.Errorf("Menu Delete: %v", err)
		return err
	}
	if err := os.Remove("./storage/menu-images/" + menu.Image); err != nil {
		log.Errorf("Menu Delete: %v", err)
		return err
	}
	if err := database.DB.Delete(&menu).Error; err != nil {
		log.Errorf("menu delete: %v", err)
		return err
	}
	return helpers.OkResponse(c, "Delete data successfully", nil)
}
