package presentation

import (
	"net/http"
	"strconv"
	"todo/internal/shared/primitive"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (p *Presentation) DeleteById(c *fiber.Ctx) error {
	paramId := utils.CopyString(c.Params("id"))
	activityId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		var out primitive.BaseResponse
		out.Status = primitive.ResponseStatusBadRequest
		out.Message = "invalid param"
		out.Data = struct{}{}

		c.Status(http.StatusBadRequest)
		return c.JSON(out)
	}

	out := p.Service.DeleteById(c.Context(), activityId)
	out.Message = out.GetMessage()
	out.Data = struct{}{}

	c.Status(out.GetCode())
	return c.JSON(out)
}
