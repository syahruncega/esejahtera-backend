package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type MahasiswaService interface {
	FindAll() ([]model.Mahasiswa, error)
	FindById(id int) (model.Mahasiswa, error)
	FindByUserId(userId int) (model.Mahasiswa, error)
	FindAllRelation() ([]model.Mahasiswa, error)
	CountVerifiedIndividu(id int, kelurahanId string) (int64, error)
	CountVerifiedKeluarga(id int, kelurahanId string) (int64, error)
	DistinctKelurahan() ([]model.DistinctKelurahan, error)
	Create(mahasiswaRequest request.CreateMahasiswaRequest) (model.Mahasiswa, error)
	CreateBatch(mahasiswaRequest []request.CreateMahasiswaRequest) ([]model.Mahasiswa, error)
	Update(id int, mahasiswaRequest request.UpdateMahasiswaRequest) (model.Mahasiswa, error)
	Delete(id int) (model.Mahasiswa, error)
}

type mahasiswaService struct {
	mahasiswaRepository repository.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepository repository.MahasiswaRepository) *mahasiswaService {
	return &mahasiswaService{mahasiswaRepository}
}

func (s *mahasiswaService) FindAll() ([]model.Mahasiswa, error) {
	var mahasiswas, err = s.mahasiswaRepository.FindAll()

	return mahasiswas, err
}

func (s *mahasiswaService) FindById(id int) (model.Mahasiswa, error) {
	var mahasiswa, err = s.mahasiswaRepository.FindById(id)

	return mahasiswa, err
}

func (s *mahasiswaService) FindByUserId(userId int) (model.Mahasiswa, error) {
	var mahasiswa, err = s.mahasiswaRepository.FindByUserId(userId)

	return mahasiswa, err
}

func (s *mahasiswaService) FindAllRelation() ([]model.Mahasiswa, error) {
	var mahasiswas, err = s.mahasiswaRepository.FindAllRelation()

	return mahasiswas, err
}

func (s *mahasiswaService) CountVerifiedIndividu(id int, kelurahanId string) (int64, error) {
	var jumlah, err = s.mahasiswaRepository.CountVerifiedIndividu(id, kelurahanId)

	return jumlah, err
}

func (s *mahasiswaService) CountVerifiedKeluarga(id int, kelurahanId string) (int64, error) {
	var jumlah, err = s.mahasiswaRepository.CountVerifiedKeluarga(id, kelurahanId)

	return jumlah, err
}

func (s *mahasiswaService) DistinctKelurahan() ([]model.DistinctKelurahan, error) {
	var distinctKelurahan, err = s.mahasiswaRepository.DistinctKelurahan()

	return distinctKelurahan, err
}

func (s *mahasiswaService) Create(mahasiswaRequest request.CreateMahasiswaRequest) (model.Mahasiswa, error) {
	var mahasiswa = model.Mahasiswa{
		UserId:          mahasiswaRequest.UserId,
		NamaLengkap:     mahasiswaRequest.NamaLengkap,
		Universitas:     mahasiswaRequest.Universitas,
		JenisKelamin:    mahasiswaRequest.JenisKelamin,
		TanggalLahir:    mahasiswaRequest.TanggalLahir,
		KabupatenKotaId: mahasiswaRequest.KabupatenKotaId,
		KecamatanId:     mahasiswaRequest.KecamatanId,
		KelurahanId:     mahasiswaRequest.KelurahanId,
		UrlFotoProfil:   mahasiswaRequest.UrlFotoProfil,
	}

	newMahasiswa, err := s.mahasiswaRepository.Create(mahasiswa)

	return newMahasiswa, err
}

func (s *mahasiswaService) CreateBatch(mahasiswasRequest []request.CreateMahasiswaRequest) ([]model.Mahasiswa, error) {
	var dataMahasiswas []model.Mahasiswa

	for _, mahasiswa := range mahasiswasRequest {
		var data = model.Mahasiswa{
			UserId:          mahasiswa.UserId,
			NamaLengkap:     mahasiswa.NamaLengkap,
			Universitas:     mahasiswa.Universitas,
			JenisKelamin:    mahasiswa.JenisKelamin,
			TanggalLahir:    mahasiswa.TanggalLahir,
			KabupatenKotaId: mahasiswa.KabupatenKotaId,
			KecamatanId:     mahasiswa.KecamatanId,
			KelurahanId:     mahasiswa.KelurahanId,
			UrlFotoProfil:   mahasiswa.UrlFotoProfil,
		}

		dataMahasiswas = append(dataMahasiswas, data)
	}

	newMahasiswa, err := s.mahasiswaRepository.CreateBatch(dataMahasiswas)

	return newMahasiswa, err
}

func (s *mahasiswaService) Update(id int, mahasiswaRequest request.UpdateMahasiswaRequest) (model.Mahasiswa, error) {
	var mahasiswa, err = s.mahasiswaRepository.FindById(id)

	mahasiswa.UserId = mahasiswaRequest.UserId
	mahasiswa.NamaLengkap = mahasiswaRequest.NamaLengkap
	mahasiswa.Universitas = mahasiswaRequest.Universitas
	mahasiswa.JenisKelamin = mahasiswaRequest.JenisKelamin
	mahasiswa.TanggalLahir = mahasiswaRequest.TanggalLahir
	mahasiswa.KabupatenKotaId = mahasiswaRequest.KabupatenKotaId
	mahasiswa.KecamatanId = mahasiswaRequest.KecamatanId
	mahasiswa.KelurahanId = mahasiswaRequest.KelurahanId
	mahasiswa.UrlFotoProfil = mahasiswaRequest.UrlFotoProfil

	updatedMahasiswa, err := s.mahasiswaRepository.Update(mahasiswa)

	return updatedMahasiswa, err
}

func (s *mahasiswaService) Delete(id int) (model.Mahasiswa, error) {
	var mahasiswa, err = s.mahasiswaRepository.FindById(id)

	deletedMahasiswa, err := s.mahasiswaRepository.Delete(mahasiswa)

	return deletedMahasiswa, err
}
