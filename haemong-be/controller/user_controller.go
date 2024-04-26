package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/user/request"
	"haemong-be/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{
		userService: s,
	}
}

func (r *UserController) UserController(c *fiber.Ctx) error {
	dto := new(request.UserReqDTO)
	_ = c.BodyParser(dto)
	err := r.userService.SignUp(dto.Name, dto.UserId, dto.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("회원가입 성공")
}
