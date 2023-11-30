package presentation

import (
	"net/http"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"

	"github.com/gofiber/fiber/v2"
)

func (p *Presentation) FindAll(c *fiber.Ctx) error {
	var query dto.FindAllQuery
	if err := c.QueryParser(&query); err != nil {
		var out primitive.BaseResponse
		out.Status = primitive.ResponseStatusBadRequest
		out.Message = "invalid query"
		out.Data = struct{}{}

		c.Status(http.StatusBadRequest)
		return c.JSON(out)
	}

	out := p.Service.FindAll(c.Context(), query)

	out.Message = out.GetMessage()

	c.Status(out.GetCode())
	return c.JSON(out)
}
