package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
)

type KeluargaService interface {
	FindAll(kabupatenKotaId string) ([]model.Keluarga, error)
	FindById(kabupatenKotaId string, id int) (model.Keluarga, error)
	FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error)
	FindPenerimaByKelurahan(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error)
}

type keluargaService struct {
	keluargaRepository repository.KeluargaRepository
}

func NewKeluargaService(keluargaRepository repository.KeluargaRepository) *keluargaService {
	return &keluargaService{keluargaRepository}
}

func (s *keluargaService) FindAll(kabupatenKotaId string) ([]model.Keluarga, error) {
	var keluargas, err = s.keluargaRepository.FindAll(kabupatenKotaId)

	return keluargas, err
}

func (s *keluargaService) FindById(kabupatenKotaId string, id int) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindById(kabupatenKotaId, id)

	return keluarga, err
}

func (s *keluargaService) FindByIdKeluarga(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindByIdKeluarga(kabupatenKotaId, idKeluarga)

	return keluarga, err
}

func (s *keluargaService) FindPenerimaByKelurahan(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.FindPenerimaByKelurahan(kabupatenKotaId, penerimaParameter, nilai)

	return count, errr
}
