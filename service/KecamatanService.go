package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KecamatanService interface {
	FindByKabupatenKota(kabupatenKotaId string) ([]model.Kecamatan, error)
	FindById(id string) (model.Kecamatan, error)
}

type kecamatanService struct {
	kecamatanRepository repository.KecamatanRepository
}

func NewKecamatanService(kecamatanRepository repository.KecamatanRepository) *kecamatanService {
	return &kecamatanService{kecamatanRepository}
}

func (s *kecamatanService) FindByKabupatenKota(kabupatenKotaId string) ([]model.Kecamatan, error) {
	var kecamatans, err = s.kecamatanRepository.FindByKabupatenKota(kabupatenKotaId)

	return kecamatans, err
}

func (s *kecamatanService) FindById(id string) (model.Kecamatan, error) {
	var kecamatan, err = s.kecamatanRepository.FindById(id)

	return kecamatan, err
}
