package service

import (
	"fmt"
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

func (r *UserService) SignUp(userName, userId, password string) error {
	isUserIdDuplicate := r.repo.IsUserIdDuplicate(userId)
	if isUserIdDuplicate {
		return fmt.Errorf("아이디 중복")
	}
	return r.repo.SaveUser(userId, userName, password)
}
