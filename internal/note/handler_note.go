package note

import (
	//"context"
	"fmt"
	"net/http"
	"time"

	//"os"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"

	"app/go-fb/hlp"
	"app/go-fb/pkg/templadapter"
	"app/go-fb/pkg/validator"
	"app/go-fb/views/components"
	"app/go-fb/views/entitypages"
)


type HandlerNote struct {
	router fiber.Router
	repositoryNote *RepositoryNote
}

func NewHandlerNote(router fiber.Router, repositoryNote *RepositoryNote) {
	hn := &HandlerNote{
		router: router,
		repositoryNote: repositoryNote,
	}	
	noteGr := hn.router.Group("/note")
	noteGr.Get("/add", hn.addNote)
	noteGr.Get("/getAlljson", hn.getAlljson)

	noteGr.Post("/", hn.fromform)

}


func (hn *HandlerNote) addNote(c *fiber.Ctx) error {
	component := entitypages.NoteAddForm()
	return templadapter.Render(c, component,200)
}

func (hn *HandlerNote) fromform(c *fiber.Ctx) error {
	var component templ.Component

	strPriority :=c.FormValue("priority")
	priority, err := hlp.ConvStrToInt(strPriority)
	if err != nil {
		fmt.Printf("HandlerNote#fromform: priority -ConvStrToInt err-%v \n", err)
	}
	//fmt.Printf("HandlerNote#fromform: type uPriority-%T \n", priority)

	strByDate := c.FormValue("byDate")
	date, err := hlp.ConvStrToDate(strByDate)
	if err != nil {
		fmt.Printf("HandlerNote#fromform: byDate -ConvStrToDate err-%v \n", err)
	}
	//fmt.Printf("HandlerNote#fromform: type date-%T \n", date)

	form := ModelNoteForm {
		Task: c.FormValue("task"),
		Category: c.FormValue("category"),
		Priority: priority,
		Description: c.FormValue("description"),
		ByDate: date,
		Createdat: time.Now(),
	}
	errors :=validate.Validate(
		&validators.StringIsPresent{Name: "Заметка", Field: c.FormValue("task"), Message: "filling error (некорректный ввод)"},
		&validators.StringIsPresent{Name: "Категория", Field: c.FormValue("category"), Message: "filling error (некорректный ввод)"},
		&validators.StringIsPresent{Name: "Приоритет", Field: c.FormValue("priority"), Message: "filling error (некорректный ввод)"},
		&validators.StringIsPresent{Name: "Описание", Field: c.FormValue("description"), Message: "filling error (некорректный ввод)"},
		&validators.StringIsPresent{Name: "Дата", Field: c.FormValue("byDate"), Message: "filling error (некорректный ввод)"},
	)

	//fmt.Println("HandlerNote#add- fields ",form.Task,form.Category, form.Priority, form.Description, form.ByDate)

	// sleep -> form class "htmx-request"
	time.Sleep(time.Second * 1)

	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NOTIFICATION_FAIL)
		return templadapter.Render(c, component,http.StatusBadRequest)
	}

	err = hn.repositoryNote.addNote(form)
	if err != nil {
		component = components.Notification("Ошибка при сохранении заметки", components.NOTIFICATION_FAIL)
		fmt.Println("HandlerNote#add- err from repo")
		return templadapter.Render(c, component,http.StatusBadRequest)
	}
	component = components.Notification("Ok", components.NOTIFICATION_SUCCESS)
	return templadapter.Render(c, component, http.StatusOK)
	

	//return c.SendString("HandlerNote")	
}


func (hn *HandlerNote) getAlljson(c *fiber.Ctx) error {
	LIMIT := 2 //on page
	pageFromUrl := c.QueryInt("page", 1) //currentPage   default 1, page-1 -first	
	offset := LIMIT * (pageFromUrl - 1)//skip
	notes, err := hn.repositoryNote.GetAll(LIMIT, offset)
	if err !=nil {
		fmt.Printf("HandlerNote#getAlljson- from repo %v\n",err)
	}
	return c.JSON(notes)
}
