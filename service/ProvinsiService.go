package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type ProvinsiService interface {
	FindAll() ([]model.Provinsi, error)
	FindById(id string) (model.Provinsi, error)
}

type provinsiService struct {
	provinsiRepository repository.ProvinsiRepository
}

func NewProvinsiService(provinsiRepository repository.ProvinsiRepository) *provinsiService {
	return &provinsiService{provinsiRepository}
}

func (s *provinsiService) FindAll() ([]model.Provinsi, error) {
	var provinsis, err = s.provinsiRepository.FindAll()

	return provinsis, err
}

func (s *provinsiService) FindById(id string) (model.Provinsi, error) {
	var provinsi, err = s.provinsiRepository.FindById(id)

	return provinsi, err
}
