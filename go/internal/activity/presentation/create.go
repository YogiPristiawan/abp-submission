package presentation

import (
	"net/http"
	"todo/internal/activity/dto"
	"todo/internal/shared/primitive"

	"github.com/gofiber/fiber/v2"
)

func (p *Presentation) Create(c *fiber.Ctx) error {
	var req dto.CreateReq
	if err := c.BodyParser(&req); err != nil {
		var out primitive.BaseResponse
		out.Status = primitive.ResponseStatusBadRequest
		out.Message = "invalid body"
		out.Data = struct{}{}

		c.Status(http.StatusBadRequest)
		return c.JSON(out)
	}

	out := p.Service.Create(c.Context(), req)

	c.Status(out.GetCode())
	return c.JSON(out)
}
