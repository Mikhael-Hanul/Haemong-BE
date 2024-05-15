package service

import (
	"errors"
	"haemong-be/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) SignUp(name, userId, password string) error {
	isUserIdDuplicate := r.repo.IsUserIdDuplicate(userId)
	if isUserIdDuplicate {
		return errors.New("아이디 중복")
	}
	return r.repo.SaveUser(userId, password, name)
}

func (r *UserService) ChangePassword(userId, password, newPassword string) error {
	userPassword, err := r.repo.FindPasswordById(userId)
	if err != nil {
		return err
	}
	if userPassword != password {
		return errors.New("비밀번호가 일치하지 않습니다.")
	}
	return r.repo.ChangeUserPassword(userId, newPassword)
}
