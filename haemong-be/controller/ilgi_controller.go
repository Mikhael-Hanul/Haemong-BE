package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/ilgi/request"
	"haemong-be/service"
)

type IlgiController struct {
	ilgiService *service.IlgiService
}

func NewIlgiController(s *service.IlgiService) *IlgiController {
	return &IlgiController{
		ilgiService: s,
	}
}

func (r *IlgiController) SaveIlgi(c *fiber.Ctx) error {
	dto := new(request.SaveIlgiReqDTO)
	_ = c.BodyParser(dto)
	err := r.ilgiService.SaveIlgi(dto.Title, dto.Content, dto.Date, dto.Weather)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("일기 저장 성공")

}
