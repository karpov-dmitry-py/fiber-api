package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karpov-dmitry-py/fiber-api/database"
	"github.com/karpov-dmitry-py/fiber-api/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var (
		users         []models.User
		responseUsers []User
	)

	database.Database.Db.Find(&users)
	for _, v := range users {
		responseUsers = append(responseUsers, CreateResponseUser(v))
	}

	result := map[string]interface{}{
		"total_count": len(responseUsers),
		"items":       responseUsers,
	}

	return c.Status(200).JSON(result)
}
