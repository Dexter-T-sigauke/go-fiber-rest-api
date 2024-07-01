package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/models"
)

type User struct {
	//
	ID        uuid.UUID   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.ID = uuid.New()
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id uuid.UUID, user *models.User) error {
	database.Database.Db.First(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	if err := findUser(uuid, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	err = findUser(uuid, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is a valid UUID")
	}

	err = findUser(uuid, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}