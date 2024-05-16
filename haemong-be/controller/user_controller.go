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

func (r *UserController) SignUp(c *fiber.Ctx) error {
	dto := new(request.SignUpReqDTO)
	_ = c.BodyParser(dto)
	err := r.userService.SignUp(dto.Name, dto.UserId, dto.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("회원가입 성공")
}

func (r *UserController) ChangePassword(c *fiber.Ctx) error {
	dto := new(request.ChangePasswordReqDTO)
	_ = c.BodyParser(dto)
	err := r.userService.ChangePassword(dto.UserId, dto.Password, dto.NewPassword)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("비밀번호 변경 성공")
}

func (r *UserController) Withdrawal(c *fiber.Ctx) error {
	userId := c.Params("userId")
	password := c.Params("password")
	err := r.userService.Withdrawal(userId, password)
	if err != nil {
		c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("회원 탈퇴 성공")
}
