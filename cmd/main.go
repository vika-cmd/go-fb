package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"app/go-fb/config"
	"app/go-fb/internal/home"
	"app/go-fb/internal/note"
)

func main(){
	start()
	app := fiber.New()
	app.Use(recover.New())
	app.Static("/public", "./public")

	//Handlers
	home.NewHandlerHome(app)
	note.NewHandlerNote(app)
	
	app.Listen("localhost:3031")
}

func start(){
	config.Init()
	dbConf :=config.NewDatebaseConfig()
	fmt.Println(dbConf)
/* 	tumbler := config.NewFlag()
	fmt.Println(tumbler) */
	val := config.NewCurrentVal()
	fmt.Println(val)
}