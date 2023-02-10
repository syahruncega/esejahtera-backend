package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndividuService interface {
	FindAll() ([]model.Individu, error)
	FindById(id int) (model.Individu, error)
	FindByIdKeluarga(idKeluarga string) ([]model.Individu, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error)
	Update(id int, individuReqeust request.UpdateIndividuRequest) (model.Individu, error)
	CountJumlahIndividu(places string, placesId string) (int64, error)
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

type individuService struct {
	individuRepository repository.IndividuRepository
}

func NewIndividuService(individuRepository repository.IndividuRepository) *individuService {
	return &individuService{individuRepository}
}

func (s *individuService) FindAll() ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindAll()

	return individus, err
}

func (s *individuService) FindById(id int) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(id)

	return individu, err
}

func (s *individuService) FindByIdKeluarga(idKeluarga string) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindByIdKeluarga(idKeluarga)

	return individus, err
}

func (s *individuService) FindBySearch(whereClause map[string]interface{}) ([]model.Individu, error) {
	var individus, err = s.individuRepository.FindBySearch(whereClause)

	return individus, err
}

func (s *individuService) Update(id int, individuRequest request.UpdateIndividuRequest) (model.Individu, error) {
	var individu, err = s.individuRepository.FindById(id)

	individu.UserId = individuRequest.UserId
	individu.MahasiswaId = individuRequest.MahasiswaId
	individu.StatusVerifikasi = individuRequest.StatusVerifikasi

	updatedIndividu, err := s.individuRepository.Update(individu)

	return updatedIndividu, err
}

func (s *individuService) CountJumlahIndividu(places string, placesId string) (int64, error) {
	var jumlah, err = s.individuRepository.CountJumlahIndividu(places, placesId)

	return jumlah, err
}

func (s *individuService) CountJumlahDesil(places string, placesId string, desil string) (int64, error) {
	var jumlah, err = s.individuRepository.CountJumlahDesil(places, placesId, desil)

	return jumlah, err
}

func (s *individuService) CountJumlahPenerima(places string, placesId string, penerima string, penerimaValue string) (int64, error) {
	var jumlah, err = s.individuRepository.CountJumlahPenerima(places, placesId, penerima, penerimaValue)

	return jumlah, err
}

func (s *individuService) DistinctKabupatenKota() ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota, err = s.individuRepository.DistinctKabupatenKota()

	return distinctKabupatenKota, err
}

func (s *individuService) DistinctCountKabupatenKota() (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKabupatenKota()

	return distinctCount, err
}

func (s *individuService) DistinctKecamatan() ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.individuRepository.DistinctKecamatan()

	return distinctKecamatan, err
}

func (s *individuService) DistinctKecamatanByKabupatenKota(kabupatenKotaId string) ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.individuRepository.DistinctKecamatanByKabupatenKota(kabupatenKotaId)

	return distinctKecamatan, err
}

func (s *individuService) DistinctCountKecamatan(kabupatenKotaId string) (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKecamatan(kabupatenKotaId)

	return distinctCount, err
}

func (s *individuService) DistinctKelurahanByKecamatan(kecamatanId string) ([]model.DistinctKelurahan, error) {
	var distinctKelurahan, err = s.individuRepository.DistinctKelurahanByKecamatan(kecamatanId)

	return distinctKelurahan, err
}

func (s *individuService) DistinctCountKelurahan(kecamatanId string) (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKelurahan(kecamatanId)

	return distinctCount, err
}
