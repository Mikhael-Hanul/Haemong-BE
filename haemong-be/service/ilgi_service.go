package service

import (
	"github.com/google/uuid"
	"haemong-be/dto/ilgi/response"
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

func (r *IlgiService) DeleteIlgi(ilgiId string) error {
	return r.repo.DeleteIlgi(ilgiId)
}

func (r *IlgiService) SearchIlgi(keyword string) ([]response.SearchIlgiResDTO, error) {
	ilgis, err := r.repo.SearchIlgi(keyword)
	if err != nil {
		return nil, err
	}

	var ilgisDTO []response.SearchIlgiResDTO
	for _, i := range ilgis {
		ilgiDTO := response.SearchIlgiResDTO{
			IlgiId:  i.IlgiId,
			Title:   i.Title,
			Content: i.Content,
			Date:    i.Date,
			Weather: i.Weather,
		}
		ilgisDTO = append(ilgisDTO, ilgiDTO)
	}

	return ilgisDTO, nil
}
