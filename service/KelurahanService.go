package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KelurahanService interface {
	FindByKecamatanId(kecamatanId string) ([]model.Kelurahan, error)
	FindById(id string) (model.Kelurahan, error)
}

type kelurahanService struct {
	kelurahanRepository repository.KelurahanRepository
}

func NewKelurahanService(kelurahanRepository repository.KelurahanRepository) *kelurahanService {
	return &kelurahanService{kelurahanRepository}
}

func (s *kelurahanService) FindByKecamatanId(kecamatanId string) ([]model.Kelurahan, error) {
	var kelurahans, err = s.kelurahanRepository.FindByKecamatanId(kecamatanId)

	return kelurahans, err
}

func (s *kelurahanService) FindById(id string) (model.Kelurahan, error) {
	var kelurahan, err = s.kelurahanRepository.FindById(id)

	return kelurahan, err
}
