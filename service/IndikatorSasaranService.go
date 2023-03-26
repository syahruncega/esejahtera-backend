package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndikatorSasaranService interface {
	FindAll() ([]model.IndikatorSasaran, error)
	FindById(id int) (model.IndikatorSasaran, error)
	FindByProgramId(programId int) ([]model.IndikatorSasaran, error)
	Create(indikatorSasaranRequest request.CreateIndikatorSasaranRequest) (model.IndikatorSasaran, error)
	Update(id int, indikatorSasaranRequest request.UpdateIndikatorSasaranRequest) (model.IndikatorSasaran, error)
	Delete(id int) (model.IndikatorSasaran, error)
}

type indikatorSasaranService struct {
	indikatorSasaranRepository repository.IndikatorSasaranRepository
}

func NewIndikatorSasaranService(indikatorSasaranRepository repository.IndikatorSasaranRepository) *indikatorSasaranService {
	return &indikatorSasaranService{indikatorSasaranRepository}
}

func (s *indikatorSasaranService) FindAll() ([]model.IndikatorSasaran, error) {
	var indikatorSasarans, err = s.indikatorSasaranRepository.FindAll()

	return indikatorSasarans, err
}

func (s *indikatorSasaranService) FindById(id int) (model.IndikatorSasaran, error) {
	var indikatorSasaran, err = s.indikatorSasaranRepository.FindById(id)

	return indikatorSasaran, err
}

func (s *indikatorSasaranService) FindByProgramId(programId int) ([]model.IndikatorSasaran, error) {
	var indikatorSasarans, err = s.indikatorSasaranRepository.FindByProgramId(programId)

	return indikatorSasarans, err
}

func (s *indikatorSasaranService) Create(indikatorSasaranRequest request.CreateIndikatorSasaranRequest) (model.IndikatorSasaran, error) {
	var indikatorSasaran = model.IndikatorSasaran{
		ProgramId:            indikatorSasaranRequest.ProgramId,
		NamaIndikatorSasaran: indikatorSasaranRequest.NamaIndikatorSasaran,
	}

	newIndikatorSasaran, err := s.indikatorSasaranRepository.Create(indikatorSasaran)

	return newIndikatorSasaran, err
}

func (s *indikatorSasaranService) Update(id int, indikatorSasaranRequest request.UpdateIndikatorSasaranRequest) (model.IndikatorSasaran, error) {
	var indikatorSasaran, err = s.indikatorSasaranRepository.FindById(id)

	indikatorSasaran.ProgramId = indikatorSasaranRequest.ProgramId
	indikatorSasaran.NamaIndikatorSasaran = indikatorSasaranRequest.NamaIndikatorSasaran

	updatedIndikatorSasaran, err := s.indikatorSasaranRepository.Update(indikatorSasaran)

	return updatedIndikatorSasaran, err
}

func (s *indikatorSasaranService) Delete(id int) (model.IndikatorSasaran, error) {
	var indikatorSasaran, err = s.indikatorSasaranRepository.FindById(id)

	deletedIndikatorSasaran, err := s.indikatorSasaranRepository.Delete(indikatorSasaran)

	return deletedIndikatorSasaran, err
}
