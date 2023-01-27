package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type KeluargaService interface {
	FindAll(kabupatenKotaId string) ([]model.Keluarga, error)
	FindById(kabupatenKotaId string, id int) (model.Keluarga, error)
	FindByIdKeluargaByKabupatenKota(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error)
	CountPenerimaByKabupatenKota(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKecamatan(kecamatanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountDesilByKabupatenKota(kabupatenKotaId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKecamatan(kecamatanId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKelurahan(kelurahanId string, nilaiDesil string) (jumlah int64, err error)
	Update(kabupatenKotaId string, id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error)
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

func (s *keluargaService) FindByIdKeluargaByKabupatenKota(kabupatenKotaId string, idKeluarga string) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindByIdKeluargaByKabupatenKota(kabupatenKotaId, idKeluarga)

	return keluarga, err
}

func (s *keluargaService) CountPenerimaByKabupatenKota(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountPenerimaByKabupatenKota(kabupatenKotaId, penerimaParameter, nilai)

	return count, errr
}

func (s *keluargaService) CountPenerimaByKecamatan(kecamatanId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountPenerimaByKecamatan(kecamatanId, penerimaParameter, nilai)

	return count, errr
}

func (s *keluargaService) CountPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountPenerimaByKelurahan(kelurahanId, penerimaParameter, nilai)

	return count, errr
}

func (s *keluargaService) CountDesilByKabupatenKota(kabupatenKotaId string, nilaiDesil string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountDesilByKabupatenKota(kabupatenKotaId, nilaiDesil)

	return count, errr
}

func (s *keluargaService) CountDesilByKecamatan(kecamatanId string, nilaiDesil string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountDesilByKecamatan(kecamatanId, nilaiDesil)

	return count, errr
}

func (s *keluargaService) CountDesilByKelurahan(kelurahanId string, nilaiDesil string) (jumlah int64, err error) {
	var count, errr = s.keluargaRepository.CountDesilByKelurahan(kelurahanId, nilaiDesil)

	return count, errr
}

func (s *keluargaService) Update(kabupatenKotaId string, id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindById(kabupatenKotaId, id)

	keluarga.UserId = keluargaRequest.UserId
	keluarga.StatusVerifikasi = keluargaRequest.StatusVerifikasi
	keluarga.MahasiswaId = keluargaRequest.MahasiswaId

	updatedKeluarga, err := s.keluargaRepository.Update(keluarga)

	return updatedKeluarga, err
}
