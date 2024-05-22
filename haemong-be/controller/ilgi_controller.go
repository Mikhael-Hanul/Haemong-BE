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

func (r *IlgiController) ModifyIlgi(c *fiber.Ctx) error {
	dto := new(request.ModifyIlgiReqDTO)
	_ = c.BodyParser(dto)
	err := r.ilgiService.ModifyIlgi(dto.IlgiId, dto.Title, dto.Content, dto.Date, dto.Weather)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("일기 수정 성공")
}

func (r *IlgiController) DeleteIlgi(c *fiber.Ctx) error {
	ilgiId := c.Params("id")
	err := r.ilgiService.DeleteIlgi(ilgiId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("일기 삭제 성공")
}

func (r *IlgiController) SearchIlgi(c *fiber.Ctx) error {
	keyword := c.Params("search")
	ilgis, err := r.ilgiService.SearchIlgi(keyword)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(ilgis)
}
