package activity

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type IActivityPresentation interface {
	FindAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
	UpdateById(c *fiber.Ctx) error
}

func (a Activity) Route(app *fiber.App) {
	if app == nil {
		log.Fatal("[x] app is required on activitiy module")
	}

	app.Get("/activity-groups", a.Presentation.FindAll)
	app.Get("/activity-groups/:id", a.Presentation.GetById)
	app.Post("/activity-groups", a.Presentation.Create)
	app.Delete("/activity-groups/:id", a.Presentation.DeleteById)
	app.Patch("/activity-groups/:id", a.Presentation.UpdateById)
}
