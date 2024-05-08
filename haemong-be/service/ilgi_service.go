package service

import (
	"github.com/google/uuid"
	"haemong-be/repository"
)

type IlgiService struct {
	repo *repository.IlgiRepositroy
}

func NewIlgiService(repo *repository.IlgiRepositroy) *IlgiService {
	return &IlgiService{
		repo: repo,
	}
}

func (r *IlgiService) SaveIlgi(title, content, date, weather string) error {
	ilgiId := uuid.New()
	return r.repo.SaveIlgi(ilgiId.String(), title, content, date, weather)
}

func (r *IlgiService) ModifyIlgi(ilgiId, title, content, date, weather string) error {
	return r.repo.ModifyIlgi(ilgiId, title, content, date, weather)
}
