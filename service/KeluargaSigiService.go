package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KeluargaSigiService interface {
	FindAll() ([]model.KeluargaSigi, error)
	FindById(id int) (model.KeluargaSigi, error)
	FindByIdKeluarga(idKeluarga string) (model.KeluargaSigi, error)
}

type keluargaSigiService struct {
	keluargaRepository repository.KeluargaSigiRepository
}

func NewKeluargaSigiService(keluargaRepository repository.KeluargaSigiRepository) *keluargaSigiService {
	return &keluargaSigiService{keluargaRepository}
}

func (s *keluargaSigiService) FindAll() ([]model.KeluargaSigi, error) {
	var keluargaSigis, err = s.keluargaRepository.FindAll()

	return keluargaSigis, err
}

func (s *keluargaSigiService) FindById(id int) (model.KeluargaSigi, error) {
	var keluargaSigi, err = s.keluargaRepository.FindById(id)

	return keluargaSigi, err
}

func (s *keluargaSigiService) FindByIdKeluarga(idKeluarga string) (model.KeluargaSigi, error) {
	var keluargaSigi, err = s.keluargaRepository.FindByIdKeluarga(idKeluarga)

	return keluargaSigi, err
}
