package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type MahasiswaService interface {
	FindAll() ([]model.Mahasiswa, error)
	FindById(id int) (model.Mahasiswa, error)
	FindAllRelation() ([]model.Mahasiswa, error)
	Create(mahasiswaRequest request.CreateMahasiswaRequest) (model.Mahasiswa, error)
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

func (s *mahasiswaService) FindAllRelation() ([]model.Mahasiswa, error) {
	var mahasiswas, err = s.mahasiswaRepository.FindAllRelation()

	return mahasiswas, err
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
	}

	newMahasiswa, err := s.mahasiswaRepository.Create(mahasiswa)

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

	updatedMahasiswa, err := s.mahasiswaRepository.Update(mahasiswa)

	return updatedMahasiswa, err
}

func (s *mahasiswaService) Delete(id int) (model.Mahasiswa, error) {
	var mahasiswa, err = s.mahasiswaRepository.FindById(id)

	deletedMahasiswa, err := s.mahasiswaRepository.Delete(mahasiswa)

	return deletedMahasiswa, err
}
