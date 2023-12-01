package presentation

import (
	"net/http"
	"strconv"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (p *Presentation) UpdateById(c *fiber.Ctx) error {
	paramId := utils.CopyString(c.Params("id"))
	var req dto.UpdateByIdReq

	todoId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		var out primitive.BaseResponse
		out.Status = primitive.ResponseStatusBadRequest
		out.Message = "invalid param"
		out.Data = struct{}{}

		c.Status(http.StatusBadRequest)
		return c.JSON(out)
	}
	if err := c.BodyParser(&req); err != nil {
		var out primitive.BaseResponse
		out.Status = primitive.ResponseStatusBadRequest
		out.Message = "invalid body"
		out.Data = struct{}{}

		c.Status(http.StatusBadRequest)
		return c.JSON(out)
	}

	out := p.Service.UpdateById(c.Context(), todoId, req)
	out.Message = out.GetMessage()
	if out.GetCode() >= 400 {
		out.Data = struct{}{}
	}

	c.Status(out.GetCode())
	return c.JSON(out)
}
