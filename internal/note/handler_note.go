package note

import (
	//"context"
	"fmt"
	//"os"

	"github.com/gofiber/fiber/v2"
	"github.com/a-h/templ"

	//"app/go-fb/views"
	"app/go-fb/views/components"
	"app/go-fb/views/widgets"
	"app/go-fb/pkg/templadapter"
)


type HandlerNote struct {
	router fiber.Router
}

func NewHandlerNote(router fiber.Router) {
	hn := &HandlerNote{
		router: router,
	}	
	noteGr := hn.router.Group("/note")
	noteGr.Get("/add", hn.addNote)

	noteGr.Post("/", hn.fromform)

}

func (hn *HandlerNote) addNote(c *fiber.Ctx) error {
	component := widgets.NoteAddForm()
	return templadapter.Render(c, component,200)
}

func (hn *HandlerNote) fromform(c *fiber.Ctx) error {
	var component templ.Component
	task := c.FormValue("task")
	fmt.Println("HandlerNote#add- task ", task)
	if task == ""{
		component = components.Notification("task - err!", components.NOTIFICATION_FAIL)
		return templadapter.Render(c, component,200)
	} else {
		component = components.Notification("Ok", components.NOTIFICATION_SUCCESS)
		return templadapter.Render(c, component,200)
	}

	//return c.SendString("HandlerNote")	
}
