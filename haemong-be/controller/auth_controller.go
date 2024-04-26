package controller

import (
	"github.com/gofiber/fiber/v2"
	"haemong-be/dto/auth/request"
	"haemong-be/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{
		authService: s,
	}
}

func (r *AuthController) SignIn(c *fiber.Ctx) error {
	dto := new(request.SignInReqDTO)
	_ = c.BodyParser(dto)
	err := r.authService.SignIn(dto.UserId, dto.Password)
	if err != nil { // 없는 계정 or 비밀번호가 틀리면 401 해주기(나중에)
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).SendString("로그인 성공")
}
