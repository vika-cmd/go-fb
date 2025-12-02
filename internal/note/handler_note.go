package note

import (
	//"context"
	"fmt"
	//"os"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"

	//"app/go-fb/views"
	"app/go-fb/pkg/templadapter"
	"app/go-fb/pkg/validator"
	"app/go-fb/views/components"
	"app/go-fb/views/widgets"
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

	form := ModelNoteForm{
		Task: c.FormValue("task"),
	}
	errors :=validate.Validate(
		&validators.StringIsPresent{Name: "task", Field: c.FormValue("task"), Message: "filling error!"},
	)
	//task := c.FormValue("task")
	fmt.Println("HandlerNote#add- task ",form.Task)
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NOTIFICATION_FAIL)
		return templadapter.Render(c, component,200)
	} else {
		component = components.Notification("Ok", components.NOTIFICATION_SUCCESS)
		return templadapter.Render(c, component,200)
	}

	//return c.SendString("HandlerNote")	
}
