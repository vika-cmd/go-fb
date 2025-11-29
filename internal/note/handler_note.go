package note

import (
	//"context"
	"fmt"
	//"os"

	"github.com/gofiber/fiber/v2"

	//"app/go-fb/views"
	//"app/go-fb/views/components"
	//"app/go-fb/pkg/templadapter"
)


type HandlerNote struct {
	router fiber.Router
}

func NewHandlerNote(router fiber.Router) {
	hn := &HandlerNote{
		router: router,
	}	
	noteGr := hn.router.Group("/note")
	noteGr.Post("/add", hn.add)

}

func (hn *HandlerNote) add(c *fiber.Ctx) error {
	fmt.Println("HandlerNote#add")
	return c.SendString("HandlerNote")
}
