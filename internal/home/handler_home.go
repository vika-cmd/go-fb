package home

import (
	//"context"
	"fmt"
	//"os"

	"github.com/gofiber/fiber/v2"

	"app/go-fb/views"
	//"app/go-fb/views/components"
	"app/go-fb/pkg/templadapter"
)

type HandlerHome struct {
	router fiber.Router
}

func NewHandlerHome(router fiber.Router) {
	hn := &HandlerHome{
		router: router,
	}	
/* 	api := hn.router.Group("/api")
	api.Get("/", hn.home)
	api.Get("/error", hn.errorPage) */
	hn.router.Get("/", hn.home)
	hn.router.Get("/404", hn.errorPage)
}

func (hn *HandlerHome) home(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#home")
	//component := components.Title(components.TitleProps{Message: "HandlerHome#h"})
	component := views.MainPage()
	//component.Render(context.Background(), os.Stdout) //cod in consol
	//return c.SendString("HandlerHome#home") //page

	return templadapter.Render(c, component,200)
}

func (hn *HandlerHome) errorPage(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#errorPage")
	//return c.SendString("HandlerHome#errorPage")
	//return fiber.ErrBadGateway
	//return fiber.NewError(400, "Its mess about my err !")
	return fiber.NewError(fiber.StatusBadRequest, "Its mess about my err !")
}