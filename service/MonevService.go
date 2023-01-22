package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type MonevService interface {
	FindAll(kabupatenKotaId string) ([]model.Monev, error)
	FindById(kabupatenKotaId string, id int) (model.Monev, error)
}

type monevService struct {
	monevRepository repository.MonevRepository
}

func NewMonevService(monevRepository repository.MonevRepository) *monevService {
	return &monevService{monevRepository}
}

func (s *monevService) FindAll(kabupatenKotaId string) ([]model.Monev, error) {
	var monevs, err = s.monevRepository.FindAll(kabupatenKotaId)

	return monevs, err
}

func (s *monevService) FindById(kabupatenKotaId string, id int) (model.Monev, error) {
	var monev, err = s.monevRepository.FindById(kabupatenKotaId, id)

	return monev, err
}
