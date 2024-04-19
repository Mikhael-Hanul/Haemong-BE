package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{
		userService: s,
	}
}

func (r *UserController) UserController(c *fiber.Ctx) error {
	userName := "유찬홍"
	userId := "chanhong1206"
	password := "password"
	err := r.userService.SignUp(userName, userId, password)
	if err != nil {
		c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("회원가입 성공")
}
