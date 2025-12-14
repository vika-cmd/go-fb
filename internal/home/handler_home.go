package home

import (
	//"context"
	"fmt"	
	"net/http"

	//"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"

	"app/go-fb/internal/note"
	"app/go-fb/views"
	"app/go-fb/pkg/templadapter"
	//"app/go-fb/views/components"
	"app/go-fb/hlp"
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

//LIMIT 20 OFFSET 40 вернет 20 записей, начиная с 41-й
//http://localhost:3031/?page=2
func (hn *HandlerHome) home(c *fiber.Ctx) error {
	var component templ.Component
	LIMIT := 2 //on page
	pageFromUrl := c.QueryInt("page", 1) //currentPage   default 1, page-1 -first   1	
	offset := LIMIT * (pageFromUrl - 1)//skip   0	
	notes, err := hn.repositoryNote.GetAll(LIMIT, offset)
	if err !=nil {
		fmt.Printf("HandlerHome#home-notes from repo %v\n",err)
		c.SendStatus(500)
	}

	total := hn.repositoryNote.TotalRecords()  //3	
	quantityOfPages := hlp.DivideAndCeil(total,LIMIT) //2	
	//fmt.Printf("HandlerHome#quantityOfPages %d\n",quantityOfPages)
	
	//component := components.Title(components.TitleProps{Message: "HandlerHome#h"})
	component = views.MainPage(notes,  quantityOfPages, pageFromUrl)
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