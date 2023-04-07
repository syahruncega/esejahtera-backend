package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type SasaranService interface {
	FindAll() ([]model.Sasaran, error)
	FindById(id int) (model.Sasaran, error)
	FindByProgramId(programId int) ([]model.Sasaran, error)
	Create(sasaranRequest request.CreateSasaranRequest) (model.Sasaran, error)
	Update(id int, sasaranRequest request.UpdateSasaranRequest) (model.Sasaran, error)
	Delete(id int) (model.Sasaran, error)
}

type sasaranService struct {
	sasaranRepository repository.SasaranRepository
}

func NewSasaranService(sasaranRepository repository.SasaranRepository) *sasaranService {
	return &sasaranService{sasaranRepository}
}

func (s *sasaranService) FindAll() ([]model.Sasaran, error) {
	var sasarans, err = s.sasaranRepository.FindAll()

	return sasarans, err
}

func (s *sasaranService) FindById(id int) (model.Sasaran, error) {
	var sasaran, err = s.sasaranRepository.FindById(id)

	return sasaran, err
}

func (s *sasaranService) FindByProgramId(programId int) ([]model.Sasaran, error) {
	var sasarans, err = s.sasaranRepository.FindByProgramId(programId)

	return sasarans, err
}

func (s *sasaranService) Create(sasaranRequest request.CreateSasaranRequest) (model.Sasaran, error) {
	var sasaran = model.Sasaran{
		RencanaProgramId: sasaranRequest.RencanaProgramId,
		NamaSasaran:      sasaranRequest.NamaSasaran,
	}

	newSasaran, err := s.sasaranRepository.Create(sasaran)

	return newSasaran, err
}

func (s *sasaranService) Update(id int, sasaranRequest request.UpdateSasaranRequest) (model.Sasaran, error) {
	var sasaran, err = s.sasaranRepository.FindById(id)

	sasaran.RencanaProgramId = sasaranRequest.RencanaProgramId
	sasaran.NamaSasaran = sasaranRequest.NamaSasaran

	updatedSasaran, err := s.sasaranRepository.Update(sasaran)

	return updatedSasaran, err
}

func (s *sasaranService) Delete(id int) (model.Sasaran, error) {
	var sasaran, err = s.sasaranRepository.FindById(id)

	deletedSasaran, err := s.sasaranRepository.Delete(sasaran)

	return deletedSasaran, err
}
