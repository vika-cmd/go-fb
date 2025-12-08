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
	startEnv()
/* 	config.Init() //1
	val := config.NewCurrentVal()
	fmt.Println(val) */

	app := fiber.New()
	app.Use(recover.New())
	app.Static("/public", "./public")

	//dbpool
	dbConf :=config.NewDatebaseConfig() //2
	fmt.Println(dbConf)
	dbpool :=database.CreateDbPool(dbConf)// 3
	defer dbpool.Close()// 4	

	//repo
	repoNote := note.NewRepositoryNote(dbpool)
	
	//Handlers
	home.NewHandlerHome(app, repoNote)
	note.NewHandlerNote(app,repoNote)
	
	app.Listen("localhost:3031")
}

func startEnv(){
	config.Init()  ///1

// 	tumbler := config.NewFlag()
//	fmt.Println(tumbler) 
 	val := config.NewCurrentVal()
	fmt.Println(val)
}
