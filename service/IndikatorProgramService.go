package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type IndikatorProgramService interface {
	FindAll() ([]model.IndikatorProgram, error)
	FindById(id int) (model.IndikatorProgram, error)
	Create(indikatorProgramRequest request.CreateIndikatorProgramRequest) (model.IndikatorProgram, error)
	Update(id int, indikatorProgramRequest request.UpdateIndikatorProgramRequest) (model.IndikatorProgram, error)
	Delete(id int) (model.IndikatorProgram, error)
}

type indikatorProgramService struct {
	indikatorProgramRepository repository.IndikatorProgramRepository
}

func NewIndikatorProgramService(indikatorProgramRepository repository.IndikatorProgramRepository) *indikatorProgramService {
	return &indikatorProgramService{indikatorProgramRepository}
}

func (s *indikatorProgramService) FindAll() ([]model.IndikatorProgram, error) {
	var indikatorPrograms, err = s.indikatorProgramRepository.FindAll()

	return indikatorPrograms, err
}

func (s *indikatorProgramService) FindById(id int) (model.IndikatorProgram, error) {
	var indikatorProgram, err = s.indikatorProgramRepository.FindById(id)

	return indikatorProgram, err
}

func (s *indikatorProgramService) Create(indikatorProgramRequest request.CreateIndikatorProgramRequest) (model.IndikatorProgram, error) {
	var indikatorProgram = model.IndikatorProgram{
		InstansiId:              indikatorProgramRequest.InstansiId,
		ProgramId:               indikatorProgramRequest.ProgramId,
		Sasaran:                 indikatorProgramRequest.Sasaran,
		IndikatorSasaran:        indikatorProgramRequest.IndikatorSasaran,
		Kebijakan:               indikatorProgramRequest.Kebijakan,
		IndikatorKinerjaProgram: indikatorProgramRequest.IndikatorKinerjaProgram,
		PaguProgram:             indikatorProgramRequest.PaguProgram,
	}

	var newIndikatorProgram, err = s.indikatorProgramRepository.Create(indikatorProgram)

	return newIndikatorProgram, err
}

func (s *indikatorProgramService) Update(id int, indikatorProgramRequest request.UpdateIndikatorProgramRequest) (model.IndikatorProgram, error) {
	var indikatorProgram, err = s.indikatorProgramRepository.FindById(id)

	indikatorProgram.InstansiId = indikatorProgramRequest.InstansiId
	indikatorProgram.ProgramId = indikatorProgramRequest.ProgramId
	indikatorProgram.Sasaran = indikatorProgramRequest.Sasaran
	indikatorProgram.IndikatorSasaran = indikatorProgramRequest.IndikatorSasaran
	indikatorProgram.Kebijakan = indikatorProgramRequest.Kebijakan
	indikatorProgram.IndikatorKinerjaProgram = indikatorProgramRequest.IndikatorKinerjaProgram
	indikatorProgram.PaguProgram = indikatorProgramRequest.PaguProgram

	updatedIndikatorProgram, err := s.indikatorProgramRepository.Update(indikatorProgram)

	return updatedIndikatorProgram, err
}

func (s *indikatorProgramService) Delete(id int) (model.IndikatorProgram, error) {
	var indikatorProgram, err = s.indikatorProgramRepository.FindById(id)

	deletedIndikatorProgram, err := s.indikatorProgramRepository.Delete(indikatorProgram)

	return deletedIndikatorProgram, err
}
