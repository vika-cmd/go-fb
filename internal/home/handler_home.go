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
	hn.router.Get("/", hn.home)
}

func (hn *HandlerHome) home(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#home")
	return c.SendString("HandlerHome#home")
}