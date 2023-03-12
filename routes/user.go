package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/karpov-dmitry-py/fiber-api/database"
	"github.com/karpov-dmitry-py/fiber-api/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getResponseUser(user models.User) User {
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
	responseUser := getResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var (
		users         []models.User
		responseUsers []User
	)

	database.Database.Db.Find(&users)
	for _, v := range users {
		responseUsers = append(responseUsers, getResponseUser(v))
	}

	result := map[string]interface{}{
		"total_count": len(responseUsers),
		"items":       responseUsers,
	}

	return c.Status(200).JSON(result)
}

func GetUser(c *fiber.Ctx) error {
	var (
		user models.User
	)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("user id must be passed in")
	}

	if err = findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := getResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var (
		user models.User
	)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("user id must be passed in")
	}

	if err = findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var requestUser updateUser

	if err = c.BodyParser(&requestUser); err != nil {
		return c.Status(400).JSON("first_name and last_name must be passed in")
	}

	user.FirstName = requestUser.FirstName
	user.LastName = requestUser.LastName

	database.Database.Db.Save(&user)

	responseUser := getResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	var (
		user models.User
	)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("user id must be passed in")
	}

	if err = findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("user deleted successfully")
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return fmt.Errorf("user with id %d does not exist in db", id)
	}

	return nil
}
