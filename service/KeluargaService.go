package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type KeluargaService interface {
	FindAll() ([]model.Keluarga, error)
	FindById(id int) (model.Keluarga, error)
	FindByIdKeluarga(idKeluarga string) (model.Keluarga, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Keluarga, error)
	CountPenerimaByKabupatenKota(kabupatenKotaId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKecamatan(kecamatanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountPenerimaByKelurahan(kelurahanId string, penerimaParameter string, nilai string) (jumlah int64, err error)
	CountDesilByKabupatenKota(kabupatenKotaId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKecamatan(kecamatanId string, nilaiDesil string) (jumlah int64, err error)
	CountDesilByKelurahan(kelurahanId string, nilaiDesil string) (jumlah int64, err error)
	Update(id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error)
}

type keluargaService struct {
	keluargaRepository repository.KeluargaRepository
}

func NewKeluargaService(keluargaRepository repository.KeluargaRepository) *keluargaService {
	return &keluargaService{keluargaRepository}
}

func (s *keluargaService) FindAll() ([]model.Keluarga, error) {
	var keluargas, err = s.keluargaRepository.FindAll()

	return keluargas, err
}

func (s *keluargaService) FindById(id int) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindById(id)

	return keluarga, err
}

func (s *keluargaService) FindByIdKeluarga(idKeluarga string) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindByIdKeluarga(idKeluarga)

	return keluarga, err
}

func (s *keluargaService) FindBySearch(whereClause map[string]interface{}) ([]model.Keluarga, error) {
	var keluargas, err = s.keluargaRepository.FindBySearch(whereClause)

	return keluargas, err
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

func (s *keluargaService) Update(id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindById(id)

	keluarga.UserId = keluargaRequest.UserId
	keluarga.StatusVerifikasi = keluargaRequest.StatusVerifikasi
	keluarga.MahasiswaId = keluargaRequest.MahasiswaId

	updatedKeluarga, err := s.keluargaRepository.Update(keluarga)

	return updatedKeluarga, err
}
