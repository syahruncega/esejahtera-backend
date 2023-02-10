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
	CountData(whereClause map[string]interface{}) (int64, error)
	FindBySearch(whereClause map[string]interface{}, limit int, offset int) ([]model.Keluarga, error)
	Update(id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error)
	CountJumlahKeluarga(places string, placesId string) (int64, error)
	CountJumlahDesil(places string, placesId string, desil string) (int64, error)
	CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error)
	DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error)
	DistinctCountKabupatenKota() (map[string]int64, error)
	DistinctKecamatan() ([]model.DistinctKecamatan, error)
	DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error)
	DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error)
	DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error)
	DistinctCountKelurahan(kecamatanId string) (map[string]int64, error)
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

func (s *keluargaService) CountData(whereClause map[string]interface{}) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountData(whereClause)

	return jumlah, err
}

func (s *keluargaService) FindBySearch(whereClause map[string]interface{}, limit int, offset int) ([]model.Keluarga, error) {
	var keluargas, err = s.keluargaRepository.FindBySearch(whereClause, limit, offset)

	return keluargas, err
}

func (s *keluargaService) Update(id int, keluargaRequest request.UpdateKeluargaRequest) (model.Keluarga, error) {
	var keluarga, err = s.keluargaRepository.FindById(id)

	keluarga.UserId = keluargaRequest.UserId
	keluarga.StatusVerifikasi = keluargaRequest.StatusVerifikasi
	keluarga.MahasiswaId = keluargaRequest.MahasiswaId

	updatedKeluarga, err := s.keluargaRepository.Update(keluarga)

	return updatedKeluarga, err
}

func (s *keluargaService) CountJumlahKeluarga(places string, placesId string) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountJumlahKeluarga(places, placesId)

	return jumlah, err
}

func (s *keluargaService) CountJumlahDesil(places string, placesId string, desil string) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountJumlahDesil(places, placesId, desil)

	return jumlah, err
}

func (s *keluargaService) CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountJumlahPenerima(places, placesId, penerima, penerimaValue)

	return jumlah, err
}

func (s *keluargaService) DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota, err = s.keluargaRepository.DistinctKabupatenKota()

	return distinctKabupatenKota, err
}

func (s *keluargaService) DistinctCountKabupatenKota() (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKabupatenKota()

	return distinctCount, err
}

func (s *keluargaService) DistinctKecamatan() ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.keluargaRepository.DistinctKecamatan()

	return distinctKecamatan, err
}

func (s *keluargaService) DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.keluargaRepository.DistinctKecamatanByKabupatenKota(kabupatenKotaId)

	return distinctKecamatan, err
}

func (s *keluargaService) DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKecamatan(kabupatenKotaId)

	return distinctCount, err
}

func (s *keluargaService) DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error) {
	var distinctKelurahan, err = s.keluargaRepository.DistinctKelurahanByKecamatan(kecamatanId)

	return distinctKelurahan, err
}

func (s *keluargaService) DistinctCountKelurahan(kecamatanId string) (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKelurahan(kecamatanId)

	return distinctCount, err
}
