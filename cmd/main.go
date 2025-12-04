package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"app/go-fb/config"
	"app/go-fb/pkg/database"
	"app/go-fb/internal/home"
	"app/go-fb/internal/note"
)

func main(){
	start()
	app := fiber.New()
	app.Use(recover.New())
	app.Static("/public", "./public")

/* 	dbConf :=config.NewDatebaseConfig() //2
	fmt.Println(dbConf)
	dbpool :=database.CreateDbPool(dbConf)// 3
	defer dbpool.Close()// 4 */
	db()
	
	//Handlers
	home.NewHandlerHome(app)
	note.NewHandlerNote(app)
	
	app.Listen("localhost:3031")
}

func start(){
	config.Init() //1

/* 	tumbler := config.NewFlag()
	fmt.Println(tumbler) */
	val := config.NewCurrentVal()
	fmt.Println(val)
}

func db(){
	dbConf :=config.NewDatebaseConfig() //2
	fmt.Println(dbConf)
	dbpool :=database.CreateDbPool(dbConf)// 3
	defer dbpool.Close()// 4	
}