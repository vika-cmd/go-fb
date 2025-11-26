package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main(){
	app := fiber.New()
	app.Get("/", func (c *fiber.Ctx)  error{
		return c.SendString("Hi")
	})
	fmt.Println("Start")
	app.Listen("localhost:3031")
}