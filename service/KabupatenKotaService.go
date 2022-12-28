package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KabupatenKotaService interface {
	FindAll() ([]model.Kabupaten_Kota, error)
	FindById(id string) (model.Kabupaten_Kota, error)
}

type kabupatenKotaService struct {
	kabupatenKotaRepository repository.KabupatenKotaRepository
}

func NewKabupatenKotaService(kabupatenKotaRepository repository.KabupatenKotaRepository) *kabupatenKotaService {
	return &kabupatenKotaService{kabupatenKotaRepository}
}

func (s *kabupatenKotaService) FindAll() ([]model.Kabupaten_Kota, error) {
	var kabupatenKotas, err = s.kabupatenKotaRepository.FindAll()

	return kabupatenKotas, err
}

func (s *kabupatenKotaService) FindById(id string) (model.Kabupaten_Kota, error) {
	var kabupatenKota, err = s.kabupatenKotaRepository.FindById(id)

	return kabupatenKota, err
}
