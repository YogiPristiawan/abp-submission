package todo

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ITodoPresentation interface {
	FindAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
	UpdateById(c *fiber.Ctx) error
}

func (t Todo) Route(app *fiber.App) {
	if app == nil {
		log.Fatal("[x] app is required on todo module")
	}

	app.Get("/todo-items", t.Presentation.FindAll)
	app.Get("/todo-items/:id", t.Presentation.GetById)
	app.Post("/todo-items", t.Presentation.Create)
	app.Delete("/todo-items/:id", t.Presentation.DeleteById)
	app.Patch("/todo-items/:id", t.Presentation.UpdateById)
}
