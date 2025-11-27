package home

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

type HandlerHome struct {
	router fiber.Router
}

func NewHandlerHome(router fiber.Router) {
	hn := &HandlerHome{
		router: router,
	}
	//hn.router.Get("/", hn.home)
	api := hn.router.Group("/api")
	api.Get("/", hn.home)
	api.Get("/error", hn.errorPage)
}

func (hn *HandlerHome) home(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#home")
	return c.SendString("HandlerHome#home")
}

func (hn *HandlerHome) errorPage(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#errorPage")
	return c.SendString("HandlerHome#errorPage")
}