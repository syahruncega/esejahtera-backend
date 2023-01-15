package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KeluargaDonggalaService interface {
	FindAll() ([]model.KeluargaDonggala, error)
	FindById(id int) (model.KeluargaDonggala, error)
	FindByIdKeluarga(idKeluarga string) (model.KeluargaDonggala, error)
}

type keluargaDonggalaService struct {
	keluargaDonggalaRepository repository.KeluargaDonggalaRepository
}

func NewKeluargaDonggalaService(keluargaDonggalaRepository repository.KeluargaDonggalaRepository) *keluargaDonggalaService {
	return &keluargaDonggalaService{keluargaDonggalaRepository}
}

func (s *keluargaDonggalaService) FindAll() ([]model.KeluargaDonggala, error) {
	var keluargaDonggalas, err = s.keluargaDonggalaRepository.FindAll()

	return keluargaDonggalas, err
}

func (s *keluargaDonggalaService) FindById(id int) (model.KeluargaDonggala, error) {
	var keluargaDonggala, err = s.keluargaDonggalaRepository.FindById(id)

	return keluargaDonggala, err
}

func (s *keluargaDonggalaService) FindByIdKeluarga(idKeluarga string) (model.KeluargaDonggala, error) {
	var keluargaDonggala, err = s.keluargaDonggalaRepository.FindByIdKeluarga(idKeluarga)

	return keluargaDonggala, err
}
