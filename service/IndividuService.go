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
	CountVerifiedByMahasiswa(mahasiswaId int) (int64, error)
	DistinctKabupatenKota(kabupatenKotaId string) ([]model.DistinctKabupatenKota, error)
	DistinctCountKabupatenKota(kabupatenKotaId string) (map[string]int64, error)
	DistinctKecamatan(kecamatanId string) ([]model.DistinctKecamatan, error)
	DistinctCountKecamatan(kecamatanId string) (map[string]int64, error)
	DistinctKelurahan(kelurahanId string) ([]model.DistinctKelurahan, error)
	DistinctCountKelurahan(kelurahanId string) (map[string]int64, error)
	CountJumlahIndividuByIdKeluarga(places string, placesId string, idKeluarga string) (int64, int64, error, error)
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

func (s *individuService) CountVerifiedByMahasiswa(mahasiswaId int) (int64, error) {
	var jumlah, err = s.individuRepository.CountVerifiedByMahasiswa(mahasiswaId)

	return jumlah, err
}

func (s *individuService) DistinctKabupatenKota(kabupatenKotaId string) ([]model.DistinctKabupatenKota, error) {
	var distinctKabupatenKota, err = s.individuRepository.DistinctKabupatenKota(kabupatenKotaId)

	return distinctKabupatenKota, err
}

func (s *individuService) DistinctCountKabupatenKota(kabupatenKotaId string) (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKabupatenKota(kabupatenKotaId)

	return distinctCount, err
}

func (s *individuService) DistinctKecamatan(kecamatanId string) ([]model.DistinctKecamatan, error) {
	var distinctKecamatan, err = s.individuRepository.DistinctKecamatan(kecamatanId)

	return distinctKecamatan, err
}

func (s *individuService) DistinctCountKecamatan(kecamatanId string) (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKecamatan(kecamatanId)

	return distinctCount, err
}

func (s *individuService) DistinctKelurahan(kelurahanId string) ([]model.DistinctKelurahan, error) {
	var distinctKelurahan, err = s.individuRepository.DistinctKelurahan(kelurahanId)

	return distinctKelurahan, err
}

func (s *individuService) DistinctCountKelurahan(kelurahanId string) (map[string]int64, error) {
	var distinctCount, err = s.individuRepository.DistinctCountKelurahan(kelurahanId)

	return distinctCount, err
}

func (s *individuService) CountJumlahIndividuByIdKeluarga(places string, placesId string, idKeluarga string) (int64, int64, error, error) {
	var jumlahIndividu, jumlahIndividuVerified, err1, err2 = s.individuRepository.CountJumlahIndividuByIdKeluarga(places, placesId, idKeluarga)

	return jumlahIndividu, jumlahIndividuVerified, err1, err2
}
