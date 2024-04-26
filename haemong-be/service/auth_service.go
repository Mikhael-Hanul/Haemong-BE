package service

import (
	"errors"
	"haemong-be/repository"
)

type AuthService struct {
	auth *repository.AuthRepository
	user *repository.UserRepository
}

func NewAuthService(auth *repository.AuthRepository, user *repository.UserRepository) *AuthService {
	return &AuthService{
		auth: auth,
		user: user,
	}
}

func (r *AuthService) SignIn(userId, password string) error {
	isUserIdDuplicate := r.user.IsUserIdDuplicate(userId)
	if !isUserIdDuplicate {
		return errors.New("등록되지 않은 유저")
	}
	isUserPasswordMatching := r.auth.IsUserPasswordMatching(userId, password)
	if !isUserPasswordMatching {
		return errors.New("비밀번호가 맞지 않음")
	}
	return nil
}
