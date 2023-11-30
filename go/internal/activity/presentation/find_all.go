package presentation

import (
	"github.com/gofiber/fiber/v2"
)

func (p *Presentation) FindAll(c *fiber.Ctx) error {
	out := p.Service.FindAll(c.Context())

	out.Message = out.GetMessage()

	c.Status(out.GetCode())
	return c.JSON(out)
}
