package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	fmt.Println("Hello World!")

	app.Listen(":7500")
}
