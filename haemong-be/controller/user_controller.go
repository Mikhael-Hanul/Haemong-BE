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
	return nil
}
