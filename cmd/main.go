package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"

	"app/go-fb/internal/home"
)

func main(){
	app := fiber.New()

	home.NewHandlerHome(app)

	fmt.Println("Start")
	app.Listen("localhost:3031")
}