package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/ai/request"
	"haemong-be/dto/ai/response"
	"haemong-be/service"
)

type AiController struct {
	aiService *service.AiService
}

func NewAiController(s *service.AiService) *AiController {
	return &AiController{
		aiService: s,
	}
}

func (r *AiController) Haemong(c *fiber.Ctx) error {
	dto := new(request.HaemongReqDTO)

	_ = c.BodyParser(dto)
	res, err := r.aiService.Haemong(dto.Content)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	resp := response.HaemongResDTO{
		Gotjr: res,
	}

	return c.Status(200).JSON(resp)
}
