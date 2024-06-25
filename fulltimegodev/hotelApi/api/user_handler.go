package api

import (
	"github.com/alikazai/golearn/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Ali",
		LastName:  "Kazai",
	}
	return c.JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("Ali", "James")
}
