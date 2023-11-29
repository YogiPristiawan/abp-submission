package activity

import "github.com/gofiber/fiber/v2"

type IActivityPresentation interface {
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}

func Route(app *fiber.App) {

}
