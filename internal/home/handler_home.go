package home

import (
	//"context"
	"fmt"
	"net/http"
	//"os"

	"github.com/gofiber/fiber/v2"
	"github.com/a-h/templ"

	"app/go-fb/internal/note"
	"app/go-fb/views"
	//"app/go-fb/views/components"
	"app/go-fb/pkg/templadapter"	
)

type HandlerHome struct {
	router fiber.Router
	repositoryNote *note.RepositoryNote
}

func NewHandlerHome(router fiber.Router, repositoryNote *note.RepositoryNote) {
	hn := &HandlerHome{
		router: router,
		repositoryNote: repositoryNote,
	}	
/* 	api := hn.router.Group("/api")
	api.Get("/", hn.home)
	api.Get("/error", hn.errorPage) */
	hn.router.Get("/", hn.home)
	hn.router.Get("/404", hn.errorPage)
	hn.router.Get("/contact", hn.getContact)
}

/* func (hn *HandlerHome) contact(c *fiber.Ctx) error {
	redir := "https://yandex.ru/maps/39/rostov-na-donu/?ll=39.709844%2C47.226728&mode=poi&poi%5Bpoint%5D=39.732830%2C47.224679&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D1138347645&z=13"
	return c.Redirect(redir, 200)
} */

func (hn *HandlerHome) home(c *fiber.Ctx) error {
	var component templ.Component
	QUANTITY_ON_PAGE := 2 //limit
	page := c.QueryInt("page", 1) //default 1, page-1 -first
	//http://localhost:3031/?page=2
	quantityOfNotes := QUANTITY_ON_PAGE * (page - 1)// all notes
	notes, err := hn.repositoryNote.GetAll(QUANTITY_ON_PAGE, quantityOfNotes)
	if err !=nil {
		fmt.Printf("HandlerHome#home-notes from repo %v\n",err)
		c.SendStatus(500)
	}
	
	//component := components.Title(components.TitleProps{Message: "HandlerHome#h"})
	component = views.MainPage(notes)
	//component.Render(context.Background(), os.Stdout) //cod in consol
	//return c.SendString("HandlerHome#home") //page

	return templadapter.Render(c, component, http.StatusOK)
}

func (hn *HandlerHome) errorPage(c *fiber.Ctx) error {
	fmt.Println("HandlerHome#errorPage")
	//return c.SendString("HandlerHome#errorPage")
	//return fiber.ErrBadGateway
	//return fiber.NewError(400, "Its mess about my err !")
	return fiber.NewError(fiber.StatusBadRequest, "Its mess about my err !")
}

func (hn *HandlerHome) getContact(c *fiber.Ctx) error{
	fmt.Println("HandlerHome# getContact")
	return nil
}