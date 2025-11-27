package main

import (	
	"github.com/gofiber/fiber/v2"
	"fmt"

	"app/go-fb/internal/home"
	"app/go-fb/config"
)

func main(){
	start()
	app := fiber.New()

	home.NewHandlerHome(app)
	
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