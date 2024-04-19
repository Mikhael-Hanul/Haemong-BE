package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"haemong-be/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) SignUp(c *fiber.Ctx) error {
	userName := "유찬홍"
	userId := "chanhong1206"
	password := "password"
	isUserIdDuplicate := r.repo.IsUserIdDuplicate(userId)
	if isUserIdDuplicate {
		return fmt.Errorf("아이디 중복")
	}
	return r.repo.SaveUser(userId, userName, password)
}
