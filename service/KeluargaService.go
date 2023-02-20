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
	CountVerified(places string, placesId string, status int) (int64, error)
	CountVerifiedByMahasiswa(mahasiswaId int) (int64, error)
	DistinctKabupatenKota(kabupatenKotaId string) ([]model.DistinctKabupatenKota, error)
	DistinctCountKabupatenKota(kabupatenKotaId string) (map[string]int64, error)
	DistinctKecamatan(kecamatanId string) ([]model.DistinctKecamatan, error)
	DistinctCountKecamatan(kecamatanId string) (map[string]int64, error)
	DistinctKelurahan(kelurahanId string) ([]model.DistinctKelurahan, error)
	DistinctCountKelurahan(kelurahanId string) (map[string]int64, error)
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

func (s *keluargaService) CountVerified(places string, placesId string, statusVerified int) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountVerified(places, placesId, statusVerified)

	return jumlah, err
}

func (s *keluargaService) CountVerifiedByMahasiswa(mahasiswaId int) (int64, error) {
	var jumlah, err = s.keluargaRepository.CountVerifiedByMahasiswa(mahasiswaId)

	return jumlah, err
}

func (s *keluargaService) DistinctKabupatenKota(kabupatenKotaId string) ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota, err = s.keluargaRepository.DistinctKabupatenKota(kabupatenKotaId)

	return distinctKabupatenKota, err
}

func (s *keluargaService) DistinctCountKabupatenKota(kabupatenKotaId string) (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKabupatenKota(kabupatenKotaId)

	return distinctCount, err
}

func (s *keluargaService) DistinctKecamatan(kecamatanId string) ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.keluargaRepository.DistinctKecamatan(kecamatanId)

	return distinctKecamatan, err
}

func (s *keluargaService) DistinctCountKecamatan(kecamatanId string) (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKecamatan(kecamatanId)

	return distinctCount, err
}

func (s *keluargaService) DistinctKelurahan(kelurahanId string) ([]model.DistinctKelurahan, error) {
	var distinctKelurahan, err = s.keluargaRepository.DistinctKelurahan(kelurahanId)

	return distinctKelurahan, err
}

func (s *keluargaService) DistinctCountKelurahan(kelurahanId string) (map[string]int64, error) {
	var distinctCount, err = s.keluargaRepository.DistinctCountKelurahan(kelurahanId)

	return distinctCount, err
}
