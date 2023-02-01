package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndividuVerifikasiService interface {
	FindAll() ([]model.IndividuVerifikasi, error)
	FindById(id int) (model.IndividuVerifikasi, error)
	FindByIdKeluarga(idKeluarga string) ([]model.IndividuVerifikasi, error)
	FindBySearch(whereClause map[string]interface{}) ([]model.IndividuVerifikasi, error)
	Create(individuVerifikasiRequest request.CreateIndividuVerifikasiRequest) (model.IndividuVerifikasi, error)
	Update(id int, individuVerifikasiRequest request.UpdateIndividuVerifikasiRequest) (model.IndividuVerifikasi, error)
	Delete(id int) (model.IndividuVerifikasi, error)
}

type individuVerifikasiService struct {
	individuVerifikasiRepository repository.IndividuVerifikasiRepository
}

func NewIndividuVerifikasiService(individuVerifikasiRepository repository.IndividuVerifikasiRepository) *individuVerifikasiService {
	return &individuVerifikasiService{individuVerifikasiRepository}
}

func (s *individuVerifikasiService) FindAll() ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis, err = s.individuVerifikasiRepository.FindAll()

	return individuVerifikasis, err
}

func (s *individuVerifikasiService) FindById(id int) (model.IndividuVerifikasi, error) {
	var individuVerifikasi, err = s.individuVerifikasiRepository.FindById(id)

	return individuVerifikasi, err
}

func (s *individuVerifikasiService) FindByIdKeluarga(idKeluarga string) ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis, err = s.individuVerifikasiRepository.FindByIdKeluarga(idKeluarga)

	return individuVerifikasis, err
}

func (s *individuVerifikasiService) FindBySearch(whereClause map[string]interface{}) ([]model.IndividuVerifikasi, error) {
	var individuVerifikasis, err = s.individuVerifikasiRepository.FindBySearch(whereClause)

	return individuVerifikasis, err
}

func (s *individuVerifikasiService) Create(individuVerifikasiRequest request.CreateIndividuVerifikasiRequest) (model.IndividuVerifikasi, error) {
	var individuVerifikasi = model.IndividuVerifikasi{
		Id:                 individuVerifikasiRequest.Id,
		IdKeluarga:         individuVerifikasiRequest.IdKeluarga,
		ProvinsiId:         individuVerifikasiRequest.ProvinsiId,
		KabupatenKotaId:    individuVerifikasiRequest.KabupatenKotaId,
		KecamatanId:        individuVerifikasiRequest.KecamatanId,
		KelurahanId:        individuVerifikasiRequest.KelurahanId,
		DesilKesejahteraan: individuVerifikasiRequest.DesilKesejahteraan,
		Alamat:             individuVerifikasiRequest.Alamat,
		IdIndividu:         individuVerifikasiRequest.IdIndividu,
		Nama:               individuVerifikasiRequest.Nama,
		Nik:                individuVerifikasiRequest.Nik,
		PadanDukcapil:      individuVerifikasiRequest.PadanDukcapil,
		JenisKelamin:       individuVerifikasiRequest.JenisKelamin,
		Hubungan:           individuVerifikasiRequest.Hubungan,
		TanggalLahir:       individuVerifikasiRequest.TanggalLahir,
		StatusKawin:        individuVerifikasiRequest.StatusKawin,
		Pekerjaan:          individuVerifikasiRequest.Pekerjaan,
		Pendidikan:         individuVerifikasiRequest.Pendidikan,
		KategoriUsia:       individuVerifikasiRequest.KategoriUsia,
		PenerimaBPNT:       individuVerifikasiRequest.PenerimaBPNT,
		PenerimaBST:        individuVerifikasiRequest.PenerimaBST,
		PenerimaPKH:        individuVerifikasiRequest.PenerimaPKH,
		PenerimaSembako:    individuVerifikasiRequest.PenerimaSembako,
		PenerimaLainnya:    individuVerifikasiRequest.PenerimaLainnya,
		StatusResponden:    individuVerifikasiRequest.StatusResponden,
		UserId:             individuVerifikasiRequest.UserId,
		MahasiswaId:        individuVerifikasiRequest.MahasiswaId,
	}

	newIndividuVerifikasi, err := s.individuVerifikasiRepository.Create(individuVerifikasi)

	return newIndividuVerifikasi, err
}

func (s *individuVerifikasiService) Update(id int, individuVerifikasiRequest request.UpdateIndividuVerifikasiRequest) (model.IndividuVerifikasi, error) {
	var individuVerifikasi, err = s.individuVerifikasiRepository.FindById(id)

	individuVerifikasi.Id = individuVerifikasiRequest.Id
	individuVerifikasi.IdKeluarga = individuVerifikasiRequest.IdKeluarga
	individuVerifikasi.ProvinsiId = individuVerifikasiRequest.ProvinsiId
	individuVerifikasi.KabupatenKotaId = individuVerifikasiRequest.KabupatenKotaId
	individuVerifikasi.KecamatanId = individuVerifikasiRequest.KecamatanId
	individuVerifikasi.KelurahanId = individuVerifikasiRequest.KelurahanId
	individuVerifikasi.DesilKesejahteraan = individuVerifikasiRequest.DesilKesejahteraan
	individuVerifikasi.Alamat = individuVerifikasiRequest.Alamat
	individuVerifikasi.IdIndividu = individuVerifikasiRequest.IdIndividu
	individuVerifikasi.Nama = individuVerifikasiRequest.Nama
	individuVerifikasi.Nik = individuVerifikasiRequest.Nik
	individuVerifikasi.PadanDukcapil = individuVerifikasiRequest.PadanDukcapil
	individuVerifikasi.JenisKelamin = individuVerifikasiRequest.JenisKelamin
	individuVerifikasi.Hubungan = individuVerifikasiRequest.Hubungan
	individuVerifikasi.TanggalLahir = individuVerifikasiRequest.TanggalLahir
	individuVerifikasi.StatusKawin = individuVerifikasiRequest.StatusKawin
	individuVerifikasi.Pekerjaan = individuVerifikasiRequest.Pekerjaan
	individuVerifikasi.Pendidikan = individuVerifikasiRequest.Pendidikan
	individuVerifikasi.KategoriUsia = individuVerifikasiRequest.KategoriUsia
	individuVerifikasi.PenerimaBPNT = individuVerifikasiRequest.PenerimaBPNT
	individuVerifikasi.PenerimaBPUM = individuVerifikasiRequest.PenerimaBPUM
	individuVerifikasi.PenerimaBST = individuVerifikasiRequest.PenerimaBST
	individuVerifikasi.PenerimaPKH = individuVerifikasiRequest.PenerimaPKH
	individuVerifikasi.PenerimaSembako = individuVerifikasiRequest.PenerimaSembako
	individuVerifikasi.PenerimaLainnya = individuVerifikasiRequest.PenerimaLainnya
	individuVerifikasi.StatusResponden = individuVerifikasiRequest.StatusResponden
	individuVerifikasi.UserId = individuVerifikasiRequest.UserId
	individuVerifikasi.MahasiswaId = individuVerifikasiRequest.MahasiswaId

	updatedIndividuVerifikasi, err := s.individuVerifikasiRepository.Update(individuVerifikasi)

	return updatedIndividuVerifikasi, err

}

func (s *individuVerifikasiService) Delete(id int) (model.IndividuVerifikasi, error) {
	var individuVerifikasi, err = s.individuVerifikasiRepository.FindById(id)

	deletedIndividuVerifikasi, err := s.individuVerifikasiRepository.Delete(individuVerifikasi)

	return deletedIndividuVerifikasi, err
}
