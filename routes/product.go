package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/models"
)

type Product struct {
	//serializer
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.ID = uuid.Must(uuid.NewV4())
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id uuid.UUID, product *models.Product) error {
	database.Database.Db.First(&product, "id = ?", id)
	if product.ID == uuid.Nil {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.FromString(id)

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	if err := findProduct(uuid, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.FromString(id)

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	err = findProduct(uuid, &product)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.FromString(id)

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	err = findProduct(uuid, &product)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Product")
}
